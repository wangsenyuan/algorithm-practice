# E - Prefix Equality

[Problem link](https://atcoder.jp/contests/abc250/tasks/abc250_e)

**Contest:** [AtCoder Beginner Contest 250](https://atcoder.jp/contests/abc250)

time limit: 4 sec

memory limit: 1024 MiB

score: 500 points

You are given integer sequences A = (a_1, ..., a_N) and B = (b_1, ..., b_N), each of length N.

For i = 1, ..., Q, answer the query in the following format.

- If the set of values contained in the first x_i terms of A, (a_1, ..., a_{x_i}), and the set of
  values contained in the first y_i terms of B, (b_1, ..., b_{y_i}), are equal, then print `Yes`;
  otherwise, print `No`.

## Constraints

- 1 <= N, Q <= 2 * 10^5
- 1 <= a_i, b_i <= 10^9
- 1 <= x_i, y_i <= N
- All values in input are integers

## Input

```text
N
a_1 ... a_N
b_1 ... b_N
Q
x_1 y_1
...
x_Q y_Q
```

## Output

Print Q lines. The i-th line should contain the response to the i-th query.

## Sample Input 1

```text
5
1 2 3 4 5
1 2 2 4 3
7
1 1
2 2
2 3
3 3
4 4
4 5
5 5
```

## Sample Output 1

```text
Yes
Yes
Yes
No
No
Yes
No
```

Sets care only about whether each value is contained. For the 3rd query, the first 2 terms of A
contain one 1 and one 2, while the first 3 terms of B contain one 1 and two 2's. The sets of values
are both {1, 2}, so they are equal. For the 6th query, the values appear in different orders, but
the sets are still equal.
