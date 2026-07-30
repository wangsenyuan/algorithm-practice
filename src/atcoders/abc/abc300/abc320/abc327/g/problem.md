# G - Many Good Tuple Problems

[Problem link](https://atcoder.jp/contests/abc327/tasks/abc327_g)

**Contest:** [AtCoder Beginner Contest 327](https://atcoder.jp/contests/abc327)

time limit: 2 sec

memory limit: 1024 MiB

score: 650 points

## Problem Statement

> The definition of a good pair of sequences in this problem is the same as in
> Problem D.

A pair of sequences of length `M` consisting of positive integers at most `N`,
`(S, T) = ((S_1, ..., S_M), (T_1, ..., T_M))`, is said to **be a good pair of
sequences** when there exists a sequence `X = (X_1, ..., X_N)` of length `N`
consisting of `0` and `1` such that:

- `X_{S_i} != X_{T_i}` for each `i = 1, 2, ..., M`.

Among the `N^{2M}` possible pairs of sequences `(A, B)` of length `M`
consisting of positive integers at most `N`, find the number of good pairs
modulo `998244353`.

## Constraints

- `1 <= N <= 30`
- `1 <= M <= 10^9`
- `N` and `M` are integers.

## Input

```text
N M
```

## Output

Print the number, modulo `998244353`, of good pairs of sequences of length `M`
consisting of positive integers at most `N`.

## Sample Input 1

```text
3 2
```

## Sample Output 1

```text
36
```

For example, `A = (1, 2)`, `B = (2, 3)` is a good pair (`X = (0, 1, 0)` works).
There are 36 good pairs in total.

## Sample Input 2

```text
3 3
```

## Sample Output 2

```text
168
```

## Sample Input 3

```text
12 34
```

## Sample Output 3

```text
539029838
```

## Sample Input 4

```text
20 231104
```

## Sample Output 4

```text
966200489
```


## ideas
1. N <= 30, n个节点, 要能分成2部图. 
2. w个节点为红色, n - w 个为奇数, 
3. tot = w * (n - w) >= m, nCr(tot, m)?
4. 不对, 应该可以出现重复的pair
5. 