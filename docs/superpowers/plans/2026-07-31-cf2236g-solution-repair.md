# CF2236G Solution Repair Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Complete the existing rooted-tree/LCA solution so it counts all and only path subsegments without repeated set bits.

**Architecture:** Preserve binary lifting, propagated forbidden-depth boundaries, and root-path prefix sums. Partition each query into two arm-only counts and an LCA-containing count; construct the latter from at most 21 cumulative bitmask states per arm, with zero runs aggregated by depth gaps.

**Tech Stack:** Go standard library, package-local tests, deterministic randomized brute-force oracle.

---

### Task 1: Add independent correctness tests

**Files:**
- Modify: `src/codeforces/set2/set22/set223/set2236/g/solution_test.go`

- [ ] **Step 1: Add a direct brute-force oracle**

Add helpers that reconstruct a query path with BFS parent pointers, then
enumerate every subsegment and count it when `xor == sum`. Keep this helper
independent of `fp`, LCA, masks, and binary lifting.

- [ ] **Step 2: Add focused regressions**

Add cases for:

```text
two-node all-zero path: expected 3
opposite branches [1,0,1]: expected 5
same-arm path [1,1,0]: expected 4
sibling branches where DFS state must not leak
long zero runs with sparse nonzero values
```

- [ ] **Step 3: Add a deterministic randomized oracle test**

For `n = 2..10`, generate many parent-linked random trees, values below `2^5`,
and all ordered pairs of distinct vertices. Compare `solve` against the direct
oracle. Use a fixed `math/rand` seed so failures are reproducible.

- [ ] **Step 4: Run tests and verify RED**

Run:

```bash
env GOCACHE=/tmp/learn-go-gocache \
  go test ./src/codeforces/set2/set22/set223/set2236/g/ -count=1
```

Expected: focused and randomized tests fail because the current solution
omits arm-only subsegments, forgets boundaries, leaks DFS state, and treats
cross-LCA choices as an unconditional product.

### Task 2: Repair rooted-tree preprocessing

**Files:**
- Modify: `src/codeforces/set2/set22/set223/set2236/g/solution.go`

- [ ] **Step 1: Use the propagated boundary invariant**

Initialize every `fp` entry to `-1`. On DFS entry:

```go
if p != u {
    fp[u] = fp[p]
}
for bit := range 20 {
    if a[u]>>bit&1 == 1 {
        fp[u] = max(fp[u], lowDepthOfBit[bit])
        lowDepthOfBit[bit] = depth[u]
    }
}
```

- [ ] **Step 2: Make DFS rollback frame-local**

Inside each DFS invocation, use:

```go
var keep [20]int
copy(keep[:], lowDepthOfBit)
defer copy(lowDepthOfBit, keep[:])
```

so recursive children cannot overwrite the caller's snapshot.

- [ ] **Step 3: Precompute count and nonzero links**

Use `[]int64` for `pr`, add `depth[u]-fp[u]` to the parent's prefix, and store
the strict nearest nonzero ancestor:

```go
if p == u {
    prevNZ[u] = -1
} else if a[p] != 0 {
    prevNZ[u] = p
} else {
    prevNZ[u] = prevNZ[p]
}
```

### Task 3: Implement the query partition

**Files:**
- Modify: `src/codeforces/set2/set22/set223/set2236/g/solution.go`

- [ ] **Step 1: Find monotone arm boundaries**

Implement a helper that returns the deepest ancestor of `x` at or below `w`
whose propagated boundary satisfies either `fp[z] <= depth[w]` or the strict
variant `fp[z] < depth[w]`. When `x` is bad, binary-lift upward to the
shallowest bad vertex and return its parent.

- [ ] **Step 2: Count arm-only subsegments**

For deepest non-strict-good vertex `g`, return:

```go
d := int64(depth[g] - depth[w])
triangle := d * (d + 1) / 2
tail := pr[x] - pr[g]
return triangle + tail
```

The prefix difference covers exactly the deeper endpoints with
`fp[z] > depth[w]`.

- [ ] **Step 3: Build cumulative mask states**

For the deepest strict-good endpoint, collect nonzero ancestors strictly below
`w`, reverse them into outward order, and aggregate endpoint counts by depth
gaps:

```go
type state struct {
    mask int
    ways int64
}
```

The initial mask is zero and includes endpoint `w`; each encountered nonzero
value changes the cumulative mask.

- [ ] **Step 4: Combine LCA-containing states**

Replace the Cartesian product of arm lengths with:

```go
for _, left := range leftStates {
    for _, right := range rightStates {
        if left.mask&right.mask == 0 {
            cross += left.ways * right.ways
        }
    }
}
```

Return:

```go
countArm(x, w) + countArm(y, w) + cross
```

### Task 4: Verify correctness and limits

**Files:**
- Test: `src/codeforces/set2/set22/set223/set2236/g/solution_test.go`

- [ ] **Step 1: Format changed Go files**

Run:

```bash
gofmt -w \
  src/codeforces/set2/set22/set223/set2236/g/solution.go \
  src/codeforces/set2/set22/set223/set2236/g/solution_test.go
```

- [ ] **Step 2: Run the focused and randomized suite**

Run:

```bash
env GOCACHE=/tmp/learn-go-gocache \
  go test ./src/codeforces/set2/set22/set223/set2236/g/ -count=1
```

Expected: all samples, regressions, and deterministic random comparisons pass.

- [ ] **Step 3: Add and run a maximum zero-chain regression**

Build a chain with `n = 100000`, query its endpoints, and require:

```go
int64(n) * int64(n+1) / 2
```

Run it separately and confirm it finishes within the package's normal test
timeout.

- [ ] **Step 4: Check whitespace and scope**

Run:

```bash
git diff --check -- \
  src/codeforces/set2/set22/set223/set2236/g/solution.go \
  src/codeforces/set2/set22/set223/set2236/g/solution_test.go
git status --short
```

Expected: no whitespace errors and no unrelated files modified by the repair.
