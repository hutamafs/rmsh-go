package main

import (
	"database/sql"
	"fmt"
	"log"
	"personal-blog/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = ""
		dbname   = "blog"
)

psqlInfo := fmt.Sprintf("host=localhost port=%d user=%s dbname=%s sslmode=disable",
 port, user, dbname)

// Open DB
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	log.Fatal("Failed to open DB:", err)
}
log.Println("✅ Connected to the database!")
	router := gin.Default();
	router.GET("/posts", handlers.GetAllPosts(db))
	router.POST("/posts", handlers.CreatePost(db))
	router.PUT("/posts/:id", handlers.UpdatePost(db))
	router.Run("localhost:8080");
}
