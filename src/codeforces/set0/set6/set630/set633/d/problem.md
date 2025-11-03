Yash has recently learnt about the Fibonacci sequence and is very excited about it. He calls a sequence Fibonacci-ish if:

- the sequence consists of at least two elements
- $f_0$ and $f_1$ are arbitrary
- $f_{n+2} = f_{n+1} + f_n$ for all $n \geq 0$

You are given some sequence of integers $a_1, a_2, \ldots, a_n$. Your task is rearrange elements of this sequence in such a
way that its longest possible prefix is Fibonacci-ish sequence.

### Input

The first line of the input contains a single integer $n$ ($2 \leq n \leq 1000$) — the length of the sequence $a_i$.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($|a_i| \leq 10^9$).

### Output

Print the length of the longest possible Fibonacci-ish prefix of the given sequence after rearrangement.

### Examples

#### Example 1

**Input**
```
3
1 2 -1
```

**Output**
```
3
```

#### Example 2

**Input**
```
5
28 35 7 14 21
```

**Output**
```
4
```

### Note

In the first sample, if we rearrange elements of the sequence as $-1, 2, 1$, the whole sequence $a_i$ would be
Fibonacci-ish.

In the second sample, the optimal way to rearrange elements is $7, 14, 21, 35, 28$.


### ideas
1. 如果 a[i] + a[j] 这个是f序列中的，那么这个序列其实是确定的
2. dp[i][j] 表示以a[i], a[j]开始的序列的最长
3. a, b, a => b = 0, 如果全是0的，也是ok的
4. 0可以添加到所有序列的开始，也可以添加到两个相同数字的中间
5. a, 0, a, a, 2 * a, 3 * a, 5 * a, ...
6. 如果 a != 0, b != 0
7. a, b, a + b, a + 2 * b, 2 * a + 3 * b, 3 * a + 5 * b
8. a = b是可以的
9. 如果 a != b, 但是还是会出现 a + b = 0的情况
10. 但是x = f[i-1] * a + f[i] * b
11. 还有一个就是，这个序列增长会非常快，长度不会超过20？
12. 所以，其实是可以brute force的？
13. 