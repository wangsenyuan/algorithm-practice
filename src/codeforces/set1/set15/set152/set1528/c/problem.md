Soroush and Keshi each have a labeled and rooted tree on 𝑛 vertices. Both of their trees are rooted from vertex 1.

Soroush and Keshi used to be at war. After endless decades of fighting, they finally became allies to prepare a Codeforces round. To celebrate this fortunate event, they decided to make a memorial graph on 𝑛 vertices.

They add an edge between vertices 𝑢 and 𝑣 in the memorial graph if both of the following conditions hold:

- One of 𝑢 or 𝑣 is the ancestor of the other in Soroush's tree.
- Neither of 𝑢 or 𝑣 is the ancestor of the other in Keshi's tree.

Here vertex 𝑢 is considered ancestor of vertex 𝑣 if 𝑢 lies on the path from 1 (the root) to 𝑣.

Popping out of nowhere, Mashtali tried to find the maximum clique in the memorial graph for no reason. He failed because the graph was too big.

Help Mashtali by finding the size of the maximum clique in the memorial graph.

As a reminder, a clique is a subset of vertices of the graph such that every two of them are connected by an edge.

## Input

The first line contains an integer 𝑡 (1≤𝑡≤3⋅10⁵) — the number of test cases. The description of the test cases follows.

The first line of each test case contains an integer 𝑛 (2≤𝑛≤3⋅10⁵).

The second line of each test case contains 𝑛−1 integers 𝑎₂,…,𝑎ₙ (1≤𝑎ᵢ<𝑖), where 𝑎ᵢ is the parent of vertex 𝑖 in Soroush's tree.

The third line of each test case contains 𝑛−1 integers 𝑏₂,…,𝑏ₙ (1≤𝑏ᵢ<𝑖), where 𝑏ᵢ is the parent of vertex 𝑖 in Keshi's tree.

It is guaranteed that the given graphs are trees.

It is guaranteed that the sum of 𝑛 over all test cases does not exceed 3⋅10⁵.

## Output

For each test case print a single integer — the size of the maximum clique in the memorial graph.

## Examples

**Input**

```
4
4
1 2 3
1 2 3
5
1 2 3 4
1 1 1 1
6
1 1 1 1 2
1 2 1 2 2
7
1 1 3 4 4 5
1 2 1 4 2 5
```

**Output**

```
1
4
1
3
```

## Note

In the first and third test cases, you can pick any vertex.

In the second test case, one of the maximum cliques is {2,3,4,5}.

In the fourth test case, one of the maximum cliques is {3,4,6}.

---

## Solution

Take any clique 𝐶 in the memorial graph. The vertices of 𝐶 are a subset of a path from root to some leaf in Soroush's tree. So it's sufficient to solve the task for every leaf in Soroush's tree — consider subsets of paths from the root to a leaf in Soroush's tree.

Assume a data structure that supports:

1. Insert a vertex.
2. Erase a vertex.
3. Among the vertices in it, find the largest set of vertices 𝑆 such that no two are ancestor/descendant in Keshi's tree.

Algorithm: run DFS from the root of Soroush's tree. When you enter a vertex 𝑣, insert 𝑣 (operation 1). When you leave 𝑣, erase 𝑣 (operation 2). The data structure then always represents a path from root to some vertex 𝑣 in Soroush's tree.

The answer is the maximum size of 𝑆 in operation 3 over all leaves 𝑢 of Soroush's tree (when the last operation was adding 𝑢); i.e. 𝑎𝑛𝑠 = max(𝑎𝑛𝑠, 𝑥) where 𝑥 is that size when reaching a leaf.

**When adding a vertex 𝑣:** if no vertex 𝑢 ∈ 𝑆 lies in 𝑣's subtree in Keshi's tree, then:

- If no ancestor of 𝑣 is in 𝑆, add 𝑣 to 𝑆.
- Otherwise let that ancestor be 𝑤: remove 𝑤 from 𝑆 and add 𝑣 instead.

If such a 𝑢 already exists in 𝑆, do not add 𝑣 to 𝑆 (by the greedy above).

**Notation below:** all references are to Keshi's tree unless stated.

Run DFS to compute start time 𝑠𝑡 and finish time 𝑓𝑡 for each vertex. Then: 𝑣 is an ancestor of 𝑢 ⟺ 𝑠𝑡ᵥ ≤ 𝑠𝑡ᵤ and 𝑓𝑡ᵥ ≥ 𝑓𝑡ᵤ.

Observation: for any two vertices 𝑢 and 𝑣, the segments [𝑠𝑡ᵤ, 𝑓𝑡ᵤ] and [𝑠𝑡ᵥ, 𝑓𝑡ᵥ] either do not intersect or one contains the other.

**Data structure:** maintain a set 𝑆 of pairs (𝑠𝑡ᵥ, 𝑣) representing a maximal clique.

- **Check if some vertex in 𝑆 lies in 𝑣's subtree:** take the first pair in 𝑆 with first element ≥ 𝑠𝑡ᵥ. If that pair's vertex has finish time < 𝑓𝑡ᵥ, then that vertex is in 𝑣's subtree; otherwise not.
- **Check if some ancestor of 𝑣 is in 𝑆:** take the first pair in 𝑆 with first element < 𝑠𝑡ᵥ. If an ancestor of 𝑣 is in 𝑆, it can be shown that this pair is (𝑠𝑡ᵤ, 𝑢) for that ancestor; then check whether 𝑣 is in 𝑢's subtree using the observation above.
- **Erase:** keep a history of removed elements so erasures can be done (e.g. for backtracking).
