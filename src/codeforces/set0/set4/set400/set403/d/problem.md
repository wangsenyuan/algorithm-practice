# D. Beautiful Pairs of Numbers

[Problem link](https://codeforces.com/problemset/problem/403/D)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

## Problem

The sequence of integer pairs `(a1, b1), (a2, b2), ..., (ak, bk)` is beautiful
if:

- `1 <= a1 <= b1 < a2 <= b2 < ... < ak <= bk <= n`, where `n` is a given
  positive integer;
- all numbers `b1 - a1, b2 - a2, ..., bk - ak` are distinct.

For the given number `n` find the number of beautiful sequences of length `k`.
As the answer can be rather large, print it modulo `1000000007` (`10^9 + 7`).

## Constraints

- `1 <= t <= 2 * 10^5`
- `1 <= k <= n <= 1000`

## Input

The first line contains integer `t` — the number of test cases.

Each of the next `t` lines contains two integers `n` and `k`.

## Output

For each test print the answer modulo `10^9 + 7`, in the same order as the
tests.

## Sample

Input:

```
6
1 1
2 1
2 2
3 1
3 2
3 3
```

Output:

```
1
3
0
6
2
0
```

### Notes

- First test: only `(1, 1)`.
- Second test: `(1, 1)`, `(1, 2)`, `(2, 2)`.
- Fourth test: six length-1 sequences covering all pairs with
  `1 <= a <= b <= 3`.
- Fifth test: `(1, 1), (2, 3)` and `(1, 2), (3, 3)`.
- Third and sixth tests: no beautiful sequences.


### ideas
1. a1 <= b1 < a2 <= b2 < ... ai <= bi ... < an <= bn
2. 且 b1 - a1, b2 - a2, ... bi - ai 都要不一样
3. 这里没有要求a2 - b1 的限制(只有大于的情况)
4. c1 = a1,
5. c2 = b1 = a1 + s1
6. c3 = a2 > c2
7. c4 = b2 = a1 + s1 + ? + s2
8. ...
9. ci >= a1 + s1 + s2 + si + i / 2 (这里i是偶数)
10. 假设从 0, 1, 2, 3, ... n 中选择了k个数 (这里还必须保证 b[i] <= n)
11. dp[i][j][sum] 表示到为止, 选择了j个数, 且和是sum时的方案数
12. dp[i+1][j][sum] = dp[i][j][sum] + dp[i][j-1][sum - i] * (j-1) (有j-1个位置可以选择)
13. 然后在给定a1的情况下, bn = a1 + sum + k - 1 <= n, 还有 n - bn 的数, 可以分配到k-1个位置上去 (这里甚至可以把a1给去点, 那就是分配到k个位置上去)
14. 上面那个 sum <= n - k

## Solution

### 1. The flaw in the original implementation

The original `solve(n,k)` rebuilt a DP whose complexity was approximately:

```text
O(n * k * (n-k))
```

That can be acceptable for one test case, but the problem allows as many as
`2 * 10^5` test cases. Repeating the DP for every query cannot finish within
the time limit.

The fixed implementation precomputes every relevant `answers[n][k]` once.
After that, each call to `solve` is only an array lookup:

```go
return answers[n][k]
```

### 2. Convert every pair into an interval length

For a pair `(a_i,b_i)`, define its occupied length as:

```text
length_i = b_i - a_i + 1
```

The problem requires all values `b_i-a_i` to be distinct. Adding one to each
of them does not change distinctness, so this is equivalent to requiring all
occupied interval lengths to be distinct positive integers.

Suppose the `k` interval lengths have total:

```text
s = length_1 + length_2 + ... + length_k
```

We separate the counting into three parts:

1. choose `k` distinct positive lengths with total `s`;
2. arrange those distinct lengths from left to right;
3. distribute the `n-s` unused positions around the intervals.

### 3. Meaning of `ways[k][s]`

Define:

```text
ways[k][s] = number of unordered sets of k distinct positive lengths
             whose sum is s
```

For example:

```text
ways[2][5] = 2
```

because the two possible sets are:

```text
{1,4}
{2,3}
```

The recurrence used by the code is:

```text
ways[k][s] = ways[k][s-k] + ways[k-1][s-k]
```

To derive it, write a selected set as:

```text
x1 < x2 < ... < xk
```

and subtract one from every value. The total decreases by `k`.

- If `x1 > 1`, all `k` resulting values are still distinct positive values.
  They are counted by `ways[k][s-k]`.
- If `x1 = 1`, it becomes zero. Remove that zero; the remaining `k-1`
  values are distinct and positive, and are counted by
  `ways[k-1][s-k]`.

The base state is:

```text
ways[0][0] = 1
```

### 4. Arrange the lengths

`ways[k][s]` counts an unordered set, but the corresponding intervals appear
in a left-to-right order.

Because the lengths are all different, every selected set has exactly:

```text
k!
```

orders. This is the `fact` multiplier in the code.

### 5. Distribute the empty positions

The intervals occupy `s` positions, leaving:

```text
n - s
```

empty positions. These positions can be placed in `k+1` gaps:

```text
g0: before the first interval
g1: between interval 1 and interval 2
...
g(k-1): between interval k-1 and interval k
gk: after the last interval
```

The gap sizes are nonnegative and satisfy:

```text
g0 + g1 + ... + gk = n - s
```

By stars and bars, their number is:

```text
C((n-s)+(k+1)-1, (k+1)-1)
= C(n-s+k, k)
```

Therefore, the value stored in the answer table is:

```text
answers[n][k]
    = k! * sum over s of ways[k][s] * C(n-s+k, k)
```

### 6. How one prefix sum changes `cur`

The code does not explicitly calculate the combination above. Instead, it
starts with:

```go
cur := ways[k]
```

and repeatedly executes an in-place prefix sum:

```go
for s := 1; s < N; s++ {
    cur[s] = add(cur[s], cur[s-1])
}
```

The actual implementation currently names this loop index `n`; `s` is used
here to emphasize that it is just the running capacity index.

Because the loop goes from left to right, `cur[s-1]` has already been updated.
Thus, after one pass:

```text
newCur[n] = oldCur[0] + oldCur[1] + ... + oldCur[n]
```

### 7. Follow one original `ways[k][s]`

To see the coefficient produced by repeated prefix sums, focus on one original
entry `ways[k][s] = 1` and pretend all other entries are zero.

Initially its contribution is concentrated at index `s`:

```text
index:         s  s+1  s+2  s+3
contribution:  1   0    0    0
```

After one prefix sum:

```text
1  1  1  1
```

At index `n`, its coefficient is:

```text
1 = C(n-s, 0)
```

After two prefix sums:

```text
1  2  3  4
```

The coefficient is:

```text
n-s+1 = C(n-s+1, 1)
```

After three prefix sums:

```text
1  3  6  10
```

The coefficient is:

```text
C(n-s+2, 2)
```

In general, after `r` prefix sums, the coefficient is:

```text
C(n-s+r-1, r-1)
```

The code applies the prefix sum `k+1` times:

```go
for rep := 0; rep <= k; rep++ {
```

Substituting `r = k+1` gives:

```text
C(n-s+(k+1)-1, (k+1)-1)
= C(n-s+k, k)
```

This is exactly the stars-and-bars coefficient needed for the `k+1` gaps.

### 8. Formal proof of the repeated-prefix-sum identity

Let `P_r[n]` be the value at index `n` after `r` prefix sums. We claim:

```text
P_r[n] = sum from s=0 to n of
         ways[k][s] * C(n-s+r-1, r-1)
```

For `r=1`, one prefix sum gives:

```text
P_1[n] = ways[k][0] + ways[k][1] + ... + ways[k][n]
```

This matches the formula because `C(n-s,0)=1`.

Now assume the formula holds after `r` prefix sums. One more prefix sum gives:

```text
P_(r+1)[n] = sum from x=0 to n of P_r[x]
```

Substitute the induction hypothesis and group together the contribution of
each original index `s`:

```text
P_(r+1)[n]
    = sum over s of ways[k][s]
      * sum from x=s to n of C(x-s+r-1, r-1)
```

By the hockey-stick identity:

```text
sum from x=s to n of C(x-s+r-1, r-1)
    = C(n-s+r, r)
```

Therefore:

```text
P_(r+1)[n]
    = sum over s of ways[k][s] * C(n-s+r, r)
```

This is exactly the claimed formula for `r+1`. Setting `r=k+1` proves that
the final `cur[n]` is:

```text
cur[n] = sum over s of ways[k][s] * C(n-s+k, k)
```

Finally, the code multiplies by `k!`:

```go
answers[n][k] = mul(cur[n], fact)
```

### 9. Small example

Take `n=3` and `k=2`.

The only set of two distinct positive interval lengths that fits is:

```text
{1,2}
```

Its total length is `s=3`, so:

```text
ways[2][3] = 1
```

There are `2! = 2` possible orders. There are no unused positions, and hence
only one gap distribution:

```text
C(3-3+2, 2) = C(2,2) = 1
```

Thus:

```text
answers[3][2] = 1 * 2 * 1 = 2
```

### 10. Correctness summary

Every beautiful sequence uniquely determines:

1. its unordered set of `k` distinct positive interval lengths;
2. an ordering of those lengths, counted by `k!`;
3. a distribution of the unused positions among `k+1` gaps, counted by
   `C(n-s+k,k)`.

Conversely, every choice of these three objects constructs exactly one
beautiful sequence. Therefore, `answers[n][k]` counts every valid sequence
exactly once.

### 11. Complexity

Let `N=1000` and `K=50` be the maximum relevant limits.

- Computing `ways` costs `O(KN)`.
- For each `k`, the code performs `k+1` prefix sums over all `n`, for a total
  of `O(K^2N)`.
- Each test case is answered in `O(1)`.
- The `ways` and `answers` tables use `O(KN)` space.

```text
Precomputation: O(K^2 N)
Per query:      O(1)
Space:          O(KN)
```
