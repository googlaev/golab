package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<H1>Hello</H1>")
	})
	http.ListenAndServe(":80", nil)
}
