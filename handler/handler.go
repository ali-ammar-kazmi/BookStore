package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ali-ammar-kazmi/book-store/model"
	"github.com/gorilla/mux"
)

var Book model.Book

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Book-Store!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := Book.GetAllBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error)
	}

	json.Unmarshal(bytes, &Book)
	book := Book.CreateBook()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)
	id, _ := strconv.ParseInt(v["id"], 0, 0)
	book := Book.GetBookById(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)
	id, _ := strconv.ParseInt(v["id"], 0, 0)

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error)
	}
	json.Unmarshal(bytes, &Book)

	book := Book.UpdateBook(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)

	id, _ := strconv.ParseInt(v["id"], 0, 0)
	book := Book.DeleteBook(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
