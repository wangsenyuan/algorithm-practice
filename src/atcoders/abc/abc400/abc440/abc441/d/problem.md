# D - Paid Walk (ABC441)

**Contest:** [ABC441](https://atcoder.jp/contests/abc441) — AtCoder Beginner Contest 441  
**Task:** [https://atcoder.jp/contests/abc441/tasks/abc441_d](https://atcoder.jp/contests/abc441/tasks/abc441_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

There is a directed graph, not necessarily simple, with `N` vertices numbered
`1, 2, ..., N` and `M` edges.

The `i`-th edge (`1 <= i <= M`) goes from vertex `U_i` to vertex `V_i` with cost
`C_i`. The out-degree of each vertex is at most `4`.

Find all vertices `v` (`1 <= v <= N`) that satisfy the following condition:

- There exists a path from vertex `1` to vertex `v` that traverses exactly `L`
  edges.
- The sum of the costs of the traversed edges is at least `S` and at most `T`.

If the same edge is traversed multiple times, each traversal is counted, and its
cost is added each time.

The out-degree of a vertex `u` is the number of edges going out from `u`. Even if
multiple edges go to the same vertex, they are counted separately.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= M <= 2 * 10^5`
- `1 <= L <= 10`
- `1 <= S <= T <= 10^9`
- `1 <= U_i, V_i <= N`
- `1 <= C_i <= 10^8`
- The out-degree of each vertex is at most `4`.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N M L S T
U_1 V_1 C_1
U_2 V_2 C_2
...
U_M V_M C_M
```

## Output

Print the vertices that satisfy the condition in ascending order, separated by
spaces.

If there are no vertices that satisfy the condition, print an empty line.

## Sample Input 1

```text
5 8 3 80 100
1 2 20
1 3 70
2 1 30
2 5 10
3 2 10
3 4 30
3 5 20
5 1 70
```

## Sample Output 1

```text
1 5
```

For vertex `1`, the path `1 -> 2 -> 5 -> 1` traverses exactly `3` edges and has
total cost `20 + 10 + 70 = 100`, so it satisfies the condition.

For vertex `5`, the path `1 -> 3 -> 2 -> 5` traverses exactly `3` edges and has
total cost `70 + 10 + 10 = 90`, so it satisfies the condition.

Vertices `2`, `3`, and `4` do not satisfy the condition.

## Sample Input 2

```text
10 1 1 1 100
2 3 1
```

## Sample Output 2

```text

```

If there are no vertices that satisfy the condition, print an empty line.

## Sample Input 3

```text
2 5 3 1 100
1 1 1
2 2 100
1 2 1
1 2 1
1 2 100
```

## Sample Output 3

```text
1 2
```

The graph may contain self-loops and multiple edges. In this test case, the
out-degrees from vertices `1` and `2` are `4` and `1`, respectively.
