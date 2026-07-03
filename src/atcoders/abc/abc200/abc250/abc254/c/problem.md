# C - K Swap

[Problem link](https://atcoder.jp/contests/abc254/tasks/abc254_c)

**Contest:** [AtCoder Beginner Contest 254](https://atcoder.jp/contests/abc254)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

You are given a length-`N` sequence `A = (a_1, ..., a_N)` and an integer `K`.

You may perform the following operation zero or more times:

- Choose `i` with `1 <= i <= N - K`, then swap `a_i` and `a_{i+K}`.

Determine whether `A` can be sorted in non-decreasing order.

## Constraints

- `2 <= N <= 2 * 10^5`
- `1 <= K <= N - 1`
- `1 <= a_i <= 10^9`
- All input values are integers

## Input

```text
N K
a_1 ... a_N
```

## Output

Print `Yes` if sorting is possible, otherwise print `No`.

## Sample Input 1

```text
5 2
3 4 1 3 4
```

## Sample Output 1

```text
Yes
```

One valid sequence:

- Swap `a_1` and `a_3` → `(1, 4, 3, 3, 4)`
- Swap `a_2` and `a_4` → `(1, 3, 3, 4, 4)`

## Sample Input 2

```text
5 3
3 4 1 3 4
```

## Sample Output 2

```text
No
```

## Sample Input 3

```text
7 5
1 2 3 4 5 5 10
```

## Sample Output 3

```text
Yes
```

The sequence is already sorted.

### ideas
1. 
