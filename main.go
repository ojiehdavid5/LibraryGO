package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Book struct {
	Author string `json:"author"`

	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func main() {

	connStr := "postgresql://user:library@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	fmt.Println("connected to the database")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return getHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})
	app.Listen(":8080")

}

func getHandler(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT author, title FROM books")
	if err != nil {
		return c.Status(500).SendString("Error querying the database")
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Title, &book.Author); err != nil {
			return c.Status(500).SendString("Error scanning row")
		}
		books = append(books, book)
	}

	return c.JSON(books)

}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	fmt.Println(book)
	_, err := db.Exec("INSERT INTO books ( author, title, publisher) VALUES ($1, $2, $3)", book.Author, book.Title, book.Publisher)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Error inserting into the database")
	}
	c.SendString("Book added successfully")

	return c.SendStatus(201)

}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	_, err := db.Exec("DELETE FROM books WHERE title = $1", book.Title)
	if err != nil {
		return c.Status(500).SendString("Error deleting from the database")
	}

	return c.SendStatus(204)

}
