# F. Fragmented Nim

[Problem link](https://codeforces.com/problemset/problem/2181/F)

**Contest:** [Codeforces contest 2181](https://codeforces.com/contest/2181)

## Problem

Classical Nim: `n` piles with `a_i` stones. Players alternate; Alice starts. A
player picks a non-empty pile and removes any positive number of stones. The
player who takes the last stone wins.

**Fragmented Nim** changes who chooses the pile: the opponent chooses which
non-empty pile the current player must play in; the current player still chooses
how many stones to remove (at least one).

- On Alice's turn: Bob chooses a non-empty pile, then Alice removes stones from it.
- On Bob's turn: Alice chooses a non-empty pile, then Bob removes stones from it.

Alice still moves first. Determine the winner with optimal play.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 2 * 10^5`
- `1 <= a_i <= 10^9`
- Sum of `n` over all test cases `<= 2 * 10^5`

## Input

```text
t
case_1
...
case_t
```

Each test case:

```text
n
a1 a2 ... an
```

## Output

For each test case print `Alice` or `Bob`.

## Sample 1

```text
Input
3
1 2 3

Output
Bob
```

## Sample 2

```text
Input
1
1

Output
Alice
```

## Sample 3

```text
Input
5
10 3 4 7 4

Output
Alice
```
