package onlinestockspan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockSpanner_Next(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			input: []int{100, 80, 60, 70, 60, 75, 85},
			want:  []int{1, 1, 1, 2, 1, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			var result []int

			for _, price := range tt.input {
				result = append(result, obj.Next(price))
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
