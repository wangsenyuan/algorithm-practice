# C. Levels and Regions

[Problem link](https://codeforces.com/problemset/problem/643/C)

**Contest:** [Codeforces Round 348 (VK Cup 2016 Round 2, Div. 1 Edition)](https://codeforces.com/contest/643)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` levels, divided into `k` regions. Each region is a non-empty contiguous segment of levels, and the regions partition `1..n`.

The game repeats the following process:

1. If every region is finished, the game ends. Otherwise let `X` be the first region that still has an unbeaten level.
2. Build a bag of tokens. For every already beaten level `i` in `X`, put `t_i` tokens of level `i` into the bag. Let `j` be the first unbeaten level in `X`, and put `t_j` tokens of level `j` into the bag.
3. Draw one token uniformly at random and play that level. This always takes one hour and counts as beating that level once, even if it was already beaten.

Print the minimum possible expected number of hours, over all ways to split the levels into `k` regions.

## Constraints

- `1 ≤ n ≤ 200000`
- `1 ≤ k ≤ min(50, n)`
- `1 ≤ t_i ≤ 100000`

## Input

```text
n k
t_1 t_2 ... t_n
```

## Output

Print one real number. An absolute or relative error of at most `10^-4` is accepted: if the answer is `a` and the jury answer is `b`, then `|a - b| ≤ 10^-4 · max(1, |b|)` is enough.

## Sample Input 1

```text
4 2
100 3 5 7
```

## Sample Output 1

```text
5.7428571429
```

The first region contains only level 1. The second region contains levels 2, 3, and 4.

## Sample Input 2

```text
6 2
1 2 4 8 16 32
```

## Sample Output 2

```text
8.5000000000
```

Split the levels into two regions of three levels each.

## Solution

区域从左到右打。当前区域还有没打通的关时，袋子里只有这个区域的代币，所以总期望是各段期望之和。

一段 `t[l..r]` 里，关卡必须按顺序第一次打通。打第 `j` 关时，袋子大小是这段自己的前缀和 `P[j] = t[l] + … + t[j]`，其中 `t[j]` 个是当前关。每次抽取花 1 小时，成功概率是 `t[j] / P[j]`，这一步的期望是 `P[j] / t[j]`。整段期望是

```text
cost(l, r) = Σ P[j] / t[j]    （j 从 l 到 r）
```

`t[j]` 是代币个数，不是小时数。

### 前缀和

记全局前缀和 `s[i] = t[1] + … + t[i]`。区间 `(p, i]`（上一段停在 `p`）里，第 `j` 关的段内前缀和是 `s[j] - s[p]`，所以

```text
cost(p+1, i) = Σ (s[j] - s[p]) / t[j]     （j 从 p+1 到 i）
```

`s[p]` 与 `j` 无关，把它从求和里提出来：

```text
Σ (s[j] - s[p]) / t[j]
  = Σ s[j] / t[j]  -  s[p] · Σ 1 / t[j]
```

两段求和都只覆盖 `j = p+1 .. i`，用从 1 开始的前缀相减：

```text
Σ s[j] / t[j] = a[i] - a[p]
Σ 1 / t[j]    = b[i] - b[p]
```

因此

```text
cost(p+1, i) = (a[i] - a[p]) - s[p] · (b[i] - b[p])
```

其中

```text
a[i] = Σ s[j] / t[j]
b[i] = Σ 1 / t[j]
```

`j` 都从 1 累到 `i`。一段的代价可以 `O(1)` 算出来。`b[i]` 严格递增，`s[i]` 也严格递增。

### 分段 DP

`fp[i][j]` 是前 `i` 关切成 `j` 段的最小期望。

```text
fp[0][0] = 0
fp[i][j] = min over p < i of  fp[p][j-1] + cost(p+1, i)
```

`j` 段至少要 `j` 个关卡，所以 `i ≥ j`，且 `p ≥ j - 1`。后面还要留 `k - j` 段，每段至少一个关卡，因此 `i ≤ n - (k - j)`。答案是 `fp[n][k]`。

直接枚举 `p` 是 `O(n²k)`。`n ≤ 2·10⁵`、`k ≤ 50`，过不了。实现里只留一层 `fp[i]`：第一层直接等于 `a[i]`（整段 `[1, i]`），之后每多一段就原地覆盖。读 `fp[j]` 时它还是上一层的值，算完再写回这一层。

### 斜率优化

把 `cost(p+1, i)` 代进转移：

```text
fp[i] = min over p of  fp[p] + (a[i] - a[p]) - s[p] · (b[i] - b[p])
```

`a[i]` 与 `p` 无关，提到 min 外面。剩下和 `p` 有关的项收成 `y[p]`：

```text
fp[p] - a[p] - s[p] · b[i] + s[p] · b[p]
  = (fp[p] - a[p] + s[p] · b[p]) - b[i] · s[p]
  = y[p] - b[i] · s[p]
```

所以

```text
fp[i] = a[i] + min over p of ( y[p] - b[i] · s[p] )
y[p] = fp[p] - a[p] + s[p] · b[p]
```

每个决策是平面上的点 `(x, y) = (s[p], y[p])`。`y - b·x` 是过该点、斜率为 `b` 的直线在 y 轴上的截距：

```text
y = b · x + c    ⇒    c = y - b · x
```

最小化 `c`，就是用斜率固定为 `b[i]` 的直线从下方去切这些点，切到的那个点就是最优的 `p`。切点一定落在下凸壳上，壳内部或壳上方的点不可能更优。

```text
y
│            D
│        C
│    B           E
│ A
└────────────────── x = s[p]
  下凸壳 A—B—E
  C、D 在壳上方，加入时会被弹掉

查询时用斜率 b[i] 的直线从下面顶住凸壳：

y
│         E
│     B ／
│    ／
│ A／
└────────────────── x
  切到 B。b[i] 变大以后切点往右移，下一刀会切到 E
```

`p` 从左到右加入，`s[p]` 递增，新点总在已有点的右边，凸壳用一个队列维护。

加入新点 `C` 时看队尾的两个点 `A、B`。`A→B→C` 右转或共线，说明 `B` 凹进壳里或落在 `AC` 上，弹出 `B`。叉积 `(B-A) × (C-B) > 0` 才是左转，`B` 留在下凸壳上。

```text
右转，弹出 B                左转，保留 B

      B                         A       C
     / \                         \     /
    A   C                         \   /
                                   \ /
                                    B
```

查询斜率 `b[i]` 随 `i` 增大，而壳上相邻边的斜率从左到右也递增，所以切点只会往右走。队头两个点 `P、Q`（`P` 在左）里，边 `PQ` 的斜率不超过 `b[i]` 时，`Q` 的截距已经不比 `P` 差：

```text
(y[Q] - y[P]) / (s[Q] - s[P])  ≤  b[i]
⇔  y[Q] - b[i] · s[Q]  ≤  y[P] - b[i] · s[P]
```

这时弹出 `P`。每个点最多进队一次、出队一次。

队头先放哨兵 `(0, 1e99)`。这一层真正的决策要从 `p = j-1` 才开始有，哨兵的纵坐标比任何真实点都大，一旦有真实点就会被弹掉。`j` 的循环范围是 `seg-1 .. n-(k-seg)`，正好是这一层合法的右端点。

### 样例 1

`100 | 3 5 7`。

第一段只有 `100`，期望 `100/100 = 1`。

第二段：

```text
3/3 + (3+5)/5 + (3+5+7)/7 = 1 + 1.6 + 15/7 = 4.742857…
```

总和 `5.7428571429`。

### 复杂度

预处理 `O(n)`。每一层凸壳上每个点最多进出一次，`k` 层一共 `O(nk)`。
