package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.Handle("/", MyHandler{})
	http.ListenAndServe(":8080", nil)
}
