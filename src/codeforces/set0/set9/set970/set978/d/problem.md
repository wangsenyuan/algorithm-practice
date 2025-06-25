# Problem D: Arithmetic Progression

## Description

Polycarp likes arithmetic progressions. A sequence $[a_1, a_2, \ldots, a_n]$ is called an **arithmetic progression** if for each $i$ $(1 \leq i < n)$ the value $a_{i+1} - a_i$ is the same.

### Examples of Arithmetic Progressions:
- $[42]$
- $[5, 5, 5]$
- $[2, 11, 20, 29]$
- $[3, 2, 1, 0]$

### Examples of Non-Arithmetic Progressions:
- $[1, 0, 1]$
- $[1, 3, 9]$
- $[2, 3, 1]$

**Note:** It follows from the definition that any sequence of length one or two is an arithmetic progression.

## Problem

Polycarp found some sequence of positive integers $[b_1, b_2, \ldots, b_n]$. He agrees to change each element by at most one. In other words, for each element there are exactly three options:
- An element can be decreased by 1
- An element can be increased by 1
- An element can be left unchanged

Determine the **minimum possible number of elements** in $b$ which can be changed (by exactly one), so that the sequence $b$ becomes an arithmetic progression, or report that it is impossible.

**Note:** It is possible that the resulting sequence contains elements equal to 0.

## Input

- The first line contains a single integer $n$ $(1 \leq n \leq 100000)$ — the number of elements in $b$.
- The second line contains a sequence $b_1, b_2, \ldots, b_n$ $(1 \leq b_i \leq 10^9)$.

## Output

- If it is impossible to make an arithmetic progression with the described operations, print `-1`.
- Otherwise, print a non-negative integer — the minimum number of elements to change to make the given sequence become an arithmetic progression.

**Constraint:** The only allowed operation is to add/subtract one from an element (can't use operation twice to the same position).