# ARC201 B - Knapsack Problem

## Problem Description

There are $N$ items numbered $1, 2, \ldots, N$. Item $i$ has weight $2^{X_i}$ and value $Y_i$.

When choosing $0$ or more items such that the sum of weights is at most $W$, find the maximum possible sum of values.

$T$ test cases are given, so find the answer for each.

## Constraints

- $1 \leq T \leq 10^5$
- $1 \leq N \leq 2 \times 10^5$
- $1 \leq W \leq 10^{18}$
- $0 \leq X_i < 60$
- $1 \leq Y_i \leq 10^9$
- The sum of $N$ over all test cases is at most $2 \times 10^5$
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```
T
case_1
case_2
⋮
case_T
```

Each test case is given in the following format:

```
N W
X_1 Y_1
X_2 Y_2
⋮
X_N Y_N
```

## Output

Output $T$ lines.

The $i$-th line should contain the maximum possible sum of values for the $i$-th test case.

## Sample Input 1

```
3
4 16
0 3
3 2
4 5
3 4
1 576460752303423487
59 1000000000
15 23752394551518
42 457687868
42 527769950
41 336200204
42 555090674
32 452384
42 697175056
38 62269946
39 24293218
40 214670437
37 6306990
39 103832403
41 205334902
37 20326312
35 4723060
42 176783630
```

## Sample Output 1

```
7
0
3102936938
```

## Explanation

**For the first test case:**

The pairs of weight and value for items $1, 2, 3, 4$ are $(1, 3), (8, 2), (16, 5), (8, 4)$, respectively.

If we choose items $1$ and $4$, the sum of weights is $1 + 8 = 9 \leq 16$ and the sum of values is $3 + 4 = 7$.

**For the second test case:**

Note that it is allowed to choose $0$ items.

### ideas
1. 体积越大，value越大才可以
2. 所以可以按照体积,value递增排序