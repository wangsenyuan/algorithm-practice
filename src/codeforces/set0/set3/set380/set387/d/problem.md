# Problem Description

George loves graphs. Most of all, he loves interesting graphs. We will assume that a directed graph is interesting, if it meets the following criteria:

1. The graph doesn't contain any multiple arcs;
2. There is vertex v (we'll call her the center), such that for any vertex of graph u, the graph contains arcs (u, v) and (v, u). Please note that the graph also contains loop (v, v).
3. The outdegree of all vertexes except for the center equals two and the indegree of all vertexes except for the center equals two. The outdegree of vertex u is the number of arcs that go out of u, the indegree of vertex u is the number of arcs that go in u. Please note that the graph can contain loops.

However, not everything's that simple. George got a directed graph of n vertices and m arcs as a present. The graph didn't have any multiple arcs. As George loves interesting graphs, he wants to slightly alter the presented graph and transform it into an interesting one. In one alteration he can either remove an arbitrary existing arc from the graph or add an arbitrary arc to the graph.

George wonders: what is the minimum number of changes that he needs to obtain an interesting graph from the graph he's got as a present? Help George and find the answer to the question.

## Input

The first line contains two space-separated integers n and m (2 ≤ n ≤ 500, 1 ≤ m ≤ 1000) — the number of vertices and arcs in the presented graph.

Each of the next m lines contains two space-separated integers aᵢ, bᵢ (1 ≤ aᵢ, bᵢ ≤ n) — the descriptions of the graph's arcs. Pair (aᵢ, bᵢ) means that the graph contains an arc from vertex number aᵢ to vertex number bᵢ. It is guaranteed that the presented graph doesn't contain multiple arcs.

Assume that the graph vertices are numbered 1 through n.

## Output

Print a single integer — the answer to George's question.

## Examples

### Input
```
3 7
1 1
2 2
3 1
1 3
3 2
2 3
3 3
```

### Output
```
0
```

### Input
```
3 6
1 1
2 2
3 1
3 2
2 3
3 3
```

### Output
```
1
```

### Input
```
3 1
2 2
```

### Output
```
6
```

## Note

For more information about directed graphs, please visit: http://en.wikipedia.org/wiki/Directed_graph

In the first sample the graph already is interesting, its center is vertex 3.

## Solution

To solve this problem you should know about bipartite matching.

### Step-by-Step Explanation

#### Step 1: Understand the Target Graph Structure

For an "interesting" graph with center vertex `c`:
- **Center vertex `c`**: Must have arcs to/from all other vertices (including itself)
  - Required arcs: `(c, u)` and `(u, c)` for all vertices `u` (including `c`)
  - Total: `2n - 1` arcs (since `(c, c)` is counted once)
- **Non-center vertices**: Each must have exactly:
  - Outdegree = 2 (exactly 2 outgoing arcs)
  - Indegree = 2 (exactly 2 incoming arcs)

#### Step 2: Try Each Vertex as Center

For each vertex `i` from 1 to n, try it as the potential center and calculate the minimum changes needed.

#### Step 3: Separate Center-Related Arcs

When considering vertex `i` as center:
- Count arcs involving center: `cntWithI` = number of arcs `(i, u)` or `(u, i)`
- Count other arcs: `Other = m - cntWithI` (arcs between non-center vertices)

#### Step 4: Build Bipartite Graph for Non-Center Vertices

For the remaining `n-1` non-center vertices, we need to satisfy:
- Each vertex needs exactly 2 incoming arcs (indegree = 2)
- Each vertex needs exactly 2 outgoing arcs (outdegree = 2)

**Key insight**: We model this as a bipartite matching problem:
- **Left side**: Represents vertices that need incoming arcs (indegree requirements)
- **Right side**: Represents vertices that need outgoing arcs (outdegree requirements)
- **Edge in bipartite graph**: If original graph has arc `(u, v)` where both `u` and `v` are non-center, this creates an edge from left-`u` to right-`v`

**Why bipartite matching?**
- Each non-center vertex appears twice on the left (needs 2 incoming arcs) and twice on the right (needs 2 outgoing arcs)
- A matching in this bipartite graph represents which existing arcs we can keep
- Maximum matching = maximum number of existing arcs we can preserve

#### Step 5: Calculate Changes for Current Center

For center vertex `i`, the total cost is calculated as:

```
cost = 2*(n-1) + 1 - cntWithI + Other - len + (n - 1) - len
```

Breaking down each component:

1. **`2*(n-1) + 1 - cntWithI`**: Changes needed for center connections
   - We need exactly `2n - 1` arcs involving the center (arcs `(c, u)` and `(u, c)` for all `u`, where `(c, c)` is counted once)
   - Currently we have `cntWithI` arcs involving the center
   - Net changes needed: `2n - 1 - cntWithI`
   - Note: This can be positive (need to add arcs) or negative (need to remove arcs)

2. **`Other - len`**: Arcs to remove from non-center subgraph
   - We have `Other` existing arcs between non-center vertices
   - Maximum matching size `len` = maximum arcs we can keep
   - Arcs to remove: `Other - len`

3. **`(n - 1) - len`**: Arcs to add for non-center vertices
   - Each of the `n-1` non-center vertices needs exactly 2 outgoing arcs
   - Total outgoing arcs needed: `2 * (n - 1)`
   - But wait: each arc `(u, v)` contributes 1 to `u`'s outdegree and 1 to `v`'s indegree
   - Since we also need indegree = 2 for each vertex, and one arc satisfies both requirements:
     - We need exactly `n - 1` arcs total between non-center vertices (not `2*(n-1)`)
   - We can keep `len` arcs from the matching
   - Arcs to add: `(n - 1) - len`

**Simplified formula**:
```
cost = 2n - 2 + 1 - cntWithI + Other - len + n - 1 - len
     = 3n - 2 - cntWithI + Other - 2len
```

**Why `n-1` arcs for non-center vertices?**
- We need 2 outgoing arcs per non-center vertex = `2*(n-1)` total outdegree
- We need 2 incoming arcs per non-center vertex = `2*(n-1)` total indegree  
- Each arc contributes to both outdegree and indegree
- So we need exactly `n-1` arcs to satisfy both requirements simultaneously

#### Step 6: Find Minimum Across All Centers

Try each vertex as center, calculate the cost, and return the minimum.

### Why Bipartite Matching Works

- **Maximum matching** finds the maximum number of existing arcs we can keep while satisfying degree constraints
- Each matched edge represents an arc we can preserve
- Unmatched requirements must be satisfied by adding new arcs
- This minimizes the total number of changes (removals + additions)

### Important Notes

- Self-loops are allowed, which simplifies the problem significantly
- The bipartite graph has `2(n-1)` vertices on each side (each non-center vertex needs 2 incoming and 2 outgoing arcs)
- We use the Hungarian algorithm or DFS-based matching to find maximum bipartite matching

### Complexity

- Time: O(n²m) - For each of n centers, we run bipartite matching on a graph with O(m) edges
- Space: O(n + m) - Adjacency list and matching arrays