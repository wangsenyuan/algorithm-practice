# E - Crossing Table Cloth (ABC457)

**Contest:** [ABC457](https://atcoder.jp/contests/abc457) — AtCoder Beginner Contest 457  
**Task:** [https://atcoder.jp/contests/abc457/tasks/abc457_e](https://atcoder.jp/contests/abc457/tasks/abc457_e)

**Time limit:** 2.5 sec / **Memory limit:** 1024 MiB  
**Score:** 475 points

## Problem Statement

There are N cells arranged in a horizontal row. The i-th cell from the left (1 ≤ i ≤ N) is called cell i.

There are M pieces of cloth. Laying cloth i (1 ≤ i ≤ M) covers cells L_i through R_i.

Answer Q queries. For the q-th query (1 ≤ q ≤ Q), integers S_q and T_q are given. Determine whether it is possible to choose exactly two pieces of cloth from the M pieces and lay them so that:

- Cells S_q through T_q are covered by at least one piece of cloth, and
- No other cells are covered by any cloth.

## Constraints

- 1 ≤ N ≤ 2 × 10^5
- 2 ≤ M ≤ 2 × 10^5
- 1 ≤ L_i ≤ R_i ≤ N
- 1 ≤ Q ≤ 2 × 10^5
- 1 ≤ S_q ≤ T_q ≤ N
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N M
L_1 R_1
L_2 R_2
⋮
L_M R_M
Q
S_1 T_1
S_2 T_2
⋮
S_Q T_Q
```

## Output

Output the answers for the queries, separated by newlines.

For each query, output `Yes` if it is possible to choose two pieces of cloth satisfying the condition, and `No` otherwise.

## Sample Input 1

```text
4 3
1 3
1 1
2 4
4
1 4
2 4
1 3
1 1
```

## Sample Output 1

```text
Yes
No
Yes
No
```

For the first query, the condition can be satisfied by choosing cloth 1 and cloth 3.

For the third query, the condition can be satisfied by choosing cloth 1 and cloth 2.

For the second and fourth queries, no choice of two pieces of cloth can satisfy the condition.

## Sample Input 2

```text
7 10
2 6
2 5
3 6
1 6
1 2
5 6
2 3
3 7
2 3
1 2
10
1 2
3 5
1 4
1 5
1 5
5 7
1 6
2 3
5 7
2 4
```

## Sample Output 2

```text
Yes
No
No
Yes
Yes
No
Yes
Yes
No
No
```

## Solution

Exactly two cloths must be chosen so that their union is **exactly** the query interval
`[S, T]`: every cell in the range is covered, and no cell outside it is covered.

### Sweep by right endpoint

Process `r = 1 … N` in increasing order.

- `todo[r]` — left endpoints of cloths with right endpoint `r` (sorted when `r` is reached).
- `pending[r]` — query indices with `T_q = r`.

For each `r`, answer all queries `(l, r)` in `pending[r]` **before** inserting the cloths
that end at `r`. At that moment:

- cloths ending at `r` are available in `todo[r]`;
- cloths ending before `r` are summarized in the sweep state below.

### Two feasible patterns

The `check(l, r)` routine accepts `Yes` in two situations (matching the code comments):

1. **Split / connecting cover** — one cloth covers a prefix of `[l, r]` and another covers
   the suffix, meeting with no gap and without spilling outside `[l, r]`.
2. **Full cover plus inner cloth** — one cloth is exactly `[l, r]`, and a second cloth lies
   strictly inside that range (so the union is still `[l, r]`).

### Sweep state

- `dp[l]` — among cloths ending **before** the current `r`, the maximum length of a cloth
  that starts at `l` (equivalently, how far right a completed cloth from `l` reached).
- `nearest` — the largest left endpoint among all cloths that ended before the current `r`.

### Query test (`check(l, r)`)

Using binary search on the sorted list `todo[r]`:

1. **Two cloths ending at `r` with left `l`** — count entries equal to `l`; if at least two,
   answer `Yes`.
2. **One cloth `[l, r]` plus a partner** — if exactly one cloth starts at `l` and ends at `r`,
   answer `Yes` when either:
   - another cloth ending at `r` starts to the right of `l`, or
   - `nearest ≥ l` (some earlier-ending cloth starts at or after `l`, giving an inner partner).
3. **Split using historical prefix** — if `dp[l] > 0`, let `r1 = l + dp[l]`. Find the rightmost
   cloth ending at `r` whose left endpoint lies in `(l, r1]`. Such a cloth connects with the
   best completed prefix from `l`, so the union is exactly `[l, r]`.

If none of the cases fire, answer `No`.

After all queries for this `r`, update the state: for each `l` in `todo[r]`, set
`dp[l] = r - l + 1` and `nearest = max(nearest, l)`.

### Complexity

Each cloth is sorted once into `todo[r]`; each query does `O(log M)` binary searches on
`todo[r]`. Overall:

```text
Time:  O((N + M + Q) log M)
Space: O(N + M + Q)
```
