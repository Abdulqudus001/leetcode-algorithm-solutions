package lettercombinationsofaphonenumber

import (
	"strconv"
	"testing"
)

type TestCase struct {
	digits string
	output []string
}

var testCases = []TestCase{
	TestCase{
		digits: "23",
		output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	TestCase{
		digits: "",
		output: []string{},
	},
	TestCase{
		digits: "2",
		output: []string{"a", "b", "c"},
	},
	TestCase{
		digits: "78",
		output: []string{"pt", "pu", "pv", "qt", "qu", "qv", "rt", "ru", "rv", "st", "su", "sv"},
	},
	TestCase{
		digits: "234",
		output: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
	},
}

func TestCases(t *testing.T) {
	for key, val := range testCases {
		t.Run(strconv.Itoa(key), func(t *testing.T) {
			returnVal := letterCombinations(val.digits)
			if !Equals(val.output, returnVal) {
				t.Errorf("input %v, expected %#v, found %#v", val.digits, val.output, returnVal)
			}
		})
	}
}

func Equals(first []string, second []string) bool {
	if len(first) == len(second) {
		for _, val := range first {
			if !Contains(second, val) {
				return false
			}
		}
		return true
	}
	return false
}

func Contains(arr []string, key string) bool {
	for _, val := range arr {
		if val == key {
			return true
		}
	}
	return false
}
