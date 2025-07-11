# Problem E: Clones and Lab Escape

## Problem Description

The evil Bumbershoot corporation produces clones for gruesome experiments in a vast underground lab. On one occasion, the corp cloned a boy Andryusha who was smarter than his comrades. Immediately Andryusha understood that something fishy was going on there. He rallied fellow clones to go on a feud against the evil corp, and they set out to find an exit from the lab. The corp had to reduce to destroy the lab complex.

The lab can be pictured as a connected graph with n vertices and m edges. k clones of Andryusha start looking for an exit in some of the vertices. Each clone can traverse any edge once per second. Any number of clones are allowed to be at any vertex simultaneously. Each clone is allowed to stop looking at any time moment, but he must look at his starting vertex at least. The exit can be located at any vertex of the lab, hence each vertex must be visited by at least one clone.

Each clone can visit at most vertices before the lab explodes.

Your task is to choose starting vertices and searching routes for the clones. Each route can have at most vertices.

## Input

The first line contains three integers n, m, and k (1 ≤ n ≤ 2·10⁵, n - 1 ≤ m ≤ 2·10⁵, 1 ≤ k ≤ n) — the number of vertices and edges in the lab, and the number of clones.

Each of the next m lines contains two integers xi and yi (1 ≤ xi, yi ≤ n) — indices of vertices connected by the respective edge. The graph is allowed to have self-loops and multiple edges.

The graph is guaranteed to be connected.

## Output

You should print k lines. i-th of these lines must start with an integer ci — the number of vertices visited by i-th clone, followed by ci integers — indices of vertices visited by this clone in the order of visiting. You have to print each vertex every time it is visited, regardless if it was visited earlier or not.

It is guaranteed that a valid answer exists.

## Examples

### Example 1

**Input:**
```
3 2 1
2 1
3 1
```

**Output:**
```
3 2 1 3
```

**Explanation:** In this case there is only one clone who may visit vertices in order (2, 1, 3), which fits the constraint of 6 vertices per clone.

### Example 2

**Input:**
```
5 4 2
1 2
1 3
1 4
1 5
```

**Output:**
```
3 2 1 3
3 4 1 5
```

**Explanation:** The two clones can visited vertices in order (2, 1, 3) and (4, 1, 5), which fits the constraint of 5 vertices per clone.


## Ideas

1. **Self-loops and multiple edges**: 这部分是多余的，要给舍弃掉
2. **Clone visit constraint**: Each clone can visit at most $2 \times n \div k$ vertices
3. eular tour。。。。