package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID int 
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt string
}

func GetAllPosts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM posts")
		if err != nil {
			log.Println("Error fetching posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
			return
		}
		defer rows.Close()

		posts := []Post{}

		for rows.Next() {
			var post Post

			err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.PublishedAt)
			if err != nil {
				log.Println("Error scanning post:", err)
				continue
			}
			posts = append(posts, post)
		}
		c.JSON(http.StatusOK, posts)
	}
}

func CreatePost(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		var newPost Post

		if err := c.BindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
			return
		}

		_, err := db.Exec("INSERT INTO posts (title, content)VALUES($1,$2)", &newPost.Title, &newPost.Content)
		if err != nil {
			fmt.Println("Err", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
			return
		 }
		 c.JSON(http.StatusOK, gin.H{"message" : "post created"})
	}
}

func UpdatePost(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		id := c.Param("id");
		post := &Post{}

		if err:= c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		_, err := db.Exec("UPDATE posts SET title=$1, content=$2 where id=$3",post.Title, post.Content, id)
		if err != nil {
			fmt.Println("Err", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to find post"})
			return
		}
		message := fmt.Sprintf("post with id %s has been updated", id)
		c.JSON(http.StatusOK, gin.H{"message" : message})
	}
}
