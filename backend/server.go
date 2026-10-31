package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server works!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server Running on port:3000!")
	http.ListenAndServe(":3000", nil)
}
