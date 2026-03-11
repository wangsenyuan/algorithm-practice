# Problem

**题意简述**：给定两个长度为 `n` 的数组 `a` 和 `b`。你可以任意次选择位置 `i (1 ≤ i < n)`，执行 `a[i] = b[i+1]` 或 `b[i] = a[i+1]`。在所有操作开始前，至多还能选择一个位置 `i`，同时删除 `a[i]` 和 `b[i]`。定义两数组的匹配数为对应位置上相等的个数，要求最大化最终匹配数。

## Input

- The first line contains an integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.
- The first line of each test case contains an integer `n` (`2 ≤ n ≤ 2 * 10^5`) — the length of `a` and `b`.
- The second line contains `n` integers `a1, a2, ..., an` (`1 ≤ ai ≤ n`) — the elements of `a`.
- The third line contains `n` integers `b1, b2, ..., bn` (`1 ≤ bi ≤ n`) — the elements of `b`.
- It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print a single integer — the maximum number of matches.

## Examples

### Example 1

**Input**

```text
10
4
1 3 1 4
4 3 2 2
6
2 1 5 3 6 4
3 2 4 5 1 6
2
1 2
2 1
6
2 5 1 3 6 4
3 5 2 3 4 6
4
1 3 2 2
2 1 3 4
8
3 1 4 6 2 2 5 7
4 2 3 7 1 1 6 5
10
5 1 2 7 3 9 4 10 6 8
6 2 3 6 4 10 5 1 7 9
5
3 2 4 1 5
2 4 5 1 3
7
2 2 6 4 1 3 5
3 1 6 5 1 4 2
5
4 1 3 2 5
3 2 1 5 4
```

**Output**

```text
3
3
0
4
3
5
6
4
5
2
```
