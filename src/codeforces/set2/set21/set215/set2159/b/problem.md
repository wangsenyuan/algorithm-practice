# B. Rectangles

[Problem link](https://codeforces.com/problemset/problem/2159/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a binary grid `G` of dimensions `n × m`.

A rectangle is a tuple `(u, d, l, r)` that satisfies:

- `1 <= u < d <= n`
- `1 <= l < r <= m`
- Cells `(u, l)`, `(u, r)`, `(d, l)`, and `(d, r)` all contain a `1`

The area of rectangle `(u, d, l, r)` is `(d - u + 1) * (r - l + 1)`.

For each cell `(i, j)`, find the minimum area of any rectangle `(u, d, l, r)` such that
`u <= i <= d` and `l <= j <= r`. If no such rectangle exists, the answer is `0`.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n, m` and `n * m <= 250000`
- Each cell of `G` is `0` or `1`
- Sum of `n * m` over all test cases does not exceed `250000`

## Input

The first line contains `t` — the number of test cases.

For each test case:

- The first line contains `n` and `m`
- Each of the next `n` lines contains a binary string of length `m`

## Output

For each test case, print an `n × m` grid. Entry `(i, j)` is the minimum covering
rectangle area for that cell, or `0` if none exists.

## Example

```text
Input
3
3 5
10101
10100
00101
4 6
011101
010001
100010
101110
5 5
11100
10110
11111
01101
00111

Output
6 6 6 9 9
6 6 6 9 9
0 0 9 9 9
0 10 8 8 10 10
0 10 8 8 10 10
10 10 8 8 10 0
10 10 8 8 10 0
6 6 6 0 0
6 6 4 4 0
6 4 4 4 6
0 4 4 6 6
0 0 6 6 6
```

### Note

In the first test case, the grid is:

```text
1 0 1 0 1
1 0 1 0 0
0 0 1 0 1
```

There are two rectangles: `(1, 2, 1, 3)` with area `6`, and `(1, 3, 3, 5)` with area `9`.

In the third test case, six rectangles achieve the minimum area for at least one cell:

- `(1, 3, 1, 2)` area `6`
- `(1, 2, 1, 3)` area `6`
- `(3, 4, 2, 3)` area `4`
- `(2, 3, 3, 4)` area `4`
- `(3, 5, 4, 5)` area `6`
- `(4, 5, 3, 5)` area `6`


### ideas
1. 如果一个矩形区域的4个角都是1, 那么这个区域内的点, f(r, c) 可以被赋值为这个区域的面积
2. n * m = 250000, 如果 n <= m => n <= 500
3. n * m * n = 就太大了
4. 处理(r, c) 如果a(r, c) = 1, 那么这一行a(r, c1) = 1的位置, 找到c1列, 最大的r1, a(r1, c1) = 1
5. 但是这里不能保证a(r1, c) = 1
6. 完全不知道咋搞~
7. 如果在(r, c), 可以很快的知道, 存在一对(r1, c1), a(r1, c1) = 1, a(r1, c) = 1, a(r, c1) = 1
8. a(r1, c1) = 1 的时候, 放入某个结构, 当在r1行, 遇到c的时候, 怎么记录一下 c1
9. (r1 + c1)的这个结构上, 记录r和c
10. 还真是 n * n * m ~~ 1.25 * 1e8
11. 那还是可以搞的. 