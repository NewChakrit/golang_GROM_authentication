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
	Price       uint   `json:"price"`
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
