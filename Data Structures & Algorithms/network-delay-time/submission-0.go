func networkDelayTime(times [][]int, n int, k int) int {

	dist := make([]int, n+1)
	
	// next, cost
	graph := make([][][]int, n+1)
	const INF = 1<<60
	
	for i :=0 ;i<=n; i++{
		dist[i] = INF
	}

	for i:=0; i<len(times); i++{
		curNode := times[i][0]
		nextNode := times[i][1]
		cost := times[i][2]

		graph[curNode] = append(graph[curNode], []int{nextNode, cost})
	}

	pq := &minHeap{}
	dist[k] = 0
	heap.Push(pq, item{node: k, cost: 0})

	for pq.Len()> 0{
		it := heap.Pop(pq).(item)
		if it.cost > dist[it.node]{
			continue
		}

		for _ ,v := range graph[it.node]{
    		d := it.cost + v[1]
    		
			if d < dist[v[0]] {
        		dist[v[0]] = d
        		heap.Push(pq, item{v[0], d})
    		}
		}	
	}

	ans := 0
	for i := 1; i <= n; i++ {
    if dist[i] == INF { 
		return -1 }
    if dist[i] > ans { ans = dist[i] }
}

	return ans 
}


type item struct{
	node int
	cost int

}

type minHeap []item

func (m minHeap) Len() int{
	return len(m)
}

func (m minHeap) Less(i,j int) bool{
	return m[i].cost < m[j].cost
}

func (m minHeap) Swap(i,j int){
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Pop () any{
	n := len(*m)
	old := *m
	
	it := old[n-1]
	*m = old[:n-1]

	return it
}

func (m *minHeap) Push(x any){
	*m = append(*m, x.(item))
}