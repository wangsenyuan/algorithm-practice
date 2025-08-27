# Problem Statement

There is a linear strip divided into $m$ cells, numbered from $1$ to $m$ from left to right.

You are given $n$ segments. Each segment is defined by four numbers: $l$, $r$, $p$ and $q$ — the segment covers cells from $l$ to $r$ inclusively and exists with probability $\frac{p}{q}$ (independently).

Your task is to calculate the probability that each cell is covered by exactly one segment.

## Input

The first line contains two integers $n$ and $m$ ($1 \leq n, m \leq 2 \cdot 10^5$).

Then $n$ lines follow. The $i$-th of them contains four integers $l_i$, $r_i$, $p_i$ and $q_i$ ($1 \leq l_i \leq r_i \leq m$; $1 \leq p_i < q_i < 998244353$).

## Output

Print a single integer — the probability that each cell is covered by exactly one segment, taken modulo $998244353$.

Formally, the probability can be expressed as an irreducible fraction $\frac{x}{y}$. You have to print the value of $x \cdot y^{-1} \bmod 998244353$, where $y^{-1}$ is an integer such that $y \cdot y^{-1} \bmod 998244353 = 1$.

## Examples

### Example 1
**Input:**
```
3 3
1 2 1 3
3 3 1 2
1 3 2 3
```

**Output:**
```
610038216
```

### Example 2
**Input:**
```
2 3
1 2 1 2
2 3 1 2
```

**Output:**
```
0
```

### Example 3
**Input:**
```
8 5
1 3 1 2
1 5 1 6
1 4 4 5
5 5 1 7
4 5 1 2
4 5 2 5
3 3 2 7
1 2 1 3
```

**Output:**
```
94391813
```

## Note

In the first example, the probability is equal to $\frac{5}{18}$.