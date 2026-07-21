
func insert(intervals [][]int, newInterval []int) [][]int {

      ans := [][]int{}
      idx := 0

      // left side
      for i:=0; i< len(intervals); i++{
        if newInterval[0] > intervals[i][1]{
        ans = append(ans, []int{intervals[i][0], intervals[i][1]})
        idx++
        }
      }

      for i:=idx; i<len(intervals); i++{
      if newInterval[1] >= intervals[i][0]{
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        idx++
      }
    }
      ans = append(ans, []int{newInterval[0], newInterval[1]})
      
      for i:=idx; i<len(intervals);i++{
      if newInterval[1] < intervals[i][0]{
        ans = append(ans, []int{intervals[i][0], intervals[i][1]})
    }
}

return ans
}