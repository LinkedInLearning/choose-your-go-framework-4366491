package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Static file server
	app.Static("/", "./files")
	

	// Start server
	log.Fatal(app.Listen(":3000"))
}
