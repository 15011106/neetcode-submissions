func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	if len(cost) == 2{
		if dp[0] > dp[1]{
			return dp[1]
		}
		return dp[0]
	}


	for i:=2 ;i<len(cost); i++{
		min := 0

		if dp[i-1] > dp[i-2]{
			min = dp[i-2]
		}else{
			min = dp[i-1]
		}

		dp[i] = min + cost[i]
	}

	ans :=0 
	if dp[len(cost)-1] > dp[len(cost)-2] {
		ans = dp[len(cost)-2]
}else{
	ans = dp[len(cost)-1]
}
	return ans 
}
