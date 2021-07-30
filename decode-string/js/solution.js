// Find single pair of square brackets
const findPair = str => {
	let indexOfOpening = str.indexOf('[');
	let indexOfClosing = -1;
	let openingBracket = 0;
	for (let i = 0; i < str.length; i++) {
		const char = str[i];
		if (char === '[') {
			openingBracket += 1;
		} else if (char === ']') {
			openingBracket -= 1;
		}

		if (openingBracket === 0 && i > indexOfOpening) {
			indexOfClosing = i;
			return [indexOfOpening, indexOfClosing];
		}
	}

}

// Expand string inside pair or square brackets
const expandString = str => {
	const [start, end] = findPair(str);
	const numOfRepeat = str[start - 1];
	const strToRepeat = str.slice(start + 1, end);
	const arr = str.split('');
	arr.splice(start - 1, end - start + 2, strToRepeat.repeat(numOfRepeat));
	return arr.join('');
}

// Entry point
const decode = str => {
	if (str.indexOf('[') < 0) {
		return str;
	}
	return decode(expandString(str));
}
