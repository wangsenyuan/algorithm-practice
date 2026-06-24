# F - Total Product is N (ABC461)

**Contest:** [ABC461](https://atcoder.jp/contests/abc461) — AtCoder Beginner Contest 461  
**Task:** [https://atcoder.jp/contests/abc461/tasks/abc461_f](https://atcoder.jp/contests/abc461/tasks/abc461_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 500 points

## Problem Statement

You are given a positive integer `N`. A non-empty sequence of positive integers `A`
is called a **good sequence** if:

- all elements of `A` are distinct
- the product of all elements of `A` equals `N`

The **score** of a sequence is the sum of its elements. Find the sum, modulo
`998244353`, of the scores of all good sequences.

## Constraints

- `1 <= N <= 10^10`
- the input value is an integer

## Input

```text
N
```

## Output

Print the answer modulo `998244353`.

## Sample Input 1

```text
8
```

## Sample Output 1

```text
80
```

There are 11 good sequences:

`(1,2,4)`, `(1,4,2)`, `(1,8)`, `(2,1,4)`, `(2,4)`, `(2,4,1)`, `(4,1,2)`, `(4,2)`,
`(4,2,1)`, `(8)`, `(8,1)`.

Their scores are `7, 7, 9, 7, 6, 7, 7, 6, 7, 8, 9`, and the total is `80`.

## Sample Input 2

```text
461
```

## Sample Output 2

```text
1385
```

## Sample Input 3

```text
100
```

## Sample Output 3

```text
1702
```

## Editorial

**Official editorial:** [F - Total Product is N](https://atcoder.jp/contests/abc461/editorial/21390) (ABC461, original proposer: vwxyz)

### DP over divisors

Let `d_1, d_2, ..., d_K` be all divisors of `N` (in any fixed order). When choosing `b` of these divisors whose product is `c`, define:

- `dp0[a, b, c]` — number of such choices using only divisors among `d_1, ..., d_a`
- `dp1[a, b, c]` — sum of their scores (sum of chosen divisors) over those choices

Base cases: `dp0[0, 0, 1] = 1`, `dp1[0, 0, 1] = 0`, and all other states are `0`.

When processing divisor `d_a`:

- If `c` is **not** divisible by `d_a`, we cannot take `d_a`:
  - `dp0[a, b, c] = dp0[a - 1, b, c]`
  - `dp1[a, b, c] = dp1[a - 1, b, c]`
- If `c` **is** divisible by `d_a`, we may skip or take `d_a`:
  - `dp0[a, b, c] = dp0[a - 1, b - 1, c / d_a] + dp0[a - 1, b, c]`
  - `dp1[a, b, c] = dp1[a - 1, b - 1, c / d_a] + dp0[a - 1, b - 1, c / d_a] * d_a + dp1[a - 1, b, c]`

The second term in `dp1` adds `d_a` to every partial choice counted by `dp0[a - 1, b - 1, c / d_a]`.

### Answer

The DP counts **unordered** multisets of divisors with product `N`. Every good sequence is a permutation of such a multiset, so multiply by `b!`:

```text
answer = sum over b >= 1 of dp1[K, b, N] * b!   (mod 998244353)
```

### Complexity

- For `1 <= N <= 10^10`, `N` has at most **2304** divisors, so index `a` and product `c` only need divisor values.
- `14! > 10^10`, so no good sequence has more than **14** elements; it suffices to use `b <= 14`.
- Time: **O(14 * 2304^2)**, fast enough for the limits.

## Solution (Detailed)

Sum, modulo `998244353`, the **scores** (element sums) of every **good sequence** — a
non-empty ordered list of **distinct** positive integers whose product is `N`.

### Short recap

- **Good sequence:** non-empty, all elements distinct, product exactly `N`.
- **Score:** sum of elements in the sequence.
- **Answer:** sum of scores over **all** good sequences (order matters), mod
  `998244353`.
- **Example (`N = 8`):** 11 good sequences (see sample); their scores sum to `80`.

### What we are counting (ordered vs unordered)

The problem counts **ordered** sequences. `(2, 4)` and `(4, 2)` are different good
sequences and both contribute score `6`.

A useful split:

1. Choose an **unordered set** of distinct divisors whose product is `N`.
2. **Permute** those divisors in all `b!` orders (`b` = set size).

The DP below handles step 1; the final answer multiplies by `b!` to account for
step 2.

### Key observations

1. **Every element must divide `N`.** If some element `x` had a prime factor not
   dividing `N`, the full product would also contain that prime. So each term in a
   good sequence is a **divisor of `N`**. We never need numbers outside the divisor
   list.

2. **Elements are distinct → choose each divisor at most once.** The DP is
   “subset DP over divisors”: for each divisor `d_a`, we either take it once or
   skip it. This automatically excludes invalid choices like `(2, 2, 2)` for
   `N = 8`.

3. **State uses only divisor products.** The running product `c` in the DP is always
   a divisor of `N` (start from `1`, multiply by divisors). So `c` can be indexed
   by position in the sorted divisor list — at most **2304** values.

4. **Length bound `b ≤ 13` (loop `b ≤ 14` is safe).** The smallest possible product
   of `b` distinct positive integers is `1 · 2 · … · b = b!` (take
   `{1, 2, …, b}`). We need `b! ≤ N ≤ 10^10`:
   - `13! = 6 227 020 800 ≤ 10^10` — still possible.
   - `14! = 87 178 291 200 > 10^10` — impossible.
   So no good sequence has more than **13** elements. Running the DP for
   `b = 0 … 14` is harmless (extra layers stay zero).

5. **Special case `N = 1`.** The only good sequence is `(1)` with score `1`. Handle
   this before the general DP.

### Divisor count: why at most 2304

For `N ≤ 10^10`, the maximum number of divisors is **2304** (attained e.g. by
highly composite numbers near the limit). Any implementation only needs arrays of
size `K = τ(N) ≤ 2304` for the product dimension. Combined with `b ≤ 14`, total
state count is on the order of `14 × 2304² ≈ 7.4 × 10⁷`, which is fine in 2
seconds.

### DP definitions

Let `d_1 < d_2 < … < d_K` be **all divisors of `N`** in increasing order (`K ≤ 2304`).

When building a subset using only divisors among `d_1, …, d_a`, with exactly `b`
chosen divisors and product `c`:

| State | Meaning |
| --- | --- |
| `dp0[a, b, c]` | Number of ways to choose **exactly `b`** divisors from `{d_1,…,d_a}` whose product is **`c`**. |
| `dp1[a, b, c]` | Sum of **scores** (sum of chosen divisors) over those same choices. |

**Base case** (no divisors processed, empty subset):

```text
dp0[0, 0, 1] = 1    // one way: choose nothing, product 1
dp1[0, 0, 1] = 0    // empty subset has score 0
all other dp0[0, *, *] = dp1[0, *, *] = 0
```

**Interpretation:** `dp0` counts subsets; `dp1` is the score-sum over those
subsets. They share the same `(a, b, c)` indexing.

### Transitions (processing divisor `d_a`)

**Case A — `c` is not divisible by `d_a`.** We cannot include `d_a` without
breaking the target product `c`:

```text
dp0[a, b, c] = dp0[a - 1, b, c]
dp1[a, b, c] = dp1[a - 1, b, c]
```

**Case B — `c` is divisible by `d_a`.** Skip or take `d_a`:

```text
dp0[a, b, c] = dp0[a - 1, b - 1, c / d_a] + dp0[a - 1, b, c]
               \_________________________/   \_______________/
                     take d_a                    skip d_a
```

For `dp1`, three contributions:

```text
dp1[a, b, c] = dp1[a - 1, b - 1, c / d_a]     // take d_a: prior score sums
             + dp0[a - 1, b - 1, c / d_a] * d_a  // take d_a: add d_a to each subset
             + dp1[a - 1, b, c]                 // skip d_a
```

**Why the middle term `dp0[…] * d_a`?** Every subset counted by
`dp0[a - 1, b - 1, c / d_a]` has some score `S`. Appending `d_a` gives score
`S + d_a`. Summing over all such subsets:

```text
Σ (S + d_a) = Σ S + (number of subsets) · d_a = dp1[a-1,b-1,c/d_a] + dp0[a-1,b-1,c/d_a] · d_a
```

The first term is already in `dp1[a - 1, b - 1, c / d_a]`; the second is the
extra `d_a` per subset.

**Small example of one transition (`N = 8`, `d_a = 2`, state `(a,b,c) = (2,1,2)`):**

Divisors so far: `{1, 2}`. Want `b = 1`, product `c = 2`.

- Take `2`: from `dp0[1, 0, 1] = 1` (empty → product 1), get `{2}`, score `2`.
- Skip: `dp0[1, 1, 2] = 0` (cannot make product 2 using only `{1}` with one element).

So `dp0[2, 1, 2] = 1`, `dp1[2, 1, 2] = 2`.

### Final answer (unordered → ordered)

The DP at `(K, b, N)` aggregates **unordered** subsets of size `b` with product
`N`. Each such subset yields **`b!` ordered** good sequences with the same
multiset of elements.

```text
answer = Σ_{b = 1}^{K}  dp1[K, b, N] · b!   (mod 998244353)
```

Do **not** multiply `dp0` by `b!` — we need the **sum of scores**, not the count
of sequences. Every permutation of a fixed subset has the same score, so
multiplying the total score for that subset by `b!` is correct.

### Algorithm

**Step 0 — edge case.** If `N == 1`, return `1`.

**Step 1 — enumerate divisors.** Collect all `d | N`, sort ascending →
`d[0], …, d[K-1]`. Build a map `id[c]` = index of divisor `c` (only needed for
`c` that appear as products).

**Step 2 — allocate DP.** Two tables `dp0[b][idx]`, `dp1[b][idx]` for
`b = 0…14` and `idx` over divisor indices. Initialize `dp0[0][id[1]] = 1`,
`dp1[0][id[1]] = 0`.

**Step 3 — loop over divisors.** For each `a = 0 … K-1` with value `d = d[a]`:

1. Copy current layer to `next0`, `next1` (skip branch).
2. For each `b = 1…14`, each divisor index `j` with value `c = d[j]`:
   - If `c % d != 0`, continue.
   - Let `j0 = id[c / d]`.
   - **Take:** `next0[b][j] += dp0[b-1][j0]`, `next1[b][j] += dp1[b-1][j0] + dp0[b-1][j0] * d`.
3. Replace `dp0, dp1` with `next0, next1`.

**Step 4 — combine with factorials.** Precompute `fact[b] = b! mod MOD`. Return

```text
Σ_{b=1}^{14} dp1[b][id[N]] · fact[b]  mod 998244353
```

### Walkthrough: `N = 8`

Divisors: `d = [1, 2, 4, 8]`, `K = 4`.

**Valid unordered subsets** (product 8) and their scores:

| Subset | Score | Size `b` | Ordered sequences | Score × `b!` |
| --- | --- | --- | --- | --- |
| `{8}` | 8 | 1 | `(8)` | 8 × 1 = 8 |
| `{2, 4}` | 6 | 2 | `(2,4)`, `(4,2)` | 6 × 2 = 12 |
| `{1, 8}` | 9 | 2 | `(1,8)`, `(8,1)` | 9 × 2 = 18 |
| `{1, 2, 4}` | 7 | 3 | 6 permutations | 7 × 6 = 42 |

**Total:** `8 + 12 + 18 + 42 = 80` — matches sample output.

After the full DP, `dp1[K, b, N]` for `b = 1, 2, 3` should be `8, 15, 7`
respectively (`15 = 6 + 9` from the two 2-element subsets). Then
`8·1! + 15·2! + 7·3! = 8 + 30 + 42 = 80`.

**Why `(2,2,2)` never appears:** divisor `2` is processed once; “take” uses
`b - 1` and product `c / 2` — you cannot take `2` twice from a set of distinct
divisors.

### Why it works

- **Exhaustive subset enumeration:** Processing divisors in fixed order with
  take/skip is the standard 0/1 knapsack / subset-DP pattern. Distinctness is
  enforced because each `d_a` is considered at most once.
- **Product correctness:** We only extend states where `c` is divisible by `d_a`,
  and we divide `c` when taking `d_a`, so the invariant “current product equals
  the product of chosen divisors” holds.
- **Score aggregation:** The `dp1` recurrence is linear: adding `d_a` to every
  partial score is equivalent to adding `dp0 · d_a` to the sum.
- **Ordering:** Good sequences are permutations of exactly these subsets (each
  divisor used at most once). All permutations share one score, hence multiply by
  `b!`.

### Complexity

| | Bound |
| --- | --- |
| `K = τ(N)` | ≤ 2304 |
| `b` | ≤ 14 (only 1…13 matter) |
| **Time** | `O(14 · K²)` — for each of `K` divisors, update `O(14 · K)` product states |
| **Space** | `O(14 · K)` with two rolling layers per `dp0` / `dp1` |

Well within limits for `N ≤ 10^10`.

### Implementation note (`solution.go`)

The reference implementation in this directory is **not yet** the divisor-DP above;
tests currently fail (`go test` expects `80, 1385, 1702`). A correct solution
follows the algorithm in **Step 0–4**: enumerate divisors, run the subset DP, then
multiply `dp1[K, b, N]` by `b!` and sum over `b`.

## ideas
1. 由n的因数构成的一个序列, f[1], f[2], .. f[i]
2. f[1] * f[2] * f[i] = n
3. 对n进行质因分解, n = p[1] ** c[1] * p[2] ** c[2] .. p[i] ** c[i]
4. 假设这个序列中w个数, 那么c[1]个p[1]就需要分配到w个数里面C(w + c[1], c[1] - 1) ?
5. p[1] ** x[1] * p[2] ** x[2]... p[i]**x[i] + p[1] ** y[1] * p[2] ** y[2], ..
6. 还是得整体考虑,  n[1]是n的第一个因数(不是质因数), 要计算它出现了多少次
7. 那么就是c[1] - x[1], c[2] - x[2] ... c[i] - x[i] 这些数分配出去
8. 但是(必须保证n1不能再出现)
9. 不考虑n1重复出现, C(sum, c[1] - x1) * C{sum - .. c[2] - x[2]} ...
10. 不对, 除了n1重复, 其他数也有可能重复
11. 比如对于8, 如何排除掉 [2, 2, 2]?
12. 假设n = a * b * c * d, 且互不相同