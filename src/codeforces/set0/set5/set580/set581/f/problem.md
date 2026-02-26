# Problem

It's election time in Berland. The favorites are the **zublicanes** and the **mumocrates**. Both parties run campaigns on $n$ main squares of the capital. Each square hosts demonstrations of exactly one party; all squares must be used. The capital management will assign the squares to the two parties.

The squares and roads form a tree: $(n - 1)$ bidirectional roads connect pairs of squares so that there is a unique path between any two squares. A **dead end** square is one that is connected by a road to only one other square.

**Constraints:**

- The mayor requires that the number of dead end squares assigned to the first party equals the number assigned to the second party. (The total number of dead end squares is even.)
- Minimize the number of roads whose endpoints are in different parties (i.e., edges between squares with different parties).

Find the minimum number of such roads.

---

## Input

- Line 1: integer $n$ ($2 \le n \le 5000$) — number of squares.
- Next $n - 1$ lines: two integers $x$, $y$ ($1 \le x, y \le n$, $x \ne y$) — a road between squares $x$ and $y$. Squares are labeled $1$ to $n$.

The graph is a tree. The number of dead end squares is even.

## Output

Print one integer — the minimum number of roads that connect squares with different parties.

---

## Examples

**Input**

```
8
1 4
2 4
3 4
6 5
7 5
8 5
4 5
```

**Output**

```
1
```

### Why the answer is 1 (Example 1 in depth)

The tree has 8 nodes and 7 edges:

- **Structure:** Two “hubs” (4 and 5) connected by the edge 4–5. Hub 4 is connected to dead ends 1, 2, 3; hub 5 is connected to dead ends 6, 7, 8.

```
    1   2   3          6   7   8
     \  |  /             \  |  /
      \ | /               \ | /
        [4] =============== [5]
              (edge 4-5)
```

- **Dead ends:** Nodes 1, 2, 3, 6, 7, 8 each have degree 1 → 6 dead ends. The mayor requires 3 for one party and 3 for the other.

- **Internal nodes:** 4 and 5 have degree ≥ 2; they can be assigned to either party.

- **Optimal assignment:** Put one side of the “bridge” in one party and the other side in the other party, e.g.  
  - Party A (zublicanes): 1, 2, 3, 4  
  - Party B (mumocrates): 5, 6, 7, 8  

  Then the only road with one endpoint in each party is **4–5**, so the number of “cross-party” roads is **1**. Any other valid assignment (same dead-end balance) either keeps this single cross edge or adds more (e.g. splitting the two hubs would still give at least the edge 4–5, and mixing dead ends across the bridge would add more cross edges). So the minimum is **1**.

---

**Input**

```
5
1 2
1 3
1 4
1 5
```

**Output**

```
2
```

### Why the answer is 2 (Example 2 in depth)

The tree is a **star** with 5 nodes:

- **Center:** node 1 (degree 4).
- **Leaves (dead ends):** nodes 2, 3, 4, 5 (each degree 1) → 4 dead ends.

```
        2
        |
    3 — 1 — 4
        |
        5
```

- **Constraint:** 2 dead ends must go to one party, 2 to the other. The center 1 can be assigned to either party.

- **Edges:** Every edge is 1–2, 1–3, 1–4, or 1–5. So every “cross-party” road is between the center and a leaf.

- **Optimal assignment:** Put the center 1 in one party and two leaves in the other, e.g.  
  - Party A: 1, 2, 3  
  - Party B: 4, 5  

  Then the only cross-party edges are **1–4** and **1–5** → **2** roads. (Equivalently, Party A: 1, 4, 5 and Party B: 2, 3 gives cross edges 1–2 and 1–3.)

- **Why we can’t get 1:** The center is adjacent to all four leaves. Exactly two leaves must be in the other party (to balance dead ends 2–2). So at least 2 edges from the center go to the other party. Hence the minimum number of cross-party roads is **2**.

---

## Editorial

**Setup.** Let $c$ be the number of **leaves** in the tree (vertices of degree 1). The statement guarantees $c$ is even. If the graph has only 2 vertices, the answer is $1$. Otherwise, pick a non-leaf vertex and **root the tree** at it.

**DP state.**

- **$z_1[v][\mathit{cnt}][\mathit{col}]$** — minimum number of edges with different-colored endpoints in the subtree of $v$, given: $v$ has color $\mathit{col}$ ($0$ or $1$), and exactly $\mathit{cnt}$ leaves in the subtree of $v$ have color $0$.
- **$z_2[s][\mathit{cnt}][\mathit{col}]$** — when we have processed some children of $v$ and are at child $s$: we still need to distribute $\mathit{cnt}$ leaves with color $0$ among the subtrees of $v$'s children (including $s$ and those after $s$).

**Base case.** If $v$ is a leaf, the value is easy to compute.

**Transition for $z_1$.** If $v$ is not a leaf, $z_1[v][\mathit{cnt}][\mathit{col}] = z_2[s][\mathit{cnt}][\mathit{col}]$, where $s$ is the first child of $v$ in the adjacency list.

**Transition for $z_2$.** When computing $z_2[s][\mathit{cnt}][\mathit{col}]$, iterate over the color $\mathit{ncol}$ of child $s$ and the number $a$ of leaves with color $0$ in the subtree of $s$ ($0 \le a \le \text{leaves}(s)$, and $a$ must be consistent with subtree sizes). Then:
$$z_2[s][\mathit{cnt}][\mathit{col}] = \min_{\mathit{ncol},\,a} \Big( z_2[\mathit{ns}][\mathit{cnt}-a][\mathit{col}] + z_1[s][a][\mathit{ncol}] + \mathbf{1}[\mathit{ncol} \ne \mathit{col}] \Big),$$
where $\mathit{ns}$ is the next sibling of $s$. The term $\mathbf{1}[\mathit{ncol} \ne \mathit{col}]$ counts the edge $v$–$s$ when it connects different colors.

**Complexity.** A naive bound is $O(n^3)$. One can show the DP actually does $O(n^2)$ updates: each pair of vertices $(x, y)$ contributes at most once when $v = \operatorname{lca}(x, y)$, so the total number of updates is $O(n^2)$.