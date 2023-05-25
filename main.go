package main

import (
	"fmt"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	// use default http you handle http method and handle parameter
	app.HandleFunc("/hello", Hello)

	http.ListenAndServe(":8000", app)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
}
