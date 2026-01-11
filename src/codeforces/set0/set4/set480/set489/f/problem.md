An n × n square matrix is special, if:

- it is binary, that is, each cell contains either a 0, or a 1;
- the number of ones in each row and column equals 2.

You are given n and the first m rows of the matrix. Print the number of special n × n matrices, such that the first m rows coincide with the given ones.

As the required value can be rather large, print the remainder after dividing the value by the given number mod.

## Input

The first line of the input contains three integers n, m, mod (2 ≤ n ≤ 500, 0 ≤ m ≤ n, 2 ≤ mod ≤ 10^9). 

Then m lines follow, each of them contains n characters — the first rows of the required special matrices. Each of these lines contains exactly two characters '1', the rest characters are '0'. Each column of the given m × n table contains at most two numbers one.

## Output

Print the remainder after dividing the required value by number mod.

## Examples

### Example 1

**Input:**
```
3 1 1000
011
```

**Output:**
```
2
```

### Example 2

**Input:**
```
4 4 100500
0110
1010
0101
1001
```

**Output:**
```
1
```

## Note

For the first test the required matrices are:

```
011
101
110
```

```
011
110
101
```

In the second test the required matrix is already fully given, so the answer is 1.

## Ideas

1. 行的顺序没有关系
2. 从左到右处理，dp[c][i][j] 表示到c为止，其中有i行已经放置了1，有j行放置了0, (n - i - j 行已经放置了2)
3. 现在在第c列，如果它已经放置了2，那么根据所在行的情况，更新i和j，然后进入下一列
4. 如果这一列放置了1，可以到一个状态，也可以选择一个放置1（放置1，要么在i中变化，要么在j中变化）
5. 
