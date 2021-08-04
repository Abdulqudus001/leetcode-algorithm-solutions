package mergetwosortedlists

import (
	"strconv"
	"testing"
)

var testCases = [][]*ListNode{
	{
		makeListNodes([]int{1, 2, 4}),
		makeListNodes([]int{1, 3, 4}),
		makeListNodes([]int{1, 1, 2, 3, 4, 4}),
	},
	{
		makeListNodes([]int{}),
		makeListNodes([]int{}),
		makeListNodes([]int{}),
	},
	{
		makeListNodes([]int{}),
		makeListNodes([]int{0}),
		makeListNodes([]int{0}),
	},
	{
		makeListNodes([]int{5}),
		makeListNodes([]int{0, 2, 4}),
		makeListNodes([]int{0, 2, 4, 5}),
	},
}

func TestCases(t *testing.T) {
	for key, val := range testCases {
		t.Run(strconv.Itoa(key), func(t *testing.T) {
			l1, l2, l3 := val[0], val[1], val[2]
			returnVal := mergeTwoLists(l1, l2)
			res := l3.getList()
			for index, val := range returnVal.getList() {
				if val != res[index] {
					t.Errorf("input %v, expected %v, found %v", key, val, returnVal)
				}
			}
		})
	}
}
