# Problem

Pavel is going to make a game of his dream. However, he knows that he can't make it on his own so he founded a development company and hired `n` workers. Now he wants to pick workers from the staff who will be directly responsible for developing the game.

Each worker has a certain skill level `vi`. Besides, each worker doesn't want to work with those whose skill is very different. In other words, the `i`-th worker won't work with those whose skill is less than `li`, and with those whose skill is more than `ri`.

Pavel understands that the game of his dream isn't too hard to develop, so a worker with any skill will be equally useful. That's why he wants to pick a team of the maximum possible size. Help him pick such a team.

## Input

The first line contains a single integer `n` (`1 ≤ n ≤ 10^5`) — the number of workers Pavel hired.

Each of the following `n` lines contains three space-separated integers `li`, `vi`, `ri` (`1 ≤ li ≤ vi ≤ ri ≤ 3 * 10^5`) — the minimum skill value of the workers that the `i`-th worker can work with, the `i`-th worker's skill, and the maximum skill value of the workers that the `i`-th worker can work with.

## Output

In the first line print a single integer `m` — the number of workers Pavel must pick for developing the game.

In the second line print `m` space-separated integers — the indices of selected workers in any order.

If there are multiple optimal solutions, print any of them.

## Examples

### Example 1

**Input**

```text
4
2 8 9
1 4 7
3 6 8
5 8 10
```

**Output**

```text
3
1 3 4
```

### Example 2

**Input**

```text
6
3 5 16
1 6 11
4 8 12
7 9 16
2 10 14
8 13 15
```

**Output**

```text
4
1 2 3 5
```


### Ideas
1. 假设v[i]是目前最高的技能，那么所有其他被选中的必须有r[?] >= v[i]
2. 假设这批人中，包括i，最低的技能是v[j], 那么所有l[?] <= v[j]
3. 在i给定的情况下，如何找到最优的j？
4. 找到满足r[?] >= v[i]条件的人很好办。（这个可以用一个heap来处理，如果r增加，不断的移除就好了）
5. 在这批人中，有些人可能v[?]很大，但是l[?]也很大（选择它，就必须放弃很多v比较小的人）
6. 有办法了，就是把它们看作(l, v)的一个区间，就是最大区间覆盖（l的时候，+1， v的时候-1）
7. got

### Solution Summary

**Problem essence:** Select the largest subset of workers such that every worker in the subset tolerates every other worker's skill level. Worker `i` tolerates worker `j` iff `l_i <= v_j <= r_i`. For a valid team, if the minimum skill is `v_min` and the maximum skill is `v_max`, then every member must satisfy `l_i <= v_min` and `r_i >= v_max`.

**Algorithm: Sweep + Segment Tree + Min-Heap**

1. **Sort** workers by skill `v_i` in ascending order.

2. **Sweep** through workers in increasing `v` order. Maintain a set of "active" workers — those whose upper tolerance `r_i` is still >= the current highest skill being considered. Use a **min-heap** on `r_i` to efficiently evict workers that can no longer coexist with the current worker (i.e., `r < v_current`).

3. **Segment tree with lazy propagation** over the range `[0, max_v]`. When a worker becomes active, do a **range +1 update** on `[l_i, v_i]`. When a worker is evicted, do a **range -1 update** on `[l_i, v_i]`. The value at position `x` in the segment tree counts how many active workers have `l_i <= x` and `v_i >= x`. This means: if we pick `x` as the minimum skill cutoff, exactly that many active workers are compatible. Since all active workers also satisfy `r_i >= v_current` (the current maximum skill), position `x`'s count is a valid team size.

4. The **global maximum** across all positions in the segment tree (tracked at the root) over the entire sweep gives the answer (maximum team size).

5. **Reconstruction (second pass):** Repeat the sweep. When the segment tree maximum first equals the best found in step 4, call `GetBest()` to find the position `x` (the rightmost point achieving the max, i.e., the largest valid `l` cutoff). Then collect all active workers in the heap whose interval `[l_i, v_i]` covers `x`.

**Complexity:** O(n log V) where V = max skill value (up to 3*10^5). The segment tree has size O(V) and each worker causes O(log V) work for updates. The heap operations contribute O(n log n).