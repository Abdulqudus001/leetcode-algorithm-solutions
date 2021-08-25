/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function (candidates, target) {
  const result = [];

  const generateCombination = (resultArray, index, sum) => {
    if (sum < target && index >= candidates.length) {
      return;
    }

    if (sum > target) {
      return;
    }

    if (sum === target) {
      result.push(resultArray);
      return;
    }

    for (let i = index; i < candidates.length; i++) {
      generateCombination(
        [...resultArray, candidates[i]],
        i,
        sum + candidates[i]
      );
    }
  };

  for (let i = 0; i < candidates.length; i++) {
    generateCombination([candidates[i]], i, candidates[i]);
  }

  return result;
};
