package coinsCombinations

func coinConbination(a, b, c, x int) int {
	ans := 0
	backtrack(x, 0, []int{a, b, c}, []int{500, 100, 50}, &ans)
	return ans
}

func backtrack(target, idx int, coins, coinsValue []int, ans *int) {
	// End condition we are found a solution
	if target == 0 {
		*ans += 1
		return
	}

	// the branch is dead
	if target < 0 || idx == len(coins) {
		return
	}

	// Search all cases
	for i := idx; i < len(coins); i++ {
		if coins[i] == 0 {
			continue
		}
		// Save the current target
		coins[i] -= 1
		backtrack(target-coinsValue[i], i, coins, coinsValue, ans)
		// Revert the condtion
		coins[i] += 1
	}
}
