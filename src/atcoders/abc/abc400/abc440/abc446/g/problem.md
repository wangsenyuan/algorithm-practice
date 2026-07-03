# G - 221 Subsequence

[Problem link](https://atcoder.jp/contests/abc446/tasks/abc446_g)

**Contest:** [AtCoder Beginner Contest 446](https://atcoder.jp/contests/abc446)

time limit: 2 sec

memory limit: 1024 MiB

score: 600 points

For a sequence `X = (X_1, ..., X_n)` of positive integers, `X` is a **221 sequence** if and only if,
in its run-length encoding, every run has **length equal to its value**.

More formally, for any `(l, r)` with `1 <= l <= r <= n`, if all of the following hold:

- `l = 1` or (`l >= 2` and `X_{l-1} != X_l`)
- `r = n` or (`r <= n-1` and `X_{r+1} != X_r`)
- `X_l = X_{l+1} = ... = X_r`

then `(r - l + 1) = X_l`.

Examples: `(2,2,3,3,3,1,2,2)` is a 221 sequence; `(1,1)` and `(4,4,1,4,4)` are not.

You are given a length-`N` sequence `A = (A_1, ..., A_N)` of positive integers. Count the number of
non-empty (not necessarily contiguous) subsequences of `A` that are 221 sequences, modulo
`998244353`. Equal sequences taken from different positions count once.

## Constraints

- `1 <= N <= 500000`
- `1 <= A_i <= N`
- All input values are integers

## Input

```text
N
A_1 A_2 ... A_N
```

## Output

Print the answer modulo `998244353`.

## Sample Input 1

```text
8
2 1 2 1 1 2 7 2
```

## Sample Output 1

```text
5
```

The five distinct 221 subsequences are:

- `(1)`
- `(1,2,2)`
- `(2,2)`
- `(2,2,1)`
- `(2,2,1,2,2)`

## Sample Input 2

```text
5
2 3 4 5 4
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
20
2 2 3 1 1 4 1 4 1 4 2 4 1 2 1 4 4 1 1 4
```

## Sample Output 3

```text
15
```

## Solution

Equal subsequence values chosen from different positions are counted only once. To avoid duplicates,
represent every value sequence by its greedy occurrence: scan the wanted values from left to right,
and for each value choose the earliest possible index after the previous one. This gives one canonical
index sequence for each distinct subsequence value sequence.

For a valid 221 sequence, each block is `v` repeated exactly `v` times. In its greedy occurrence, it
is enough to remember the ending index of each block. Suppose the current block ends at position `i`
and `A_i = v`. If the previous block ended at position `q`, then the current block is valid exactly
when the interval `(q, i]` contains exactly `v` occurrences of value `v`. Also, this automatically
keeps adjacent block values different: if `q` were at a previous `v`, then `(q, i]` would not include
that `v`, so it would not be the earliest greedy boundary for a block of `v`.

Use 1-based positions in the DP:

- `dp[i]` = number of valid canonical block-end sequences whose last block ends at position `i`.
- `dp[0] = 1` is the empty sequence before the first block.
- `fp[i] = dp[0] + dp[1] + ... + dp[i]` modulo `998244353`.
- `pos[v]` stores all positions where value `v` has appeared so far.

Process positions from left to right. After appending the current position `i` to `pos[v]`, let
`m = len(pos[v])`. If `m < v`, there are not enough copies of `v` to form a block ending at `i`, so
`dp[i] = 0`.

Otherwise, among the occurrences of `v`, the current block must use the last `v` occurrences ending
at `i`. Let

```text
r = pos[v][m-v]
```

be the first occurrence in those last `v` copies. Then the previous endpoint `q` must satisfy
`q < r`; otherwise `(q, i]` would contain fewer than `v` copies of `v`.

If there is an occurrence of `v` immediately before this block, let

```text
l = pos[v][m-v-1]
```

Then `q <= l` is also invalid, because `(q, i]` would contain more than `v` copies of `v`. Therefore
the valid previous endpoints are:

```text
l < q < r
```

or, if there is no such `l`, simply:

```text
0 <= q < r
```

So the transition is a prefix-sum range query:

```text
dp[i] = fp[r-1]                  if m == v
dp[i] = fp[r-1] - fp[l]          if m > v
```

All arithmetic is modulo `998244353`.

Finally, every non-empty 221 subsequence has exactly one last block endpoint, so the answer is:

```text
dp[1] + dp[2] + ... + dp[N]
```

The algorithm stores each position once in `pos`, and every DP transition is `O(1)`, so the total
complexity is `O(N)` time and `O(N)` memory.
