# Problem Description

Once the mischievous and wayward shooter named Shel found himself on a rectangular field of size 𝑛×𝑚, divided into unit squares. Each cell either contains a target or not.

Shel only had a lucky shotgun with him, with which he can shoot in one of the four directions:
- right-down
- left-down
- left-up
- right-up

When fired, the shotgun hits all targets in the chosen direction, the Manhattan distance to which does not exceed a fixed constant 𝑘. The Manhattan distance between two points (𝑥₁,𝑦₁) and (𝑥₂,𝑦₂) is equal to |𝑥₁−𝑥₂|+|𝑦₁−𝑦₂|.

Possible hit areas for 𝑘=3.

Shel's goal is to hit as many targets as possible. Please help him find this value.

## Input

Each test consists of several test cases. The first line contains a single integer 𝑡 (1≤𝑡≤1000) — the number of test cases. Then follows the description of the test cases.

The first line of each test case contains field dimensions 𝑛, 𝑚, and the constant for the shotgun's power 𝑘 (1≤𝑛,𝑚,𝑘≤10⁵, 1≤𝑛⋅𝑚≤10⁵).

Each of the next 𝑛 lines contains 𝑚 characters — the description of the next field row, where:
- The character '.' means the cell is empty
- The character '#' indicates the presence of a target

It is guaranteed that the sum of 𝑛⋅𝑚 over all test cases does not exceed 10⁵.

### ideas
1. 想象从左上到右下的处理（一条斜线一条斜线的处理）
2. 当a-b两条线的距离超过k的时候，b的上贡献要全部清理掉
3. 然后再考虑这个贡献要怎么计算。range updae
4. 不对。少了一个关键的东西，k
5. 考虑点(0, 0) 在处理斜线0时，它的贡献只在位置(0)
6. 当处理斜线(1)时，它的贡献范围是(0, 1), (1, 0)
7. 当处理斜线(2)时，它的贡献范围是(0, 2), (1, 1), (2, 0)
8. dp[i][j] = 在i,j处的结果
9. 在不考虑k的情况下，dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
10. 考虑k的情况下，就是把上面和左边k+1这条线上的抵消掉，就可以了
11. 还有k+2处的