# Problem D: Road Repair

## Problem Statement

Everything is great about Ilya's city, except the roads. The thing is, the only ZooVille road is represented as n holes in a row. We will consider the holes numbered from 1 to n, from left to right.

Ilya is really keen on helping his city. So, he wants to fix at least k holes (perhaps he can fix more) on a single ZooVille road.

The city has m building companies, the i-th company needs ci money units to fix a road segment containing holes with numbers of at least li and at most ri. The companies in ZooVille are very greedy, so, if they fix a segment containing some already fixed holes, they do not decrease the price for fixing the segment.

Determine the minimum money Ilya will need to fix at least k holes.

## Input

The first line contains three integers n, m, k (1 ≤ n ≤ 300, 1 ≤ m ≤ 10^5, 1 ≤ k ≤ n). The next m lines contain the companies' description. The i-th line contains three integers li, ri, ci (1 ≤ li ≤ ri ≤ n, 1 ≤ ci ≤ 10^9).

## Output

Print a single integer — the minimum money Ilya needs to fix at least k holes.

If it is impossible to fix at least k holes, print -1.

**Note:** Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Examples

### Example 1
**Input:**
```
10 4 6
7 9 11
6 9 13
7 7 7
3 5 6
```

**Output:**
```
17
```

### Example 2
**Input:**
```
10 7 1
3 4 15
8 9 8
5 6 8
9 10 6
1 4 2
1 4 10
8 10 13
```

**Output:**
```
2
```

### Example 3
**Input:**
```
10 1 9
5 10 14
```

**Output:**
```
-1
```


### ideas
1. 假设在r处，选择供应商i, 那么这时候，有好几种情况
2. 在l前面有j个人
3. 有一部分重叠的，
4. 有一部分被包含的
5. 包含的这部分可以认为肯定不合适，可以直接不考虑
6. 那就考虑重叠和在前面的情况
7. 不好搞， dp[l][r] 表示，如果使用供应商，将l...r都修复的情况下的，最优解
8. dp[l][r] = dp[l][i] + dp[j][r] 如果存在一个供应商可以修理 l...i, j <= i + 1
9. 但是这样子的复杂性 = n * n * n + m (用m先去更新一波)
10. fp[r][j] = fp[r][i] + fp[r - i][j - i] done
11. 