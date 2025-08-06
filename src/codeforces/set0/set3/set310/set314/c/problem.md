# Problem Description

Sereja has a sequence that consists of n positive integers, a₁, a₂, ..., aₙ.

First Sereja took a piece of squared paper and wrote all distinct non-empty non-decreasing subsequences of sequence a. Then for each sequence written on the squared paper, Sereja wrote on a piece of lines paper all sequences that do not exceed it.

A sequence of positive integers x = x₁, x₂, ..., xᵣ doesn't exceed a sequence of positive integers y = y₁, y₂, ..., yᵣ, if the following inequation holds: x₁ ≤ y₁, x₂ ≤ y₂, ..., xᵣ ≤ yᵣ.

Now Sereja wonders, how many sequences are written on the lines piece of paper. Help Sereja, find the required quantity modulo 1000000007 (10⁹ + 7).

## Input

The first line contains integer n (1 ≤ n ≤ 10⁵). The second line contains n integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁶).

## Output

In the single line print the answer to the problem modulo 1000000007 (10⁹ + 7).

## Examples

### Example 1
**Input:**
```
1
42
```
**Output:**
```
42
```

### Example 2
**Input:**
```
3
1 2 2
```
**Output:**
```
13
```

### Example 3
**Input:**
```
5
1 2 3 4 5
```
**Output:**
```
719
```

### ideas
1. [1, 2, 2]
2. [1], [2], [1, 2], [2, 2], [1, 2, 2]
3. [1], [1,2], [11, 12], [11,12,21,22], [111, 112, 122, 121] 4
4. 怎么得到unique的 subsequence 呢？
5. 如果v现在出现在了序列里面，那么它就可以贡献v种可能性(1, 2, ...v)
6. 那么它可以出现在多少个序列中呢？
7. 