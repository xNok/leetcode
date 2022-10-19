package inttoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "3 -> III",
			num:  3,
			want: "III",
		}, {
			name: "58 -> LVIII",
			num:  58,
			want: "LVIII",
		}, {
			name: "1994 -> MCMXCIV",
			num:  1994,
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
