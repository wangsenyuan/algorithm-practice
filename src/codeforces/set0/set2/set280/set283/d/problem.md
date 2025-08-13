# Problem D: Cool Sequences

## Problem Description

Bessie and the cows have recently been playing with "cool" sequences and are trying to construct some. Unfortunately they are bad at arithmetic, so they need your help!

## Definitions

A pair $(x, y)$ of positive integers is **"cool"** if $x$ can be expressed as the sum of $y$ consecutive integers (not necessarily positive).

A sequence $(a_1, a_2, \ldots, a_n)$ is **"cool"** if the pairs $(a_1, a_2), (a_2, a_3), \ldots, (a_{n-1}, a_n)$ are all cool.

## Task

The cows have a sequence of $n$ positive integers, $a_1, a_2, \ldots, a_n$. In one move, they may replace some $a_i$ with any other positive integer (there are no other limits on the new value of $a_i$).

**Determine the smallest number of moves needed to make the resulting sequence cool.**

## Input

- The first line contains a single integer, $n$ ($2 \leq n \leq 5000$)
- The next line contains $n$ space-separated integers, $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^{15}$)

> **Note for C++ users:** Please do not use the `%lld` specifier to read or write 64-bit integers. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Output

A single integer, the minimum number of $a_i$ that must be changed to make the sequence cool.

## Examples

### Example 1
**Input:**
```
3
6 4 1
```

**Output:**
```
0
```

**Explanation:** The sequence is already cool, so we don't need to change any elements.

### Example 2
**Input:**
```
4
20 6 3 4
```

**Output:**
```
2
```

**Explanation:** We can change $a_2$ to 5 and $a_3$ to 10 to make $(20, 5, 10, 4)$ which is cool. This changes 2 elements.

## Solution

Here is a full solution to Codeforces #174 div 1 D.

### Key Observations

Let $\nu_2(n)$ denote the exponent of the largest power of 2 that divides $n$. For example:
- $\nu_2(5) = 0$
- $\nu_2(96) = 5$

Let $f(n)$ denote the largest odd factor of $n$.

### Mathematical Formula

The sum of an arithmetic series can be expressed as:

$$\sum_{i=1}^{y} (a + i-1) = y \cdot a + \frac{y(y-1)}{2}$$

### Cool Pair Characterization

**Claim:** The pair $(x, y)$ is cool if and only if $f(y)$ divides $f(x)$ and one of the following is true:

1. $\nu_2(x) + 1 = \nu_2(y)$
2. $\nu_2(y) = 0$

**Proof:** This can be proven by casework on the parity of $y$.

- **If $y$ is odd:** The average term of the arithmetic sequence is an integer, so $f(y) = y$ divides $f(x)$ and $\nu_2(y) = 0$.

- **If $y$ is even:** The average is of the form $0.5 \cdot k$ where $k$ is odd, so $2y$ divides $x$. It follows that $y$ divides $x$, so $f(y)$ divides $f(x)$, and furthermore $\nu_2(x) + 1 = \nu_2(y)$.

### Sequence Construction

From this observation, it follows that for fixed $a_i, a_j$ ($i < j$), we can construct a cool sequence $a_i = b_i, b_{i+1}, \ldots, b_{j-1}, b_j = a_j$ if and only if:

1. $f(a_j - a_i)$ divides $f(a_j)$
2. Either $\nu_2(a_i) + j - i = \nu_2(a_j)$ or $\nu_2(a_j) \leq j - i - 1$

### Dynamic Programming Solution

We can solve this problem using dynamic programming where the $k$-th state is the maximum number of $a_i$ ($i \leq k$) we can keep so that it is possible to make $a_1, \ldots, a_k$ cool.

Then the answer is just $n - \max(dp[1], dp[2], \ldots, dp[n])$.

