# E — XOR neighbor operations

## Problem description

You are given an array **a** of length **n**. For each index **i** with 1 ≤ i < n, you may perform the following operation **at most once**:

- Assign **a[i] := a[i] ⊕ a[i+1]**, where **⊕** is bitwise XOR.

You may choose which indices to use and in **any order** (sequentially).

You are also given an array **b** of length **n**. Determine whether **a** can be transformed into **b**.

## Input

The first line contains an integer **t** (1 ≤ t ≤ 10⁴) — the number of test cases.

For each test case:

- The first line contains an integer **n** (2 ≤ n ≤ 2⋅10⁵).
- The second line contains **n** integers **a₁, a₂, …, aₙ** (0 ≤ aᵢ < 2³⁰).
- The third line contains **n** integers **b₁, b₂, …, bₙ** (0 ≤ bᵢ < 2³⁰).

It is guaranteed that the sum of **n** over all test cases does not exceed 2⋅10⁵.

## Output

For each test case, print **YES** if **a** can become **b**, otherwise **NO**. You may use any letter case (e.g. `yes`, `Yes`, `YES`).

## Example

**Input:**

```
7
5
1 2 3 4 5
3 2 7 1 5
3
0 0 1
1 0 1
3
0 0 1
0 0 0
4
0 0 1 2
1 3 3 2
6
1 1 4 5 1 4
0 5 4 5 5 4
3
0 1 2
2 3 2
2
10 10
11 10
```

**Output:**

```
YES
NO
NO
NO
YES
NO
NO
```

## Note

In the first test case, one valid sequence is:

1. Choose **i = 3**: **a[3] := a[3] ⊕ a[4] = 7** → **a** becomes `[1, 2, 7, 4, 5]`.
2. Choose **i = 4**: **a[4] := a[4] ⊕ a[5] = 1** → **a** becomes `[1, 2, 7, 1, 5]`.
3. Choose **i = 1**: **a[1] := a[1] ⊕ a[2] = 3** → **a** becomes `[3, 2, 7, 1, 5]`.
