Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 525 points

## Problem Statement

A positive integer $X$ is called a good integer if and only if it satisfies the following condition:

When $X$ is written in decimal notation, the ones digit, tens digit, … form a non-increasing sequence.

More formally, the unique non-negative integer sequence $(d_0, d_1, \ldots)$ satisfying

\[
X = \sum_{i=0}^{\infty} d_i 10^i \quad (0 \le d_i < 10)
\]

forms a non-increasing sequence.

For example, $112389$, $1$, and $777$ are good integers, but $443$ and $404$ are not good integers.

You are given a positive integer $N$.

Determine whether there exists a good integer that is a multiple of $N$, and if it exists, find its minimum value.

## Constraints

- $1 \le N \le 3 \times 10^6$
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
N
```

## Output

If there does not exist a good integer that is a multiple of $N$, output `-1`.

If it exists, output the minimum value of a good integer that is a multiple of $N$.

## Sample Input 1

```
21
```

## Sample Output 1

```
126
```

$126$ is a multiple of $21$, and we have $6 \ge 2 \ge 1 \ge 0 \ge \ldots$, so it is a good integer. There does not exist a good integer less than $126$ that is a multiple of $21$, so output $126$.

## Sample Input 2

```
10
```

## Sample Output 2

```
-1
```

## Sample Input 3

```
3
```

## Sample Output 3

```
3
```

## Sample Input 4

```
1089
```

## Sample Output 4

```
9999999999999999999999
```

The answer may be $2^{64}$ or greater.

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 525 points

## Problem Statement

A positive integer $X$ is called a good integer if and only if it satisfies the following condition:

When $X$ is written in decimal notation, the ones digit, tens digit, … form a non-increasing sequence.

More formally, the unique non-negative integer sequence $(d_0, d_1, \ldots)$ satisfying
\[
X = \sum_{i=0}^{\infty} d_i 10^i \quad (0 \le d_i < 10)
\]
forms a non-increasing sequence.

For example, $112389$, $1$, and $777$ are good integers, but $443$ and $404$ are not good integers.

You are given a positive integer $N$.

Determine whether there exists a good integer that is a multiple of $N$, and if it exists, find its minimum value.

## Constraints

- $1 \le N \le 3 \times 10^6$
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
N
```

## Output

If there does not exist a good integer that is a multiple of $N$, output `-1`.

If it exists, output the minimum value of a good integer that is a multiple of $N`.

## Sample Input 1

```
21
```

## Sample Output 1

```
126
```

$126$ is a multiple of $21$, and we have $6 \ge 2 \ge 1 \ge 0 \ge \ldots$, so it is a good integer. There does not exist a good integer less than $126$ that is a multiple of $21$, so we output $126$.

## Sample Input 2

```
10
```

## Sample Output 2

```
-1
```

## Sample Input 3

```
3
```

## Sample Output 3

```
3
```

## Sample Input 4

```
1089
```

## Sample Output 4

```
9999999999999999999999
```

The answer may be $2^{64}$ or greater.

Time Limit: 2 sec / Memory Limit: 1024 MiB
