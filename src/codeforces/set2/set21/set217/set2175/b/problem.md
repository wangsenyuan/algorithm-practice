# B. XOR Array

You are given three integers `n`, `l`, and `r`.

Construct an array `a` of length `n` of **positive** integers with `1 ≤ a_i ≤ 10^9`. For `1 ≤ x ≤ y ≤ n`, let

`f(x, y) = a_x ⊕ a_{x+1} ⊕ … ⊕ a_y`

(bitwise XOR of the subarray `a[x..y]`).

You must satisfy:

- `f(x, y) = 0` **exactly when** `x = l` and `y = r`;
- for every other pair `(x, y)` with `1 ≤ x ≤ y ≤ n`, **`f(x, y) ≠ 0`**.

> **\*** `⊕` is [bitwise XOR](https://en.wikipedia.org/wiki/Bitwise_operation#XOR).

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

Each test case is one line with three integers `n`, `l`, and `r` (`2 ≤ n ≤ 4 · 10^5`, `1 ≤ l < r ≤ n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `5 · 10^5`.

## Output

For each test case, print one line with `n` integers `a_1, a_2, …, a_n`.

A solution always exists. If several arrays work, print any of them.

## Example

**Input**

```text
4
3 1 3
4 1 3
8 2 4
4 3 4
```

**Output**

```text
9 8 1
2 7 5 4
9 1 9 8 10 5 4 9
85484 130377 6031 6031
```

## Note

In the first test case, `f(1, 3) = 9 ⊕ 8 ⊕ 1 = 0`, while every **other** non-empty subarray has non-zero XOR:

- `f(1, 2) = 9 ⊕ 8 = 1 ≠ 0`
- `f(2, 3) = 8 ⊕ 1 = 9 ≠ 0`
- `f(1, 1) = 9 ≠ 0`
- `f(2, 2) = 8 ≠ 0`
- `f(3, 3) = 1 ≠ 0`

In the second test case, `2 ⊕ 7 ⊕ 5 = 0`, while for example `7 ⊕ 5 ⊕ 4 = 6 ≠ 0`.
