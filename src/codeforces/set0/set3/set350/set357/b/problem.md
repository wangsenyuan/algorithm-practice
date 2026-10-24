# B. Flag Day

[Problem link](https://codeforces.com/problemset/problem/357/B)

**Contest:** [Codeforces Round 207 (Div. 1)](https://codeforces.com/contest/357)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` dancers and `m` dances. Each dance uses exactly three dancers, and those three must wear three different colors: `1` white, `2` red, and `3` blue. A dancer keeps the same color for every dance.

The dances are given in order. Each dance shares at most one dancer with all earlier dances, and every dancer appears in at least one dance. A valid coloring exists. Print any one.

## Constraints

- `3 ≤ n ≤ 10^5`
- `1 ≤ m ≤ 10^5`
- Each dance lists three distinct dancer indices from `1` to `n`

## Input

```text
n m
a_1 b_1 c_1
a_2 b_2 c_2
...
a_m b_m c_m
```

## Output

Print `n` integers, the color of dancers `1` through `n`.

## Sample Input 1

```text
7 3
1 2 3
1 4 5
4 6 7
```

## Sample Output 1

```text
1 2 3 3 2 2 1
```

## Sample Input 2

```text
9 3
3 6 9
2 5 8
1 4 7
```

## Sample Output 2

```text
1 1 1 2 2 2 3 3 3
```

## Sample Input 3

```text
5 2
4 1 5
3 1 2
```

## Sample Output 3

```text
2 3 1 1 3
```

## Solution

舞会按出场顺序处理。每一场里最多只有一个舞者在更早的舞会里出现过，所以这一场最多只有一个人已经有颜色。

三个人都还没染色时，直接染成 `1, 2, 3`。否则把已经染色的那个人换到三个位置里的第一个，另外两个人染上剩下的两种颜色。已有颜色是 `c` 时，第二个人取 `c + 1`（`c = 3` 时回到 `1`），第三个人取 `6 - c - 第二个人的颜色`，也就是 `{1, 2, 3}` 里剩下的那个。已经染过的人不再改色，所以每一场都恰好用遍三种颜色。

### 样例 1

`(1, 2, 3)` 都是新的，染成 `1, 2, 3`。`(1, 4, 5)` 里只有 `1` 已经是 `1`，于是 `4, 5` 染成 `2, 3`。`(4, 6, 7)` 里只有 `4` 已经是 `2`，于是 `6, 7` 染成 `3, 1`。最终是 `1 2 3 2 3 3 1`。官方样例 `1 2 3 3 2 2 1` 也合法，任意一种都可以。

### 复杂度

每场舞会只看三个人。`n, m ≤ 10^5`，总时间 `O(n + m)`。
