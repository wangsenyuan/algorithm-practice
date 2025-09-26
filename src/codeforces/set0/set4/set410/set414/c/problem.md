# Problem Description

Mashmokh's boss, Bimokh, didn't like Mashmokh. So he fired him. Mashmokh decided to go to university and participate in ACM instead of finding a new job. He wants to become a member of Bamokh's team. In order to join he was given some programming tasks and one week to solve them. Mashmokh is not a very experienced programmer. Actually he is not a programmer at all. So he wasn't able to solve them. That's why he asked you to help him with these tasks. One of these tasks is the following.

You have an array `a` of length $2^n$ and $m$ queries on it. The $i$-th query is described by an integer $q_i$. In order to perform the $i$-th query you must:

1. Split the array into $2^{n-q_i}$ parts, where each part is a subarray consisting of $2^{q_i}$ numbers; the $j$-th subarray ($1 \leq j \leq 2^{n-q_i}$) should contain the elements $a[(j-1) \cdot 2^{q_i} + 1], a[(j-1) \cdot 2^{q_i} + 2], \ldots, a[(j-1) \cdot 2^{q_i} + 2^{q_i}]$
2. Reverse each of the subarrays
3. Join them into a single array in the same order (this array becomes new array `a`)
4. Output the number of inversions in the new `a`

Given initial array `a` and all the queries. Answer all the queries. Please, note that the changes from some query is saved for further queries.

## Input

The first line of input contains a single integer $n$ ($0 \leq n \leq 20$).

The second line of input contains $2^n$ space-separated integers $a[1], a[2], \ldots, a[2^n]$ ($1 \leq a[i] \leq 10^9$), the initial array.

The third line of input contains a single integer $m$ ($1 \leq m \leq 10^6$).

The fourth line of input contains $m$ space-separated integers $q_1, q_2, \ldots, q_m$ ($0 \leq q_i \leq n$), the queries.

**Note:** Since the size of the input and output could be very large, don't use slow output techniques in your language. For example, do not use input and output streams (`cin`, `cout`) in C++.

## Output

Output $m$ lines. In the $i$-th line print the answer (the number of inversions) for the $i$-th query.

## Examples

### Example 1

**Input:**
```
2
2 1 4 3
4
1 2 0 2
```

**Output:**
```
0
6
6
0
```

### Example 2

**Input:**
```
1
1 2
3
0 1 1
```

**Output:**
```
0
1
0
```

## Notes

- If we reverse an array $x[1], x[2], \ldots, x[n]$ it becomes new array $y[1], y[2], \ldots, y[n]$, where $y[i] = x[n - i + 1]$ for each $i$.

- The number of inversions of an array $x[1], x[2], \ldots, x[n]$ is the number of pairs of indices $i, j$ such that: $i < j$ and $x[i] > x[j]$.

### ideas
1. 假设在第n层的增加量是x，如果交换的都是下层的，那么这个交换量不会变化
2. 