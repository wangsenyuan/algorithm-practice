# E. Fair and Square

[Problem link](https://codeforces.com/problemset/problem/2241/E)

**Contest:** [Codeforces Round (contest 2241)](https://codeforces.com/contest/2241)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A tree is an undirected connected graph with no cycles.

You are given a tree having `n` vertices. Each vertex `i` has an integer value `a_i`
written on it.

For any two vertices `u` and `v` (`u != v`), define `p(u, v)` as the product of the values
written on the vertices lying on the unique simple path from `u` to `v`.

An unordered triplet of three distinct vertices `{u, v, w}` is called good if and only if
`p(u, v) * p(v, w) * p(w, u)` is a perfect square.

Determine the number of good unordered triplets in the given tree.

A simple path from vertex `u` to vertex `v` is a sequence of distinct vertices
`u = x_0, x_1, ..., x_k = v` such that there exists an edge between `x_{i-1}` and `x_i`
for all `1 <= i <= k`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case begins with an integer `n` (`3 <= n <= 2 * 10^5`) — the number of vertices.

The second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^6`) — the
integer values written on the vertices.

Each of the next `n - 1` lines contains two integers `u, v` (`1 <= u, v <= n`), denoting
an edge of the tree. It is guaranteed that the edges form a tree.

It is guaranteed that the sum of `n` over all the test cases does not exceed `2 * 10^5`.

## Output

For each test case output the number of good triplets in the tree.

## Example

### Input

```text
4
5
1 1 1 1 1
1 2
2 3
2 4
4 5
10
1 2 3 4 5 6 7 8 9 10
1 3
2 6
6 7
5 4
8 3
3 4
4 6
9 1
10 2
6
12 6 3 18 9 2
3 4
4 5
2 6
6 1
4 2
8
3 16 9 1 8 16 4 9
2 1
3 1
4 3
3 5
6 3
4 7
8 1
```

### Output

```text
10
48
0
40
```

## Note

For the first test case, all the unordered triplets of three distinct vertices are good
(`C(5, 3) = 10`).

For the second test case, `{2, 5, 8}` is a good triplet.


## ideas
1. in image tree.svg, p(u, v) = a[2] * a[1] * a[4]
2. p(u, w) = a[2] * a[6]
3. p(v, w) = a[4] * a[1] * a[2] * a[6]
4. 只有a[2]出现了3次, 其他都是出现了2次
5. 如果 w = 3, 
6. p(u, w) = a[2] * a[1] * a[3]
7. p(v, w) = a[4] * a[1] * a[3]
8. 只有a[1]出现了3次 
9. 如果节点p, 是 (u, v, w) 的中心点, 它会出现3次
