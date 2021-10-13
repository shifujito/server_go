package main

import (
	"net/http"
)

func main() {
	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Hello, This is go server"))
	}
	http.HandleFunc("/hello", hh)
	http.ListenAndServe("", nil)
}
