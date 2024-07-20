package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
)

var listFilePath string = "C:\\Development\\web\\No-Story-Recipes-Backend\\data\\menuList.json"
var recipeFilePath string = "C:\\Development\\web\\No-Story-Recipes-Backend\\data\\recipes.json"

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/list", getList)
	router.GET("/recipe/:key", getRecipeByKey)

	router.Run("localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func readData(filePath string) *gabs.Container {

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	if filePath == "C:\\Development\\web\\No-Story-Recipes-Backend\\data\\recipes.json" {
		arrayParsed, _ := gabs.ParseJSON([]byte(fileData))
		return arrayParsed
	}
	jsonParsed, _ := gabs.ParseJSON(fileData)

	return jsonParsed
}

func getList(c *gin.Context) {
	parsedData := readData(listFilePath)
	fmt.Println("JSON PARSED AND SENT")
	c.IndentedJSON(http.StatusOK, parsedData)
}

func getRecipeByKey(c *gin.Context) {
	key := c.Param("key")

	parsedData := readData(recipeFilePath)

	gTitle := parsedData.Path("recipes." + key + ".Title")
	gIngredients := parsedData.Path("recipes." + key + ".ingredients")
	gDirections := parsedData.Path("recipes." + key + ".directions")

	title := gTitle.String()
	ingredients := gIngredients.String()
	directions := gDirections.String()

	var msg struct {
		Title       string
		Ingredients string
		Directions  string
	}

	msg.Title = title
	msg.Ingredients = ingredients
	msg.Directions = directions

	c.JSON(http.StatusOK, msg)
}
