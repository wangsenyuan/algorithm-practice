# B. Karen and Test

[Problem link](https://codeforces.com/problemset/problem/815/B)

**Contest:** [Codeforces Round 419 (Div. 1)](https://codeforces.com/contest/815)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

There are `n` integers in a row. Build the next row from each pair of adjacent values, alternating addition and subtraction. The first operation is addition.

If a row ends with addition, the next row starts with subtraction. If it ends with subtraction, the next row starts with addition.

Repeat until one integer remains. Print that integer modulo `10^9+7`, as a non-negative remainder.

## Constraints

- `1 ≤ n ≤ 200000`
- `1 ≤ a_i ≤ 10^9`

## Input

```text
n
a_1 a_2 ... a_n
```

## Output

Print one integer, the final value modulo `10^9+7`.

## Sample Input 1

```text
5
3 6 9 12 15
```

## Sample Output 1

```text
36
```

The rows are `3 6 9 12 15`, then `9 -3 21 -3`, then `6 -24 18`, then `30 -6`, then `36`.

## Sample Input 2

```text
4
3 7 5 2
```

## Sample Output 2

```text
1000000006
```

The final value is `-1`. Its non-negative remainder modulo `10^9+7` is `10^9+6`.

## Solution

答案对每个 `a[i]` 都是线性的。`n = 1` 时答案就是 `a[0]`。`n` 是奇数时先做掉这一行，后面只处理偶数长度。

奇数长度有偶数次运算，最后一次是减法，所以下一行以加法开头，长度 `n-1` 是偶数：

```text
b[i] = a[i] + (-1)^i * a[i+1]
```

### 每两行消掉一列

偶数长度、并且这一行以 `+` 开头时：

```text
b[i] = a[i] + (-1)^i * a[i+1]
```

这一行有奇数次运算，最后一次是 `+`，下一行以 `-` 开头：

```text
c[j] = b[j] - (-1)^j * b[j+1]
```

把 `b` 代进去，`a[j+1]` 的系数是 `(-1)^j + (-1)^{j+1} = 0`，于是

```text
c[j] = a[j] + a[j+2]
```

```text
a0  a1  a2  a3  a4  a5
│       │       │
└── + ──┘── + ──┘      偶数下标
    │       │       │
    └── + ──┘── + ──┘  奇数下标
```

以 `-` 开头再做两行，中间项同样抵消，得到的还是 `a[j] + a[j+2]`，并且下一行回到以 `+` 开头。所以每两行，奇偶两组各自做一次相邻相加。

一组有 `m = n/2` 个数，这样的加法要做 `m-1` 次。全是加法时，第 `i` 个数（从 0 开始）的系数是 `C(m-1, i)`。这就是 `play`：

```text
x = Σ C(m-1, i) * a[2i]
y = Σ C(m-1, i) * a[2i+1]
```

### 最后一次运算

做到只剩两个数时，开头的符号一共翻转了 `n/2 - 1` 次。

- `n % 4 == 2`：翻转偶数次，最后仍是加法，答案是 `x + y`
- `n % 4 == 0`：翻转奇数次，最后是减法，答案是 `x - y`

减法会产生负数，输出前折回 `[0, 10^9+7)`。

### 样例

样例 1 先做掉奇数行，得到 `9 -3 21 -3`，此时 `n = 4`。

```text
x = C(1,0)*9 + C(1,1)*21 = 30
y = C(1,0)*(-3) + C(1,1)*(-3) = -6
30 - (-6) = 36
```

样例 2 的 `n = 4`：

```text
x = 3 + 5 = 8
y = 7 + 2 = 9
8 - 9 = -1 ≡ 1000000006
```

### 复杂度

两组长度都是 `O(n)`，阶乘和逆元各算一遍。总时间 `O(n)`。
