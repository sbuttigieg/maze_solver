package api

import (
	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
)

// Check if the level is rectangular by checking if all rows in the level have the same length as the first row.
// Returns nil if there is no error. Returns error 1000 if level is not rectangular
func CheckLevelRectangular(level [][]int) error {
	funcErrCode := 1000             // error ID for error thrown by this function
	firstRowLength := len(level[0]) // find length of first row
	for i := 1; i < len(level); i++ {
		if len(level[i]) != firstRowLength {
			return app_errors.ErrorMap[funcErrCode]
		}
	}
	return nil
}

// Check if the level size is less or equal to the value of constants.MaxLevelSize both in the X and Y directions
// Returns nil if there is no error. Returns error 1001 or 1002 if Y or X exceeds the max size.
func CheckLevelSize(level [][]int) error {
	funcErrCodeYMax := 1001 // error ID for error thrown by this function when the number of rows is greater than the maximum size
	funcErrCodeXMax := 1002 // error ID for error thrown by this function when the number of tiles in a row is greater than the maximum size
	if len(level) > constants.MaxLevelSize {
		return app_errors.ErrorMap[funcErrCodeYMax]
	}
	// Assumes level has already been tested for rectangular shape
	if len(level[0]) > constants.MaxLevelSize {
		return app_errors.ErrorMap[funcErrCodeXMax]
	}
	return nil
}

// Check if the level contains only valid tile IDs as per constants.LevelObjects
// Returns nil if there is no error. Returns error 1003 if invalid tiles are found.
func CheckValidTiles(level [][]int) error {
	funcErrCode := 1003 // error ID for error thrown by this function
	for y := 0; y < len(level); y++ {
		for x := 0; x < len(level[y]); x++ {
			var invalidObject bool = true
			for _, value := range constants.LevelObjects {
				if value == level[y][x] {
					invalidObject = false
				}
			}
			if invalidObject {
				return app_errors.ErrorMap[funcErrCode]
			}
		}
	}
	return nil
}
