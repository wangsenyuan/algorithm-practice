# Problem E: Inversions After Swaps

## Problem Statement

There is an infinite sequence consisting of all positive integers in the increasing order: $p = \{1, 2, 3, \ldots\}$. We performed $n$ swap operations with this sequence. A `swap(a, b)` is an operation of swapping the elements of the sequence on positions $a$ and $b$. 

Your task is to find the number of inversions in the resulting sequence, i.e. the number of such index pairs $(i, j)$, that $i < j$ and $p_i > p_j$.

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 10^5$) — the number of swap operations applied to the sequence.

Each of the next $n$ lines contains two integers $a_i$ and $b_i$ ($1 \leq a_i, b_i \leq 10^9$, $a_i \neq b_i$) — the arguments of the swap operation.

## Output

Print a single integer — the number of inversions in the resulting sequence.

## Examples

### Example 1
**Input:**
```
2
4 2
1 4
```

**Output:**
```
4
```

### Example 2
**Input:**
```
3
1 6
3 4
2 5
```

**Output:**
```
15
```

## Note

In the first sample the sequence is being modified as follows: $\{1, 2, 3, 4\} \rightarrow \{1, 4, 3, 2\} \rightarrow \{2, 4, 3, 1\}$. It has 4 inversions formed by index pairs $(1, 4)$, $(2, 3)$, $(2, 4)$ and $(3, 4)$.


### ideas
1. 考虑这样一种操作 (1, a), => a.2 .....a - 1, 1, a + 1
2. 然后(1, b) => a.2..1...a-1.b.a-1
3. 那么1的前面比它大的数，就包括a, 2.... b - 1 
4. a没有问题，但是2...b-1因为没有出现，就麻烦了～
5. 所以，还需要知道1现在的位置，假设是pos，正常情况下，它应该在位置i,
6. 如果i < pos 那么它i+1....pos-1 的数都比它大
7. 如果这个中间出现了x个数（不管具体跑哪去了，都排除掉，剩下没有出现的，都要计入）
8. 