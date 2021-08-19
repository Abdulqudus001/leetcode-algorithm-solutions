package lettercombinationsofaphonenumber

import (
	"math"
	"strconv"
	"strings"
)

// taking each numbers as having their array i.e
// 7 -> [p, q, r, s] and 8 -> [t, u, v]
// if 7 and 8 were pressed, maxNumber to pass to permutation will be
// (len(largestArrayOfDigits) ^ numberOfDigits) - 1
// len(largestArrayofDigits) = len([p, q, r, s]) = 4
// numberofDigits = 2
// finalNumber of combinations should be the length of all digit arrays multiplied together i.e len([p, q, r, s]) * len([t, u, v]) = 12
// base to use to represent each digit will also be len(largestArrayOfDigits)
//baseRepresentation (xy) where maxNumber will be (4^2)-1 = 15[base 4] i.e possible combinations are
// 00 = [(p), q, r, s] [(t), u, v] = pt
// 01 = [(p), q, r, s] [t, (u), v] = pu
// 02 = [(p), q, r, s] [t, u, (v)] = pv
// 03 (skipped because 3 index does not exist for number 8)
// 10 = [p, (q), r, s] [(t), u, v] = qt
// 11 = [p, (q), r, s] [t, (u), v] = qu
// 12 = [p, (q), r, s] [t, u, (v)] = qv
// 13 (skipped because 3 index does not exist for number 8)
// 20 = [p, q, (r), s] [(t), u, v] = rt
// 21 = [p, q, (r), s] [t, (u), v] = ru
// 22 = [p, q, (r), s] [t, u, (v)] = rv
// 23 (skipped because 3 index does not exist for number 8)
// 30 = [p, q, r, (s)] [(t), u, v] = st
// 31 = [p, q, r, (s)] [t, (u), (v)] = su
// 32 = [p, q, r, (s)] [t, u, (v)] = sv
// 33 (skipped because 3 index does not exist for number 8)

// leaving us with the 12 intended combinations

var keyPad = map[string][]string{
	"":  []string{},
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var arrays [][]string
	longestArrayLength := 1

	if len(digits) <= 1 {
		return keyPad[digits]
	}

	// get arrays of digit into an array of arrays
	// e.g 78 => [[p, q, r, s], [t, u, v]]
	for _, number := range digits {
		values := keyPad[string(number)]
		if len(values) >= longestArrayLength {
			longestArrayLength = len(values)
		}
		arrays = append(arrays, values)
	}
	maxNumber := math.Pow(float64(longestArrayLength), float64(len(digits)))
	return permutate(arrays, int(maxNumber-1), longestArrayLength)
}

func permutate(arrays [][]string, maxNumber int, base int) []string {
	var result []string
	lenOfRepresentation := len(arrays)
	i := 0
	for i <= maxNumber {
		var combination string
		// convert numner to base e.g 5 -> 11 [base 4]
		baseRepresentation := strconv.FormatInt(int64(i), base)
		representation := baseRepresentation
		// pad represention to make sure it reaches required length
		// e.g if 3 digits were pressed then we pad 11 -> 011 and 1 -> 001
		if len(baseRepresentation) < lenOfRepresentation {
			representation = strings.Repeat("0", lenOfRepresentation-len(baseRepresentation)) + baseRepresentation
		}

		for key, val := range representation {
			array := arrays[key]
			// golang craziness to convert a single rune character to string
			stringChar, _ := strconv.Unquote(strconv.QuoteRune(val))
			intChar, _ := strconv.Atoi(stringChar)
			// skip this inner for loop as index i.e intChar does not exist for array
			if intChar > len(array)-1 {
				continue
			}
			combination = combination + array[intChar]
		}
		// the combination must be equal to the number of keys pressed
		// i.e 2 letter and 1 letter combinations are not allowed if 3 buttons
		// were pressed
		if len(combination) == lenOfRepresentation {
			result = append(result, combination)
		}
		i++
	}
	return result
}
