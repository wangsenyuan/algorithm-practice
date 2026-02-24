# Problem C

To confuse the opponents, the Galactic Empire represents fractions in an unusual way: as two sets of integers. The **numerator** is the product of the numbers in the first set; the **denominator** is the product of the numbers in the second set. The programs that work with fractions in this form are incomplete — they do not support reducing fractions. Implement reduction: given two such sets, output two sets that represent the same fraction in lowest terms.

---

## Input

- First line: two integers $n$, $m$ ($1 \le n, m \le 10^5$) — the number of factors in the numerator set and the denominator set.
- Second line: $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le 10^7$) — the numerator factors.
- Third line: $m$ integers $b_1, b_2, \ldots, b_m$ ($1 \le b_i \le 10^7$) — the denominator factors.

## Output

Print the reduced fraction in the same format as the input:

- First line: two integers $n_{out}$, $m_{out}$ ($1 \le n_{out}, m_{out} \le 10^5$).
- Second line: $n_{out}$ integers — the numerator factors ($1 \le a_{out,i} \le 10^7$).
- Third line: $m_{out}$ integers — the denominator factors ($1 \le b_{out,i} \le 10^7$).

The printed fraction must be in lowest terms: no integer $x > 1$ may divide both the product of the numerator set and the product of the denominator set. If multiple valid answers exist, output any.

---

## Examples

### Example 1

**Input:**

```
3 2
100 5 2
50 10
```

**Output:**

```
2 3
2 1
1 1 1
```

**Note:** Numerator $100 \cdot 5 \cdot 2 = 1000$, denominator $50 \cdot 10 = 500$. $\gcd(1000, 500) = 500$, so the reduced fraction is $1000/500 = 2/1$.

### Example 2

**Input:**

```
4 3
2 5 10 20
100 1 3
```

**Output:**

```
1 1
20
3
```

**Note:** Numerator $2 \cdot 5 \cdot 10 \cdot 20 = 2000$, denominator $100 \cdot 1 \cdot 3 = 300$. $\gcd(2000, 300) = 100$, so the reduced fraction is $20/3$.
