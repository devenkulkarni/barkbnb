package main

import (
	"log"

	"github.com/devenkulkarni/barkbnb/backend/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	routes.Register(app)

	log.Println("🐶 BarkBnB API started on :8080")

	log.Fatal(app.Listen(":8080"))
}
