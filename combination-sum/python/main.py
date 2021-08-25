#! /bin/env python
# Time Complexity: O(n x 2^n)
# Space Complexity: n^n
# Uses Multi-Source BFS to construct a Binary Tree from each element
# and compute the sum down the tree until it is less that or equal to the target
# [2, 3, 6, 7]
#         2
#    2         3
#  2   3     3   6
# 2 3 3 6   3 6 6 7
# Calculates the sum of each node at each level and return accordingly

from typing import List 

class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        solution = []

        def generate(index, current_sum, target, __arr):
            __arr.append(candidates[index])

            if current_sum > target:
                return

            if current_sum == target:
                solution.append(list(__arr))
                return

            for i in range(len(candidates[index:])):
                generate(index+i, current_sum + candidates[index+i], target, __arr)
                __arr.pop()

        for i, v in enumerate(candidates):
            generate(i, v, target, [])

        return solution

if __name__ == "__main__":
    sol = Solution()

    arr = [3,5,8]
    target = 11
    solution_arr = sol.combinationSum(arr, target)
    print(solution_arr)

     
