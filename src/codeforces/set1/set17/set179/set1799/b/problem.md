# B. Equalize by Divide

[Problem link](https://codeforces.com/problemset/problem/1799/B)

**Contest:** [Codeforces Round 854 by cybercats (Div. 1 + Div. 2)](https://codeforces.com/contest/1799)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a positive integer array `a1, a2, ..., an`. You may perform the following operation any number of times, including zero:

- Choose indices `i` and `j` with `i ≠ j`.
- Replace `a_i` with `ceil(a_i / a_j)`.

If you can make every element equal, print any sequence of at most `30n` operations. Otherwise print `-1`.

If a solution exists, some sequence of at most `30n` operations always exists.

## Constraints

- `1 ≤ t ≤ 1000`
- `1 ≤ n ≤ 100`
- `1 ≤ a_i ≤ 10^9`
- The sum of `n` over all test cases is at most `1000`

## Input

```text
t
n
a_1 a_2 ... a_n
```

Each test case is the pair of lines `n` and the array.

## Output

For each test case, print `q`. If it is impossible, `q = -1`. Otherwise `0 ≤ q ≤ 30n`, and the next `q` lines are the operations `i j`.

Any valid sequence is accepted.

## Sample Input 1

```text
1
100
```

## Sample Output 1

```text
0
```

## Sample Input 2

```text
3
1 1 1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
2
2 1
```

## Sample Output 3

```text
-1
```

## Sample Input 4

```text
2
5 5
```

## Sample Output 4

```text
0
```

## Sample Input 5

```text
3
4 3 2
```

## Sample Output 5

```text
2
1 3
2 1
```

One possible sequence is `[4, 3, 2] → [2, 3, 2] → [2, 2, 2]`.

## Sample Input 6

```text
4
3 3 4 4
```

## Sample Output 6

```text
4
3 1
4 2
1 3
2 4
```

## Sample Input 7

```text
2
2 100
```

## Sample Output 7

```text
6
2 1
2 1
2 1
2 1
2 1
2 1
```

## Sample Input 8

```text
5
5 3 6 7 8
```

## Sample Output 8

```text
8
5 2
4 2
3 2
1 3
1 3
2 1
4 1
5 1
```

## Sample Input 9

```text
6
3 3 80 3 8 3
```

## Sample Output 9

```text
4
5 1
3 1
3 1
3 1
```

## Sample Input 10

```text
4
19 40 19 55
```

## Sample Output 10

```text
9
4 2
2 1
1 2
1 2
3 2
3 2
1 4
2 4
3 4
```

The printed operations are one valid sequence. Any other sequence that makes the array equal in at most `30n` operations is also correct.
