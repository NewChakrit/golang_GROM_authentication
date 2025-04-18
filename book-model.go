package main

import (
	"gorm.io/gorm"
	"log"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Authur      string `json:"authur"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name string
	Book []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
}

// ----- Create Book ----- //

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ----- Get Book ----- //

func getBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error geting book: %v", result.Error)
	}

	return &book
}

// ----- Get All Book ----- //

func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books) // Find ถ้าไม่ใส่เงื่อนไขจะเท่ากับ Select *

	if result.Error != nil {
		log.Fatalf("Error geting book: %v", result.Error)
	}

	return books
}

// ----- Update Book ----- //

func updateBook(db *gorm.DB, book *Book) error {
	//result := db.Save(&book)

	result := db.Model(&book).Updates(*book) // handle created_at != nil

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ----- Delete Book ----- //

func deleteBook(db *gorm.DB, id int) error {
	var book Book
	//result := db.Delete(&book, id)

	// ถ้า hard delete ต้องใช้ db.Unscoped().Delete(&book, id)

	result := db.Unscoped().Delete(&book, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ----- Search Book ----- //

func searchBook(db *gorm.DB, bookName string) *Book {
	var book Book

	result := db.Where("name = ?", bookName).First(&book) // First คือ quiry limit 1 row
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}

	return &book
}

// ----- Search Books Great Than 1 Book ----- //

func searchBooks(db *gorm.DB, bookName string) []Book {
	var books []Book

	result := db.Where("name = ?", bookName).Order("price").Find(&books) // First คือ quiry limit 1 row
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}

	return books
}

// ----- Create Publisher ----- //

func createPublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Create(publisher)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func createAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func createBookWithAuthor(db *gorm.DB, book *Book) error {
	// First, create the book
	if err := db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

func getBookWithPublisher(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Publisher").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func getBookWithAuthors(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Authors").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func listBooksOfAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	var books []Book
	result := db.Joins("JOIN author_books on author_books.book_id = books.id").
		Where("author_books.author_id = ?", authorID).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
