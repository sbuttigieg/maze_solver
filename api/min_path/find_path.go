package min_path

// result variables
var paths [][]int
var winningPaths [][]int

func findPath(
	level [][]int,
	startPos, exitPos, startPath []int,
	cycle int,
	incomingDir string,
) error {
	// Variable for storing error
	var err error

	// Make a copy of startPath to use it as a starting point for the new path
	newPath := make([]int, len(startPath))
	copy(newPath, startPath)

	// Determine the current location
	currentLocation := []int{newPath[len(newPath)-2], newPath[len(newPath)-1]}

	// Variable for new location coordinates
	var newLocation []int

	// Not coming from front direction
	if incomingDir != "front" {
		newLocation, err = move(level, currentLocation, newPath, "up")
		if err != nil {
			return err
		}
		if compareLocations(newLocation, currentLocation) {
			setPathAsDead(cycle, "front", newPath, newLocation)
		} else {
			nextPath := make([]int, len(newPath))
			copy(nextPath, newPath)
			nextPath = append(nextPath, newLocation[0], newLocation[1])
			if !checkForWin(cycle, exitPos, newLocation, nextPath) {
				err = findPath(level, startPos, exitPos, nextPath, cycle+1, "behind")
				if err != nil {
					return err
				}
			}
		}
	}

	// Not coming from left direction
	if incomingDir != "left" &&
		currentLocation[0] != 1 {
		newLocation, err = move(level, currentLocation, newPath, "left")
		if err != nil {
			return err
		}
		if compareLocations(newLocation, currentLocation) {
			setPathAsDead(cycle, "left", newPath, newLocation)
		} else {
			nextPath := make([]int, len(newPath))
			copy(nextPath, newPath)
			nextPath = append(nextPath, newLocation[0], newLocation[1])
			if !checkForWin(cycle, exitPos, newLocation, nextPath) {
				err = findPath(level, startPos, exitPos, nextPath, cycle+1, "right")
				if err != nil {
					return err
				}
			}
		}
	}

	// Not coming from right direction
	if incomingDir != "right" &&
		currentLocation[0] != len(level[0])-2 {
		newLocation, err = move(level, currentLocation, newPath, "right")
		if err != nil {
			return err
		}
		if compareLocations(newLocation, currentLocation) {
			setPathAsDead(cycle, "right", newPath, newLocation)
		} else {
			nextPath := make([]int, len(newPath))
			copy(nextPath, newPath)
			nextPath = append(nextPath, newLocation[0], newLocation[1])
			if !checkForWin(cycle, exitPos, newLocation, nextPath) {
				err = findPath(level, startPos, exitPos, nextPath, cycle+1, "left")
				if err != nil {
					return err
				}
			}
		}
	}

	// Not coming from behind :)
	if incomingDir != "behind" &&
		currentLocation[1] != len(level)-2 {
		newLocation, err = move(level, currentLocation, newPath, "down")
		if err != nil {
			return err
		}
		if compareLocations(newLocation, currentLocation) {
			setPathAsDead(cycle, "behind", newPath, newLocation)
		} else {
			nextPath := make([]int, len(newPath))
			copy(nextPath, newPath)
			nextPath = append(nextPath, newLocation[0], newLocation[1])
			if !checkForWin(cycle, exitPos, newLocation, nextPath) {
				err = findPath(level, startPos, exitPos, nextPath, cycle+1, "front")
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Append the dead path to the list of paths, adding special code 696969
func setPathAsDead(cycle int, direction string, deadPath, location []int) {
	toAppend := make([]int, len(deadPath))
	copy(toAppend, deadPath)
	paths = append(paths, deadPath)
	paths[len(paths)-1] = append(paths[len(paths)-1], 696969)
}

// Compare to locations to find if they are equal
func compareLocations(loc1, loc2 []int) bool {
	for i, v := range loc1 {
		if v != loc2[i] {
			return false
		}
	}
	return true
}

// If the new location is the exit location then append the winning path to the list of paths
// Adding special code 999999 followed by the number of moves
func checkForWin(cycle int, exitPos, location, winPath []int) bool {
	if location[0] == exitPos[0] && location[1] == exitPos[1] {
		moves := (len(winPath) - 2) / 2
		toAppend := make([]int, len(winPath))
		copy(toAppend, winPath)
		paths = append(paths, winPath)
		winningPaths = append(winningPaths, toAppend)
		paths[len(paths)-1] = append(paths[len(paths)-1], 999999, moves)
		return true
	}
	return false
}
