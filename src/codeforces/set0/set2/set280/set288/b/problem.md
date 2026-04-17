# Problem

Little penguin Polo loves his home village. The village has $n$ houses, indexed from $1$ to $n$.
Each house has a plaque with an integer; the $i$-th house contains integer $p_i$ ($1 \le p_i \le n$).

Polo walks as follows:

- he starts at house $x$,
- then goes to house $p_x$,
- then to house $p_{p_x}$,
- and so on.

You know that:

1. Starting from any house with index from $1$ to $k$ (inclusive), he can eventually reach house $1$.
2. Starting from any house with index from $k+1$ to $n$ (inclusive), he can never reach house $1$.
3. Starting from house $1$, he can return to house $1$ after a non-zero number of moves.

Find the number of ways to assign plaque numbers so that all three conditions hold.
Print the answer modulo $1000000007$ ($10^9 + 7$).

## Input

A single line containing two integers $n$ and $k$:

- $1 \le n \le 1000$
- $1 \le k \le \min(8, n)$

## Output

Print one integer — the required count modulo $1000000007$.

## Examples

### Example 1

**Input**

```text
5 2
```

**Output**

```text
54
```

### Example 2

**Input**

```text
7 4
```

**Output**

```text
1728
```

## Explanation Notes

Think of `p` as a directed graph where each node has out-degree exactly `1`.

- Condition 1 says every node in `[1..k]` must eventually flow to node `1`.
- Condition 3 says node `1` must lie on a directed cycle (cycle length at least `1`).
- Therefore, inside nodes `[1..k]`, there is exactly one cycle, it contains node `1`, and all other nodes in `[1..k]` feed into that cycle.
- Condition 2 says nodes `[k+1..n]` cannot reach `1`, so they cannot point to any node in `[1..k]` (because all `[1..k]` nodes reach `1`).

So the count factorizes:

1. number of valid mappings on nodes `[1..k]`;
2. number of mappings on nodes `[k+1..n]` that stay inside `[k+1..n]`.

The second factor is immediate:

\[
(n-k)^{(n-k)}
\]

because each of the `n-k` nodes freely chooses one of `n-k` destinations.

### Counting the `[1..k]` part by cycle length

Let cycle length containing node `1` be `s` (`1 <= s <= k`).

1. Choose the other `s-1` cycle nodes:
\[
\binom{k-1}{s-1}
\]
2. Arrange these `s` chosen nodes into one directed cycle that contains `1`:
\[
(s-1)!
\]
3. Attach the remaining `m = k-s` nodes as a directed forest feeding into the cycle nodes.
   - roots are exactly those `s` cycle nodes.
   - count is:
   \[
   \begin{cases}
   1, & m=0 \\
   s \cdot k^{m-1}, & m>0
   \end{cases}
   \]

So contribution of fixed `s` is:

\[
\binom{k-1}{s-1}(s-1)! \times
\begin{cases}
1, & s=k\\
s\cdot k^{k-s-1}, & s<k
\end{cases}
\]

Summing over `s=1..k`, then multiplying by `(n-k)^{n-k}`, gives the final answer.

### Sample 1 (`n=5, k=2`)

For nodes `{1,2}`, valid internal mappings:

- `1->1, 2->1`
- `1->2, 2->1`

So there are `2` ways.

Nodes `{3,4,5}` must stay inside `{3,4,5}`:

\[
3^3 = 27
\]

Total:

\[
2 \cdot 27 = 54
\]

### Sample 2 (`n=7, k=4`)

For nodes `{5,6,7}`:

\[
3^3 = 27
\]

For nodes `{1,2,3,4}`, valid internal count is:

- cycle length `1`: `16`
- cycle length `2`: `24`
- cycle length `3`: `18`
- cycle length `4`: `6`

sum:

\[
16+24+18+6=64
\]

Total:

\[
64 \cdot 27 = 1728
\]

## Cayley Formula (Tree) and Forest Formula

### Cayley formula (labeled tree)

The number of trees on labeled vertices `{1,2,...,n}` is:

\[
n^{n-2}
\]

#### Proof sketch via Prüfer code

- Every labeled tree on `n` vertices maps to a Prüfer sequence of length `n-2`.
- Construction: repeatedly remove the smallest leaf, append its neighbor label.
- This mapping is bijective: every sequence of length `n-2` over `[1..n]` reconstructs a unique tree.
- Number of sequences is `n^{n-2}`, so number of trees is `n^{n-2}`.

### Rooted forest formula (with fixed roots)

Let roots be a fixed set `R` of size `r` among `n` labeled vertices.
Count forests where each component is rooted at one root in `R` and every edge is directed toward the root.

The count is:

\[
r \cdot n^{n-r-1}
\]

for `n>r`, and equals `1` when `n=r`.

#### Proof from Cayley (super-root trick)

- Add a new vertex `0` and connect `0` to every root in `R`.
- For any rooted forest on original `n` vertices with root set `R`, adding these `r` edges yields a tree on `n+1` vertices.
- Conversely, in a tree on `{0,1,...,n}`, if degree of `0` is exactly `r` and its neighbors are exactly `R`, removing `0` recovers exactly one such rooted forest.
- In Prüfer sequence of length `n-1` for trees on `n+1` nodes:
  - vertex `0` appears exactly `deg(0)-1 = r-1` times,
  - the other `n-r` positions can be any labels in `[1..n]`.
- Equivalent counting gives \(r \cdot n^{n-r-1}\).

This is exactly the formula used above with:

- total nodes `n = k`,
- roots `r = s` (the cycle nodes),
- non-root nodes `k-s`.