# Codeforces 2106E - Wolf

https://codeforces.com/problemset/problem/2106/E

## Statement

Wolf has found `n` sheep with tastiness values `p1, p2, ..., pn`, where `p` is a
permutation. Wolf wants to perform binary search on `p` to find the sheep with
tastiness `k`, but `p` may not necessarily be sorted.

The success of binary search on the range `[l, r]` for `k` is represented as
`f(l, r, k)`, defined as follows:

- If `l > r`, then `f(l, r, k)` fails.
- Otherwise, let `m = floor((l + r) / 2)`.
- If `p[m] = k`, then `f(l, r, k)` is successful.
- If `p[m] < k`, then `f(l, r, k) = f(m + 1, r, k)`.
- If `p[m] > k`, then `f(l, r, k) = f(l, m - 1, k)`.

Cow the Nerd decides to help Wolf out. Cow the Nerd is given `q` queries, each
consisting of three integers `l`, `r`, and `k`.

Before the search begins, Cow the Nerd may choose a non-negative integer `d`,
and `d` indices:

```text
1 <= i1 < i2 < ... < id <= n
```

such that `p[ij] != k` for all `1 <= j <= d`. Then, he may reorder the elements:

```text
p[i1], p[i2], ..., p[id]
```

however he likes.

For each query, output the minimum integer `d` that Cow the Nerd must choose so
that `f(l, r, k)` can be successful, or report that it is impossible.

The queries are independent, and the reordering is not actually performed.

A permutation of length `n` is an array that contains every integer from `1` to
`n` exactly once.

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of
test cases.

For each test case:

- The first line contains two integers `n` and `q`
  (`1 <= n, q <= 2 * 10^5`) — the length of `p` and the number of queries.
- The second line contains `n` integers `p1, p2, ..., pn` — the tastiness of each
  sheep. It is guaranteed that every integer from `1` to `n` appears exactly
  once in `p`.
- The following `q` lines each contain three integers `l`, `r`, and `k`
  (`1 <= l <= r <= n`, `1 <= k <= n`) — the binary-search range and the target
  value.

It is guaranteed that the sum of `n` and the sum of `q` over all test cases do
not exceed `2 * 10^5`.

## Output

For each query, output the minimum integer `d` needed so that `f(l, r, k)` is
successful. If it is impossible, output `-1`.

## Sample

```text
8
5 3
1 2 3 4 5
1 5 4
1 3 4
3 4 4
7 4
3 1 5 2 7 6 4
3 4 2
2 3 5
1 5 6
1 7 3
2 1
2 1
1 2 1
1 1
1
1 1 1
7 1
3 4 2 5 7 1 6
1 7 1
16 1
16 10 12 6 13 9 14 3 8 11 15 2 7 1 5 4
1 16 4
16 1
14 1 3 15 4 5 6 16 7 8 9 10 11 12 13 2
1 16 14
13 1
12 13 10 9 8 4 11 5 7 6 2 1 3
1 13 2
```

```text
0 -1 0
2 0 -1 4
-1
0
-1
-1
-1
-1
```

## Note

In the first example, second query, `4` does not exist in the first three
elements, so it is impossible to find it when searching in that range.

In the second example, first query, choose indices `2` and `3`, and swap them so:

```text
p = [3, 5, 1, 2, 7, 6, 4]
```

Then `f(3, 4, 2)` works:

- `m = floor((3 + 4) / 2) = 3`;
- `p[3] = 1 < 2`, so continue with `f(4, 4, 2)`;
- `m = floor((4 + 4) / 2) = 4`;
- `p[4] = 2 = k`, so the search succeeds.

The total number of chosen indices is `2`, which is minimum. For this query, we
cannot choose index `4`, because `p[4] = k = 2`.

In the last query of the second example, choose indices `2, 3, 4, 5` and
rearrange them so:

```text
p = [3, 5, 2, 7, 1, 6, 4]
```

Then `f(1, 7, 3)` is successful.

## Solution summary (from `solution.go`)

**Preprocessing**

- Build `pos[v]` = 1-based index of each value `v` in the permutation `p`.

**Queries**

- Answer each query in input order: for index `i`, read `(l, r, k)` and run the same path simulation (no extra bucketing or sweep over `r`).

**Per query `(l, r, k)`**

- Let `exp = pos[k]`. If `k` is not inside `[l, r]`, output `-1`.
- Otherwise simulate the **binary-search recursion** on indices `(l, r)` that always splits at `mid = floor((l + r) / 2)` and recurses toward `exp` (the index where `p[exp] = k`), until `mid == exp` (success path in index space).

**Counters along that path**

- When the search goes **right** (`mid < exp`): compare `p[mid]` with `p[exp]` (= `k`).
  - If `p[mid] > k`, the value at `mid` would send standard BS the wrong way; count `lt`.
  - Else count `lt1` (already consistent with going right toward `k`).
- When the search goes **left** (`mid > exp`): compare `p[mid]` with `k`.
  - If `p[mid] < k`, count `gt` (wrong for going left).
  - Else count `gt1`.

**Feasibility**

- If `lt + lt1 >= k` or `gt + gt1 > n - k`, output `-1` (not enough room in values `< k` or `> k` to fix the path with allowed reorderings of non-`k` positions).

**Answer**

- Otherwise `d = lt + gt + abs(lt - gt)` — minimum number of indices to touch so that `f(l, r, k)` can succeed, derived from the imbalance between “wrong” left vs right branch comparisons along the simulated path.

**Complexity**

- Building `pos` is `O(n)`. Each of the `q` queries walks one binary-search path of length `O(log n)`. Total per test case: `O(n + q log n)` time and `O(n + q)` space.
