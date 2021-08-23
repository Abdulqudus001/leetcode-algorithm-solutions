/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePaths = function (m, n) {
  const myArray = Array(m)
    .fill(0)
    .map(() => Array(n).fill(0));

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i == 0 || j == 0) {
        myArray[i][j] = 1;
      } else {
        myArray[i][j] =
          parseInt(myArray[i - 1][j]) + parseInt(myArray[i][j - 1]);
      }
    }
  }

  return myArray[m - 1][n - 1];
};
