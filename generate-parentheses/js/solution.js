/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function (n) {
  let result = [];

  const build = (currentString, left, right) => {
    if (left === 0 && right === 0) {
      result.push(currentString);
      return;
    }

    if (left > 0) {
      build(currentString + '(', left - 1, right);
    }

    if (right > 0 && right > left) {
      build(currentString + ')', left, right - 1);
    }
  };

  build('', n, n);

  return result;
};
