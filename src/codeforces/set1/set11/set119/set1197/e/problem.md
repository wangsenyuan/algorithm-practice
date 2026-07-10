# E. Culture Code

[Problem link](https://codeforces.com/problemset/problem/1197/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` matryoshkas (Russian nesting dolls). Doll `i` has outer volume
`out_i` and inner volume `in_i` (`out_i > in_i`).

Doll `i` can be nested inside doll `j` if `out_i <= in_j`.

A set of dolls is **nested** if they can be ordered so that each fits inside the
next. For a nested chain `i_1, i_2, ..., i_k` (innermost to outermost), the
**extra space** is:

```text
in_{i_1} + (in_{i_2} - out_{i_1}) + (in_{i_3} - out_{i_2}) + ... + (in_{i_k} - out_{i_{k-1}})
```

A nested subset is **big enough** if no other doll from the store can be added
while keeping the nested property.

Count the number of big enough nested subsets whose extra space is minimum
possible among all big enough nested subsets. Print the answer modulo
`10^9 + 7`.

Two subsets are different if their sets of doll indices differ.

## Constraints

- `1 <= n <= 2 * 10^5`
- `1 <= in_i < out_i <= 10^9`

## Input

```text
n
out_1 in_1
...
out_n in_n
```

## Output

Print one integer — the answer modulo `10^9 + 7`.

## Example

```text
Input
7
4 1
4 2
4 2
2 1
5 4
6 4
3 2

Output
6
```

### Note

There are 6 big enough nested subsets with minimum extra space `1`:

- `{1, 5}`
- `{1, 6}`
- `{2, 4, 5}`
- `{2, 4, 6}`
- `{3, 4, 5}`
- `{3, 4, 6}`

For example, `{6, 7}` is not big enough (doll `4` can be added), and
`{4, 6, 7}` has extra space `2`.


## Solution

Think of a chain from innermost to outermost.

Let `dp[i]` be the minimum extra space for a valid nested chain whose outermost doll is
`i`. Also store `ways[i]`, the number of chains that achieve this minimum.

If doll `i` contains nothing inside it, then:

```text
dp[i] = in_i
ways[i] = 1
```

If the current inner chain has outermost doll `j`, it can be placed inside `i` when:

```text
out_j <= in_i
```

The new extra space is:

```text
dp[j] + (in_i - out_j)
= in_i + (dp[j] - out_j)
```

So for each `i`, among all dolls `j` with `out_j <= in_i`, we need the minimum value of:

```text
dp[j] - out_j
```

and the number of ways to achieve that minimum. This is a prefix minimum query by
`out_j`, so the implementation compresses all `in` and `out` values and uses a segment
tree. Each tree node stores:

- `val`: the minimum `dp[j] - out_j`;
- `ways`: how many chains achieve that `val`.

When two candidates have the same minimum value, their ways are added modulo `1e9+7`.

The sweep goes over compressed coordinates in increasing order:

1. When reaching a value equal to some doll's `out`, insert that doll's current
   `dp[doll] - out` into the segment tree.
2. When reaching a value equal to some doll's `in`, query all inserted dolls with
   `out <= in`, and use the result to improve `dp[doll]`.

This order is valid because `out > in` for every doll, so a doll is only inserted after
all possible inner chains for it have already been computed.

Now handle the "big enough" condition.

For the outermost doll `i`, no other doll may be placed outside it. That means there must
be no doll `j` with:

```text
out_i <= in_j
```

The code checks this with a suffix count over inner volumes.

What about adding a doll inside the innermost doll? The DP minimum already handles this.
If another doll can be placed inside the current innermost doll, adding it decreases the
extra space by:

```text
out_new - in_new > 0
```

so a non-maximal inner chain can never be a minimum `dp` chain for the same outermost
doll. Therefore the minimum DP chains are already maximal on the inside.

Finally, among all dolls that can be valid outermost dolls, take the minimum `dp[i]` and
sum `ways[i]` for those achieving it.

The time complexity is `O(n log n)` from coordinate compression and segment tree
operations, and the memory complexity is `O(n)`.
