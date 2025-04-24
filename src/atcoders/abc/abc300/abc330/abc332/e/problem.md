# Lucky Bag Division Problem

## Problem Description

There are N items remaining in the company. The weight of the i-th item (1≤i≤N) is W[i].

Takahashi will sell these items as D lucky bags. He wants to minimize the variance of the total weights of the items in the lucky bags.

## Variance Definition

The variance V is defined as:

\[
V = \frac{1}{D} \sum_{i=1}^{D} (x_i - \bar{x})^2
\]

where:
- x[1], x[2], ..., x[D] are the total weights of the items in the lucky bags
- \(\bar{x} = \frac{1}{D}(x_1 + x_2 + ... + x_D)\) is the average of x[1], x[2], ..., x[D]

## Rules
- It is acceptable to have empty lucky bags (in which case the total weight of the items in that bag is defined as 0)
- Each item must be in exactly one of the D lucky bags
- The goal is to find the minimum possible variance when dividing the items optimally

## Constraints
- 2 ≤ D ≤ N ≤ 15
- 1 ≤ W[i] ≤ 10^8
- All input values are integers

### ideas
1. avg 是确定的 = sum(a) / D
2. (xi - avg) ** 2 = xi ** 2 + avg ** 2 - 2 * avg * xi
3. xi ** 2 - 2 * avg * xi = xi * (xi - 2 * avg)
4. xi要尽量和avg接近
5. 不失一般性，假设 xi是非递减的
6. 一个猜想是 x[d] - x[1]越小越好
7. 所以是否存在2分的可能性呢？主要是这个序列不是连续的。也就是说如果差值为2能够得到，不代表3一定能得到
8. 