# C. Mister B and Boring Game

[Problem link](https://codeforces.com/problemset/problem/820/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Sometimes Mister B has free evenings when he doesn't know what to do.
Fortunately, Mister B found a new game, where the player can play against
aliens.

All characters in this game are lowercase English letters. There are two
players: Mister B and his competitor.

Initially the players have a string `s` consisting of the first `a` English
letters in alphabetical order (for example, if `a = 5`, then `s` equals
`"abcde"`).

The players take turns appending letters to string `s`. Mister B moves first.

Mister B must append exactly `b` letters on each his move. He can arbitrarily
choose these letters. His opponent adds exactly `a` letters on each move.

Mister B quickly understood that his opponent was just a computer that used a
simple algorithm. The computer on each turn considers the suffix of string `s`
of length `a` and generates a string `t` of length `a` such that all letters
in the string `t` are distinct and don't appear in the considered suffix.
From multiple variants of `t` the lexicographically minimal is chosen (if
`a = 4` and the suffix is `"bfdd"`, the computer chooses string `t` equal to
`"aceg"`). After that the chosen string `t` is appended to the end of `s`.

Mister B soon found the game boring and came up with the following question:
what can be the minimum possible number of different letters in string `s` on
the segment between positions `l` and `r`, inclusive. Letters of string `s`
are numbered starting from `1`.

## Constraints

- `1 <= a, b <= 12`
- `1 <= l <= r <= 10^9`

## Input

One line with four space-separated integers: `a`, `b`, `l` and `r`.

## Output

Print one integer — the minimum possible number of different letters in the
segment from position `l` to position `r`, inclusive, in string `s`.

## Samples

### Sample 1

Input:

```
1 1 1 8
```

Output:

```
2
```

One optimal strategy generates `s = "abababab..."`.

### Sample 2

Input:

```
4 2 2 6
```

Output:

```
3
```

For example `s = "abcdbcaefg..."`, the segment is `"bcdbc"`.

### Sample 3

Input:

```
3 7 4 6
```

Output:

```
1
```

For example `s = "abczzzacad..."`, the segment is `"zzz"`.

## Ideas

### Optimal policy for Mister B

A standard optimal strategy: on every turn, **append the current last
letter of `s`, repeated `b` times** (“dump last letter”). Then let the
computer append its forced lex-smallest block of length `a`.

Because `a, b ≤ 12`, it is enough to simulate this construction on a long
enough prefix (until the B-block + computer-block pattern / alphabet
stabilizes), then read off the answer for `[l, r]` from that string
(compressing huge `r` via the period when needed).

### Case `b ≥ a`

After Mister B’s move the length-`a` suffix is monochromatic (`x…x`).
The computer’s reply `T` is then forced: the lex-smallest `a` distinct
letters ≠ `x`. The next dump-last move again makes a monochromatic suffix,
so play stays in this shape.

- Letters used in the repeating part: at most **`a + 1`**.
- So for any `[l, r]`, this strategy proves `answer ≤ a + 1`.
- Locally it can be better: if `[l, r]` sits inside one run of `x`’s, the
  answer can be `1` (sample 3: `a=3, b=7, [4,6] → "zzz"`).

(The same ceiling is achieved by always dumping one fixed letter when
`b ≥ a`; dump-last specializes to that after the suffix becomes mono.)

### Case `b < a`

Dumping the same *fixed* letter forever is **not** automatically safe:
Mister B cannot wipe the whole length-`a` window, so the computer’s reply
depends on a mixed suffix, and a bad choice of letter can inflate the
alphabet on `[l, r]` (e.g. after `"abcd"`, dumping `"aa"` makes
`[2,6] = "bcdaa"` with 4 letters, while the optimum is 3).

Dump-last still works here: it reuses the end letter and shrinks the
suffix alphabet toward the last part of the old window. Sample 2 under
dump-last:

```text
abcd → dd → abcddd → computer aefg
[2,6] = bcddd → {b,c,d} → 3
```

same as the statement’s `"bc"` construction. So `"bc"` is another optimal
first move for that query; dump-last is not contradicted.

### Takeaway

| Case | Strategy |
|------|----------|
| `b ≥ a` | Dump-last (= monochromatic suffix); `answer ≤ a + 1`, often less on short windows |
| `b < a` | Still dump-last (not “always one fixed letter”); simulate and query the segment |

The answer for `[l, r]` is the number of distinct letters that this
construction places on that segment (minimum over the usual handling of
the initial length-`a` prefix when `l` is small).

## Solution Summary

There are two kinds of moves for Mister B: a **reset move** and an
**ordinary move**.

Let the initial computer block be

```text
P = abc...  // the first a letters
```

### Reset move: try every initial letter

On a reset move, Mister B's block can be normalized to `b` copies of one
letter `x`. It is sufficient to try the `a` letters of `P`:

- using several different letters can only introduce more distinct letters;
- using a letter outside `P` introduces a completely new letter;
- the chosen `x` controls which letters the computer excludes from its next
  block.

There is no single best `x` for every query because the result depends on the
alignment of `[l, r]`. Therefore the solution simulates the game once for each
`x` and takes the minimum:

```go
ans := 26
for x := range a {
	ans = min(ans, play(a, b, l, r, x))
}
```

### Ordinary move: repeat the last letter

Suppose the computer has just appended a sorted block of distinct letters:

```text
T = t1 t2 ... ta
```

If `b < a`, after Mister B appends his block, the suffix still contains the
last `a-b` letters of `T`, including `ta`. Repeating `ta` therefore:

- introduces no new letter into the suffix;
- joins the adjacent `ta`, creating the longest possible run;
- keeps the suffix's set of letters as small as possible;
- consequently gives the computer the least harmful next block.

If `b >= a`, Mister B replaces the entire length-`a` suffix. Repeating one
letter makes that suffix contain only one distinct letter, which is again
optimal.

Thus, after the special reset move, ordinary moves append the current last
letter `b` times.

### Why the construction becomes periodic

A computer block is completely determined by its letter mask: its letters are
distinct and always written in sorted order. If the same mask appears again,
the same computer state has returned, so Mister B can repeat the same choices
and the following string is periodic.

There is one subtle case. When the repeated computer block is `P`, the game is
back at the reset state, so Mister B uses the selected `x` again rather than
the last letter of `P`. The simulation stops at this boundary; repeating the
detected cycle automatically repeats the original `x` move as required.

Prefix counts for every letter are stored while constructing the finite
prefix and one cycle. They allow the number of occurrences in `[1, pos]` to
be calculated even when `pos` is as large as `10^9`. A letter is present in
`[l, r]` exactly when

```text
count(r, letter) - count(l-1, letter) > 0
```

The answer for one `x` is the number of such letters, and the final answer is
the minimum over all `a` choices of `x`.

### Sample 6

For `a = 3`, `b = 1`, choose `x = 'b'`:

```text
abc | b | ade | e | abc | b | ...
```

The period is `abcbadee`. Positions `4..10` are

```text
badeeab
```

which contains four distinct letters: `{a, b, d, e}`. Trying only `x = 'c'`,
as in the original implementation, produces five. The essential point is not
to abandon the repeat-last rule, but to enumerate the special choice made at
each reset state.
