# Problem D: Mike and !Mike

## Problem Statement

Mike and !Mike are old childhood rivals, they are opposite in everything they do, except programming. Today they have a problem they cannot solve on their own, but together (with you) — who knows?

Every one of them has an integer sequences `a` and `b` of length `n`. Being given a query of the form of pair of integers `(l, r)`, Mike can instantly tell the value of `max(a[l], a[l+1], ..., a[r])` while !Mike can instantly tell the value of `min(b[l], b[l+1], ..., b[r])`.

Now suppose a robot (you!) asks them all possible different queries of pairs of integers `(l, r)` (`1 ≤ l ≤ r ≤ n`) (so he will make exactly `n(n + 1) / 2` queries) and counts how many times their answers coincide, thus for how many pairs `max(a[l], ..., a[r]) = min(b[l], ..., b[r])` is satisfied.

How many occasions will the robot count?

## Input

The first line contains only integer `n` (`1 ≤ n ≤ 200,000`).

The second line contains `n` integer numbers `a₁, a₂, ..., aₙ` (`-10⁹ ≤ aᵢ ≤ 10⁹`) — the sequence `a`.

The third line contains `n` integer numbers `b₁, b₂, ..., bₙ` (`-10⁹ ≤ bᵢ ≤ 10⁹`) — the sequence `b`.

## Output

Print the only integer number — the number of occasions the robot will count, thus for how many pairs `max(a[l], ..., a[r]) = min(b[l], ..., b[r])` is satisfied.

## Examples

### Example 1

**Input:**

```
6
1 2 3 2 1 4
6 7 1 2 3 2
```

**Output:**

```
2
```

### Example 2

**Input:**

```
3
3 3 3
1 1 1
```

**Output:**

```
0
```

## Note

The occasions in the first sample case are:

1. `l = 4, r = 4` since `max{2} = min{2}`
2. `l = 4, r = 5` since `max{2, 1} = min{2, 3}` which equals `2`

There are no occasions in the second sample case since Mike will answer `3` to any query pair, but !Mike will always answer `1`.


### ideas
1. 有意思～
2. 从小到大的迭代a，从大到小的迭代b
3. 假设a[i] = b[j], 且包好i的区间为[l1, r1], 包含j的区间是[l2, r2]
4. 那么这个区间的交集，l = max(l1, l2), r = min(r1, r2) 贡献 (min(i, j) - l1 + 1) * (r - max(i, j) + 1)
5. 但是还有一个情况是相同的a[i], 相同的a[j] 要怎么处理？
6. 不大对～
7. 先找出，a[i]作为最大值的[li, ri], ...
8. 然后处理b[j]作为最小值的区间, [l, r], 然后找出和它有重叠的区间
9. 区间更新和区间更新。貌似思路对了～
10. 先考虑，只有一个区间和b[j]重叠，那么上面的那个推导还是成立的
11. 但是如果有多个区间呢？如果这些区间不重叠，也还是ok的。如果有重叠要怎么算呢？
12. 似乎也还好～
13. 因为每个数a[i]和b[j],只会被使用一次，所以整体的时间复杂度，还是(n + m)?