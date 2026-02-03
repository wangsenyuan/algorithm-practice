You are given an array $a_i$ of length $n$. You may consecutively apply two operations to this array:

remove some subsegment (continuous subsequence) of length $m < n$ and pay $m \cdot a$ coins;
change some elements of the array by at most $1$, and pay $b$ coins for each change.
Please note that each operation may be applied at most once (and may be not applied at all), so you can remove only one segment and each number may be changed (increased or decreased) by at most $1$. Also note that you are not allowed to delete the whole array.

Your goal is to calculate the minimum number of coins that you need to spend in order to make the greatest common divisor of the elements of the resulting array be greater than 1.

## Input

The first line contains integers $n$, $a$ and $b$ ($1 \le n \le 10^6$, $0 \le a, b \le 10^9$) — the length of the array, the cost of removing a single element in the first operation, and the cost of changing an element, respectively.

The second line contains $n$ integers $a_i$ ($2 \le a_i \le 10^9$) — elements of the array.

## Output

Print a single number — the minimum cost of changes needed to obtain an array such that the greatest common divisor of all its elements is greater than $1$.

## Examples

### Input
```
3 1 4
4 2 3
```

### Output
```
1
```

### Input
```
5 3 2
5 17 13 5 6
```

### Output
```
8
```

### Input
```
8 3 4
3 7 5 4 3 12 9 4
```

### Output
```
13
```

## Note

In the first sample the optimal way is to remove number $3$ and pay $1$ coin for it.

In the second sample you need to remove a segment $[17, 13]$ and then decrease number $6$. The cost of these changes is equal to $2 \cdot 3 + 2 = 8$ coins.


### ideas
1. 如果有删除操作，那么必须是第一步就完成的（在被删除的元素上，进行操作2没有意义）
2. dp[i] 表示在没有删除操作的情况下，将前缀i，只进行操作2时，得到gcd(a[:i]) > 1 的最小代价
3. fp[i]... 将后缀.... 的最小代价
4. dp[l] + fp[r] + (r - l - 1) * a
5. 但是还有个问题，就是 前缀和后缀的gcd > 1
6. 但是，如果进行操作，那么最优的操作，是不是得到gcd = 2？
7. 如果前后缀的gcd > 1 => 0; 否则的话，是不是肯定最优的gcd就是2呢？
8. 好像不一定; 假设奇数的部分很多。
9. 在指定g的情况下，计算最优解？
10. 按照a的大小降序处理。处理x的时候， dp[i] = i - 前缀变成x的倍数的最小值
11. 假设 y % x != 0, 它的cost  = min(x - y % x, y % x) * b
12. 如果不考虑顺序，那么可以用freq去处理
13. 但是因为有删除中间一段的操作，就没法搞了～
14. 搞错了。每个数只能变一次
15. 要么大家都变成2，这个肯定能达到，要么变成另外一个数x（但是大家也只能操作1次）
16. 那么意味着这个数，只可能是a[0] - 1, a[0], a[0] + 1 的一个质数因子（那么这个数量是很少的）
17. dp[i][j] = 其中某个变化的值，就很好处理了