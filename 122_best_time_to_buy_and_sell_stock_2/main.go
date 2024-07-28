package besttimetobuyandsellstock2

func maxProfit(prices []int) int {
	profitSum := 0
	prevNumber := prices[0]

	// We can abuse the condition that we can only hold for a day
	// This means a long as we can mae a profit on a day we should
	// there is not reason we should not an this would give us the optimal solution
	for i := 1; i < len(prices); i++ {
		if prices[i]-prevNumber > 0 {
			profitSum += prices[i] - prevNumber
		}
		prevNumber = prices[i]
	}

	return profitSum
}
