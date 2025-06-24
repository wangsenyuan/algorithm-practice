# Problem B: Array Partition

## Problem Description

An array $b$ is called to be a **subarray** of $a$ if it forms a continuous subsequence of $a$, that is, if it is equal to $a_l, a_{l+1}, \ldots, a_r$ for some $l, r$.

Suppose $m$ is some known constant. For any array, having $m$ or more elements, let's define its **beauty** as the sum of $m$ largest elements of that array. For example:

- For array $x = [4, 3, 1, 5, 2]$ and $m = 3$, the 3 largest elements of $x$ are $5$, $4$ and $3$, so the beauty of $x$ is $5 + 4 + 3 = 12$.
- For array $x = [10, 10, 10]$ and $m = 2$, the beauty of $x$ is $10 + 10 = 20$.

You are given an array $a_1, a_2, \ldots, a_n$, the value of the said constant $m$ and an integer $k$. You need to split the array $a$ into exactly $k$ subarrays such that:

1. Each element from $a$ belongs to exactly one subarray.
2. Each subarray has at least $m$ elements.
3. The sum of all beauties of $k$ subarrays is maximum possible.

## Input

The first line contains three integers $n$, $m$ and $k$ ($2 \leq n \leq 2 \cdot 10^5$, $1 \leq m$, $2 \leq k$, $m \cdot k \leq n$) — the number of elements in $a$, the constant $m$ in the definition of beauty and the number of subarrays to split to.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($-10^9 \leq a_i \leq 10^9$).

## Output

In the first line, print the maximum possible sum of the beauties of the subarrays in the optimal partition.

In the second line, print $k-1$ integers $p_1, p_2, \ldots, p_{k-1}$ ($1 \leq p_1 < p_2 < \ldots < p_{k-1} < n$) representing the partition of the array, in which:

- All elements with indices from $1$ to $p_1$ belong to the first subarray.
- All elements with indices from $p_1 + 1$ to $p_2$ belong to the second subarray.
- $\ldots$
- All elements with indices from $p_{k-1} + 1$ to $n$ belong to the last, $k$-th subarray.

If there are several optimal partitions, print any of them.


### ideas
1. 将a分成k段，每段至少有m个数，且选择其中m个最大的数，作为score，求最大的score
2. 也就是说要舍弃掉n - m * k 个数
3. 一定能舍弃掉最小的这么多的数吗？似乎是可以的
4. 考虑就舍弃一个数，那么只要把这一个数放在某一段m+1中就可以了，也就是让它所在的段有m+1个数即可