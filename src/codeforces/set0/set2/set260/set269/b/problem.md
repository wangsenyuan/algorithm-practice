# Problem B: Greenhouse Arrangement

## Problem Description

Emuskald is an avid horticulturist and owns the world's longest greenhouse — it is effectively infinite in length.

Over the years Emuskald has cultivated $n$ plants in his greenhouse, of $m$ different plant species numbered from 1 to $m$. His greenhouse is very narrow and can be viewed as an infinite line, with each plant occupying a single point on that line.

Emuskald has discovered that each species thrives at a different temperature, so he wants to arrange $m - 1$ borders that would divide the greenhouse into $m$ sections numbered from 1 to $m$ from left to right with each section housing a single species. He is free to place the borders, but in the end all of the $i$-th species plants must reside in $i$-th section from the left.

Of course, it is not always possible to place the borders in such way, so Emuskald needs to replant some of his plants. He can remove each plant from its position and place it anywhere in the greenhouse (at any real coordinate) with no plant already in it. Since replanting is a lot of stress for the plants, help Emuskald find the minimum number of plants he has to replant to be able to place the borders.

## Input

The first line of input contains two space-separated integers $n$ and $m$ ($1 \leq n, m \leq 5000$, $n \geq m$), the number of plants and the number of different species.

Each of the following $n$ lines contain two space-separated numbers:
- One integer number $s_i$ ($1 \leq s_i \leq m$), the species of the $i$-th plant
- One real number $x_i$ ($0 \leq x_i \leq 10^9$), the position of the $i$-th plant

Each $x_i$ will contain no more than 6 digits after the decimal point.

**Constraints:**
- All $x_i$ are different
- There is at least one plant of each species
- The plants are given in order "from left to the right", that is in the ascending order of their $x_i$ coordinates ($x_i < x_{i+1}$, $1 \leq i < n$)

## Output

Output a single integer — the minimum number of plants to be replanted.

## Examples

### Example 1

**Input:**
```
3 2
2 1
1 2.0
1 3.100
```

**Output:**
```
1
```

**Explanation:** Emuskald can replant the first plant to the right of the last plant, so the answer is 1.

### Example 2

**Input:**
```
3 3
1 5.0
2 5.5
3 6.0
```

**Output:**
```
0
```

**Explanation:** The species are already in the correct order, so no replanting is needed.

### Example 3

**Input:**
```
6 3
1 14.284235
2 17.921382
1 20.328172
3 20.842331
1 25.790145
1 27.204125
```

**Output:**
```
2
```


### ideas
1. 第i个品种要放在第i段房子中
2. 一开始的位置x[i], 最后的位置是y[i]， 这个是确定的
3. dp[i][j], 表示前i个放置到正确的位置后，且最后一个的位置是x[j]时的最优解
4. 现在处理i+1, 要把它摆到最后面去
5. 这里有两种处理方式，一种是吧i+1,移动到最后面去(这时候 + 1)
6. 还有一个办法是把i+1后面的植物，全部移动到i+1的前面去
7. 这里怎么都感觉是第一个方式更优（最多操作1次）但是整体考虑时，这个方案是不对的
8. 问题出在哪里？
9. 我们考虑最后一段，假设有w个x[i] = m 的
10. 那么最后一段开始的位置，有m+1个，分别是第1个，第2个，。。。 第w个，所有位置的最后吗
11. dp[m][0,1....w] = 这个是能算出来的
12. dp[m][w] = w
13. dp[m][w-1] = w - 1 + pos[w]后面的植物的个数 
14. dp[m][w - i] = w  - i + pow[w-i]后面的植物的个数（这些它后面的植物，其实可以直接忽略掉了，它只需要搬一次）
15. dp[m-1][x] = dp[m][y] + x前面第m-1种类 + x...y中间不是 < m - 1种类的植物
16. 似乎可以搞了