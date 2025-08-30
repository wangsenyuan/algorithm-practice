# Problem Statement

## Definitions

- **AB string**: A non-empty string where every character is either 'A' or 'B'
- **Good set**: A set X consisting of AB strings is called a good set when it satisfies the following condition:
  - Every AB string of length $10^{100}$ has some element of X as a prefix

## Problem

You are given N distinct AB strings: $S_1, S_2, \ldots, S_N$.

For each $k = 1, 2, \ldots, N$, find the number (modulo $998244353$) of subsets of the set $\{S_1, S_2, \ldots, S_k\}$ that are good sets.

## Output

For each $k$, output the number of good subsets modulo $998244353$.

