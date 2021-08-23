const generatePairs = (base, combination, words) => {
  const result = [];

  for (let i = 0; i <= combination; i++) {
    const indexInBase = i.toString(base);
    const currentCombo = indexInBase.toString().padStart(words.length, '0');

    const currentWordCombo = currentCombo
      .split('')
      .map((el, index) => words[index][parseInt(el)])
      .join('');

    if (currentWordCombo.length === currentCombo.length) {
      result.push(currentWordCombo);
    }
  }

  return result;
};
/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function (digits) {
  const lookup = {
    2: 'abc',
    3: 'def',
    4: 'ghi',
    5: 'jkl',
    6: 'mno',
    7: 'pqrs',
    8: 'tuv',
    9: 'wxyz',
  };
  const words = digits.split('').map((el) => lookup[el]);

  const base = Math.max(...words.map((el) => el.length));

  const combinations = parseInt(
    Array(digits.length)
      .fill(base - 1)
      .join(''),
    base
  );

  const res = generatePairs(base, combinations, words);

  return res;
};
