package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
