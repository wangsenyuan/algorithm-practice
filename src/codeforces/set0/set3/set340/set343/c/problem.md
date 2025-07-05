# Problem: Mad Scientist Mike's Hard Drive

Mad scientist Mike does not use slow hard disks. His modification of a hard drive has not one, but **n** different heads that can read data in parallel.

When viewed from the side, Mike's hard drive is an endless array of tracks. The tracks of the array are numbered from left to right with integers, starting with 1. In the initial state, the i-th reading head is above the track number $h_i$. For each of the reading heads, the hard drive's firmware can move the head exactly one track to the right or to the left, or leave it on the current track. During the operation, each head's movement does not affect the movement of the other heads: the heads can change their relative order; there can be multiple reading heads above any of the tracks. A track is considered read if at least one head has visited this track. In particular, all of the tracks numbered $h_1, h_2, \ldots, h_n$ have been read at the beginning of the operation.

Mike needs to read the data on **m** distinct tracks with numbers $p_1, p_2, \ldots, p_m$. Determine the minimum time the hard drive firmware needs to move the heads and read all the given tracks. Note that an arbitrary number of other tracks can also be read.

---

## Input
- The first line of the input contains two space-separated integers $n$, $m$ ($1 \leq n, m \leq 10^5$) — the number of disk heads and the number of tracks to read, accordingly.
- The second line contains $n$ **distinct** integers $h_i$ in ascending order ($1 \leq h_i \leq 10^{10}$, $h_i < h_{i+1}$) — the initial positions of the heads.
- The third line contains $m$ **distinct** integers $p_i$ in ascending order ($1 \leq p_i \leq 10^{10}$, $p_i < p_{i+1}$) — the numbers of tracks to read.

---

**Note:**
- Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is recommended to use the `cin`, `cout` streams or the `%I64d` specifier.

### ideas
1. 对于某一个p[i], 它要么被它左边最近的h[j], 要么被它右边最近的h[j+1]处理
2. 假设依次处理完第i个head后，（它前面的track都被处理了）的最优解是f(i)
3. 然后考虑i+1, 那么在i...i+1中间的最优解，要怎么处理呢？
4. 如果不考虑回头的问题，那么就是找到中间相邻的两个track，一个分配给i，一个分配给i+1
5. 那么增加的移动次数 = p[x] - h[i] + h[i+1] - p[x+1] = (h[i+1] - h[i]) - (p[x+1] - p[x])
6. 貌似是选择相邻离的最远的两个（最好）
7. 但是考虑到要回头，就有点麻烦了
8. 可以二分
9. 限定单个head处理的上限，只要能处理，就分配给它