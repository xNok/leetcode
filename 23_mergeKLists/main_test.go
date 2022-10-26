package mergeklists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name: "[[1,4,5],[1,3,4],[2,6]]",
			lists: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var lists []*ListNode
			var want *ListNode

			for _, v := range tt.lists {
				lists = append(lists, createListNode(v))
			}

			want = createListNode(tt.want)

			if got := mergeKLists(lists); !reflect.DeepEqual(got, want) {
				t.Errorf("mergeKLists() = %v, want %v", got, want)
			}
		})
	}
}

func Test_createListNode(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want *ListNode
	}{
		{
			name: "3 node",
			list: []int{1, 2, 3},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createListNode(tt.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
