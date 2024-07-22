package removeduplicatefromsortedarrays2

func removeDuplicates(nums []int) int {
	var k, r int
	m := make(map[int]bool)

	for i, n := range nums {
		// duplicate
		if ok := m[n]; ok {
			k++
			continue
		}

		if i > 0 && k > 0 {
			nums[i-k] = n
		}

		m[n] = true
		r++
	}

	return r
}
