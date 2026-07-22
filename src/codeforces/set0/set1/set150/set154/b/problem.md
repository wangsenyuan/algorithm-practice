# B. Colliders

[Problem link](https://codeforces.com/problemset/problem/154/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` colliders numbered `1..n`. Initially all are off.

Activating several colliders is safe iff all activated numbers are pairwise
relatively prime (`gcd = 1`). If two activated numbers share a common divisor
greater than `1`, there is a conflict.

Process `m` requests in order:

- `+ i` — try to activate collider `i`
- `- i` — try to deactivate collider `i`

Responses for `+ i`:

- `Success` — activated successfully
- `Already on` — already activated
- `Conflict with j` — some active collider `j` is not relatively prime to `i`
  (do not activate `i`; print any conflicting `j`)

Responses for `- i`:

- `Success` — deactivated successfully
- `Already off` — already deactivated

## Constraints

- `1 <= n, m <= 10^5`
- `1 <= i <= n` for each request

## Input

```text
n m
+ i | - i
...
```

## Output

Print `m` lines — one response per request.

## Example

```text
Input
10 10
+ 6
+ 10
+ 5
- 10
- 5
- 6
+ 10
+ 3
+ 6
+ 3

Output
Success
Conflict with 6
Success
Already off
Success
Success
Success
Success
Conflict with 10
Already on
```

The ninth request could also receive `Conflict with 3`.

## Solution

Two numbers conflict iff they share a prime factor. With up to `10^5`
operations, scanning all currently active colliders per request is too slow.
Instead, track **which active collider owns each prime**.

### Preprocessing

Linear sieve up to `n`:

- `lpf[x]` — least prime factor of `x`
- `ord[p]` — index of prime `p` in the primes list

### State

- `flag[x]` — whether collider `x` is on
- `status[ord[p]]` — if some active collider is divisible by prime `p`, store
  that collider's id; otherwise `-1`

### Operations

`check(x)`: walk prime factors of `x` via `lpf`. If any factor `p` already has
`status[ord[p]] >= 0`, return that conflicting collider id; else return `-1`.

`play(x, on)`: for each prime factor `p` of `x`, set `status[ord[p]] = x` when
turning on, or `-1` when turning off.

For each request:

- `+ x`
  - if `flag[x]`: `Already on`
  - else if `check(x) = j >= 0`: `Conflict with j` (do not activate)
  - else: `play(x, true)`, set `flag[x]`, `Success`
- `- x`
  - if not `flag[x]`: `Already off`
  - else: `play(x, false)`, clear `flag[x]`, `Success`

Because each number has `O(log x)` prime factors (with multiplicity handled by
repeated division), each request is fast enough for `n, m <= 10^5`.

### Complexity

- Sieve: `O(n)`
- Per request: `O(log x)` factorizations
- Total: `O(n + m log n)` time, `O(n)` memory
