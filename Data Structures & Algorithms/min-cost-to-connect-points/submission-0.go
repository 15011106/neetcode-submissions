func minCostConnectPoints(points [][]int) int {

	var manhattanDist func (x1,x2,y1,y2 int) (cost int)
	var abs func (x,y int) int
	manhattanDist = func(x1,x2,y1,y2 int) (cost int){
		x := abs(x1, x2)
		y := abs(y1, y2)
		return x+y
	}

	abs = func(a,b int) int{
		if a>b {
			return a-b
		}
		return b-a
	}
	
	n := len(points)
	visited := make([]bool, n)
	mh := &MinHeap{{0,0}}
	total, cnt := 0, 0
	heap.Init(mh)

	for cnt < n {
		item := heap.Pop(mh).(Item)
		if visited[item.node] {
			continue
		}
		visited[item.node] = true
		total += item.cost
		cnt++

		cur := item.node
		for next :=0; next <n; next++{
			if !visited[next]{
				cost := manhattanDist(points[cur][0],points[next][0],points[cur][1],points[next][1])
				heap.Push(mh, Item{cost, next})
			}
		}
	}
	return  total
}

type Item struct{
	cost, node int
}

type MinHeap []Item

func (mh MinHeap) Len() int{
	return len(mh)
}

func (mh MinHeap) Less(a, b int) bool{
	return mh[a].cost < mh[b].cost
}

func (mh MinHeap) Swap(a, b int){
	mh[a], mh[b] = mh[b], mh[a]
}

func (mh *MinHeap) Pop() (item any){
	old := *mh
	n := len(old)
	item = old[n-1]
	*mh = old[:n-1]

	return item
}

func (mh *MinHeap) Push(item any){
	*mh = append(*mh, item.(Item))
}
