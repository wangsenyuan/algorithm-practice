# Pikachu's Subsequence Problem

## Problem Description

Pikachu had an array with him. He wrote down all the non-empty subsequences of the array on paper. Note that an array of size n has $2^n - 1$ non-empty subsequences in it.

Pikachu being mischievous as he always is, removed all the subsequences in which:
```
Maximum_element_of_the_subsequence - Minimum_element_of_subsequence ≥ d
```

Pikachu was finally left with X subsequences.

However, he lost the initial array he had, and now is in serious trouble. He still remembers the numbers X and d. He now wants you to construct any such array which will satisfy the above conditions. All the numbers in the final array should be positive integers less than $10^{18}$.

**Note:** The number of elements in the output array should not be more than $10^4$. If no answer is possible, print -1.

## Input

The only line of input consists of two space separated integers X and d (1 ≤ X, d ≤ $10^9$).

## Output

Output should consist of two lines:

1. **First line:** A single integer n (1 ≤ n ≤ 10,000) — the number of integers in the final array.
2. **Second line:** n space separated integers — a₁, a₂, ..., aₙ (1 ≤ aᵢ < $10^{18}$).

If there is no answer, print a single integer -1. If there are multiple answers, print any of them.


### ideas
1. $2^{h}$ > X
2. 如果正好有个h得到X个非空子集，那么找到这样h个数（最大值 - 最小值 < d)
3. 但是如果不能完全得到这样的集合，那么先找出一个h ($2^h$ - 1 < X)
4. 这个集合内的数，满足d的条件， 剩余X1个集合, 再找一个新的集合 最小值 == set1 + d
5. 上面那个h还受到d的控制