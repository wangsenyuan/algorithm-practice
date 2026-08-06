# E. Excuse for Breaks

[Problem link](https://codeforces.com/problemset/problem/2248/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given integers `n`, `m`, and `d`, and two arrays `p_1, p_2, ..., p_m`
and `r_1, r_2, ..., r_m`. The array `p` is strictly increasing.

For any positive-length binary array `a`, define `f(a)` by scanning the array
from left to right:

- Start with `v = 0` and `c = 0`.
- If `a_i = 1`, add `d` to `v` and increase `c` by `1`.
- If `a_i = 0`, reset `c` to `0`.
- For every `j`, if `c = p_j`, add `r_j` to `v`.
- If `c = n`, reset `c` to `0`.

Let `I(a)` be the all-ones array with the same length as `a`.

Determine whether there exists a binary array `a` such that:

```text
f(a) > f(I(a))
```

## Constraints

- `1 <= t <= 2000`
- `1 <= n <= 10^9`
- `0 <= m <= 2000`
- `0 <= d <= 10^9`
- `1 <= p_i <= n`
- `1 <= r_i <= 10^9`
- `p` is strictly increasing.
- The sum of `m` over all test cases does not exceed `2000`.

## Input

The first line contains the number of test cases `t`.

For each test case:

- The first line contains `n`, `m`, and `d`.
- The next `m` lines each contain `p_i` and `r_i`.

## Output

For each test case, output `YES` if such a binary array exists, and `NO`
otherwise.

The answer is case-insensitive.

## Example

### Input

```text
3
6 4 3
2 5
3 9
4 1
5 3
7 3 5
2 5
4 5
7 10
684492057 3 386217943
367971233 991739271
612599954 429216213
684492056 402931836
```

### Output

```text
YES
NO
YES
```

## Note

In the first test case, the statement gives an example array with runs of ones
of lengths `9`, `4`, and `3`, whose value is larger than the same-length
all-ones array.

## Solution

Let `S_i = F(I(i))`, the value of an all-ones array of length `i`. Since the
counter resets after `n` consecutive ones:

```text
S_{i+n} = S_i + S_n
```

For any binary array, zeros only split the array into runs of ones. The key
claim is that if any binary array can beat the all-ones array of the same
length, then two runs are already enough:

```text
I(x), 0, I(y)
```

So the answer is `YES` iff there exist `x, y >= 1` such that:

```text
S_x + S_y > S_{x+y+1}
```

For fixed `y`, the difference `S_x + S_y - S_{x+y+1}` can only increase when
`x` reaches a reward point. By the same argument for `y`, it is enough to check
only pairs `x = p_i`, `y = p_j`.

The implementation precomputes prefix sums of rewards. A helper `score(i)`
computes `S_i` using the periodic formula, and all `m^2` reward-point pairs are
tested.

In the pair loop, `x = p_i`, so `S_x` can be computed directly as:

```text
S_x = x * d + pref_i
```

Here `pref_i = r_1 + r_2 + ... + r_i`. This works because `p` is strictly
increasing: if the run length is exactly `x = p_i`, then the scan reaches every
reward checkpoint `p_1, p_2, ..., p_i`, and reaches no later checkpoint.

For example, if:

```text
p = [2, 3, 5]
r = [5, 9, 3]
```

then a run of length `3` triggers the rewards at checkpoints `2` and `3`, but
not the checkpoint `5`. Its value is:

```text
3 * d + 5 + 9
```

Complexity: `O(m^2 log m)` per test case, with total `m <= 2000`.
