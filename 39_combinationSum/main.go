package combinationsum

func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    backtrack(candidates, 0, []int{}, target, &res)
    return res
}

func backtrack(candidates []int, idx int, current []int, target int, res *[][]int) {
	// End condition
    if target < 0 {
        return
    }
    if target == 0 {
		*res = append(*res, append([]int{}, current...))
		return
	}

	// We start a the index idx
	for i := idx; i < len(candidates); i++ {
        // Do the operation
        current = append(current, candidates[i])
        // Look deeper
        backtrack(candidates, i, current, target-candidates[i], res)
        // Revert the operation
        current = current[:len(current)-1]
	}
}