package stockpricefluctuation

import "slices"

type StockPrice struct {
	Stack   [][2]int
	current int
	max     int
	min     int
	len     int
}

func Constructor() StockPrice {
	return StockPrice{
		Stack:   [][2]int{},
		current: -1,
		max:     -1,
		min:     -1,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.len == 0 {
		this.current = price
		this.max = price
		this.min = price
		this.Stack = [][2]int{{timestamp, price}}
		this.len = 1
		return
	}

	if price > this.max {
		this.max = price
	}

	if price < this.min {
		this.min = price
	}

	// Add this to the stack
	for i := this.len - 1; i >= 0; i-- {
		prev := this.Stack[i]

		// compare timestamps
		// prev=10 timestamp=11
		if prev[0] < timestamp && i == this.len-1 {
			// happy case we simply update
			this.current = price
			this.Stack = append(this.Stack, [2]int{timestamp, price})
			this.len++
			return
		}

		// in place update
		if prev[0] == timestamp {
			prev[1] = price
			return
		}

		// prev=11 timestamp=10 - insert
		if prev[0] > timestamp {
			this.Stack = slices.Insert(this.Stack, i, [2]int{timestamp, price})
			this.len++
			return
		}
	}

}

func (this *StockPrice) Current() int {
	return this.current
}

func (this *StockPrice) Maximum() int {
	return this.max
}

func (this *StockPrice) Minimum() int {
	return this.min
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
