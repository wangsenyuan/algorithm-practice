# Problem E: Array Queries

## Problem Description

You've got an array, consisting of n integers: a₁, a₂, ..., aₙ. Your task is to quickly run the queries of two types:

1. **Assign value x to all elements from l to r inclusive.** After such query the values of the elements of array aₗ, aₗ₊₁, ..., aᵣ become equal to x.
2. **Calculate and print sum** $\sum_{i=l}^{r}{a_i\times(r-l+1)}^k$, where k doesn't exceed 5. As the value of the sum can be rather large, you should print it modulo 1000000007 (10⁹ + 7).

## Input

- The first line contains two integers n and m (1 ≤ n, m ≤ 10⁵), showing how many numbers are in the array and the number of queries, correspondingly.
- The second line contains n integers: a₁, a₂, ..., aₙ (0 ≤ aᵢ ≤ 10⁹) — the initial values of the array elements.
- Then m queries follow, one per line:
  - The assign query has the following format: `"1 l r x"` (1 ≤ l ≤ r ≤ n; 0 ≤ x ≤ 10⁹).
  - The query to calculate the sum has the following format: `"2 l r k"` (1 ≤ l ≤ r ≤ n; 0 ≤ k ≤ 5).

All numbers in the input are integers.

## Output

For each query to calculate the sum, print an integer — the required sum modulo 1000000007 (10⁹ + 7).

## ideas
1. For $(a - b)^3$, using binomial expansion:
   $(a - b)^3 = a^3 - 3a^2b + 3ab^2 - b^3$

2. **Matrix Exponentiation for Power Sums**: 
   Let $f(n) = 1^k + 2^k + 3^k + ... + n^k$
   
   For small values of $k$, this can be represented using matrix exponentiation:
   
   **For k = 1**: $f(n) = \frac{n(n+1)}{2}$
   
   **For k = 2**: $f(n) = \frac{n(n+1)(2n+1)}{6}$
   
   **For k = 3**: $f(n) = \frac{n^2(n+1)^2}{4}$
   
   **For k = 4**: $f(n) = \frac{n(n+1)(2n+1)(3n^2+3n-1)}{30}$
   
   **For k = 5**: $f(n) = \frac{n^2(n+1)^2(2n^2+2n-1)}{12}$
   
   **Matrix Approach**: For larger $k$ or when you need to compute $f(n)$ for many values, you can use matrix exponentiation with a $(k+1) \times (k+1)$ matrix that represents the recurrence relation.

3. **Combining Power Sums with Shifts**:
   Given four numbers $[a, b, c, d]$:
   - $L = a \cdot 1^k + b \cdot 2^k$
   - $R = c \cdot 1^k + d \cdot 2^k$
   
   To get $F = a \cdot 1^k + b \cdot 2^k + c \cdot 3^k + d \cdot 4^k$:
   
   **Key Insight**: $3^k = (1+2)^k$ and $4^k = (2+2)^k = (2 \cdot 1 + 2)^k$
   
   Using binomial expansion:
   - $3^k = (1+2)^k = \sum_{i=0}^k \binom{k}{i} \cdot 1^{k-i} \cdot 2^i$
   - $4^k = (2+2)^k = 2^k \cdot (1+1)^k = 2^k \cdot \sum_{i=0}^k \binom{k}{i} \cdot 1^{k-i} \cdot 1^i = 2^k \cdot 2^k = 4^k$
   
   **For k = 1**: $F = L + R + 2 \cdot (c + d)$
   **For k = 2**: $F = L + R + 4 \cdot (c + d) + 4 \cdot c$
   **For k = 3**: $F = L + R + 8 \cdot (c + d) + 12 \cdot c + 8 \cdot c$
   
   **General Formula**: $F = L + R + \sum_{i=1}^k \binom{k}{i} \cdot 2^i \cdot (c \cdot 1^{k-i} + d \cdot 2^{k-i})$

