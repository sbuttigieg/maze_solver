package min_path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Locations are equal
func TestCompareLocationsSame(t *testing.T) {
	location1 := []int{2, 1}
	location2 := []int{2, 1}
	result := compareLocations(location1, location2)
	assert.Equal(t, true, result)
}

// Locations are not equal
func TestCompareLocationsNotSame(t *testing.T) {
	location1 := []int{2, 2}
	location2 := []int{1, 2}
	result := compareLocations(location1, location2)
	assert.Equal(t, false, result)
}

// Locations are equal
func TestCheckForWinTrue(t *testing.T) {
	exitPos := []int{4, 1}
	location := []int{4, 1}
	winPath := []int{4, 4, 4, 3, 4, 2, 4, 1}
	result := checkForWin(0, exitPos, location, winPath)
	assert.Equal(t, true, result)
}

func TestCheckForWinFalse(t *testing.T) {
	exitPos := []int{4, 1}
	location := []int{4, 2}
	winPath := []int{4, 4, 4, 3, 4, 2}
	result := checkForWin(2, exitPos, location, winPath)
	assert.Equal(t, false, result)
}
