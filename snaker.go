// Package snaker provides methods to convert CamelCase names to snake_case and back.
// It considers the list of allowed initialisms used by github.com/golang/lint/golint (e.g. ID or HTTP)
package snaker

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// S is a snaker with its own list of initialisms
type S struct {
	initialisms map[string]bool
	min, max    int
}

var def = NewDefault()

// New returns a new snaker without any initialisms
func New(initialism ...string) *S {
	s := &S{
		initialisms: make(map[string]bool),
		min:         1<<31 - 1,
	}
	s.Add(initialism...)
	return s
}

// NewDefault returns a new snaker with common initialisms
func NewDefault() *S {
	return New(commonInitialisms...)
}

// Add an initialism.
func (ss *S) Add(initialism ...string) *S {
	for _, i := range initialism {
		ss.initialisms[i] = true
		if len(i) > ss.max {
			ss.max = len(i)
		}
		if len(i) < ss.min {
			ss.min = len(i)
		}
	}
	return ss
}

// CamelToSnake converts a given string to snake case
func (ss *S) CamelToSnake(s string) string {
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := ss.startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}
	return strings.ToLower(strings.Join(words, "_"))
}

// startsWithInitialism returns the initialism if the given string begins with it
func (ss *S) startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := ss.min - 1; i <= ss.max; i++ {
		if len(s) > i-1 && ss.initialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// SnakeToCamel returns a string converted from snake case to uppercase
func (ss *S) SnakeToCamel(s string) string {
	return ss.snakeToCamel(s, true)
}

// SnakeToCamelLower returns a string converted from snake case to lowercase
func (ss *S) SnakeToCamelLower(s string) string {
	return ss.snakeToCamel(s, false)
}

func (ss *S) snakeToCamel(s string, upperCase bool) string {
	words := strings.Split(s, "_")

	for i, word := range words {
		if upperCase || i > 0 {
			if upper := strings.ToUpper(word); ss.initialisms[upper] {
				words[i] = upper
				continue
			}
		}

		if (upperCase || i > 0) && len(word) > 0 {
			r, s := utf8.DecodeRuneInString(word)
			words[i] = string(unicode.ToUpper(r)) + word[s:]
		}
	}
	return strings.Join(words, "")
}

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string {
	return def.CamelToSnake(s)
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string) string {
	return def.SnakeToCamel(s)
}

// SnakeToCamelLower returns a string converted from snake case to lowercase
func SnakeToCamelLower(s string) string {
	return def.SnakeToCamelLower(s)
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = []string{
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"OS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
}
