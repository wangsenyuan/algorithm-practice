# Problem B: Maximum Remainder

## Problem Description

You are given a sequence $a$ consisting of $n$ integers. Find the maximum possible value of $(a_i \bmod a_j)$, where $1 \leq i, j \leq n$ and $a_i \geq a_j$.

## Input

- The first line contains integer $n$ — the length of the sequence $(1 \leq n \leq 2 \cdot 10^5)$.
- The second line contains $n$ space-separated integers $a_i$ $(1 \leq a_i \leq 10^6)$.

## Output

Print the answer to the problem.

## Examples

### Example 1

**Input:**
```
3
3 4 5
```

**Output:**
```
2
```

## Notes

- The modulo operation $\bmod$ returns the remainder after division
- We need to find the maximum remainder when dividing any element by any other element (where the dividend is greater than or equal to the divisor)
- In the example: $3 \bmod 4 = 3$, $3 \bmod 5 = 3$, $4 \bmod 5 = 4$, $5 \bmod 3 = 2$, $5 \bmod 4 = 1$, $4 \bmod 3 = 1$. The maximum is $2$.

## ideas
1. sort a firstly
2. 假设 x是y的一个因子(y % x = 0)
3. 那么 y % (x + 1) = ?
4. n * x = y
5. (n - 1) * x = y - x
6. (n - 1) * (x + 1) = (n - 1) * x + (n - 1) = y - x + n - 1
7. 也就是说 y % (x + 1) = y/x - (1 + x )
8. x + 2 呢？
9. (n - 1) * (x + 2) = (n - 1) * x + n - 2 = y - x + n - 2
10. 那么似乎 x+1是最优的选择，因为y/x是相同的，那么(x+1)越小越好
11. 那么就好办了
12. 对于a[j]来说，找到它的因数，然后看x + 1是否存在，然后就可以就算出来了
13. 但是不够快～～
14. 对于x来说， 2 * x, 3 * x, 4 * x 都是它的倍数，所以它们 y % x = 0
15. 如果存在某个 y = n * x + r那么它求余x的话，应该是怎么样的？
16. y % x = r % x (r < x)的
17. 所以 就等于 r
18. 那么也就是反过来，就是把 x的倍数给去掉，那么这样子，似乎是logn * logn de 