#! /bin/env python

from typing import List

# for each element
# count = index * element - (sum of previous elements)


def arrayChallenge(n: List[int]) -> List[int]:
    solution = []
    sum = 0
    for i, j in enumerate(n):
        solution.append(i*j - sum)
        sum += j

    return solution


if __name__ == "__main__":
    arr = [1, 2, 2, 3]
    print(arrayChallenge(arr))
