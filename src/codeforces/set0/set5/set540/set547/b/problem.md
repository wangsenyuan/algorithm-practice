# Problem Description

Mike is the president of country What-The-Fatherland. There are n bears living in this country besides Mike. All of them are standing in a line and they are numbered from 1 to n from left to right. The i-th bear is exactly ai feet high.

A **group of bears** is a non-empty contiguous segment of the line. The **size** of a group is the number of bears in that group. The **strength** of a group is the minimum height of the bear in that group.

Mike is curious to know for each x such that 1 ≤ x ≤ n, the maximum strength among all groups of size x.

## Input

- The first line of input contains integer **n** (1 ≤ n ≤ 2 × 10⁵), the number of bears.
- The second line contains **n** integers separated by space, a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁹), heights of bears.

## Output

Print **n** integers in one line. For each x from 1 to n, print the maximum strength among all groups of size x.

## Examples

### Example 1

**Input:**
```
10
1 2 3 4 5 4 3 2 1 6
```

**Output:**
```
6 4 4 3 3 2 2 1 1 1
```

## Analysis

For each group size x, we need to find the maximum strength (minimum height) among all contiguous segments of length x.

- For x=1: The maximum strength is 6 (from the single bear with height 6)
- For x=2: The maximum strength is 4 (from segments like [4,5] or [5,4])
- For x=3: The maximum strength is 4 (from segment [4,5,4])
- And so on...

## Approach

This problem can be solved using a monotonic stack or deque approach to efficiently find the maximum minimum in sliding windows of different sizes.

For each possible group size k from 1 to n:
1. Find all contiguous segments of size k
2. For each segment, calculate its strength (minimum element)
3. Return the maximum strength among all segments of size k

The key insight is to use sliding window minimum technique with a deque to efficiently compute the minimum for each window size.