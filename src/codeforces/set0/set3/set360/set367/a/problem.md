Sereja loves all sorts of algorithms. He has recently come up with a new algorithm, which receives a string as an input. Let's represent the input string of the algorithm as $q = q_1q_2... q_k$. 

The algorithm consists of two steps:

1. Find any continuous subsequence (substring) of three characters of string $q$, which doesn't equal to either string "zyx", "xzy", "yxz". If $q$ doesn't contain any such subsequence, terminate the algorithm, otherwise go to step 2.
2. Rearrange the letters of the found subsequence randomly and go to step 1.

Sereja thinks that the algorithm works correctly on string $q$ if there is a non-zero probability that the algorithm will be terminated. But if the algorithm anyway will work for infinitely long on a string, then we consider the algorithm to work incorrectly on this string.

Sereja wants to test his algorithm. For that, he has string $s = s_1s_2... s_n$, consisting of $n$ characters. The boy conducts a series of $m$ tests. As the $i$-th test, he sends substring $s_{l_i}s_{l_i + 1}... s_{r_i}$ ($1 \le l_i \le r_i \le n$) to the algorithm input. Unfortunately, the implementation of his algorithm works too long, so Sereja asked you to help. For each test $(l_i, r_i)$ determine if the algorithm works correctly on this test or not.

## Input

The first line contains non-empty string $s$, its length ($n$) doesn't exceed $10^5$. It is guaranteed that string $s$ only contains characters: 'x', 'y', 'z'.

The second line contains integer $m$ ($1 \le m \le 10^5$) — the number of tests. Next $m$ lines contain the tests. The $i$-th line contains a pair of integers $l_i, r_i$ ($1 \le l_i \le r_i \le n$).

## Output

For each test, print "YES" (without the quotes) if the algorithm works correctly on the corresponding test and "NO" (without the quotes) otherwise.

## Example

### Input

```
zyxxxxxxyyz
5
5 5
1 3
1 11
1 4
3 6
```

### Output

```
YES
YES
NO
YES
NO
```

## Note

In the first example, in test one and two the algorithm will always be terminated in one step. In the fourth test you can get string "xzyx" on which the algorithm will terminate. In all other tests the algorithm doesn't work correctly.


### ideas
1. 给定q，在q中找到一个连续长度为3的字符串，是"zyx", "xzy", "yxz"的任何一个，就终止
2. 如果找不到，就继续，打乱继续？
3. 如果这个序列可以表示成，zyxzyxzyx, 或者 xzyxzy, 或者 yxzyxz, 那么就能终止，否则肯定不行
4. 也就是说cnt[x], cnt[y], cnt[z]的差值，不能超过1，