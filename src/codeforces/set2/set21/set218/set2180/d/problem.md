# D. Insolvable Disks

[Problem link](https://codeforces.com/problemset/problem/2180/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given `n` different integer points `x_1 < x_2 < ... < x_n` on the X-axis of
the plane. For each `1 <= i <= n`, you have to pick a real value `r_i > 0` and draw
a disk with radius `r_i` and center `x_i` such that no two disks overlap.

You want to choose values `r_i` in such a way that the number of tangent pairs of
disks is maximized, and you have to find this maximum value.

We say that two disks overlap if the area of their intersection is positive. In
particular, two outer tangent disks do not overlap.

## Input

Each test contains multiple test cases. The first line contains the number of test
cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` (`1 <= n <= 2 * 10^6`)
— the number of points in that test case.

The second line of each test case contains `n` integers `x_1, x_2, ... x_n`
(`0 <= x_1 < x_2 < ... < x_n <= 10^9`) — the coordinates of the points in that test
case.

It is also guaranteed that the sum of `n` over all test cases is at most `2 * 10^6`.

## Output

For each test case, print a single integer denoting the maximum number of tangent
pairs of disks in that test case.

## Example

### Input

```text
4
3
1 2 3
4
1 2 4 5
6
1 2 11 12 21 22
7
0 1 2 3 5 8 13
```

### Output

```text
2
2
3
6
```

## Note

In the first test case, you can let `r_1 = r_3 = 0.25` and `r_2 = 0.75` so the
1-st and 2-nd disks and the 2-nd and 3-rd disks will be pairwise tangent.

In the second test case, you can let `r_1 = 0.4`, `r_2 = 0.6`, `r_3 = 0.2`, and
`r_4 = 0.8` so the 1-st and 2-nd disks and 3-rd and 4-th disks will be pairwise
tangent.

In the third test case, you can set `r_i = 0.5` for each `1 <= i <= 6` so the 1-st
and 2-nd disks, 3-rd and 4-th disks, and 5-th and 6-th disks will be pairwise
tangent.

## Solution

Place disks on sorted points so no two overlap, and maximize the number of
**tangent** pairs (touching but not overlapping).

### Short recap

- Disk `i` has center `x_i` and radius `r_i > 0`.
- Disks overlap iff the intersection area is positive; outer tangency is allowed.
- Two disks are tangent iff `r_i + r_j = |x_j - x_i|`.
- Output the maximum possible number of tangent pairs.

### Key observations

1. **Only adjacent pairs matter.** If disks `i` and `i + k` (`k ≥ 2`) were tangent,
   every disk between them would lie inside one of the two, so such a pair cannot
   appear in a valid non-overlapping layout. An optimal answer uses tangency only
   between consecutive indices.

2. **Upper bound `(n - 1)`.** There are `n - 1` consecutive pairs; the answer is
   `(n - 1)` minus the number of **mandatory breaks** where a tangent chain cannot
   continue.

3. **Linear system on a chain.** If consecutive disks `j … i` are all pairwise
   tangent, let `d_k = x_{k+1} - x_k`. Then
   `r_k + r_{k+1} = d_k` for `j ≤ k < i`. Eliminating radii gives
   `r_j = S(j, i) ± r_{i+1}`, where `S(j, i)` is the alternating sum of gaps from
   `j` to `i` (sign `+` on `d_i`, then alternating). The code maintains this as
   `sum` with `sum ← d_i - sum` at each step.

4. **Positivity ⇒ interval for `r_j`.** All `r_k > 0` in the chain impose lower
   and upper bounds on the start radius `r_j`. As the chain grows rightward,
   feasible values lie in an interval `(lo, hi)` (strict). When `lo ≥ hi`, no
   `r_j > 0` can satisfy every constraint — the chain must end before index `i + 1`.

5. **Parity splits the updates.** Whether `i` and the segment start `j` have the
   same parity determines whether `r_j = S + r_{i+1}` or `r_j = S - r_{i+1}`:
   - **Same parity** (`i & 1 == j & 1`): `hi = min(hi, sum)`; if `i + 1 < n`,
     `lo = max(lo, sum - d_{i+1})`.
   - **Different parity**: `lo = max(lo, -sum)`; if `i + 1 < n`,
     `hi = min(hi, d_{i+1} - sum)`.

### Algorithm (matching `solve`)

**Step 0 — trivial case.** If `n == 1`, return `0`.

**Step 1 — initialize a scan over gaps** (`i = 1 … n - 1`, 0-based array `x`):

```go
res := 0
j := 1              // 1-based start index of the current chain
lo, hi := 0, inf
sum := 0
```

**Step 2 — extend or break** at each gap `d_i = x[i] - x[i-1]`:

1. `sum = d_i - sum` (alternating prefix from `j`).
2. Apply parity-based bound updates (see observation 5).
3. If `lo >= hi`:
   - `res--` (one fewer tangent pair than a full chain would give),
   - reset `j = i + 1`, `lo = 0`, `hi = inf`, `sum = 0`.

**Step 3 — answer.** Return `res + n - 1`.

Each break splits the line into independent chains; tangencies inside a chain of
length `L` contribute `L - 1` pairs, so total `(n - 1) + res` with `res ≤ 0`.

### Why it works

Fix a chain starting at `j`. Writing every radius in terms of `r_j` and the
alternating gap sums shows that **all** positivity conditions are equivalent to
`r_j` lying in an interval derived from the same `sum` and next gap `d_{i+1}`.
The loop maintains the tightest such interval while extending rightward.

If `lo ≥ hi`, no positive radii make every disk in `j … i` tangent to its
neighbor — so no tangent edge can connect disk `i` to disk `i + 1`. Any feasible
global layout must break the chain there. Conversely, when the scan finishes
without a break, the interval is non-empty and one can choose `r_j` in `(lo, hi)`
and recover all other radii positively, making every adjacent pair in the segment
tangent.

Summing over disjoint chains gives the maximum: start from the `(n - 1)` upper
bound and subtract one for each forced break recorded in `res`.

### Walkthrough: `[1, 2, 4, 5]` (sample 2, answer `2`)

Gaps `d_1 = 1`, `d_2 = 2`, `d_3 = 1`. Start `j = 1`.

| step | `i` | `sum` | `lo` | `hi` | event |
| --- | --- | --- | --- | --- | --- |
| 1 | 1 | 1 | 0 | 1 | same parity; `lo = max(0, 1 - 2) = 0` |
| 2 | 2 | 1 | 0 | 0 | diff parity; `hi = min(1, 1 - 1) = 0` → `lo ≥ hi`, **break** |
| 3 | 3 | 1 | 0 | 1 | new chain from `j = 3` |

One break: `res = -1`, answer `-1 + 3 = 2`. Tangencies only on `(1,2)` and
`(4,5)`; the large gap `2 → 4` forces a break.

### Walkthrough: `[1, 2, 11, 12, 21, 22]` (sample 3, answer `3`)

Gaps `1, 9, 1, 9, 1`. Each block `1–2`, `11–12`, `21–22` has gap `1` internally,
but bridging across a gap of `9` makes `hi` negative on the second step of each
block, so the scan breaks twice after the first two-point chains. Three separate
pairs → `5 - 2 = 3` tangencies.

### Complexity

- **Time:** `O(n)` per test case — one left-to-right pass, `O(1)` work per gap.
- **Space:** `O(1)` besides the input array (or `O(n)` for storing `x`).

The sum of `n` over tests is at most `2 · 10^6`, so the total work fits easily.

### ideas

1. Upper bound `(n - 1)`; only adjacent pairs can be tangent in an optimal layout.
2. Consecutive tangencies form a linear system; alternating gap sums give bounds on
   the first radius of a chain.
3. Scan gaps once; when `(lo, hi)` becomes empty, decrement a counter and restart
   the chain — answer `(n - 1) + res`.
