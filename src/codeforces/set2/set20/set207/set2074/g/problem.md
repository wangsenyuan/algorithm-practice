# G. Game With Triangles: Season 2

[Problem link](https://codeforces.com/problemset/problem/2074/G)

time limit per test: 4 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

> The Frontman greets you to this final round of the survival game.

There is a regular polygon with `n` sides (`n >= 3`). The vertices are indexed as `1, 2, ..., n`
in clockwise order. On each vertex `i`, the pink soldiers have written a positive integer `a_i`.
With this regular polygon, you will play a game defined as follows.

Initially, your score is `0`. Then, you perform the following operation any number of times to
increase your score.

Select 3 different vertices `i`, `j`, `k` that you have not chosen before, and draw the triangle
formed by the three vertices.

- Then, your score increases by `a_i * a_j * a_k`.
- However, you cannot perform this operation if the triangle shares a **positive common area**
  with any of the triangles drawn previously.

Your objective is to maximize the score. Find the maximum score you can get from this game.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t`
(`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` — the number of vertices
(`3 <= n <= 400`).

The second line of each test case contains `a_1, a_2, ..., a_n` — the integers written on
vertices (`1 <= a_i <= 1000`).

It is guaranteed that the sum of `n^3` over all test cases does not exceed `400^3`.

## Output

For each test case, output the maximum score you can get on a separate line.

## Example

### Input

```text
6
3
1 2 3
4
2 1 3 4
6
2 1 2 1 1 1
6
1 2 1 3 1 5
9
9 9 8 2 4 4 3 5 3
9
9 9 3 2 4 4 8 5 3
```

### Output

```text
6
24
5
30
732
696
```

## Note

- Test case 1: only one triangle can be drawn. Maximum score `6` uses vertices `(1, 2, 3)`.
- Test case 2: only one triangle. Maximum score `24` uses vertices `(1, 3, 4)`.
- Test case 3: two triangles can be drawn; maximum score is `5`.
- Test case 4: two triangles are possible, but one triangle gives maximum score `30` (vertices
  `(2, 4, 6)`).
- Test case 5: three triangles; maximum score `732`.

### ideas
