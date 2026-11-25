package main

import (
	"fmt"
	"net/http"

	"github.com/Sahan619/url_shortner/backend/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/api/health", handlers.HealthCheck)
	router.Post("/api/urls", handlers.CreateURL)
	router.Get("/{shortCode}", handlers.RedirectURL)
	router.Get("/api/urls", handlers.GetURLs)
	router.Delete("/api/urls/{shortCode}", handlers.DeleteURL)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
