### Problem

Petya loves lucky numbers. A **lucky number** is a positive integer whose decimal representation consists **only** of digits `4` and `7`.  
For example, `4`, `47`, `744` are lucky, while `5`, `17`, `467` are not.

Petya dreams of being president of an island country. The country:

- Has `n` islands (nodes) connected by `m` **undirected** roads (edges).
- Some islands cannot reach others, even via intermediate islands, so the country is divided into **connected components**, which we will call **regions**.

Formally:

- Each island belongs to exactly one region.
- In a region, any two islands are connected by some path.
- There is no path between islands from different regions.

A region is **lucky** if the **number of islands** (its size) is a lucky number.

Petya wants to build his presidential palace in a lucky region. If initially there is no lucky region, he may **add roads between islands in different regions**, thereby merging regions. Your task is to find the **minimum number of new roads** Petya needs to add to create **at least one lucky region**.

### Input

- First line: two integers `n` and `m` (`1 ≤ n, m ≤ 10^5`) — the number of islands and existing roads.
- Next `m` lines: each contains two integers `u v` (`1 ≤ u, v ≤ n`) describing an undirected road between islands `u` and `v`.
  - Self-loops (`u = v`) are possible.
  - Multiple edges between the same pair of islands are allowed.

### Output

- If it is impossible to obtain a lucky region by adding roads, print `-1`.
- Otherwise, print a single integer `r` — the **minimum number of roads** that need to be added to create at least one lucky region.

### Examples

**Input**
```text
4 3
1 2
2 3
1 3
```

**Output**
```text
1
```

**Input**
```text
5 4
1 2
3 4
4 5
3 5
```

**Output**
```text
-1
```


### ideas
1. dp[x] = 到达x-size的region所需要的最少road数
2. dp[x] = dp[x - w] + 1 如果当前region的sz = w. 但是这样的问题是，time complexity = O(n * n)
3. 在 <= 1e5的范围内，lucky数字，似乎是比较少的
4. 4, 7, 44, 47, 74, 77, 444, 2 + 4 + 8 + 16 = 30个
5. 也就是说，可以按照这个列表去检查，是否能够到达所需的最优解？
6. 特别是 777的情况下，dp[x] 应该还是很快的
7. 但是假设一直都没有达到结果？还是会退化到n.
8. 不知道～
9. 