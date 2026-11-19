# K. Addition Robot

[Problem link](https://codeforces.com/problemset/problem/1252/K)

**Contest:** [2019-2020 ICPC, Asia Jakarta Regional Contest](https://codeforces.com/contest/1252)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

The robot stores a string `S` of length `n`. Every character is `A` or `B`. It then receives `q` commands:

- `1 L R`: toggle every character on `[L, R]`. `A` becomes `B`, and `B` becomes `A`.
- `2 L R A B`: return `f(L, R, A, B)`:

```text
function f(L, R, A, B):
    for i from L to R:
        if S[i] = 'A':
            A = A + B
        else:
            B = A + B
    return (A, B)
```

There is at least one command of the second type. Print each returned pair modulo `10^9+7`.

## Constraints

- `1 ≤ n, q ≤ 10^5`
- `S` has length `n` and contains only `A` and `B`
- For a toggle: `1 ≤ L ≤ R ≤ n`
- For a query: `1 ≤ L ≤ R ≤ n` and `0 ≤ A, B ≤ 10^9`

## Input

```text
n q
S
```

Then `q` commands, each either:

```text
1 L R
```

or:

```text
2 L R A B
```

## Output

For each command of type `2`, print `A` and `B` modulo `10^9+7`.

## Sample Input 1

```text
5 3
ABAAA
2 1 5 1 1
1 3 5
2 2 5 0 1000000000
```

## Sample Output 1

```text
11 3
0 1000000000
```

The first call starts from `(1, 1)` and walks `ABAAA`: `(2, 1)`, `(2, 3)`, `(5, 3)`, `(8, 3)`, `(11, 3)`. The toggle turns `S` into `ABBBB`. The last call starts from `(0, 1000000000)`, and every step adds `A` into `B`, so `A` stays `0`.

## Solution

把 `(A, B)` 写成行向量 `[A, B]`，乘在矩阵左边。字符从左到右作用，所以一段的矩阵是这些字符矩阵按从左到右的顺序相乘。线段树结点存的就是这个乘积，左儿子在前：

```text
val[i] = val[left] * val[right]
```

遇到 `A` 时 `A` 变成 `A+B`，`B` 不变：

```text
[A, B] * [[1, 0], [1, 1]] = [A+B, B]
```

遇到 `B` 时 `B` 变成 `A+B`，`A` 不变：

```text
[A, B] * [[1, 1], [0, 1]] = [A, A+B]
```

查询 `[L, R]` 取出这段的矩阵 `M`，答案是 `[A, B] * M`，两个分量再对 `10^9+7` 取模。

### 操作 1

记交换矩阵

```text
P = [[0, 1],
     [1, 0]]
```

`P` 乘在左边是交换两行，乘在右边是交换两列，而且 `P * P = I`。一个矩阵 `M = [[a, b], [c, d]]` 做 `P * M * P` 就是主对角线互换、副对角线互换：

```text
[[a, b],    →    [[d, c],
 [c, d]]          [b, a]]
```

单个字符刚好因此互换：

```text
A = [[1, 0], [1, 1]]  →  [[1, 1], [0, 1]] = B
B = [[1, 1], [0, 1]]  →  [[1, 0], [1, 1]] = A
```

用 `(3, 5)` 核对。`A` 得到 `(8, 5)`，`B` 得到 `(3, 8)`。

一段里面每个字符都翻转时，第 `i` 个矩阵从 `M_i` 变成 `P * M_i * P`。乘积是

```text
(P M_1 P) (P M_2 P) ... (P M_k P)
= P M_1 (P P) M_2 (P P) ... (P P) M_k P
= P (M_1 M_2 ... M_k) P
```

相邻的 `P` 两两抵消，所以整段乘积也做同一种翻转，不用拆开结点。再翻一次：`P (P M P) P = M`，标记打两次就恢复。

以 `AB` 为例。

```text
A * B = [[1, 0], [1, 1]] * [[1, 1], [0, 1]] = [[1, 1], [1, 2]]
```

翻转后是 `[[2, 1], [1, 1]]`。直接算翻转后的字符串 `BA`：

```text
B * A = [[1, 1], [0, 1]] * [[1, 0], [1, 1]] = [[2, 1], [1, 1]]
```

两边相同。从 `(1, 1)` 走：`AB` 得到 `(2, 3)`，`BA` 得到 `(3, 2)`，也分别等于 `[1, 1]` 乘这两个矩阵。

因此懒标记只存一个布尔值。区间被整段覆盖时，把结点矩阵按上面的规则翻转，并把标记传给两个儿子。

### 样例

`ABAAA` 的矩阵乘到 `[1, 1]` 上是 `[11, 3]`。`1 3 5` 翻转最后三个 `A`，字符串变成 `ABBBB`。最后一次查询从 `(0, 1000000000)` 出发，每一步都是 `B`，只把 `A` 加进 `B`，`A` 一直是 `0`。

### 复杂度

每次修改和查询都走 `O(log n)` 个结点，矩阵是 `2×2`。总时间 `O((n + q) log n)`。