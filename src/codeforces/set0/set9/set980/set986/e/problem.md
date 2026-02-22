# Problem E

Let the main characters of this problem be personages from some recent movie. New Avengers seem to make a lot of buzz. I didn't watch any part of the franchise and don't know its heroes well, but it won't stop me from using them in this problem statement. So, Thanos and Dr. Strange are doing their superhero and supervillain stuff, but then suddenly they stumble across a regular competitive programming problem.

You are given a tree with $n$ vertices.

In each vertex $v$ there is a positive integer $a_v$.

You have to answer $q$ queries.

Each query has the form $u$ $v$ $x$.

You have to calculate $\prod_{w \in P} \gcd(x, a_w) \bmod (10^9+7)$, where $P$ is the set of vertices on the path from $u$ to $v$. In other words, you are to calculate the product of $\gcd(x, a_w)$ for all vertices $w$ on the path from $u$ to $v$. As it might be large, compute it modulo $10^9+7$. Here $\gcd(s,t)$ denotes the greatest common divisor of $s$ and $t$.

Note that the numbers in vertices do not change after queries.

I suppose that you are more interested in superhero business of Thanos and Dr. Strange than in them solving the problem. So you are invited to solve this problem instead of them.

---

## Input

- First line: one integer $n$ ($1 \le n \le 10^5$) — the size of the tree.
- Next $n-1$ lines: the edges of the tree. The $i$-th edge is described with two integers $u_i$ and $v_i$ ($1 \le u_i, v_i \le n$) and connects vertices $u_i$ and $v_i$. It is guaranteed that the graph is a tree.
- Next line: $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_v \le 10^7$).
- Next line: one integer $q$ ($1 \le q \le 10^5$) — the number of queries.
- Next $q$ lines: each query is described with three integers $u_i$, $v_i$, and $x_i$ ($1 \le u_i, v_i \le n$, $1 \le x_i \le 10^7$).

## Output

Print $q$ numbers — the answers to the queries in the order given. Print each answer modulo $10^9+7 = 1000000007$, one per line.

---

## Examples

### Example 1

**Input:**

```
4
1 2
1 3
1 4
6 4 9 5
3
2 3 6
2 3 2
3 4 7
```

**Output:**

```
36
4
1
```

### Example 2

**Input:**

```
6
1 2
2 3
2 4
1 5
5 6
100000 200000 500000 40000 800000 250000
3
3 5 10000000
6 2 3500000
4 1 64000
```

**Output:**

```
196000
12250
999998215
```
