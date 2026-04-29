# C1. Maple and Tree Beauty (Easy Version)

Maple is given a rooted tree consisting of `n` vertices numbered from `1` to `n`, where the root has index `1`. Each vertex of the tree is labeled either zero or one. Unfortunately, Maple forgot how the vertices are labeled and only remembers that there are exactly `k` zeros and `n - k` ones.

For each vertex, we define the **name** of the vertex as the binary string formed by concatenating the labels of the vertices from the root to that vertex. More formally, `name_1 = label_1` and `name_u = name_{p_u} + label_u` for all `2 ≤ u ≤ n`, where `p_u` is the parent of vertex `u` and `+` means string concatenation.

The **beauty** of the tree is the length of the **longest common subsequence** (see note below) of the names of **all leaves** (vertices with no children). Determine the **maximum beauty** over all labelings of the tree with exactly `k` zeros and `n - k` ones.

> **\*** A sequence `a` is a subsequence of a sequence `b` if `a` can be obtained from `b` by deleting zero or more elements from arbitrary positions. The longest common subsequence of strings `s_1, s_2, …, s_m` is the longest string that is a subsequence of every `s_i`.

> **†** A leaf is any vertex without children.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 50`). The description of the test cases follows.

For each test case:

- The first line contains two integers `n` and `k` (`2 ≤ n ≤ 1000`, `0 ≤ k ≤ n`) — the number of vertices and the number of vertices labeled with zero, respectively.
- The second line contains `n - 1` integers `p_2, p_3, …, p_n` (`1 ≤ p_i ≤ i - 1`) — the parent of vertex `i`.

There is **no** constraint on the sum of `n` over all test cases.

## Output

For each test case, output a single integer: the maximum beauty among all labelings with exactly `k` zeros and `n - k` ones.

## Examples

### Example 1

**Input**

```text
5
7 3
1 1 2 2 3 3
7 2
1 1 2 3 1 1
5 0
1 2 3 4
5 2
1 1 1 1
5 4
1 1 1 1
```

**Output**

```text
3
2
5
1
2
```

### Example 2

**Input**

```text
5
2 0
1
2 1
1
3 0
1 1
3 1
1 2
3 1
1 1
```

**Output**

```text
2
2
2
3
2
```

## Note

In the first test case of **Example 1** (`7 3`), the maximum beauty is `3` when the vertices are labeled with `[0, 0, 0, 1, 1, 1, 1]`; the longest common subsequence is `001`.

In the second test case of **Example 1** (`7 2`), the maximum beauty is `2` when the vertices are labeled with `[1, 0, 0, 1, 1, 1, 1]`; the longest common subsequence is `11`.


### ideas
1. 为n个点分配0/1，（k个0，n-k个1），使的叶子节点的name的子序列有最长的相同的字符
2. 假设这个长度是w <= 最短的路径