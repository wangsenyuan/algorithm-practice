# A. MAD Interactive Problem (Codeforces 2159A)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2159/A](https://codeforces.com/problemset/problem/2159/A)

## Problem Statement

This problem was originally interactive.

There is a hidden sequence `a[1], a[2], ..., a[2n]`. Every integer from `1` to
`n` appears exactly twice in the sequence.

The allowed query is:

- choose `k` distinct indices `j[1], j[2], ..., j[k]`;
- the jury returns `MAD([a[j[1]], a[j[2]], ..., a[j[k]]])`.

For a sequence, `MAD` means the largest value that appears at least twice. If no
value appears at least twice, `MAD` is `0`.

Examples:

- `MAD([1, 2, 1]) = 1`;
- `MAD([2, 2, 3, 3]) = 3`;
- `MAD([1, 2, 3, 4]) = 0`.

The original goal is to recover the whole hidden sequence using at most `3n`
queries.

## Hack Format

For hacks, the hidden sequence is given directly in the input.

The first line contains an integer `t` (`1 <= t <= 3000`) — the number of test
cases.

For each test case:

- The first line contains an integer `n` (`2 <= n <= 300`).
- The second line contains `2n` integers `a[1], a[2], ..., a[2n]`
  (`1 <= a[i] <= n`).

Each number from `1` to `n` appears exactly twice.

It is guaranteed that the sum of `n^2` over all test cases does not exceed
`10^5`.

## Original Interaction

In the original interactive version, after reading `n`, the solution could print
queries of the form:

```text
? k j1 j2 ... jk
```

where all selected indices must be distinct and `1 <= k <= 2n`.

After each query, the jury returned one integer: the `MAD` value of the selected
subsequence.

Once the sequence was determined, the solution printed:

```text
! a1 a2 ... a(2n)
```

The final answer did not count toward the query limit.

## Example

The original page shows an interaction. The corresponding hack input is:

```text
2
2
2 2 1 1
2
1 2 1 2
```

One possible interaction for the same hidden sequences is:

### Input

```text
2
2

2

0

1

2

0

1

1
```

### Output

```text
? 2 2 1

? 2 1 3
? 3 1 3 4

! 2 2 1 1

? 2 1 2

? 2 1 3

? 3 1 3 4

! 1 2 1 2
```

## Note

For the first hidden sequence `a = [2, 2, 1, 1]`:

- Querying indices `2, 1` selects `[2, 2]`, so the answer is `2`.
- Querying indices `1, 3` selects `[2, 1]`, so the answer is `0`.
- Querying indices `1, 3, 4` selects `[2, 1, 1]`, so the answer is `1`.

The example interaction is only illustrative; it does not imply that the hidden
sequence is uniquely determined by only those displayed queries.


### ideas
1. 在一个长度为n+1的区间里面，至少一个数是重复的，假设f(1) = ask(1...n+1) = x
2. f(2) = ask(2, n+2) = y
3. 如果 x != y, 可以立马得到 a[1] = x, 或者 a[n+2] = y (看哪个大)
4. 如果 x = y? 这样扫一遍过去，就可以找到n的两个位置？
5. 它没有说需要连续。
6. 能在差不多n次，找到n的两个位置