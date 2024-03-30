package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Book struct {
	Id          int64  `gorm:"id; primary key; autoIncrement" json:"id"`
	Name        string `gorm:"name" json:"name"`
	Author      string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Db = db
	return Db
}

func Init() {
	db := connect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	Db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	Db.Find(&books)
	return books
}

func (b *Book) GetBookById(id int64) Book {
	var book Book
	Db.Where("id=?", id).First(&book)
	return book
}

func (b *Book) UpdateBook(id int64) *Book {
	Db.Where("id=?", id).Updates(&b)
	return b
}

func (b *Book) DeleteBook(id int64) Book {
	var book Book
	Db.Where("id=?", id).Delete(&book)
	return book
}
