# Problem


- **Profiles:** `1..n` with undirected friendship edges.
- **Doubles:** `i` and `j` are doubles if every other profile `k` has **identical adjacency** to `i` and `j` (either friends with both or with neither); their mutual friendship does not matter.
- **Goal:** Count unordered pairs `(i, j)` that satisfy this property.

---

## Input

- The first line contains two integers `n` and `m`  
  (`1 ≤ n ≤ 10^6`, `0 ≤ m ≤ 10^6`) — the number of profiles and the number of friendship pairs.

- Each of the next `m` lines contains two integers `v` and `u`  
  (`1 ≤ v, u ≤ n`, `v ≠ u`) — profiles that are friends with each other.

It is guaranteed that each unordered pair of friends appears at most once, and no profile is friends with itself.

---

## Output

Print a single integer — the number of unordered pairs of profiles that are doubles.

---

## Examples

### Example 1

Input:

```text
3 3
1 2
2 3
1 3
```

Output:

```text
3
```

### Example 2

Input:

```text
3 0
```

Output:

```text
3
```

### Example 3

Input:

```text
4 1
1 3
```

Output:

```text
2
```

---

## Note

- In the first and second samples, **any two profiles** are doubles.
- In the third sample, the doubles are pairs `(1, 3)` and `(2, 4)`.

---


### ideas
1. 先构造联通块，那些独立存在的点，相互之间构成double
2. 在一个联通块中，(a, b) 如果 a和b的邻接节点，完全一致时，也满足条件
3. 

---

## Algorithm

### Open / closed neighborhoods

For a vertex `u`:

- The **open neighborhood** of `u`, denoted `N(u)`, is the set of all vertices adjacent to `u`.
- The **closed neighborhood** of `u`, denoted `N[u]`, is `N(u) ∪ {u}`.

Example:

- If `u` is connected to `{2, 5, 7}`, then:
  - `N(u) = {2, 5, 7}`
  - `N[u] = {u, 2, 5, 7}`

### Key observation

There are two cases for a pair `(u, v)`:

1. `u` and `v` are **not adjacent**
   - Then they are doubles iff:
     - `N(u) = N(v)`

2. `u` and `v` are **adjacent**
   - Then they are doubles iff:
     - `N[u] = N[v]`

Reason:

- If `u` and `v` are not connected, then for every other vertex `k`, the statement says exactly that `k` is adjacent to `u` iff it is adjacent to `v`, so their open neighborhoods must match.
- If `u` and `v` are connected, then the same equivalence holds for all other vertices, and each one also contains the other in its neighbor set. After adding the vertex itself, the two closed neighborhoods become identical.

### Counting method

1. Build the undirected graph.
2. Sort each vertex's adjacency list.
3. Compute a signature for each vertex's **open neighborhood** `N(u)`.
4. Group equal open-neighborhood signatures.
   - If a group has size `c`, it contributes `c * (c - 1) / 2`.
5. Compute a signature for each vertex's **closed neighborhood** `N[u]`.
6. Group equal closed-neighborhood signatures.
   - If a group has size `c`, it contributes `c * (c - 1) / 2`.
7. Add both contributions.

These two contributions do not overlap:

- equal open neighborhoods count only **non-adjacent** double pairs
- equal closed neighborhoods count only **adjacent** double pairs

### Complexity

- Build graph: `O(n + m)`
- Sort adjacency lists: `O(m log m)` total
- Hash/group signatures: `O(n log n)` or near-linear depending on implementation

This is fast enough for `n, m ≤ 10^6`.