func subsets(nums []int) [][]int {
	ans := [][]int{}
	var dfs func(int, []int)
	
	dfs = func(level int, curArr []int) {
		for i:=level; i<len(nums); i++{
			curArr = append(curArr, nums[i])
			tmp := make([]int, len(curArr))
			
			copy(tmp, curArr)
			ans = append(ans, tmp)
			
			dfs(i+1, curArr)
			curArr = curArr[:len(curArr)-1]
		}
}

	dfs(0, []int{})
	ans = append(ans, []int{})
	return ans
}


