package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("GET /say-bye", handler2)
	http.HandleFunc("GET /say/{word}", handler3)
	http.HandleFunc("GET /json/{word}", handler4)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye"))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")
	w.Write([]byte(word))
}

func handler4(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")

	w.Header().Set("Content-Type", "application/json")

	ans := map[string]string{
		"word": word,
	}

	json.NewEncoder(w).Encode(ans)
}
