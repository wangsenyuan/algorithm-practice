# Problem E2

This is the hard version of the problem. The difference between the versions is that in this version, $m \leq 10^6$. You can hack only if you solved all versions of this problem.

A valid configuration is defined as an arrangement of $n$ piles of stones such that:

- The number of stones in each pile is an integer between $1$ and $m$ (both inclusive).

Given a valid configuration of $n$ piles of stones, some indices from $1$ to $n$ are marked as good. Alice and Bob start playing a game taking $n-1$ turns alternately with Alice going first. In each turn, they have to perform the following operation:

- Choose any integer $i$ such that $1 \leq i \leq p$ (where $p$ is the number of piles left) and $i$ is good, and remove the $i$-th pile completely.

Note that after performing the operation once, the number of piles decreases by $1$ and the remaining piles are re-indexed. The game will end when there is only one pile left. It is guaranteed that the index $1$ is always good.

Let $x$ denote the number of stones in the final remaining pile. Alice wants to maximize $x$, whereas Bob wants to minimize it. Both Alice and Bob play optimally.

Find the sum of $x$ over all the possible valid configurations modulo $10^9 + 7$.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each testcase contains two integers $n$ ($1 \leq n \leq 20$) and $m$ ($1 \leq m \leq 10^6$) — the number of piles and the upper bound on the number of stones in a pile.

The second line of each testcase contains a single integer $k$ ($1 \leq k \leq n$) — the number of indices marked as good.

The third line of each testcase contains $k$ integers $c_1, c_2, \ldots, c_k$ ($1 = c_1 < c_2 < \ldots < c_k \leq n$) — the good indices. It is guaranteed that $1$ is always a good index (i.e. $c_1 = 1$).

It is guaranteed that the sum of $2^n$ over all test cases does not exceed $2^{20}$.

It is guaranteed that the sum of $m$ over all test cases does not exceed $10^6$.

## Output

For each testcase, print a single integer — the sum of $x$ over all the possible valid configurations modulo $10^9 + 7$.

## Example

### Example Input

```text
5
2 3
1
1
7 4
3
1 4 6
12 31
6
1 3 5 7 9 11
11 121
11
1 2 3 4 5 6 7 8 9 10 11
19 6969
2
1 19
```

### Example Output

```text
18
33664
909076242
683044824
901058932
```

## Note

For the first test case, the valid configurations are: $[1,1]$, $[1,2]$, $[1,3]$, $[2,1]$, $[2,2]$, $[2,3]$, $[3,1]$, $[3,2]$, $[3,3]$.

As there are only $2$ piles of stones, Alice can only choose the first pile and remove it. Thus, the sum of $x$ is equal to $1+1+1+2+2+2+3+3+3=18$.



### ideas
1. dp[state] state表示位置的状态,0/1表示是否还存在这个位置，状态转移，勉强可以搞，但是怎么赋值呢？
2. 肯定要考虑贡献; 
3. 先考虑n为奇数，如果m=x, 小于等于x的数 = h, 大于等于x的数也有h, 那么结果就是x
4. 那么这个时候，有一个位置确定是放x的，剩余(n-1)个位置中，选择一半的位置，C(n, h), 然后每个位置放置(1...x)
5. 剩余一半的位置,(x....m)
6. 不对。还有good position的影响
7. 如果是想让最后的结果是x，要怎么安排，才能保证最后的结果一定是x呢？
8. 假设只有两个good的位置 [1....i]
9. 如果现在是alice操作, 那么最后保留的数 = max(a[i-1], a[i])
10. 如果 a[i-1] > a[i], 那么就删除i, 如果 a[i-1] < a[i], 那么就删除1
11. 但是如果当前是bob的操作，那么结果反过来 min(a[i-1], a[i])
12. 如果 [1......i.....j] 情况就有点复杂了