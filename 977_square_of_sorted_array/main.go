package squareofsortedarray

// sortedSquares
// No power function in go for intergers (simply use multiplication)
// The pointer move one at the time for the candidate selected
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j, n := 0, len(nums)-1, len(nums)-1

	for i < j {
		i2, j2 := nums[i]*nums[i], nums[j]*nums[j]

		if i2 > j2 {
			result[n] = i2
			i++
		} else {
			result[n] = j2
			j--
		}

		n--
	}

	// the middle number need to be added back alone (i=j)
	result[0] = nums[i] * nums[i]

	return result
}
