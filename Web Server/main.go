package main

import (
	"fmt"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from Web Server</h1>")
}

func main() {
	fmt.Println("Starting server on :8000")
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}
