# GukiZiana Function

Professor GukiZ was playing with arrays again and accidentally discovered a new function, which he called **GukiZiana**. 

## Problem Description

For a given array `a`, indexed with integers from 1 to `n`, and a number `y`, the function `GukiZiana(a, y)` represents the maximum value of `j - i`, such that `a[j] = a[i] = y`. 

If there is no `y` as an element in `a`, then `GukiZiana(a, y)` is equal to **-1**.

## Query Types

GukiZ prepared a problem for you with two types of queries:

### Type 1: Range Update
- **Format**: `1 l r x`
- **Description**: Increase values of all `a[i]` such that `l ≤ i ≤ r` by the non-negative integer `x`

### Type 2: Function Query
- **Format**: `2 y`
- **Description**: Find the value of `GukiZiana(a, y)`

For each query of type 2, print the answer and make GukiZ happy!

## Input

- **First line**: Two integers `n, q` (1 ≤ n ≤ 5×10^5, 1 ≤ q ≤ 5×10^4) - size of array `a` and number of queries
- **Second line**: `n` integers `a[1], a[2], ..., a[n]` (1 ≤ a[i] ≤ 10^9) forming array `a`
- **Next `q` lines**: Each line contains either four or two numbers:
  - If line starts with `1`: Query format `1 l r x` (1 ≤ l ≤ r ≤ n, 0 ≤ x ≤ 10^9)
  - If line starts with `2`: Query format `2 y` (1 ≤ y ≤ 10^9)

## Output

For each query of type 2, print the value of `GukiZiana(a, y)` for the given `y` value.

## Examples

### Example 1
**Input:**
```
4 3
1 2 3 4
1 1 2 1
1 1 1 1
2 3
```

**Output:**
```
2
```

### Example 2
**Input:**
```
2 3
1 2
1 2 2 1
2 3
2 4
```

**Output:**
```
0
-1
```