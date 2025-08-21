package partition_assignment

import "testing"

func Test_partition(t *testing.T) {
	type args struct {
		p      []int32
		p_used []int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple",
			args: args{
				p: []int32{10,20,30,40,50},
				p_used: []int32{40,10,5,6,7,8},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.p, tt.args.p_used); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
