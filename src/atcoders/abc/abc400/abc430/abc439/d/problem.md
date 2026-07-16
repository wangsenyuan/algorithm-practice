# D - Kadomatsu Subsequence

[Problem link](https://atcoder.jp/contests/abc439/tasks/abc439_d)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 425 points

## Problem Statement

You are given an integer sequence `A = (A_1, A_2, ..., A_N)` of length `N`.
Find the number of triples of integers `(i, j, k)` that satisfy all of the
following:

- `1 <= i, j, k <= N`
- `A_i : A_j : A_k = 7 : 5 : 3`
- `min(i, j, k) = j` or `max(i, j, k) = j`

## Constraints

- All input values are integers
- `1 <= N <= 3 * 10^5`
- `1 <= A_i <= 10^9`

## Input

```
N
A_1 A_2 ... A_N
```

## Output

Output the answer.

## Samples

### Sample 1

Input:

```
10
3 10 7 10 7 6 7 6 5 14
```

Output:

```
7
```

### Sample 2

Input:

```
6
210 210 210 210 210 210
```

Output:

```
0
```

### Sample 3

Input:

```
21
49 30 50 21 35 15 21 70 35 9 50 70 21 49 30 50 70 15 9 21 30
```

Output:

```
34
```

### ideas
1. a[i]:a[j]:a[k] = 7:5:3
2. j是最小的下标, 或者是最大的下标. 
3. 假设j是最小的下标, 那么就是, 假设a[j] = w, 那么就要找到右边 a[i] = 7 * w / 5的全部的i, a[k] = 3 * w / 5的全部的k
4. 