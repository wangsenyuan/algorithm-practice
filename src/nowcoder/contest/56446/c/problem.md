# C. Gold Rong 的数组

[Problem link](https://ac.nowcoder.com/acm/contest/56446/C)

**Contest:** [第五届太原理工大学程序设计新生赛决赛（重现赛）](https://ac.nowcoder.com/acm/contest/56446)

time limit per test: 1 second

memory limit per test: 256 megabytes

## Problem

Gold Rong 有一个长度为 `n` 的数组，下标从 1 开始，初始全是 `0`。

进行 `q` 次修改。第 `i` 次给出位置 `x` 和正整数 `y`，把每个 `a[i]` 加上

```text
max(y - |x - i|, 0)
```

输出全部修改结束后的数组。

## Constraints

- `1 ≤ n, q ≤ 10^5`
- `1 ≤ x ≤ n`
- `1 ≤ y ≤ 10^6`
- 所有 `y` 的总和不超过 `10^6`

## Input

```text
n q
x_1 y_1
x_2 y_2
...
x_q y_q
```

## Output

输出一行 `n` 个整数，为最终数组。

## Sample Input 1

```text
5 2
2 3
4 1
```

## Sample Output 1

```text
2 3 2 2 0
```

第一次 `x = 2, y = 3` 得到 `2 3 2 1 0`。第二次 `x = 4, y = 1` 只给下标 4 加 1，得到 `2 3 2 2 0`。

## Sample Input 2

```text
6 3
2 6
2 5
2 4
```

## Sample Output 2

```text
12 15 12 9 6 3
```

三次修改的位置都是 2，高度分别是 6、5、4。

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
