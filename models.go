package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Book struct {
	gorm.Model
	Name        string
	Authur      string
	Description string
	Price       uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}

	fmt.Println("Create Book Successful!")
}

// ----- Get Book ----- //

func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error geting book: %v", result.Error)
	}

	return &book
}

// ----- Update Book ----- //

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatalf("Update Book Failed!: %v", result.Error)
	}

	fmt.Println("Update Book Successful!")
}

// ----- Delete Book ----- //

func deleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)

	// ถ้า hard delete ต้องใช้ db.Unscoped().Delete(&book, id)

	if result.Error != nil {
		log.Fatalf("Delete Book Failed!: %v", result.Error)
	}

	fmt.Println("Delete Book Successful!")
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
