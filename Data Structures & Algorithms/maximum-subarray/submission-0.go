func maxSubArray(nums []int) int {


	dp := make([]int,len(nums))
	dp[0] = nums[0]
	for i:=1; i<len(nums); i++{
		if nums[i] < dp[i-1] + nums[i]{
			dp[i] = dp[i-1] + nums[i]
		}else{
			dp[i] = nums[i]
		}
}
	max := nums[0]
	
		for i:=1; i<len(dp); i++{
		if max < dp[i] {
			max = dp[i]
		}
	}


	
	return max
}
