# CNF

"In Boolean logic, a formula is in conjunctive normal form (CNF) or clausal normal form if it is a conjunction of clauses, where a clause is a disjunction of literals" (cited from https://en.wikipedia.org/wiki/Conjunctive_normal_form)

In other words, CNF is a formula of type (clause1) ∧ (clause2) ∧ ... ∧ (clause_n), where ∧ represents logical "AND" (conjunction), each clause is a disjunction (OR) of literals, and literals are boolean variables or their negations.

You are given a CNF containing variables x1, ..., xm and their negations. Each variable occurs in at most two clauses (with negation and without negation in total). Your task is to determine whether this CNF is satisfiable, that is, whether there exist values of variables such that the CNF evaluates to true. If the CNF is satisfiable, you also need to output the values of the variables.

It is guaranteed that each variable occurs at most once in each clause.

## Input

The first line contains integers n and m (1 ≤ n, m ≤ 2·10^5) — the number of clauses and the number of variables.

Next n lines contain the descriptions of each clause. The i-th line first contains ki (ki ≥ 1) — the number of literals in the i-th clause. Then follow space-separated literals vij (1 ≤ |vij| ≤ m). A literal corresponds to x|vij|: with negation if vij is negative, or without negation otherwise.

## Output

If the CNF is not satisfiable, print a single line "NO". Otherwise print two lines: "YES", and then a string of m numbers (0 or 1) — the values of variables x1 to xm in order.

## Examples

### Example 1

**Input:**

```
2 2
2 1 -2
2 2 -1
```

**Output:**

```
YES
11
```

### Example 2

**Input:**

```
4 3
1 1
1 2
3 -1 -2 3
1 -3
```

**Output:**

```
NO
```

### Example 3

**Input:**

```
5 6
2 1 2
3 1 -2 3
4 -3 5 4 6
2 -6 -4
1 5
```

**Output:**

```
YES
100010
```

## Note

In the first sample, the formula is (x1 ∨ ¬x2) ∧ (x2 ∨ ¬x1). One possible answer is x1 = 1, x2 = 1.


### editorial
Firstly let's assign values to variables occurring in our fomula only with negation or only without negation. After that we can throw away the disjuncts which contained them, since they are already true, and continue the process until it is possible. To make it run in time limit, one should use dfs or bfs algorithm to eliminate these variables and disjuncts.

So now we have only variables which have both types of occurrences in disjucnts. Let's build a graph with the vertices corresponding to disjuncts, and for each varible a make an edge between the disjuncts that contain a and !a. Now we should choose the directions of edges in this graph in such a way that every vertex has at least one incoming edge.

We can notice that if some connected component of this graph is a tree, the solution is not possible: on each step we can take some leaf of this tree, and we have to orient its only edge to it, and then erase the leaf. In the end we'll have only one vertex, and it'll not have any incoming edges.

Otherwise, take some cycle in the component and orient the edges between neighbouring vertices along it. Then run dfs from every vertex of the cycle and orient each visited edge in the direction we went along it. It is easy to easy that after this process every vertex will have at least one incoming edge.

So, we should consider cases with the variables which values can be assigned at once, than construct a graph from disjuncts and variables and find whether each connected component has a cycle. If so, we also should carefully construct the answer, assigning the remaining variables their values, looking at the directions of the edges in the graph. The overall complexity is O(n + m).

### solution summary (bipartite matching approach)

Each clause needs at least one literal set to true. Each variable appears in at most 2 clauses.
The key insight: satisfying all clauses is equivalent to finding a **system of distinct representatives** --
each clause picks one variable/literal as its "responsible satisfier".

**Step 1 -- Eliminate "same-sign double" variables**

For each variable v, build `g[v]` = clauses where v appears positively, `g[m+v]` = clauses where v appears negatively.

- If `|g[v]| == 2` (v appears positive in 2 clauses): set v=1, both clauses are satisfied. Mark them as **eliminated**.
- If `|g[m+v]| == 2` (v appears negative in 2 clauses): set v=0, both clauses are satisfied. Mark them eliminated.

These variables satisfy all their clauses with a single assignment -- no conflict possible.

**Step 2 -- Build bipartite graph for remaining clauses**

- Left side: variables (1..m)
- Right side: non-eliminated clauses
- Edge from variable v to clause c if v appears (positive or negative) in c and c is not eliminated

**Step 3 -- Find maximum bipartite matching**

Use Hungarian algorithm (augmenting path DFS). Each round: reset visited flags, try all unmatched
variables, find one augmenting path, repeat until no more found.

**Step 4 -- Check feasibility**

If matching size != number of non-eliminated clauses, some clause has no responsible variable -> "NO".

This works because: in a tree component with V vertices and V-1 edges, the maximum matching is V-1,
but we need V clauses satisfied -- impossible. In a component with a cycle (E >= V), matching can
cover all V vertices.

**Step 5 -- Determine variable assignments from matching**

For each matched clause c with `match[c] = variable u`:
- If c is in `g[u]` (positive list) -> set u=1 (`marked[u] = true`)
- If c is in `g[m+u]` (negative list) -> set u=0 (default)

Unmatched variables default to 0.

**Complexity**: O(n * m) worst case due to Hungarian matching (but fast in practice with the
augmenting path heuristic).
