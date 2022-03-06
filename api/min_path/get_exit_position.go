package min_path

import (
	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
)

func GetExitPosition(level [][]int) (int, int, error) {
	funcErrCodeNoExit := 1008        // error ID for no exit tile in top row of level
	funcErrCodeMultipleExits := 1009 // error ID for more than one exit tile in top row of level
	var xPos, yPos int               // store x and y positions returned by function
	var located bool

	// Find key for starting point. If not found throw configuration error.
	exitKey := constants.LevelObjects["open tile"] //getLevelObjectKey("open tile")
	if exitKey < 0 {
		return xPos, yPos, app_errors.ErrorMap[999]
	}

	// Locate position of exit point in top row. If multiple exit points are found, throw error
	if !located {
		for xIndex, xValue := range level[0] {
			if xValue == exitKey {
				if !located {
					xPos = xIndex
					yPos = 0
					located = true
				} else {
					return xPos, yPos, app_errors.ErrorMap[funcErrCodeMultipleExits]
				}
			}
		}
	}
	// If no exit point is found, throw error
	if !located {
		return xPos, yPos, app_errors.ErrorMap[funcErrCodeNoExit]
	}
	return xPos, yPos, nil
}
