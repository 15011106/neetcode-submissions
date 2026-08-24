func combinationSum(nums []int, target int) [][]int {
  

ans := [][]int{}

var dfs func(temp[] int, idx int)
dfs = func(temp []int, idx int){
    if sum(temp) == target{
      ans = append(ans, append([]int{}, temp...))
      return
   }
    if sum(temp) > target{
      return 
    }

     for i:=idx; i<len(nums); i++{
        temp = append(temp, nums[i])
        dfs(temp, i)
        temp = temp[:len(temp)-1]
  }
}
	dfs([]int{},0)
    return ans
}

func sum (arr []int) int{

    val := 0
    for _,v := range arr{
      val += v
  }

    return val
}