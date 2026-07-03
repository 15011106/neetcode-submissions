func search(nums []int, target int) int {

	l := 0
	h := len(nums)-1

	for l <= h{
		mid := (h+l)/2

		if target >= nums[mid] {
			if target == nums[mid]{
				return mid
			}
			l = mid+1
		}else{
			h = mid-1
		}
}
	
	return -1
}
