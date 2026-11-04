# A. Learning Languages

[Problem link](https://codeforces.com/problemset/problem/277/A)

**Contest:** [Codeforces Round 170 (Div. 1)](https://codeforces.com/contest/277)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

BerCorp has `n` employees and `m` official languages. Each employee knows a (possibly empty) subset of those languages. Teaching one language to one employee costs 1 berdollar.

Two employees can correspond directly if they share a language, and indirectly through a chain of employees. Find the minimum cost so that every employee can correspond with every other employee.

## Constraints

- `2 ≤ n, m ≤ 100`
- `0 ≤ k_i ≤ m`
- `1 ≤ a_ij ≤ m`
- languages in one employee's list are distinct

## Input

```text
n m
k1 a11 a12 ... a1k1
...
kn an1 an2 ... ankn
```

## Output

Print the minimum cost.

## Sample Input 1

```text
5 5
1 2
2 2 3
2 3 4
2 4 5
1 5
```

## Sample Output 1

```text
0
```

Languages already connect everyone: `2–3–4–5`.

## Sample Input 2

```text
8 7
0
3 1 2 3
1 1
2 5 4
2 6 7
1 3
2 7 4
1 1
```

## Sample Output 2

```text
2
```

Employee 1 can learn language 2, and employee 8 can learn language 4.

## Sample Input 3

```text
2 2
1 2
0
```

## Sample Output 3

```text
1
```

Employee 2 must learn language 2.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
