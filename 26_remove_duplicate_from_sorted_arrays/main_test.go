package removeduplicatefromsortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantNums []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2, 0},
		},
		{
			name:     "example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, got)
			for i := 0; i < got-1; i++ {
				assert.Equal(t, tt.wantNums[i], tt.nums[i])
			}
		})
	}
}
