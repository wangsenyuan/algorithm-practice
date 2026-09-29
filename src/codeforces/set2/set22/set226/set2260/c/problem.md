# C. Maximize XOR, Minimize Operations

[题目链接](https://codeforces.com/problemset/problem/2260/C)

**比赛：** [Educational Codeforces Round 194 (Rated for Div. 2)](https://codeforces.com/contest/2260)

**标签：** bitmasks, greedy, \*1300

time limit per test: 2 seconds

memory limit per test: 512 megabytes

## 题意

给定两个非负整数 `x` 和 `y`。一次操作可以把 `x` 减 `1`，同时把 `y` 加 `1`（当 `x = 0` 时不能操作）。

对每组初始值，做任意次（可以为 0）操作，使得 `x ⊕ y`（按位异或）尽量大；在所有能达到最大异或值的方案中，选择操作次数最少的一种。

输出最大的 `x ⊕ y`，以及达到它所需的最少操作次数。

## 约束

- 多测：第一行 `t`（`1 ≤ t ≤ 10^4`）
- 每组一行两个整数 `x y`（`0 ≤ x, y < 2^{29}`）

## 输入

第一行一个整数 `t` — 测试组数。

接下来 `t` 行，每行两个整数 `x` 和 `y`。

## 输出

对每组测试输出两个整数：最大的 `x ⊕ y`，以及最少操作次数。

## 样例

### 样例输入

```
3
3 1
0 5
6 4
```

### 样例输出

```
4 3
5 0
10 4
```

将官方样例拆成单组时：

| 输入 | 输出 |
| --- | --- |
| `3 1` | `4 3` |
| `0 5` | `5 0` |
| `6 4` | `10 4` |
