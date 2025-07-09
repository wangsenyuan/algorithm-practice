# Problem F: Subsegments with Given Sum and Maximum

## Problem Description

You are given an array of integers $a_1, a_2, \ldots, a_n$ and two integers $s$ and $x$. Count the number of subsegments of the array whose sum of elements equals $s$ and whose maximum value equals $x$.

More formally, count the number of pairs $1 \leq l \leq r \leq n$ such that:

- $a_l + a_{l+1} + \ldots + a_r = s$
- $\max(a_l, a_{l+1}, \ldots, a_r) = x$

## Input

Each test consists of multiple test cases. The first line contains a single integer $t$ $(1 \leq t \leq 10^4)$ — the number of test cases. The description of the test cases follows.

The first line of each test case contains three integers $n$, $s$, and $x$:
- $1 \leq n \leq 2 \cdot 10^5$
- $-2 \cdot 10^{14} \leq s \leq 2 \cdot 10^{14}$
- $-10^9 \leq x \leq 10^9$

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ $(-10^9 \leq a_i \leq 10^9)$.

It is guaranteed that the sum of $n$ across all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output the number of subsegments of the array whose sum of elements equals $s$ and whose maximum value equals $x$.

## Example

### Input
```
9
1 0 0
0
1 -2 -1
-2
3 -1 -1
-1 1 -1
6 -3 -2
-1 -1 -1 -2 -1 -1
8 3 2
2 2 -1 -2 3 -1 2 2
9 6 3
1 2 3 1 2 3 1 2 3
13 7 3
0 -1 3 3 3 -2 1 2 2 3 -1 0 3
2 -2 -1
-2 -1
2 -2 -1
-1 -2
```

### Output
```
1
0
2
0
2
7
8
0
0
```

## Explanation

- **Test case 1**: The suitable subsegment is $l=1$, $r=1$.

- **Test case 3**: The suitable subsegments are $l=1$, $r=1$ and $l=3$, $r=3$.

- **Test case 5**: The suitable subsegments are $l=1$, $r=3$ and $l=6$, $r=8$.

- **Test case 6**: The suitable subsegments are those for which $r=l+2$.

- **Test case 7**: The following subsegments are suitable:
  - $l=1$, $r=7$
  - $l=2$, $r=7$
  - $l=3$, $r=6$
  - $l=4$, $r=8$
  - $l=7$, $r=11$
  - $l=7$, $r=12$
  - $l=8$, $r=10$
  - $l=9$, $r=13$