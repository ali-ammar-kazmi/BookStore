package test

import (
	"fmt"

	model "github.com/ali-ammar-kazmi/book-store/model"
)

func TestCRUDOperations() {

	// Create a new book
	newBook := &model.Book{Name: "Test Book", Author: "Test Author", Publication: "Test Publication"}
	createdBook := newBook.CreateBook()
	fmt.Println("Created Book:", createdBook)

	// Get all books
	allBooks := createdBook.GetAllBooks()
	fmt.Println("All Books:", allBooks)

	// Update book
	updatedBook := &model.Book{Id: createdBook.Id, Name: "Updated Book", Author: "Updated Author", Publication: "Updated Publication"}
	updatedBook.UpdateBook(createdBook.Id)
	fmt.Println("Updated Book:", updatedBook)

	// Get book by ID
	bookByID := createdBook.GetBookById(createdBook.Id)
	fmt.Println("Book by ID:", bookByID)

	// Delete book
	deletedBook := updatedBook.DeleteBook(updatedBook.Id)
	fmt.Println("Deleted Book:", deletedBook)
}
