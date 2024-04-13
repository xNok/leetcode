package stockpricefluctuation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockPrice_Update(t *testing.T) {
	type args struct {
		timestamp int
		price     int
	}
	tests := []struct {
		name   string
		fields StockPrice
		args   args
		want   StockPrice
	}{
		{
			name: "add to tail",
			fields: StockPrice{
				current: -1,
				max:     -1,
				min:     -1,
				Stack:   [][2]int{},
			},
			args: args{1, 10},
			want: StockPrice{
				current: 10,
				max:     10,
				min:     10,
				Stack:   [][2]int{{1, 10}},
				len:     1,
			},
		},
		{
			name: "add to tail",
			fields: StockPrice{
				current: 10,
				max:     10,
				min:     10,
				Stack:   [][2]int{{1, 10}},
				len:     1,
			},
			args: args{2, 5},
			want: StockPrice{
				current: 5,
				max:     10,
				min:     5,
				Stack:   [][2]int{{1, 10}, {2, 5}},
				len:     2,
			},
		}, {
			name: "update in place",
			fields: StockPrice{
				current: 5,
				max:     10,
				min:     5,
				Stack:   [][2]int{{1, 10}, {2, 5}},
				len:     2,
			},
			args: args{1, 3},
			want: StockPrice{
				current: 5,
				max:     5,
				min:     3,
				Stack:   [][2]int{{1, 3}, {2, 5}},
				len:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Update(tt.args.timestamp, tt.args.price)

			assert.Equal(t, tt.want, tt.fields)
		})
	}
}
