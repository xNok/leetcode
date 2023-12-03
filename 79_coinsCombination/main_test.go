package coinsCombinations

import "testing"

func Test_coinConbination(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple1",
			args: args{
				a: 1,
				b: 1,
				c: 1,
				x: 550,
			},
			want: 1,
		},
		{
			name: "simple2",
			args: args{
				a: 2,
				b: 2,
				c: 2,
				x: 100,
			},
			want: 2,
		}, {
			name: "simple3",
			args: args{
				a: 5,
				b: 1,
				c: 0,
				x: 150,
			},
			want: 0,
		}, {
			name: "hard1",
			args: args{
				a: 30,
				b: 40,
				c: 50,
				x: 6000,
			},
			want: 213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinConbination(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("coinConbination() = %v, want %v", got, tt.want)
			}
		})
	}
}
