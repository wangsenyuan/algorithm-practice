# E. Tree-String Problem

[Problem link](https://codeforces.com/problemset/problem/291/E)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a rooted tree with `n` vertices (root is `1`). On each edge from
parent `p_v` to child `v` (`2 <= v <= n`) there is a non-empty string `s_v`
written top-to-bottom.

A **position** is a pair `(v, x)` meaning the `x`-th character of `s_v`
(`0 <= x < |s_v|`).

A pair of positions `(v, x)` and `(u, y)` defines a string `z` if the path from
the first position to the second always goes **down** the tree; `z` is the
concatenation of all letters along that path in order.

Given a string `t`, count the number of pairs of positions that define exactly
`t`.

## Constraints

- `2 <= n <= 10^5`
- Strings on edges are non-empty and consist of lowercase English letters
- `|t| >= 2`, lowercase English letters
- Total number of English letters in the input is at most `3 * 10^5`

## Input

The first line contains integer `n`.

The next `n - 1` lines describe edges to vertices `2, 3, ..., n` in order. The
line for vertex `v` contains `p_v` and string `s_v`.

The last line contains string `t`.

## Output

Print a single integer — the number of pairs of positions that define `t`.

## Examples

### Sample 1

```text
Input
7
1 ab
5 bacaba
1 abacaba
2 aca
5 ba
2 ba
aba

Output
6
```

### Sample 2

```text
Input
7
1 ab
5 bacaba
1 abacaba
2 aca
5 ba
2 ba
bacaba

Output
4
```

## Solution

### KMP state along a root-to-node path

While traversing the tree from the root, concatenate the labels of the edges on
the current path. We run KMP on this conceptual string without constructing it.

For the pattern `t`, a KMP state `q` means that the processed path ends with the
prefix `t[0:q]`. When DFS enters an edge, it starts with the state inherited from
the parent, processes every character on that edge, and passes the resulting
state to the child. Different children receive the same parent state, so their
searches remain independent.

Whenever consuming a character completes all of `t`, one valid pair of
positions ends at that character. We add one to the answer and continue from
the prefix-function value `pi[|t|-1]`, which also handles overlapping matches.

### Why ordinary KMP can time out on a tree

For one ordinary text, the usual KMP transition is amortized `O(1)`:

```text
while q > 0 and t[q] != ch:
    q = pi[q-1]
```

That amortization does not carry across tree branches. Suppose a long edge has
matched a long prefix of `t`, and its endpoint has many children. Every child
starts from the same large state. If the first character on every child edge
mismatches, every child independently walks through the same long chain of
prefix-function links. Thus the worst case can be proportional to the number
of branches times `|t|`, even though the total number of edge characters is
small enough.

### Building the KMP automaton

There are only `|t|` useful states and 26 possible next characters. Precompute
the result of every pair `(state, character)` in `transition[state][c]`.

For a state `q` and character `ch`:

```text
if t[q] == ch:
    next = q + 1
else if q == 0:
    next = 0
else:
    next = transition[pi[q-1]][ch]
```

The last case is exactly the ordinary KMP fallback: after failing to extend
`t[0:q]`, try extending its longest proper prefix that is also a suffix. Since
`pi[q-1] < q`, that smaller state's transitions have already been computed
when states are processed in increasing order.

For example, let `t = "abab"`. Its states and selected transitions are:

| State | Matched suffix | Read `a` | Read `b` |
|------:|----------------|---------:|---------:|
| 0 | `""` | 1 | 0 |
| 1 | `"a"` | 1 | 2 |
| 2 | `"ab"` | 3 | 0 |
| 3 | `"aba"` | 1 | match, then 2 |

After the last transition, `"abab"` has been found and state 2 is retained
because `"ab"` is both a suffix of the match and a prefix of the pattern. This
allows the two overlapping occurrences in `"ababab"` to be counted.

### Encoding a completed match

Each table entry needs to store both the next state and whether a match was
completed. The implementation packs both into one `int32`:

```text
no match:  code = nextState
match:     code = -nextState - 1
```

During DFS:

```text
code = transition[state][character]
if code < 0:
    answer++
    state = -code - 1
else:
    state = code
```

The extra minus one is necessary because a completed match may continue from
state 0; plain `-0` would be indistinguishable from a normal transition to
state 0.

### Correctness

The automaton transition is equivalent to one ordinary KMP transition. On a
match it advances by one character, and on a mismatch it uses the same
prefix-function fallback until it finds the longest prefix of `t` that is a
suffix of the processed path.

By induction over the characters on a root-to-node path, the state carried by
DFS is therefore exactly the KMP state of that path string. A match is counted
precisely when the current path suffix equals `t`. Its ending position uniquely
determines its starting position `|t|-1` characters earlier, so every valid
pair is counted once and no invalid pair is counted.

### Complexity

Let `L` be the total length of all edge labels and `m = |t|`.

- Building the prefix function takes `O(m)` time.
- Building all automaton transitions takes `O(26m)` time and memory.
- DFS processes every vertex and every edge character once, taking `O(n + L)`
  time.

The total complexity is `O(26m + n + L)` time and `O(26m + n + L)` memory.
The transition table uses `int32`, so even for `m = 3 * 10^5` it occupies about
31.2 MB.
