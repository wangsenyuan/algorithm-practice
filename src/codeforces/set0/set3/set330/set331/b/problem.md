# Beaver Shaving Problem

## Problem Description

The Smart Beaver has recently designed and built an innovative nanotechnologic all-purpose beaver mass shaving machine, "Beavershave 5000". Beavershave 5000 can shave beavers by families! How does it work? Very easily!

There are n beavers, each of them has a unique id from 1 to n. Consider a permutation a₁, a₂, ..., aₙ of n these beavers. Beavershave 5000 needs one session to shave beavers with ids from x to y (inclusive) if and only if there are such indices i₁ < i₂ < ... < iₖ, that aᵢ₁ = x, aᵢ₂ = x + 1, ..., aᵢₖ₋₁ = y - 1, aᵢₖ = y. And that is really convenient. For example, it needs one session to shave a permutation of beavers 1, 2, 3, ..., n.

If we can't shave beavers from x to y in one session, then we can split these beavers into groups [x, p₁], [p₁ + 1, p₂], ..., [pₘ + 1, y] (x ≤ p₁ < p₂ < ... < pₘ < y), in such a way that the machine can shave beavers in each group in one session. But then Beavershave 5000 needs m + 1 working sessions to shave beavers from x to y.

All beavers are restless and they keep trying to swap. So if we consider the problem more formally, we can consider queries of two types:

1. What is the minimum number of sessions that Beavershave 5000 needs to shave beavers with ids from x to y, inclusive?
2. Two beavers on positions x and y (the beavers aₓ and aᵧ) swapped.

You can assume that any beaver can be shaved any number of times.

## Input

- The first line contains integer n — the total number of beavers, 2 ≤ n
- The second line contains n space-separated integers — the initial beaver permutation
- The third line contains integer q — the number of queries, 1 ≤ q ≤ 10⁵
- The next q lines contain the queries. Each query i looks as pᵢ xᵢ yᵢ, where:
  - pᵢ is the query type (1 is to shave beavers from xᵢ to yᵢ, inclusive, 2 is to swap beavers on positions xᵢ and yᵢ)
  - All queries meet the condition: 1 ≤ xᵢ < yᵢ ≤ n

### Constraints
- **Subproblem B1 (30 points)**: n ≤ 100
- **Subproblem B2 (100 points)**: n ≤ 3·10⁵
- Note that the number of queries q is limited 1 ≤ q ≤ 10⁵ in both subproblem B1 and subproblem B2.

## Output

For each query with pᵢ = 1, print the minimum number of Beavershave 5000 sessions.

## Examples

### Input
```
5
1 3 4 2 5
6
1 1 5
1 3 4
2 2 3
1 1 5
2 1 5
1 1 5
```

### Output
```
2
1
3
5
```

### ideas
1. 先不考虑swap的情况， 对于区间[x...y] (不是位置，是数字)，如果, 假设x,x+1...y
2. 的位置位 pos[x], pos[x+1], ... pos[y]
3. 如果 pos[x] < pos[x+1] ... < pos[y]那么就可以在一次删除掉
4. 特别的如果 [x, x+1], 如果 pos[x] < pox[x+1], 那么它们就能一次处理
5. 如果 pos[x] > pos[x+1], 那么它们只能分布到两批次
6. 如果 pos[...] 中，有 pos[i] > pos[i+1], 那么就产生了额外的一次
7. 那么 [x...y]的处理次数 = pos[x]...pos[y]里面的次数
8. 排列 3, 1, 2 需要两次， 3， 2， 1需要3次
9. 