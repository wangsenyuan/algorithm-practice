# P3834 【模板】可持久化线段树 2（静态区间第 k 小）

[Problem link](https://www.luogu.com.cn/problem/P3834)

time limit per test: 1.00s

memory limit per test: 1.00GB

This is a classic persistent weighted segment tree (chairman tree) template: the static range k-th smallest.

You are given a sequence `a` of `n` integers. Answer `m` queries. Each query gives a closed interval `[l, r]` and an integer `k`, and asks for the k-th smallest value in that interval.

## Input

The first line contains two integers `n` and `m`, the length of the sequence and the number of queries.

The second line contains `n` integers. The i-th integer is `a_i`.

Each of the next `m` lines contains three integers `l`, `r`, and `k`, asking for the k-th smallest value in `[l, r]`.

## Output

Print `m` lines. The i-th line is the answer to the i-th query.

## Sample

### Input

```text
5 5
25957 6405 15770 26287 26465
2 2 1
3 4 1
4 5 1
1 2 2
4 4 1
```

### Output

```text
6405
15770
26287
25957
26287
```

### Note

The sequence is `{25957, 6405, 15770, 26287, 26465}`.

- `[2, 2]`, 1st smallest: `6405`
- `[3, 4]`, 1st smallest: `15770`
- `[4, 5]`, 1st smallest: `26287`
- `[1, 2]`, 2nd smallest: `25957`
- `[4, 4]`, 1st smallest: `26287`

## Constraints

- For 10% of tests, `1 <= n, m <= 10`.
- For 25% of tests, `1 <= n, m <= 10^3`.
- For 40% of tests, `1 <= n, m <= 10^5`.
- For 50% of tests, `1 <= n, m <= 2 * 10^5` and `0 <= a_i <= 2 * 10^5`.
- For all tests, `1 <= n, m <= 2 * 10^5`, `0 <= a_i <= 10^9`, `1 <= l <= r <= n`, and `1 <= k <= r - l + 1`.
