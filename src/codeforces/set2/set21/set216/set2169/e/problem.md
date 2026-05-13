# E. Points Selection (Codeforces 2169E)

**Limits:** 2 seconds per test, 512 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2169/E](https://codeforces.com/problemset/problem/2169/E)

---

Alice and Bob play on the **XY plane**. Initially there are `n` points; the `i`-th point is at `(x_i, y_i)` with cost `c_i`.

The game has two stages:

1. **Alice** chooses some points (**possibly none**, but **not all**) and **removes** them from the field.
2. **Bob** draws an **axis-aligned rectangle** (edges parallel to the axes) so that **every remaining** point lies **inside or on the boundary**. The rectangle may **degenerate** to a segment or a single point.

The **total score** is:

- the sum of costs of points **removed** by Alice, **plus**
- the **perimeter** of Bob’s rectangle.

Alice maximizes the score; Bob minimizes it. Both play optimally. Find the resulting score.

### Perimeter convention

The perimeter is the sum of the **four** side lengths. If the “rectangle” degenerates to a segment of length `k`, the perimeter is **`2k`**. If it degenerates to a **point**, the perimeter is **`0`**.

## Input

The first line contains an integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

For each test case:

- One line: integer `n` (`1 ≤ n ≤ 3·10^5`) — number of points.
- One line: `n` integers `x_1, …, x_n` (`0 ≤ x_i ≤ 10^15`).
- One line: `n` integers `y_1, …, y_n` (`0 ≤ y_i ≤ 10^15`).
- One line: `n` integers `c_1, …, c_n` (`0 ≤ c_i ≤ 10^9`) — costs.

**Additional constraints:**

- In one test case, all points are **pairwise distinct**.
- The sum of `n` over all test cases does not exceed `3·10^5`.

## Output

For each test case, print one integer — the final score under optimal play.

## Example

### Input

```text
4
1
42
42
1000
4
5 10 5 0
0 5 10 5
1 1 1 1
4
6 7 8 9
3 3 3 3
9 0 9 0
2
1000000000 10
10 1000000000
12345 54321
```

### Output

```text
0
40
22
3999999960
```

## Note

In the **first** test case there is only one point; Alice **cannot** remove all points, so none are removed. Bob chooses a **degenerate** rectangle (a single point); perimeter **`0`**.

In the **second** test case an optimal play is: Alice removes **nothing**; Bob covers all points with rectangle corners `(0, 0)`–`(10, 10)`, perimeter `40`.

In the **third** test case an optimal play is: Alice removes the **first and third** points (cost `9 + 9`); Bob covers the rest with rectangle `(7, 3)`–`(9, 3)`, perimeter `4`. Total `18 + 4 = 22`.

## Solution

Bob always chooses the minimum-perimeter rectangle containing the remaining points.  That rectangle is exactly the bounding box of the remaining points.

So Alice is really choosing a non-empty set of points to **keep**.  If the kept set is `K`, the score is:

```text
sum(c_i for removed points) + perimeter(bounding_box(K))
```

Equivalently:

```text
sum(c_i over all points) + perimeter(bounding_box(K)) - sum(c_i for i in K)
```

The total cost of all points is constant.  Instead of writing it separately, the implementation uses the following equivalent process:

- if a point is removed, add `c_i`;
- if a point is kept, use it to define one or more sides of the bounding box, and add that side's perimeter contribution.

At the end all four sides must be defined.

### Only Boundary Points Need to Be Kept

Suppose Alice keeps a point strictly inside the bounding box, not on any of the four sides.

Removing this point:

- does not change the bounding box;
- adds its cost to the removed-cost sum.

So keeping an interior point is never better.  Therefore, in an optimal strategy, every kept point lies on at least one side of the final bounding box.

This means every kept point can be viewed as defining some non-empty subset of the four sides:

```text
left, right, bottom, top
```

A corner point may define two sides.  If the rectangle degenerates to a line or a point, one point may define even more sides.

### Perimeter as Side Contributions

Let the final rectangle have:

```text
left   = L
right  = R
bottom = B
top    = T
```

Its perimeter is:

```text
2 * (R - L) + 2 * (T - B)
```

Rewrite it as four independent side contributions:

```text
(-2L) + (2R) + (-2B) + (2T)
```

So if a point `(x, y)` defines a side, it contributes:

```text
left:   -2x
right:   2x
bottom: -2y
top:     2y
```

The implementation represents these four sides by four bits:

```text
0: left    -> -2x
1: right   ->  2x
2: bottom  -> -2y
3: top     ->  2y
```

In code this is stored using `d & 1`:

- side `0` and `1` use `x`;
- side `2` and `3` use `y`;
- side `< 2` is negative;
- side `>= 2` is positive.

The exact numbering does not matter as long as all four signed contributions are present.

### Dynamic Programming State

Let:

```text
dp[mask] = best score after processing some prefix of points,
           where the sides in mask are already defined
```

The score already includes:

- costs of removed points among the processed prefix;
- side contributions of kept points among the processed prefix.

Initially:

```text
dp[0] = 0
dp[other masks] = -infinity
```

For each point `(x, y, c)`, there are two choices.

### Transition 1: Remove This Point

If the point is removed, the side mask does not change and we add its cost:

```text
ndp[mask] = max(ndp[mask], dp[mask] + c)
```

### Transition 2: Keep This Point

If the point is kept, it must define at least one side that has not been defined yet.

Let:

```text
sub = non-empty subset of undefined sides
```

Then:

```text
newMask = mask | sub
ndp[newMask] = max(ndp[newMask], dp[mask] + contribution(point, sub))
```

where:

```text
contribution(point, sub)
```

is the sum of the signed side contributions for all sides in `sub`.

A side never needs to be defined twice.  Once its coordinate is chosen, defining it again would only add another point without changing the side represented by the DP state.  The DP therefore only assigns each side once.

### Why This Covers Every Valid Answer

Take any optimal final kept set.

Every kept point lies on at least one side.  Assign each of the four sides to one kept point lying on that side.  If a point lies on multiple still-unassigned sides, it may define multiple sides.

Now process the points in input order:

- remove every point not in the kept set;
- for the kept points, define exactly the sides assigned to them;
- if a kept point is not needed to define any side, removing it would not hurt, so such a point is unnecessary.

This sequence is one of the DP paths.  Its total value is exactly:

```text
removed costs + perimeter of the bounding box
```

Thus the DP can achieve the optimal valid score.

### Why Invalid Rectangles Do Not Win

The DP may also build invalid side assignments, for example:

```text
left coordinate > right coordinate
bottom coordinate > top coordinate
```

Such a state has a "perimeter contribution" with a negative width or height.  For example, if `L > R`, then:

```text
2R - 2L < 0
```

Swapping to a valid pair of sides would not make the value worse.  More simply, the real optimal answer is represented by some valid DP path as shown above, while invalid paths only reduce at least one dimension contribution.  Therefore the maximum over all complete masks is attained by a valid construction.

This is why the DP can ignore explicit coordinate-order checks.

### Important Detail: Negative Partial States Are Valid

A partial DP state can be negative.

Example:

```text
points: (10, 10), (20, 20)
costs:  0, 0
```

If the first point defines `left` and `bottom`, the partial contribution is:

```text
-2 * 10 - 2 * 10 = -40
```

This is a valid partial state.  Later, the second point can define `right` and `top`:

```text
2 * 20 + 2 * 20 = 80
```

The final value becomes:

```text
40
```

So the DP must not discard states just because `dp[mask] < 0`.  It should only discard states equal to `-infinity`, meaning unreachable.

This is exactly why the implementation checks:

```go
if dp[mask] > -inf
```

not:

```go
if dp[mask] >= 0
```

### Complexity

There are only `16` masks.  For every point, the implementation:

- precomputes the contribution for every side subset;
- tries removing the point;
- tries assigning every non-empty subset of still-undefined sides.

So the complexity is:

```text
O(n * 16 * 16)
```

with `O(16)` memory per test case, which is easily enough for total `n <= 3 * 10^5`.
