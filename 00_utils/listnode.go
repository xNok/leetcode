package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToListNode(in []int) *ListNode {
	var out ListNode
	var head *ListNode

	head = &out

	l := len(in)
	for index, i := range in {
		head.Val = i
		if index == l-1 {
			break
		}
		head.Next = &ListNode{}
		head = head.Next
	}

	return &out
}

func ListNodeToList(in *ListNode) []int {
	var out []int
	var head *ListNode

	head = in

	for {
		out = append(out, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	return out
}
