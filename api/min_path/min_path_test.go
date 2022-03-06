package min_path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// One possible Winning Path
func TestCalculateMinPathOneWin(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 0, 2, 1, 1}, {1, 1, 1, 1, 1}}
	sizeY := len(level)
	sizeX := len(level[0])
	minPath, possiblePaths, winningPaths, err := CalculateMinPath(level, sizeX, sizeY)
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, minPath)
	assert.Equal(t, 8, possiblePaths)
	assert.Equal(t, 1, winningPaths)
}

// Two possible Winning Paths
func TestCalculateMinPathTwoWins(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 2, 0, 1}, {1, 1, 1, 1, 1}}
	sizeY := len(level)
	sizeX := len(level[0])
	minPath, possiblePaths, winningPaths, err := CalculateMinPath(level, sizeX, sizeY)
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, minPath)
	assert.Equal(t, 15, possiblePaths)
	assert.Equal(t, 2, winningPaths)
}

// No possible Winning Paths
func TestCalculateMinPathNoWins(t *testing.T) {
	level := [][]int{{1, 1, 1, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 0, 2, 1, 1}, {1, 1, 1, 1, 1}}
	sizeY := len(level)
	sizeX := len(level[0])
	minPath, possiblePaths, winningPaths, err := CalculateMinPath(level, sizeX, sizeY)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, minPath)
	assert.Equal(t, 4, possiblePaths)
	assert.Equal(t, 0, winningPaths)
}
