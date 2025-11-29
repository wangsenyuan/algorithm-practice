# Problem G

You are given an array a consisting of positive integers and q queries to this array. There are two types of queries:

1 l r x — for each index i such that l ≤ i ≤ r set ai = x.
2 l r — find the minimum among such ai that l ≤ i ≤ r.

We decided that this problem is too easy. So the array a is given in a compressed form: there is an array b consisting of n elements and a number k in the input, and before all queries a is equal to the concatenation of k arrays b (so the size of a is n·k).

## Input

The first line contains two integers n and k (1 ≤ n ≤ 105, 1 ≤ k ≤ 104).

The second line contains n integers — elements of the array b (1 ≤ bi ≤ 109).

The third line contains one integer q (1 ≤ q ≤ 105).

Then q lines follow, each representing a query. Each query is given either as 1 l r x — set all elements in the segment from l till r (including borders) to x (1 ≤ l ≤ r ≤ n·k, 1 ≤ x ≤ 109) or as 2 l r — find the minimum among all elements in the segment from l till r (1 ≤ l ≤ r ≤ n·k).

## Output

For each query of type 2 print the answer to this query — the minimum on the corresponding segment.

## Examples

### Example 1

**Input:**

```text
3 1
1 2 3
3
2 1 3
1 1 2 4
2 1 3
```

**Output:**

```text
1
3
```

### Example 2

**Input:**

```text
3 2
1 2 3
5
2 4 4
1 4 4 5
2 4 4
1 1 6 1
2 6 6
```

**Output:**

```text
1
5
1
```


### ideas
1. 如果没有压缩，那么就是segment tree with lazy update 就可以了
2. 现在有压缩的情况下，就很难处理了。首先肯定没法展开成完全形态
3. 考虑更新区间[l...r], 如果l...r很大，那么它必然包含了很多完整的压缩后的区间
4. 那么这些区间还是可以当作单个点去处理
5. 可以假装解压，然后重新根据用到的点，进行重压缩就可以了
6. 不对， 假设k = 3, b = [1, 2, 3, 4]
7. 那么 a = [1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4]
8. a先不管，就在query的(压缩后)区间上处理
9. 然后对于query，找到一个结果，还需要知道，在l...r区间中，有没有未被覆盖到的部分
10. 比如对于上面的a, 如果查询[1, 5]， 但是前面的更新，只覆盖到了[2..4], 那么a[1], a[5]也需要被挑出来
11. 