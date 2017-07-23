package snaker

import (
	"fmt"
	"testing"
)

func test(t *testing.T, cases map[string]string, fn func(string) string) {
	for input, expected := range cases {
		actual := fn(input)
		if actual != expected {
			t.Errorf(
				"given input %q: expected output %q, got %q",
				input, expected, actual,
			)
		}
	}
}

func TestCamelToSnake(t *testing.T) {
	cases := map[string]string{
		"":    "",
		"One": "one",
		"ONE": "o_n_e",
		"ID":  "id",
		"i":   "i",
		"I":   "i",
		"ThisHasToBeConvertedCorrectlyID": "this_has_to_be_converted_correctly_id",
		"ThisIDIsFine":                    "this_id_is_fine",
		"ThisHTTPSConnection":             "this_https_connection",
		"HelloHTTPSConnectionID":          "hello_https_connection_id",
		"HTTPSID":                         "https_id",
	}
	test(t, cases, CamelToSnake)
}

func TestSnakeToCamel(t *testing.T) {
	cases := map[string]string{
		"":                          "",
		"id":                        "ID",
		"potato_":                   "Potato",
		"this_has_to_be_uppercased": "ThisHasToBeUppercased",
		"this_is_an_id":             "ThisIsAnID",
		"this_is_an_identifier":     "ThisIsAnIdentifier",
	}
	test(t, cases, SnakeToCamel)
}

func TestSnakeToCamelLower(t *testing.T) {
	cases := map[string]string{
		"":                          "",
		"id":                        "id",
		"potato_":                   "potato",
		"this_has_to_be_uppercased": "thisHasToBeUppercased",
		"this_is_an_id":             "thisIsAnID",
		"this_is_an_identifier":     "thisIsAnIdentifier",
		"id_me_please":              "idMePlease",
	}
	test(t, cases, SnakeToCamelLower)
}

func TestS_CamelToSnake_withoutInitialisms(t *testing.T) {
	cases := map[string]string{
		"":                                  "",
		"One":                               "one",
		"ONE":                               "o_n_e",
		"ID":                                "i_d",
		"i":                                 "i",
		"I":                                 "i",
		"HTTPSID":                           "h_t_t_p_s_i_d",
		"ThisHasToBeConvertedIncorrectlyID": "this_has_to_be_converted_incorrectly_i_d",
		"ThisIDIsNotFine":                   "this_i_d_is_not_fine",
		"ThisHTTPSConnection":               "this_h_t_t_p_s_connection",
		"HelloHTTPSConnectionID":            "hello_h_t_t_p_s_connection_i_d",
	}
	s := New()
	test(t, cases, s.CamelToSnake)
}

func TestS_SnakeToCamel_withoutInitialisms(t *testing.T) {
	cases := map[string]string{
		"":                          "",
		"potato_":                   "Potato",
		"this_has_to_be_uppercased": "ThisHasToBeUppercased",
		"this_is_an_id":             "ThisIsAnId",
		"this_is_an_identifier":     "ThisIsAnIdentifier",
		"id": "Id",
	}
	s := New()
	test(t, cases, s.SnakeToCamel)
}

func TestS_SnakeToCamelLower_withoutInitialisms(t *testing.T) {
	cases := map[string]string{
		"":                          "",
		"potato_":                   "potato",
		"this_has_to_be_uppercased": "thisHasToBeUppercased",
		"this_is_an_id":             "thisIsAnId",
		"this_is_an_identifier":     "thisIsAnIdentifier",
		"id":           "id",
		"id_me_please": "idMePlease",
	}
	s := New()
	test(t, cases, s.SnakeToCamelLower)
}

var benchSnakeToCamel = []string{
	"",
	"potato_",
	"this_has_to_be_uppercased",
	"this_is_an_id",
	"this_is_an_identifier",
	"id",
	"id_me_please",
	"a_b_c_d_e_f_g_h",
}

func BenchmarkSnakeToCamel(b *testing.B) {
	var dummy string
	for i := 0; i < b.N; i++ {
		for _, in := range benchSnakeToCamel {
			dummy = SnakeToCamel(in)
		}
	}
	_ = dummy
}

func BenchmarkSnakeToCamelLower(b *testing.B) {
	var dummy string
	for i := 0; i < b.N; i++ {
		for _, in := range benchSnakeToCamel {
			dummy = SnakeToCamelLower(in)
		}
	}
	_ = dummy
}

var benchCamelToSnake = []string{
	"",
	"One",
	"ONE",
	"ID",
	"i",
	"I",
	"ThisHasToBeConvertedCorrectlyID",
	"ThisIDIsFine",
	"ThisHTTPSConnection",
	"HelloHTTPSConnectionID",
	"HTTPSID",
}

func BenchmarkCamelToSnake(b *testing.B) {
	var dummy string
	for i := 0; i < b.N; i++ {
		for _, in := range benchCamelToSnake {
			dummy = CamelToSnake(in)
		}
	}
	_ = dummy
}

func ExampleNew() {
	s := New("ID", "IMDB")
	fmt.Println(s.CamelToSnake("IMDBID"))
	fmt.Println(s.SnakeToCamel("imdb_name"))

	// Output:
	// imdb_id
	// IMDBName
}
