# B. Help King

[Problem link](https://codeforces.com/problemset/problem/98/B)

**Contest:** [Codeforces contest 98](https://codeforces.com/contest/98)

## Problem

There are `n` knights. A fair coin must be used to pick one knight uniformly at
random (probability `1/n` each), following an optimal tossing strategy that
minimizes the expected number of tosses.

Output that minimum expected number of tosses as an irreducible fraction
`a/b`.

Note: this is the archive-modified statement; the original contest version was
replaced because the author's solution was wrong.

## Constraints

- `1 <= n <= 10000`

## Input

```text
n
```

## Output

Print the expected number of tosses as an irreducible fraction `a/b` with no
leading zeroes.

## Sample 1

```text
Input
2

Output
1/1
```

## Sample 2

```text
Input
3

Output
8/3
```

## Sample 3

```text
Input
4

Output
2/1
```

## Sample 4

```text
Input
5

Output
18/5
```

## ideas

Let `f(n)` be the minimum expected number of fair coin tosses to pick one of
`n` outcomes uniformly. The optimal process is the Knuth–Yao free-node tree
(editorial / Knuth–Yao 1976):

- Start with `1` free node.
- Each toss doubles every free node.
- Whenever there are at least `n` free nodes, turn `n` of them into leaves
  (one per knight) and keep the remainder free.
- Paths that hit a leaf stop; residual free nodes recurse.

Let `E[k]` be the expected remaining tosses with `k` free nodes (`1 <= k < n`).
One toss expands to `2k` nodes; write `r = (2k) mod n`. Then

```text
E[k] = 1                         if r = 0
E[k] = 1 + (r / (2k)) · E[r]     if r > 0
```

The answer is `E[1]` as an irreducible fraction. Dependencies form a functional
graph (cycles with trees), so solve the linear equations over rationals along
each component.

### Special cases (corollaries, not a full substitute)

- `f(1) = 0`.
- When `n` is even, equal splitting is optimal and gives `f(n) = 1 + f(n/2)`
  (matches the DP, e.g. `f(6) = 11/3`).
- Plain rejection `h · 2^h / n` with `h = ceil(log2 n)` is **not** always
  optimal for odd `n`: for `n = 5` it gives `24/5`, while the DP gives `18/5`.
  It does match some odds (e.g. `n = 3` → `8/3`, `n = 7` → `24/7`).

## summary

### Goal

Minimum expected fair-coin tosses to pick one of `n` knights with equal
probability `1/n`. Output that expectation as a reduced fraction `a/b`.

### Process (Knuth–Yao free nodes)

Treat unfinished work as *free nodes* (partial random outcomes):

1. Start with `1` free node.
2. Each toss doubles every free node → count becomes `2k`.
3. Whenever there are at least `n` free nodes, peel off `n` leaves (one per
   knight); those paths finish.
4. Keep the leftover `r = (2k) mod n` free nodes and continue.

### DP

Let `E[k]` be expected remaining tosses from state `k` (`1 <= k < n`), and
`r = (2k) % n`:

```text
E[k] = 1                         if r = 0
E[k] = 1 + (r / (2k)) · E[r]     if r > 0
```

Answer is `E[1]`.

After one toss you always pay `1`. Probability `(2k-r)/(2k)` finishes;
probability `r/(2k)` continues in state `r`.

### Tiny examples

- `n = 2`: `k=1` → `r=0` → `E[1]=1` → `1/1`.
- `n = 4`: `E[2]=1`, `E[1]=1+E[2]=2` → `2/1`.
- `n = 5`: next is `(2k)%5`, orbit `1→2→4→3→1` (a cycle). Solving the system
  gives `E[1]=18/5`, not rejection's `24/5`.

### Why only walk from `1`

From any `k`, next state is unique: `(2k)%n` (or `0`). The graph is a
functional graph. From `1` you either hit `0` or enter a cycle — only that
orbit is needed.

### Solving a cycle (closed form)

On cycle `c[0],…,c[L-1]`:

```text
E[c[j]] = 1 + (c[j+1] / (2·c[j])) · E[c[j+1]]
```

Product of coefficients around the cycle is `1/2^L`, so:

```text
E[c0] = (Σ_j c[j] · 2^{L-j}) / (c[0] · (2^L - 1))
```

That avoids chaining thousands of rationals (e.g. cycle length ~4268 for
`n=8947`).

If the walk is `prefix → cycle`, compute `E` at the cycle entry, then
back-substitute along the prefix with the same recurrence.

### Mental picture

```text
start E[1]
   │
   ▼
follow k → (2k) mod n
   │
   ├─ hit 0  → back-substitute; E=1 when r=0
   └─ hit a cycle → closed form for E[cycle start]
                    then back-substitute along the prefix
```

One-line: each state is “how many unfinished partial outcomes you hold”; one
toss doubles them and peels multiples of `n`; expectation satisfies
`E[k]=1+(r/(2k))E[r]`; answer is `E[1]`, solved only along the orbit of `1`,
with a closed form on cycles.
