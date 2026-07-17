# D. Two Digit Strings

[Problem link](https://codeforces.com/problemset/problem/2242/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given two strings `a` and `b`, consisting only of digits. In one
operation, you can choose two adjacent characters in either string and replace
them with their sum modulo `10`. For example, from `57246` you can obtain in
one operation:

- `57246 → 2246` (`5+7=12 ≡ 2`)
- `57246 → 5946`
- `57246 → 5766`
- `57246 → 5720`

After such an operation, the length of the string decreases by exactly `1`.

You may perform any number of operations (including zero). Make the two
strings equal (`|a| = |b|` and corresponding characters match). Find the
**maximum** possible length of the resulting equal strings. If they cannot be
made equal, print `-1`.

## Constraints

- `1 ≤ t ≤ 2 · 10^3`
- `1 ≤ |a|, |b| ≤ 5 · 10^3`
- `a` and `b` consist only of digits
- Sum of `|a| + |b|` over all test cases ≤ `10^4`

## Input

```
t
a
b
...
```

## Output

For each test case, print one integer — the maximum possible length of the
resulting equal strings, or `-1` if impossible.

## Samples

### Sample 1

Input:

```
1
5147
44441
```

Output:

```
2
```

One optimal way:

- `5147 → 647 → 61`
- `44441 → 8441 → 241 → 61`

### Sample 2

Input:

```
1
21945
60
```

Output:

```
-1
```

### Sample 3

Input:

```
1
123450
012345
```

Output:

```
5
```

One optimal way:

- `123450 → 12345`
- `012345 → 12345`

## Ideas

### Operations are equivalent to partitioning the string

After repeatedly merging adjacent digits, every digit in the resulting string
comes from one nonempty contiguous segment of the original string. Its value
is the sum of that segment modulo `10`.

For example, partitioning a string into three segments produces:

```text
[segment 1] [segment 2] [segment 3]
     ↓           ↓           ↓
 sum % 10    sum % 10    sum % 10
```

Conversely, any partition into contiguous segments can be produced by merging
all digits inside each segment. Therefore, the problem is to partition `a`
and `b` into the maximum number of segments such that corresponding segment
sums are equal modulo `10`.

### Prefix sums modulo 10

Define:

```text
pa[i] = digit sum of a[0:i] modulo 10
pb[j] = digit sum of b[0:j] modulo 10
```

with `pa[0] = pb[0] = 0`.

Suppose the previous matched segment ends at positions `(i0, j0)`, and the
next one ends at `(i, j)`. The two new segment values are equal exactly when:

```text
(pa[i] - pa[i0]) mod 10 = (pb[j] - pb[j0]) mod 10
```

All previous segment values were equal, so their accumulated sums are also
equal:

```text
pa[i0] = pb[j0]
```

Consequently, the new segments are equal exactly when:

```text
pa[i] = pb[j]
```

Thus, every pair of equal prefix sums gives the next pair of segment
boundaries. We want the longest sequence of such pairs with both positions
strictly increasing. This is precisely the longest common subsequence of:

```text
pa[1], pa[2], ..., pa[n]
pb[1], pb[2], ..., pb[m]
```

### The final boundary

The final segments must end at `n` and `m`. Hence the total digit sums modulo
`10` must be equal:

```text
pa[n] = pb[m]
```

If they differ, no common resulting string exists, so the answer is `-1`.

If they are equal, the two LCS sequences have the same final element. An
optimal common subsequence can therefore include this final match, ensuring
that the selected segment boundaries cover both complete strings.

### Dynamic programming

Let `dp[i][j]` be the LCS length between:

```text
pa[1:i+1] and pb[1:j+1]
```

Using one-based DP dimensions, the transitions for `i < n` and `j < m` are:

```text
dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])

if pa[i+1] == pb[j+1]:
    dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j] + 1)
```

The answer is `dp[n][m]`.

### Correctness summary

1. Every sequence of operations represents a partition into contiguous
   segments, and every such partition can be realized with adjacent merges.
2. Once the preceding boundaries have equal prefix sums, the next two segment
   sums are equal exactly when their ending prefix sums are equal.
3. Therefore, a common resulting string corresponds exactly to an increasing
   common subsequence of the two modulo-`10` prefix-sum arrays.
4. Equal total sums provide the required final boundary, so their LCS length
   is the maximum possible resulting length.

### Complexity

For strings of lengths `n` and `m`:

```text
Time:   O(nm)
Memory: O(nm)
```
