# B2 - Sub-RBS (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2190/B2)

**Contest:** [Codeforces Round 1073 (Div. 1)](https://codeforces.com/contest/2190)

time limit: 2 seconds per test

memory limit: 256 MB

This is the hard version of Sub-RBS. Sum the scores of all **non-empty subsequences** of `s`.
The string `s` is not necessarily a regular bracket sequence, and `n` is smaller than in B1.

Sequence `a` is **better** than sequence `b` if either:

- `b` is a proper prefix of `a`, or
- at the first index where they differ, `a` has `(` and `b` has `)`.

For a bracket sequence `t`, its **score** is:

- `0` if `t` is not a regular bracket sequence (RBS)
- otherwise, the maximum length of an RBS subsequence `r` of `t` such that `r` is better than `t`
- `0` if no such `r` exists

Equivalently, the score is the length of the longest RBS subsequence of `t` that is better than `t`,
or `0`.

Find the sum of scores over all non-empty subsequences of `s`, modulo `998244353`.

## Constraints

- 1 <= t <= 30
- 1 <= n <= 100
- sum of `n` over all test cases <= 100
- `s` consists of `(` and `)`

## Input

The first line contains `t` — the number of test cases.

Each test case:

```text
n
s
```

## Output

For each test case, print one integer — the required sum modulo `998244353`.

## Sample Input

```text
5
1
(
6
()()()
6
(())()
8
(())()()
22
()()())()()(()()()((()
```

## Sample Output

```text
0
4
0
22
563070
```

## Note

In the first test case, the only non-empty subsequence is `(`, which is not an RBS, so the answer is
`0`.

In the second test case, for subsequence `g = ()()()`, choosing `r = (())` shows the score is `4`.
All other non-empty subsequences have score `0`.

## Solution

Consider one chosen subsequence `t` of the original string `s`.

If `t` is not a regular bracket sequence, its score is `0`, so we only need to
understand regular bracket sequences.

For a regular bracket sequence `t`, any better subsequence `r` cannot be longer
than `t`, because `r` is a subsequence of `t`. If `r` is shorter and better than
`t`, then at the first position where they differ, `r` must have `'('` while `t`
has `')'`. This means `r` is obtained by skipping some earlier `')'` of `t` and
using some later `'('` instead. To keep the final balance zero, we must also
delete one later `'('`. Therefore the best possible positive score is always:

```text
len(t) - 2
```

Thus every subsequence contributes either `0` or `len(t)-2`.

The remaining question is when a regular bracket sequence `t` has positive
score.

Suppose we delete an earlier `')'`. From that point onward, the balance becomes
one larger than before, so no prefix becomes invalid. Later we delete a `'('` to
bring the final balance back to zero. This is possible exactly when, after the
first `')'` of `t`, there are at least two selected `'('` characters. One of
them can be used to make `r` lexicographically better, and the extra deletion
keeps the resulting sequence regular with length `len(t)-2`.

Examples:

- `((()))` has no `'('` after the first `')'`, so its score is `0`.
- `(())()` has only one `'('` after the first `')'`, so its score is `0`.
- `()(())` has at least two `'('` after the first `')'`, so its score is `4`.
- `()()()` also has at least two `'('` after the first `')'`, so its score is `4`.

So the task becomes:

Count every subsequence `t` of `s` that is a regular bracket sequence and has at
least two `'('` after its first `')'`. Add `len(t)-2` for each such subsequence.

### DP state

Process the original string from left to right. For every prefix of `s`, keep
subsequences chosen from that prefix.

The DP state is:

```text
dp[len][bal][cnt][phase]
```

where:

- `len` is the length of the chosen subsequence.
- `bal` is its current bracket balance.
- `cnt` is the number of chosen `'('` after the first chosen `')'`, capped at `2`.
- `phase = 0` means no `')'` has been chosen yet.
- `phase = 1` means at least one `')'` has already been chosen.

The cap at `2` is enough because the final condition only asks whether there are
at least two such `'('`.

Initially:

```text
dp[0][0][0][0] = 1
```

For each character of `s`, there are two choices:

1. Skip it. The current implementation handles skipping by adding the newly
   created states into `dp` while keeping old states unchanged.
2. Take it. These transitions are stored in `ndp`, then merged into `dp`.

### Taking `'('`

If `phase = 0`, no `')'` has appeared yet. Taking `'('` only increases the
balance:

```text
len -> len + 1
bal -> bal + 1
cnt unchanged
phase remains 0
```

If `phase = 1`, then this `'('` is after the first `')'`, so it contributes to
the positive-score condition:

```text
len -> len + 1
bal -> bal + 1
cnt -> min(2, cnt + 1)
phase remains 1
```

### Taking `')'`

We may only take `')'` when the current balance is positive. Then:

```text
len -> len + 1
bal -> bal - 1
cnt unchanged
phase becomes 1
```

Once a `')'` has appeared, `phase` stays `1`.

### Adding the answer

After processing a character, every state newly created in `ndp` represents a
subsequence whose last character is the current character.

A subsequence contributes if:

```text
bal == 0
phase == 1
cnt == 2
```

`bal == 0` means it is a regular bracket sequence. `phase == 1` means it ended
after choosing at least one `')'`, which is necessary for any non-empty RBS.
`cnt == 2` means it has at least two `'('` after the first `')'`, so its score is
positive.

For such a subsequence, add:

```text
len - 2
```

to the answer.

### Correctness

For any fixed subsequence `t`, the DP records its length, balance, whether a
right bracket has appeared, and the number of later left brackets capped at two.
The transitions exactly match either skipping or taking each character, so every
subsequence is counted once.

If a counted subsequence has `bal == 0`, it is a regular bracket sequence because
the DP never allows the balance to become negative when taking `')'`.

The condition `cnt == 2` is equivalent to having at least two chosen `'('` after
the first chosen `')'`. By the characterization above, exactly those regular
bracket sequences have positive score, and their score is `len-2`.

Therefore the DP adds exactly the score of every subsequence of `s`, and the
final answer is the required sum modulo `998244353`.

### Complexity

There are `O(n)` possible lengths, `O(n)` possible balances, `3` values for
`cnt`, and `2` values for `phase`. Each input character scans all states.

The time complexity is:

```text
O(n^3)
```

The memory complexity is:

```text
O(n^2)
```
