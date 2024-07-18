package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListToListNode(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want *ListNode
	}{
		{
			name: "",
			in:   []int{1, 2, 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListToListNode(tt.in)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestListNodeToList(t *testing.T) {
	tests := []struct {
		name string
		in   *ListNode
		want []int
	}{
		{
			name: "123",
			in: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListNodeToList(tt.in)

			require.Equal(t, tt.want, got)
		})
	}
}
