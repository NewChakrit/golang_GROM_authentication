package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})
	fmt.Println("Migrate Successful!")

	// ----- Create Book ----- //
	//newBook := &Book{
	//	Name:        "New",
	//	Authur:      "Christ Nolan",
	//	Description: "History",
	//	Price:       355,
	//}

	//createBook(db, newBook)

	// ----- Get Book ----- //
	//currentBook := getBook(db, 1)

	//fmt.Println(currentBook)

	// ----- Update Book ----- //
	//currentBook.Name = "Name"
	//currentBook.Price = 280
	//
	//updateBook(db, currentBook)

	// ----- Delete Book ----- //
	// db.Delete คือถ้าใน table มี field deleted_at มันจะเป็นการ soft delete ไม่ใช่การลบ database ออกไปจริงๆ
	// ใน pgadmin จะ stamp deleted_at
	currentBook := getBook(db, 1)

	// แต่ถ้าลอง get ดูจะไม่เจอแล้ว
	fmt.Println(currentBook) // record not found
}
