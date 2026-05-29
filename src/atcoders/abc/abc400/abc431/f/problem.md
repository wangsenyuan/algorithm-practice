# Problem F - Almost Sorted 2

https://atcoder.jp/contests/abc431/tasks/abc431_f

**Time Limit:** 2 sec / **Memory Limit:** 1024 MiB

**Score:** 500 points

## Problem Statement

You are given an integer sequence
$A=(A_1,A_2,\ldots,A_N)$ of length $N$ and a positive integer $D$.

Find the number, modulo $998244353$, of integer sequences
$B=(B_1,B_2,\ldots,B_N)$ that can be obtained by rearranging $A$ and satisfy the
following condition:

- $B_{i+1} \geq B_i - D$ holds for all $i$ $(1 \leq i \leq N-1)$.

## Constraints

- $2 \leq N \leq 2 \times 10^5$
- $1 \leq D \leq 10^6$
- $1 \leq A_i \leq 10^6$
- All input values are integers.

### ideas
1. a升序排列；(这种情况肯定是一个有效的排列)
2. 然后依次处理；对于a[i], 它能够前移到位置j, 要满足, b[j+1] >= a[i] - d
3. a[i] >= b[j-1] - d => b[j] > a[i] + d 的最大的j,
4. * （i-j) ?

### summary

Sort `A` first, then process values from small to large. Think of inserting the
current value `v = a[r]` into a valid arrangement of the previous sorted prefix.

If `v` is placed immediately before some smaller value `x`, the condition requires
`x >= v - D`. Therefore, `v` cannot be inserted before values smaller than
`v - D`. Maintain a left pointer `l` such that `a[l] >= v - D`; then all insertion
positions from `l` through `r` are valid, giving `r - l + 1` choices for `a[r]`.

The two-pointer loop advances `l` while `a[l] < v - D`, so the valid window is
maintained in linear time after sorting. Multiply the answer by the number of
choices for each `r`.

Equal values are indistinguishable in the final integer sequence, but the insertion
process counts them as separate indexed elements. For each block of equal values
with size `cnt`, divide the answer by `cnt!` modulo `998244353`.

The total complexity is `O(N log N)` due to sorting, with an `O(N)` sweep afterward.
