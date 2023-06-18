package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()

	// use default http you handle http method and handle parameter
	app.HandleFunc("/hello/{id}", Hello).Methods(http.MethodGet)
	http.ListenAndServe(":8000", app)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println(r.Method)
	fmt.Println("id : " + id)

}
