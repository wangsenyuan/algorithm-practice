# Problem D

## Description

Nastya received one more array on her birthday, this array can be used to play a traditional Byteland game on it. However, to play the game the players should first select such a subsegment of the array that $\frac{p}{s} = k$, where:
- $p$ is the product of all integers on the given array
- $s$ is their sum
- $k$ is a given constant for all subsegments

Nastya wonders how many subsegments of the array fit the described conditions. A subsegment of an array is several consecutive integers of the array.

## Input

The first line contains two integers $n$ and $k$ ($1 \leq n \leq 2 \cdot 10^5$, $1 \leq k \leq 10^5$), where $n$ is the length of the array and $k$ is the constant described above.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^8$) — the elements of the array.

## Output

In the only line print the number of subsegments such that the ratio between the product and the sum on them is equal to $k$.

## Examples

### Example 1
**Input:**
```
1 1
1
```

**Output:**
```
1
```

### Example 2
**Input:**
```
4 2
6 3 8 1
```

**Output:**
```
2
```

## Note

In the first example the only subsegment is $[1]$. The sum equals $1$, the product equals $1$, so it suits us because $\frac{1}{1} = 1$.

There are two suitable subsegments in the second example — $[6, 3]$ and $[3, 8, 1]$. 
- Subsegment $[6, 3]$ has sum $9$ and product $18$, so it suits us because $\frac{18}{9} = 2$.
- Subsegment $[3, 8, 1]$ has sum $12$ and product $24$, so it suits us because $\frac{24}{12} = 2$.

### ideas
1. 考虑 k = 1 的情况， a * b = a + b => a * (b - 1) = b
2. 如果b = 1, 显然不会有解。所以考虑 b > 1, 那么不会有解
3. 所以k=1的时候，答案 = n
4. 考虑k > 1的情况
5. prod = sum * k 
6. 但是可能存在，连续的1, 一个很大的数这样的配置。如果*brute force*就有很多无效的检查
7. 似乎也只有1是需要特殊处理的，如果没有1的情况下，连续60个2，肯定就足够了
8. 所以，把1的序列特殊拿出来，进行处理就好了
9. 假设都是1， 那么在这个序列中，不会产生符合条件的子序列
10. 假设目前是 p/s p/(s + x) = k 那么x是可以计算出来的，如果 x <= len(seq of 1)
11. 那么增加1， 然后将 s += len(seq), 继续处理就好了
12. 这样子不会超过60次检查 1e9 * 