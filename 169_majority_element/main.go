package majorityelement

func majorityElement(nums []int) int {
	var r int
	var rmax int
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	for key, val := range m {
		if val > rmax {
			rmax = val
			r = key
		}
	}

	return r
}
