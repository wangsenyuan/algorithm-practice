# Problem

A cat lives on a tree with $n$ nodes. The cat starts on node $1$, and you live on node $n$. You will leave the cat a note in the **parkour language** so that it can reach you.

## Parkour language

Two instruction types:

1. **`1`** — The cat moves to any adjacent node (if there are several, it picks one arbitrarily). If it has no adjacent nodes, it does not move.
2. **`2 u`** — Destroy node $u$ and all edges incident to it. If the cat is on node $u$, it dies, so this must be avoided. If $u$ is already destroyed, nothing happens.

**Constraint:** Two instructions of type 2 cannot appear consecutively.

The language is ambiguous because the cat can choose among several moves for each `1`. You must construct a sequence of at most $3n$ instructions such that, no matter how the cat chooses, it ends at node $n$. It can be shown that such a sequence exists for any tree.

---

## Input

- First line: number of test cases $t$ ($1 \le t \le 10^4$).
- For each test case:
  - Line 1: integer $n$ ($2 \le n \le 2 \cdot 10^5$) — size of the tree.
  - Next $n - 1$ lines: two integers $u$, $v$ ($1 \le u, v \le n$, $u \ne v$) — an edge between $u$ and $v$.
- The graph is a tree (no loops, no multiple edges).
- Sum of $n$ over all test cases $\le 2 \cdot 10^5$.

## Output

For each test case:

- Line 1: integer $k$ ($0 \le k \le 3n$) — number of instructions.
- Next $k$ lines, each either:
  - `1` — cat moves to an adjacent node (if any).
  - `2 u` ($1 \le u \le n$) — delete node $u$ and all its incident edges.

---

## Example

**Input**

```
4
5
1 2
2 3
1 5
5 4
2
1 2
4
1 2
1 3
1 4
6
1 2
1 3
3 4
4 5
4 6
```

**Output**

```
2
2 2
1

1
1

5
2 2
1
1
2 3
1

9
2 2
1
2 1
1
2 3
1
1
2 5
1
```

*(Extra blank lines between test cases are for clarity only; you do not need to print them.)*

---

## Note

- In the first test case, the only possible path leads the cat to node $5$.
- An invalid sequence for the first test case is `1`, `2 2`: the cat might be on node $2$ when node $2$ is destroyed and die.


### ideas
1. 给节点图色
2. 那么每次操作1，都是改变小猫的节点的颜色，操作2删除离n最远的，且颜色和小猫所在节点不一致的节点（如果最远的点是和小猫颜色一致的）那么随便删一个点