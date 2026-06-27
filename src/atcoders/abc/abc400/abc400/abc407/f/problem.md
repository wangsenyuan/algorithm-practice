# F - Sums of Sliding Window Maximum

[Problem link](https://atcoder.jp/contests/abc407/tasks/abc407_f)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

input: stdin

output: stdout

You are given a sequence of non-negative integers `A = (A_1, ..., A_N)` of
length `N`.

For each `k = 1, ..., N`, solve the following problem:

- `A` has `N - k + 1` contiguous subarrays of length `k`. Take the maximum of
  each of them, and output the sum of these maxima.

## Constraints

- `1 <= N <= 2 * 10^5`
- `0 <= A_i <= 10^7` (`1 <= i <= N`)
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N
A_1 A_2 ... A_N
```

## Output

Output `N` lines. The `i`-th line (`1 <= i <= N`) should contain the answer for
`k = i`.

## Examples

### Input 1

```text
4
5 3 4 2
```

### Output 1

```text
14
13
9
5
```

For `k = 2`, the length-2 subarrays are `(5,3)`, `(3,4)`, `(4,2)` with maxima
`5`, `4`, `4`, summing to `13`.

### Input 2

```text
8
2 0 2 5 0 5 2 4
```

### Output 2

```text
20
28
27
25
20
15
10
5
```

### Input 3

```text
11
9203973 9141294 9444773 9292472 5507634 9599162 497764 430010 4152216 3574307 430010
```

### Output 3

```text
61273615
68960818
69588453
65590626
61592799
57594972
47995810
38396648
28797486
19198324
9599162
```

## Solution

Instead of computing the maximum for every window directly, count each `A_i`'s contribution to all
answers.

For a fixed index `i`, find the maximal interval `[L_i, R_i]` such that `A_i` is chosen as the
representative maximum for every window inside this interval that contains `i`.

To avoid double counting equal values, ties are broken consistently:

- on the left, stop at the previous index with value `> A_i`;
- on the right, stop at the next index with value `>= A_i`.

This means equal maximum values are assigned to exactly one position. The implementation gets these
boundaries with two monotonic stack passes:

- left pass pops `<= A_i`, so `L_i` is after the previous strictly greater value;
- right pass pops `< A_i`, so `R_i` is before the next greater-or-equal value.

Now only one question remains: for each window length `k`, how many valid windows of length `k` contain
`i` and stay inside `[L_i, R_i]`?

Let:

```text
x = i - L_i + 1          // choices/space on the left side including i
y = R_i - i + 1          // choices/space on the right side including i
m = R_i - L_i + 1        // interval length
k1 = min(x, y)
k2 = max(x, y)
```

The number of length-`k` windows containing `i` is:

```text
k                 if 1 <= k <= k1
k1                if k1 < k <= k2
m - k + 1         if k2 < k <= m
0                 otherwise
```

Why this shape?

- For small `k`, the window can slide around `i` freely, giving `k` placements.
- After the shorter side is exhausted, the count stays flat at `k1`.
- Once `k` is longer than the longer side too, the window almost fills the whole interval, and the
  number of possible placements decreases by one each time.

So `A_i` contributes:

```text
A_i * count(k)
```

to the answer for each `k`.

Each of the three count formulas is linear in `k`:

```text
k          = 1 * k + 0
k1         = 0 * k + k1
m - k + 1  = -1 * k + (m + 1)
```

The implementation keeps two difference arrays, representing coefficients `a` and `b` for expressions
of the form:

```text
a * k + b
```

For every index `i`, it adds these three ranges:

```text
[1, k1]       :  A_i * k
[k1+1, k2]    :  A_i * k1
[k2+1, m]     :  A_i * (m - k + 1)
```

After all range additions, prefix-sum the two difference arrays. For each length `k`, if the accumulated
coefficients are `coef[k]` and `constant[k]`, then:

```text
answer[k] = k * coef[k] + constant[k]
```

### Correctness

The monotonic-stack boundaries assign every window maximum to exactly one index. If a window has a unique
maximum, it is assigned to that maximum. If a window has several equal maximum values, the asymmetric
strict/non-strict boundaries make only one of those equal positions own the window, so there is no double
counting.

For a fixed owner index `i`, a length-`k` window containing `i` is valid exactly when its left endpoint is
inside `[L_i, i]` and its right endpoint is inside `[i, R_i]`. Counting those placements gives the three
ranges above: increasing, flat, then decreasing. Therefore the range updates add `A_i` once for every
window whose assigned maximum is `A_i`, and never for any other window.

Since every window has exactly one assigned maximum, summing all contributions gives exactly the sum of
sliding window maximums for every length.

### Complexity

The two monotonic-stack passes are `O(N)`. Each index performs only three range additions, and the final
prefix pass is `O(N)`.

Total complexity is `O(N)` time and `O(N)` memory.
