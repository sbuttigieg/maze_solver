package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Initialise the API and its endpoints
func InitialiseApi() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/levels", getLevels)
	router.GET("/levels/:id", getLevelById)
	router.DELETE("/:id", deleteLevelById)
	router.PATCH("/:id", updateLevelById)
	router.POST("/level", addLevel)
	return router
}

// Responds with the list of all levels as JSON.
func getLevels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, [][]int{{1, 1}, {1, 1}})
}

// Responds with a level as JSON.
func getLevelById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, [][]int{{id}, {1, 1}})
}

// Responds with the ID of the deleted level.
func deleteLevelById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Level %v deleted", id))
}

// Validates and stores the updated level in the database
// Responds with the ID of the updated level in the database
func updateLevelById(c *gin.Context) {
	updatedLevel := [][]int{} // variable for storing the received level

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	// Call BindJSON to bind the received JSON to updatedLevel
	if err := c.BindJSON(&updatedLevel); err != nil {
		panic(err)
	}

	// TODO: Validate the level
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Level %v Updated", id))
}

// Validates and stores the received level in the database
// Responds with the ID of the new level in the database.
func addLevel(c *gin.Context) {
	newLevel := [][]int{} // variable for storing the received level
	var levelId int       // variable for retrieved level ID from database

	// Call BindJSON to bind the received JSON to newLevel
	if err := c.BindJSON(&newLevel); err != nil {
		panic(err)
	}

	// TODO: Validate the level
	c.IndentedJSON(http.StatusCreated, levelId)
}
