package server

import (
	"fmt"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func SimpleServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)
}
