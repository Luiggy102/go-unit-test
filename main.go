package main

import (
	"fmt"
	"net/http"

	"github.com/Luiggy102/go-unit-test/controller"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", controller.GetPokemon).
		Methods("GET")

	fmt.Println("Server started at port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Print("Error found")
	}
}
