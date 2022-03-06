package min_path

// Calculate the minimum survivable path for the level
func CalculateMinPath(level [][]int, sizeX, sizeY int) (int, int, int, error) {
	// Determine starting point
	startX, startY, startErr := GetStartPosition(level)
	if startErr != nil {
		return 0, 0, 0, startErr
	}
	startPos := []int{startX, startY}

	// Determine exit point
	exitX, exitY, exitErr := GetExitPosition(level)
	exitPos := []int{exitX, exitY}
	if exitErr != nil {
		return 0, 0, 0, exitErr
	}

	// Reset result variables
	paths = paths[:0]
	winningPaths = winningPaths[:0]

	// Retrieve the path information from the level
	errPath := findPath(level, startPos, exitPos, startPos, 0, "start")
	if errPath != nil {
		return 0, 0, 0, errPath
	}

	// Result variables
	// minPath, fastestWinningPaths := fastestWinningPath(true)
	minPath, _ := fastestWinningPath()
	possiblePaths := len(paths)
	numOfWinningPaths := len(winningPaths)

	// Return the results
	return minPath, possiblePaths, numOfWinningPaths, nil
}
