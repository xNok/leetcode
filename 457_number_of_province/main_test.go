package numberofprovince

import "testing"

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name: "simple 3 - 2 connected",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "simple 3 - not connected",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.isConnected); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
