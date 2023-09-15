package main

import (
	"hexagonal/handler"
	"hexagonal/newprovider"
	"hexagonal/route"
	"hexagonal/usecase"
	"net/http"
)

func main() {
	// Initialize Repositories (DB, Cache, etc.)
	repos := newprovider.NewRepositories("A", "B", "C", "D")

	// Initialize Use Cases
	ucs := usecase.NewUseCases(repos)

	// Initialize Handlers
	handlers := handler.NewHandlers(ucs)

	// Set up routes
	route.InitializeRoutes(handlers)

	// Run the server
	http.ListenAndServe(":8080", nil)
}
