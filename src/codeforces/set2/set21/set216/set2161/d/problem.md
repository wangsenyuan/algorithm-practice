# Problem Description

An array $b$ is good if there do not exist indices $1 \leq i < j \leq |b|$ such that $b_j - b_i = 1$.

You are given an integer array $a_1, a_2, \ldots, a_n$. Determine the minimum number of elements that need to be removed from the given array so that it becomes a good array.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 6 \cdot 10^4$). The description of the test cases follows.

The first line of each test case contains an integer $n$ ($1 \leq n \leq 3 \cdot 10^5$) — the number of elements in the array.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq n$) — the elements of the array.

It is guaranteed that the sum of $n$ over all test cases does not exceed $3 \cdot 10^5$.

## Output

For each test case, print a single integer — the minimum number of elements that need to be removed from the array to make it a good array.

## Example

### Example Input

```text
6
1
1
5
1 2 3 4 5
5
5 4 3 2 1
5
5 5 5 4 4
7
1 7 1 2 5 7 1
6
1 2 5 6 5 5
```

### Example Output

```text
0
2
0
0
1
2
```

## Note

In the first test case, the array is good from the very beginning, so nothing needs to be removed from it.

In the second test case, the optimal solution is to remove the second and fourth elements of the array, which will result in $1, 3, 5$ array.

In the third test case, the array is good from the very beginning.

In the fourth test case, the array is good from the very beginning.

In the fifth test case, the optimal solution is to remove the fourth element of the array, which will result in $1, 7, 1, 5, 7, 1$ array.

In the sixth test case, one of the optimal solutions is to remove the first and fourth elements of the array, which will result in $2, 5, 5, 5$ array.



### ideas
1. 搞不出来～～～
2. 1 2 3 不可以, 3 2 1是有可能的
3. 1.。。3 是可以的
4. 也就是说，如果是连续的数字
5. 必然是 5 4 3 2 1 这种结构
6. 如果v出现了，且v-1也出现了，那么v-1只能出现在v的后面
7. 假设 dp[v][1] 表示v出现时的计数 = dp[v-1][1] + ? 或者 = dp[v-1][0] + 全部的v
8. 第二个表示v不出现时的最优解
9. fp[i] 表示 a[i] = v, 且i就是v的最后一个位置时的最优解
10. fp[i] = fp[j] i < j 且a[j] = a[i] - 1,  + i前面v的个数 - j前面v-1的个数。 
11. 但是 1， 2， 3 这种结构怎么处理呢？
12. 2在被选中的时候1是不能选的，现在2不选择了, 那么1就可以被选择了
13. 所以还需要 + fp[?]