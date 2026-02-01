package main

import (
	"log"
	"net/http"
	"url-shortner/internal/config"
	"url-shortner/internal/handler"
	"url-shortner/internal/middleware"
	"url-shortner/internal/service"
	"url-shortner/internal/store"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()

	port := cfg.Port
	baseURL := cfg.BaseURL + port

	MemoryStore := store.NewMemoryStore()
	shortenerService := service.NewShortener(MemoryStore)

	shortenHandler := handler.NewShortenHandler(shortenerService, baseURL)
	redirectHandler := handler.NewRedirectHandler(MemoryStore)

	mux := http.NewServeMux()
	mux.Handle("/shorten", shortenHandler)
	mux.Handle("/", redirectHandler)

	log.Printf("🚀 URL shortener running on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(":"+port, middleware.Logger(mux)))

}
