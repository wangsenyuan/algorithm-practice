# Problem B: Lucky Number

## Problem Description

Bike loves looking for the second maximum element in the sequence. The second maximum element in the sequence of distinct numbers $x_1, x_2, ..., x_k$ ($k > 1$) is such maximum element $x_j$, that the following inequality holds: $x_j < \max(x_1, x_2, ..., x_k)$.

The **lucky number** of the sequence of distinct positive integers $x_1, x_2, ..., x_k$ ($k > 1$) is the number that is equal to the bitwise excluding OR of the maximum element of the sequence and the second maximum element of the sequence.

You've got a sequence of distinct positive integers $s_1, s_2, ..., s_n$ ($n > 1$). Let's denote sequence $s_l, s_{l+1}, ..., s_r$ as $s[l..r]$ ($1 \leq l < r \leq n$). Your task is to find the maximum number among all lucky numbers of sequences $s[l..r]$.

**Note:** As all numbers in sequence $s$ are distinct, all the given definitions make sense.

## Input

- The first line contains integer $n$ ($1 < n \leq 10^5$)
- The second line contains $n$ distinct integers $s_1, s_2, ..., s_n$ ($1 \leq s_i \leq 10^9$)

## Output

Print a single integer — the maximum lucky number among all lucky numbers of sequences $s[l..r]$.

## Examples

### Example 1
**Input:**
```
5
5 2 1 4 3
```

**Output:**
```
7
```

**Explanation:** You can choose $s[4..5] = \{4, 3\}$ and its lucky number is $(4 \oplus 3) = 7$. You can also choose $s[1..2]$.

### Example 2
**Input:**
```
5
9 8 3 5 7
```

**Output:**
```
15
```

**Explanation:** You must choose $s[2..5] = \{8, 3, 5, 7\}$.

## Notes

- The bitwise XOR operation ($\oplus$) is used to calculate the lucky number
- All numbers in the input sequence are distinct
- You need to find the maximum lucky number among all possible subarrays $s[l..r]$


### ideas
1. 像绕口令一样的题目～～～
2. lucky number是指一个序列中，最大值和第二大的值的 xor
3. 对于给定的序列，找出所有连续子序列的lucky number，最终输出这些lucky number中的最大值
4. 假设有个最大值的最高位h， 如果不是所有的位，这一位都set了
5. 那么最后结果中，h必然是set的
6. 假设有一个区间 (l...r) (很容易能够找到最大值和第二大的值)
7. 有两个最高位都设置过的区间， 不能放在一起
8. 对于l，l的最低位置可以通过最高位找到 (也就是说, 在给定r的时候， 最大值是确定的 = 左边最靠近的， h被设置的值)
9. l增加，那么第二大的值，会不断减小
10. 然后看i前面，l...i组成一个递减序列（如果是递增序列的话，没有意义，good(l, r) = good(l+1, r) 了）
11. 然后找到 和 a[i]不一样最高位的位置 比如 a[i] = 1111001 如果左边 1110。。。
12. 这里可以设定i = r, 再反过来计算一遍
13.  如果 a[r]的最高位和最大值一样，然后进行处理，它左边形成一个降序序列 （它们是第二大的值）
14.  然后找 a[r] 和这个序列的最优解
15.  可以用trie