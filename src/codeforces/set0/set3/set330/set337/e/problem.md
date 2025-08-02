# Divisor Tree

## Problem Description

A **divisor tree** is a rooted tree that meets the following conditions:

- Each vertex of the tree contains a positive integer number.
- The numbers written in the leaves of the tree are prime numbers.
- For any inner vertex, the number within it is equal to the product of the numbers written in its children.

Manao has $n$ distinct integers $a_1, a_2, \ldots, a_n$. He tries to build a divisor tree which contains each of these numbers. That is, for each $a_i$, there should be at least one vertex in the tree which contains $a_i$. 

Manao loves compact style, but his trees are too large. Help Manao determine the **minimum possible number of vertices** in the divisor tree sought.

## Input

- The first line contains a single integer $n$ ($1 \leq n \leq 8$).
- The second line contains $n$ distinct space-separated integers $a_i$ ($2 \leq a_i \leq 10^{12}$).

## Output

Print a single integer — the minimum number of vertices in the divisor tree that contains each of the numbers $a_i$.

## Examples

### Example 1
**Input:**
```
2
6 10
```

**Output:**
```
7
```

### Example 2
**Input:**
```
4
6 72 8 4
```

**Output:**
```
12
```

### Example 3
**Input:**
```
1
7
```

**Output:**
```
1
```

## Note

**Sample 1:** The smallest divisor tree looks this way:

### ideas
1. a[i]中的质数，可以忽略掉（当作一个叶子节点即可）
2. 按照从小到大排列
3. 只考虑一个a[i], 对a[i]进行质因数分解，那么节点数 = 质因数个数的sum
4. 比如 72 = 2 ** 3 * 3 ** 2 => 5 (然后加上自己， 共6个节点)
5. 如果 a[i] % a[j] = 0, (可以整除的情况下)
6. a[j]单独产生一个节点（a[j]不是质数的情况下）
7. dp[state]表示由state组成的数字的树的最小表示
8. dp[next] = 加入一个新的数字, x, 如果 x可以整除 state表示的数， x / num(state)剩余的数字进行分解
9. 如果x不能整除，那么它们就需要组合出一个新的节点出来