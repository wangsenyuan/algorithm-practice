# Problem Description

The new ITone 6 has been released recently and George got really keen to buy it. Unfortunately, he didn't have enough money, so George was going to work as a programmer. Now he faced the following problem at the work.

Given a sequence of n integers p₁, p₂, ..., pₙ. You are to choose k pairs of integers:

**[l₁, r₁], [l₂, r₂], ..., [lₖ, rₖ]** (1 ≤ l₁ ≤ r₁ < l₂ ≤ r₂ < ... < lₖ ≤ rₖ ≤ n; rᵢ - lᵢ + 1 = m)

in such a way that the value of sum is maximal possible. Help George to cope with the task.

## Input

The first line contains three integers n, m and k (1 ≤ (m × k) ≤ n ≤ 5000). The second line contains n integers p₁, p₂, ..., pₙ (0 ≤ pᵢ ≤ 10⁹).

## Output

Print an integer in a single line — the maximum possible value of sum.

## Examples

### Example 1
**Input:**
```
5 2 1
1 2 3 4 5
```

**Output:**
```
9
```

### Example 2
**Input:**
```
7 1 3
2 10 7 18 5 33 0
```

**Output:**
```
61
```

### ideas
1. 长度为m这个限制
2. dp[i][j] 表示到i为止，有j段的最优解
3. dp[i][j] = dp[i-1][j] max sum[i-m + 1:i] + dp[i-m][j-1]
4. 