### Caisa and the Tree Queries

Caisa is now at home and his son has a simple task for him.

Given a rooted tree with \(n\) vertices, numbered from \(1\) to \(n\) (vertex \(1\) is the root). Each vertex of the tree has a value. You should answer \(q\) queries. Each query is one of the following:

- **Query type 1**: `1 v`  
  Let's write out the sequence of vertices along the path from the root to vertex \(v\): \(u_1, u_2, \ldots, u_k\) \((u_1 = 1; u_k = v)\).  
  You need to output such a vertex \(u_i\) that \(\gcd(\text{value of } u_i, \text{ value of } v) > 1\) and \(i < k\).  
  If there are several possible vertices \(u_i\), pick the one with maximum value of \(i\).  
  If there is no such vertex, output \(-1\).

- **Query type 2**: `2 v w`  
  You must change the value of vertex \(v\) to \(w\).

You are given all the queries, help Caisa to solve the problem.

---

### Input

The first line contains two space-separated integers \(n, q\) \((1 \le n, q \le 10^5)\).

The second line contains \(n\) integers \(a_1, a_2, \ldots, a_n\) \((1 \le a_i \le 2 \cdot 10^6)\), where \(a_i\) represents the value of node \(i\).

Each of the next \(n - 1\) lines contains two integers \(x_i\) and \(y_i\) \((1 \le x_i, y_i \le n; x_i \ne y_i)\), denoting the edge of the tree between vertices \(x_i\) and \(y_i\).

Each of the next \(q\) lines contains a query in the format that is given above. For each query the following inequalities hold: \(1 \le v \le n\) and \(1 \le w \le 2 \cdot 10^6\). Note that there are no more than \(50\) queries that change the value of a vertex.

---

### Output

For each query of the first type output the result of the query.

---

### Example

**Input**

```text
4 6
10 8 4 3
1 2
2 3
3 4
1 1
1 2
1 3
1 4
2 1 9
1 4
```

**Output**

```text
-1
1
2
-1
1
```

---

### Note

\(\gcd(x, y)\) is the greatest common divisor of two integers \(x\) and \(y\).


### ideas
1. 要找到离u最近的祖先节点，和它的gcd(a[u], a[v]) > 1
2. x最多log(n)的质数因子，假设把所有的质数上构建一棵树，那么感觉应该只有logn * n 的节点
3. 但是问题在于，如何维护这棵树？
4. 还有就是，怎么保证在查询的时候，找到的都是它的祖先节点？
5.  如果没有修改, 可以直接让u节点，根据质因数指向对应的祖先节点
6.  但是修改怎么办呢？
7.  肯定无法去更新所有指向它的子节点了～
8.  如果使用heavy-light分解呢？
9.  关键的信息丢了，就是最多50次修改
10. 那就简单了，每次修改，都重新构建树（就可以了）