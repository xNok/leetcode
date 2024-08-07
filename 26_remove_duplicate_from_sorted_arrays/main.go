package removeduplicatefromsortedarrays

func removeDuplicates(nums []int) int {
	var k, r int
	m := make(map[int]int)

	for i, n := range nums {
		// duplicate
		if ok := m[n]; ok > 1 {
			k++
			continue
		}

		if i > 0 && k > 0 {
			nums[i-k] = n
		}

		m[n]++
		r++
	}

	return r
}
