# Problem E

You are given an array $a$ of size $n$ and an integer $k$. You do the following procedure $k$ times:

For each element $a_i$, you set $a_i$ to $\text{mex}^*$($a_1, a_2, \ldots, a_{i-1}, a_{i+1}, a_{i+2}, \ldots, a_n$). In other words, you set $a_i$ to the mex of all other elements in the array. This is done for all elements in the array at the same time.

Please find the sum of elements in the array after all $k$ operations.

$^*$ The minimum excluded (MEX) of a collection of integers $d_1, d_2, \ldots, d_k$ is defined as the smallest non-negative integer $x$ which does not occur in the collection $d$.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line contains two integers $n$ and $k$ ($2 \leq n \leq 2 \cdot 10^5$, $1 \leq k \leq 10^9$) – the number of elements in $a$ and the number of operations done.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($0 \leq a_i \leq n$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output the sum of elements after all $k$ operations on a new line.

## Example

**Input:**
```
5
3 3
0 2 1
2 4
0 2
4 1
0 0 1 1
8 7
6 6 2 4 3 0 1 8
2 2
0 0
```

**Output:**
```
3
1
8
25
0
```

## Note

In the first test case, we performed the operation on the array $[0,2,1]$ three times. Let's compute the result after the first time:

- The first element becomes $\text{MEX}(2,1) = 0$
- The second element becomes $\text{MEX}(0,1) = 2$
- The third element becomes $\text{MEX}(0,2) = 1$

So, after the first operation, $[0,2,1]$ becomes $[0,2,1]$ again. It can be shown that the array will not change no matter how many times we perform the operation, so the final array after three operations is still $[0,2,1]$. The sum is $0+2+1=3$.

In the third test case, the array becomes $[2,2,2,2]$.