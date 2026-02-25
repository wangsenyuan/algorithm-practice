# Problem

A non-empty string $s$ is called **binary** if it consists only of characters `0` and `1`. Number the characters of binary string $s$ from 1 to the string's length and denote the $i$-th character by $s_i$.

A binary string $s$ of length $n$ is **periodical** if there exists an integer $1 \le k < n$ such that:

- $k$ is a divisor of $n$
- For all $1 \le i \le n - k$: $s_i = s_{i+k}$

For example, binary strings `101010` and `11` are periodical; `10` and `10010` are not.

A positive integer $x$ is **periodical** if its binary representation (without leading zeroes) is a periodical string.

**Task:** Count how many periodical numbers are in the interval $[l, r]$ (both ends included).

## Input

A single line with two integers $l$ and $r$ ($1 \le l \le r \le 10^{18}$), separated by a space.

> For C++: do not use `%lld` for 64-bit integers; use `cin`/`cout` or `%I64d`.

## Output

One integer: the number of periodical numbers in $[l, r]$.

## Examples

**Input**

```
1 10
```

**Output**

```
3
```

**Input**

```
25 38
```

**Output**

```
2
```

## Note

- First sample: periodical numbers are 3, 7, and 10.
- Second sample: periodical numbers are 31 and 36.


### ideas
1. 数位dp
2. dp[0/1][i] 表示到目前为止是否小于等于num, 且当前是第i位个位置（加上剩余的就能知道n）
3. 如果(i+1) 能整除n，且小于num, 那么可以把这个数当作循环节
4. 但是这里存在一个问题（就是似乎有重复的问题？）如果(i+1)本身就已经重复了，就不大行
5. dp[0/1][i] 如果目前小于num， 且n % (i+1) = 0
6. 那么把 dp[0][i] + fp[i+1] (就是那个inclusion/exclusion)的方式？
7. 是不是应该反过来计算，有多少不是周期性的数字？
8. 对于固定的长度m, 一共有 2 ^^ m 个数字
9. 有多少个是周期性的呢？
10. 