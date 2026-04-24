# Problem

Let `n` and `h` be positive integers. Swampy the Alligator needs water for his bathtub, and you are tasked with getting it to him.

Above his house, there is an `h × n` grid of dirt and water tiles, with columns numbered `1..n`.

In column `i`, the bottom `ai` tiles are dirt, and the remaining top tiles are water.

You may place a drain on any non-dirt tile. Placing a drain removes all water tiles that can reach that drain by moving only **down**, **left**, or **right**, without crossing dirt tiles.

What is the maximum amount of water you can get to Swampy by placing at most **two** drains?

If a water tile can be removed by either drain, it is counted only once.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^3`) — the number of test cases.

For each test case:

- First line: two integers `n` and `h` (`1 ≤ n ≤ 2000`, `1 ≤ h ≤ 10^9`) — number of columns and grid height.
- Second line: `n` integers `a1, a2, ..., an` (`1 ≤ ai ≤ h`) — dirt height in each column.

It is guaranteed that the sum of `n` over all test cases does not exceed `2000`.

## Output

For each test case, print one integer — the maximum number of water tiles that can be drained using at most two drains.

## Example

**Input**

```text
6
7 4
1 3 1 2 3 1 1
8 10
7 5 1 3 2 5 6 8
1 1
1
10 20
5 2 1 2 1 3 6 7 1 1
5 1000000000
1 420420420 1 420420420 1
1 1000000000
1
```

**Output**

```text
14
43
0
170
3738738738
999999999
```

## Note

In the first test case, it is possible to drain `14` tiles of water by placing two drains as shown in the statement figure.

In the second test case, it is possible to drain all `43` water tiles.

In the third test case, the grid has only one dirt tile and no drain can be placed, so the answer is `0`.


### ideas
1. 考虑只有1个drain时的情况
2. 假设在某个i处，放置了一个drain，那么它可以抽取的水量
3. = 左边 + 右边。左边是一个递减的stack；右边是一个递增的stack
4. (h - h1) * w1 + (h - h2) * w2 + ... (h - hi) * wi
5. h * w1 + h * w2 + ... - h1 * w1 - h2 * w2... 
6. h * i - h1 * w1 - h2 * w2 ...
7. 单个的可以这样算; 
8. 考虑两个要怎么算呢？
9. 首先，可以判断的是，这两个里面，肯定有一个是最优，流出最大的那个
10. 假设最优答案不是最优的，那么通过交换成最优的一个，答案会更好
11. 这时候，它不可能包含第二个（因为包含第二个的时候，第二个就是最优的）
12. 然后，把被第一个抽干的部分，全部填满，然后再跑一次就可以了
13. 不大对～
14. 虽然做对了，但是不像1600分呐

## key insights

1. For one drain at column `i`, drained water equals:
   - sum of `(h - a[x])` over all columns `x` that can connect to `i` through moves left/right/down.
   - In 1D column terms, this is determined by nearest higher barriers on both sides.

2. The code precomputes directional contribution stacks via `play(a)`:
   - It scans left to right with a monotonic decreasing stack by height `a[idx]`.
   - For each position `i`, it computes cumulative drainable water from the left side.
   - `dp[i]` stores a snapshot of this stack, where stack top contains the best cumulative value up to `i`.

3. To get right-side symmetric information, it reuses the same routine on reversed array:
   - `fp = play(reverse(a))`, then reverse back.
   - So `top(dp[x]).second` and `top(fp[x]).second` act like one-drain totals seen from left/right preprocessing.

4. Then it enumerates pairs `(l, r)` of drain columns (`O(n^2)`):
   - start from sum of independent left/right totals,
   - identify `hi` = index of maximum `a` in `[l..r]`,
   - apply inclusion-exclusion style corrections to avoid double counting overlap between two drains,
   - special-case `l == r` as single-drain answer.

5. Why this passes limits:
   - `sum n <= 2000`, so `O(n^2)` per test case (plus linear preprocessing) is acceptable.
   - Stack operations in preprocessing are amortized linear.

6. Final answer is the maximum corrected value over all `(l, r)`.