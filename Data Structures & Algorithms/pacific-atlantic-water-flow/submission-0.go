
func pacificAtlantic(heights [][]int) [][]int {
pacificVisited := make([][]bool, len(heights))
atlanticVisited := make([][]bool, len(heights))

for i := 0; i < len(heights); i++ {
    pacificVisited[i] = make([]bool, len(heights[0]))
    atlanticVisited[i] = make([]bool, len(heights[0]))
}

	var pacificDfs func(y, x int)
	var atlanticDfs func(y, x int)
	pacificDfs = func(y, x int) {
		pacificVisited[y][x] = true
		next := [4][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
		for i := 0; i < 4; i++ {
			dy := next[i][0]
			dx := next[i][1]

			nextX := x + dx
			nextY := y + dy
			if nextX >= 0 && nextX < len(heights[0]) && nextY >= 0 && nextY < len(heights) {
				if !pacificVisited[nextY][nextX] && heights[y][x] <= heights[nextY][nextX] {
					pacificDfs(nextY, nextX)
				}
			}
		}
	}

	atlanticDfs = func(y, x int) {
		atlanticVisited[y][x] = true
		next := [4][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
		for i := 0; i < 4; i++ {
			dy := next[i][0]
			dx := next[i][1]

			nextX := x + dx
			nextY := y + dy
			if nextX >= 0 && nextX < len(heights[0]) && nextY >= 0 && nextY < len(heights) {
				if !atlanticVisited[nextY][nextX] && heights[y][x] <= heights[nextY][nextX] {
					atlanticDfs(nextY, nextX)
				}
			}
		}
	}

	for i := 0; i < len(heights); i++ {
		if i == 0 {
			for j := 0; j < len(heights[0]); j++ {
				pacificDfs(i, j)
			}
		} else {
			pacificDfs(i, 0)
		}

		if i == len(heights)-1 {
			for j := 0; j < len(heights[0]); j++ {
				atlanticDfs(i, j)
			}
		} else {
			atlanticDfs(i, len(heights[0])-1)
		}

	}
	ans := [][]int{}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if pacificVisited[i][j] && atlanticVisited[i][j] {
				tempArr := []int{i, j}
				ans = append(ans, tempArr)
			}
		}
	}

	return ans
}
