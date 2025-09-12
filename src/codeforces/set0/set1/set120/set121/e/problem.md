# Lucky Numbers Problem

## Problem Description

Petya loves lucky numbers. Everybody knows that lucky numbers are positive integers whose decimal representation contains only the lucky digits **4** and **7**. 

**Examples:**
- Lucky numbers: `47`, `744`, `4`
- Not lucky numbers: `5`, `17`, `467`

Petya has an array consisting of `n` numbers. He wants to perform `m` operations of two types:

### Operations

1. **`add l r d`** — Add an integer `d` to all elements whose indexes belong to the interval from `l` to `r`, inclusive
   - Constraints: `1 ≤ l ≤ r ≤ n`, `1 ≤ d ≤ 10⁴`

2. **`count l r`** — Find and print how many lucky numbers there are among elements with indexes that belong to the interval from `l` to `r`, inclusive
   - Constraints: `1 ≤ l ≤ r ≤ n`
   - Each lucky number should be counted as many times as it appears in the interval

**Note:** The operations are such that after all additions, the array won't have numbers that would exceed `10⁴`.

## Input

- **Line 1:** Two integers `n` and `m` (`1 ≤ n, m ≤ 10⁵`) — the number of numbers in the array and the number of operations
- **Line 2:** `n` positive integers, none of which exceeds `10⁴` — the array numbers
- **Next `m` lines:** Operations, one per line

**Guarantee:** After all operations are fulfilled, each number in the array will not exceed `10⁴`.

## Output

For each operation of the second type (`count`), print a single number on a single line — the number of lucky numbers in the corresponding interval.

## Examples

### Example 1

**Input:**
```
3 6
2 3 4
count 1 3
count 1 2
add 1 3 2
count 1 3
add 2 3 3
count 1 3
```

**Output:**
```
1
0
1
1
```

**Explanation:**
- Initial array: `[2, 3, 4]`
- After first addition (`add 1 3 2`): `[4, 5, 6]`
- After second addition (`add 2 3 3`): `[4, 8, 9]`

### Example 2

**Input:**
```
4 5
4 4 4 4
count 1 4
add 1 4 3
count 1 4
add 2 3 40
count 1 4
```

**Output:**
```
4
4
4
```

**Explanation:**
- Initial array: `[4, 4, 4, 4]`
- After first addition (`add 1 4 3`): `[7, 7, 7, 7]`
- After second addition (`add 2 3 40`): `[7, 47, 47, 7]`