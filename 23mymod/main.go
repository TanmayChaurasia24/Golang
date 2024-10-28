package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey there mod user")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>learning golang</h1>"))
}
