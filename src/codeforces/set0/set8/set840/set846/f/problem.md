# Problem F

## Description

You are given an array `a` consisting of `n` positive integers. You pick two integer numbers `l` and `r` from 1 to `n`, inclusive (numbers are picked randomly, equiprobably and independently). If `l > r`, then you swap values of `l` and `r`. You have to calculate the expected value of the number of unique elements in segment of the array from index `l` to index `r`, inclusive (1-indexed).

## Input

The first line contains one integer number `n` (1 ≤ n ≤ 10^6). The second line contains `n` integer numbers `a1, a2, ..., an` (1 ≤ ai ≤ 10^6) — elements of the array.

## Output

Print one number — the expected number of unique elements in chosen segment.

Your answer will be considered correct if its absolute or relative error doesn't exceed 10^-4 — formally, the answer is correct if |x - y| / max(1, |x|) ≤ 10^-4, where `x` is jury's answer, and `y` is your answer.

## Examples

### Example 1
**Input:**
```
2
1 2
```

**Output:**
```
1.500000
```

### Example 2
**Input:**
```
2
2 2
```

**Output:**
```
1.000000
```