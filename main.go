package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Recipe Define struct for recipe
type Recipe struct {
	ID           string    `json:"ID"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Instructions []string  `json:"instructions"`
	Ingredients  []string  `json:"ingredients"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func createRecipes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	router := gin.Default()
	router.POST("/", createRecipes)
	router.Run(":5000")
}
