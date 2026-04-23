# Problem

Eulampius has created a game with the following rules:

- There are two players: a human and a computer.
- The game lasts for at most `n` rounds. Initially both players have `0` points. In the `j`-th round the human gains `aj` points and the computer gains `bj` points simultaneously.
- The game ends when one player reaches `k` or more points. That player **loses**. If both reach `k` or more points at the same time, **both** lose.
- If after `n` rounds both still have fewer than `k` points, the game ends in a **tie**.
- After each round the human may press the **Reset** button. If before reset the human had `x` points and the computer had `y` points (with `x < k` and `y < k`), then after reset:
  - human has `x' = max(0, x - y)`
  - computer has `y' = max(0, y - x)`
- For example, reset maps `(x = 3, y = 5)` to `(x' = 0, y' = 2)`, and `(x = 8, y = 2)` to `(x' = 6, y' = 0)`.

The sequences `a` and `b` are fixed in advance (reset does not change future `aj`, `bj`).

Polycarpus wants a plan to **win** while pressing Reset **as few times as possible**. Find that minimum number of presses, or report that winning is impossible.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

Each test case:

- First line: two integers `n` and `k` (`1 ≤ n ≤ 2 * 10^5`, `2 ≤ k ≤ 10^9`) — max rounds and the losing threshold.
- Second line: `n` integers `a1, a2, ..., an` (`1 ≤ aj < k`) — human points per round.
- Third line: `n` integers `b1, b2, ..., bn` (`1 ≤ bj < k`) — computer points per round.

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

Print answers for all test cases in order.

If Polycarpus cannot win, print a single line:

`-1`

(and nothing else for that test case).

Otherwise:

- First line: integer `d` — minimum number of Reset presses.
- Second line (optional if `d = 0`): `d` distinct integers `r1, r2, ..., rd` (`1 ≤ ri < n`) — round indices **after which** Reset must be pressed, in any order.

If `d = 0`, you may omit the second line or leave it empty.

If multiple optimal solutions exist, print any.

## Example

**Input**

```text
3
4 17
1 3 5 7
3 5 7 9
11 17
5 2 8 2 4 6 1 2 7 2 5
4 6 3 3 5 1 7 4 2 5 3
6 17
6 1 2 7 2 5
1 7 4 2 5 3
```

**Output**

```text
0

2
2 4
-1
```

## Note

In the second test case, if the human presses Reset after rounds `2` and `4`, the game proceeds as follows:

1. After round 1: human `5`, computer `4`.
2. After round 2: human `7`, computer `10`.
3. Reset → human `0`, computer `3`.
4. After round 3: human `8`, computer `6`.
5. After round 4: human `10`, computer `9`.
6. Reset → human `1`, computer `0`.
7. After round 5: human `5`, computer `5`.
8. After round 6: human `11`, computer `6`.
9. After round 7: human `12`, computer `13`.
10. After round 8: human `14`, computer `17`.

The computer reaches at least `k` points while the human stays strictly below `k`, so the human wins.


### ideas
1. 注意到，pref[a[:i]] - pref[b[:i]] 无论操作多少次，是不会变的
2. 假设在i处进行了一次reset，那么此时分数，就等于 (0, diff), 或者 (diff, 0)
3. 如果在i处进行了一次reset，必须找到一个位置j < i, 也进行了一次reset，且在这个过程中，不能出现超过pref[a[:i-1]] > k, 或者 pref[b[:i-1]] > k
4. 在所有的位置，dp[i] = 在位置i处reset时的少次数（能安全到达它这里）
5. dp[i] = dp[j] + 1 where j 到i对于双方都是安全的，且dp[j]最小
6. score_b[j] 表示在j处reset后，computer的得分(0, 或者 diff)
7. 考虑在位置i处，human获胜，那么此时必须存在一个j, pref_b[i] - pref_b[j] + score_b[j] >= k
8. 且dp[j]要最小（但是必须保证从j到i，human不会落败）
9. 或者dp[j]最小的时候，需要 score_b[j] - pref_b[j]越大越好
10. 要保证human不落败，那么pref_a[i] - pref_a[j] + score_a[j] < k
11. 假设在i处reset了
12. 保证它不失败，然后检查computer是否会失败， 如果可以，那么它就是最优解？
13. 所以，需要知道i处reset后能否成功
14. pref_b[j] - pref_b[i] + score_b[i] >= k => pref_b[j] >= pref_b[i] + k - score_b[i]
15. 因为pref_b是递增的，所以可以双指针处理

## key insights

1. Prefix sums and post-reset state:
   - `prefA[i] = sum(a1..ai)`, `prefB[i] = sum(b1..bi)`.
   - After a reset at round `i`, state becomes:
     - `(scoreA[i], scoreB[i])`
     - where one side is `0`, and the other is `|prefA[i] - prefB[i]|`.

2. For each reset position `i`, precompute whether game can be won **without any more resets**:
   - Earliest round where human reaches `k` after `i`:
     - `j1 = first j with prefA[j]-prefA[i]+scoreA[i] >= k`
   - Earliest round where computer reaches `k` after `i`:
     - `j2 = first j with prefB[j]-prefB[i]+scoreB[i] >= k`
   - `win[i] = (j2 < j1)` (computer loses earlier).
   - This is done with binary search on monotone prefix sums.

3. DP meaning:
   - `dp[i]` = minimum number of resets needed so that reset at round `i` is reachable safely.
   - Transition from previous reset `j` to next reset `r` is valid only if human does not lose before `r`:
     - `prefA[r]-prefA[j]+scoreA[j] < k`.
   - Then `dp[r] = min(dp[j] + 1)`.

4. Data-structure optimization used in code:
   - A heap `pq` tracks currently valid previous reset points `j` by their “human safety limit”.
   - While scanning `r = 1..n`, expired `j` are removed when they would already make human lose by `r`.
   - Another min-heap `dp` gives the minimum `dp[j]` among still-valid `j`.
   - This turns transitions into near `O(log n)` updates per position.

5. Building final answer:
   - For each `r`, if `win[r]` is true, candidate answer is `dp[r] + 1`.
   - Keep the best ending reset position and parent pointer `fp[r]`.
   - Reconstruct reset rounds by backtracking `fp`.
   - If no ending position works, output `-1`.

6. Complexity:
   - Prefix + scoring: `O(n)`
   - `win[i]` checks via binary search: `O(n log n)`
   - DP with heaps: `O(n log n)`
   - Total per test case: `O(n log n)`.