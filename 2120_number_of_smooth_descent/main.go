package numberofsmoothdescent

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	ans := int64(n)
	var max int64 = 0
	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1] - 1 {
			max++
			ans += max
		} else {
			max = 0
		}
	}
	return ans
}