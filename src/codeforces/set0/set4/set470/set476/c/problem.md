# Dreamoon and Sums

Dreamoon loves summing up something for no reason. One day he obtains two integers a and b occasionally. He wants to calculate the sum of all nice integers. Positive integer x is called nice if mod(x, b) != 0 and div(x, b) / mod(x, b) = k, where k is some integer in range [1, a].

The answer may be large, so please print its remainder modulo 1 000 000 007 (10^9 + 7). Can you compute it faster than Dreamoon?

## Input

The single line of the input contains two integers a, b (1 ≤ a, b ≤ 10^7).

## Output

Print a single integer representing the answer modulo 1 000 000 007 (10^9 + 7).

## Examples

### Example 1

**Input:**

```
1 1
```

**Output:**

```
0
```

### Example 2

**Input:**

```
2 2
```

**Output:**

```
8
```

## Note

For the first sample, there are no nice integers because mod(x, b) is always zero when b = 1.

For the second sample, the set of nice integers is {3, 5}.
