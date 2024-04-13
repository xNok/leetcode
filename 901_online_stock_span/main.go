package onlinestockspan

type StockSpanner struct {
	// stack of current price, max span
	Stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	n := len(this.Stack)
	max := 1
	// look at the stack in rever order
	for i := n - 1; i > 0; i-- {
		prev := this.Stack[i]
		if prev[0] <= price {
			// add the pre calculated max
			max += prev[1]
			// remove this one from the stack the new one is better
			this.Stack = this.Stack[:i]
		}
	}
	// add the new price to the stack
	this.Stack = append(this.Stack, [2]int{price, max})
	return max
}
