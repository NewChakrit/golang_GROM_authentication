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

func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error geting book: %v", result.Error)
	}

	return &book
}

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatalf("Update Book Failed!: %v", result.Error)
	}

	fmt.Println("Update Book Successful!")
}

func deleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)

	if result.Error != nil {
		log.Fatalf("Delete Book Failed!: %v", result.Error)
	}

	fmt.Println("Delete Book Successful!")
}
