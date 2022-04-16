module github.com/ringoc/gobucket/hello

go 1.18

require (
	github.com/ringoc/gobucket/greetings v0.0.0-00010101000000-000000000000
	github.com/ringoc/gobucket/stringutil v1.0.0
)

replace github.com/ringoc/gobucket/stringutil v1.0.0 => ../stringutil

replace github.com/ringoc/gobucket/greetings => ../greetings
