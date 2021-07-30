package validparentheses

import "testing"

var testCases = map[string]bool{
	`()`:           true,
	`()[]{}`:       true,
	`(]`:           false,
	`([)]`:         false,
	`{[]}`:         true,
	`{}((()))[()]`: true,
}

func TestCases(t *testing.T) {
	for key, val := range testCases {
		t.Run(key, func(t *testing.T) {
			returnVal := isValid(key)
			if returnVal != val {
				t.Errorf("input %v, expected %v, found %v", key, val, returnVal)
			}
		})
	}
}
