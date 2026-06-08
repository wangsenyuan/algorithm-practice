# D. Cute Subsequences

[Problem link](https://codeforces.com/problemset/problem/2079/D)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an array of `n` positive integers `a_1, a_2, ..., a_n`, and a
positive integer `k`.

You need to divide the array into `k` non-empty subsequences such that each
element of the array belongs to exactly one subsequence.

A subsequence is a sequence that can be obtained from another sequence by
deleting some elements without changing the order of the remaining elements.

Suppose the `i`-th subsequence contains elements with indices:

```text
j_1 < j_2 < ... < j_l
```

The value of this subsequence is:

```text
max(a[j_m] + m) for all 1 <= m <= l
```

The cost of a division is the sum of the values of all `k` subsequences.

Find the maximum possible cost.

## Input

The first line contains two positive integers `n` and `k`
(`1 <= k <= n <= 500000`) — the size of the array and the number of
subsequences.

The second line contains `n` positive integers `a_1, a_2, ..., a_n`
(`1 <= a_i <= 10^9`) — the elements of the array.

## Output

Output the maximum cost of dividing the array into `k` non-empty subsequences.

## Example

### Input

```text
5 3
3 7 10 1 2
```

### Output

```text
24
```

## Note

In the sample test, one valid division is:

```text
[3, 10], [7], [1, 2]
```

The cost is:

```text
(10 + 2) + (7 + 1) + (2 + 2) = 12 + 8 + 4 = 24
```

## Scoring

The tests consist of six groups. Points for each group are awarded only if all
tests of the group and all required groups are passed.

- Group 0: examples.
- Group 1: `n <= 8`.
- Group 2: `k = 2`.
- Group 3: `a_{i+1} <= a_i`.
- Group 4: `a_{i+1} >= a_i - 1`.
- Group 5: `n <= 1000`.
- Group 6: full constraints.


### ideas
1. 假设到i为止的时候，已经得到了m个序列,sum(value)
2. 现在加入一个新的数v, 去替换前面的某个序列的u
3. 那么应该是 sum - u + v + 1 (+1是对应序列的长度, 但是成立吗？)
4. 那么这里要替换的就是最小的u
5. 这里不对，因为这里，假设了前m个，必须是属于不同的序列（但有可能这是错的）
6. 考虑最后选择的value所在的a中的位置p1, p2, ... pm, p1 < p2 ... < pm
7. 那么p1, p2, ...pm-1 都是对应序列的长度， pm不一定
8. 这是因为，p1如果不是b[1]的长度，假设它的后面还有一个元素，如果把它拿出来，给了b[2], 那么只会增加p2的位置，会得到更优的结果
9. 所以, p1肯定是a[1]的长度， 同理p2是b2的长度， until p[m-1], p[m-1]后面的都是b[m](但是p[m-1]前面也可能有b[m]的元素)
10. a[p1] + (len(b[1])) + a[p2] + len(b[2]) + ... + a[p[m-1]] + len(b[m-1]) + a[p[m]] + len(b[p[m]])
11. sum := len(b[1]) + len(b[2]) + ... + len(b[p[m]]) = n
12. a[p1] + a[p2] + ... a[p[m-1]] + a[p[m]] 的最大值？ 找到最大的m个值？不对
13. sum 不是等于n, 它等于 p[m], 所以就是选定p[m]后，找到 m-1个最大值，计算最优解

## Solution summary

For each final division into `k` subsequences, consider the element that
determines the value of each subsequence.

If one subsequence has elements:

```text
j_1 < j_2 < ... < j_l
```

then its value is:

```text
max(a[j_m] + m)
```

So the chosen determining element contributes two parts:

- its array value `a[pos]`;
- its position inside that subsequence.

### Key structure

Assume the determining positions of the `k` subsequences are:

```text
p_1 < p_2 < ... < p_k
```

For the first `k-1` determining positions, we can make `p_i` exactly the end of
its subsequence prefix. If there were extra elements after `p_i` in that same
subsequence, moving those elements to a later subsequence would only increase
the later determining position and would not hurt the earlier value.

Therefore, when the last determining position is `p_k`, the total contribution
of all position-inside-subsequence terms is:

```text
p_k
```

using 1-based indexing.

So if we fix the last determining position `p`, the answer has the form:

```text
p + a[p] + best sum of k-1 values among a[1..p-1]
```

That is the main reduction.

### Scanning positions

The code scans the array from left to right. Suppose the current position is
`pos` with value `v = a[pos]`.

Treat this position as `p_k`, the last determining position. Then:

```go
tmp := tr.GetBest(k - 1)
best = max(best, tmp+v+(pos+1))
```

Here:

- `tmp` is the sum of the largest `k-1` values among previous elements;
- `v` is the value of the current last determining element;
- `pos + 1` is `p_k` in 1-based indexing.

After evaluating the current position, insert `v` into the data structure so it
can be used by later positions:

```go
tr.Set(i, v)
```

### Segment tree

The values `a_i` can be large, so the code first compresses them:

```go
arr := slices.Clone(a)
slices.Sort(arr)
arr = slices.Compact(arr)
```

The segment tree stores, for each compressed value range:

- `cnt`: how many inserted elements are in this range;
- `val`: the sum of those inserted elements.

`GetBest(k)` returns the sum of the largest `k` inserted values.

It works by always trying the right child first, because larger compressed values
are stored on the right side of the segment tree:

- if the right child has at least `k` elements, recurse only into the right
  child;
- otherwise, take all values from the right child and continue taking the
  remaining elements from the left child.

This gives the sum of the largest `k` previous values in `O(log n)`.

### Complexity

Sorting and coordinate compression cost:

```text
O(n log n)
```

Each array position performs one query and one update on the segment tree, both
in `O(log n)`.

Overall complexity:

```text
O(n log n)
```

Space complexity:

```text
O(n)
```
