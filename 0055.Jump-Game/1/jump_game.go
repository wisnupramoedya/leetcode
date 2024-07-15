package jump_game

func canJump(nums []int) bool {
	reach := 0
	for i := 0; i <= reach; i++ {
		reach = max(reach, nums[i]+i)
		if reach >= len(nums)-1 {
			return true
		}
	}
	return false
}
