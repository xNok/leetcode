package addtwosum



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	result := head
	var increase int
	for l1 != nil || l2 != nil || increase != 0 {

	}

	return result
}

func sum(a, b, increase int) (int, int) {
	res := a + b + increase
	digit := res / 10

	return digit, res % 10
}
