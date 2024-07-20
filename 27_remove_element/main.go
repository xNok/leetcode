package removeelement

func removeElement(nums []int, val int) int {
	var p int
	l := len(nums)

	for i, num := range nums {
		if num == val {
			p++
			continue
		}

		if i > 0 && p > 0 {
			nums[i-p] = nums[i]
		}
	}

	return l - p
}
