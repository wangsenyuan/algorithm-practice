# Symmetric and Transitive

## 题意

统计大小为 n 的集合上满足以下条件的二元关系 ρ 的数量（模 10^9 + 7）：

- **对称**：若 (a, b) ∈ ρ，则 (b, a) ∈ ρ
- **传递**：若 (a, b) ∈ ρ 且 (b, c) ∈ ρ，则 (a, c) ∈ ρ
- **非自反**：至少存在一个元素 a 使得 (a, a) ∉ ρ

## 结构分析

若 a ρ b，由对称得 b ρ a，再由传递得 a ρ a。因此参与关系的元素必然自反，非自反要求至少有一个**孤立元素**（不参与任何关系）。

将 n 个元素分为：**孤立集合**（非空）和**参与关系的集合**。参与关系的元素形成若干等价类，等价类划分数 = 贝尔数 (Bell number)。

## 公式

\[
\text{答案} = \sum_{j=1}^{n} \binom{n}{j} \cdot B_{n-j}
\]

其中 j 为孤立元素个数，\(B_k\) 为 k 个元素的贝尔数。

## 算法

贝尔数递推：\(B_{n+1} = \sum_{k=0}^{n} \binom{n}{k} B_k\)。复杂度 O(n²)。

## Input

A single line contains a single integer n (1 ≤ n ≤ 4000).

## Output

In a single line print the answer to the problem modulo 10^9 + 7.

## Examples

### Example 1

**Input:**

```
1
```

**Output:**

```
1
```

### Example 2

**Input:**

```
2
```

**Output:**

```
3
```

### Example 3

**Input:**

```
3
```

**Output:**

```
10
```

## Note

- n = 1：唯一关系为空关系 ρ = ∅
- n = 2：三种关系为 ρ = ∅，ρ = {(x, x)}，ρ = {(y, y)}
