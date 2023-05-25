package main

import "net/http"

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.ListenAndServe(":8000", app)
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {

}
