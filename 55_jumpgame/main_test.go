package jumpgame

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		}, {
			name: "example 2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		}, {
			name: "example 3",
			nums: []int{2, 3, 1, 1, 1, 1, 1, 1, 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
