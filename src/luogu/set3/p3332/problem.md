# P3332 [ZJOI2013] K 大数查询

[Problem link](https://www.luogu.com.cn/problem/P3332)

time limit per test: 2.00s

memory limit per test: 512.00MB

You need to maintain `n` integer multisets, numbered from `1` to `n`. All sets start empty. There
are `m` operations:

- `1 l r c` — insert `c` into every set whose index is in `[l, r]`
- `2 l r k` — among the multiset union of sets whose indices are in `[l, r]`, what is the k-th
  largest number?

The union of multisets keeps duplicate elements. For example
`{1, 1, 4} ∪ {5, 1, 4} = {1, 1, 4, 5, 1, 4}`.

## Input

The first line contains two positive integers `n` and `m`, the number of sets and the number of
operations.

Each of the next `m` lines contains four integers describing one operation, as above.

## Output

For each type-`2` operation, print one integer on its own line.

## Sample

### Input

```text
2 5
1 1 2 1
1 1 2 2
2 1 1 2
2 1 1 1
2 1 2 3
```

### Output

```text
1
2
1
```

### Note

- Operation 1 inserts `1` into sets `1` and `2`.
- Operation 2 inserts `2` into sets `1` and `2`. Set `1` is now `{1, 2}`.
- Query set `1`, 2nd largest: `1`
- Query set `1`, 1st largest: `2`
- Query the union of sets `1` and `2`, which is `{1, 2, 1, 2}`, 3rd largest: `1`

## Constraints

- `1 <= n, m <= 5 * 10^4`
- `1 <= l <= r <= n`
- For type-`1` operations, `|c| <= n`
- For type-`2` operations, `1 <= k < 2^63`, and the k-th largest number exists

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
