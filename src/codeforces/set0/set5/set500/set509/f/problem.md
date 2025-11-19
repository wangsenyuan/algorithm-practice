# Problem F

Programming teacher Dmitry Olegovich is going to propose the following task for one of his tests for students:

You are given a tree T with n vertices, specified by its adjacency matrix a[1... n, 1... n]. What is the output of the following pseudocode?

```pseudocode
used[1 ... n] = {0, ..., 0};

procedure dfs(v):
    print v;
    used[v] = 1;
    for i = 1, 2, ..., n:
        if (a[v][i] == 1 and used[i] == 0):
            dfs(i);

dfs(1);
```

In order to simplify the test results checking procedure, Dmitry Olegovich decided to create a tree T such that the result is his favorite sequence b. On the other hand, Dmitry Olegovich doesn't want to provide students with same trees as input, otherwise they might cheat. That's why Dmitry Olegovich is trying to find out the number of different trees T such that the result of running the above pseudocode with T as input is exactly the sequence b. Can you help him?

Two trees with n vertices are called different if their adjacency matrices a1 and a2 are different, i. e. there exists a pair (i, j), such that 1 ≤ i, j ≤ n and a1[i][j] ≠ a2[i][j].

## Input

The first line contains the positive integer n (1 ≤ n ≤ 500) — the length of sequence b.

The second line contains n positive integers b1, b2, ..., bn (1 ≤ bi ≤ n). It is guaranteed that b is a permutation, or in other words, each of the numbers 1, 2, ..., n appears exactly once in the sequence b. Also it is guaranteed that b1 = 1.

## Output

Output the number of trees satisfying the conditions above modulo 10^9 + 7.

## Examples

### Example 1

**Input:**

```text
3
1 2 3
```

**Output:**

```text
2
```

### Example 2

**Input:**

```text
3
1 3 2
```

**Output:**

```text
1
```

### ideas
1. dfs是 in-order travese
2. 当b[1]是root的时候，假设它有k个children， 必须满足 c[1] < c[2] < ... c[k]
3. 且c[1] = b[2]
4. dp[l...r] 表示由l...r组成一个由b[l]为root的子树的计数
5. 假设目前需要考虑l...r的情况
6. 那么root就是b[l], 那么root的子节点，应该可以组成一个递增序列
7. 但是这个选择太多了～
8. 貌似就是 n * n * n 的复杂性呐