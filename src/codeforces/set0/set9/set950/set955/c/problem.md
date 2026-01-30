You're given $Q$ queries of the form $(L, R)$.

For each query you have to find the number of such $x$ that $L \le x \le R$ and there exist integer numbers $a > 0$, $p > 1$ such that $x = a^p$.

## Input

The first line contains the number of queries $Q$ ($1 \le Q \le 10^5$).

The next $Q$ lines contains two integers $L, R$ each ($1 \le L \le R \le 10^{18}$).

## Output

Output $Q$ lines — the answers to the queries.

## Example

### Input
```
6
1 4
9 9
5 7
12 29
137 591
1 1000000
```

### Output
```
2
1
0
3
17
1111
```

## Note

In query one the suitable numbers are $1$ and $4$.

### ideas
1. 显然p越小(比如2), a越大（起点）
2. 可以迭代p，从2开始，（当 pow(2, p) > R的时候，停止
3. 所以p <= 60
4. 然后对于每个p计算出最小的a,最大的b,(但是 4^^2 = 2^^4)所以，还得再处理一下
5. 当p= 2的时候， a = [l1,r1], 当 p = 3 的时候， a = [l2, r2]
6. 有没有可能某个数 x, pow(x, 2), pow(x, 3)都满足条件（存在的）
7. 所以还是的用incluesion, excludesion 的方式