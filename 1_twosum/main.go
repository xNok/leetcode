package twosum

func twoSum(nums []int, target int) []int {
	// save what we already know
	lookup := make(map[int]int)

	for index, v := range nums {
		// what number do we need
		need := target - v
		if val, ok := lookup[need]; ok {
			return []int{val, index}
		} else {
			lookup[v] = index
		}
	}

	return []int{}
}
