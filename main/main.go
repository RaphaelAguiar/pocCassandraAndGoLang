package main

import (
	"brajd/cliente"
	"brajd/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Init()
	router := mux.NewRouter()
	var clienteResource cliente.Resource
	router.HandleFunc("/cliente", clienteResource.GetTodosClientes).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", router))
}
