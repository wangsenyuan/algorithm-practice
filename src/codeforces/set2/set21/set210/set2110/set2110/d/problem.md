# Problem D - Robot Battery Game

In 2077, when robots took over the world, they decided to compete in the following game.

There are $n$ checkpoints, and the $i$-th checkpoint contains $b_i$ batteries. Initially, the Robot starts at the 1-st checkpoint with no batteries and must reach the $n$-th checkpoint.

There are a total of $m$ one-way passages between the checkpoints. The $i$-th passage allows movement from point $s_i$ to point $t_i$ ($s_i < t_i$), but not the other way. Additionally, the $i$-th passage can only be used if the robot has at least $w_i$ charged batteries; otherwise, it will run out of power on the way.

When the robot arrives at point $v$, it can additionally take any number of batteries from $0$ to $b_v$, inclusive. Moreover, it always carries all previously collected batteries, and at each checkpoint, it recharges all previously collected batteries.

Find the minimum number of batteries that the robot can have at the end of the journey, or report that it is impossible to reach from the first checkpoint to the last.

---

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

For each test case:

- The first line contains two integers $n, m$ ($2 \leq n \leq 2 \cdot 10^5$, $0 \leq m \leq 3 \cdot 10^5$) — the number of checkpoints and the number of passages, respectively.
- The second line contains $n$ numbers $b_i$ ($0 \leq b_i \leq 10^9$) — the number of batteries at the $i$-th checkpoint.
- The next $m$ lines contain three integers $s_i, t_i, w_i$ ($1 \leq s_i < t_i \leq n$, $1 \leq w_i \leq 10^9$) — the endpoints of the passage and the minimum number of batteries required to pass through it.

It is guaranteed that the sum of $n$ does not exceed $2 \cdot 10^5$.

It is guaranteed that the sum of $m$ does not exceed $3 \cdot 10^5$.

---

## Output

For each test case, output the minimum number of batteries that you can have at the end of the journey, or $-1$ if it is impossible to reach point $n$.

## ideas
1. 假设最后需要x个电池，那么这x个电池越早拿到越好
2. 这是因为，前面拿到的后面也可以用。且前面拿到也更有利
3. 同时，如果x可以，那么x+1肯定也可以（如果可以拿到x+1个的话）
4. 所以可以二分
5. 所以，可以贪心的处理。在给定x的情况下，尽量的多拿。拿到以后，检查是否可以到达下一个点。如果能到达，就到达
6. 可以分成两个阶段。 先贪心的看，能到达什么位置。然后从这些位置开始移动
7. 对于任何一个节点，先计算出，它到达终点需要的最小的电池数。
8. 这个就好处理了
9. dp[u] = min(max(dp[v], min(w[u,v])))