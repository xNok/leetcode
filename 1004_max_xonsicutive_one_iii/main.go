package maxxonsicutiveoneiii

// longestOnesflip
//
func longestOnesflip(nums []int, k int) int {
	var right, left, cursor, result int

	for i := 0; i < len(nums); i++ {
		// add the 0 me might be able to flip it
		if nums[i] == 0 {
			cursor++
		}

		// no room left flipping left most
		for cursor > k {
			if nums[left] == 0 {
				cursor--
			}
			left++
		}

		// increasing window to the right
		right++

		result = max(result, right-left)
	}

	return result
}
