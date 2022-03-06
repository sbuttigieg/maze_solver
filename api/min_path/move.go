package min_path

import (
	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
)

func move(level [][]int, location, path []int, direction string) ([]int, error) {
	// New location coordinates after moving.
	// Return error if direction is invalid
	var newLocationX, newLocationY int
	switch direction {
	case "up":
		newLocationX = location[0]
		newLocationY = location[1] - 1
	case "left":
		newLocationX = location[0] - 1
		newLocationY = location[1]
	case "right":
		newLocationX = location[0] + 1
		newLocationY = location[1]
	case "down":
		newLocationX = location[0]
		newLocationY = location[1] + 1
	default:
		return location, app_errors.ErrorMap[999]
	}

	// If new location has already been passed through do not go again
	for i, v := range path {
		// Check if index is an even number (as X location is always on even index)
		if i%2 == 0 {
			if v == newLocationX && path[i+1] == newLocationY {
				return location, nil
			}
		}
	}

	// Evaluate the tile to move to and return new location if move is possible
	// Return error if tile is invalid
	switch level[newLocationY][newLocationX] {
	case constants.LevelObjects["open tile"]:
		// fmt.Println("open tile", newLocationX, newLocationY)
		return []int{newLocationX, newLocationY}, nil
	case constants.LevelObjects["wall"]:
		// fmt.Println("wall", newLocationX, newLocationY)
		return location, nil
	case constants.LevelObjects["starting point"]:
		// fmt.Println("back to start", newLocationX, newLocationY)
		return location, nil
	default:
		return location, app_errors.ErrorMap[999]
	}
}
