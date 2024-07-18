package addtwosum

import (
	"testing"

	utils "github.com/mekitmedia/leetcode/00_utils"
	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "exp 1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		}, {
			name: "exp 1",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := utils.ListToListNode(tt.l1)
			l2 := utils.ListToListNode(tt.l2)

			got := addTwoNumbers(l1, l2)
			gotList := utils.ListNodeToList(got)

			require.Equal(t, tt.want, gotList)
		})
	}
}
