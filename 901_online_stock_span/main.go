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
	var max int = 1
	for i := n - 1; i >= 0; i-- {
		prices := this.Stack[i]
		if prices[0] <= price {
			max += this.Stack[i][1]
			// remove what we calculated from the stack
			this.Stack = this.Stack[:i]
		}
	}
	// the new day we just calculated
	this.Stack = append(this.Stack, [2]int{price, max})
	return max
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
