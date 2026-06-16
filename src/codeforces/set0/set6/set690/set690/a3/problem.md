# A3. Collective Mindsets (hard)

[Problem link](https://codeforces.com/problemset/problem/690/A3)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

After dinner, at midnight, the `N` attendees play a guessing game.

Each guest gets a number `n_i` (`1 <= n_i <= N`) written on their forehead. No one
sees their own number, but everyone sees all `N - 1` others. Numbers may repeat.

From midnight on, no communication is allowed. Everyone simultaneously guesses the
number on their own forehead:

- If **at least one** guest guesses correctly, **all** survive.
- If **everyone** guesses wrong, **all** die.

Heidi can reprogram every guest’s guessing strategy in advance. Design a strategy
so that **for every possible assignment** of forehead numbers, at least one guest
guesses correctly.

You do not output the strategy directly. For each test scenario you are given a
rank `R` and the `N - 1` numbers visible to that guest; output the number that
rank `R` should guess. The judge checks that your answers define a valid global
strategy.

## Input

The first line contains an integer `T` (`1 <= T <= 50000`) — the number of
scenarios.

Each scenario consists of two lines:

- Two integers `N` and `R` (`2 <= N <= 6`, `1 <= R <= N`) — total guests and the
  rank of the guest who must guess.
- `N - 1` integers: the forehead numbers of all **other** guests, listed in
  **increasing rank order** (every guest knows everyone’s rank).

## Output

For each scenario, print one integer — the guess for rank `R` given those visible
numbers.

## Examples

### Input

```text
4
2 1
1 2
2 1
2 1
2 2
1 2
2 2
2
```

### Output

```text
1
2
2
1
```

### Input

```text
2
5 2
2 2 2 6
6 4
3 2 2 1 2
```

### Output

```text
5
2
```

## Note

For `N = 2`, one valid strategy is:

- Rank `1` always guesses the number on rank `2`’s forehead.
- Rank `2` always guesses the opposite of rank `1`’s number (in `{1, 2}`,
  opposite means `3 - n`).

Then exactly one guest is always correct.

## Solution

Use the sum of all forehead numbers modulo `N`.

Before the game, assign each rank one residue class:

```text
rank R is responsible for total sums S where S ≡ R - 1 (mod N)
```

For a fixed complete assignment, the total sum `S` has exactly one residue modulo
`N`, so exactly one rank is responsible for that assignment. If that rank can make
its guess force the total sum into its assigned residue, then that rank will be
correct.

Suppose rank `R` sees the other `N - 1` numbers, and their sum is `sum`. If their
own hidden number is `x`, then:

```text
S = sum + x
```

Rank `R` wants:

```text
sum + x ≡ R - 1 (mod N)
```

So:

```text
x ≡ R - 1 - sum (mod N)
```

There is exactly one valid forehead number in `{1, 2, ..., N}` with this residue.
The code computes it as:

```go
guess := (r - 1 - sum%n + n) % n
if guess == 0 {
    return n
}
return guess
```

The residue `0` is represented by number `N`, because guesses must be in
`1 .. N`.

### Correctness

Consider any actual assignment of forehead numbers. Let:

```text
S = n_1 + n_2 + ... + n_N
```

There is exactly one rank `R` such that:

```text
S ≡ R - 1 (mod N)
```

For that rank, the visible sum is `S - n_R`. The strategy makes rank `R` guess
the unique value `x` satisfying:

```text
(S - n_R) + x ≡ R - 1 (mod N)
```

Since `S ≡ R - 1 (mod N)`, the true value `n_R` satisfies the same congruence.
Both `x` and `n_R` are in `{1, ..., N}`, and each residue modulo `N` appears
exactly once in that set, so `x = n_R`.

Therefore rank `R` guesses correctly. This proves that every possible assignment
has at least one correct guess.

### Complexity

For each scenario, the algorithm sums `N - 1` visible numbers.

```text
Time:  O(N)
Space: O(1) besides the input slice
```
