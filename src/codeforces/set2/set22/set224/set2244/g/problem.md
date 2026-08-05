# G. Yura's Homework

[Problem link](https://codeforces.com/problemset/problem/2244/G)

time limit per test: *(unknown)*

memory limit per test: *(unknown)*

input: standard input

output: standard output

## Problem

Yura has `n` homework assignments. For each assignment, its weight `a_i` is
known — the number of course points Yura will receive if he completes it.

Yura wants to choose a subset of assignments to maximize the total number of
points. However, he has one problem: some assignments are too time-consuming. If
Yura works on two assignments at positions `i` and `j` (`i ≠ j`), there must be
enough other assignments between them; otherwise, he will get distracted and
fail to complete them.

Formally, for any two chosen assignments with indices `i` and `j`, the following
condition must hold:

```text
|i − j| > max(a_i, a_j)
```

Find the maximum total weight Yura can obtain by choosing a subset of
assignments satisfying this condition.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `0 ≤ a_i ≤ 10^9`
- Sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains a single integer `t` — the number of test cases.

The first line of each test case contains a single integer `n` — the size of the
array `a`.

The second line contains `n` integers `a_1, a_2, …, a_n` — the elements of the
array.

## Output

For each test case, output a single integer — the maximum total weight of the
selected assignments.

## Example

### Input

```text
3
5
3 1 1 1 3
3
2 1 2
7
4 1 5 1 1 4 1
```

### Output

```text
6
2
8
```

## Note

In the first example, it is optimal to choose assignments with indices `1`
and `5`.

In the third example, it is optimal to choose assignments with indices `1`
and `6`.

## ideas
1. let i > j, then abs(i - j) > max(a[i], a[j]) => i - j > max(a[i], a[j])
2. 考虑 j = i - 1
3. i - j = 1 > max(a[i], a[j]) (不一定成立)
4. 假设 a[j] > a[i]
5. 那么在计算到j的时候, i > j + a[j], 如果a[i] <= a[j], 就可以更新
6. 如果 a[i] > a[j], 那么在计算i的时候, i - a[i] > j

For the implemented left-to-right DP, when the current position is `j` and a
previous chosen position is `i`, the compatibility condition is:

```text
i + a_i < j
i < j - a_j
```

The inequalities are strict. The Fenwick tree only stores previous positions
after their own blocking interval has ended (`i + a_i < j`), and the query for
the current value `a_j` must therefore use the last allowed previous index
`j - a_j - 1`.

## Submission notes

- RE cause: the Fenwick tree was initially allocated with size `n + 1`. The
  implementation stores one-based positions by calling `update(i, ...)` and
  then doing `i++` inside `Fenwick.update`. When `i == n`, this touches internal
  index `n + 1`, so the tree must be allocated as `n + 2`.
- Correctness boundary: the pair condition is strict. For current position `j`,
  the query must stop at `j - a_j - 1`, not `j - a_j`; otherwise cases like
  `[1, 1, 2]` incorrectly allow positions `1` and `2` together.
