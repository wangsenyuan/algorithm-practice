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

### ideas
1. a * b / 2 = n * m / k 
2. a * b = 2 * n * m / k
3. 如果 2 * n * m 不能整除k => false
4. a * b = 2 * n * m / k 是整数的情况下