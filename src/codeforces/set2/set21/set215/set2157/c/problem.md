# C. Meximum Array 2

[Problem link](https://codeforces.com/problemset/problem/2157/C)

**Contest:** [Codeforces Round 1066 (Div. 1 + Div. 2)](https://codeforces.com/contest/2157)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given three positive integers `n`, `k`, and `q`, and `q` tuples `(c, l, r)` with `1 <= c <= 2` and `1 <= l <= r <= n`.

An array `a_1, a_2, …, a_n` is meximum if `0 <= a_i <= 10^9` for each `i`, and for each given tuple `(c, l, r)`:

- if `c = 1`, then `min(a_l, …, a_r) = k`;
- if `c = 2`, then `MEX(a_l, …, a_r) = k`.

`MEX` of a collection is the smallest non-negative integer that does not appear in it.

Find a meximum array of length `n`. A valid array always exists. If there are multiple possible arrays, print any of them.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 500`).

The first line of each test case contains three integers `n`, `k`, and `q` (`1 <= k <= n <= 100`, `1 <= q <= 100`).

Then `q` lines follow. The `i`-th line contains a tuple `(c, l, r)`.

It is guaranteed that a valid array exists.

## Output

For each test case, print a single line containing a meximum array `a_1, a_2, …, a_n`.

## Example

### Input

```text
4
6 2 2
1 1 3
2 2 6
3 3 1
2 1 3
3 3 2
1 1 1
1 3 3
3 2 2
2 1 2
2 2 3
```

### Output

```text
2 5 4 3 0 1
2 0 1
3 3 3
1 0 1
```

### Note

In the first test case, `min(a_1, a_2, a_3)` must be `2` and `MEX(a_2, …, a_6)` must be `2`. One valid array is `[2, 5, 4, 3, 0, 1]`.

## Solution

Difference arrays mark every index that sits in a `min` constraint and
every index that sits in a `MEX` constraint. After the prefix sums, each
position falls into one of four kinds:

- **min only** — must be `>= k` and should equal `k` so the min is hit;
- **MEX only** — must not be `k`, and should be `< k` so the range can
  collect `{0, …, k-1}`;
- **both** — must be `>= k` and `!= k`, so `k+1`;
- **neither** — unconstrained, left as `0`.

The remaining MEX-only positions, in left-to-right order, are filled with
`i mod k`. A valid input always has at least `k` such positions inside
every MEX range, so that cyclic fill puts every residue `0..k-1` into the
range and never writes `k`.

### Correctness sketch

A min-only cell cannot lie in a MEX range (else it would be marked both),
so writing `k` there satisfies the min without injecting `k` into a MEX.
Both-covered cells are `k+1`, which is legal for min and invisible to MEX.
MEX-only cells in one range form a contiguous block of the global leftover
sequence; length at least `k` plus `i mod k` yields all of `0..k-1` and
omits `k`, so the MEX is exactly `k`. Unconstrained zeros sit outside every
query.

### Complexity

Difference arrays and one fill pass: `O(n + q)` time and `O(n)` memory.
`n` and `q` are at most `100`.
