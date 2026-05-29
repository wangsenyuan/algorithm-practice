# F - Second Gap (ABC457)

**Contest:** [ABC457](https://atcoder.jp/contests/abc457) — AtCoder Beginner Contest 457  
**Task:** [https://atcoder.jp/contests/abc457/tasks/abc457_f](https://atcoder.jp/contests/abc457/tasks/abc457_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

You are given an integer N and an integer sequence D = (D_1, D_2, …, D_{N-1}) of length N − 1.

Find the number, modulo 998244353, of permutations P = (P_1, P_2, …, P_N) of (1, 2, …, N) that satisfy the following condition.

- For each 1 ≤ i ≤ N − 1, let P_a and P_b be the elements with the largest and the second largest values, respectively, among (P_i, P_{i+1}, …, P_N); then |a − b| = D_i.

## Constraints

- 2 ≤ N ≤ 2 × 10^5
- 1 ≤ D_i ≤ N − i
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N
D_1 D_2 … D_{N-1}
```

## Output

Output the number, modulo 998244353, of permutations satisfying the condition.

## Sample Input 1

```text
3
1 1
```

## Sample Output 1

```text
4
```

For example, we can verify that (2, 3, 1) satisfies the condition as follows.

- We have (P_1, P_2, P_3) = (2, 3, 1). The largest value is P_2 and the second largest value is P_1, and |2 − 1| = 1 = D_1.
- We have (P_2, P_3) = (3, 1). The largest value is P_2 and the second largest value is P_3, and |2 − 3| = 1 = D_2.

Four permutations satisfy the condition: (1, 2, 3), (1, 3, 2), (2, 3, 1), (3, 2, 1). Thus, output 4.

## Sample Input 2

```text
5
1 2 2 1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
15
4 4 4 4 4 4 3 2 2 2 2 2 1 1
```

## Sample Output 3

```text
70270200
```


## ideas
1. 这个题目好拗口
2. 对于排列P, 后缀P[i:], 满足其中最大的两个数的位置距离 = D[i]
3. 比如 P = [2, 3, 1], 
4. when i = 1, abs(1 - 2) = 1 = D[1]
5. when i = 2, abs(2 - 3) = 1 = D[2]
6. 所以，必须有D[n-1] = 1 (因为只有两个位置)
7. 完全没有头绪～
8. dp[i][a][b] 表示到i为止，最大值的位置是a,第二大的值的位置是b时的方案数
9. dp[i-1][a][b] = dp[i-1][a][b] * ? 表示在保持最大值、最小值位置不变时的方案数，那么就是用当前的数，替换a，a处的替换b，其他的数替换到当前位置（D[i-1] = D[i]) ? =  (n - (i-1) - 2)
10. dp[i-1][i-1][a] = dp[i-1][a][?] 如果 (a - (i-1)) = D[i] 当前新加入的（最大值）放在位置i-1上，那么 a = i-1 + D[i]
11. dp[i-1][a][i-1] = dp[i-1][a][?] 如果 (a - (i-1)) = D[i], 用a处的和当前新加入的进行交换, a = i-1+D[i]
12. 所以，如果新加入最大值，不是放置在特殊位置上，那么已有的计数方案数 * （n - (i - 1) - 2)
13. 如果，新加入最大值放置在位置i-1上，i-1的计数 = 以 i-1+D[i]为最大值的方案数
14. dp[a] 表示以位置a为最大值的方案数， 
15. dp[a] *= (n - i - 2) D[i] = D[i+1] 必须成立才可以
16. dp[i] = sum(dp[i + d[i]]) (作为最大值)
17. dp[a] 不变（i作为第二大的值）