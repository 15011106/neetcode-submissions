func numIslands(grid [][]byte) int {
    

	visited := make([][]bool, len(grid))

	for i:=0 ;i<len(grid); i++{
		visited[i] = make([]bool, len(grid[i]))
	}

	next := [4][2]int{{0,1},{1,0},{-1,0},{0,-1}}
	
	var dfs func (x,y int)
	dfs = func (x,y int){

		visited[y][x] = true
		for i:=0;i<4;i++{
			nextY := y + next[i][0]
			nextX := x + next[i][1]

			if nextY >= 0 && nextY < len(grid) && nextX >=0 && nextX < len(grid[0]){
				if !visited[nextY][nextX] && string(grid[nextY][nextX]) == "1"{
					dfs(nextX, nextY)
				}
			}
		}
	}

	ans:= 0
	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[i]); j++{
			if string(grid[i][j]) =="1" && !visited[i][j]{
				dfs(j,i)
				ans++
			}
		}
	}

	return ans
}
