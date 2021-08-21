/**
 * @param {number[]} nums
 * @return {boolean}
 */

var canJump = function (nums) {
  //  If length of array is less than 2, array is automatically at the end
  if (nums.length < 2) {
    return true;
  }

  // If the first number is 0, no movement is possible, so return false
  if (nums[0] === 0) {
    return false;
  }

  let maxReachable = 0;
  for (let i = 0; i < nums.length - 1; i++) {
    const num = nums[i];
    if (num === 0 && i === maxReachable && i !== nums.length - 1) {
      return false;
    }
    const currentMaxReachable = num + i;
    maxReachable = Math.max(maxReachable, currentMaxReachable);
  }
  return maxReachable >= nums.length - 1;
};
