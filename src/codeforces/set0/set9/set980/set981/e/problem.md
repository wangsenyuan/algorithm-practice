# Problem E: Grisha and the Array

Grisha came to a contest and faced the following problem.

## Problem Description

You are given an array of size $n$, initially consisting of zeros. The elements of the array are enumerated from $1$ to $n$. You perform $q$ operations on the array. The $i$-th operation is described with three integers $l_i$, $r_i$ and $x_i$ ($1 \leq l_i \leq r_i \leq n$, $1 \leq x_i \leq n$) and means that you should add $x_i$ to each of the elements with indices $l_i, l_i+1, \ldots, r_i$. After all operations you should find the maximum in the array.

Grisha is clever, so he solved the problem quickly.

However, something went wrong inside his head and now he thinks of the following question: "Consider we applied some subset of the operations to the array. What are the possible values of the maximum in the array?"

Help Grisha find all integers $y$ between $1$ and $n$ such that if you apply some subset (possibly empty) of the operations, then the maximum in the array becomes equal to $y$.

## Input

The first line contains two integers $n$ and $q$ ($1 \leq n, q \leq 10^4$) — the length of the array and the number of queries in the initial problem.

The following $q$ lines contain queries, one per line. The $i$-th of these lines contains three integers $l_i$, $r_i$ and $x_i$ ($1 \leq l_i \leq r_i \leq n$, $1 \leq x_i \leq n$), denoting a query of adding $x_i$ to the segment from $l_i$-th to $r_i$-th elements of the array, inclusive.

## Output

In the first line print the only integer $k$, denoting the number of integers from $1$ to $n$, inclusive, that can be equal to the maximum in the array after applying some subset (possibly empty) of the given operations.

In the next line print these $k$ integers from $1$ to $n$ — the possible values of the maximum. Print these integers in increasing order.

## Examples

### Example 1

**Input:**
```
4 3
1 3 1
2 4 2
3 4 4
```

**Output:**
```
4
1 2 3 4
```

**Explanation:**
- If you consider the subset only of the first query, the maximum is equal to $1$
- If you take only the second query, the maximum equals to $2$
- If you take the first two queries, the maximum becomes $3$
- If you take only the third query, the maximum becomes $4$
- If you take the third query and something more, the maximum becomes greater than $n$, so you shouldn't print it

### Example 2

**Input:**
```
7 2
1 5 1
3 7 2
```

**Output:**
```
3
1 2 3
```

**Explanation:**
- You can take the first query to obtain $1$
- You can take only the second query to obtain $2$
- You can take all queries to obtain $3$

### Example 3

**Input:**
```
10 3
1 1 2
1 1 3
1 1 6
```

**Output:**
```
6
2 3 5 6 8 9
```

**Explanation:**
- You can achieve the maximum of $2$ by using queries: $(1)$
- You can achieve the maximum of $3$ by using queries: $(2)$
- You can achieve the maximum of $5$ by using queries: $(1,2)$
- You can achieve the maximum of $6$ by using queries: $(3)$
- You can achieve the maximum of $8$ by using queries: $(1,3)$
- You can achieve the maximum of $9$ by using queries: $(2,3)$

### ideas
1. 也就是说，采用一部分区间，是否能够得到 max(a) = y 的结果
2. 这里有一个点，就是对于某个位置i, 它假设可取值的最大值时1, 2, 3.... x
3. 那么这个取值可以作为y来使用（就是整个区间的最大值）那些不包含i的区间都不保留
4. n * q / 64?
5. 假设现在处理位置i, 保留那些覆盖i的区间，然后dp