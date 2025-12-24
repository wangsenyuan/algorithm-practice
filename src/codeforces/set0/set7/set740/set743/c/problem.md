# Problem Description

Vladik and Chloe decided to determine who of them is better at math. Vladik claimed that for any positive integer $n$ he can represent fraction $\frac{1}{n}$ as a sum of three distinct positive fractions in form $\frac{1}{x} + \frac{1}{y} + \frac{1}{z}$.

Help Vladik with that, i.e for a given $n$ find three distinct positive integers $x$, $y$ and $z$ such that $\frac{1}{n} = \frac{1}{x} + \frac{1}{y} + \frac{1}{z}$. Because Chloe can't check Vladik's answer if the numbers are large, he asks you to print numbers not exceeding $10^9$.

If there is no such answer, print $-1$.

## Input

The single line contains single integer $n$ ($1 \leq n \leq 10^4$).

## Output

If the answer exists, print $3$ distinct numbers $x$, $y$ and $z$ ($1 \leq x, y, z \leq 10^9$, $x \neq y$, $x \neq z$, $y \neq z$). Otherwise print $-1$.

If there are multiple answers, print any of them.

## Examples

### Example 1

**Input:**

```text
3
```

**Output:**

```text
2 7 42
```

### Example 2

**Input:**

```text
7
```

**Output:**

```text
7 8 56
```


### ideas
