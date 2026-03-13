# Problem

- **Given:** `n` roads with repair costs, and `m` interval races `[lb, ub]` each with profit `p`.
- **Constraint:** a race `[lb, ub]` can be held only if **all** roads `lb..ub` are repaired.
- **Profit:** total race rewards for all held races, minus total repair costs for all repaired roads.
- **Goal:** choose which roads to repair (and thus which races become feasible) to **maximize** net profit (or choose none and get `0`). 

---

## Input

- The first line contains two integers `n` and `m`  
  (`1 ≤ n, m ≤ 2·10^5`) — the number of roads and the number of races.

- The next `n` lines each contain a single non-negative integer `cᵢ`  
  (`0 ≤ cᵢ ≤ 10^9`) — the cost to repair road `i`, given in order from road `1` to road `n`.

- Then `m` lines follow, each describing one race by three integers `lb`, `ub`, and `p`:
  - `1 ≤ lb ≤ ub ≤ n`
  - `1 ≤ p ≤ 10^9`
  - this race uses all roads in `[lb, ub]` (inclusive) and pays `p` if held.

---

## Output

Print a single integer — the **maximum possible profit** you can gain.

---

## Examples

### Example 1

Input:

```text
7 4
3
2
3
2
1
2
3
1 2 5
2 3 5
3 5 3
7 7 5
```

Output:

```text
4
```

### Example 2

Input:

```text
2 1
0
3
1 2 5
```

Output:

```text
2
```

### Example 3

Input:

```text
3 1
10
10
10
1 3 10
```

Output:

```text
0
```

---

## Note

In the first sample, the optimal solution is to repair roads `1, 2, 3, 7`.  
Then three races are held, paying `5 + 5 + 5 = 15`,  
and repair cost is `3 + 2 + 3 + 3 = 11`, so the profit is `15 − 11 = 4`.

---

### ideas
1. dp[i] = 表示到road i时的最大收益
2. dp[i] = 如果区间 j...i 被修复了，dp[j-1] + 在区间[j..i]中能够举办的race的收益 - 修复[j...i]中间的花费
3. 这些j，必然是某个race的开始区间（修理更前面的，没有好处）
4. fp[j] 表示到i为止，j处的收益
5. dp[i] = dp[j-1] + fp[j] - (pref[i] - pref[j-1])
6.      =  dp[j-1] + fp[j] + pref[j-1] - pref[i]
7.   那么只要找到最大的j就可以了（区间更新）+ 区间查询