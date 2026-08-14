type node struct{
  i int
  j int
  l int
}

func islandsAndTreasure(grid [][]int) {
nextArr := [4][2]int{{0,-1},{0,1},{-1,0},{1,0}}

var bfs func(i,j,l int) 
bfs = func (i,j,l int){
	q := []node{}
	q = append(q, node{i,j,l})
for len(q) > 0{
  cur := q[0]
  q = q[1:]
  for i :=0; i<4; i++{

    nextX := cur.i + nextArr[i][0]
    nextY := cur.j + nextArr[i][1]

    if nextX >= 0 && nextX < len(grid) && nextY >=0 && nextY < len(grid[0]){
      curVal := grid[nextX][nextY]
      if curVal != -1 && curVal !=0{
        if curVal > cur.l + 1{
			grid[nextX][nextY] = cur.l + 1
          q = append(q, node{nextX, nextY, cur.l + 1})
      }
  }
  }
  }
}
}

for i :=0; i<len(grid); i++{
  for j:=0; j<len(grid[i]); j++{
   if grid[i][j] == 0{
        bfs(i,j,0)
        }
       }
    }


}