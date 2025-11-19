Chris the Rabbit has been interested in arrays ever since he was a child. At the moment he is researching arrays with the length of n, containing only integers from 1 to n. He is not good at math, that's why some simple things drive him crazy. For example, yesterday he grew keen on counting how many different beautiful arrays there are. Chris thinks that an array is beautiful if it meets one of the two conditions:

- Each element, starting from the second one, is no more than the preceding one
- Each element, starting from the second one, is no less than the preceding one

Having got absolutely mad at himself and at math, Chris came to Stewie and Brian to ask them for help. However, they only laughed at him and said that the answer is too simple and not interesting. Help Chris the Rabbit to find the answer at last.

## Input

The single line contains an integer n which is the size of the array (1 ≤ n ≤ 10^5).

## Output

You must print the answer on a single line. As it can be rather long, you should print it modulo 1000000007.

## Examples

### Example 1

**Input:**
```
2
```

**Output:**
```
4
```

### Example 2

**Input:**
```
3
```

**Output:**
```
17
```


### ideas
1. n = 2, [1, 1], [1, 2], [2, 1], [2, 2]
2. n = 3, 怎么考虑呢？ dp[i][j] 表示到i为止，最大值为j时的计数, dp[i][j] = sum(dp[i-1][1], dp[i-1][2], .... dp[i-1][j])
3. 假设是一个递增的序列，a[1], a[2], a[3]... a[n]
4. 变成 a[1], a[2] - a[1], a[3] - a[2], .... a[n] - a[n-1]
5. 那么 let d[1] = a[1], d[2] = a[2] - a[1] , ... d[n-1] = a[n] - a[n-1]
6. 那么 a[n] = d[1] + d[2] + ... d[n-1] <= n
7. d[i] >= 0, d[1] >= 1
8. 可以根据a[n]的取值，确定不同的计数方案
9. a[n] = 1 (1, 只能a[1] = 1, 后面的等于0)
10. a[n] = 2, (n - 2, + 1) = n - 1
11. a[n] = x, 如果 d[1] 也可以等于0， C(n - 1 + x, x)种方案数， 然后减去  C(n - 2 + x, x)种d[1] = 0 的方案数
12. 考虑 n = 3的情况 
13. a[n] = 1, 1
14. a[n] = 2, C(n - 1 + 2, 2) - C(n - 2 + 2, 2) = C(4, 2) - C(3, 2) = 6 - 3 = 3
15. a[n] = 3, C(5, 3) - C(4, 3) = 5 * 4 * 3 / 6 - 4 = 6
16. 1 + 3 + 6 = 10
17. (10 - 3) * 2 + 3 = 17
18. 