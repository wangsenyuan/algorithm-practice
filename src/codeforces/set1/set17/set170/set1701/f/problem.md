A triple of points $i$, $j$ and $k$ on a coordinate line is called beautiful if $i < j < k$ and $k - i \leq d$.

You are given a set of points on a coordinate line, initially empty. You have to process queries of three types:

- add a point;
- remove a point;
- calculate the number of beautiful triples consisting of points belonging to the set.

### Input

The first line contains two integers $q$ and $d$ ($1 \leq q, d \leq 2 \times 10^5$) — the number of queries and the parameter for defining if a triple is beautiful, respectively.

The second line contains $q$ integers $a_1, a_2, \ldots, a_q$ ($1 \leq a_i \leq 2 \times 10^5$) denoting the queries. The integer $a_i$ denotes the $i$-th query in the following way:

- if the point $a_i$ belongs to the set, remove it; otherwise, add it;
- after adding or removing the point, print the number of beautiful triples.

### Output

For each query, print one integer — the number of beautiful triples after processing the respective query.

### Example

**Input**
```
7 5
8 5 3 2 1 5 6
```

**Output**
```
0
0
1
2
5
1
5
```

