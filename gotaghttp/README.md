# gotaghttp

This package defines `gotags.StructTagHandler`s to handle `net/http.Request`s.

## Example

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type ExampleRequest struct {
	UserId   int64  `pathParam:"userId"`
	External bool   `queryParam:"external"`
	Name     string `formValue:"name"`
	Age      int    `formValue:"age"`
}

func main() {
	exampleRequestDecoder := gotag.NewDecoder[ExampleRequest](gotaghttp.DefaultWriters...)
	http.HandleFunc("POST /user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		p, err := exampleRequestDecoder.Decode(r)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(p.UserId, p.External, p.Name, p.Age)
		w.WriteHeader(200)
	})
	http.ListenAndServe(":0", nil)
}
```
