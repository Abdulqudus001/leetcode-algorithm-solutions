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
			returnResult := mergeTwoLists(l1, l2).getList()
			expectedResult := l3.getList()
                        // check equality between elements of expected result and returned result
			for index, val := range returnResult {
				if val != expectedResult[index] {
					t.Errorf("input %v, expected %v, found %v", key, val, returnResult)
				}
			}
		})
	}
}
