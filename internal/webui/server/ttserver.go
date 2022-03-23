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
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})
	http.ListenAndServe(":80", nil)
}

//https://tproger.ru/translations/go-web-server/
