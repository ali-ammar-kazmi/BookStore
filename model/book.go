package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type CRUD interface {
	Create()
	Update(int64)
	Delete(int64)
	RetrieveOne(int64) interface{}
	RetrieveAll() []interface{}
}

type Book struct {
	Id          int64  `gorm:"id; primary key; autoIncrement" json:"id"`
	Name        string `gorm:"name" json:"name"`
	Author      string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

type User struct {
	Id       int64  `gorm:"id; primary key; autoIncrement" json:"id"`
	Name     string `gorm:"name" json:"name"`
	Email    string `gorm:"email; unique" json:"email"`
	Password []byte `gorm:"password" json:"password"`
}

func DbConnect() {
	data, err := gorm.Open(sqlite.Open("DataBase.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	db := data.Migrator()
	if !db.HasTable(&Book{}) {
		db.CreateTable(&Book{})
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	DB = data
}

func (user *User) Create() {
	DB.Create(&user)
}

func (user *User) RetrieveAll() []User {
	var users []User
	DB.Find(&users)
	return users
}

func (user *User) RetrieveOne(email string) {
	DB.Where("email=?", email).First(&user)
}

func (user *User) Update(id int64) {
	DB.Where("id=?", id).Updates(&user)
}

func (user *User) Delete(id int64) {
	DB.Where("id=?", id).Delete(&user)
}

func (book *Book) Create() {
	DB.Create(&book)
}

func (book *Book) RetrieveAll() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func (book *Book) RetrieveOne(id int64) error {

	if db := DB.Where("id=?", id).First(&book); db.Error != nil {
		return db.Error
	}
	return nil
}

func (book *Book) Update(id int64) {
	DB.Where("id=?", id).Updates(&book)
}

func (book *Book) Delete(id int64) {
	DB.Where("id=?", id).Delete(&book)
}
