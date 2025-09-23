# Problem C: Tile Painting

## Problem Statement

Little Joty has got a task to do. She has a line of n tiles indexed from 1 to n. She has to paint them in a strange pattern.

**Painting Rules:**
- An unpainted tile should be painted **Red** if its index is divisible by `a`
- An unpainted tile should be painted **Blue** if its index is divisible by `b`
- A tile with a number divisible by both `a` and `b` can be painted either **Red** or **Blue**

**Rewards:**
- After painting is done, she will get `p` chocolates for each tile painted Red
- She will get `q` chocolates for each tile painted Blue

**Note:** She can paint tiles in any order she wants.

**Goal:** Find the maximum number of chocolates Joty can get.

## Input

The only line contains five integers `n`, `a`, `b`, `p` and `q` (1 ≤ n, a, b, p, q ≤ 10^9).

## Output

Print the only integer `s` — the maximum number of chocolates Joty can get.

**Note:** The answer can be too large, so you should use 64-bit integer type to store it. In C++ you can use the `long long` integer type and in Java you can use `long` integer type.

## Examples

### Example 1
**Input:**
```
5 2 3 12 15
```

**Output:**
```
39
```

### Example 2
**Input:**
```
20 2 3 3 5
```

**Output:**
```
51
```


### ideas
1. x = n / a, y = n / b, z = n / lcm(a, b)
2. if p >= q, x * p + q * (y - z)
3. 