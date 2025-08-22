# Problem C: Array Operations

## Problem Description

You are given an array $a_1, a_2, \ldots, a_n$ and $m$ sets $S_1, S_2, \ldots, S_m$ of indices of elements of this array. Let's denote $S_k = \{S_{k,i}\}$ where $1 \leq i \leq |S_k|$. In other words, $S_{k,i}$ is some element from set $S_k$.

In this problem you have to answer $q$ queries of two types:

1. **Find the sum** of elements with indices from set $S_k$: $\sum_{i \in S_k} a_i$
   - Query format: `? k`

2. **Add number $x$** to all elements at indices from set $S_k$: $a_{S_{k,i}}$ is replaced by $a_{S_{k,i}} + x$ for all $i$ $(1 \leq i \leq |S_k|)$
   - Query format: `+ k x$

After each first type query, print the required sum.

## Input

- **First line**: Three integers $n, m, q$ $(1 \leq n, m, q \leq 10^5)$
- **Second line**: $n$ integers $a_1, a_2, \ldots, a_n$ $(|a_i| \leq 10^8)$ — elements of array $a$
- **Next $m$ lines**: Each line describes one set of indices:
  - The $k$-th line first contains a positive integer representing the number of elements in set $|S_k|$
  - Followed by $|S_k|$ distinct integers $S_{k,1}, S_{k,2}, \ldots, S_{k,|S_k|}$ $(1 \leq S_{k,i} \leq n)$ — elements of set $S_k$
- **Next $q$ lines**: Queries in the format `? k` or `+ k x`
  - For all queries: $1 \leq k \leq m$, $|x| \leq 10^8$
  - Queries are given in the order they need to be answered

**Constraints**: The sum of sizes of all sets $S_k$ doesn't exceed $10^5$.

## Output

After each first type query, print the required sum on a single line.

## Notes

- Please, do not write the `%lld` specifier to read or write 64-bit integers in C++
- It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier

## Examples

### Example 1

**Input:**
```
5 3 5
5 -5 5 1 -4
2 1 2
4 2 1 4 5
2 2 5
? 2
+ 3 4
? 1
+ 2 1
? 2
```

**Output:**
```
-3
4
9
```

**Explanation:**
- Initial array: $[5, -5, 5, 1, -4]$
- Set $S_1 = \{1, 2\}$, Set $S_2 = \{2, 1, 4, 5\}$, Set $S_3 = \{2, 5\}$
- Query `? 2`: Sum of elements at indices $\{2, 1, 4, 5\}$ = $(-5) + 5 + 1 + (-4) = -3$
- Query `+ 3 4`: Add 4 to elements at indices $\{2, 5\}$, array becomes $[5, -1, 5, 1, 0]$
- Query `? 1`: Sum of elements at indices $\{1, 2\}$ = $5 + (-1) = 4$
- Query `+ 2 1`: Add 1 to elements at indices $\{2, 1, 4, 5\}$, array becomes $[6, 0, 5, 2, 1]$
- Query `? 2`: Sum of elements at indices $\{2, 1, 4, 5\}$ = $0 + 6 + 2 + 1 = 9$


### ideas
1. 