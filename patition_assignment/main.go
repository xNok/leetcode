package partition_assignment

import (
	"sort"
)

func partition(p []int32, p_used []int32) int {
	sort.Slice(p, func(i, j int) bool { return p[i] > p[j] })
	sort.Slice(p_used, func(i, j int) bool { return p_used[i] > p_used[j] })

	var binIndex int
	var bins []int32

	bins = append(bins, p[binIndex])

	for _, item := range p_used {
		placed := false

		for i := range bins {
			if item <= bins[i] {
				bins[i] -= item
				placed = true
				break
			}
		}


		if !placed {
			binIndex++
			bins = append(bins, p[binIndex]-item)
		}
	}

	return len(bins)
}