package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
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

// General store
var recipes []Recipe

func initDB() {
	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	// marshal the request body into a Recipe struct
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusCreated, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run(":5000")
}
