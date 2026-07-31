# CF2236G Solution Repair Design

## Goal

Complete the existing `fp`/prefix-sum/LCA solution frame so every query counts
exactly the path subsegments whose values have no repeated set bit.

## Preserved Direction

- Root the tree once and build binary-lifting ancestors.
- During DFS, track the latest depth of each of the 20 bits.
- Split every query path at its LCA.
- Use prefix sums for subsegments contained in one LCA arm.

## Rooted-Path Boundary

For every vertex `u`, define `fp[u]` as the greatest forbidden starting depth
for a valid root-path suffix ending at `u`.

```text
fp[u] = max(
    fp[parent[u]],
    latest depth of every bit set in a[u]
)
```

The root boundary and all unseen-bit depths are `-1`. Consequently,
`depth[u] - fp[u]` is the number of valid root-path suffixes ending at `u`.
The DFS must use a frame-local snapshot of the latest-bit depths so sibling
subtrees cannot affect one another.

Let `pr[u]` be the root-path prefix sum of these ending counts.

## Arm-Only Subsegments

For an arm endpoint `x` and LCA `w`, count subsegments strictly below `w`.
Because propagated `fp` is monotone down a rooted path, split the arm at the
deepest vertex `g` satisfying:

```text
fp[g] <= depth[w]
```

Vertices from the child of `w` through `g` contribute:

```text
d * (d + 1) / 2
```

where `d = depth[g] - depth[w]`. Deeper vertices have
`fp[z] > depth[w]` and contribute through a `pr` difference.

## Subsegments Containing the LCA

For each arm, find the farthest vertex `g` satisfying the strict condition:

```text
fp[g] < depth[w]
```

This is the longest internally valid extension that includes `w`. Enumerate
only its nonzero vertices by following nearest-nonzero-ancestor links. A valid
extension contains at most 20 nonzero vertices.

Create cumulative mask states for endpoints from `w` through `g`. Consecutive
zero vertices do not change the mask; aggregate them as the state's `ways`
using depth differences.

For left and right states, a crossing subsegment is valid exactly when:

```text
left.mask & right.mask == 0
```

The crossing count is the sum of `left.ways * right.ways` over compatible
state pairs. The masks exclude `a[w]`; the strict `fp` boundary already proves
that each individual arm is compatible with `w`.

## Query Result

The three categories form a disjoint partition:

```text
left-arm-only
+ right-arm-only
+ valid subsegments containing w
```

All counts and prefix sums use `int64`.

## Complexity

- Preprocessing: `O(n log n + 20n)`
- Each query:
  - `O(log n)` boundary searches on the two arms
  - `O(20)` state construction per arm
  - `O(20^2)` compatible-state combination
- Memory: `O(n log n)`

Long zero paths are handled through depth-difference multiplicities and are
never traversed vertex by vertex per query.

## Validation

- Preserve all supplied samples.
- Add focused regressions for:
  - an all-zero path;
  - a repeated bit on opposite LCA branches;
  - an earlier same-arm collision followed by zeros;
  - sibling-subtree contamination of DFS bit state;
  - long zero runs around sparse nonzero vertices.
- Compare the optimized solver against an independent brute-force enumerator
  on many small random trees, values, and queries.
- Include a maximum-size zero-chain regression to exercise `int64` counting
  and long-path performance.
