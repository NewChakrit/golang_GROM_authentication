package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
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
	//fmt.Println("Migrate Successful!")

	// // ----- Setup Fiber ----- // //
	app := fiber.New()

	app.Get("/books", func(ctx *fiber.Ctx) error {
		return ctx.JSON(getBooks(db))
	})

	app.Get("/book/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		book := getBook(db, id)
		return ctx.JSON(book)
	})

	app.Post("/books", func(ctx *fiber.Ctx) error {
		book := new(Book)
		if err := ctx.BodyParser(book); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		err := createBook(db, book)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		return ctx.JSON(fiber.Map{
			"message": "Create Book Successful!",
		})
	})

	app.Put("/book/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		book := new(Book)
		if err := ctx.BodyParser(book); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		book.ID = uint(id)

		err = updateBook(db, book)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		return ctx.JSON(fiber.Map{
			"message": "Update Book Successful!",
		})
	})

	app.Delete("/book/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		book := new(Book)
		if err := ctx.BodyParser(book); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		book.ID = uint(id)

		err = deleteBook(db, id)
		if err := ctx.BodyParser(book); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		return ctx.JSON(fiber.Map{
			"message": "Delete Book Successful!",
		})
	})

	app.Listen(":8080")

	// --------------------------------------------------------//

	// ----- Create Book ----- //
	//newBook := &Book{
	//	Name:        "Aura",
	//	Authur:      "Grace",
	//	Description: "Study",
	//	Price:       700,
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
	//currentBook := getBook(db, 1)
	// db.Delete คือถ้าใน table มี field deleted_at มันจะเป็นการ soft delete ไม่ใช่การลบ database ออกไปจริงๆ
	// ใน pgadmin จะ stamp deleted_at
	//deleteBook(db, 1)

	//fmt.Println(currentBook)

	// ----- Search Book ----- //

	//currentBook := searchBook(db, "Aura")
	//fmt.Println(currentBook)

	// ----- Search Book ----- //

	//currentBooks := searchBooks(db, "Aura")
	////fmt.Println(currentBooks)
	//
	//for _, book := range currentBooks {
	//	fmt.Println(book.ID, book.Name, book.Authur, book.Price)
	//}

}
