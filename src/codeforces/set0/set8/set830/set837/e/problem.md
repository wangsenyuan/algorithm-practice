# Problem Description

Vasya is studying number theory. He has denoted a function f(a, b) such that:

- f(a, 0) = 0;
- f(a, b) = 1 + f(a, b - gcd(a, b)), where gcd(a, b) is the greatest common divisor of a and b.

Vasya has two numbers x and y, and he wants to calculate f(x, y). He tried to do it by himself, but found out that calculating this function the way he wants to do that might take very long time. So he decided to ask you to implement a program that will calculate this function swiftly.

## Input

The first line contains two integer numbers x and y (1 ≤ x, y ≤ 10^12).

## Output

Print f(x, y).

## Examples

### Example 1

**Input:**

```text
3 5
```

**Output:**

```text
3
```

### Example 2

**Input:**

```text
6 3
```

**Output:**

```text
1
```


### ideas
1. let gcd(a, b) = c
2. b - c ?
3. b = n * c, 
4. b - c = (n - 1) * c 
5. 如果 n > 1, 那么 (n - 1) * c 肯定可以整除 c
6. gcd(a, b - c) 有可能变大（但是不会变小）
7. gcd(6, 8) = 2, gcd(6, 8 - 2) = 6
8. 