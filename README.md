# snaker

[![Build Status](https://travis-ci.org/nochso/snaker.svg?branch=master)](https://travis-ci.org/nochso/snaker)
[![GoDoc](https://godoc.org/github.com/nochso/snaker?status.svg)](https://godoc.org/github.com/nochso/snaker)

This is a small utility to convert camel cased strings to snake case and back, except some defined words.

## Example

```go
s := New("ID", "IMDB")
s.CamelToSnake("IMDBID")    // imdb_id
s.SnakeToCamel("imdb_name") // IMDBName
```