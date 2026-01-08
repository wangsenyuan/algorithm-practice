# Problem Description

Polycarp is an experienced participant in Codehorses programming contests. Now he wants to become a problemsetter.

He sent to the coordinator a set of n problems. Each problem has it's quality, the quality of the i-th problem is ai (ai can be positive, negative or equal to zero). The problems are ordered by expected difficulty, but the difficulty is not related to the quality in any way. The easiest problem has index 1, the hardest problem has index n.

The coordinator's mood is equal to q now. After reading a problem, the mood changes by it's quality. It means that after the coordinator reads a problem with quality b, the value b is added to his mood. The coordinator always reads problems one by one from the easiest to the hardest, it's impossible to change the order of the problems.

If after reading some problem the coordinator's mood becomes negative, he immediately stops reading and rejects the problemset.

Polycarp wants to remove the minimum number of problems from his problemset to make the coordinator's mood non-negative at any moment of time. Polycarp is not sure about the current coordinator's mood, but he has m guesses "the current coordinator's mood q = bi".

For each of m guesses, find the minimum number of problems Polycarp needs to remove so that the coordinator's mood will always be greater or equal to 0 while he reads problems from the easiest of the remaining problems to the hardest.

## Input

The first line of input contains two integers n and m (1 ≤ n ≤ 750, 1 ≤ m ≤ 200 000) — the number of problems in the problemset and the number of guesses about the current coordinator's mood.

The second line of input contains n integers a1, a2, ..., an (-10^9 ≤ ai ≤ 10^9) — the qualities of the problems in order of increasing difficulty.

The third line of input contains m integers b1, b2, ..., bm (0 ≤ bi ≤ 10^15) — the guesses of the current coordinator's mood q.

## Output

Print m lines, in i-th line print single integer — the answer to the problem with q = bi.

## Example

### Example Input

```text
6 3
8 -5 -4 1 -7 4
0 7 3
```

### Example Output

```text
2
0
1
```


### ideas
1. 有意思，n比较小，m很大
2. 如果mood = x 只需要删除k个，那么 mood = x+1的时候，也最多需要删除k个
3. dp[k]表示最差的 mood = x
4. k越小, mood越大
5. 那么对于m来说，只需要二分就可以了， dp[k] >= q[i] 最小的k即可
6. 但是貌似太慢了
7. dp[i][j] 表示到i为止，删除j个，需要的最小的mood
8. dp[i][j] = dp[i][j-1] 删除越少，需要的mood越大
9. n * n * lg(inf) * lg(n)
10. dp[j][i] 表示目前队列里面有j个最小的值，且到达了i
11. dp[j][i] = -(sum(这j个数))