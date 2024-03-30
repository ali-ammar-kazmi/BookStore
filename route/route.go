package route

import (
	"github.com/ali-ammar-kazmi/book-store/handler"
	"github.com/gorilla/mux"
)

var RouteInit = func(router *mux.Router) {

	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/book/", handler.AddBook).Methods("POST")
	router.HandleFunc("/book/", handler.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", handler.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handler.DeleteBook).Methods("DELETE")
}
