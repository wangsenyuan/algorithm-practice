# G. Greedy Subsequences

[Problem link](https://codeforces.com/problemset/problem/1132/G)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

For an array `c`, a **greedy subsequence** is a sequence of indices
`p_1 < p_2 < ... < p_l` such that for each `i`, `p_{i+1}` is the **smallest**
index greater than `p_i` with `c[p_{i+1}] > c[p_i]`.

You are given an array `a_1, ..., a_n`. For each contiguous subsegment of length
`k`, find the length of its **longest** greedy subsequence.

## Constraints

- `1 <= k <= n <= 10^6`
- `1 <= a_i <= n`

## Input

```text
n k
a_1 a_2 ... a_n
```

## Output

Print `n - k + 1` integers — the answers for subsegments `a[1..k]`,
`a[2..k+1]`, ..., `a[n-k+1..n]`.

## Example 1

```text
Input
6 4
1 5 2 5 3 6

Output
2 2 3
```

- `[1, 5, 2, 5]` — longest greedy length `2` (e.g. `1,5` or `2,5`)
- `[5, 2, 5, 3]` — length `2` (`2,5`)
- `[2, 5, 3, 6]` — length `3` (`2,5,6`)

## Example 2

```text
Input
7 6
4 5 2 5 3 6 6

Output
3 3
```

- `[4, 5, 2, 5, 3, 6]` — length `3`
- `[5, 2, 5, 3, 6, 6]` — length `3`


### ideas
1. p[1] 确定的时候, 后续位置都确定了, 
2. 先构造一棵树(从右往左), fa[i] = j
3. 然后构造欧拉序列, 这样, 当一个点r移除窗口的时候, 所有的r子树中的dp - 1 (这样就可以用 segment tree 的 lazy range update)
4. 加入一个点l的时候, 这个是点 get (query fa[l])
5. 应该是可以的