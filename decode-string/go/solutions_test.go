package decodestring

import "testing"

var testCases = map[string]string{
	`3[a]2[bc]`:     "aaabcbc",
	`3[a2[c]]`:      "accaccacc",
	`2[abc]3[cd]ef`: "abcabccdcdcdef",
	`abc3[cd]xyz`:   "abccdcdcdxyz",
}

func TestCases(t *testing.T) {
	for key, val := range testCases {
		t.Run(key, func(t *testing.T) {
			returnVal := decodeString(key)
			if returnVal != val {
				t.Errorf("input %v, expected %v, found %v", key, val, returnVal)
			}
		})
	}
}
