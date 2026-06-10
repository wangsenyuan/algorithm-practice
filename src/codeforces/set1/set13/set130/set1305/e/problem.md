# E. Kuroni and the Score Distribution

[Problem link](https://codeforces.com/problemset/problem/1305/E)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

Kuroni is the coordinator of the next Mathforces round written by the "Proof by
AC" team. All the preparation has been done, and he is discussing with the team
about the score distribution for the round.

The round consists of `n` problems, numbered from `1` to `n`. The problems are
ordered in increasing order of difficulty, and no two problems have the same
difficulty. A score distribution for the round can be denoted by an array
`a_1, a_2, ..., a_n`, where `a_i` is the score of the `i`-th problem.

Kuroni thinks that the score distribution should satisfy the following
requirements:

- The score of each problem should be a positive integer not exceeding `10^9`.
- A harder problem should grant a strictly higher score than an easier problem.
  In other words, `1 <= a_1 < a_2 < ... < a_n <= 10^9`.
- The **balance** of the score distribution, defined as the number of triples
  `(i, j, k)` such that `1 <= i < j < k <= n` and `a_i + a_j = a_k`, should be
  exactly `m`.

Help the team find a score distribution that satisfies Kuroni's requirement. If
such a score distribution does not exist, output `-1`.

## Input

The first and only line contains two integers `n` and `m` (`1 <= n <= 5000`,
`0 <= m <= 10^9`) — the number of problems and the required balance.

## Output

If there is no solution, print a single integer `-1`.

Otherwise, print one line containing `n` integers `a_1, a_2, ..., a_n` —
a score distribution that satisfies all the requirements. If there are multiple
answers, print any of them.

## Example

### Input

```text
5 3
8 0
4 10
```

### Output

```text
4 5 9 13 18
10 11 12 13 14 15 16 17
-1
```

## Note

In the first example, there are `3` triples `(i, j, k)` that contribute to the
balance of the score distribution:

- `(1, 2, 3)`
- `(1, 3, 4)`
- `(2, 4, 5)`


### ideas
1. 正好要有m个3元组满足a[i] + a[j] = a[k] 
2. 好像要得到这样的三元组还挺难的.
3. 假设到i为止,得到了dp[i]个这样的三元组, 现在计算dp[i+1]
4. dp[i+1] = dp[i] + (由i+1组成的3元祖), 也就是说有多少组(j1, j2), a[j1] + a[j2] = a[i+1]?
5. 如果 a[i+1] = a[1] + a[i] = a[2] + a[i-1] = a[3] + a[i-2], ... 貌似可以最多
6. 比如 1, 2, 3, 4, 5, 6, 7, 8 (这样子肯定最多)
7. 如果 a[1] != 1, 比如 a[1] = 2, 那么貌似只能有3个数,组成一组?
8. a, b, c, d, e 
9. a + d = e, b + c = e (这个似乎很好满足)
10. 如果 d = a + c 能满足吗?
11. e = a + a + c
12. e = b + c => b = 2 * a
13. 假设[3, 6, 9, 12, 15, 18, 21, 24]
14. 奇怪,为啥能一直加下去? 难道这要 b = 2 * a 就可以?
15.  等差数列,且差值= a[1]
16.  a[1], 2 * a[1], 3 * a[1], 4 * a[1], ... 那么就
17.  那么这样子, a[i] = i * a[1], 因为 i = 1 + i - 1 = 2 + i - 2 = 3 + i - 3. 贡献 (i-1) / 2