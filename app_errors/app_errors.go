package app_errors

import (
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

// Structure for the error messages
type ErrorStruct struct {
	ErrorCode   int
	Description string
}

// String function for the error
func (e ErrorStruct) Error() string {
	return fmt.Sprintf("Error: %v, Description: %v",
		e.ErrorCode, e.Description)
}

// List of errors for this app
var ErrorMap = map[int]ErrorStruct{
	999:  {999, "Configuration Error"},
	1000: {1000, "Level Shape not rectangular"},
	1001: {1001, fmt.Sprintf("Incorrect level size. Y exceeds %v", constants.MaxLevelSize)},
	1002: {1002, fmt.Sprintf("Incorrect level size. X exceeds %v", constants.MaxLevelSize)},
	1003: {1003, "Invalid tile ID"},
	1004: {1004, "Incorrect level type format"},
	1005: {1005, "Incorrect level ID type"},
	// 1006: {1006, "No Starting point in level"},
	// 1007: {1007, "More than one Starting point in level"},
	// 1008: {1008, "No Exit point in level"},
	// 1009: {1009, "More than one Exit point in level"},
}
