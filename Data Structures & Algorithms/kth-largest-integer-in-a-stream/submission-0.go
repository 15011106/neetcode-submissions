type KthLargest struct {
  minHeap minHeap
  k int
}

type minHeap []int

func Constructor(k int, nums []int) KthLargest {
  
  h := minHeap(nums)
  heap.Init(&h)
  	for len(h) > k {
		heap.Pop(&h)
	}

   return  KthLargest{
    minHeap: h,
    k : k,
   }
}


func (this *KthLargest) Add(val int) int {
    heap.Push(&this.minHeap, val)
    
    if len(this.minHeap) > this.k{
        heap.Pop(&this.minHeap)
    }

    return this.minHeap[0]
}



func (k minHeap) Len() int{
	return len(k)
}

func (k minHeap) Less(i, j int) bool{
	return k[i] < k[j]
}

func (k minHeap) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k *minHeap) Pop() any{
	old := *k
	l := len(old)

	item := old[l-1]
	*k = old[:l-1]

	return item
}

func (k *minHeap) Push(x any){

  *k = append(*k, x.(int))
}
