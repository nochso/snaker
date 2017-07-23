# snaker

[![GoDoc](https://godoc.org/github.com/nochso/snaker?status.svg)](https://godoc.org/github.com/nochso/snaker)
[![Build Status](https://travis-ci.org/nochso/snaker.svg?branch=master)](https://travis-ci.org/nochso/snaker)
[![Coverage Status](https://coveralls.io/repos/github/nochso/snaker/badge.svg?branch=master)](https://coveralls.io/github/nochso/snaker?branch=master)

This is a small utility to convert camel cased strings to snake case and back, except some defined words.

## Example

```go
s := New("ID", "IMDB")
s.CamelToSnake("IMDBID")    // imdb_id
s.SnakeToCamel("imdb_name") // IMDBName
```