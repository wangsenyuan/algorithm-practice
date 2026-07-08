# F. Cherry Tree

[Problem link](https://codeforces.com/problemset/problem/2184/F)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a rooted tree with `n` vertices. The root is vertex `1`.

Every leaf contains one cherry. You need to collect all cherries by repeatedly choosing a
vertex `v` and shaking it. When `v` is shaken, cherries fall from every leaf in the
subtree of `v`; if `v` is a leaf, its own cherry falls.

No leaf may lose its cherry more than once. If a cherry has already fallen and another
chosen vertex would make it fall again, the tree breaks and the process is invalid.

The number of shaken vertices must be a multiple of `3`.

For each test case, determine whether it is possible to collect all cherries.

## Input

The first line contains `t` (`1 <= t <= 10^4`) -- the number of test cases.

For each test case:

- The first line contains `n` (`2 <= n <= 2 * 10^5`).
- The next `n - 1` lines contain edges `u v`
  (`1 <= u, v <= n`, `u != v`).

Each test case is a tree. The sum of `n` over all test cases is at most `2 * 10^5`.

## Output

For each test case, print `YES` if all cherries can be collected while shaking a multiple
of three vertices. Otherwise, print `NO`.

The answer is case-insensitive.

## Example

```text
Input
3
4
1 2
1 3
1 4
3
1 2
1 3
9
1 2
3 1
2 4
5 2
5 6
3 7
8 3
8 9

Output
YES
NO
YES
```

## Solution

Root the tree at vertex `1`.

If we shake a vertex `u`, then all leaves in `u`'s subtree are collected at once. After
that, no descendant of `u` and no ancestor of `u` can also be shaken in the same valid
plan, because that would make at least one leaf lose its cherry twice. So for each subtree
there are only two kinds of choices:

- shake the subtree root itself, using exactly `1` shaken vertex;
- do not shake the subtree root, and independently solve all child subtrees.

The final condition only asks whether the total number of shaken vertices is divisible by
`3`, so the DP only needs counts modulo `3`.

Let `dp[u][r]` mean: inside the subtree of `u`, assuming no ancestor of `u` is shaken, it
is possible to collect every cherry in this subtree using a number of shaken vertices
congruent to `r (mod 3)`.

For every vertex `u`:

1. Combine all children with knapsack over residues:

   ```text
   cur[(i + j) mod 3] |= previous[i] && dp[child][j]
   ```

   This is the case where `u` is not shaken, so each child subtree must be handled by its
   own valid plan.

2. Set `dp[u][1] = true`, because shaking `u` itself always collects all leaves in its
   subtree using exactly one shaken vertex.

Leaves naturally have no child contribution, and step 2 makes `dp[leaf][1] = true`, which
matches shaking the leaf itself.

After DFS finishes, the answer is `YES` iff `dp[1][0]` is true.

Each edge is processed once, and every merge checks only `3 * 3` residue pairs. Therefore
the time complexity is `O(n)`, and the memory complexity is `O(n)`.
