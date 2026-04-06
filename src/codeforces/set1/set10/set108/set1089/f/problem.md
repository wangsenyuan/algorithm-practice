# F. Fractions

You are given a positive integer `n`.

Find a sequence of fractions `a_i` / `b_i` for `i = 1 … k` (where `a_i` and `b_i` are positive integers) for some `k` such that:

- for each `i` from 1 to `k`, `b_i` divides `n`, and `1 < b_i < n`;
- for each `i` from 1 to `k`, `1 ≤ a_i < b_i`;
- `a_1/b_1 + a_2/b_2 + … + a_k/b_k = 1 - 1/n`.

## Input

The input consists of a single integer `n` (`2 ≤ n ≤ 10^9`).

## Output

In the first line print `YES` if there exists such a sequence of fractions or `NO` otherwise.

If there exists such a sequence, the next lines should contain a description of the sequence in the following format.

The second line should contain integer `k` (`1 ≤ k ≤ 100000`) — the number of elements in the sequence. It is guaranteed that if such a sequence exists, then there exists a sequence of length at most 100000.

The next `k` lines should contain the fractions of the sequence with two integers `a_i` and `b_i` on each line.

## Examples

### Example 1

**Input**

```
2
```

**Output**

```
NO
```

### Example 2

**Input**

```
6
```

**Output**

```
YES
2
1 2
1 3
```

## Note

In the second example there is a sequence 1/2, 1/3 such that 1/2 + 1/3 = 1 − 1/6.


### ideas
1. got (n - 1) / n those b > 1 and n % b = 0
2. 如果 n = 7, (prime number) => -1
3. 假设 n = u * v, u < v
4. a / u + b / v = (n - 1) / n
5. a * v + b * u = n - 1
6. b * u < n and a * v < n must hold
7. b < v must hold (becase u * v = n)
8. b * u + 1 要能整除v
9. a * v + 1 要能整除u
10. 感觉肯定有答案（但是怎么构造呢？）
11. 比如n = 8, 1/2 + 1/4 => -1
12. n = 10 => 1/2 + 2/5 = 9 / 10 => work
13. n = 15 => 1/3 + 3/5 = 14/ 15 
14. 这里让a = 1, 那么 v + b * u = n - 1
15. v + b * u = u * v - 1
16. v + 1 = u * (v - b)
17. a * v+1是u得倍数
18. 比如上面 (5 + 1) 是2的倍数，（5 + 1）是3的倍数
19. 

## key insights

1. We want
   - `a1 / b1 + a2 / b2 + ... = 1 - 1 / n = (n - 1) / n`
   - and every denominator must be a proper divisor of `n`.

2. If `n` is a prime power, the answer is impossible.
   - Suppose all denominators divide `n = p^t`.
   - Then every denominator is also a power of `p`.
   - So every fraction has denominator a power of `p`, and after reducing to denominator `n`, each term contributes a multiple of `p`.
   - Hence the whole numerator must be divisible by `p`.
   - But the target numerator is `n - 1 = p^t - 1`, which is not divisible by `p`.
   - Contradiction.

3. So a solution can exist only if `n` has at least two distinct prime factors.

4. In that case, two fractions are enough.
   - Let `x` be one prime-power factor of `n`, and let `y = n / x`.
   - Then `gcd(x, y) = 1`, and both `x` and `y` are proper divisors of `n`.
   - We try to build
     `a / x + b / y = (n - 1) / n`.

5. Multiply by `n = x * y`:
   - `a * y + b * x = x * y - 1`.

6. Since `x` and `y` are coprime, we can solve this by choosing `a` modulo `x`.
   - We need
     `a * y ≡ -1 (mod x)`.
   - Because `gcd(y, x) = 1`, `y` has an inverse modulo `x`.
   - So we can compute
     `a ≡ -y^{-1} (mod x)`,
     and then derive
     `b = (x * y - 1 - a * y) / x`.

7. This construction always gives valid bounds:
   - `1 <= a < x`
   - `1 <= b < y`
   - so both fractions are valid and their denominators are proper divisors of `n`.

8. Therefore:
   - `NO` iff `n` is a prime power
   - otherwise `YES`, and a solution with exactly two fractions always exists.
