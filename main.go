package main
import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/gofiber/fiber/v2"



)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
		Publisher   int    `json:"publisher"`}

func main(){



	connStr := "postgresql://user:library@localhost:5432/postgres?options"


	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	fmt.Println("hello")

	app:= fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return getHandler(c,db);
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c,db);
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return deleteHandler(c,db);
	})


}



func getHandler(c *fiber.Ctx, db *sql.DB) error {

}

func postHandler(c *fiber.Ctx, db *sql.DB) error {

}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	
}