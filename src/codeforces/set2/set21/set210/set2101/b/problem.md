# Problem Description

You are given a permutation $a$ of length $n$*. You are allowed to do the following operation any number of times (possibly zero):

Choose an index $1 \leq i \leq n-3$. Then, swap $a_i$ with $a_{i+2}$, and $a_{i+1}$ with $a_{i+3}$ simultaneously. In other words, permutation $a$ will be transformed from:
$$[..., a_i, a_{i+1}, a_{i+2}, a_{i+3}, ...]$$
to:
$$[..., a_{i+2}, a_{i+3}, a_i, a_{i+1}, ...]$$

Determine the lexicographically smallest permutation† that can be obtained by applying the above operation any number of times.

## Definitions

* A permutation of length $n$ is an array consisting of $n$ distinct integers from $1$ to $n$ in arbitrary order. For example:
  - $[2,3,1,5,4]$ is a permutation
  - $[1,2,2]$ is not a permutation (2 appears twice in the array)
  - $[1,3,4]$ is not a permutation ($n=3$ but there is 4 in the array)

† An array $x$ is lexicographically smaller than an array $y$ of the same size if and only if the following holds:
- In the first position where $x$ and $y$ differ, the array $x$ has a smaller element than the corresponding element in $y$.

## ideas
1. 每次操作，不会改变位置的奇偶性。
2. 所以可以单独考虑偶数位和奇数位的情况
3. 但是有个问题是， [1, 2, 3, 4] 要同时交换[3, 4, 1, 2]
4. 所以，还不能单独考虑， [1, 2], [3, 4] 要打包考虑。 只能按照第一位来考虑
5. 还有个问题是，不一定是[1, 2]配对， 也可以是[2, 3]配对 [2, 3, 4, 5]
6. 还有点麻烦的。应该看第一位能够取的最小的值是什么。（假设是x，那么这个交换肯定是把某个y一起带过来了）
7. 然后看第二位能够取的的最小值是什么（要它后面的部分，前面的不能用）
8. 所以要两个队列（改一个的时候，也要更新第二个）