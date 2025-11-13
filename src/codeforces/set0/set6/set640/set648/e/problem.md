Given a non-negative integer $k$ and $n$ non-negative integers $a_1, a_2, ..., a_n$. By writing some of these numbers one after another in any order and possibly using some of them several times (and not using some at all), the task is to create the shortest (smallest in number of digits) number divisible by $k$, or determine that this is impossible.

## Input

The first line contains two integers $n$ ($1 ≤ n ≤ 1,000,000$) and $k$ ($1 ≤ k ≤ 1000$) — the number of numbers and the required divisor, respectively.

The second line contains $n$ integers $a_1, a_2, ..., a_n$ ($0 ≤ a_i ≤ 10^9$).

## Output

If the answer exists, print "YES" (without quotation marks) on the first line, and the shortest number you're looking for, without leading zeros, on the second line. If the answer doesn't exist, print "NO" (without quotation marks) on the single line of output.

## Examples

### Example 1
#### Input
```
2 3
123 1
```
#### Output
```
YES
123
```

### Example 2
#### Input
```
1 10
1
```
#### Output
```
NO
```

### Example 3
#### Input
```
3 4
1 2 3
```
#### Output
```
YES
12
```

### Example 4
#### Input
```
3 777
12 23 345
```
#### Output
```
YES
121212
```

### ideas
1. 又是数学相关的～头疼
2. 不考虑复杂性， (a ++ b) % k = (a * pow(10, len(b)) % k + b % k) % k
3. 可以重复的情况下，似乎
4. dp[rem][len] = 在获得余数为rem, 且长度了len时的最少操作数
5. 关键这个len会一直涨，有没有办法判定，超过一定界限后，就停下来？
6. 盲猜是有的，可能不会很大
7. 那么dp[r1][l1] = dp[r2][l2] + dp[r3][l3] 
8. where r1 = (r2 * pow(10, l3) % k + r2) % k
9. l1 = l2 + l3