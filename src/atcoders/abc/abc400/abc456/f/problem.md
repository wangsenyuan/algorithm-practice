# F - Plan Holidays (ABC456)

**Contest:** [ABC456](https://atcoder.jp/contests/abc456) — AtCoder Beginner Contest 456  
**Task:** [https://atcoder.jp/contests/abc456/tasks/abc456_f](https://atcoder.jp/contests/abc456/tasks/abc456_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

Takahashi is trying to decide his schedule for `N` days. Initially, none of the days
are holidays.

He can repeat either of the following operations any number of times:

- Choose an integer `i` between `1` and `N`, inclusive, and make day `i` a holiday.
  This operation costs `A_i`.
- Choose an integer `i` between `2` and `N - 1`, inclusive, such that both day `i - 1`
  and day `i + 1` are already holidays, and make day `i` a holiday. This operation is
  free.

Find the minimum total cost required to create a consecutive block of `K` or more
holidays.

`T` test cases are given; solve each of them.

## Constraints

- `1 <= T <= 2 * 10^5`
- `1 <= K <= N <= 2 * 10^5`
- `1 <= A_i <= 10^9`
- All input values are integers.
- The sum of `N` over all test cases is at most `2 * 10^5`.

## Input

The input is given from Standard Input in the following format:

```text
T
case_1
case_2
...
case_T
```

Here, `case_i` denotes the input for the `i`-th test case. Each test case is given in
the following format:

```text
N K
A_1 A_2 ... A_N
```

## Output

Output `T` lines. The `i`-th line should contain the answer for the `i`-th test case.

## Sample Input 1

```text
3
5 2
3 1 4 1 5
6 4
24 3 22 39 4 29
15 7
220651272 302798780 874479994 657822311 613294668 479624013 241168404 610547619 762548286 256160531 823041612 951553052 226556081 649525901 153805947
```

## Sample Output 1

```text
2
29
1902064780
```

For the first test case, a consecutive block of at least two holidays can be created by
performing operations as follows:

- Make day 2 a holiday using the first type of operation. This costs 1.
- Make day 4 a holiday using the first type of operation. This costs 1.
- Make day 3 a holiday using the second type of operation. This is free.

The total cost of this sequence of operations is 2. It is impossible to create a
consecutive block of two or more holidays with a cost less than 2, so output 2.


### ideas
1. 肯定是一段连续的，l...r, r - l + 1 >= k
2. 要么是k，要么是k+1

## Solution

First fix the final consecutive block as `[l, r]`. The two endpoints `l` and
`r` must be paid directly, because the free operation can only create an
interior holiday from two already-holiday neighbors.

Inside `(l, r)`, a day can be left unpaid if both of its neighbors are eventually
holidays. However, two adjacent unpaid days cannot both be created for free:
neither of them can be created first, because each is waiting for the other one
to already be a holiday. Therefore, among the interior days, the unpaid days
must form an independent set on a path, meaning no two chosen unpaid days are
adjacent.

For a fixed interval `[l, r]`, if we initially pay for every day in the interval,
the cost is:

```text
A_l + A_{l+1} + ... + A_r
```

Then every interior day that we decide not to pay for saves `A_i`, as long as
the saved days are non-adjacent. So the minimum cost for this interval is:

```text
sum(l, r) - maximum_weight_independent_set(l+1, r-1)
```

For intervals of length `2`, there is no interior, so the cost is simply
`A_l + A_r`.

### Why Only Length `K` Or `K + 1`

All costs are positive. If an optimal block has length greater than `K + 1`, we
can shrink it from the left or right while keeping paid endpoints:

- If the neighbor of an endpoint is paid, remove just the endpoint.
- If the neighbor is unpaid, then the next day after it must be paid, because
  unpaid days are never adjacent; remove the endpoint and that unpaid neighbor.

This preserves the same "endpoints paid, unpaid interior non-adjacent" structure
and never increases the cost. Repeating this process reduces any longer block to
length `K` or `K + 1`. Therefore it is enough to check every interval of length
`K` and every interval of length `K + 1`.

### Segment Tree For Maximum Saved Cost

The remaining subproblem is range maximum-weight independent set on a path:
for many ranges `[L, R]`, find the maximum total `A_i` we can choose with no two
adjacent indices.

The code builds a segment tree. Each node stores four values:

```text
d00: maximum sum with no endpoint restriction
d01: maximum sum when the right endpoint of this segment is forbidden
d10: maximum sum when the left endpoint of this segment is forbidden
d11: maximum sum when both endpoints are forbidden
```

For a leaf, `d00 = A_i`, and the other states are `0`, because if the only point
is forbidden, we choose nothing.

When merging left node `a` and right node `b`, the only extra condition is at the
boundary between them: the right endpoint of `a` and the left endpoint of `b`
cannot both be chosen. For example:

```text
d00 = max(a.d00 + b.d10, a.d01 + b.d00)
```

Either forbid the left endpoint of `b`, or forbid the right endpoint of `a`.
The other three merged states are computed in the same way while also preserving
the requested outer endpoint restrictions.

For an interval `[l, r]`, the code computes:

```text
total = prefix[r+1] - prefix[l]
saved = query(l+1, r-1).d00
cost  = total - saved
```

Then it takes the minimum over all length-`K` and length-`K+1` intervals.

### Complexity

Building the segment tree takes `O(N)`. Each range query takes `O(log N)`, and
there are `O(N)` candidate intervals per test case. The total complexity is:

```text
O(N log N)
```

over each test case, with `O(N)` memory. Since the sum of `N` over all test cases
is bounded, this fits within the constraints.

## Official Solution

The current `solve()` uses the official formulation. It is still a DP for the
same structure, but instead of querying a segment tree for every interval, it
turns the DP transition of one day into a small matrix and maintains the product
of a sliding window.

### 2-State DP

Process the days of a candidate block from left to right. For each position, the
only information we need to carry is whether the previous day has already been
paid.

- State `0`: the previous day is not paid, so the current day must be paid.
- State `1`: the previous day is paid, so the current day may be paid or left
  unpaid.

If the current day is unpaid, then the next day becomes forced, because two
adjacent unpaid days are impossible. This is why only two states are needed.

For one day with cost `x = A_i`, the transition matrix is:

```text
| inf  0 |
|  x   x |
```

which is stored as:

```go
Mat{inf, 0, nums[i], nums[i]}
```

Its meaning is:

- `inf`: impossible transition.
- `0`: leave this day unpaid.
- `x`: pay this day.

So a single matrix describes all valid ways to process one day under the two
boundary states.

### Min-Plus Matrix Product

If matrix `A` describes the transitions of one segment, and matrix `B`
describes the transitions of the next segment, then the combined segment uses:

```text
C[i][j] = min(A[i][0] + B[0][j], A[i][1] + B[1][j])
```

This is exactly the code in `matMul()`. It is ordinary DP composition:

- `+` combines costs.
- `min` picks the better middle state.

It is called a min-plus matrix product because it looks like matrix
multiplication with:

- multiplication replaced by `+`
- addition replaced by `min`

The product of all matrices in a window gives the DP transition for that whole
window.

### Why `tmp[2]` and `tmp[3]`

Let `tmp` be the matrix product of the current length-`K` window.

- `tmp[2]` is the cost of using exactly this length-`K` window as the final
  block.
- `tmp[3]` is the cost when the window is allowed to rely on one extra paid day
  on its left.

So the code checks:

```go
ans = min(ans, tmp[2])
if l > 0 {
	ans = min(ans, nums[l-1]+tmp[3])
}
```

The first line handles length `K`. The second line handles length `K+1` by
paying day `l-1` explicitly and combining it with the best structure on
`[l, l+K-1]`.

### Why FoldableQueue Is Linear

The matrix product is associative, so for a sliding window we only need to
maintain the product of all matrices in order.

The foldable queue stores the window using two stacks:

- `back`: newly pushed matrices.
- `front`: matrices ready to be popped from the left.

For `back`, `last` stores the product of all elements in order, so `Push()` is
`O(1)`.

When `front` becomes empty, `Pop()` rebuilds it from `back`. Suppose:

```text
back = [a, b, c, d]
```

After this loop:

```go
for i := len(q.back) - 1; i > 0; i-- {
	q.back[i-1] = q.op(q.back[i-1], q.back[i])
}
```

`back` becomes:

```text
[a*b*c*d, b*c*d, c*d, d]
```

These are suffix products. After copying them into `front`, we get:

- `front[0]`: product of all current front elements
- `front[1]`: product after removing the first one
- and so on

So popping from the left is now just removing `front[0]`, and the next suffix
product is already prepared.

Finally, the product of the whole queue is:

```text
product(front) * product(back)
```

which `AllProd()` returns as:

```go
q.op(q.front[0], q.last)
```

Each matrix is:

- pushed once,
- moved from `back` to `front` at most once,
- popped once.

So the total rebuild work over the whole sliding process is linear. That is why
the official solution runs in `O(N)` per test case instead of `O(N log N)`.
