You are given two arrays $a$ and $b$ of length $n$ ($1 \le a_i \le b_i$).

In one operation, you may either:

- choose an index $i$ ($1 \le i \le n$) and set $a_i := a_i + 1$, or  
- double all elements of $a$.

Let $x$ denote the minimum number of operations needed to make $a = b$. Two arrays $a$ and $b$ of length $n$ are considered equal if $a_i = b_i$ for all $1 \le i \le n$.

Find the value of $x$. Additionally, count the number of sequences of operations that make $a = b$ using exactly $x$ operations. Two such sequences of operations are considered different if, for any $1 \le j \le x$, the $j$-th operation of each sequence differs (either in the type of operation selected or the index chosen, if applicable).

Since the number of sequences may be large, output it modulo $10^6 + 3$. Note that $10^6 + 3$ is a prime number.

## Input

The first line contains a single integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

The first line of each test case contains a single integer $n$ ($2 \le n \le 2 \cdot 10^5$).

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le 10^6$).

The third line of each test case contains $n$ integers $b_1, b_2, \ldots, b_n$ ($a_i \le b_i \le 10^6$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output two integers: the value of $x$, and the number of sequences of operations that make $a = b$ using exactly $x$ operations, modulo $10^6 + 3$. The value of $x$ should be printed exactly; that is, it should not be taken modulo $10^6 + 3$.

## Example

### Input

```
8
6
1 3 6 4 3 2
3 7 10 4 4 8
2
1 1
4 3
5
2 3 2 5 1
18 13 10 30 7
5
5 4 3 6 2
100 125 231 113 107
4
2 2 2 2
2 2 2 2
4
1 1 1 1
2 2 2 2
7
1 1 1 1 1 1 200000
200000 200000 200000 200000 200000 200000 200000
3
542264 174876 441510
641112 325241 995342
```

### Output

```
17 827116
3 1
12 288
35 567812
0 1
1 1
1199994 0
803045 366998
```

## Note

In the second sample, it is possible to convert $a$ into $b$ using only three operations. There is only one way to do so, namely:

1. Add $1$ to $a_1$, yielding $a = [2, 1]$.
2. Then double all elements of $a$, yielding $a = [4, 2]$.
3. Then add $1$ to $a_2$, yielding $a = [4, 3]$.

Then we have $a = b$, as desired. It can be shown that it is impossible to convert $a$ into $b$ using fewer than three operations.


### ideas
1. a -> b 单个数字先看看怎么变化
2. 如果b是偶数，且 b/2 >= a, 那么 double一次, b /= 2
3. 否则就必须通过操作1， 从a增加到b
4. 除非a = 1, b = 2, 那么 double 和 add的效果是一致的
5. 对单个数进行观察 d[i] 表示能对这个数进行多少次double
6. 那么 min(d[i]) = 所有数能够double的操作数，剩下的都必须经过add来操作
7. x很好算，但是seq的数量，有点难搞～
8. double的位置，相当于阶段，在这个阶段中间，对所有的add操作，进行排列
9. 假设y[1]，... y[i] 是当前阶段需要执行的操作1的数量
10. 那么就是 sum >= 1e6 + 3时，直接返回0