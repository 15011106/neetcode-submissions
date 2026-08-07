func countBits(n int) []int {
	ans := []int{}

	for i:= 0 ; i <= n; i++ {
		tempNum := 0
		curNum := i

		for curNum > 0 {
			tempBit := curNum & 1
			if tempBit == 1{
				tempNum++
			}
			curNum = curNum >> 1
		}

		ans = append(ans, tempNum)
	}

	return ans 
}
