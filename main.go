package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"platform-go-challenge/db"
	"platform-go-challenge/handlers"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "platform-go-challenge/docs"
)

func initDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	database, err := db.New(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return database, nil
}

func initServer(database *sql.DB) *http.Server {
	mux := http.NewServeMux()
	
	// Swagger API documentation - serve from docs directory
	swaggerHandler := http.FileServer(http.Dir("./docs"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerHandler))
	
	// Redirect /swagger to /swagger/ for UI
	mux.HandleFunc("/api/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})
	
	// API endpoints
	mux.HandleFunc("/users", handlers.UserRouter(database))
	mux.Handle("/users/", handlers.FavouritesRouter(database))
	mux.HandleFunc("/assets", handlers.AssetsRouter(database))
	mux.Handle("/assets/", handlers.AssetsRouter(database))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}

func main() {
	// Load .env (no-op in prod)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found (this is fine in prod)")
	}

	// ---- database ----
	database, err := initDatabase()
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer database.Close()

	// ---- server ----
	server := initServer(database)

	// ---- graceful shutdown ----
	go func() {
		log.Println("🚀 Server running on port: " + os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("✅ Server exited cleanly")
}
