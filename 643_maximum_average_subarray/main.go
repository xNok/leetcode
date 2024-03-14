package maximumaveragesubarray

// findMaxAverage
// we don't need to actually bother with the float division until the end
// Go doesn't mix numeral type we need to cast explicitly
// In this case we didn't care about the out if bound window
func findMaxAverage(nums []int, k int) float64 {
	var result, candidate int

	// initialisation
	for i := 0; i < k; i++ {
		result += nums[i]
	}
	candidate = result

	// sliding window
	for i := k; i < len(nums); i++ {
		candidate += nums[i] - nums[i-k]

		if candidate > result {
			result = candidate
		}
	}

	return float64(result) / float64(k)
}
