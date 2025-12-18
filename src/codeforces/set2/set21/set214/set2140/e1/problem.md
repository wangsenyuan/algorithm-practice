A valid configuration is defined as an arrangement of $n$ piles of stones such that:

The number of stones in each pile is an integer between $1$ and $m$ (both inclusive).

Given a valid configuration of $n$ piles of stones, some indices from $1$ to $n$ are marked as good. Alice and Bob start playing a game taking $n-1$ turns alternately with Alice going first. In each turn, they have to perform the following operation:

Choose any integer $i$ such that $1 \leq i \leq p$ (where $p$ is the number of piles left) and $i$ is good, and remove the $i$-th pile completely.

Note that after performing the operation once, the number of piles decreases by $1$ and the remaining piles are re-indexed. The game will end when there is only one pile left. It is guaranteed that the index $1$ is always good.

Let $x$ denote the number of stones in the final remaining pile. Alice wants to maximize $x$, whereas Bob wants to minimize it. Both Alice and Bob play optimally.

Find the sum of $x$ over all the possible valid configurations modulo $10^9 + 7$.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each testcase contains two integers $n$ ($1 \leq n \leq 20$) and $m$ ($1 \leq m \leq 2$) — the number of piles and the upper bound on the number of stones in a pile.

The second line of each testcase contains a single integer $k$ ($1 \leq k \leq n$) — the number of indices marked as good.

The third line of each testcase contains $k$ integers $c_1, c_2, \ldots, c_k$ ($1 = c_1 < c_2 < \ldots < c_k \leq n$) — the good indices. It is guaranteed that $1$ is always a good index (i.e. $c_1 = 1$).

It is guaranteed that the sum of $2^n$ over all test cases does not exceed $2^{20}$.

It is guaranteed that the sum of $m$ over all test cases does not exceed $10^6$.

## Output

For each testcase, print a single integer — the sum of $x$ over all the possible valid configurations modulo $10^9 + 7$.

## Example

**Input**
```
3
2 2
1
1
3 1
2
1 3
3 2
2
1 2
```

**Output**
```
6
1
11
```

## Note

For the first test case, the valid configurations are: $[1,1]$, $[1,2]$, $[2,1]$, $[2,2]$.

As there are only $2$ piles of stones, Alice can only choose the first pile and remove it. Thus, the sum of $x$ is equal to $1 + 1 + 2 + 2 = 6$.

For the second test case, the final pile is always $1$.


### ideas
1. 没啥想法～～～
2. 如果c = [1], 那么剩余的就只能是 a[n], 那么有 pow(m, n - 1)种选择，(1 + 2)
3. 如果 c = [1...n]， 是中位数
4. 在一般的情况下呢？
5. dp[state][l] 表示长度为l，且有state代表的状态
6. dp[state][l]  => dp[state][l-1] 可以根据l的奇偶性判断，当下是alice操作，还是bob操作
7. alice只会把最小值给移除掉， bob只会把最大值移出
8. 