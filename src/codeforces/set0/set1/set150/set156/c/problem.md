# C. Cipher

[Problem link](https://codeforces.com/problemset/problem/156/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

Sherlock Holmes found a mysterious correspondence of two VIPs and decided to read
it. But the correspondence is encrypted.

Let a word `s` consist of `|s|` lowercase Latin letters. In one operation, choose a
position `p` (`1 <= p < |s|`) and do one of the following:

- replace `s_p` with the next letter in the alphabet and replace `s_{p+1}` with
  the previous letter;
- replace `s_p` with the previous letter and replace `s_{p+1}` with the next
  letter.

Letter `z` has no following letter, and letter `a` has no preceding letter, so
any operation that would require such a change is forbidden.

Two words have the **same meaning** if one can be transformed into the other by
zero or more operations.

For each given word, count how many **other** words (over all lowercase strings of
the same length, not just the input words) have the same meaning but differ in at
least one character. Print the answer modulo `1000000007` (`10^9 + 7`).

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of tests.

Each of the next `t` lines contains one word consisting of lowercase Latin letters.
Word length is from 1 to 100 inclusive. Lengths may differ.

## Output

For each word, print the number of different other words that coincide with it in
meaning, modulo `1000000007`.

## Example

### Input

```text
1
ab
1
aaaaaaaaaaa
2
ya
klmbfxzb
```

### Output

```text
1
0
24
320092793
```

## Note

For each letter, the following and preceding letters are defined in the usual way:
`b` follows `a`, `c` follows `b`, ..., `z` follows `y`; and `y` precedes `z`, ...,
`a` precedes `b`. Operations never change word length.

In the first sample, the only other word with the same meaning as `ab` is `ba`.

In the second sample, no other word has the same meaning, so the answer is `0`.

In the third sample, one operation can transform `klmbfxzb` into `klmcexzb` by
choosing `p = 4`: replace `b` with `c` and `f` with `e`. Word `ya` has the same
meaning as 24 other words: `xb`, `wc`, `vd`, ..., `ay`. Word `klmbfxzb` has
`3320092814` other equivalent words; modulo `10^9 + 7` this is `320092793`.

## Solution

Map each letter to a number: `a = 0`, `b = 1`, ..., `z = 25`.

### Invariant: total sum

Each operation changes two adjacent positions by `+1` and `-1`, so the total sum
of the word never changes.

So two words can have the same meaning only if they have the same length and the
same sum. In fact, for fixed length, **same sum is enough**: any two strings with
that sum can be transformed into each other by moving value between adjacent
positions (like sliding units left/right while staying in `[0, 25]`).

### Reduce the problem

For a word `s` of length `n` with sum `S`, the answer is:

```text
(number of length-n strings with sum S) - 1
```

Subtract `1` to exclude `s` itself. If `n = 1`, the sum alone determines the
only possible string, so the answer is `0`.

### Precomputed DP

Let `ways[n][sum]` be the number of length-`n` strings whose values add up to
`sum`:

```text
ways[0][0] = 1
ways[n][sum] = Σ ways[n-1][sum - c]   over c = 0..25 with sum - c >= 0
```

Precompute all `ways[n][sum]` for `n <= 100` and `sum <= 2500` once.

For each test word, compute `S`, then output `ways[n][S] - 1 (mod 1e9+7)`.

### Sanity checks

- `ab`: sum `1`, two strings of length 2 with sum `1` (`ab`, `ba`) → answer `1`.
- `aaaaaaaaaaa`: sum `0`, only one length-11 string → answer `0`.
- `ya`: sum `24`, 25 strings of length 2 with that sum → answer `24`.

### Complexity

- Precomputation: `O(100 * 2500 * 26)`
- Per test: `O(n)` to compute the sum
