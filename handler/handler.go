package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ali-ammar-kazmi/book-store/model"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = "SECRETKEYSECRETKEYSECRETKEY"

func Index(w http.ResponseWriter, r *http.Request) {
	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi %s : Welcome to the Book-Store!", claims.Issuer)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}

	var book model.Book

	books := book.RetrieveAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}

	var book model.Book

	bytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytes, &book)
	book.Create()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}

	var book model.Book

	v := mux.Vars(r)
	id, _ := strconv.ParseInt(v["id"], 0, 0)

	book.RetrieveOne(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}

	var book model.Book

	v := mux.Vars(r)
	id, _ := strconv.ParseInt(v["id"], 0, 0)

	bytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytes, &book)
	book.Update(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := Access(r)

	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Book-Store : Login Please!")
		return
	}

	var book model.Book

	v := mux.Vars(r)
	id, _ := strconv.ParseInt(v["id"], 0, 0)
	book.Delete(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &data)

	bytes, _ := json.Marshal(data["password"])

	bytes, _ = bcrypt.GenerateFromPassword(bytes, 14)
	var user = model.User{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: bytes,
	}

	user.Create()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &data)

	var user model.User
	user.RetrieveOne(data["email"].(string))

	bytes, _ := json.Marshal(data["password"])
	if err := bcrypt.CompareHashAndPassword(user.Password, bytes); err != nil {
		fmt.Println(err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(SECRET_KEY))

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   -1,
	}

	http.SetCookie(w, cookie)

	w.Write([]byte("Logout Successfully"))
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}

func Access(r *http.Request) *jwt.StandardClaims {
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err.Error(), cookie)
		return nil
	}

	if cookie.Value == "" {
		return nil
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	return token.Claims.(*jwt.StandardClaims)
}
