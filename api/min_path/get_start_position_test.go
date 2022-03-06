package min_path

import (
	"testing"

	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/stretchr/testify/assert"
)

// test function GetStartPosition when the level is correct
func TestGetStartPosition(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	xPos, yPos, err := GetStartPosition(level)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, xPos)
	assert.Equal(t, 1, yPos)
}

// test function GetStartPosition when the level has no starting points
func TestGetStartPositionFailNoStarts(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}
	_, _, err := GetStartPosition(level)
	assert.Equal(t, app_errors.ErrorMap[1006], err)
}

// test function GetStartPosition when the level has multiple starting points
func TestGetStartPositionFailMultipleStarts(t *testing.T) {
	level := [][]int{{1, 1, 1, 2, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	_, _, err := GetStartPosition(level)
	assert.Equal(t, app_errors.ErrorMap[1007], err)
}
