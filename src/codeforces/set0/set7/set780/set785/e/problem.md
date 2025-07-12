# Anton and Permutations

## Problem Description

Anton likes permutations, especially he likes to permute their elements. Note that a permutation of n elements is a sequence of numbers {a₁, a₂, ..., aₙ}, in which every number from 1 to n appears exactly once.

One day Anton got a new permutation and started to play with it. He does the following operation q times: he takes two elements of the permutation and swaps these elements. After each operation he asks his friend Vanya, how many inversions there are in the new permutation. The number of inversions in a permutation is the number of distinct pairs (i, j) such that 1 ≤ i < j ≤ n and aᵢ > aⱼ.

Vanya is tired of answering Anton's silly questions. So he asked you to write a program that would answer these questions instead of him.

Initially Anton's permutation was {1, 2, ..., n}, that is aᵢ = i for all i such that 1 ≤ i ≤ n.

## Input

The first line of the input contains two integers n and q (1 ≤ n ≤ 200,000, 1 ≤ q ≤ 50,000) — the length of the permutation and the number of operations that Anton does.

Each of the following q lines of the input contains two integers lᵢ and rᵢ (1 ≤ lᵢ, rᵢ ≤ n) — the indices of elements that Anton swaps during the i-th operation. Note that indices of elements that Anton swaps during the i-th operation can coincide. Elements in the permutation are numbered starting with one.

## Output

Output q lines. The i-th line of the output is the number of inversions in Anton's permutation after the i-th operation.


### ideas
1. 假设交换前已知，现在进行交换(i, j)
2. i < j
3. 那么对于a[i]来说，需要知道在区间i...j中间比它大的数，假设有w，那么从a[i]的角度看，这个变化 = -w + (j - i - w) = (j - i - 2 * w)
4. 对于a[j]来说， 同样假设在区间(i...j)中间比a[j]大的数,有w, w - (j - i - w) = 2 * w - (j - i -1) 
5. -2 * w + 2 * v = 2 * (v - w)
6. 所以关键是知道在区间内，比某个数大的数
7. kdtree？