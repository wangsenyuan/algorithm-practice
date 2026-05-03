# D. A Cruel Segment's Thesis

You are given `n` segments on a number line. The `i`-th segment is `[l_i, r_i]`. Initially, all segments are **unmarked**.

Repeat the following until no unmarked segment remains:

- In the `k`-th operation, if there are **at least two** unmarked segments, choose any two unmarked segments `[l_i, r_i]` and `[l_j, r_j]`, **mark both**, and add a **new marked** segment `[x_k, y_k]` such that:
  - `l_i ≤ x_k ≤ r_i`
  - `l_j ≤ y_k ≤ r_j`
  - `x_k ≤ y_k`
- If **exactly one** unmarked segment remains, mark it.

Your task is to find the **maximum possible sum** of the lengths of **all marked segments** at the end. The length of `[l, r]` is `r - l`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

For each test case:

- The first line contains a single integer `n` (`1 ≤ n ≤ 2 · 10^5`) — the number of segments.
- Each of the next `n` lines contains two integers `l_i` and `r_i` (`1 ≤ l_i ≤ r_i ≤ 10^9`) — the `i`-th segment.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 · 10^5`.

## Output

For each test case, print a single integer — the maximum possible total length of all marked segments when the process ends.

## Example

**Input**

```text
4
2
1 1000000000
1 1000000000
3
1 10
2 15
3 9
5
1 11
2 7
15 20
1 3
11 15
1
1000000000 1000000000
```

**Output**

```text
2999999997
42
59
0
```

## Note

In the first test case, take the two given segments and form the new segment `[1, 10^9]`.

In the second test case, take `[1, 10]` and `[2, 15]` and form `[1, 15]`. Then only `[3, 9]` is still unmarked, and it is marked in the next step.
