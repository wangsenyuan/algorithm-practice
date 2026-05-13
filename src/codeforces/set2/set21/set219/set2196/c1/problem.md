# C1. Interactive Graph (Simple Version) (Codeforces 2196C1)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2196/C1](https://codeforces.com/problemset/problem/2196/C1)

---

**This is the simple version.** You may ask **at most** `32 · (n + m)` questions, and `n ≤ 15`. You can hack only if you solved **all** versions of this problem.

**This is an interactive problem.**

The jury has a **directed acyclic graph** (no self-loops, no multiple edges) with `n` vertices and `m` edges.

Your task is to determine **which edges** are present. You may ask questions of the form: what does the **`k`-th path** look like in the **lexicographically** sorted list of **all paths** in the graph.

A **path** is a sequence of vertices `u_1, u_2, …, u_l` such that for every `i < l` there is a directed edge `(u_i, u_{i+1})` in the graph.

You must succeed using **at most** `32 · (n + m)` questions.

### Lexicographic order (footnote)

A sequence `a` is lexicographically smaller than a sequence `b` if and only if either:

- `a` is a **strict prefix** of `b`, or  
- at the first index where they differ, `a` has a **smaller** vertex number than `b`.

## Input

The first line contains an integer `t` (`1 ≤ t ≤ 10`) — the number of test cases.

Each test case is one line with an integer `n` (`1 ≤ n ≤ 15`) — the number of vertices.

The jury guarantees the hidden graph is a DAG with **no multiple edges**.

**Note:** `m` is **unknown** to you until you deduce it.

## Interaction

For each test case, after reading `n`:

- You may ask up to `32 · (n + m)` questions.
- To ask a question, print `? k` (`1 ≤ k ≤ 2^30`). Then read an integer `q` — the length of the `k`-th path (number of vertices). If `q = 0`, that path does not exist; otherwise read `q` integers — the vertices on that path.
- Under the constraints, the total number of distinct paths does not exceed `2^30`.
- To answer, print `! m`, then print `m` lines `u v` meaning a directed edge from `u` to `v`. Edges may be in **any order**. Printing the answer does **not** count as a query.

After **each** query (and when required by your runtime), print a newline and **flush** output (e.g. `fflush(stdout)` / `cout.flush()` / `sys.stdout.flush()`).

If you read `-1`, your program must **exit immediately** (invalid query or protocol error).

### Hacks

First line: `t` (`1 ≤ t ≤ 10`).

Each test case: first line `n m` (`1 ≤ n ≤ 15`, `0 ≤ m ≤ n(n−1)/2`). Next `m` lines: `v u` for an edge `v → u`. Graph must be a DAG with no multiple edges.

## Example

The statement’s PDF/HTML uses two columns: **Input** = everything the interactor writes to your stdin (in chronological order for the reference solution); **Output** = everything your program prints.

### Input (interactor)

```text
3
5

1 1

2 1 2

3 1 2 4

3 1 2 5

2 1 3

3 1 3 4

3 1 3 5

1 2

1 3

1 4

1 5

1

0

2

1 1

1 2

2 2 1
```

### Output (participant)

```text
? 1

? 2

? 3

? 4

? 5

? 6

? 7

? 8

? 11

? 14

? 15

! 6
1 3
1 2
2 4
3 4
2 5
3 5

? 2

! 0

? 1

? 2

? 3

! 1
2 1
```

## Note

The first test case’s hidden graph is illustrated in the original statement (figure: [espresso image](https://espresso.codeforces.com/bf518061806cb126a06a412783ef4cafb1d04ca4.png)). There are **15** paths, in lexicographic order:

1. `1`  
2. `1 → 2`  
3. `1 → 2 → 4`  
4. `1 → 2 → 5`  
5. `1 → 3`  
6. `1 → 3 → 4`  
7. `1 → 3 → 5`  
8. `2`  
9. `2 → 4`  
10. `2 → 5`  
11. `3`  
12. `3 → 4`  
13. `3 → 5`  
14. `4`  
15. `5`


### ideas
1. 这个好难呐。完全没想法
2. 如果path(i) 从1开始，那么后面的都一直会从1开始，知道path(j) 不从1开始，那么后面就不会从j开始
3. 可以这样询问吗？
4. 那么每个节点贡献30次，似乎是可以的
5. 好像这样查询一次后，就知道了dp[v] = 有多少条路径从v开始
6. 但是怎么查询i的后续节点呢？
7. 然后查询, path(s[0])， 它的后继，肯定是第一个最小的v，这样子就可以顺着path查询下去
