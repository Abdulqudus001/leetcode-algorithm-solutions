package jumpgame

import (
	"strconv"
	"testing"
)

type TestCase struct {
	array  []int
	result bool
}

var testCases = []TestCase{
	TestCase{
		array:  []int{2, 3, 1, 1, 4},
		result: true,
	},
	TestCase{
		array:  []int{3, 3, 1, 0, 4},
		result: false,
	},
	TestCase{
		array:  []int{2, 3, 1, 3, 4, 1, 0, 1},
		result: false,
	},
}

func TestCases(t *testing.T) {
	for key, val := range testCases {
		t.Run(strconv.Itoa(key), func(t *testing.T) {
			returnVal := canJump(val.array)
			if returnVal != val.result {
				t.Errorf("input %v, expected %v, found %v", val.array, val.result, returnVal)
			}
		})
	}
}
