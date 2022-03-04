package api

import (
	"testing"

	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/constants"
)

// test function CheckLevelRectangular when the level is correct
func TestRectangularCorrect(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 2, 1}, {1, 1, 1, 1, 1}}
	if CheckLevelRectangular(level) != nil {
		t.Fatalf("Rectangular level failed - Should Pass")
	}
}

// test function CheckLevelRectangular when row 1 is incorrect
func TestRectangularIncorrectRow1(t *testing.T) {
	level := [][]int{{1, 1, 0, 1}, {1, 0, 0, 2, 1}, {1, 1, 1, 1, 1}}
	if CheckLevelRectangular(level) != app_errors.ErrorMap[1000] {
		t.Fatalf("Non Rectangular level at row 1 should Fail with error 1000")
	}
}

// test function CheckLevelRectangular when other row than row 1 is incorrect
func TestRectangularIncorrectOtherRow(t *testing.T) {
	level := [][]int{{1, 1, 0, 1, 0}, {1, 0, 1, 2, 1}, {1, 1, 1, 1, 1, 1}}
	if CheckLevelRectangular(level) != app_errors.ErrorMap[1000] {
		t.Fatalf("Non Rectangular level at other row should Fail with error 1000")
	}
}

// test function CheckLevelSize when the level is correct
func TestSizeCorrect(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 2, 1}, {1, 1, 1, 1, 1}}
	if CheckLevelSize(level) != nil {
		t.Fatalf("Correct sized level failed - Should Pass")
	}
}

// test function CheckLevelSize when Y is too large
func TestSizeIncorrectY(t *testing.T) {
	level := [][]int{
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1},
		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 2, 1}, {1, 1, 1, 0},
	}
	if CheckLevelSize(level) != app_errors.ErrorMap[1001] {
		t.Fatalf("Incorrect Y sized level should Fail with error 1001")
	}
}

// test function CheckLevelSize when X is too large
func TestSizeIncorrectX(t *testing.T) {
	level := [][]int{
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
	}
	if CheckLevelSize(level) != app_errors.ErrorMap[1002] {
		t.Fatalf("Incorrect X sized level should Fail with error 1002")
	}
}

// test function CheckValidTiles when all tiles are valid
func TestTileValid(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}
	if CheckValidTiles(level) != nil {
		t.Fatalf("Valid tiled level failed - Should Pass")
	}
}

// test function CheckValidTiles when a tile is invalid
func TestTileInvalid(t *testing.T) {
	invalidTile := len(constants.LevelObjects) + 1
	level := [][]int{{1, 1, 1, 0, 1}, {1, invalidTile, 0, 0, 1}, {1, 1, 1, 1, 1}}
	if CheckValidTiles(level) != app_errors.ErrorMap[1003] {
		t.Fatalf("Incorrect Y sized level should Fail with error 1003")
	}
}
