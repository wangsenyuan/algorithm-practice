# Problem Description

You are given n integers a₁, a₂, ..., aₙ. Find the number of pairs of indexes i, j (i < j) that aᵢ + aⱼ is a power of 2 (i.e. some integer x exists so that aᵢ + aⱼ = 2ˣ).

## Input

The first line contains the single positive integer n (1 ≤ n ≤ 10⁵) — the number of integers.

The second line contains n positive integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁹).

## Output

Print the number of pairs of indexes i, j (i < j) that aᵢ + aⱼ is a power of 2.

## Examples

### Example 1

**Input:**

```text
4
7 3 2 1
```

**Output:**

```text
2
```

### Example 2

**Input:**

```text
3
1 1 1
```

**Output:**

```text
3
```

## Note

In the first example the following pairs of indexes include in answer: (1, 4) and (2, 4).

In the second example all pairs of indexes (i, j) (where i < j) include in answer.
