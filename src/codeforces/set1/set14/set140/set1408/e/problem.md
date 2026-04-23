# Problem

You are given `m` sets of integers `A1, A2, ..., Am`. Elements of these sets are integers between `1` and `n`, inclusive.

There are two arrays of positive integers `a1, a2, ..., am` and `b1, b2, ..., bn`.

In one operation you can delete an element `j` from the set `Ai` and pay `ai + bj` coins for that.

You can make several (maybe zero) operations (some sets can become empty).

After that, you will build an edge-colored undirected graph on `n` vertices. For each set `Ai` you add an edge `(x, y)` with color `i` for all `x, y ∈ Ai` with `x < y`. Some pairs of vertices can be connected by more than one edge, but such edges have different colors.

You call a cycle

`i1 → e1 → i2 → e2 → ... → ik → ek → i1`

(where `ej` is an edge connecting vertices `ij` and `i(j+1)` in this graph, with `ik+1 = i1`) **rainbow** if all edges on it have different colors.

Find the minimum number of coins you should pay so that the resulting graph has **no** rainbow cycles.

## Input

The first line contains two integers `m` and `n` (`1 ≤ m, n ≤ 10^5`) — the number of sets and the number of vertices.

The second line contains `m` integers `a1, a2, ..., am` (`1 ≤ ai ≤ 10^9`).

The third line contains `n` integers `b1, b2, ..., bn` (`1 ≤ bi ≤ 10^9`).

Each of the next `m` lines describes one set. In the `i`-th line, the first integer `si` (`1 ≤ si ≤ n`) is the size of `Ai`. Then `si` distinct integers follow — the elements of `Ai`, each between `1` and `n`.

It is guaranteed that the sum of `si` over all `1 ≤ i ≤ m` does not exceed `2 * 10^5`.

## Output

Print one integer: the minimum total number of coins paid so that the obtained graph has no rainbow cycles.

## Examples

### Example 1

**Input**

```text
3 2
1 2 3
4 5
2 1 2
2 1 2
2 1 2
```

**Output**

```text
11
```

### Example 2

**Input**

```text
7 8
3 6 7 9 10 7 239
8 1 9 7 10 2 6 239
3 2 1 3
2 4 1
3 1 3 7
2 4 3
5 3 4 5 6 7
2 5 7
1 8
```

**Output**

```text
66
```

## Note

In the first test, you can make such operations:

1. Delete element `1` from set `1`. You pay `a1 + b1 = 5` coins.
2. Delete element `1` from set `2`. You pay `a2 + b1 = 6` coins.

You pay `11` coins in total. After these operations, the first and second sets become `{2}` and the third set becomes `{1, 2}`.

So the graph consists of one edge `(1, 2)` of color `3`.

In the second test, you can make such operations:

1. Delete element `1` from set `1`. Pay `a1 + b1 = 11`.
2. Delete element `4` from set `2`. Pay `a2 + b4 = 13`.
3. Delete element `7` from set `3`. Pay `a3 + b7 = 13`.
4. Delete element `4` from set `4`. Pay `a4 + b4 = 16`.
5. Delete element `7` from set `6`. Pay `a6 + b7 = 13`.

You pay `66` coins in total.

After these operations, the sets will be:

- `{2, 3}`
- `{1}`
- `{1, 3}`
- `{3}`
- `{3, 4, 5, 6, 7}`
- `{5}`
- `{8}`

The resulting graph has no rainbow cycles.


### ideas
1. 看得我头大～
2. 考虑3个点(a, b, c),如果存在3个set, (a, b), (b, c), (a, c)
3. 那么就存在3条线(a, b), (b, c), (a, c) 此时就出现了一个彩色的环
4. 甚至(a, b) 出现在两个set中时，也会出现环
5. 所以，最后的结构，就是用最小的代价（在不产生环）把这些点连接起来
6. 选择集合i,删除数字j，代价 = a[i] + b[j]
7. 那么把所有的都删除掉，是可以算出来的
8. 然后反过来，在不产生环的情况下，把它们尽可能的加回来？
9. 

## key insights

1. Model the problem as a **bipartite graph**:
   - left side: one node for each set `Ai`
   - right side: one node for each value `1..n`
   - if value `j` belongs to set `Ai`, add bipartite edge `(Ai, j)` with weight `ai + bj`

2. Deleting element `j` from `Ai` means **removing** this bipartite edge and paying its weight.
   So if we delete everything, total cost is:

   `sum(ai + bj)` over all memberships `j in Ai`

3. The key equivalence used by the solution:
   - a rainbow cycle in the original graph corresponds to a cycle in this bipartite incidence graph
   - therefore, to avoid rainbow cycles, the kept bipartite edges must form a **forest**

4. So the optimization becomes:
   - start from deleting all edges
   - add back as much total weight as possible
   - but only while keeping the bipartite graph acyclic

5. That is exactly a **maximum spanning forest** problem.
   The code:
   - builds all bipartite edges with weight `ai + bj`
   - sorts them in descending weight
   - uses DSU / Union-Find to keep only edges that connect two different components

6. Let:
   - `tot` = sum of weights of all bipartite edges
   - `saved` = total weight of edges kept in the maximum spanning forest

   Then the answer is:

   `tot - saved`

7. Complexity:
   - let `E = sum |Ai|` (`E <= 2 * 10^5`)
   - sorting: `O(E log E)`
   - DSU operations: `O(E alpha(m+n))`

   This easily fits the limits.