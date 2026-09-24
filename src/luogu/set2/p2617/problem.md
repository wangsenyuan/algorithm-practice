# P2617 Dynamic Rankings

[Problem link](https://www.luogu.com.cn/problem/P2617)

time limit per test: 5.00s

memory limit per test: 1.00GB

You are given a sequence `a` of `n` integers. Support two kinds of operations:

- `Q l r k` — the k-th smallest value among `a[l], ..., a[r]`
- `C x y` — set `a[x]` to `y`

This is the dynamic range k-th smallest problem (BZOJ1901).

## Input

The first line contains two positive integers `n` and `m`, the length of the sequence and the
number of operations.

The second line contains `n` integers `a_1, a_2, ..., a_n`.

Each of the next `m` lines is one operation, in one of the two formats above.

## Output

For each query, print one integer on its own line.

## Sample

### Input

```text
5 3
3 2 1 4 7
Q 1 4 3
C 2 6
Q 2 5 3
```

### Output

```text
3
6
```

### Note

The initial sequence is `{3, 2, 1, 4, 7}`.

- `Q 1 4 3`: values `{3, 2, 1, 4}`, 3rd smallest is `3`
- `C 2 6`: the sequence becomes `{3, 6, 1, 4, 7}`
- `Q 2 5 3`: values `{6, 1, 4, 7}`, 3rd smallest is `6`

## Constraints

- For 10% of tests, `1 <= n, m <= 100`
- For 20% of tests, `1 <= n, m <= 1000`
- For 50% of tests, `1 <= n, m <= 10^4`
- For all tests, `1 <= n, m <= 10^5`, `1 <= l <= r <= n`, `1 <= k <= r - l + 1`,
  `1 <= x <= n`, and `0 <= a_i, y <= 10^9`

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
