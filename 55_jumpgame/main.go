package jumpgame

func canJump(nums []int) bool {
	index := make(map[int]bool)
	return dfs(0, index, nums)
}

func dfs(index int, memory map[int]bool, nums []int) bool {
	goal := len(nums) - 1
	start := nums[0]

	if val, ok := memory[index]; ok {
		return val
	}

	if start >= goal {
		return true
	}

	if goal == 0 {
		return true
	}

	if start == 0 {
		return false
	}

	for i := start; i >= 1; i-- {
		d := i + nums[i]
		if d >= goal {
			return true
		}

		memory[index+i] = dfs(index+i, memory, nums[i:])
		if memory[index+i] {
			return true
		}
	}

	return false
}