4. **Extending to Six Numbers**:
   Given six numbers $[a, b, c, d, e, f]$:
   - $L = a \cdot 1^k + b \cdot 2^k + c \cdot 3^k$
   - $R = d \cdot 1^k + e \cdot 2^k + f \cdot 3^k$
   
   To get $F = a \cdot 1^k + b \cdot 2^k + c \cdot 3^k + d \cdot 4^k + e \cdot 5^k + f \cdot 6^k$:
   
   **Key Insight**: Shift by 3 positions
   - $4^k = (1+3)^k$, $5^k = (2+3)^k$, $6^k = (3+3)^k = 3^k \cdot 2^k$
   
   **For k = 1**: 
   - $4^1 = 1 + 3 = 4$
   - $5^1 = 2 + 3 = 5$ 
   - $6^1 = 3 + 3 = 6$
   - $F = L + R + 3 \cdot (d + e + f)$
   
   **For k = 2**:
   - $4^2 = (1+3)^2 = 1^2 + 2 \cdot 1 \cdot 3 + 3^2 = 1 + 6 + 9 = 16$
   - $5^2 = (2+3)^2 = 2^2 + 2 \cdot 2 \cdot 3 + 3^2 = 4 + 12 + 9 = 25$
   - $6^2 = (3+3)^2 = 9 \cdot (1+1)^2 = 9 \cdot 4 = 36$
   - $F = L + R + 6 \cdot (d + e + f) + 9 \cdot (d + e) + 9 \cdot f$
   
   **General Formula**: $F = L + R + \sum_{i=1}^k \binom{k}{i} \cdot 3^i \cdot (d \cdot 1^{k-i} + e \cdot 2^{k-i} + f \cdot 3^{k-i})$

5. **Detailed Binomial Expansion Combination**:
   
   **Mathematical Foundation**:
   When we shift by $m$ positions, we need to compute $(x + m)^k$ for each base $x$.
   Using binomial theorem: $(x + m)^k = \sum_{i=0}^k \binom{k}{i} \cdot x^{k-i} \cdot m^i$
   
   **Step-by-Step Process**:
   
   **Step 1**: Identify the shift amount $m$
   - For 4 numbers: $m = 2$ (shift from $[1,2]$ to $[3,4]$)
   - For 6 numbers: $m = 3$ (shift from $[1,2,3]$ to $[4,5,6]$)
   
   **Step 2**: Apply binomial expansion to each shifted term
   - $3^k = (1+2)^k = \sum_{i=0}^k \binom{k}{i} \cdot 1^{k-i} \cdot 2^i$
   - $4^k = (2+2)^k = 2^k \cdot (1+1)^k = 2^k \cdot 2^k = 4^k$
   
   **Step 3**: Collect coefficients for each original term
   For 4 numbers $[a,b,c,d]$ with $k=2$:
   - $c \cdot 3^2 = c \cdot (1+2)^2 = c \cdot (1^2 + 2 \cdot 1 \cdot 2 + 2^2) = c \cdot (1 + 4 + 4) = c \cdot 9$
   - $d \cdot 4^2 = d \cdot (2+2)^2 = d \cdot 4 \cdot (1+1)^2 = d \cdot 4 \cdot 4 = d \cdot 16$
   
   **Step 4**: Combine all terms
   $F = a \cdot 1^2 + b \cdot 2^2 + c \cdot 3^2 + d \cdot 4^2$
   $= a \cdot 1 + b \cdot 4 + c \cdot 9 + d \cdot 16$
   $= L + R + 4 \cdot (c + d) + 4 \cdot c$
   
   **For k = 3 with 6 numbers**:
   - $4^3 = (1+3)^3 = 1^3 + 3 \cdot 1^2 \cdot 3 + 3 \cdot 1 \cdot 3^2 + 3^3 = 1 + 9 + 27 + 27 = 64$
   - $5^3 = (2+3)^3 = 2^3 + 3 \cdot 2^2 \cdot 3 + 3 \cdot 2 \cdot 3^2 + 3^3 = 8 + 36 + 54 + 27 = 125$
   - $6^3 = (3+3)^3 = 3^3 \cdot (1+1)^3 = 27 \cdot 8 = 216$
   
   **Implementation Strategy**:
   1. Precompute binomial coefficients $\binom{k}{i}$ for all $i \leq k$
   2. For each shift amount $m$, compute $m^i$ for all $i \leq k$
   3. Apply the expansion formula systematically
   4. Combine like terms efficiently