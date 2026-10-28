// K Closest Points to Origin
// https://neetcode.io/problems/k-closest-points-to-origin

func kClosest(points [][]int, k int) [][]int {
    
    
    var pq mh
    heap.Init(&pq)
    
    for i:=0;i <len(points);i++{
        
        item := item{
                idx :i,
                sum : points[i][0]*points[i][0] + points[i][1]*points[i][1],
            }
        
        heap.Push(&pq, item)
        }
        
     ans := [][]int{}
     for i:=0; i< k; i++{
         
         x := heap.Pop(&pq).(item)
         ans = append(ans, points[x.idx])
         
         }
         
            return ans   
    
}


type item struct{
    idx int
    sum int
    }
    
type mh []item

func (m mh) Less (i, j int) bool{ 
    return m[i].sum < m[j].sum
    }

func (m mh)Len() int{
    return len(m)
    } 

func (m mh) Swap(i,j int){
    m[i],m[j] = m[j], m[i]
    }
    
func (m *mh) Push(x any){
    *m = append(*m, x.(item))
    }

func (m *mh) Pop() any{
    old := *m
    l := len(old)
    
    item := old[l-1]
    *m = old[:l-1]
    
    return item
    }