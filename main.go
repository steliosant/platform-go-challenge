package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"platform-go-challenge/db"
	"platform-go-challenge/handlers"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env (no-op in prod)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found (this is fine in prod)")
	}
	// now os.Getenv works as usual
	// ---- configuration ----
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// ---- database ----
	database, err := db.New(dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer database.Close()

	// ---- router ----
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.UserRouter(database))
	mux.Handle("/users/", handlers.FavouritesRouter(database))
	mux.HandleFunc("/assets", handlers.AssetsRouter(database))
	mux.Handle("/assets/", handlers.AssetsRouter(database))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

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
