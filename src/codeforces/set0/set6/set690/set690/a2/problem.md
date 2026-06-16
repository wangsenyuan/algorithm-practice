# A2. Collective Mindsets (medium)

[Problem link](https://codeforces.com/problemset/problem/690/A2)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Heidi now knows how many brains she needs to take one home (see problem A1). She
also wants to know the minimum chest size that lets her **survive at all** — she
may leave empty-handed, but must not be killed.

The setup is the same as in A1:

- `N` guests: `N - 1` zombies (ranks `1 .. N-1`) and Heidi (rank `N`).
- The highest-ranked living guest proposes a brain split; at least half must accept,
  or the proposer dies and the next rank proposes.
- On a tie, the highest-ranked living voter counts twice and votes for their own
  offer.
- Zombies are rational: survive first, then maximize brains; reject any offer that
  is not strictly better than waiting for later rounds.

Heidi proposes first. Find the **minimum** number of brains in the chest so that
some proposal of hers is accepted (she does **not** need to receive any brains).

## Input

One integer `N` (`1 <= N <= 10^9`) — the number of attendees.

## Output

One integer — the minimum chest size that allows Heidi to survive.

## Examples

### Input

```text
1
```

### Output

```text
0
```

### Input

```text
3
```

### Output

```text
1
```

### Input

```text
99
```

### Output

```text
49
```

## Solution

Same pirate-voting rules as A1. Heidi (rank `N`) proposes first and may keep **0**
brains; she only needs her offer accepted by **≥ half** of attendees.

Work **bottom-up**: for each rank, ask “would this zombie survive if they became
proposer with `B` brains in the chest?” Some ranks can guarantee survival by
**unconditionally accepting** a higher rank’s offer — they know they would die if
they ever had to propose with nothing to offer.

### Who survives with 0 brains?

Simulate from low ranks upward:

- Ranks **1** and **2** survive as proposers with an empty chest.
- Rank **3** dies as proposer (no brains to buy a vote). So rank **3** accepts **any**
  offer from rank **4** to avoid ever proposing → rank **4** survives.
- Ranks **5, 6, 7** are doomed as proposers; they accept any offer from rank **8** →
  rank **8** survives.

Pattern: with **0** brains, ranks **1, 2, 4, 8, 16, …** (powers of two) survive.

### Who survives with `B` brains?

Induct on `B`:

| `B` | Ranks that can survive as proposer |
| --- | ---------------------------------- |
| 0 | `{1, 2, 4, 8, …}` |
| 1 | `{1, 2, 3, 4, 6, 10, …} = {1, 2} ∪ (2 + {1, 2, 4, 8, …})` |
| 2 | `{1, 2, 3, 4, 5, 6, 8, 12, …} = {1, 2, 3, 4} ∪ (4 + {1, 2, 4, 8, …})` |

In general, with **`B`** brains the survivors are:

```text
{1, 2, …, 2B}  ∪  (2B + {1, 2, 4, 8, …})
```

The right-hand set contains **only even** ranks (each is `2B` plus a power of two).

### Odd `N` — why `(N - 1) / 2`

If `N` is odd, it lies in `{1, 2, …, 2B}` for `B = (N - 1) / 2`. So Heidi (rank `N`)
first appears in the **left** block when the chest has exactly `(N - 1) / 2` brains.

```text
answer = (N - 1) / 2        (N = 1 → 0)
```

Examples: `N = 3 → 1`, `N = 99 → 49`.

### Even `N` — why `(N - P) / 2`

If `N` is even, it is **not** in the left block `{1, …, 2B}` (that block is all
integers up to `2B`, but the even rank `N` in the survivor set must come from the
**right** block: `N = 2B + 2^k` for some power of two `2^k`.

So:

```text
N = 2B + 2^k   →   B = (N - 2^k) / 2
```

To minimize `B`, subtract the **largest** power of two **≤ `N`** (equivalently the
largest `2^k ≤ N`):

```text
P = max { 2^k | 2^k ≤ N }
answer = (N - P) / 2
```

Examples:

- `N = 100`: `P = 64`, answer `(100 - 64) / 2 = 18`
- `N = 4`: `P = 4`, answer `(4 - 4) / 2 = 0`
- `N = 6`: `P = 4`, answer `(6 - 4) / 2 = 1`

### Code

```go
if n&1 == 1 {
    return (n - 1) / 2
}
var h int
for 1<<(h+1) <= n {
    h++
}
return (n - (1 << h)) / 2   // P = 1 << h
```

The loop finds the largest `h` with `2^(h+1) ≤ n`, so `P = 2^h` is the largest power
of two `≤ n`.

### Relation to A1

Same coalition; A1 also keeps one brain for Heidi → **`A1 = A2 + 1`** (for `N > 1`).

### Complexity

`O(log N)` time, `O(1)` space.