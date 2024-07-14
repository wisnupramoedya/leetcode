package jump_game_ii

func jump(nums []int) int {
	position, step := 0, 0
	var forward, newJump int
	for {
		if position >= len(nums)-1 {
			break
		}

		forward = nums[position]
		newJump = 0
		for i := 1; i <= forward; i++ {
			if i+position >= len(nums)-1 {
				newJump = i
				break
			}

			if nums[position+newJump]+newJump < nums[position+i]+i {
				newJump = i
			}
		}

		position = position + newJump
		step++
	}
	return step
}
