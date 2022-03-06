package min_path

import (
	"testing"

	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
	"github.com/stretchr/testify/assert"
)

// func move(level [][]int, location []int, hits int, direction string) ([]int, int, error) {

// Move up to an open tile
func TestMoveUpOpenTile(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "up"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 1}, newLocation)
	assert.Equal(t, nil, err)
}

// Move left to an open tile
func TestMoveLeftOpenTile(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "left"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{1, 2}, newLocation)
	assert.Equal(t, nil, err)
}

// Move right to an open tile
func TestMoveRightOpenTile(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 0, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "right"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{3, 2}, newLocation)
	assert.Equal(t, nil, err)
}

// Move down to an open tile
func TestMoveDownOpenTile(t *testing.T) {
	level := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 1}
	newPath := []int{0, 0}
	direction := "down"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, nil, err)
}

// Move up to a wall tile
func TestMoveUpWallTile(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "up"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, nil, err)
}

// Move up to the starting location
func TestMoveUpStartLocation(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "up"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, nil, err)
}

// Move to an invalid direction
func TestMoveUpInvalidDirection(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "side"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, app_errors.ErrorMap[999], err)
}

// Move up to an invalid tile
func TestMoveUpInvalidTile(t *testing.T) {
	invalidTile := len(constants.LevelObjects) + 1
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, invalidTile, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{0, 0}
	direction := "up"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, app_errors.ErrorMap[999], err)
}

// Move up to a tile already passed through
func TestMoveUpWallTroddenTile(t *testing.T) {
	level := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}
	location := []int{2, 2}
	newPath := []int{2, 1}
	direction := "up"
	newLocation, err := move(level, location, newPath, direction)
	assert.Equal(t, []int{2, 2}, newLocation)
	assert.Equal(t, nil, err)
}
