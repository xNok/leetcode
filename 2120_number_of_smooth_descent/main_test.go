package numberofsmoothdescent

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int64
	}{
		{
			name:   "test 1",
			prices: []int{3, 2, 1, 4},
			want:   7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
