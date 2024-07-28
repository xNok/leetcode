package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	// Loop throught the list and access two things:
	// 1. Can we buy cheaper
	// 2. Can we increase profit
	for i := 1; i < len(prices); i++ {
		// we look for the lowest possibly bying price
		if prices[i] < minPrice {
			minPrice = prices[i]
			// otherwise we look id selling at the current price increase profit
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
