# C - Sushi (ABC460)

**Contest:** [ABC460](https://atcoder.jp/contests/abc460) — AtCoder Beginner Contest 460  
**Task:** [https://atcoder.jp/contests/abc460/tasks/abc460_c](https://atcoder.jp/contests/abc460/tasks/abc460_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 300 points

## Problem Statement

There are `N` pieces of shari (vinegared rice) and `M` pieces of neta (toppings)
as ingredients for sushi.

The weight of the `i`-th shari is `A_i`, and the weight of the `j`-th neta is
`B_j`.

You will make sushi by combining shari and neta.

To make one piece of sushi, you need to combine one shari with one neta. The
weight of the neta must be at most twice the weight of the shari. Also, the
same shari or neta cannot be used in multiple pieces of sushi.

Find the maximum number of sushi that can be made.

## Constraints

- `1 <= N, M <= 2 * 10^5`
- `1 <= A_i, B_j <= 10^9`
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```text
N M
A_1 A_2 … A_N
B_1 B_2 … B_M
```

## Output

Output the answer.

## Sample Input 1

```text
4 5
4 2 1 8
14 9 3 2 9
```

## Sample Output 1

```text
3
```

By combining the 1st shari with the 3rd neta, the 2nd shari with the 4th neta,
and the 4th shari with the 1st neta, for example, you can make three pieces of
sushi. It is impossible to make four or more pieces of sushi, so output `3`.

## Sample Input 2

```text
3 3
5 5 3
11 1000 1000
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
8 7
2 3 4 4 4 3 2 3
8 5 5 9 9 7 1
```

## Sample Output 3

```text
5
```
