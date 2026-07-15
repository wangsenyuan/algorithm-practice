# D. Vasya and Triangle

[Problem link](https://codeforces.com/problemset/problem/1030/D)

**Contest:** [Codeforces Round #505 (rated, Div. 1 + Div. 2, based on VK Cup 2018 Final)](https://codeforces.com/contest/1030)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Vasya has three integers `n`, `m` and `k`. He wants three integer points `(x1, y1)`, `(x2, y2)`,
`(x3, y3)` such that:

- `0 <= x1, x2, x3 <= n`
- `0 <= y1, y2, y3 <= m`
- the area of the triangle formed by these points equals `n * m / k`

Help Vasya find such points if possible. If there are multiple solutions, print any of them.

## Input

A single line contains three integers `n`, `m`, `k` (`1 <= n, m <= 10^9`, `2 <= k <= 10^9`).

## Output

If there are no such points, print `NO`.

Otherwise print `YES` on the first line. The next three lines should each contain integers `xi yi` —
the coordinates of one point.

Letter case does not matter for `YES` / `NO`.

## Example

### Input

```text
4 3 3
```

### Output

```text
YES
1 0
2 3
4 1
```

### Input

```text
4 4 7
```

### Output

```text
NO
```

### Note

In the first sample, the required area is `n * m / k = 4`. One valid triangle is shown in the sample
output.

In the second sample, no triangle has area `n * m / k = 16 / 7`.

## Solution

Choose an axis-aligned right triangle with vertices:

```text
(0, 0), (a, 0), (0, b)
```

Its area is `a*b/2`, so we need:

```text
a * b = 2 * n * m / k
```

We also need `a <= n` and `b <= m`.

### Remove `k` from the two side lengths

Start with:

```text
a = n
b = m
```

First remove as much of `k` as possible from `a`:

```text
g = gcd(a, k)
a /= g
k /= g
```

Then do the same with `b` and the remaining `k`:

```text
g = gcd(b, k)
b /= g
k /= g
```

Let the remaining value still be called `k`. It is now coprime with both
original side lengths after all usable factors have been removed.

For the required doubled area to be an integer, this remaining `k` must divide
the factor `2` in `2*n*m`. Therefore:

- if `k > 2`, no construction exists;
- if `k == 2`, the current `a` and `b` are already the answer;
- if `k == 1`, the current product is half of the required doubled area, so
  one of `a` and `b` must be doubled.

In the last case, doubling always fits. The original input has `k >= 2`, so at
least one of the two GCD operations removed a factor of at least `2`.

- If the factor removed from `n` was at least `2`, then `a <= n/2`, so `2*a`
  fits within `n`.
- Otherwise that factor was `1`, so a factor of at least `2` was removed from
  `m`. Hence `b <= m/2`, and `2*b` fits within `m`.

This is exactly what the implementation checks:

```text
if 2*a <= n:
    a *= 2
else:
    b *= 2
```

Finally, output `(0, 0)`, `(a, 0)`, and `(0, b)`.

### Correctness Proof

Let `g1` and `g2` be the factors removed from `k` by the GCD operations on
`n` and `m`. After the reductions:

```text
a = n / g1
b = m / g2
remaining k = original k / (g1 * g2)
```

Therefore:

```text
2 * n * m / original k = 2 * a * b / remaining k
```

If the remaining `k` is greater than `2`, it cannot divide `2`, so the
required doubled area is not an integer and no integer-coordinate triangle can
exist.

If the remaining `k` is `2`, the required doubled area is `a*b`, exactly the
doubled area of the constructed triangle.

If the remaining `k` is `1`, the required doubled area is `2*a*b`. As shown
above, the algorithm can double either `a` or `b` without leaving the
rectangle, so the constructed triangle again has exactly the required area.

Thus the algorithm returns `NO` precisely when construction is impossible,
and otherwise returns three valid points forming a triangle of area
`n*m/k`.

### Complexity

The algorithm performs two GCD computations.

- Time: `O(log(max(n, m, k)))`
- Extra space: `O(1)`
