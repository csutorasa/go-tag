# go-tags

This library simplifies the processing of custom [struct tags](https://go.dev/ref/spec#Tag).
The recommended use case are where framworks are generic, and use-cases could be more specific.
Take inspiration from [Well-known struct tags](https://go.dev/wiki/Well-known-struct-tags),
where most use-cases are structured object data based converters.

## Install

Go 1.22 is the minimum supported version.

```sh
go get github.com/csutorasa/go-tags
```

## Example

The [gotaghttp](gotaghttp/) shows an example of how the [net/http.Request](https://pkg.go.dev/net/http#Request) variables can be parsed.
Request body format, expected query and path parameters or form values are all gerenic,
but most endpoints expect a few well defined values.
