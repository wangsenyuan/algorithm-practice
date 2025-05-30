Let $c_1, c_2, \ldots, c_n$ be a permutation of integers $1, 2, \ldots, n$. Consider all subsegments of this permutation containing an integer $x$. Given an integer $m$, we call the integer $x$ **good** if there are exactly $m$ different values of maximum on these subsegments.

Cirno is studying mathematics, and the teacher asks her to count the number of permutations of length $n$ with exactly $k$ good numbers.

Since the answer may be very big, you only need to tell her the number of permutations modulo $p$.

A **permutation** is an array consisting of $n$ distinct integers from $1$ to $n$ in arbitrary order. For example, `[2, 3, 1, 5, 4]` is a permutation, but `[1, 2, 2]` is not a permutation ($2$ appears twice in the array) and `[1, 3, 4]` is also not a permutation ($n=3$ but there is $4$ in the array).

A sequence $a$ is a **subsegment** of a sequence $b$ if $a$ can be obtained from $b$ by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

---

### Input

The first line contains four integers $n, m, k, p$ ($1 \leq n \leq 100$, $1 \leq m \leq n$, $1 \leq k \leq n$, $1 \leq p \leq 10^9$).


### ideas
1. 包括i的连续的子序列，共有(l+1) * (r+1)个，要求这些子序列的最大值，组成的set的size 刚好是m，且这样的i有k个的排列数
2. 完全没有想法
3. 包括n的子序列的set的大小是1
4. 包括n-1的子序列的set的大小是2 (n, n - 1)
5. 包括n-2的子序列的set的大小，可能是2 (n, n-2), 也可能是3 (n, n - 1, n-2)
6.  当n在n-1，n-2的中间的时候，n-1就没法被选中，否则就可以被选中
7. 包括n-3的子序列的set的大小,可能是2，3，4
8. 假设，我希望这个set的大小正好是m, 那么比a[i]大的数字中，可以选择 n - a[i] + 1 - 2 个数（除了n和a[i])
9. 然后排列这m个数，怎么排列是合适的？
10. 假设a[j] ( > a[i])被选中了，把i排列到j的旁边时，似乎如果a[j]正好形成过m-1个size的集合的话，那么把i排到j的旁边
11. 无论是左边，还是右边，都可以正好得到set为m的size
12. 这是因为，如果j正好得到了m-1大小的set，那么把i放到它旁边后，这个m-1个set仍然可以保留，且增加了i自己
13. dp[i][sz] = sum of dp[j][sz-1] where i < j
14. dp[i][sz][j] j表示目前为止有多少个满足sz = k 的数字的计数
15. 还有一个问题是这个k的限制
16. 假设i是一个good的数，也就是说，包含它的子序列的最大值的set的size是m
17. 那么在i的两边的的数，都不可能是是一个good的数
18. 但是 i, j, i - 1, i和i-1可以同时是good的数
19. 那么k个good的数，是不是意味着只能有k段连续的数比如 1, 2, 3, 5(这一段只能有一个)
20. 1, 2, 3, 5, 4, 6, 8, 7, 9 
21. 2， 1， 3， 5似乎也能算作一段