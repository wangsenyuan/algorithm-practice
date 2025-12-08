Lunar New Year is approaching, and Bob is going to receive some red envelopes with countless money! But collecting money from red envelopes is a time-consuming process itself.

Let's describe this problem in a mathematical way. Consider a timeline from time 1 to n. The i-th red envelope will be available from time s_i to t_i, inclusive, and contain w_i coins. If Bob chooses to collect the coins in the i-th red envelope, he can do it only in an integer point of time between s_i and t_i, inclusive, and he can't collect any more envelopes until time d_i (inclusive) after that. Here s_i ≤ t_i ≤ d_i holds.

Bob is a greedy man, he collects coins greedily — whenever he can collect coins at some integer time x, he collects the available red envelope with the maximum number of coins. If there are multiple envelopes with the same maximum number of coins, Bob would choose the one whose parameter d is the largest. If there are still multiple choices, Bob will choose one from them randomly.

However, Alice — his daughter — doesn't want her father to get too many coins. She could disturb Bob at no more than m integer time moments. If Alice decides to disturb Bob at time x, he could not do anything at time x and resumes his usual strategy at the time x+1 (inclusive), which may lead to missing some red envelopes.

Calculate the minimum number of coins Bob would get if Alice disturbs him optimally.

## Input

The first line contains three non-negative integers n, m and k (1 ≤ n ≤ 10^5, 0 ≤ m ≤ 200, 1 ≤ k ≤ 10^5), denoting the length of the timeline, the number of times Alice can disturb Bob and the total number of red envelopes, respectively.

The following k lines describe those k red envelopes. The i-th line contains four positive integers s_i, t_i, d_i and w_i (1 ≤ s_i ≤ t_i ≤ d_i ≤ n, 1 ≤ w_i ≤ 10^9) — the time segment when the i-th envelope is available, the time moment Bob can continue collecting after collecting the i-th envelope, and the number of coins in this envelope, respectively.

## Output

Output one integer — the minimum number of coins Bob would get if Alice disturbs him optimally.

## Examples

### Example 1

**Input:**
```
5 0 2
1 3 4 5
2 5 5 8
```

**Output:**
```
13
```

### Example 2

**Input:**
```
10 1 6
1 1 2 4
2 2 6 2
3 3 3 3
4 4 4 5
5 5 5 7
6 6 6 9
```

**Output:**
```
2
```

### Example 3

**Input:**
```
12 2 6
1 5 5 4
4 6 6 2
3 8 8 3
2 9 9 5
6 10 10 7
8 12 12 9
```

**Output:**
```
11
```

## Note

In the first sample, Alice has no chance to disturb Bob. Therefore Bob will collect the coins in the red envelopes at time 1 and 5, collecting 13 coins in total.

In the second sample, Alice should disturb Bob at time 1. Therefore Bob skips the first envelope, collects the second one and can not do anything after that. So the answer is 2.


### ideas
1. dp[i][j] 表示在时刻i，有j次干扰后的最小值
2. dp[i][j] = dp[i-1][j] + gain at this monment
3.    or dp[i-1][j-1]
4. 但是这里的问题是，前面的选择，会影响到当前可以获得的红包
5. 如果take某个红包，那么在时刻d[i]前，他都不能再take其他的红包。
6. 好复杂～～～
7. 要倒过来考虑吗？
8. 假设在时刻i，按照贪心策略获得最大红包是x, dp[i][j] = w[x] + dp[d[x] + 1][j] (在这段时间内，进行干扰不划算)
9. 或者进行干扰 dp[i][j] = dp[i+1][j-1] ? 问题出来了，因为在时刻i+1, 可能最优的还是x
10. 所以，这时候, dp[i][j] = w[x] + dp[d[x] + 1][j-1]
11. 为了让这次干扰有效， dp[i][j] = dp[t[x] + 1][j-(t[x] - i + 1)] (这样子才可以)
12. 也就是说，在接下来的时间内，都必须一直干扰，直到时刻t[x]
13. 现在的选择，会对未来产生影响。就不好搞了啊
14. dp[i][x] 表示到i为止，且干扰了以x表示的红包集合的最优解，那么确实是可以搞的
15. 就贪心的选择当时可用的最好的红包就可以了
