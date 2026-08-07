# C2. Equal Multisets (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2211/C2)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

This is the hard version. The difference from the easy version is that `a` is
arbitrary.

You are given arrays `a` and `b` of length `n`, and an integer `k`. Array `a`
contains integers from `1` to `n`. Array `b` contains integers from `1` to `n`
and `-1`.

After replacing every `-1` in `b` with an integer from `1` to `n`, array `b` is
called cool with respect to `k` if every length-`k` window has the same multiset
as the corresponding window in `a`. Formally, for every `i` from `k` to `n`:

```text
[a_{i-k+1}, a_{i-k+2}, ..., a_i]
```

must be a rearrangement of:

```text
[b_{i-k+1}, b_{i-k+2}, ..., b_i]
```

Determine whether it is possible to replace all `-1` values so that `b` is cool.

## Constraints

- `1 <= t <= 10^4`
- `1 <= k <= n <= 2 * 10^5`
- `1 <= a_i <= n`
- `b_i = -1` or `1 <= b_i <= n`
- The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Input

The first line contains the number of test cases `t`.

For each test case:

- The first line contains `n` and `k`.
- The second line contains `a_1, a_2, ..., a_n`.
- The third line contains `b_1, b_2, ..., b_n`.

## Output

For each test case, output `YES` if all `-1` values can be replaced to make `b`
cool, and `NO` otherwise.

The answer is case-insensitive.

## Example

### Input

```text
5
5 5
1 2 3 4 5
3 1 5 2 4
5 2
1 2 1 2 1
2 -1 -1 -1 -1
6 1
5 6 2 2 4 3
5 -1 -1 2 -1 3
2 1
1 2
2 -1
6 4
1 2 3 4 1 2
2 -1 3 -1 4 -1
```

### Output

```text
YES
YES
YES
NO
NO
```

## Notes

- In the first sample, `k = 5`, so the only window is the whole array, and `b`
  is a rearrangement of `a`.
- In the second sample, one valid replacement is `b = [2, 1, 2, 1, 2]`.
- In the fourth sample, `k = 1`, so each fixed position must match exactly;
  since `a_1 != b_1`, the answer is `NO`.

## ideas
1. assume a[1:k] <> b[1:k]
2. move to 2, 如果 a[1] = a[k+1], 那么 b[1] = b[k+1] must hold
3. 如果 a[1] != a[k+1]呢?
4. 这个时候, a[1] = b[1] 必须成立, a[k+1] = b[k+1] must hold
5.

## Solution

Compare two adjacent windows:

```text
[i-k, ..., i-1]
[i-k+1, ..., i]
```

Their multisets differ only by removing the left endpoint and adding the right
endpoint. Since the window multisets of `b` must match those of `a`, every
transition gives one of two constraints:

- If `a_i = a_{i-k}`, then the multiset does not change, so `b_i` must equal
  `b_{i-k}`.
- If `a_i != a_{i-k}`, then the outgoing and incoming values are determined:
  `b_{i-k} = a_{i-k}` and `b_i = a_i`.

The implementation builds these constraints with DSU:

1. For every `i >= k` with `a_i = a_{i-k}`, union positions `i` and `i-k`.
2. Every fixed value in `b` assigns a value to its component.
3. Every transition with `a_i != a_{i-k}` assigns exact values to the two
   endpoint components.
4. Any component assigned two different values makes the answer `NO`.

After all transition constraints are applied, only the first window remains to
check. Its multiset must equal the first `k` values of `a`. Count those needed
values, then subtract the values already forced by the components containing
positions `0..k-1`. If any count becomes negative, the first window cannot be
completed, so the answer is `NO`; otherwise the remaining unassigned positions
can fill exactly the remaining counts.

The first-window check is necessary. For example:

```text
n = 2, k = 1
a = [2, 2]
b = [-1, 1]
```

The transition says `b_1 = b_0`, but the first window requires `b_0 = 2`, while
the fixed second value forces `b_1 = 1`; therefore the correct answer is `NO`.

Complexity: `O(n alpha(n))` per test case.
