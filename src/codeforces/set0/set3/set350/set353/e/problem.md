# Problem

You have a directed acyclic graph `G` consisting of `n` vertices numbered from `0` to `n - 1`. The graph contains `n` edges numbered from `0` to `n - 1`.

Edge `i` connects vertices `i` and `(i + 1) mod n`, and it can be directed in either direction:

- from `i` to `(i + 1) mod n`, or
- from `(i + 1) mod n` to `i`.

Operation x mod y means taking the remainder after dividing number x by number y.

Let's call two vertexes u and v in graph G comparable if the graph contains a path either from u to v or from v to u. We'll assume that an antichain is a set of vertexes of graph G, where any two distinct vertexes are not comparable. The size of an antichain is the number of vertexes in the corresponding set. An antichain is maximum if the graph doesn't have antichains of a larger size.

Your task is to find the size of the maximum antichain in graph G.

## Input

The first line contains a binary string `s = s0s1...s(n-1)` (`2 ≤ n ≤ 10^6`), consisting of characters `0` and `1`. The length of the string `n` corresponds to the number of vertices and edges in graph `G`.

If `si = 0`, then the edge between vertices `i` and `(i + 1) mod n` is directed from `i` to `(i + 1) mod n`. Otherwise (`si = 1`) it is directed in the opposite direction.

It is guaranteed that the given graph is acyclic.

## Output

Print a single integer — the size of the maximum antichain of graph `G`.

## Examples

### Example 1

Input

```text
001
```

Output

```text
1
```

### Example 2

Input

```text
110010
```

Output

```text
3
```

## Note

Consider the first sample. The edges of `G` are: `0 → 1`, `1 → 2`, `0 → 2`. We can choose the set of vertices `{0}` as a maximum antichain. We cannot choose an antichain of larger size.

## Thoughts

Because the graph is built from a directed cycle and is guaranteed to be acyclic, the edge directions must alternate between increasing and decreasing along the cycle, grouped into maximal runs of equal bits.

After compressing the binary string into run lengths:

- each run behaves like a monotone directed path
- inside one run, all vertices are mutually comparable
- therefore an antichain can contain at most one "free" choice contributed by each run

So the original graph turns into a cycle of runs, and the problem becomes a small DP on these runs.

For a run of length `len`:

- if we do not choose the boundary vertex at its end, then this run can contribute `1` only when `len > 1` (we can choose one interior vertex)
- if we choose that boundary vertex, then this run contributes exactly `1`, and the previous boundary must be unchosen

This gives a standard two-state DP over the run sequence:

- `dp0`: best answer when the current boundary vertex is not chosen
- `dp1`: best answer when the current boundary vertex is chosen

Since the runs form a cycle, we handle it with two cases:

1. the shared boundary between the first and last runs is not chosen
2. that shared boundary is chosen

Take the maximum of these two cases.

The whole solution is linear in `n`, since run compression and the DP both scan the string once.