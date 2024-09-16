package gotaghttp_test

import (
	"bytes"
	"fmt"
	"net"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
)

func doRequest[T any](pattern string, url string, contentType string, body []byte, testCreator gotag.StructDecoder[T, *http.Request], validate func(T)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST "+pattern, func(w http.ResponseWriter, r *http.Request) {
		p, err := testCreator.Decode(r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		validate(p)
		w.WriteHeader(200)
	})
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return err
	}
	port := l.Addr().(*net.TCPAddr).Port
	s := &http.Server{Handler: mux}
	go s.Serve(l)
	defer s.Close()
	baseUrl := fmt.Sprintf("http://localhost:%d", port)
	d := bytes.NewReader(body)
	r, err := http.Post(baseUrl+url, contentType, d)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("handler failed")
	}
	return nil
}
