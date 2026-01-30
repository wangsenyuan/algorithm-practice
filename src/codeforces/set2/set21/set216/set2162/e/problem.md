We call an array $[b_1, b_2, \ldots, b_m]$ of length $m$ palindromic if the following condition holds:

$$b_i = b_{m - i + 1} \text{ for all } 1 \le i \le m$$

In other words, an array is palindromic if it reads the same forward and backward.

You are given an array $[a_1, a_2, \ldots, a_n]$ of $n$ integers where $1 \le a_i \le n$ and an integer $k$.

You are required to perform the following operation exactly $k$ times:

- choose an integer $x$ such that $1 \le x \le n$,
- append $x$ to the end of the array $a$.

Your goal is to perform these $k$ operations in such a way that the number of palindromic subarrays* in the resulting array is minimized.

Output the $k$ integers you chose for each operation, in the order they were appended.

*An array $b$ is a subarray of an array $a$ if $b$ can be obtained from $a$ by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. In particular, an array is a subarray of itself.

## Input

The first line of input contains a single integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

The first line of each test case contains two integers $n$ and $k$ ($3 \le n \le 2 \cdot 10^5$, $1 \le k \le n$) — the length of the array $a$.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le n$) — the elements of the array $a$.

It is guaranteed that the sum of $n$ over all the test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, print the $k$ integers chosen for the append operations, in the order they were appended, such that the total number of palindromic subarrays in the resulting array is minimized.

If there are multiple answers, you may output any one of them.

## Example

### Input
```
5
4 1
1 3 3 4
4 2
2 2 2 2
5 1
4 1 5 2 2
6 3
1 2 3 4 5 6
5 3
3 2 5 2 3
```

### Output
```
2
1 3
3
3 4 1
4 1 5
```

## Note

For the first test case, if we append $2$ to the end of the array, $a$ becomes $[1, 3, 3, 4, 2]$. Now $a$ has only 6 palindromic subarrays — $[1]$, $[3]$, $[3]$, $[4]$, $[2]$, $[3, 3]$.
