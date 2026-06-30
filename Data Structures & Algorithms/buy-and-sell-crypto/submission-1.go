func maxProfit(prices []int) int {
	min := prices[0]
	ans := 0
	for i:=1; i< len(prices); i++{

		if prices[i] < min{
			min = prices[i]			
		}else{
			tempAns := prices[i] - min
			if tempAns > ans{
				ans = tempAns
			}
		}
	}
		
		return ans
}