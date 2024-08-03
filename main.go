package main

import (
	"book-management-system/internal/config"
	"book-management-system/internal/handler"
	"book-management-system/internal/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Printf("Applied %d migrations", n)

	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Category routes
	categories := api.Group("/categories")
	{
		categories.GET("", handler.GetAllCategories)
		categories.POST("", handler.CreateCategory)
		categories.GET("/:id", handler.GetCategory)
		categories.DELETE("/:id", handler.DeleteCategory)
		categories.GET("/:id/books", handler.GetBooksByCategory)
	}

	// Book routes
	books := api.Group("/books")
	{
		books.GET("", handler.GetAllBooks)
		books.POST("", handler.CreateBook)
		books.GET("/:id", handler.GetBook)
		books.DELETE("/:id", handler.DeleteBook)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
