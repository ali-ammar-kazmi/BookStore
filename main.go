package main

import (
	"fmt"
	"net/http"

	"github.com/ali-ammar-kazmi/book-store/model"
	route "github.com/ali-ammar-kazmi/book-store/route"
	"github.com/gorilla/mux"
)

func main() {

	model.Init()
	router := mux.NewRouter()

	route.RouteInit(router)
	http.Handle("/", router)

	fmt.Printf("Server Starting at - http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		panic(err.Error())
	}
}
