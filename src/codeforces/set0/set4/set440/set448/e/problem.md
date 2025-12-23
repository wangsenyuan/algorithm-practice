# Problem Description

Bizon the Champion isn't just friendly, he also is a rigorous coder.

## Function Definition

Let's define function `f(a)`, where `a` is a sequence of integers. Function `f(a)` returns the following sequence: first all divisors of `a1` go in the increasing order, then all divisors of `a2` go in the increasing order, and so on till the last element of sequence `a`.

For example, `f([2, 9, 1]) = [1, 2, 1, 3, 9, 1]`.

## Sequence Definition

Let's determine the sequence `Xi`, for integer `i` (i ≥ 0):

- `X0 = [X]` ([X] is a sequence consisting of a single number X)
- `Xi = f(Xi - 1)` (i > 0)

For example, at `X = 6` we get:

- `X0 = [6]`
- `X1 = [1, 2, 3, 6]`
- `X2 = [1, 1, 2, 1, 3, 1, 2, 3, 6]`

## Task

Given the numbers `X` and `k`, find the sequence `Xk`. As the answer can be rather large, find only the first 10^5 elements of this sequence.

## Input

A single line contains two space-separated integers — `X` (1 ≤ X ≤ 10^12) and `k` (0 ≤ k ≤ 10^18).

## Output

Print the elements of the sequence `Xk` in a single line, separated by a space. If the number of elements exceeds 10^5, then print only the first 10^5 elements.

## Examples

### Example 1

**Input:**

```text
6 1
```

**Output:**

```text
1 2 3 6
```

### Example 2

**Input:**

```text
4 2
```

**Output:**

```text
1 1 2 1 2 4
```

### Example 3

**Input:**

```text
10 3
```

**Output:**

```text
1 1 1 2 1 1 5 1 1 2 1 5 1 2 5 10
```


### ideas
1. [10], [1, 2, 5, 10], [1, 1, 2, 1, 5, 1, 2, 5, 10]
2. g(.., i) 表示进行i次后的序列的长度
3. g(x, 0) = 1
4. g(x, 1) = g(f[0], 0) + g(f[1], 0) + .. + g(x, 0)
5.         = g(1, 0) + g(2, 0) + g(5, 0) + g(10, 0) = 4
6. g(x, 2) = g(1, 1) + g(2, 1) + g(5, 1) + g(10, 1) = 
7.         = 1 + 2 + 2 + 4 = 8
8. g(1, i) = 1
9. g(x, i) = g(x, i - 1) + g(1, i - 1) + g(..., i - 1)
10. 2， g(2, k) = k + 1, 质数的长度
11. g(x, k) = 1 +  g(lpf[x], k - 1) + ... + g(x, k  - 1)
12.  