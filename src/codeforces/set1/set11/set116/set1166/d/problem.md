# Problem D

Given a positive integer m, we say that a sequence x₁, x₂, …, xₙ of positive integers is **m-cute** if for every index i such that 2 ≤ i ≤ n it holds that

xᵢ = xᵢ₋₁ + xᵢ₋₂ + ⋯ + x₁ + rᵢ

for some positive integer rᵢ satisfying 1 ≤ rᵢ ≤ m.

You will be given q queries consisting of three positive integers a, b and m. For each query you must determine whether or not there exists an m-cute sequence whose first term is a and whose last term is b. If such a sequence exists, you must additionally find an example of it.

## Input

The first line contains an integer q (1 ≤ q ≤ 10³) — the number of queries.

Each of the following q lines contains three integers a, b, and m (1 ≤ a, b, m ≤ 10¹⁴, a ≤ b), describing a single query.

## Output

For each query, if no m-cute sequence whose first term is a and whose last term is b exists, print −1.

Otherwise print an integer k (1 ≤ k ≤ 50), followed by k integers x₁, x₂, …, xₖ (1 ≤ xᵢ ≤ 10¹⁴). These integers must satisfy x₁ = a, xₖ = b, and that the sequence x₁, x₂, …, xₖ is m-cute.

It can be shown that under the problem constraints, for each query either no m-cute sequence exists, or there exists one with at most 50 terms.

If there are multiple possible sequences, you may print any of them.

## Example

**Input:**
```
2
5 26 2
3 9 1
```

**Output:**
```
4 5 6 13 26
-1
```

## Note

Consider the sample. In the first query, the sequence 5, 6, 13, 26 is valid since 6 = 5 + 1, 13 = 6 + 5 + 2 and 26 = 13 + 6 + 5 + 2 have the bold values all between 1 and 2, so the sequence is 2-cute. Other valid sequences, such as 5, 7, 13, 26 are also accepted.

In the second query, the only possible 1-cute sequence starting at 3 is 3, 4, 8, 16, …, which does not contain 9.

## Editorial

### Deriving the formula

From the recurrence xₙ = x₁ + x₂ + ⋯ + xₙ₋₁ + rₙ = Sₙ₋₁ + rₙ, we can rearrange:

xₙ = (xₙ₋₁ − rₙ₋₁) + xₙ₋₁ + rₙ = 2xₙ₋₁ + rₙ − rₙ₋₁

Expanding inductively, the general formula is:

xₙ = 2^(n−2)·a + 2^(n−3)·r₂ + 2^(n−4)·r₃ + ⋯ + 2·rₙ₋₂ + rₙ₋₁ + rₙ

The sum of the coefficients of r₂ through rₙ is 2^(n−2). Therefore:

- All rᵢ = 1 gives the minimum: xₙ = 2^(n−2)·(a + 1)
- All rᵢ = m gives the maximum: xₙ = 2^(n−2)·(a + m)

### Impossibility check

For b to be the last term of an m-cute sequence of length n starting at a, it must satisfy:

2^(n−2)·(a + 1) ≤ b ≤ 2^(n−2)·(a + m)

If b doesn't lie in any of these intervals for n = 1, 2, …, 50, then it's impossible. This check runs in O(log(10¹⁴)) per query.

### Construction via binary decomposition

Once we find a valid n, write b = 2^(n−2)·(a + k) + r, where:

- k = (b − 2^(n−2)·a) / 2^(n−2), satisfying 1 ≤ k ≤ m
- r = (b − 2^(n−2)·a) mod 2^(n−2), satisfying 0 ≤ r < 2^(n−2)

Let d₀, d₁, …, dₙ₋₃ be the binary digits of r. Then set:

rᵢ = k + dₙ₋₁₋ᵢ (for i = 2, …, n−1), and rₙ = k (with d₋₁ = 0)

Each rᵢ ∈ {k, k+1} ⊆ [1, m], since when r > 0 we have k ≤ m−1.

Verification: the sum of coefficient · rᵢ terms becomes k·2^(n−2) + r, which equals b − 2^(n−2)·a. ✓

### Complexity

O(q · log(10¹⁴)) — at most ~48 iterations per query.
