# E. Ship's Shortest Path

[Problem link](https://codeforces.com/problemset/problem/75/E)

**Contest:** [Codeforces Beta Round 69 (Div. 1 Only)](https://codeforces.com/contest/75)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You have got a new job as a ship captain. Your first task is to move your ship
from one point to another at the minimum cost.

The shortest distance between any two points is the length of the line segment
between them. There is an island in the sea, so sometimes you cannot move the
ship along the segment between the two points.

You can only move to safe points. A point is safe if it lies on the line segment
between the start and the end, or if it lies on the island's edge.

Workers move the ship in the sea for `1` Egyptian pound per unit of distance,
and they can carry the ship across the island for `2` Egyptian pounds per unit
of distance. Moving along the island edge counts as moving in the sea.

The starting point is `(xStart, yStart)` and the end point is `(xEnd, yEnd)`.
The two points are distinct.

The island is a convex polygon. No three polygon vertices are collinear. The
start and the end are neither inside the island nor on its boundary. Polygon
vertices are given in anti-clockwise order.

## Constraints

- `-100 ≤ xStart, yStart, xEnd, yEnd ≤ 100`
- `3 ≤ n ≤ 30`
- `-100 ≤ x, y ≤ 100`
- Polygon vertices are distinct

## Input

The first line contains four integers `xStart`, `yStart`, `xEnd`, and `yEnd`.

The second line contains an integer `n` — the number of polygon vertices.

The third line contains `n` pairs of integers `x y` — the polygon vertices.

## Output

Print one line — the minimum possible cost. The absolute or relative error must
not exceed `10^{-6}`.

## Examples

### Input 1

```text
1 7 6 7
4
4 2 4 12 3 12 3 2
```

### Output 1

```text
6.000000000
```

### Input 2

```text
-1 0 2 0
4
0 0 1 0 1 1 0 1
```

### Output 2

```text
3.000000000
```

## Solution

船只能出现在两类点上：起点 `S` 到终点 `E` 的线段，以及岛的边界。海里（含岛的边界）每单位花费 `1`，穿过岛的内部每单位花费 `2`。岛是逆时针凸多边形，`S` 和 `E` 都在岛外。

因此不用在整张平面里搜索。线段和凸多边形的交要么是空的，要么是线段上的一段 `[A, B]`。最优路径只有三种：

1. 线段根本不进入内部：整段按海里走，花费 `|SE|`。
2. 线段穿过内部：岛外仍沿 `SA`、`BE` 走，中间 `AB` 按内部计价，花费 `|SE| - |AB| + 2|AB|`。
3. 同样走到入口 `A`、出口 `B`，但 `A` 到 `B` 改沿边界走。边界算海里。凸多边形把边界分成两段，取较短的那段。

三者取最小。前两种可以合成一句：把内部那段的长度从「按 1 计」换成「按 `min(2|AB|, 两条岸)` 计」。

### 1. 用半平面交出岛上的那一段

逆时针多边形的每条边 `u → v` 把平面切成半平面：点 `p` 在内侧或边上当且仅当 `cross(u, v, p) ≥ 0`。

把 `S`、`E` 写成 `S + t(E - S)`，`t ∈ [0, 1]`。叉积对 `t` 是线性的，所以每条边最多把可行区间切掉一头：

| `cross(S)` | `cross(E)` | 含义 |
|---|---|---|
| 都 `≥ 0` | 都 `≥ 0` | 整段都在这一侧，不切 |
| 都 `< 0` | 都 `< 0` | 整段在岛外，没有交，答案就是 `|SE|` |
| 一正一负 | | 交点 `t = cross(S) / (cross(S) - cross(E))`，用它收紧下界或上界 |

所有边切完后得到 `[lo, hi]`。`A`、`B` 就是这两个参数对应的点。`hi - lo` 约等于 `0` 表示只擦到一个点，没有进入内部，答案仍是 `|SE|`。

### 2. 贴着边走不算进入内部

`[A, B]` 落在同一条边上时，这段就是边界，按海里计价。样例 2 的直线正好贴着正方形的底边，交出来的 `[A, B]` 整段都在这条边上，所以答案是整段长度 `3`，而不是把中间那一截乘 `2`。

### 3. 穿过内部时比较穿岛和绕岸

否则 `AB` 是内部的弦。岛外两段的花费是 `|SE| - |AB|`，中间换成：

- 直接穿过去：`2|AB|`
- 从 `A` 沿边界逆时针走到 `B`
- 另一侧：周长减去上一段

`A` 或 `B` 落在顶点上时，把它算作「从该顶点出发的那条边」，这样沿边界走时不会在顶点处停错边。两条岸加起来正好是周长，所以只需要走一次。

### 样例 1

岛是 `x ∈ [3, 4]`、`y ∈ [2, 12]` 的细矩形，`S = (1, 7)`，`E = (6, 7)`。直线 `y = 7` 在 `(3, 7)` 进入、`(4, 7)` 离开，内部长度 `1`，整段长度 `5`。

两条岸都要绕到 `y = 2` 或 `y = 12`，长度都是 `11`。`min(2, 11, 11) = 2`，答案是 `5 - 1 + 2 = 6`。穿过去比绕岸便宜。

### 复杂度

`n ≤ 30`，切半平面和沿边界各走一圈都是 `O(n)`。
