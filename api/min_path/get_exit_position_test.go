package min_path

import (
	"testing"

	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/stretchr/testify/assert"
)

// test function GetExitPosition when the level is correct
func TestGetExitPosition(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	xPos, yPos, err := GetExitPosition(level)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, xPos)
	assert.Equal(t, 0, yPos)
}

// test function GetExitPosition when the level has no exits
func TestGetExitPositionFailNoExits(t *testing.T) {
	level := [][]int{{1, 1, 1, 1, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	_, _, err := GetExitPosition(level)
	assert.Equal(t, app_errors.ErrorMap[1008], err)
}

// test function GetExitPosition when the level has multiple exits
func TestGetExitPositionFailMultipleExits(t *testing.T) {
	level := [][]int{{1, 1, 0, 0, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	_, _, err := GetExitPosition(level)
	assert.Equal(t, app_errors.ErrorMap[1009], err)
}
