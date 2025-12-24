# Problem Description

In this problem, you will be playing a game with Hongcow. How lucky of you!

Hongcow has a hidden $n$ by $n$ matrix $M$. Let $M_{i,j}$ denote the entry $i$-th row and $j$-th column of the matrix. The rows and columns are labeled from $1$ to $n$.

The matrix entries are between $0$ and $10^9$. In addition, $M_{i,i} = 0$ for all valid $i$. Your task is to find the minimum value along each row, excluding diagonal elements. Formally, for each $i$, you must find $\min_{j \neq i} M_{i,j}$.

To do this, you can ask Hongcow some questions.

A question consists of giving Hongcow a subset of distinct indices $\{w_1, w_2, \ldots, w_k\}$, with $1 \leq k \leq n$. Hongcow will respond with $n$ integers. The $i$-th integer will contain the minimum value of $\min_{1 \leq j \leq k} M_{i,w_j}$.

You may only ask Hongcow at most $20$ questions — he thinks you only need that many questions answered.

When you are ready to answer, print out a single integer $-1$ on its own line, then $n$ integers on the next line. The $i$-th integer should be the minimum value in the $i$-th row of the matrix, excluding the $i$-th element. Do not forget to flush the final answer as well. Printing the answer does not count as asking a question.

## Wrong Answer Conditions

You will get Wrong Answer verdict if:

- Your question or answers are not in the format described in this statement.
- You ask strictly more than $20$ questions.
- Your question contains duplicate indices.
- The value of $k$ in your question does not lie in the range from $1$ to $n$, inclusive.
- Your final answer is not correct.

You will get Idleness Limit Exceeded if you don't print anything or if you forget to flush the output, including for the final answer (more info about flushing output below).

## Input

The first line of input will contain a single integer $n$ ($2 \leq n \leq 1,000$).

## Output

To print the final answer, print out the string `-1` on its own line. Then, the next line should contain $n$ integers. The $i$-th integer should be the minimum value of the $i$-th row of the matrix, excluding elements on the diagonal. Do not forget to flush your answer!

## Interaction

To ask a question, print out a single integer $k$ on its own line, denoting the size of your subset. Then, the next line should contain $k$ integers $w_1, w_2, \ldots, w_k$. Note, you must flush your output to get a response.

Hongcow will respond by printing out a line with $n$ integers. The $i$-th integer in this line represents the minimum value of $M_{i,w_j}$ where $j$ is between $1$ and $k$.

You may only ask a question at most $20$ times, otherwise, you will get Wrong Answer.

To flush you can use (just after printing an integer and end-of-line):

- `fflush(stdout)` in C++;
- `System.out.flush()` in Java;
- `stdout.flush()` in Python;
- `flush(output)` in Pascal;
- See the documentation for other languages.

## Hacking

To hack someone, use the following format:

```text
n
M_{1,1} M_{1,2} ... M_{1,n}
M_{2,1} M_{2,2} ... M_{2,n}
...
M_{n,1} M_{n,2} ... M_{n,n}
```

Of course, contestant programs will not be able to see this input.

## Examples

### Example 1

**Input:**

```text
3
0 0 0
2 7 0
0 0 4
3 0 8
0 5 4
```

**Output:**

```text
3
1 2 3
1
3
2
1 2
1
2
1
1
-1
2 5 4
```

### Example 2

**Input:**

```text
2
0 0
0 0
```

**Output:**

```text
1
2
1
1
-1
0 0
```

## Note

In the first sample, Hongcow has the hidden matrix:

$$
\begin{bmatrix}
0 & 3 & 2 \\
5 & 0 & 7 \\
4 & 8 & 0
\end{bmatrix}
$$

Here is a more readable version demonstrating the interaction. The column on the left represents Hongcow, while the column on the right represents the contestant.

```text
3
              3
              1 2 3
0 0 0
              1
              3
2 7 0
              2
              1 2
0 0 4
              1
              2
3 0 8
              1
              1
0 5 4
              -1
              2 5 4
```

For the second sample, it is possible for off-diagonal elements of the matrix to be zero.


### ideas
1. 先理解一下问题，[w1, w2, .... wk]
2. jury回复n个数字（不是k个数字）
3. 第1个数字表示 min(M[1, w1], M[1, w2], ... M[1, w[k]])
4. 第2个数字表示 min(M[2, w2], M[2, w2], ... M[2, w[k]])
5. ...
6. 第i个数字表示 min(M[i, wi], M[i, w2], ....M[i, w[k]])
7. 如果给定的w, 包含了1， 那么 ans[1] = 0, 所以这个答案对第一行是无效的
8. 让W = [1, 2, ... n/2] 这个答案对前n/2行，是没有用的（它们的答案都是0）
9. 但是对后面n/2行是有意义的，就是知道了前n/2的最小值
10. 最好一个策略能够问道所有位的最小值， 且可以组合出来
11. 对于1来说，先问出偶数位的值，
12. [1, 5, 9, ...] [3, 7, 11, ...] d = 4
13. [1, 9, ...], [5, 13, ..], d = 8
14. 如果需要1的结果，那么就是 d = 2 + [a0 = 3, d = 4] + [a0 = 5, d + 8]
15. 如果需要3的结果， [a0 = 2, d = 2] + [a0 = 1, d = 4] + [a0 = 5, d + 8] + [a0 = 7, d + 8]
16. 好像是可以的
17. 