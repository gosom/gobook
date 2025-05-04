package main

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func main() {
	infoHandler := &helloInfoHandler{}

	http.HandleFunc("/", handler)
	http.Handle("/info", infoHandler)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

type goInfo struct {
	Version      string `json:"version"`
	OS           string `json:"os"`
	Arch         string `json:"arch"`
	NumCPU       int    `json:"num_cpu"`
	NumGoroutine int    `json:"num_goroutine"`
}

type helloInfoHandler struct{}

func (h *helloInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := goInfo{
		Version:      runtime.Version(),
		OS:           runtime.GOOS,
		Arch:         runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(resp)

	w.Write(data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
