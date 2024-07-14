package main

import (
	"log"

	routes "github.com/Chanter327/Butler_backend/routes"
)

func main() {
	r := routes.Router()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}