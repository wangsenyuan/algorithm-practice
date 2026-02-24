# Problem B

There are $n$ piles of stones of sizes $a_1, a_2, \ldots, a_n$ on the table.

In one move you may take one pile and add it to another: adding pile $i$ to pile $j$ increases the size of pile $j$ by the current size of pile $i$, and pile $i$ is removed. The **cost** of such a move is the size of the pile that was added.

Your task is to find the minimum total cost to gather all stones into one pile.

**Constraint:** Each pile may be used as the **target** of an add at most $k$ times (it can still be added to others). So each pile can "receive" at most $k$ other piles.

You are given $q$ (possibly repeated) values of $k$. For each value, find the minimum cost to merge all piles under that constraint.

---

## Input

- First line: integer $n$ ($1 \le n \le 10^5$) — the number of piles.
- Second line: $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le 10^9$) — the initial pile sizes.
- Third line: integer $q$ ($1 \le q \le 10^5$) — the number of queries.
- Fourth line: $q$ integers $k_1, k_2, \ldots, k_q$ ($1 \le k_i \le 10^5$) — the values of $k$ for each query (values may repeat).

## Output

Print $q$ space-separated integers — the minimum cost for each query in input order.

---

## Example

**Input:**

```
5
2 3 4 1 1
2
2 3
```

**Output:**

```
9 8
```

**Note:**

- **$k = 2$:** One optimal sequence: add pile 4 to pile 2 (cost 1), add pile 5 to pile 2 (cost 1), add pile 1 to pile 3 (cost 2), add pile 2 to pile 3 (cost 5; pile 2 had become 2+1+1=4 then 5 after the first two adds). Total cost $1+1+2+5 = 9$.
- **$k = 3$:** Add pile 2 to pile 3 (cost 3), pile 1 to pile 3 (cost 2), pile 5 to pile 4 (cost 1), pile 4 to pile 3 (cost 2). Total cost $3+2+1+2 = 8$.


### ideas
1. 先不考虑k的影响， 有三堆a < b
2. 和霍夫曼编码还不一样，霍夫曼编码是增加两个节点的sum，而这里是增加其中一边的值
3. 是一棵（每个节点）不超过k个（子节点）的树
4. 每个节点的贡献 = d * a[i] 
5. 所以树的高度越小越好
6. 第一层1个节点，第二层k个节点，第三层 k * k 个节点
7. 知道了