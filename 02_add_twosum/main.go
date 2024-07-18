package addtwosum

import utils "github.com/mekitmedia/leetcode/00_utils"

type ListNode = utils.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	result := head
	var carry int
	for {
		head.Val, carry = sum(l1.Val, l2.Val, head.Val)

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil && carry == 0{
			break
		}

		head.Next = &ListNode{}
		head = head.Next
		head.Val += carry

		if l1 == nil {
			l1 = &ListNode{}
		}

		if l2 == nil {
			l2 = &ListNode{}
		}
	}

	return result
}

func sum(a, b, c int) (int, int) {
	res := a + b + c
	return res % 10, res / 10
}
