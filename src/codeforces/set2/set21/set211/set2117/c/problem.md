# Problem C: Cool Partition

## Problem Description

Yousef has an array $a$ of size $n$. He wants to partition the array into one or more contiguous segments such that each element $a_i$ belongs to exactly one segment.

A partition is called **cool** if, for every segment $b_j$, all elements in $b_j$ also appear in $b_{j+1}$ (if it exists). That is, every element in a segment must also be present in the segment following it.

### Example

If $a = [1, 2, 2, 3, 1, 5]$, a cool partition Yousef can make is:
- $b_1 = [1, 2]$
- $b_2 = [2, 3, 1, 5]$

This is a cool partition because every element in $b_1$ (which are $1$ and $2$) also appears in $b_2$.

In contrast, $b_1 = [1, 2, 2]$, $b_2 = [3, 1, 5]$ is **not** a cool partition, since $2$ appears in $b_1$ but not in $b_2$.

### Important Notes

- After partitioning the array, you do not change the order of the segments.
- If an element appears several times in some segment $b_j$, it only needs to appear at least once in $b_{j+1}$.

## Task

Your task is to help Yousef by finding the **maximum number of segments** that make a cool partition.

## Input

The first line of the input contains integer $t$ $(1 \leq t \leq 10^4)$ — the number of test cases.

For each test case:
- The first line contains an integer $n$ $(1 \leq n \leq 2 \cdot 10^5)$ — the size of the array.
- The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(1 \leq a_i \leq n)$ — the elements of the array.

**Constraints:**
- It is guaranteed that the sum of $n$ over all test cases doesn't exceed $2 \cdot 10^5$.

## Output

For each test case, print one integer — the maximum number of segments that make a cool partition.

## ideas
1. b[j] 肯定包含 b[j-1], 那么也肯定包含 b[1...j-1] 中间出现的所有的数，而且越早结束越好。这样子，后面的余地更多
2. 所以 b[1] = a[1]