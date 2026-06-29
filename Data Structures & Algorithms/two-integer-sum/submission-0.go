func twoSum(nums []int, target int) []int {   

	tempMap := make(map[int]int)

	for i:=0 ;i<len(nums);i++{
		if v, ok := tempMap[target - nums[i]]; ok{
			return []int{v,i}
		}
		tempMap[nums[i]] = i
	}

	return nil
}
