# Problem Description

You are given three sticks with positive integer lengths of a, b, and c centimeters. You can increase length of some of them by some positive integer number of centimeters (different sticks can be increased by a different length), but in total by at most l centimeters. In particular, it is allowed not to increase the length of any stick.

Determine the number of ways to increase the lengths of some sticks so that you can form from them a non-degenerate (that is, having a positive area) triangle. Two ways are considered different, if the length of some stick is increased by different number of centimeters in them.

## Input

The single line contains 4 integers a, b, c, l (1 ≤ a, b, c ≤ 3·10^5, 0 ≤ l ≤ 3·10^5).

## Output

Print a single integer — the number of ways to increase the sizes of the sticks by the total of at most l centimeters, so that you can make a non-degenerate triangle from it.

## Examples

### Example 1

**Input:**

```text
1 1 1 2
```

**Output:**

```text
4
```

### Example 2

**Input:**

```text
1 2 3 1
```

**Output:**

```text
2
```

### Example 3

**Input:**

```text
10 2 1 7
```

**Output:**

```text
0
```

## Note

In the first sample test you can either not increase any stick or increase any two sticks by 1 centimeter.

In the second sample test you can increase either the first or the second stick by one centimeter. Note that the triangle made from the initial sticks is degenerate and thus, doesn't meet the conditions.

## Solution

### Approach

We use an **inclusion-exclusion** approach:
1. Count all possible ways to distribute at most `l` centimeters among the three sticks
2. Count invalid cases (where a triangle cannot be formed)
3. Result = Total ways - Invalid ways

### Counting Total Ways

For each total increase `x` (0 ≤ x ≤ l), we need to count the number of ways to distribute `x` among 3 sticks (allowing 0). This is a classic **stars and bars** problem:

- Number of ways to distribute `x` among 3 sticks = C(x+2, 2) = (x+2)(x+1)/2
- Total ways = Σ(x=0 to l) (x+2)(x+1)/2

### Counting Invalid Cases

A triangle is invalid (degenerate) if one side is **greater than or equal to** the sum of the other two sides. We need to count cases where:
- `a' ≥ b' + c'`, OR
- `b' ≥ a' + c'`, OR  
- `c' ≥ a' + b'`

where `a' = a + x`, `b' = b + y`, `c' = c + z` are the lengths after increases.

These cases are **mutually exclusive** (at most one can be true for positive values), so we can simply add them.

### Counting One Invalid Case

For the case `a' ≥ b' + c'`:
- We need: `a + x ≥ (b + y) + (c + z)`
- Rearranging: `x ≥ (b + c - a) + (y + z)`
- Let `d = b + c - a` (this can be **negative**)
- Let `s = y + z` (sum of increases for sticks b and c)

For each `s` (0 ≤ s ≤ l):
- Number of ways to get `y + z = s`: `(s + 1)` ways (y can be 0 to s, z = s - y)
- `x` must satisfy:
  - `x ≥ d + s` (from the triangle inequality failure)
  - `x + s ≤ l` (total increase constraint)
  - `x ≥ 0` (non-negative increase)
- So `x` is in range `[max(0, d + s), l - s]`
- Number of valid `x` values: `max(0, l - s - max(0, d + s) + 1)`

**Key Insight**: The deficit `d = b + c - a` can be negative. When `d < 0`, we must use `max(0, d + s)` to ensure `x ≥ 0`.

### Time Complexity

- Counting total ways: O(l)
- Counting invalid cases: O(l) for each of the 3 cases
- Overall: **O(l)** where l ≤ 3·10⁵

### Summary

The solution efficiently counts valid triangle configurations by:
1. Using combinatorics (stars and bars) to count total distributions
2. Systematically counting invalid cases where triangle inequalities fail
3. Handling negative deficits correctly to ensure non-negative increases
4. Leveraging the mutual exclusivity of invalid cases to avoid double-counting
