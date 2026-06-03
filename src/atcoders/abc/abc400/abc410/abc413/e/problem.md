# E - Reverse 2^i (ABC413)

**Contest:** [ABC413](https://atcoder.jp/contests/abc413) — AtCoder Beginner Contest 413  
**Task:** [https://atcoder.jp/contests/abc413/tasks/abc413_e](https://atcoder.jp/contests/abc413/tasks/abc413_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

You are given a permutation `P = (P_0, P_1, ..., P_{2^N-1})` of
`(1, 2, 3, ..., 2^N)`.

You can perform the following operation any number of times, possibly zero:

- Choose non-negative integers `a, b` satisfying
  `0 <= a * 2^b < (a + 1) * 2^b <= 2^N`, and reverse
  `P_{a * 2^b}, P_{a * 2^b + 1}, ..., P_{(a + 1) * 2^b - 1}`.

Here, reversing that subarray means simultaneously replacing

```text
P_{a * 2^b}, P_{a * 2^b + 1}, ..., P_{(a + 1) * 2^b - 1}
```

with

```text
P_{(a + 1) * 2^b - 1}, P_{(a + 1) * 2^b - 2}, ..., P_{a * 2^b}
```

Find the lexicographically smallest permutation obtainable by repeating this
operation.

You are given `T` test cases, so solve each of them.

### ideas
1. let w = pow(2, b), l = a * w
2. w只有20个取值,1,2,4...
3. 但是起点，比如w = 4, 起点只能是1,5,9,...
4. 那么当w = 8的时候，可以交5...12吗？不行，因为不存在这样的a
5. 所以，只能交换相邻的区间
6. 有点知道了
7. 从大到小。没次找一半区间内的最小值

## Solution

Every allowed operation reverses a block whose length is a power of two and whose
starting position is aligned to that length. Because of this alignment, inside any
fixed block of length `2^k`, the only effect we can create at the top level is:

- keep its left half before its right half, or
- swap the two halves.

After that choice, we can recursively do the same thing inside each half.

So the whole problem becomes:

- for each power-of-two block,
- decide which half should appear first,
- then solve both halves independently.

### Greedy choice by the minimum element

To make the final permutation lexicographically smallest, the first element of the
current block should be as small as possible.

Consider a block `[l, r)` of length at least `2`. Let `mid = (l + r) / 2`.
Find the position `mn` of the minimum value in this block.

- If `mn < mid`, then the minimum lies in the left half, so the left half must come
  first.
- If `mn >= mid`, then the minimum lies in the right half, so the right half must
  come first.

This choice is forced: whichever half contains the smallest number must be placed
first, otherwise the first element of the block could be made smaller by swapping
the halves.

After fixing which half comes first, we recursively solve the first half and then
the second half in that order.

### Why the recursion is correct

The only freedom at one block is the order of its two equal halves. Once that order
is fixed, operations inside one half never mix elements with the other half, so the
two subproblems are independent.

Lexicographic order also fits this recursion perfectly:

- first minimize the first half that appears,
- then, if that is fixed, minimize the second half.

Therefore, choosing the half containing the minimum element first is optimal, and
recursing gives the lexicographically smallest reachable permutation.

### How the code works

The function `f(l, r, l1, r1)` processes original subarray `arr[l:r]` and writes its
optimal arrangement into `res[l1:r1]`.

- Base case: when the segment length is `1`, copy that value into `res`.
- Otherwise, scan `arr[l:r]` to find the index of the minimum value.
- If the minimum is in the left half, recursively place the left half into the first
  half of `res[l1:r1]`, then place the right half into the second half.
- Otherwise do the two recursive calls in the opposite order.

That exactly simulates the best sequence of allowed half-swaps.

### Complexity

At each recursion depth, every element is scanned once as part of exactly one block.
There are `log N` depths where `N` is the permutation length, so the total time is:

```text
O(N log N)
```

The extra memory is `O(N)` for the answer array and recursion.
