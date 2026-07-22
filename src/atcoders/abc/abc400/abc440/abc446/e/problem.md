# E - Multiple-Free Sequences

[Problem link](https://atcoder.jp/contests/abc446/tasks/abc446_e)

Among the integer pairs $(x, y)$ satisfying $0 \leq x,y \leq M-1$, count the pairs for which the
infinite sequence $(s_1, s_2, \ldots)$ defined by the following recurrence contains no multiples of
$M$:

- $s_1 = x$
- $s_2 = y$
- $s_n = A s_{n-1} + B s_{n-2}$ for $n \geq 3$

## Constraints

- $2 \leq M \leq 1000$
- $0 \leq A, B \leq M-1$
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
M A B
```

## Output

Output the answer on one line.

## Sample Input 1

```text
4 1 2
```

## Sample Output 1

```text
7
```

The seven valid pairs are $(1,1)$, $(1,3)$, $(2,1)$, $(2,2)$, $(2,3)$, $(3,1)$, and $(3,3)$.

## Sample Input 2

```text
446 1 1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
1000 784 385
```

## Sample Output 3

```text
995373
```
