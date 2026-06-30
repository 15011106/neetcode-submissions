
func maxArea(heights []int) int {
	j := len(heights)-1
	i := 0
	ans := 0

	 for i < j{
var curVal int		 
		 if heights[i] >= heights[j]{

			 curVal = (j-i) * heights[j]
			 j--
}else{
			 curVal = (j-i) * heights[i]
			 i++
}
		if curVal > ans{
			ans = curVal
}
}
	
	return ans
}