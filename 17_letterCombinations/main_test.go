package lettercombinations

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "23",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		}, {
			name:   "none",
			digits: "",
			want:   []string{},
		}, {
			name:   "2",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digitConvert(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want []string
	}{
		{
			name: "2",
			b:    byte(50),
			want: []string{"a", "b", "c"},
		}, {
			name: "3",
			b:    byte(51),
			want: []string{"d", "e", "f"},
		}, {
			name: "4",
			b:    byte(52),
			want: []string{"g", "h", "i"},
		}, {
			name: "5",
			b:    byte(53),
			want: []string{"j", "k", "l"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitConvert(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digitConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
