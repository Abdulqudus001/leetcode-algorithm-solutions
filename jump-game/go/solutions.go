package jumpgame

func canJump(nums []int) bool {
	if len(nums) >= 1 && nums[0] == 0 {
		return false
	}
	if len(nums) <= nums[0] {
		return true
	}
	return canJump(nums[nums[0]:])

}
