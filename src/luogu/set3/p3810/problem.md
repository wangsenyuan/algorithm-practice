# P3810 【模板】三维偏序 / 陌上花开

[Problem link](https://www.luogu.com.cn/problem/P3810)

time limit per test: 1.00s

memory limit per test: 500.00MB

This is a template problem. It can be solved with bitset, CDQ divide-and-conquer, Fenwick-in-Fenwick, K-D trees, and similar techniques.

There are `n` elements. The `i`-th element has three attributes `a_i`, `b_i`, and `c_i`. Let `f(i)` be the number of indices `j` such that `a_j <= a_i`, `b_j <= b_i`, `c_j <= c_i`, and `j != i`.

For every `d` in `[0, n)`, count how many indices `i` satisfy `f(i) = d`.

## Input

The first line contains two integers `n` and `k`, the number of elements and the maximum attribute value.

Each of the next `n` lines contains three integers `a_i`, `b_i`, and `c_i`.

## Output

Print `n` lines. Line `d + 1` is the number of indices `i` with `f(i) = d`.

## Sample

### Input

```text
10 3
3 3 3
2 3 3
2 3 1
3 1 1
3 1 2
1 3 1
1 1 2
1 2 2
1 3 2
1 2 1
```

### Output

```text
3
1
3
0
1
0
1
0
0
1
```

## Constraints

For all tests, `1 <= n <= 10^5` and `1 <= a_i, b_i, c_i <= k <= 2 * 10^5`.

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
