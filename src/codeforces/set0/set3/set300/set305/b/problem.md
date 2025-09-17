# Problem B: Continued Fraction Comparison

A continued fraction of height n is a fraction of the form:

```
a₁ + 1/(a₂ + 1/(a₃ + 1/(... + 1/aₙ)))
```

You are given two rational numbers:
- One is represented as a simple fraction p/q
- The other is represented as a finite continued fraction of height n

Check if they are equal.

## Input

The first line contains two space-separated integers p, q (1 ≤ q ≤ p ≤ 10^18) — the numerator and the denominator of the first fraction.

The second line contains integer n (1 ≤ n ≤ 90) — the height of the second fraction. 

The third line contains n space-separated integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10^18) — the continued fraction coefficients.

**Note:** Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Output

Print "YES" if these fractions are equal and "NO" otherwise.

## Examples

### Example 1
**Input:**
```
9 4
2
2 4
```
**Output:**
```
YES
```

### Example 2
**Input:**
```
9 4
3
2 3 1
```
**Output:**
```
YES
```

### Example 3
**Input:**
```
9 4
3
1 2 4
```
**Output:**
```
NO
```

## Note

In the first sample: 9/4 = 2 + 1/4 = 2.25

In the second sample: 9/4 = 2 + 1/(3 + 1/1) = 2 + 1/4 = 2.25

In the third sample: 1 + 1/(2 + 1/4) = 1 + 1/(9/4) = 1 + 4/9 = 13/9 ≠ 9/4