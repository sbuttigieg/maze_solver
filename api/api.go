package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
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
	var levelId int           // variable for retrieved level ID from database
	var errMsg error          // variable used to store the error message

	// Retrieve id param from request and convert to an integer
	// Returns error 1005 if conversion to integer fails fails
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg = app_errors.ErrorMap[1005]
	}

	// Call BindJSON to bind the received JSON to updatedLevel
	// Returns error 1004 if the received level is not in the format [][]int
	if err := c.BindJSON(&updatedLevel); err != nil {
		errMsg = app_errors.ErrorMap[1004]
	}

	// Validate the level
	validatedLevel, errValidation := validateLevel(updatedLevel)
	if errValidation != nil {
		errMsg = errValidation
	}

	if errMsg == nil {
		// Format the level data as a string
		// Returns error 999 if JSON marshalling fails
		level, err := json.Marshal(validatedLevel.Level)
		if err != nil {
			errMsg = app_errors.ErrorMap[999]
		}

		if errMsg == nil {
			fmt.Println(id, level)
			// Update level in the database
		}
	}

	// Respond by the level ID if successful or by the error if unsuccessful
	if errMsg == nil {
		if levelId > 0 {
			c.IndentedJSON(http.StatusOK, fmt.Sprintf("Level %v Updated", levelId))
		} else {
			c.IndentedJSON(http.StatusOK, "Level not found")
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, errMsg)
	}
}

// Validates and stores the received level in the database
// Responds with the ID of the new level in the database.
func addLevel(c *gin.Context) {
	newLevel := [][]int{} // variable for storing the received level
	var levelId int       // variable for retrieved level ID from database
	var errMsg error      // variable used to store the error message

	// Call BindJSON to bind the received JSON to newLevel
	// Returns error 1004 if the received level is not in the format [][]int
	if err := c.BindJSON(&newLevel); err != nil {
		errMsg = app_errors.ErrorMap[1004]
	}

	if errMsg == nil {
		// Validate the level
		validatedLevel, errValidation := validateLevel(newLevel)
		if errValidation != nil {
			errMsg = errValidation
		}

		if errMsg == nil {
			// Format the level data as a string
			// Returns error 1004 if JSON marshalling fails
			level, err := json.Marshal(validatedLevel.Level)
			if err != nil {
				errMsg = app_errors.ErrorMap[1004]
			}

			if errMsg == nil {
				// Insert level in the database
				fmt.Println(level)
			}
		}
	}

	// Respond by the level ID if successful or by the error if unsuccessful
	if errMsg == nil {
		c.IndentedJSON(http.StatusCreated, levelId)
	} else {
		c.IndentedJSON(http.StatusBadRequest, errMsg)
	}
}

func validateLevel(levelToValidate [][]int) (constants.LevelTableFields, error) {
	level := constants.LevelTableFields{Level: levelToValidate}

	var errMsg error // variable to store the error message from validation

	// Sequence of the validation functions to perform on the new level before storing it into the database
	var validateSequence = map[int]func(level [][]int) error{
		0: CheckLevelRectangular,
		1: CheckLevelSize,
		2: CheckValidTiles,
	}

	// Loop through the validation sequence and determine any error detected.
	// Will not complete the validation sequence if an error is detected
	for i := 0; i < len(validateSequence); i++ {
		if errMsg == nil {
			errMsg = validateSequence[i](levelToValidate)
		}
	}

	if errMsg == nil {
		// Determine the size of the level
		level.Size_y = len(levelToValidate)
		level.Size_x = len(levelToValidate[0])

		// TODO: Calculate the minimum survivable path for the level
	}
	if errMsg == nil {
		return level, nil
	} else {
		return level, errMsg
	}
}
