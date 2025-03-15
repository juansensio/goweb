package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./templates", ".html")

	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/", static.New("./public"))

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/layout", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"Body":  "Hello, World!",
		})
	})

	// Start the server on port 4269
	log.Fatal(app.Listen(":4269"))
}
