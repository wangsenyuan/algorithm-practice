You would like to construct a string 𝑠
, consisting of lowercase Latin letters, such that the following condition holds:

For every pair of indices 𝑖
and 𝑗
such that 𝑠𝑖=𝑠𝑗
, the difference of these indices is even, that is, |𝑖−𝑗|mod2=0
.
Constructing any string is too easy, so you will be given an array 𝑐
of 26
numbers — the required number of occurrences of each individual letter in the string 𝑠
. So, for every 𝑖∈[1,26]
, the 𝑖
-th letter of the Latin alphabet should occur exactly 𝑐𝑖
times.

Your task is to count the number of distinct strings 𝑠
that satisfy all these conditions. Since the answer can be huge, output it modulo 998244353
.

### ideas

1. 一共n个位置（c[1] + .. + c[26])
2. 选择n/2个位置，给其中的字符（它们加起来，必须是 n/2, 不能多不能少）
3. 然后对它们进行排列，剩下的也计算排列。两边*起来
4. dp[i][x] = 前i个选择达到sum = x 的
5. dp[i][x] = dp[i-1][x] or dp[i-1][x - c[i]]
6. 26 * 1e6
7. 