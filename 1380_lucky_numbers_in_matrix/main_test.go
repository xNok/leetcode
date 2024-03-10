package luckynumbersinmatrix

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "1",
			matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			want:   []int{15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luckyNumbers(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
