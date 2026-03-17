### Problem

The Smart Beaver from ABBYY became a screenwriter for a TV series. Among the main characters there are `n` single men and `n` single women.

An opinion poll showed that viewers like some specific couples. Each such couple `(h, w)` (man `h`, woman `w`) has an associated **delight value** `r` — how happy the audience is if those two marry. Any other (unlisted) couples leave the audience indifferent and are not allowed in the script.

You are given `k` possible couples, each as a triple `(h, w, r)`.

The script may include **several marriages** or none at all. A subset of these `k` marriages is considered **acceptable** if:

- Each man appears in **at most one** chosen marriage.
- Each woman appears in **at most one** chosen marriage.

In other words, an acceptable set is a **matching** between men and women, chosen from the given `k` edges. The **value** of an acceptable set is the sum of `r` over all marriages in the set.

There is a finite number of acceptable sets, each describing a variant of the script. The screenwriters do **not** want the one with maximum value (too predictable). Instead, they:

1. List all acceptable sets.
2. Sort them by increasing total value.
3. Choose the `t`‑th set in this order.

Note:

- `t = 1` corresponds to the empty set (no marriages, total value 0).
- `t = 2` is the single marriage with the smallest delight value, etc.

Your task: compute the **value** (total delight) of the `t`‑th acceptable set in this sorted order.

### Input

- First line: integers `n`, `k`, `t`  \n
  `1 ≤ k ≤ min(100, n^2)`, `1 ≤ t ≤ 2·10^5`\n
  (space-separated)
- Next `k` lines: triples `h w r`  \n
  `1 ≤ h, w ≤ n`, `1 ≤ r ≤ 1000`

Guarantees:

- Each pair `(h, w)` appears in **at most one** triple.
- `t` does not exceed the total number of acceptable sets.

**Subtasks:**

- For 30 points: `1 ≤ n ≤ 5`
- For 100 points: `1 ≤ n ≤ 20`

### Output

Print a single integer — the value of the `t`‑th acceptable set (the sum of `r` for that matching).

### Examples

**Input**
```text
2 4 3
1 1 1
1 2 2
2 1 3
2 2 7
```

**Output**
```text
2
```

**Input**
```text
2 4 7
1 1 1
1 2 2
2 1 3
2 2 7
```

**Output**
```text
8
```

### ideas
1. 二部图，有k条边（连接左右）
2. 从k条边中选择部分，组成匹配，得到第t小的值
3. 按照r升序排列进行处理，假设目前是第w个，现在要找第w+1个
4. dp[x] 表示能够达到x分数的个数（x最多是20000）
5.  dp[x] ? 怎么算？

### key insights

**Min-heap enumeration (priority queue approach)**

- Sort edges by `r` (optional but keeps low-value edges at low indices).
- Represent each matching as a state `(value, lastI, menMask, womenMask)`:
  - `lastI`: index of the last edge included (edges sorted by index).
  - `menMask` / `womenMask`: bitmasks of used men/women (fits in int for n ≤ 20).
- Seed the heap with the empty matching `(0, -1, 0, 0)`.
- Each pop yields the next matching in non-decreasing order; extend it by pushing all compatible edges with index `> lastI`.

**Uniqueness**: Every matching `{e_{i1} < e_{i2} < ... < e_{im}}` is generated exactly once — it can only be produced by extending `{e_{i1}, ..., e_{i(m-1)}}` with `e_{im}`. The "only add edges with higher index" rule enforces this canonical form.

**Complexity**: O(t × k × log(t×k)).  With t ≤ 2×10⁵ and k ≤ 100, total heap pushes ≤ 2×10⁷.
