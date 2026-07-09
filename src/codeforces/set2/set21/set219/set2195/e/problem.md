# E. Idiot First Search

[Problem link](https://codeforces.com/problemset/problem/2195/E)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

There is a binary tree of `n + 1` vertices (`n` is odd), labeled `0, 1, ..., n`.
The root is vertex `0`. Vertex `0` is the parent of vertex `1`. Every other vertex has
either `2` children or `0` children. All vertices start with nothing written on them.

Bob is lost on some vertex and wants to reach vertex `0` using **Idiot First Search**.
When Bob is on vertex `v` (`1 <= v <= n`):

- If `v` is a leaf, Bob always moves to the parent of `v`.
- Otherwise:
  - If nothing is written on `v`, write `'L'` and move to the left child.
  - If `'L'` is written on `v`, overwrite it to `'R'` and move to the right child.
  - If `'R'` is written on `v`, erase it and move to the parent of `v`.

Each move takes exactly `1` second.

For each starting vertex `k = 1, 2, ..., n`, find the time to reach vertex `0`,
modulo `10^9 + 7`.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 300001`, `n` is odd
- For each vertex `i = 1..n`, either `l_i = r_i = 0` (leaf), or `l_i` and `r_i` are
  its left and right children
- The input defines a valid binary tree as described
- Sum of `n` over all test cases does not exceed `300001`

## Input

The first line contains `t` — the number of test cases.

For each test case:

- The first line contains `n`.
- Each of the next `n` lines contains two integers `l_i` and `r_i` — the children of
  vertex `i`.

## Output

For each test case, print `n` integers `τ_1, τ_2, ..., τ_n` — the escape times from
vertices `1..n`, modulo `10^9 + 7`.

## Example

```text
Input
3
1
0 0
5
2 3
0 0
4 5
0 0
0 0
7
2 3
4 5
0 0
6 7
0 0
0 0
0 0

Output
1
9 10 14 15 15
13 22 14 27 23 28 28
```

### Note

In the first test case, Bob starts at vertex `1` and reaches `0` in `1` second.

In the second test case, escaping from vertex `3` takes `14` seconds. One possible
move sequence is:

```text
3L→4X→3R→5X→3X→1L→2X→1R→3L→4X→3R→5X→3X→1X→0
```

where the letter above each arrow is the mark on the current vertex before the move
(`X` means nothing written).


### ideas
1. f(v) 表示,如果从v出发, 访问完所有的子节点, 并回到v得时间
2. 如果v是一个leaf, f(v) = 0
3. else f(v) = 1 + f(left(v)) + 1 + 1 + f(right(v)) + 1 = 4 + f(left(v)) + f(right(v)) 
4. 这里没有考虑回到0的时间
5. f(v) + 1 到了 parent(v), 如果v是parent的左子树, 那么就可以当作parent出发(并回到0)
6. 如果 v是parent的子树, 那么 f(parent) + f(v) + 1