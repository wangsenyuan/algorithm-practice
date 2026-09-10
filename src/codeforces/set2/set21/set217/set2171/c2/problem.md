# C2. Renako Amaori and XOR Game (hard version)

[Problem link](https://codeforces.com/problemset/problem/2171/C2)

**Contest:** [Codeforces Round 1065 (Div. 3)](https://codeforces.com/contest/2171)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. The only difference between the easy and hard versions is that in the hard version, `a_i, b_i <= 10^6`.

Ajisai and Mai are given arrays `a` and `b` of length `n` (`0 <= a_i, b_i <= 10^6`). They will play a game that lasts for `n` turns, where Ajisai moves on odd-numbered turns and Mai moves on even-numbered turns. On the `i`-th turn, the player to move may choose to swap `a_i` and `b_i`, or pass.

Note that if a swap occurs, the index that is being swapped must match the turn number. For example, on the first turn, Ajisai may choose to swap `a_1` and `b_1`, or pass. On the second turn, Mai may choose to swap `a_2` and `b_2`, or pass. This continues for `n` turns. Thus, only Ajisai can swap odd indices, and only Mai can swap even indices.

At the end of the game, Ajisai achieves a score of `a_1 ⊕ a_2 ⊕ … ⊕ a_n`, and Mai achieves a score of `b_1 ⊕ b_2 ⊕ … ⊕ b_n`. The player with the higher score wins. If the players have the same score, the game ends in a tie.

Determine the outcome of the game with optimal play. More formally, one player is considered to win with optimal play if there exists a strategy for them such that they always win, regardless of their opponent's choices. The game is considered a tie with optimal play if neither player has such a strategy.

`⊕` denotes the bitwise XOR operation.

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of test cases.

The first line of each test case contains a single integer `n` (`1 <= n <= 2 * 10^5`).

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`0 <= a_i <= 10^6`).

The third line of each test case contains `n` integers `b_1, b_2, …, b_n` (`0 <= b_i <= 10^6`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output on a single line `"Ajisai"` if Ajisai wins with optimal play, `"Mai"` if Mai wins with optimal play, or `"Tie"` if the game ends in a tie with optimal play.

You may output the answer in any case (upper or lower). For example, the strings `"mAi"`, `"mai"`, `"MAI"`, and `"maI"` will be recognized as `"Mai"`.

## Example

### Input

```text
6
4
1 4 6 1
3 2 3 7
6
20 11 1 7 7 0
14 8 3 6 17 6
4
2 6 3 6
3 4 7 1
5
1 4 5 5 3
6 7 1 2 13
6
9 5 9 17 17 6
1 13 6 13 1 15
5
2 3 8 1 5
3 1 6 14 7
```

### Output

```text
Mai
Ajisai
Tie
Ajisai
Mai
Tie
```

### Note

In the first example, one way the game might play out is as follows:

- On turn 1, Ajisai chooses to swap `a_1` and `b_1`. Now the arrays are `a = [3, 4, 6, 1]` and `b = [1, 2, 3, 7]`.
- On turn 2, Mai chooses to swap `a_2` and `b_2`. Now the arrays are `a = [3, 2, 6, 1]` and `b = [1, 4, 3, 7]`.
- On turn 3, Ajisai chooses to pass.
- On turn 4, Mai chooses to swap `a_4` and `b_4`. Now the arrays are `a = [3, 2, 6, 7]` and `b = [1, 4, 3, 1]`.

Now, Ajisai's final score is `3 ⊕ 2 ⊕ 6 ⊕ 7 = 0` and Mai's final score is `1 ⊕ 4 ⊕ 3 ⊕ 1 = 7`. Therefore, Mai wins the game.

It is not guaranteed that the above description is representative of optimal play.

## ideas
1. 在最高位进行考虑, 如果a[i] & hi != b[i] & hi, 且只有一个, i是奇数
2. 那么alice获胜, 因为alice可以得到一个最高位被设置的结果
3. 如果存在偶数位这样的情况, 那么bob可以抵消掉
4. 假设这样的个数奇数位更多, 那还是alice获胜
5. 如果一样多, 那么最高位比不出来, 但是某些结果已经被确定了
