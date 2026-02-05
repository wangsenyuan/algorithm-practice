Rae, an Earth mage, has mastered the spell of growing trees. Manaria, however, claims that she can grow a more impressive species of trees. Rae remembers that the rarest type of tree can be grown using a formula represented by a certain permutation — please help her construct it.

You are given a permutation\* $p$ of length $n$.

Determine if there exists an undirected tree with $n$ vertices labeled $1, 2, \ldots, n$, satisfying the following condition:

- Let $u$ and $v$ ($1 \le u < v \le n$) be any two vertices connected by an edge. Then $u$ appears before $v$ in $p$.

Additionally, if there exists such a tree, output any of them.

\*A permutation of length $n$ is an array that contains every integer from $1$ to $n$ exactly once, in any order.

## Input

The first line contains a single integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

The first line of each test case contains a single integer $n$ ($2 \le n \le 2 \cdot 10^5$).

The second line of each test case contains $n$ integers $p_1, p_2, \ldots, p_n$ ($1 \le p_i \le n$). It is guaranteed that all $p_i$ are distinct.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output on a single line `Yes` if there exists a tree satisfying the given condition, and `No` otherwise.

Then, if the answer is `Yes`, output $n - 1$ lines. The $i$-th of these lines should contain two integers $u$ and $v$, denoting an edge connecting vertices $u$ and $v$.

You may output the answer in any case (upper or lower). For example, the strings `yEs`, `yes`, `YES`, and `yeS` will be recognized as `Yes`.

## Example

### Input

```
9
6
1 3 4 5 2 6
4
3 4 1 2
5
4 3 5 1 2
4
1 2 3 4
7
4 3 5 7 6 2 1
6
2 4 6 1 3 5
3
2 1 3
4
2 4 1 3
6
4 2 6 5 1 3
```

### Output

```
Yes
3 1
4 1
6 5
6 2
6 1
No
No
Yes
2 1
4 3
4 1
No
Yes
4 2
6 2
3 1
5 1
5 2
Yes
3 2
3 1
Yes
4 2
3 1
3 2
Yes
6 4
6 2
3 1
5 4
2 3
```

## Note

In the first example, we can construct the tree given in the sample output. We have that:

- $1 < 3$, and $1$ appears before $3$ in $p$,
- $1 < 4$, and $1$ appears before $4$ in $p$,
- $5 < 6$, and $5$ appears before $6$ in $p$,
- $2 < 6$, and $2$ appears before $6$ in $p$,
- $1 < 6$, and $1$ appears before $6$ in $p$.

In the second example, it can be shown that there does not exist a tree satisfying the given constraints.


### ideas
1. 构造一棵树，如果(u, v) 有连接，且 (u < v) 那么 pos[u] < pos[v]
2. 考虑前面构造好了，现在出现了一个新的u, 如果前面存在一个数w < u, 那么把它们接起来
3. 如果不存在 w > u, （不能连接，否则就违反规则了）那么后面必须存在一个，还没有被连接的v, w < v, u < v
4. 把它们连接起来
5. 所以，出现一个v的时候，就把所有小于它的连接起来，最后要只有一个component就可以了