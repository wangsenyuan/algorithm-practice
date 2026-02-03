Vasya is sitting in an extremely boring math class. To have fun, he took a piece of paper and wrote out $n$ numbers on a single line. After that, Vasya began to write out different ways to put pluses (`+`) in the line between certain digits so that the result was a correct arithmetic expression. Formally, no two pluses in such a partition can stand together (between any two adjacent pluses there must be at least one digit), and no plus can stand at the beginning or the end of a line.

For example, in the string `100500`, the ways `100500` (no pluses), `1+00+500`, or `10050+0` are correct, and the ways `100++500`, `+1+0+0+5+0+0`, or `100500+` are incorrect.

The lesson was long, and Vasya has written all the correct ways to place exactly $k$ pluses in a string of digits. At this point, he got caught having fun by a teacher and he was given the task to calculate the sum of all the resulting arithmetic expressions by the end of the lesson (when calculating the value of an expression the leading zeros should be ignored). As the answer can be large, Vasya is allowed to get only its remainder modulo $10^9 + 7$. Help him!

## Input

The first line contains two integers, $n$ and $k$ ($0 \le k < n \le 10^5$).

The second line contains a string consisting of $n$ digits.

## Output

Print the answer to the problem modulo $10^9 + 7$.

## Examples

### Input
```
3 1
108
```

### Output
```
27
```

### Input
```
3 2
108
```

### Output
```
9
```

## Note

In the first sample the result equals $(1 + 08) + (10 + 8) = 27$.

In the second sample the result equals $1 + 0 + 8 = 9$.


### ideas
1. 一共n个字符，那么有n-1个位置，所以有 C(n-1, k)种分割方式
2. 可以计算单个s[i]的贡献吗？
3. 完全没有想法， n * n * k 的算法肯定是不行的
4. 