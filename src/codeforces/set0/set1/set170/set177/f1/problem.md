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
4. 这时候有几种情况，一种是用一条更大的边，替换一条更小的边？
5. 倒着替换，知道替换完为止
6. 下一步只能加入一条新的边
7. 替换的时候，如果这条边关联的man（或者woman）已经处在一个关系中？
8. 那么就必须match一次，找到这条路径上，最大的边
9. dp[x] 表示能够达到x分数的个数（x最多是20000）
10. dp[x] ? 怎么算？
11. 从h开始，找到目前匹配的路径，然后用这个路径去交换r