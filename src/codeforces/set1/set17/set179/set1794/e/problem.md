# E. Good Vertices (Tree Labeling)

You are given an **unweighted tree** of \(n\) vertices numbered \(1\) to \(n\), and a list of \(n - 1\) integers \(a_1, a_2, \ldots, a_{n-1}\). A tree is a connected undirected graph without cycles.

You will use **each** element of the list to label **exactly one** vertex. **No** vertex gets two labels. The **one** vertex that does not receive a label from the list may be labeled with **any** integer.

A vertex \(x\) is called **good** if it is possible to assign the list to vertices so that for **every** vertex \(i\), its label equals the **distance** between \(x\) and \(i\) in the tree (where distance is the minimum number of edges on a path between two vertices).

Find **all** good vertices.

## Input

The first line contains one integer \(n\) (\(2 \le n \le 2 \cdot 10^5\)) — the number of vertices.

The second line contains \(n - 1\) integers \(a_1, a_2, \ldots, a_{n-1}\) (\(0 \le a_i < n\)) — the given list.

Then \(n - 1\) lines follow; each contains two integers \(u\) and \(v\) (\(1 \le u, v \le n\)) — an edge between \(u\) and \(v\). The edges form a tree.

## Output

In the first line print the number of good vertices.

In the second line, print the indices of all good vertices in **ascending** order.

## Examples

### Example 1

**Input**

```
6
2 1 0 1 2
1 2
2 3
2 4
4 5
4 6
```

**Output**

```
2
2 4
```

### Example 2

**Input**

```
5
1 2 1 2
1 2
3 2
3 4
5 4
```

**Output**

```
1
3
```

### Example 3

**Input**

```
3
2 2
1 2
2 3
```

**Output**

```
0
```

### ideas
1. A vertex `x` is good iff the multiset of distances from `x` to all `n` vertices is equal to the multiset `a` plus one extra number.
   - The extra number corresponds to the unique vertex whose label is arbitrary.
   - So we only need to compare distance frequencies, not the exact vertex-to-label matching.

2. Let `cnt_x[d]` be the number of vertices whose distance from `x` equals `d`.
   - Then `x` is good iff there exists some `k` such that
     `cnt_x[d] = freq_a[d]` for all `d != k`, and `cnt_x[k] = freq_a[k] + 1`.
   - In particular, distance `0` can appear at most once, so if `freq_a[0] > 1` the answer is immediately empty.

3. Encode the whole distance-frequency array by a polynomial hash:
   - `H(x) = sum cnt_x[d] * base^d`
   - choose `base = n`, and use double mod / double hash to avoid collisions.
   - Likewise, encode the target multiset
     `T = sum freq_a[d] * base^d`.

4. Because one extra label may be arbitrary, every valid root must satisfy
   - `H(x) = T + base^k` for some `k`.
   - So we can precompute all `n` possible hashes
     `T + base^0, T + base^1, ..., T + base^(n-1)`
     and store them in a set.

5. Now the main task is to compute `H(x)` for every root `x`.
   - This is a rerooting DP.
   - If `dp[u]` is the hash of the subtree of `u`, where distance is measured from `u`,
     then
     `dp[u] = 1 + sum dp[v] * base` over all children `v`.
   - Reason: every node inside child `v`'s subtree becomes one step farther from `u`.

6. In the second DFS, reroot from parent to child.
   - Suppose `cur` is the full hash for root `u`.
   - For a child `v`, remove `v`'s contribution `dp[v] * base` from `cur`.
   - The remaining part is the "outside" contribution seen from `u`.
   - When moving root from `u` to `v`, all those outside nodes become one step farther, so multiply by `base` and pass it down.
   - Therefore rerooting is:
     `outside_v = cur - dp[v] * base`
     and the full hash at `v` is
     `dp[v] + outside_v * base`.

7. For each vertex, check whether its full hash is in the valid set.
   - If yes, this vertex is good.

8. Complexity:
   - `O(n)` DFS/rerooting.
   - `O(n)` candidate hashes.
   - Overall `O(n)` time and `O(n)` memory, ignoring hash map constants.
