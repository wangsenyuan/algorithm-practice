# Problem Statement

You are given positive integers $A$, $B$, $C$, $D$ satisfying $\frac{B}{A} < \frac{D}{C}$.

Find the smallest positive integer $q$ satisfying the following condition:

There exists a positive integer $p$ such that $\frac{B}{A} < \frac{p}{q} < \frac{D}{C}$.

$T$ test cases are given, so solve each.

## Constraints

- $1 \leq T \leq 2 \times 10^5$
- $1 \leq A, B, C, D \leq 10^{18}$
- $\frac{B}{A} < \frac{D}{C}$
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
T
case_1
case_2
⋮
case_T
```

Here, $case_i$ represents the $i$-th test case.

Each test case is given in the following format:

```
A B C D
```

## Output

Output $T$ lines. The $i$-th line should contain the answer for the $i$-th test case.

## Sample Input 1

```
5
3 2 2 1
5 2 8 3
1 2 2 1
60 191 11 35
40 191 71 226
```

## Sample Output 1

```
3
5
1
226
4
```

Consider the first test case.

For example, if $p = 5$, $q = 3$, then $\frac{2}{3} < \frac{3}{5} < \frac{1}{2}$, so $q = 3$ satisfies the condition.

There is no positive integer less than $3$ that satisfies the condition for $q$, so the answer for the first test case is $3$.
