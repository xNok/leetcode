package permutations

func permute(nums []int) [][]int {
	res := [][]int{}

	backtrack(nums, 0, &res)
	return res
}

func backtrack(nums []int, idx int, res *[][]int) {
	// End condition
	if idx == len(nums) {
		p := make([]int, len(nums))
		copy(p, nums)
		(*res) = append(*res, p)
		return
	}

	// We start a the index idx
	for i := idx; i < len(nums); i++ {
		// No need to swap the first element with itself
		if idx == i {
			backtrack(nums, idx+1, res)
		} else {
			// Do the operation
			nums[i], nums[idx] = nums[idx], nums[i]
			// Look deeper
			backtrack(nums, idx+1, res)
			// Revert the operation
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
}
