package besttimetobuyandsellstock

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name string
		prices []int
		want int
	}{
		{
			name: "example 1",
			prices: []int{7,1,5,3,6,4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
