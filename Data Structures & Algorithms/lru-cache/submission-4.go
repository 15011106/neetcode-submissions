

type LRUCache struct {
    cache map[int]*node
    capacity int
    tail *node
    head *node
}

type node struct{
key, val int
next *node
prev *node
}

func Constructor(capacity int) LRUCache {

    tail := &node{}
    head := &node{}
    head.next = tail
    tail.prev = head

    return LRUCache{
        cache : make(map[int]*node),
        capacity: capacity,
        tail : tail,
        head: head,
}
}

func (this *LRUCache) Get(key int) int {
    if n, ok := this.cache[key]; ok{
        n.prev.next = n.next
        n.next.prev = n.prev

        n.prev = this.head
        n.next = this.head.next
      
        this.head.next.prev = n        
        this.head.next = n
		return n.val
	}else{
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
    if n, ok := this.cache[key]; ok{

        n.key = key
        n.val = value

        n.prev.next = n.next
        n.next.prev = n.prev

        n.prev = this.head
        n.next = this.head.next

        
        this.head.next.prev = n        
        this.head.next = n   
	}else{
        next := this.head.next
        prev := this.head
        this.cache[key] = &node{
            key: key,
            val: value,
            next: next,
            prev: prev,
        }
 
        this.head.next.prev = this.cache[key]        
        this.head.next = this.cache[key]

        if len(this.cache) > this.capacity{
          delNode := this.tail.prev

          delNode.prev.next = delNode.next
          delNode.next.prev = delNode.prev

          delete(this.cache, delNode.key)        }
	}
}


