# Problem Description

For each string s consisting of characters '0' and '1' one can define four integers a00, a01, a10 and a11, where axy is the number of subsequences of length 2 of the string s equal to the sequence {x, y}.

In this problem you are given four integers a00, a01, a10, a11 and have to find any non-empty string s that matches them, or determine that there is no such string. One can prove that if at least one answer exists, there exists an answer of length no more than 1 000 000.

## Input

The only line of the input contains four non-negative integers a00, a01, a10 and a11. Each of them doesn't exceed 10^9.

## Output

If there exists a non-empty string that matches four integers from the input, print it in the only line of the output. Otherwise, print "Impossible". The length of your answer must not exceed 1 000 000.

## Examples

### Example 1

**Input:**

```text
1 2 3 4
```

**Output:**

```text
Impossible
```

### Example 2

**Input:**

```text
1 2 2 1
```

**Output:**

```text
0110
```


### ideas
1. subsequence 表示非连续的
2. a00 = x, 那么 0的个数 = w，w * (w - 1) / 2 = x
3. a11 = y, 那么 同理可以计算出1的个数v
4. 有w个0和v个1，那么 01的数量 >= 0 (前面全部1，后面全部0) <= w * v (前面全部0，后面全部1)
5. 假设 0101的pattern，且每段的数量 =  a[0], a[1], a[2], ...
6. 感觉上应该每段要均匀
7. 