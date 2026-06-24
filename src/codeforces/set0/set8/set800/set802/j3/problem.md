# J3. Send the Fool Further! (hard)

[Problem link](https://codeforces.com/problemset/problem/802/J3)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Heidi is terrified by your estimate and she found it unrealistic that her friends would
collaborate to drive her into debt. She expects that, actually, each person will just
pick a random friend to send Heidi to. (This randomness assumption implies, however,
that she can now visit the same friend an arbitrary number of times...) Moreover, if
a person only has one friend in common with Heidi (i.e., if that person is in a leaf of
the tree), then that person will not send Heidi back (so that Heidi's travel will end
at some point).

Heidi also found unrealistic the assumption that she can make all the travels in one
day. Therefore now she assumes that every time she travels a route between two friends,
she needs to buy a new ticket. She wants to know: how much should she expect to spend
on the trips?

For what it's worth, Heidi knows that Jenny has at least two friends.

Heidi's `n` friends are labeled `0` through `n - 1`, and their network of connections
forms a tree. Jenny is given the number `0`.

## Input

The first line contains the number of friends `n` (`3 <= n <= 10^5`). The next `n - 1`
lines each contain three space-separated integers `u`, `v` and `c` (`0 <= u, v <= n - 1`,
`1 <= c <= 10^4`) meaning that `u` and `v` are friends and the cost for traveling
between `u` and `v` is `c` (paid every time!).

It is guaranteed that the social network of the input forms a tree.

## Output

Assume that the expected cost of the trips is written as an irreducible fraction `a / b`
(that is, `a` and `b` are coprime). Output a single integer — the value of `a * b^{-1}`
modulo `10^9 + 7` (that is, an integer between `0` and `10^9 + 6`).

## Examples

### Input 1

```text
3
0 1 10
0 2 20
```

### Output 1

```text
15
```

### Input 2

```text
4
0 1 3
0 2 9
0 3 27
```

### Output 2

```text
13
```

### Input 3

```text
7
0 1 3
0 5 7
1 2 2
1 3 1
1 4 5
5 6 8
```

### Output 3

```text
400000019
```

### Input 4

```text
11
1 0 6646
2 0 8816
3 2 9375
4 2 5950
5 1 8702
6 2 2657
7 2 885
8 7 2660
9 2 5369
10 6 3798
```

### Output 4

```text
153869806
```

### Input 5

```text
6
0 1 8
0 2 24
1 3 40
1 4 16
4 5 8
```

### Output 5

```text
39
```

## Note

In the first example, with probability `1 / 2` Heidi will go to `1` from `0`, and with
probability `1 / 2` she will go to `2`. In the first case the cost would be `10`, and
in the second it would be `20`. After reaching `1` or `2` she will stop, as `1` and `2`
are leaves of the social tree. Hence, the expected cost she has to pay is
`10 * 1 / 2 + 20 * 1 / 2 = 15`.

In the third example, the expected cost is `81 / 5`. You should output `400000019`.

## Solution

Root the tree at Jenny, vertex `0`. Let `deg(u)` be the original degree of vertex `u`.
A leaf vertex other than `0` stops the process, so its remaining expected cost is `0`.

For every non-root vertex `u`, define `E[u]` as the expected future cost when Heidi is
currently at `u`. The transition at `u` chooses uniformly from all `deg(u)` neighbors,
including the parent. If the edge from `u` to its parent `p` has cost `w(u, p)`, then:

```text
E[u] = (w(u, p) + E[p]
        + sum over children v of u (w(u, v) + E[v])) / deg(u)
```

This equation is not a normal tree DP because `E[u]` depends on its parent `E[p]`.
However, after rooting the tree, each child subtree only depends on its parent through
one value. Therefore express every non-root state as a linear function of its parent:

```text
E[u] = A[u] * E[parent[u]] + B[u]
```

For a leaf, `E[u] = 0`, so `A[u] = 0` and `B[u] = 0`.

For an internal non-root vertex `u`, substitute every child expression
`E[v] = A[v] * E[u] + B[v]` into the expectation formula:

```text
E[u] = (w(u, p) + E[p]
        + sum(w(u, v) + A[v] * E[u] + B[v])) / deg(u)
```

Move the terms containing `E[u]` to the left:

```text
E[u] * (deg(u) - sum A[v])
  = E[p] + w(u, p) + sum(w(u, v) + B[v])
```

So:

```text
A[u] = 1 / (deg(u) - sum A[v])
B[u] = (w(u, p) + sum(w(u, v) + B[v])) / (deg(u) - sum A[v])
```

The values `A[u]` and `B[u]` are computed bottom-up. The implementation builds a DFS
order from root `0`, then processes that order in reverse so every child has already
computed its linear expression.

The root has no parent. It also chooses uniformly from all its neighbors, and each
neighbor `v` satisfies `E[v] = A[v] * E[0] + B[v]`. Thus:

```text
E[0] = sum(w(0, v) + A[v] * E[0] + B[v]) / deg(0)
```

After moving the `E[0]` terms:

```text
E[0] = sum(w(0, v) + B[v]) / (deg(0) - sum A[v])
```

All divisions are performed modulo `1e9 + 7` by multiplying with modular inverses.
The problem guarantees that these denominators are invertible, so this modular form
matches the requested irreducible fraction value.

## Correctness

We prove that the algorithm returns the expected total travel cost from vertex `0`.

First, consider any leaf vertex `u != 0`. Once Heidi reaches `u`, the statement says
the process stops. Therefore the remaining expected cost is exactly `0`, and the
algorithm sets `A[u] = B[u] = 0`, which represents `E[u] = 0`.

Now consider an internal non-root vertex `u`, and assume all children `v` of `u`
already satisfy `E[v] = A[v] * E[u] + B[v]`. From vertex `u`, Heidi chooses uniformly
among all `deg(u)` neighbors. The next step is either to the parent, paying the parent
edge cost and then having expectation `E[parent[u]]`, or to a child `v`, paying that
edge cost and then having expectation `E[v]`. The expectation equation written in the
solution follows directly from this law of total expectation. Substituting the already
correct child expressions and solving for `E[u]` gives exactly the formulas used by
the algorithm for `A[u]` and `B[u]`. Hence the computed expression for `u` is correct.
By reverse DFS order, this induction applies to every non-root vertex.

Finally, at root `0`, the same total-expectation equation applies, except there is no
parent transition. Every neighbor expression has already been proven correct, so
substituting them and solving for `E[0]` gives the exact expected cost from Jenny.
The algorithm returns that value modulo `1e9 + 7`, so it returns the required answer.

## Complexity

The tree is traversed once to build parent order and once in reverse to compute the DP.
Each edge is inspected `O(1)` times.

Time complexity: `O(n)`.

Memory complexity: `O(n)`.
