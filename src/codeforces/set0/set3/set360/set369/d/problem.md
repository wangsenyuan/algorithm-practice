One fine morning, n fools lined up in a row. After that, they numbered each other with numbers from 1 to n, inclusive. Each fool got a unique number. The fools decided not to change their numbers before the end of the fun.

Every fool has exactly k bullets and a pistol. In addition, the fool number i has probability of pi (in percent) that he kills the fool he shoots at.

The fools decided to have several rounds of the fun. Each round of the fun looks like this: each currently living fool shoots at another living fool with the smallest number (a fool is not stupid enough to shoot at himself). All shots of the round are perfomed at one time (simultaneously). If there is exactly one living fool, he does not shoot.

Let's define a situation as the set of numbers of all the living fools at the some time. We say that a situation is possible if for some integer number j (0 ≤ j ≤ k) there is a nonzero probability that after j rounds of the fun this situation will occur.

Valera knows numbers p1, p2, ..., pn and k. Help Valera determine the number of distinct possible situations.

## Input

The first line contains two integers n, k (1 ≤ n, k ≤ 3000) — the initial number of fools and the number of bullets for each fool.

The second line contains n integers p1, p2, ..., pn (0 ≤ pi ≤ 100) — the given probabilities (in percent).

## Output

Print a single number — the answer to the problem.

## Examples

### Example 1

Input:
```
3 3
50 50 50
```

Output:
```
7
```

### Example 2

Input:
```
1 1
100
```

Output:
```
1
```

### Example 3

Input:
```
2 1
100 100
```

Output:
```
2
```

### Example 4

Input:
```
3 3
0 0 0
```

Output:
```
1
```

## Note

In the first sample, any situation is possible, except for situation {1, 2}.

In the second sample there is exactly one fool, so he does not make shots.

In the third sample the possible situations are {1, 2} (after zero rounds) and the "empty" situation {} (after one round).

In the fourth sample, the only possible situation is {1, 2, 3}.


### ideas
1. 完全一头雾水～～～
2. 最多有k轮，没次都要射击一颗子弹，共有k颗子弹
3. 如果i存活了，那么i+1也存活吗？
4. 也不一定；如果i是目前最小存活的下标，且还有子弹的情况下，虽然i会被它后面存活的人射击，但是它可射击i+1, 所以下轮有可能出现i存活，i+1死亡的情况
5. 但是概率要怎么应用呢？
6. 可能会存在i存活，i是最小的下标，接下来有x个不存活（都是被i消灭的），i+x后的都存活；只要i和i后面有一个人存活，之后的人都是安全的
7. dp[i][x] 的概率 = 到目前为止i，i存活，且i的后面有x个人被杀死的概率
8. dp[0][0] = 1.0
9. dp[i][j] => dp[i+j][0] * (i在这轮被杀死的概率), dp[i][j+1] * i在这轮幸存的概率
10. 但是这样子的时间复杂性是, k * n * k; 
11. 好像j不需要， 迭代到第k轮，如果i还存活，（它是最低的那个）， 那么j = k - i, j >= 0, => i <= k
12. 貌似可以搞