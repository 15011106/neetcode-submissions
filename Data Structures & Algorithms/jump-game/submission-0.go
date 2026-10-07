func canJump(nums []int) bool {

    if len(nums) == 1 {
        return true
    }

    goal := len(nums) - 1

    for i := len(nums) - 2; i >= 0; i-- {

        if goal <= nums[i]+i {
            goal = i
        }

    }

    if goal == 0 {
        return true
    }

    return false
}
