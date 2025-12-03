## Problem Statement

Sereja has two sequences `a1, a2, ..., an` and `b1, b2, ..., bm`, consisting of integers. One day Sereja got bored and decided to play with them.

### Game Rules

Sereja makes several moves. In one move, he can perform one of the following actions:

1. **Action 1 (Match and Remove):**
   - Choose several (at least one) first elements of sequence `a` (non-empty prefix of `a`)
   - Choose several (at least one) first elements of sequence `b` (non-empty prefix of `b`)
   - The element of sequence `a` with the maximum index among the chosen ones must be equal to the element of sequence `b` with the maximum index among the chosen ones
   - Remove the chosen elements from the sequences
   - **Cost:** `e` energy units
   - **Reward:** Adds one dollar to Sereja's electronic account

2. **Action 2 (End Game):**
   - Remove all elements of both sequences
   - **Cost:** Number of energy units equal to the number of elements Sereja removed from the sequences before performing this action
   - **Reward:** Sereja gets all the money earned on his electronic account during the game

### Constraints

- Initially Sereja has `s` energy units and no money on his account
- The amount of Sereja's energy mustn't be negative at any time
- **Goal:** Find the maximum number of dollars Sereja can earn

## Input

The first line contains integers `n`, `m`, `s`, `e` (`1 ≤ n, m ≤ 10^5`; `1 ≤ s ≤ 3·10^5`; `10^3 ≤ e ≤ 10^4`).

The second line contains `n` integers `a1, a2, ..., an` (`1 ≤ ai ≤ 10^5`).

The third line contains `m` integers `b1, b2, ..., bm` (`1 ≤ bi ≤ 10^5`).

## Output

Print a single integer — the maximum number of dollars that Sereja can get.

## Examples

### Example 1
**Input:**
```
5 5 100000 1000
1 2 3 4 5
3 2 4 5 1
```

**Output:**
```
3
```

### Example 2
**Input:**
```
3 4 3006 1000
1 2 3
1 2 4 3
```

**Output:**
```
2
```


### ideas
1. 第二步必须是最后一步，在第二步前如果删除了w个(a+b)， 那么就需要消耗w的能量
2. dp[i][k] 表示最后一次删除是i，且得分为k时，删除的b的最小的下标
3. let j1 = dp[i1][k-1]
4.  dp[i][k] = j (j1 < j 且a[i] = b[j])
5.  还有一个点，就是这个k的大小，不影响最后一步，越大越好
6.  最后一步只被(i + j)影响到
7.  省略掉其他的数的话，就是找到了(a, b)中间最长的两个子序列，且它们是相等的