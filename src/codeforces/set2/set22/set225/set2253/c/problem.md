# C. Sum of Distinct Values in a Matrix

[Problem link](https://codeforces.com/problemset/problem/2253/C)

**Contest:** [Educational Codeforces Round 193 (Rated for Div. 2)](https://codeforces.com/contest/2253)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

You are given a matrix with `n` rows and `m` columns. Initially, all of its elements are equal to zero.

You are also given two arrays of positive integers `a = [a_1, a_2, …, a_x]` and `b = [b_1, b_2, …, b_y]`. The elements in each array are strictly increasing.

You may perform any number of operations, possibly zero. Each operation is one of the following types:

- choose a number `c` from array `a` and a row of the matrix; set every element in this row to `c`;
- choose a number `d` from array `b` and a column of the matrix; set every element in this column to `d`.

The operations may be performed in any order. You may choose the same row, column, or value multiple times.

The cost of a matrix is the sum of all distinct numbers that occur in it at least once. Find the maximum possible cost of the matrix.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

The descriptions of the test cases follow.

The first line of each test case contains four integers `n`, `m`, `x`, and `y` (`1 <= n, m <= 10^5`, `1 <= x, y <= n + m`) — the number of rows, the number of columns, the length of array `a`, and the length of array `b`, respectively.

The second line of each test case contains `x` integers `a_1, a_2, …, a_x` (`1 <= a_1 < a_2 < … < a_x <= n + m`) — the elements of array `a`.

The third line of each test case contains `y` integers `b_1, b_2, …, b_y` (`1 <= b_1 < b_2 < … < b_y <= n + m`) — the elements of array `b`.

Additional constraints on the input:

- the sum of `n` over all test cases does not exceed `10^5`;
- the sum of `m` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer — the maximum possible cost of the matrix.

## Example

### Input

```text
7
1 3 3 3
1 2 3
1 2 3
2 2 2 2
1 4
2 3
2 2 1 1
1
1
4 1 1 5
5
1 2 3 4 5
1 1 2 2
1 2
1 2
7 2 9 1
1 2 3 4 5 6 7 8 9
9
9 9 12 12
1 3 4 6 7 9 10 12 13 15 16 18
2 3 5 6 8 9 11 12 14 15 17 18
```

### Output

```text
6
9
1
9
2
44
170
```

### Note

In the first test case, you can first assign `3` to the only row, and then assign `1` and `2` to the first and second columns. The matrix will contain `1`, `2`, and `3`, so its cost is `6`.

In the second test case, you can first assign `2` and `3` to the columns, and then assign `4` to the first row. The matrix will then contain `2`, `3`, and `4`, so its cost is `9`.


## ideas
1. 考虑一种策略, 先把最大的m个b设置到矩阵里面
2. 然后再放入最大的(n-1)个a, 可以在保留b的情况下, 得到(n-1)个不同的值
3. 这个不完全对, 因为两边会有重复的部分
4. 应该是找到最大的m + n - 1 个不同的数?
5. 不对, 因为a中的不能放到col中
6. 那是不是就是从大往下处理就好了?