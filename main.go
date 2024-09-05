package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/damaisme/go-mysql-bookstore-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
