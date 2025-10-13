## Problem Description

Simon has a prime number `x` and an array of non-negative integers `a₁, a₂, ..., aₙ`.

Simon loves fractions very much. Today he wrote out the sum `1/x^a₁ + 1/x^a² + ... + 1/x^aₙ` on a piece of paper. 

After Simon led all fractions to a common denominator and summed them up, he got a fraction `s/t`, where number `t` equals `x^(a₁ + a₂ + ... + aₙ)`. 

Now Simon wants to reduce the resulting fraction.

Help him, find the greatest common divisor of numbers `s` and `t`. As GCD can be rather large, print it as a remainder after dividing it by number `1000000007 (10⁹ + 7)`.

## Input

The first line contains two positive integers `n` and `x` (`1 ≤ n ≤ 10⁵`, `2 ≤ x ≤ 10⁹`) — the size of the array and the prime number.

The second line contains `n` space-separated integers `a₁, a₂, ..., aₙ` (`0 ≤ a₁ ≤ a₂ ≤ ... ≤ aₙ ≤ 10⁹`).

## Output

Print a single number — the answer to the problem modulo `1000000007 (10⁹ + 7)`.

## Examples

### Example 1
**Input:**
```
2 2
2 2
```
**Output:**
```
8
```

### Example 2
**Input:**
```
3 3
1 2 3
```
**Output:**
```
27
```

### Example 3
**Input:**
```
2 2
29 29
```
**Output:**
```
73741817
```

### Example 4
**Input:**
```
4 5
0 0 0 0
```
**Output:**
```
1
```

## Notes

- In the first sample: `1/2² + 1/2² = 1/4 + 1/4 = 2/4 = 1/2`. After conversion: `t = 2^(2+2) = 2⁴ = 16`, and `s = 8`, so `GCD(8, 16) = 8`. Thus, the answer is **8**.

- In the second sample: `1/3¹ + 1/3² + 1/3³ = 9/27 + 3/27 + 1/27 = 13/27`. Here `t = 3^(1+2+3) = 3⁶ = 729`, and `s = 351`. The answer is `GCD(351, 729) = 27`, as `351 = 13·27` and `729 = 27·27`.

- In the third sample: the answer to the problem is `1073741824 mod 1000000007 = 73741817`.

- In the fourth sample: `1/5⁰ + 1/5⁰ + 1/5⁰ + 1/5⁰ = 1 + 1 + 1 + 1 = 4`. Here `t = 5⁰ = 1`, so `s/t = 4/1`, and `GCD(4, 1) = 1`. Thus, the answer is **1**.


### ideas
1. sum = $x^{-a_1} + x^{-a_2} + ... + x^{-a_n}$
2. t = $x^{a_1 + a_2 + ... + a_n}$
3. $sum \times t \div t$
4. let s[i] = $x^{-a_i}$
5. let w[i] = $t \times s_i$ = $x^{sa - a_i}$
6. 那么找到最小的 $sa - a_i$ 是不是就可以了？
7. 1/4 + 1 / 4 = 4 * (1 + 1) / 16
8. 4   2 / 4, 还有2也是要搞出来的
9. 假设a按照降序处理, $x^{a_1}$是第一部分
10. $1 + x^{a_1-a_2} + ... + x^{a_1-a_i} + .. + x^{a_1-a_n}$
11. 