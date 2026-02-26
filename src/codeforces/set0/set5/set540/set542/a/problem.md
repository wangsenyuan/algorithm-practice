# Problem

Ivan Anatolyevich's agency has made $n$ TV commercial videos. The $i$-th video can only be shown during time segment $[l_i, r_i]$ (the broadcast window must lie inside this segment).

There are $m$ TV channels. The $j$-th channel has $c_j$ viewers and sells time $[a_j, b_j]$ for the commercial.

Ivan must choose:

- exactly one video $i$ and one channel $j$,
- a time segment $[x, y]$ such that $[x, y] \subseteq [l_i, r_i]$ and $[x, y] \subseteq [a_j, b_j]$.

**Efficiency** of the broadcast is $(y - x) \cdot c_j$ — total viewer-time. Find a choice that **maximizes** efficiency.

---

## Input

- Line 1: two integers $n$, $m$ ($1 \le n, m \le 2 \cdot 10^5$) — number of videos and channels.
- Next $n$ lines: two integers $l_i$, $r_i$ ($0 \le l_i \le r_i \le 10^9$) — allowed time segment for the $i$-th video.
- Next $m$ lines: three integers $a_j$, $b_j$, $c_j$ ($0 \le a_j \le b_j \le 10^9$, $1 \le c_j \le 10^9$) — time slot and viewer count for the $j$-th channel.

## Output

- Line 1: maximum possible efficiency. If no valid choice gives strictly positive efficiency, print $0$.
- If the maximum is strictly positive: line 2 — two integers $i$, $j$ ($1 \le i \le n$, $1 \le j \le m$) — the video and channel in an optimal broadcast. If multiple answers exist, output any.

---

## Examples

**Input**

```
2 3
7 9
1 4
2 8 2
0 4 1
8 9 3
```

**Output**

```
4
2 1
```

**Input**

```
1 1
0 0
1 1 10
```

**Output**

```
0
```

---

## Note

- **Sample 1:** Optimal is video $2$, channel $1$, time $[2, 4]$. Efficiency $= (4 - 2) \cdot 2 = 4$.
- **Sample 2:** Video segment $[0, 0]$ and channel segment $[1, 1]$ do not intersect, so the answer is $0$.


### ideas
1. 对于一个show i和一个channel j
2. 那么 x = max(l[i], a[j]), y = min(r[i], b[j]) (重叠的区间最大化)
3. 假设对于某个show i, 可以将channel分成了3类
4. 第一类是，a[j] >= r[i]的部分, 或者b[j] <= l[i] （无法采用)
5. 第二类是完全覆盖i的区间,(这个应该选择最大的c[j])
6. 还有就是覆盖一部分（可能在两头）
7. 考虑 l[i] < a[j] 的部分 (r[i] - a[j]) * c[j]
8. 要用到凸优化的?
9. 如果是a[j] < l[i] < b[j]的部分 （b[j] - l[i]) * c[i] 
10.  换个角度，对于任意一个channel，找到和它重叠区间最大的show
11.  处理的频道j的结束节点时（如果存在某个show没有结束）那么需要找到最早的那个show（l[i]最小的）
12.  这个可以用heap维护
13.  但是对于那些已经结束的show (r[i] <= b[j]) 的，应该找在区间a[j]后面，区间内的最大值
14.  但是还有一些是在a[j]前开始，a[j]后结束的，那么这些应该是在a[j]的时候，找到最大的r[i]