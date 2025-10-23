# Problem: Beautiful Array

Let's call an array consisting of $n$ integer numbers $a_1, a_2, ..., a_n$ **beautiful** if it has the following
property:

- Consider all pairs of numbers $x, y$ ($x \neq y$), such that number $x$ occurs in the array $a$ and number $y$ occurs
  in the array $a$;
- For each pair $x, y$ must exist some position $j$ ($1 \leq j < n$), such that at least one of the two conditions are
  met, either $a_j = x, a_{j+1} = y$, or $a_j = y, a_{j+1} = x$.

Sereja wants to build a beautiful array $a$, consisting of $n$ integers. But not everything is so easy, Sereja's friend
Dima has $m$ coupons, each contains two integers $q_i, w_i$. Coupon $i$ costs $w_i$ and allows you to use as many
numbers $q_i$ as you want when constructing the array $a$. Values $q_i$ are distinct.

Sereja has no coupons, so Dima and Sereja have made the following deal. Dima builds some beautiful array $a$ of $n$
elements. After that he takes $w_i$ rubles from Sereja for each $q_i$ which occurs in the array $a$. Sereja believed
his friend and agreed to the contract, and now he is wondering, what is the maximum amount of money he can pay.

Help Sereja, find the maximum amount of money he can pay to Dima.

## Input

The first line contains two integers $n$ and $m$ ($1 \leq n \leq 2 \cdot 10^6, 1 \leq m \leq 10^5$).

Next $m$ lines contain pairs of integers. The $i$-th line contains numbers $q_i, w_i$ ($1 \leq q_i, w_i \leq 10^5$).

It is guaranteed that all $q_i$ are distinct.

## Output

In a single line print maximum amount of money (in rubles) Sereja can pay.

Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout
streams or the %I64d specifier.

## Examples

### Example 1

**Input:**
```
5 2
1 2
2 3
```

**Output:**
```
5
```

### Example 2

**Input:**
```
100 3
1 2
2 1
3 1
```

**Output:**
```
4
```

### Example 3

**Input:**
```
1 2
1 1
2 100
```

**Output:**
```
100
```

## Note

In the first sample Sereja can pay 5 rubles, for example, if Dima constructs the following array: `[1, 2, 1, 2, 2]`.
There are another optimal arrays for this test.

In the third sample Sereja can pay 100 rubles, if Dima constructs the following array: `[2]`.



### ideas
1. 如果只有两个不同的数，肯定满足条件
2. 如果是3个数呢？ a, b, c ， c, a, b, c (至少要4个位置)
3. 如果是4个数呢？ a, b, c, d. n个节点， 组成一个完全图，需要n * (n - 1) / 2 条边，每个边要访问一次的话，n * (n - 1) / 2个节点
4. 如果 n - 1 是偶数， 那么需要 1 + n * (n - 1) / 2个节点（欧拉回路）
5. 如果 n - 1 是奇数？这个不大行，
6. 选择的数字，必须组成一个完全图，且其中边的数量 = n - 1
7. 如果存在奇数i, 满足 1 + i * (i - 1) / 2 <= n 
8. 多出来的点，缺少的那些点，使用已有的添加在后面就可以了
9. 