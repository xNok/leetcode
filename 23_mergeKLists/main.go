package mergeklists

import (
	"container/heap"

	utils "github.com/mekitmedia/leetcode/00_utils"
)

func mergeKLists(lists []*ListNode) *ListNode {
	var head, pointer *ListNode = nil, nil
	if len(lists) == 0 {
		return head
	}
	// Create a heap of [i, list.Val] where lists[i] = list
	h := &utils.IntHeap{}
	heap.Init(h)
	for i, list := range lists {
		if list != nil {
			heap.Push(h, [2]int{i, list.Val})
		}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) setNext(next *ListNode) *ListNode {
	n.Next = next
	return next
}

func createListNode(list []int) *ListNode {

	if len(list) == 0 {
		return &ListNode{}
	}

	var next *ListNode

	node := ListNode{
		Val: list[0],
	}

	next = &node

	for _, v := range list[1:] {
		next = next.setNext(&ListNode{
			Val: v,
		})
	}

	return &node
}
