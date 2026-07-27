# B. Maxim and Restaurant

[Problem link](https://codeforces.com/problemset/problem/261/B)

**Contest:** [Codeforces Round #160 (Div. 1)](https://codeforces.com/contest/261)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Maxim has opened his own restaurant! The restaurant has got a huge table, the
table's length is `p` meters.

Maxim has got a dinner party tonight, `n` guests will come to him. Let's index
the guests of Maxim's restaurant from `1` to `n`. Maxim knows the sizes of all
guests that are going to come to him. The `i`-th guest's size (`a_i`) represents
the number of meters the guest is going to take up if he sits at the restaurant
table.

Long before the dinner, the guests line up in a queue in front of the restaurant
in some order. Then Maxim lets the guests in, one by one. Maxim stops letting
the guests in when there is no place at the restaurant table for another guest
in the queue. There is no place at the restaurant table for another guest in the
queue, if the sum of sizes of all guests in the restaurant plus the size of this
guest from the queue is larger than `p`. In this case, not to offend the guest
who has no place at the table, Maxim doesn't let any other guest in the
restaurant, even if one of the following guests in the queue would have fit in
at the table.

Maxim is now wondering, what is the average number of visitors who have come to
the restaurant for all possible `n!` orders of guests in the queue. Help Maxim,
calculate this number.

## Constraints

- `1 <= n <= 50`
- `1 <= a_i <= 50`
- `1 <= p <= 50`

## Input

The first line contains integer `n` — the number of guests in the restaurant.

The next line contains integers `a_1, a_2, ..., a_n` — the guests' sizes in
meters.

The third line contains integer `p` — the table's length in meters.

The numbers in the lines are separated by single spaces.

## Output

In a single line print a real number — the answer to the problem. The answer
will be considered correct, if the absolute or relative error doesn't exceed
`10^{-4}`.

## Sample 1

```text
Input
3
1 2 3
3

Output
1.3333333333
```

### Note

In the first sample the people will come in the following orders:

- `(1, 2, 3)` — there will be two people in the restaurant;
- `(1, 3, 2)` — there will be one person in the restaurant;
- `(2, 1, 3)` — there will be two people in the restaurant;
- `(2, 3, 1)` — there will be one person in the restaurant;
- `(3, 1, 2)` — there will be one person in the restaurant;
- `(3, 2, 1)` — there will be one person in the restaurant.

In total we get `(2 + 1 + 2 + 1 + 1 + 1) / 6 = 8 / 6 = 1.(3)`.

## ideas

1. `dp[s][i]` 表示使用 `i` 个人, 组成恰好 `sum = s` 的方案数
2. `dp[s][i] = dp[s-v][i-1] + dp[s][i]` 如果 `a[i] = v`
3. 但是这里的问题是, `s` 必须正好是不超过 `p` 的数, 也就是说剩余的部分中,
   不能出现 `s + w <= p`
4. `dp[s][x][i]` 表示有 `i` 个人组成的, `sum = s`, 且没有被选中的其中最小值是 `x`
   时的方案数
5. `dp[s][x][i] += dp[s-v][i-1][min(x, v)]`

### Pitfalls of the min-leftover DP

Sample expects `1.333…`; a buggy implementation that uses the ideas above can get
`1.5`. Typical mistakes:

1. **Wrong quantity in the numerator.** Accumulating `s * dp[s][i][x]` uses the
   **size sum** `s`, but the expectation is over **how many people** sit, so it
   must be `i`, not `s`.

2. **`min` leftover is too strong.** `x = min(not taken)` and `s + x > p` only
   keeps sets where **every** leftover guest overflows. You miss cases where
   some leftover guest still fits, but the **next** person in the queue does
   not. Sample: `(1, 3, 2)` seats `{1}` (sum `1`), next is `3`. Leftover
   `{2,3}`, `min=2`, and `1+2 ≤ 3`, so the min-condition drops this permutation
   — but it is valid and contributes `1`.

3. **Wrong combinatorial weight.** For a seated set `S` (`|S|=i`) and a valid
   next guest `j` (overflow), the count of permutations is
   `i! · (n - i - 1)!` (not just `i!`). Any of the blocking guests can be next;
   the rest of the suffix is free.

4. **Early return when everyone fits.** If `sum(a) ≤ p`, every permutation seats
   all guests — return `n` (guest count), not the size sum.

### What to count instead

Subsets `S` with `sum(S) ≤ p`, and guests `j ∉ S` with `sum(S) + a_j > p`. Each
such `(S, j)` contributes

```text
|S| · |S|! · (n - |S| - 1)!
```

to the total, then divide by `n!`.

## summary

Notation: `Pr(...)` means **probability** (fraction of the `n!` permutations).

### Process

Guests arrive in a random order `π`. Walk from the front: keep seating while the
sum stays `≤ p`; stop at the first overflow. Let `X(π)` be how many sat.
Answer = `E[X]` (average over all permutations).

### Tail formula

For `X ∈ {0,1,…,n}`:

```text
E[X] = Σ_{k=1..n} Pr(X ≥ k)
```

Because `X = I(X≥1) + I(X≥2) + … + I(X≥X)`.

### What `X ≥ k` means

`X ≥ k` iff the **first `k` guests** all sat, i.e.

```text
a[π_1] + … + a[π_k] ≤ p
```

This only depends on the set `{π_1,…,π_k}` occupying the first `k` positions.
It does **not** mention `π_{k+1}`. The blocker check `sum + a[next] > p`
belongs to the event `X = k`, not to `X ≥ k`.

### Counting

Let `ways(k)` = number of `k`-subsets with sum `≤ p`. Then

```text
#{π : X ≥ k} = ways(k) · k! · (n-k)!
Pr(X ≥ k)    = ways(k) / C(n,k)
```

so

```text
E[X] = Σ_k ways(k) / C(n,k)
```

No explicit `s + a[?] > p` is needed: overflow of the next guest only decides
whether `X` stops at `k` or continues, which is captured by later terms
`Pr(X ≥ k+1)`, …

Equivalent views:

```text
E[X] = Σ_k k · Pr(X = k)     ← needs a blocker for each exact k
     = Σ_k Pr(X ≥ k)         ← no blocker
```

### Sample `n=3`, `a=[1,2,3]`, `p=3`

Permutations give `X ∈ {2,1,2,1,1,1}`, average `8/6 = 4/3`.

| k | ways(k) | C(3,k) | Pr(X≥k) |
|---|---------|--------|---------|
| 1 | 3       | 3      | 1       |
| 2 | 1 (`{1,2}` only) | 3 | 1/3 |
| 3 | 0       | 1      | 0       |

`E = 1 + 1/3 + 0 = 4/3`.

### Code

1. 0-1 knapsack: `dp[j][s]` = # of `j`-subsets with sum `s`.
2. `ways(k) = Σ_{s=0..p} dp[k][s]`.
3. Add `ways(k) / C(n,k)` for each `k` (plain `float64`; no `big.Int`).

One-line: average number seated = sum over lengths `k` of (fraction of
`k`-sets that fit as a prefix).
