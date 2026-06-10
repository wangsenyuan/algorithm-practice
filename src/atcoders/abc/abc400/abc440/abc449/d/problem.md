# D - Make Target 2

**Contest:** [ABC449](https://atcoder.jp/contests/abc449) — AtCoder Beginner Contest 449  
**Task:** [https://atcoder.jp/contests/abc449/tasks/abc449_d](https://atcoder.jp/contests/abc449/tasks/abc449_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 425 points

## Problem Statement

There is a two-dimensional coordinate plane. A lattice point with coordinates
`(x, y)` is painted black if:

```text
max(|x|, |y|)
```

is even, and white if it is odd.

Among all pairs of integers `(x, y)` satisfying:

```text
L <= x <= R
D <= y <= U
```

find the number of points painted black.

## Constraints

- `-10^6 <= L <= R <= 10^6`
- `-10^6 <= D <= U <= 10^6`
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
L R D U
```

## Output

Output the answer.

## Sample Input 1

```text
-4 3 1 3
```

## Sample Output 1

```text
10
```

As shown in the figure in the original statement, the answer is `10`.

## Sample Input 2

```text
-14 14 -14 14
```

## Sample Output 2

```text
449
```


### ideas
1. 