package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ali-ammar-kazmi/book-store/model"
	"github.com/ali-ammar-kazmi/book-store/route"
)

func main() {

	// Return a Router Instance to register and Listen routes.
	router := mux.NewRouter()
	headerOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Initialize a DB Session
	model.DbConnect()

	// method to define routes
	route.RouteInit(router)

	fmt.Println("Server Starting at - http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", handlers.CORS(originOk, headerOk, methodsOk)(router)); err != nil {
		fmt.Println(err.Error())
	}
}
