package main

import (
	"log"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"

	"github.com/devenkulkarni/barkbnb/backend/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatal("Failed to create application: ", err)
	}

	if err := application.Run(); err != nil {
		log.Fatal("Failed to run application: ", err)
	}
}
