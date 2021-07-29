package validparentheses

func isValid(s string) bool {
	sRune := []rune(s)
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	openingBraces := []rune{}
	closingBraces := []rune{}
	openingBracesOrder := []rune{}
	start := sRune[0]
	length := len(sRune)
	end := sRune[len(sRune)-1]

	for key, val := range match {
		closingBraces = append(closingBraces, key)
		openingBraces = append(openingBraces, val)
	}

	if contains(openingBraces, end) || contains(closingBraces, start) {
		return false
	}

	if length%2 != 0 {
		return false
	}

	for _, char := range sRune {
		if contains(openingBraces, char) {
			openingBracesOrder = append(openingBracesOrder, char)
		} else {
			lastOpening := openingBracesOrder[len(openingBracesOrder)-1]

			if match[char] == lastOpening {
				openingBracesOrder = openingBracesOrder[:len(openingBracesOrder)-1]
			} else {
				return false
			}

		}
	}
	if len(openingBracesOrder) == 0 {
		return true
	}

	return false
}

func contains(arr []rune, str rune) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
