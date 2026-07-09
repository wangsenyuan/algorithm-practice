# D - No-Subsequence Substring

[Problem link](https://atcoder.jp/contests/abc452/tasks/abc452_d)

**Contest:** [AtCoder Beginner Contest 452](https://atcoder.jp/contests/abc452)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

You are given strings `S` and `T` consisting of lowercase English letters.

Among the non-empty substrings `s` of `S`, count those that do **not** contain `T`
as a (not necessarily contiguous) subsequence.

Two substrings of `S` are distinguished if they are taken from different positions,
even if they are equal as strings.

A substring of `X` is obtained by deleting zero or more characters from the beginning
and zero or more from the end of `X`. A subsequence of `X` is obtained by deleting
zero or more characters from `X` while keeping the relative order of the rest.

## Constraints

- `S` consists of lowercase English letters
- `1 <= |S| <= 2 * 10^5`
- `T` consists of lowercase English letters
- `1 <= |T| <= 50`

## Input

```text
S
T
```

## Output

Print the answer.

## Sample Input 1

```text
abrakadabra
aba
```

## Sample Output 1

```text
51
```

For example, the substring `abr` from the first through third characters of `S` does
not contain `T` as a subsequence. Including this, there are 51 valid substrings, such
as `k` (only the fifth character) and `akada` (the fourth through eighth characters).

Note that `abr` appears both as positions `1..3` and as positions `8..10`; they are
counted separately because the positions differ.

## Sample Input 2

```text
aaaaa
a
```

## Sample Output 2

```text
0
```

Every non-empty substring of `S` contains `T` as a subsequence.

## Sample Input 3

```text
rdddrdtdcdrrdcredctdordoeecrotet
dcre
```

## Sample Output 3

```text
263
```

## Solution

Count substrings of `S` that do **not** contain `T` as a subsequence by scanning
right endpoints and using DP over the matched prefix of `T`.

### Key idea

Fix the right endpoint `r`. Among all left endpoints `l`, the substrings
`S[l..r]` that contain `T` as a subsequence are exactly those with
`l <= L(r)`, where `L(r)` is the **rightmost** start position of any subsequence
match of `T` ending at or before `r`.

So the good left endpoints are `L(r) + 1, ..., r`, and there are
`r - L(r)` of them. If no match exists yet, every `l` in `0..r` is good, so add
`r + 1`.

### DP

Let `dp[j]` be the rightmost start index of a match of `T[0..j]` using characters
of `S` up to the current position (or `-1` if impossible).

Process `S` from left to right. When reading `S[i]`, update from the end of `T`
backward:

```text
if S[i] == T[j]:
  if j == 0: dp[0] = i
  else:      dp[j] = max(dp[j], dp[j-1])
```

Updating from large `j` to small `j` ensures each character is used at most once
in the same transition step.

After processing `i`:

- if `dp[|T|-1] != -1`, add `i - dp[|T|-1]`
- otherwise add `i + 1`

### Complexity

`|T| <= 50` and `|S| <= 2 * 10^5`, so the time complexity is
`O(|S| * |T|)`, which is fine. The answer can be up to about `|S|^2 / 2`, so use
a 64-bit integer.
