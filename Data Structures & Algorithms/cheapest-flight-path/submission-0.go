
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	dist := make([]int, n)

    for i:=0; i< n;i++{ 
      dist[i] = 1<<60
  }
    dist[src] = 0


  for i:=0 ;i< k+1; i++{
    tmp := make([]int, n)
    copy(tmp, dist)
    
    for i:=0; i<len(flights);i++{
       s, d, c := flights[i][0], flights[i][1], flights[i][2]
       if dist[s] == 1<<60{
        continue
      }
      if dist[s] + c < tmp[d]{
        tmp[d] = dist[s] + c
      }
    }
    dist = tmp
  }

    ans := -1
    if dist[dst] == 1<<60{
      return ans
    }
      
    return dist[dst]
}