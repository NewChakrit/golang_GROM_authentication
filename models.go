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
