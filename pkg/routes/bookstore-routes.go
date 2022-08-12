package routes

import (
	"github.com/gorilla/mux"
	"github.com/kaifie/bookstore-management/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("POST")
}