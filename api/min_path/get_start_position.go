package min_path

import (
	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
)

func GetStartPosition(level [][]int) (int, int, error) {
	funcErrCodeNoStart := 1006        // error ID for no start tile in level
	funcErrCodeMultipleStarts := 1007 // error ID for more than one start tile in level
	var xPos, yPos int                // store x and y positions returned by function
	var located bool

	// Find key for starting point. If not found throw configuration error.
	startKey := constants.LevelObjects["starting point"] //getLevelObjectKey("starting point")
	if startKey < 0 {
		return xPos, yPos, app_errors.ErrorMap[999]
	}

	// Locate position of starting point. If multiple starting points are found, throw error
	for yIndex := range level {
		for xIndex, xValue := range level[yIndex] {
			if xValue == startKey {
				if !located {
					xPos = xIndex
					yPos = yIndex
					located = true
				} else {
					return xPos, yPos, app_errors.ErrorMap[funcErrCodeMultipleStarts]
				}
			}
		}
	}
	// If no starting point is found, throw error
	if !located {
		return xPos, yPos, app_errors.ErrorMap[funcErrCodeNoStart]
	}
	return xPos, yPos, nil
}
