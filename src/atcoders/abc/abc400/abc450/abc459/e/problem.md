# E - Select from Subtrees (ABC459)

**Contest:** [ABC459](https://atcoder.jp/contests/abc459) — AtCoder Beginner Contest 459  
**Task:** [https://atcoder.jp/contests/abc459/tasks/abc459_e](https://atcoder.jp/contests/abc459/tasks/abc459_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

There is a rooted tree `T` with `N` vertices numbered `1, 2, ..., N`. Vertex `1` is
the root. For each `2 <= i <= N`, the parent of vertex `i` is `P_i`.

Vertex `i` (`1 <= i <= N`) has `C_i` candies. All
`C_1 + C_2 + ... + C_N` candies are distinguishable.

Takahashi gives instructions to `N` squirrels. Squirrel `i` (`1 <= i <= N`) must
choose and collect exactly `D_i` candies from the subtree rooted at vertex `i`.
Different squirrels cannot take the same candy.

Find the number of valid assignments, modulo `998244353`.

Even if the final set of chosen candies is the same, count assignments separately
when different squirrels choose them. If no valid assignment exists, output `0`.

## Constraints

- `2 <= N <= 2 * 10^5`
- `1 <= P_i <= N`
- `1 <= C_i <= 10^9`
- `1 <= D_i`
- `D_1 + D_2 + ... + D_N <= 10^6`
- All input values are integers.
- `T` is a rooted tree with vertex `1` as the root.

## Input

The input is given from standard input in the following format:

```text
N
P_2 P_3 ... P_N
C_1 C_2 ... C_N
D_1 D_2 ... D_N
```

## Output

Print the number of valid ways to choose candies, modulo `998244353`.

## Sample Input 1

```text
5
1 1 3 3
1 2 1 2 3
1 1 3 1 1
```

## Sample Output 1

```text
144
```

The tree is:

```text
1 (1, 1)
├── 2 (2, 1)
└── 3 (1, 3)
    ├── 4 (2, 1)
    └── 5 (3, 1)
```

Label the candies as follows: candy `1` at vertex `1`; candies `2, 3` at vertex `2`;
candy `4` at vertex `3`; candies `5, 6` at vertex `4`; candies `7, 8, 9` at vertex
`5`.

One valid assignment:

- Squirrel `1` takes candy `3` from the subtree of vertex `1`.
- Squirrel `2` takes candy `2` from the subtree of vertex `2`.
- Squirrel `3` takes candies `4, 7, 9` from the subtree of vertex `3`.
- Squirrel `4` takes candy `6` from the subtree of vertex `4`.
- Squirrel `5` takes candy `8` from the subtree of vertex `5`.

Swapping which squirrel takes which candy counts as a different assignment even if
the chosen candy set is unchanged.

## Sample Input 2

```text
2
1
1 1
2 1
```

## Sample Output 2

```text
0
```

The tree has only two candies in total, so both squirrels cannot satisfy their
requirements.

## Sample Input 3

```text
3
3 1
1000000000 1 1
1 1 1
```

## Sample Output 3

```text
1755647
```

The tree is:

```text
1 (root)
└── 3
    └── 2
```

Label the candies as follows:

- vertex `1`: candies `1, 2, ..., 1,000,000,000`
- vertex `2`: candy `1,000,000,001`
- vertex `3`: candy `1,000,000,002`

One valid assignment:

- Squirrel `1` collects candy `1` from the subtree of vertex `1` (vertices `1, 2, 3`).
- Squirrel `2` collects candy `1,000,000,001` from the subtree of vertex `2` (vertex `2` only).
- Squirrel `3` collects candy `1,000,000,002` from the subtree of vertex `3` (vertices `3, 2`).

Squirrels `2` and `3` are forced: squirrel `2` must take the only candy in its
subtree, and squirrel `3` must then take the only remaining candy in its subtree.
Squirrel `1` may choose any of the `10^9` candies at vertex `1`, so there are
`10^9` valid assignments. Output `10^9 mod 998244353 = 1755647`.


### ideas
1. let s[i] = sum of C[j] where j is in sub-tree of i
2. if d[i] > s[i] => 0
3. 假设在v上可以获取到的candy的数量为w, 
4. C(w, d[v]) ？ 但是w非常的大呐 
5. 对于example1, dp[5] = 3, dp[4] = 2, dp[3] = C(4, 3) * dp[4] * dp[5] = 4 * 3 * 2 = 24
6. dp[2] = 2, dp[1] = dp[2] * dp[3] * C(3, 1) = 2 * 24 * 6 = 144?
7. 似乎是对的。但是w还是太大。
8. w = sum(C[?]) - Sum(D[?]) + c[v] 
9. 好吧～