type maxHeap struct{
	item []int
}

func (h maxHeap) Len () int{
	return len(h.item)
}
func (h maxHeap) Less (a, b int) bool{
	item := h.item
	return item[a] > item[b]
}
func (h maxHeap) Swap (a, b int){
	h.item[a], h.item[b] = h.item[b], h.item[a]
}

func (h *maxHeap) Push(a any){
	h.item = append(h.item, a.(int))
}

func (h *maxHeap) Pop() any{
	old := h.item[len(h.item)-1]
	h.item = h.item[0:len(h.item)-1]

	return old
}

func lastStoneWeight(stones []int) int {

	h := &maxHeap{}
	heap.Init(h)
	for _,v := range stones{
		heap.Push(h, v)
	}

	if h.Len() == 1{
		return heap.Pop(h).(int)
	}

	for h.Len() > 1{
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		if first > second {
			heap.Push(h, first-second)
		}else{
			heap.Push(h, second-first)
		}
	}
	
	ans := heap.Pop(h).(int)
	return ans
}
