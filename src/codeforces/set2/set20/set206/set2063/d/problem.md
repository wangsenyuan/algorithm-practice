# D. Game With Triangles

[Problem link](https://codeforces.com/problemset/problem/2063/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n + m` distinct points on the plane:
`(a_1, 0), ..., (a_n, 0)` and `(b_1, 2), ..., (b_m, 2)`.

Initially the score is `0`. You may repeatedly:

1. Choose three distinct non-collinear points
2. Add the area of their triangle to the score
3. Erase those three points

Let `k_max` be the maximum number of operations possible. For each
`k = 1, ..., k_max`, independently find `f(k)` — the maximum score achievable
with exactly `k` operations.

## Constraints

- `1 <= t <= 3 * 10^4`
- `1 <= n, m <= 2 * 10^5`
- `-10^9 <= a_i, b_i <= 10^9`
- All `a_i` are pairwise distinct; all `b_i` are pairwise distinct
- Sum of `n` and sum of `m` over all test cases do not exceed `2 * 10^5`
- All `f(x)` are integers at most `10^16`

## Input

The first line contains `t`.

For each test case:

```text
n m
a_1 ... a_n
b_1 ... b_m
```

## Output

For each test case:

- Print `k_max` on the first line
- If `k_max > 0`, print `f(1), ..., f(k_max)` on the second line

## Example

```text
Input
5
1 3
0
0 1 -1
2 4
0 100
-100 -50 0 50
2 4
0 1000
-100 -50 0 50
6 6
20 1 27 100 43 42
100 84 1 24 22 77
8 2
564040265 -509489796 469913620 198872582 -400714529 553177666 131159391 -20796763
-1000000000 1000000000

Output
1
2
2
150 200
2
1000 200
4
99 198 260 283
2
2000000000 2027422256
```

### Note

In the first test case, points are `(0,0), (0,2), (1,2), (-1,2)`. At most one
operation is possible; choosing `(0,0), (-1,2), (1,2)` gives area `2`.


### ideas
1. 假设以上下两条线上的两点作为底边的三角形个数为x, y
2. 那么有 2 * x + y <= n, 2 * y + x <= m
3. 3 * (x + y) <= n + m => (x+y)的上限是 (n + m) / 3
4. kmax = (n + m) / 3
5. 同时 x <= n, y <= m => kmax <= min(n, m, (m + n) / 3)
6. f(i) = ?
7. 计算f(2), 假设这个时候, 上面只有两个点, 但是下面有4个点
8. (a, b, c, d), (d - a) + (c - b) = (c - a) + (d - b)
9. 前半部分和后半部分进行匹配
10. f(i) = 如果有以(y = 0)边为底的三角形x个, 那么 就有以y = 2 为底的三角形(i - x)个
11. 那么计算出dp1[i](表示y=0中)组成i个底的sum, dp1[i] = 最后的i个sum - 前面i个sum(不能重叠)
12. dp2[i]同理, 那么 fp[i] = max(dp1[x] + dp2[i - x]) dp1和dp2都是单调递增的, 所以貌似可以用3分的方式进行check?

## Detailed solution

### 1. Reduce every triangle to a horizontal segment

All points lie on `y=0` or `y=2`. A non-degenerate triangle must therefore
use two points from one line and one point from the other line.

The two points on the same line form a horizontal base. Because the height is
always `2`, the area is:

```text
base_length * height / 2 = base_length
```

The horizontal position of the third point does not affect the area. Thus an
operation can be viewed as:

- select a pair from one line and earn the distance between them;
- consume any one unused point from the other line as the third vertex.

### 2. Maximum possible number of operations

Suppose we perform `k` operations. Every triangle consumes three points, so:

```text
3k <= n+m
```

Every triangle also consumes at least one point from each line, so:

```text
k <= n
k <= m
```

Therefore:

```text
kmax = min(n, m, (n+m)/3)
```

These conditions are also sufficient. The feasible-split derivation below
shows that for every `k <= kmax`, there is at least one way to distribute the
bases between the two lines.

### 3. Best score from a fixed number of bases on one line

Sort one line:

```text
c[0] < c[1] < ... < c[len(c)-1]
```

Let `best[t]` be the maximum total length of `t` disjoint segments made from
points on this line. The optimum uses the `t` smallest points as left endpoints
and the `t` largest points as right endpoints:

```text
best[t] = (c[n-1] + c[n-2] + ... + c[n-t])
        - (c[0]   + c[1]   + ... + c[t-1])
```

Why are the extremes optimal? For any chosen `2t` points, sort them. Pairing
the lower half with the upper half makes the total length equal to:

```text
sum(upper half) - sum(lower half)
```

The exact matching between the two halves does not change that sum. It is then
maximized by choosing the globally largest `t` points for the upper half and
the globally smallest `t` points for the lower half.

The `play` function computes this array with prefix sums. Let:

```text
dp1[x] = best score from x bases on y=0
dp2[y] = best score from y bases on y=2
```

### 4. Feasible split for exactly `k` operations

Let:

```text
x = number of triangles whose base is on y=0
y = number of triangles whose base is on y=2
```

Then `x+y=k`.

A triangle based on `y=0` consumes two bottom points and one top point. A
triangle based on `y=2` consumes one bottom point and two top points. Hence:

```text
2x + y <= n
x + 2y <= m
```

Substitute `y=k-x`:

```text
x + k <= n       => x <= n-k
2k - x <= m      => x >= 2k-m
0 <= x <= k
```

So the complete feasible interval is:

```text
left  = max(0, 2k-m)
right = min(k, n-k)
```

For a feasible `x`, the best score is:

```text
score_k(x) = dp1[x] + dp2[k-x]
```

We must find the maximum of `score_k(x)` on `[left,right]`.

The interval is non-empty for every `k <= kmax`. In particular:

- `k <= n` gives `0 <= n-k`;
- `k <= m` gives `2k-m <= k`;
- `3k <= n+m` gives `2k-m <= n-k`.

Thus the lower bound never exceeds both parts of the upper bound.

### 5. Why ternary search works

It is not enough that `dp1` and `dp2` are increasing. Ternary search works
because these arrays are **discrete concave**: their marginal gains never
increase.

For sorted `c`, define the gain from increasing the number of pairs from `t-1`
to `t`:

```text
gain(t) = best[t] - best[t-1]
        = c[n-t] - c[t-1]
```

The next gain is:

```text
gain(t+1) = c[n-t-1] - c[t]
```

Because the array is sorted:

```text
c[n-t-1] <= c[n-t]
c[t]     >= c[t-1]
```

therefore:

```text
gain(t+1) <= gain(t)
```

So both `dp1` and `dp2` have non-increasing first differences.

Now fix `k` and write:

```text
F(x) = dp1[x] + dp2[k-x]
```

Its first difference is:

```text
F(x+1) - F(x)
  = (dp1[x+1] - dp1[x])
  - (dp2[k-x] - dp2[k-x-1])
```

When `x` increases:

- the first marginal gain does not increase, by concavity of `dp1`;
- the index of the second marginal gain decreases, so that gain does not
  decrease; after negation, its contribution also does not increase.

Therefore `F(x+1)-F(x)` is non-increasing. Hence `F` is discrete concave: it
increases, may stay constant for a while, and then decreases. In particular,
it is unimodal, which is exactly the property needed by ternary search.

For two distinct probes `m1 < m2`:

- if `F(m1) <= F(m2)`, an optimum remains in `[m1,right]`, so values strictly
  left of `m1` may be discarded;
- if `F(m1) > F(m2)`, an optimum remains in `[left,m2]`, so values strictly
  right of `m2` may be discarded.

The code continues while there are more than three candidates. Once at most
three remain, it checks every remaining value directly.

### 6. The ternary-search boundary flaw and its fix

The loop originally used:

```go
for right-left+1 >= 3 {
```

When exactly three candidates remained, for example `[0,1,2]`, both probe
formulas produced the same middle index:

```text
m1 = 1
m2 = 1
```

The comparison was then a value against itself. Since `F(m1) <= F(m2)` was
automatically true, the code set `left=m1` and discarded the old left endpoint
without comparing it. A concave sequence can have its unique maximum at that
endpoint.

For example:

```text
n=4, m=4
a=[16, 2, -5, 17]
b=[14, -18, -19, 20]
k=2
```

The three feasible splits have scores:

| `x` | `dp1[x]` | `dp2[2-x]` | total |
|---:|---:|---:|---:|
| 0 | 0 | 71 | 71 |
| 1 | 22 | 39 | 61 |
| 2 | 36 | 0 | 36 |

The old loop discarded `x=0` and returned `61` instead of `71`.

The fix is:

```go
for right-left+1 > 3 {
```

Now ternary search only runs when the two probes are distinct. An interval of
three elements is left to the final exhaustive loop. This case is preserved by
`TestThreeFeasibleSplitsCanHaveOptimumAtLeftEndpoint` in `solution_test.go`.

### 7. Correctness summary

For each `k`:

1. Every valid triangle is represented by one base on one line and one third
   point on the other line, and its area equals its base length.
2. The feasible interval enumerates exactly all possible counts `x` of bottom
   bases.
3. For a fixed split `(x,k-x)`, `dp1[x]+dp2[k-x]` is the largest possible score
   because each line's best disjoint bases use its extreme points.
4. This score is discrete concave as a function of `x`, so the ternary search
   never removes all optimal candidates.
5. The final exhaustive scan returns the maximum among the remaining
   candidates.

Therefore the algorithm returns the maximum score for every operation count
from `1` through `kmax`.

### 8. Complexity

Let `K = kmax`.

```text
Sorting:              O(n log n + m log m)
Building dp1 and dp2: O(n + m)
All ternary searches: O(K log K)
Extra space:          O(n + m)
```
