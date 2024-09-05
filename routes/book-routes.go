package routes

import (
	"github.com/damaisme/go-mysql-bookstore-api/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookstoreRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{ID}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{ID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{ID}", controllers.UpdateBook).Methods("PUT")
}
