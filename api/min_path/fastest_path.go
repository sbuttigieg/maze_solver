package min_path

func fastestWinningPath() (int, [][]int) {
	var minMoves int
	var index []int
	var fastestWinningPaths [][]int

	// If there are winningPaths ==>
	// 1. Set minMoves as the moves for the first winning path and save its index (0)
	// 2. Loop through the remaining paths ( if any) and for each path =>
	//		3. If number of moves for the path is less than minMoves,
	//			3a. Replace index value with the one for the path
	//			3b. Set minMoves to this path's number of moves
	//		4. If the number of moves for the path is equal to minMoves,
	//			4a.	Append index value with the one for the path
	//	5. Loop through index and for each index =>
	//		6. Append the fastest winning paths to fastestWinningPaths
	if len(winningPaths) > 0 {
		minMoves = (len(winningPaths[0]) - 2) / 2
		index = append(index, 0)
		for x := 1; x < len(winningPaths); x++ {
			if (len(winningPaths[x])-2)/2 < minMoves {
				index = index[:0]
				index = append(index, x)
				minMoves = (len(winningPaths[x]) - 2) / 2
			} else if (len(winningPaths[x])-2)/2 == minMoves {
				index = append(index, x)
			}
		}
		for x := 0; x < len(index); x++ {
			fastestWinningPaths = append(fastestWinningPaths, winningPaths[index[x]])
		}
		return minMoves, fastestWinningPaths
	}
	return 0, [][]int{}
}
