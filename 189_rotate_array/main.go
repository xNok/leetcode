package rotatearray

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	l := len(nums)
	if l == 1 || l == k {
		return
	}

	k = k % l
	p := k % 2
	buff := append(nums[l-k:], nums[:l-k+p]...)

	for i, _ := range nums {
		nums[i] = buff[i]
	}
}
