# F - 1122 Subsequence 2 (ABC433)

**Contest:** [ABC433](https://atcoder.jp/contests/abc433) — AtCoder Beginner Contest 433  
**Task:** [https://atcoder.jp/contests/abc433/tasks/abc433_f](https://atcoder.jp/contests/abc433/tasks/abc433_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 500 points

## Problem Statement

You are given a string `S` consisting of digits.

A string `T` is called a **1122-string** if it satisfies all of the following
(the same definition as in Problem C):

- `T` is a non-empty string consisting of digits
- `|T|` is even
- The first `|T|/2` characters of `T` are all the same digit
- The last `|T|/2` characters of `T` are all the same digit
- Adding `1` to the first digit of `T` gives the last digit of `T`

For example, `1122`, `01`, and `444555` are 1122-strings, but `1222` and `90`
are not.

Find the number of **subsequences** of `S` (not necessarily contiguous) that are
1122-strings, modulo `998244353`.

Count subsequences by the positions chosen: two subsequences are different if they
use different index sets, even if they are equal as strings.

## Constraints

- `S` is a string of digits with length between `1` and `10^6`, inclusive

## Input

The input is given from Standard Input in the following format:

```text
S
```

## Output

Output the number of subsequences of `S` that are 1122-strings, modulo
`998244353`.

## Sample Input 1

```text
1122
```

## Sample Output 1

```text
5
```

The valid subsequences are:

- `12` from positions 1 and 3
- `12` from positions 1 and 4
- `12` from positions 2 and 3
- `12` from positions 2 and 4
- `1122` from positions 1, 2, 3, and 4

## Sample Input 2

```text
2025
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
0777468889971
```

## Sample Output 3

```text
30
```

## Solution

Split the answer by the digit pair `(d, d+1)` for `d = 0, 1, ..., 8`. A 1122-string uses only
one such pair: first half is all `d`, second half is all `d+1`. Sum `play(d+1)` over
`d = 0..8` (9 passes total).

### Fix the last `d` in the first half

Scan `S` left to right for a fixed pair `(x-1, x)`. When `S[i] = x-1`, treat index `i` as
the **last chosen position** among the first-half digits `(x-1)`.

- `pref` — count of `(x-1)` in `S[0..i]` (after including `i`)
- `suf[i]` — count of `x` in `S[i..n-1]` (precomputed by a backward scan)

Every valid subsequence anchored this way has even length `2d`:

- pick `d` positions with digit `(x-1)`, **including** `i`
- pick `d` positions with digit `x`, all from the suffix starting at `i`

### The per-length sum

For a fixed length parameter `d`:

```text
C(pref - 1, d - 1)  ·  C(suf[i], d)
```

Choose the other `d-1` copies of `(x-1)` before `i`, and choose `d` copies of `x` from
the suffix.

Sum over all feasible `d`:

```text
Σ_{d=1}^{w} C(pref - 1, d - 1) · C(suf[i], d)     where w = min(pref, suf[i])
```

### Why `C(pref - 1 + suf[i], pref)`

Substitute `j = d - 1`:

```text
Σ_{j=0}^{w-1} C(pref - 1, j) · C(suf[i], j + 1)
```

Apply Vandermonde with `a = pref - 1`, `b = suf[i]`, `n = pref`:

```text
C(a + b, n) = Σ_{k=0}^{n} C(a, k) · C(b, n - k)
```

So:

```text
C(pref - 1 + suf[i], pref)
  = Σ_{k=0}^{pref} C(pref - 1, k) · C(suf[i], pref - k)
```

Replace `k = pref - d` (so `pref - k = d`):

```text
= Σ_{d=0}^{pref} C(pref - 1, d - 1) · C(suf[i], d)
```

The term `d = 0` is zero (`C(pref - 1, -1) = 0`). For `d > w`, at least one binomial
is zero, so the sum equals the original sum over `d = 1..w`.

Therefore, at each `(x-1)`-position `i`:

```text
cur = C(pref - 1 + suf[i], pref)
```

No inner loop over `d` is needed.

### Full algorithm

1. Precompute factorials / inverse factorials modulo `998244353` for `nCr`.
2. For each `x = 1..9`:
   - build `suf[i]` = number of digit `x` in `S[i..]`
   - scan `i = 0..n-1`; when `S[i] = x-1`, increment `pref` and add
     `C(pref - 1 + suf[i], pref)` to the answer
3. Output the total modulo `998244353`.

### Example (`S = "1122"`, pair `(1, 2)`)

| `i` | `S[i]` | `pref` | `suf[i]` | contribution |
| --- | ------ | ------ | -------- | ------------ |
| 0   | `1`    | 1      | 2        | `C(2, 1) = 2` |
| 1   | `1`    | 2      | 2        | `C(3, 2) = 3` |
| 2,3 | `2`    | —      | —        | skip |

Total `2 + 3 = 5`, matching sample 1.

### Complexity

- Time: `O(9 · n)` — nine digit passes, `O(1)` work per index
- Space: `O(n)` for `suf` and factorial tables