# Problem Statement

It turns out that you are a great fan of rock band **AC/PE**. Peter learned that and started the following game: he plays the first song of the list of $n$ songs of the group, and you have to find out the name of the song. After you tell the song name, Peter immediately plays the following song in order, and so on.

The $i$-th song of AC/PE has its recognizability $p_i$. This means that if the song has not yet been recognized by you, you listen to it for exactly one more second and with probability of $p_i$ percent you recognize it and tell its name. Otherwise, you continue listening. Note that you can only try to guess it only when it is an integer number of seconds after the moment the song starts playing.

In all AC/PE songs, the first words of chorus are the same as the title, so when you've heard the first $t_i$ seconds of $i$-th song and its chorus starts, you immediately guess its name for sure.

For example, in the song *Highway To Red* the chorus sounds pretty late, but the song has high recognizability. In the song *Back In Blue*, on the other hand, the words from the title sound close to the beginning of the song, but it's hard to name it before hearing those words. You can name both of these songs during a few more first seconds.

Determine the expected number of songs you will recognize if the game lasts for exactly $T$ seconds (i.e., you can make the last guess on the second $T$, after that the game stops).

If all songs are recognized faster than in $T$ seconds, the game stops after the last song is recognized.

---

## Input

The first line of the input contains numbers $n$ and $T$ ($1 \leq n \leq 5000$, $1 \leq T \leq 5000$), separated by a space. 

Next $n$ lines contain pairs of numbers $p_i$ and $t_i$ ($0 \leq p_i \leq 100$, $1 \leq t_i \leq T$). The songs are given in the same order as in Petya's list.

## Output

Output a single number — the expected number of the number of songs you will recognize in $T$ seconds. Your answer will be considered correct if its absolute or relative error does not exceed $10^{-6}$.

## ideas
```
输入 n(1≤n≤5000) 和 T(1≤T≤5000)。
有 n 首歌，每首歌输入两个整数 pi(0≤pi≤100) 和 t(1≤t≤T)。

你在听歌识曲，按输入顺序依次播放。
每首歌从头开始听。每过一秒，识别出这首歌的概率是 p。在这首歌的第 t 秒，你可以立刻识别出这首歌。
成功识别后，立刻开始播放下一首歌。
注：相当于有 t 次抽卡机会，且第 t 次（最后一次）一定抽中。
注：如果所有歌曲都播放完毕，则识别结束，不会重复循环。

输出在 T 秒内识别出的歌曲个数的期望值。
与正确答案的绝对（相对）误差必须 ≤ 1e-6。
```

- dp[i]表示听完前i首歌的期望时间（能反过来算吗？）
- dp[i] = dp[i-1] + 听完第i首歌的期望时间
  ```
    s[j] = s[j-1] * p[i] + 1 - p[i]  
  ```
- 感觉不大对

```
由期望的线性性质可知，累加每首歌的期望，就是总期望。
考虑其中一首歌，设其在 T 秒内被识别出的概率是 q，那么这首歌对总期望的贡献是 1*q + 0*(1-q) = q，正好等于概率。
所以累加每首歌在 T 秒内被识别出的概率，就是总期望。

「T 秒内」等价于「第 1,2,3,...,T 秒」。
计算每首歌分别在恰好第 1,2,3,...,T 秒被识别出的概率，累加，就是这首歌在 T 秒内被识别出的概率。
注：比如第二首歌，无法在第 1 秒被识别出，那么这个时刻识别出它的概率为 0。

定义 f[i][j] 表示第 i 首歌恰好在第 j 秒被识别出的概率。
比如 1 秒就识别出第 i 首歌，那么问题变成第 i-1 首歌恰好在第 j-1 秒被识别出的概率，即 f[i-1][j-1]。

以 t[i] = 4 为例，我们有（下文 p 表示 p[i]，t 表示 t[i]）
f[i][j] = f[i-1][j-1] * p + f[i-1][j-2] * (1-p)*p + f[i-1][j-3] * (1-p)^2*p + f[i-1][j-4] * (1-p)^3
对比 f[i][j-1] 和 f[i][j] 的转移方程，比较相似之处（请在草稿纸上推导，回想一下错位相减法），化简可得
f[i][j] = f[i-1][j-1] * p + (f[i][j-1] - f[i-1][j-5] * (1-p)^3) * (1-p) + f[i-1][j-4] * (1-p)^4
一般地
f[i][j] = f[i-1][j-1] * p + (f[i][j-1] - f[i-1][j-t-1] * (1-p)^(t-1)) * (1-p) + f[i-1][j-t] * (1-p)^t
（我代码中把 f[i-1][j-1] * p 移到了右边，只是改了下位置）

初始值 f[0][0] = 1。
答案为 f[i][j] 之和，其中 1<=i<=n 且 1<=j<=T。
```