# G - Prefix Concatenation

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_g)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 2 sec

memory limit: 1024 MiB

score: 600 points

You are given two strings S and T consisting of lowercase English letters.

Find the minimum positive integer k such that T can be written as the concatenation of k (not
necessarily distinct) prefixes of S.

Let S_i denote the prefix of S consisting of its first i characters. You need the minimum k such
that there exist integers a_1, ..., a_k with 1 <= a_j <= |S| and

T = S_{a_1} + S_{a_2} + ... + S_{a_k}

If no such k exists, print -1.

## Constraints

- 1 <= |S| <= 5 * 10^5
- 1 <= |T| <= 5 * 10^5
- S and T consist of lowercase English letters

## Input

```text
S
T
```

## Output

Print the minimum k, or -1 if impossible.

## Sample Input 1

```text
aba
ababaab
```

## Sample Output 1

```text
3
```

T = `ababaab` = `ab` + `aba` + `ab`, and each piece is a prefix of S = `aba`. It cannot be done with
two or fewer prefixes.

## Sample Input 2

```text
atcoder
ac
```

## Sample Output 2

```text
-1
```

T cannot be formed as a concatenation of prefixes of S.

## Solution

Let `dp[i]` be the minimum number of prefixes of `S` needed to build the suffix `T[i:]`.
Then `dp[|T|] = 0`.

If we start at position `i` of `T`, we may choose any prefix of `S` that matches `T[i:]`. Suppose the
longest match length is `z`. Then every length from `1` to `z` is also valid, because any prefix of a
matching prefix is still a prefix of `S`.

So the transition is:

```text
dp[i] = 1 + min(dp[i+1], dp[i+2], ..., dp[i+z])
```

The remaining task is to compute `z` quickly for every position of `T`.

Build the string:

```text
S + "#" + T
```

and compute its Z-function. For a position corresponding to `T[i]`, the Z value gives the longest
prefix of `S` that matches `T[i:]`. This is exactly the `z` needed by the DP transition.

The code computes `dp` from right to left. Since each transition asks for a minimum over a range of
already-computed `dp` values, it stores those values in a segment tree:

1. Initialize `dp[|T|] = 0`.
2. Process positions of `T` from right to left.
3. If the Z value is positive, query the minimum `dp` value over the reachable next positions.
4. Set the current `dp` to that minimum plus `1`.

If `dp[0]` is still infinity, then `T` cannot be formed.

### Correctness

For each position `i`, the Z-function gives the maximum length `z` such that `T[i:i+z]` equals a
prefix of `S`. Therefore the valid first pieces from `T[i:]` have exactly the lengths
`1, 2, ..., z`.

After choosing a first piece of length `l`, the remaining suffix is `T[i+l:]`, whose optimal cost is
`dp[i+l]`. Hence the best choice from position `i` is
`1 + min(dp[i+1], ..., dp[i+z])`, which is precisely the transition used by the algorithm.

The DP is evaluated from right to left, so every needed future state has already been computed when
`dp[i]` is calculated. The segment tree returns the minimum over exactly those reachable future
states, so each `dp[i]` is computed correctly.

By induction from the base case `dp[|T|] = 0`, all DP values are correct. Therefore `dp[0]` is the
minimum number of prefixes needed to build all of `T`, and if it is unreachable the correct answer is
`-1`.

### Complexity

The Z-function is computed in `O(|S| + |T|)`. There are `|T|` DP states, and each segment-tree query
and update costs `O(log |T|)`.

The total time complexity is `O((|S| + |T|) + |T| log |T|)`, and the memory usage is
`O(|S| + |T|)`.
