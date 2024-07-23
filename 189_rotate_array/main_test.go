package rotatearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    0,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		}, {
			name: "example 3",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		}, {
			name: "example 4",
			args: args{
				nums: []int{1, 2},
				k:    1,
			},
			want: []int{2, 1},
		}, {
			name: "example 5",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		}, {
			name: "example 6",
			args: args{
				nums: []int{1, 2, 3},
				k:    4,
			},
			want: []int{3, 1, 2},
		}, {
			name: "example 7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: []int{6, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
