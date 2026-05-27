# E. Hexagons

[Problem link](https://codeforces.com/problemset/problem/615/E)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Ayrat is looking for the perfect code. He decided to start his search from an
infinite field tiled by hexagons. For convenience the coordinate system is
introduced; take a look at the picture to see how the coordinates of hexagon are
defined.

Ayrat is searching through the field. He started at point `(0, 0)` and is moving
along the spiral (see second picture). Sometimes he forgets where he is now. Help
Ayrat determine his location after `n` moves.

## Solution

Ayrat walks an infinite hex grid in a fixed spiral order. Each move steps to an
adjacent hexagon. Given `n` (up to `10^18`), we must output the `(x, y)` coordinates
after exactly `n` moves, using the skew coordinate system from the statement.

The path is organized into **rings**. Ring `0` is the center hex `(0, 0)` and
requires `0` moves. For `k >= 1`, ring `k` is the next layer outward: it has `6`
sides and each side contributes exactly `k` moves, so ring `k` has `6k` moves in
total. Completing rings `1, 2, ..., k` therefore takes

```text
6(1 + 2 + ... + k) = 3k(k + 1)
```

moves. The implementation treats `count(k) = 3k(k + 1)` as the prefix length
through the end of ring `k`.

### Key observations

1. **Ring index by binary search.** The answer depends only on which ring contains
   move `n` and how far we are along that ring. Because `count(k)` grows as
   `O(k^2)`, the ring index fits in `O(log n)` via binary search.

2. **Exact ring completion lands at `(2k, 0)`.** If `n = count(k)`, Ayrat has just
   finished ring `k`. In this coordinate system that corner is always `(2k, 0)`.
   The code returns this directly (`return k * 2, 0`) without simulating every
   side.

3. **Each ring side has a fixed step vector.** On ring `k`, side `s` (in spiral
   order) always uses the same displacement vector, repeated `k` times (or fewer
   on a partial ring). The six vectors in `solution.go` are:

   ```text
   (-1,  2), (-2,  0), (-1, -2), (1, -2), (2,  0), (1,  2)
   ```

   They sum to `(0, 0)`, so a *completed* ring does not drift sideways; only the
   closed-form corner `(2k, 0)` captures where the spiral exits ring `k`.

4. **Why the walk starts at `(k * 2, 0)` inside ring `k`.** After subtracting
   `count(k - 1)` remaining moves, the code sets `x, y := k * 2, 0` and walks the
   current ring side by side. This base point is not “add up all previous rings
   from the origin”; it is the canonical start corner of ring `k` in the spiral
   parameterization used by the problem. Adding the partial side vectors on top of
   it yields the correct absolute coordinates.

5. **Overflow-safe `count`.** For `n <= 10^18`, the product `3k(k + 1)` can
   overflow 64-bit integers if computed naively for large `k`. The helper returns
   `n + 1` (a sentinel strictly larger than any valid answer) when
   `k(k + 1) > n/3`, so binary search never evaluates an overflowing product.

### Algorithm (matching `solution.go`)

**Step 0 — base case.** `n == 0` → `(0, 0)`.

**Step 1 — find the ring.** Binary search the smallest `k >= 0` with
`count(k) >= n`:

```go
k := sort.Search(1e9, func(i int) bool {
    return count(i) >= n
})
```

**Step 2 — finished exactly on a ring boundary.** If `count(k) == n`, return
`(2k, 0)`.

**Step 3 — moves left on ring `k`.** Otherwise:

```text
n := n - count(k - 1)    // 1 <= n <= 6k - 1
x, y := 2k, 0             // start corner of ring k
```

**Step 4 — consume sides in spiral order.** For each of the six direction
vectors `(dx, dy)`:

- If `n <= k`, add `(dx * n, dy * n)` and stop (partial last side).
- Else add `(dx * k, dy * k)`, subtract `k` from `n`, and continue to the next
  side.

This is the loop at lines 49–59; `k` here is the **ring number** (side length),
not a doubled index—the doubling appears only in the starting `x` coordinate
`k * 2`.

### Why `k * 2` appears

There are two separate uses of `k`:

| Expression | Meaning |
|---|---|
| `count(k) = 3k(k + 1)` | Total moves through ring `k` |
| Side length `k` | Each full side of ring `k` has `k` steps |
| Starting `x = k * 2` | Spiral geometry: the corner where ring `k` begins / ends is at x-coordinate `2k` |

The factor `2` comes from the hex coordinate system in the statement (adjacent
hexes along the “horizontal” direction change `x` by `2` when moving between
certain neighbors). It is **not** “double the ring index for the side length”; the
side length remains `k`, while the base corner uses `2k`.

### Walkthrough: `n = 3`

```text
count(0) = 0
count(1) = 6  →  binary search gives k = 1 (first ring with count >= 3)
count(1) != 3, so we are inside ring 1
n ← 3 - count(0) = 3
start at (2, 0)
```

Ring `1` has side length `1`. Three sides are fully used:

| Side | Vector | Steps | Position |
|---|---|---|---|
| 1 | `(-1,  2)` | 1 | `(1,  2)` |
| 2 | `(-2,  0)` | 1 | `(-1, 2)` |
| 3 | `(-1, -2)` | 1 | `(-2, 0)` |

Output: `(-2, 0)` (matches sample 1).

### Walkthrough: `n = 7`

```text
count(1) = 6, count(2) = 12  →  k = 2
n ← 7 - 6 = 1
start at (4, 0)
```

Only the first side of ring `2` is used (`1 <= k = 2`):

```text
(4, 0) + 1 * (-1, 2) = (3, 2)
```

Output: `(3, 2)` (matches sample 2).

### Walkthrough: exact boundary `n = 6`

`count(1) = 6 = n`, so the answer is `(2 * 1, 0) = (2, 0)` without entering the
side loop. This is the corner after completing the first full ring.

### Complexity

- **Time:** `O(log n)` for binary search plus `O(1)` for at most six sides →
  `O(log n)` overall.
- **Space:** `O(1)`.

### ideas

1. Split the path into rings. Ring `k` has length `6 * k`, and the total number of
   moves in the first `k` rings is `3 * k * (k + 1)`.
2. Binary search the smallest ring `r` such that `3 * r * (r + 1) >= n`.
3. The remaining moves inside ring `r` are `pos = n - 3 * (r - 1) * r`.
4. Each ring has six sides with `r` moves per side. From `pos`, get the side index
   and the step inside that side.
5. For every completed ring `i`, each of the six move types appears exactly `i`
   times. Add the partial counts from the current ring and multiply by the six
   direction vectors to obtain `(x, y)`.
