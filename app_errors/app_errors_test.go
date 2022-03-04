package app_errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests Error stringer function
func TestError(t *testing.T) {
	testInput := ErrorMap[1000]
	expectedResult := "Error: 1000, Description: Level Shape not rectangular"
	result := testInput.Error()
	assert.Equal(t, expectedResult, result)
}
