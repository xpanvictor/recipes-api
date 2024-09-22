// Recipes API
//
// This is a sample recipes API. You can find out more about
// the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
// Schemes: http
// Host: localhost:5000
// BasePath: /
// Version: 1.0.0
// Contact: xpanvictor@gmail.com
// <xpanvictor@gmail.com> https://xpanvictor.github.io
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"strings"
	"time"
)

// Recipe Define struct for recipe
// swagger:parameters recipes newRecipe
type Recipe struct {
	// swagger:ignore
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

// swagger:operation POST /recipes recipes newRecipe
// Create a new recipe
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'400':
//	    description: Invalid input
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

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation GET /recipes/search/{tag} recipes searchRecipe
// ---
// description: Returns list of recipes by tag
// parameters:
//   - in: path
//     name: tag
//     required: true
//     type: string
//     description: The tag to search for
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//	  description: Successful operation
//	'404':
//	  description: Recipe not found
func SearchRecipeHandler(c *gin.Context) {
	var tag = c.Param("tag")
	var listRecipes []Recipe

	// going through list of recipes
	for i := 0; i < len(recipes); i++ {
		for _, recipeTag := range recipes[i].Tags {
			if strings.EqualFold(tag, recipeTag) {
				listRecipes = append(listRecipes, recipes[i])
				break
			}
		}
	}

	c.JSON(http.StatusOK, listRecipes)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// ---
// description: Updates a recipe by ID
// parameters:
//   - in: path
//     name: id
//     required: true
//     type: string
//     description: The ID of the recipe to update
//   - in: body
//     name: body
//     required: true
//     schema:
//     $ref: "#/definitions/Recipe"
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//	  description: Successful operation
//	'400':
//	  description: Invalid input
//	'404':
//	  description: Recipe not found
func UpdateRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var recipeID = c.Param("id")
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == recipeID {
			recipes[i] = recipe
			c.JSON(http.StatusOK, recipe)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Recipe not found",
	})
}

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// ---
// description: Deletes a recipe by ID
// parameters:
//   - in: path
//     name: id
//     required: true
//     type: string
//     description: The ID of the recipe to delete
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//	  description: Successful operation
//	'404':
//	  description: Recipe not found
func DeleteRecipeHandler(c *gin.Context) {
	var recipeID = c.Param("id")
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == recipeID {
			recipes = append(recipes[:i], recipes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Recipe deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Recipe not found",
	})
}

func main() {
	initDB()
	router := gin.Default()
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/search/:tag", SearchRecipeHandler)
	router.POST("/recipes", NewRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.Run(":5000")
}
