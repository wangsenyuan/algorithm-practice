### ideas
1. 同一个类型必须是连续出现的
2. 对于i来说，如果不选 => 
3. 如果选i的话，有两种情况出现
4. 要么这个是第一个 s[i], 要么它和最后选择的s[j]相同
5. dp[i][state][x] = 到i为止最后一个选择是x时的计数
6. dp[i][state][x] = dp[i-1][state][x]
7. dp[i][state][x = (s[i])] += dp[i-1][state][x] + dp[i-1][state ^ (1 << x)][y]
