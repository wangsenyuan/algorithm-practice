# E - Popcount Sum 3

[Problem link](https://atcoder.jp/contests/abc406/tasks/abc406_e)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 450 points

## Problem Statement

You are given positive integers `N` and `K`.
Find the **sum**, modulo `998244353`, of all positive integers `x` that do not
exceed `N` and satisfy:

- the popcount of `x` is exactly `K`.

You are given `T` test cases; solve each of them.

`popcount(y)` is the number of `1` bits in the binary representation of `y`.
For example, `popcount(5) = 2`, `popcount(16) = 1`, `popcount(25) = 3`.

## Constraints

- `1 <= T <= 100`
- `1 <= N < 2^60`
- `1 <= K <= 60`
- `T`, `N`, and `K` are integers

## Input

```
T
case_1
case_2
...
case_T
```

Each test case:

```
N K
```

## Output

Output `T` lines. The `i`-th line should contain the answer for the `i`-th
test case.

## Samples

### Sample 1

Input:

```
1
20 2
```

Output:

```
100
```

Positive integers `<= 20` with popcount `2`:
`3, 5, 6, 9, 10, 12, 17, 18, 20`. Their sum is `100`.

### Sample 2

Input:

```
1
1234567890 17
```

Output:

```
382730918
```

## Solution

Use binary digit DP. Process the bits of `N` from the most significant bit to
the least significant bit and construct every valid `x` bit by bit.

The implementation stores the binary digits of `N` in `ds`. If `m` is the
number of bits, bit position `p` has value:

```text
2^(m-p-1).
```

The array `pw` stores these powers of two modulo `998244353`.

### DP states

For the already processed prefix, define:

```text
dp[c][e]
```

as the number of prefixes that:

- contain exactly `c` set bits;
- have tightness state `e`.

The tightness state is:

- `e = 1`: the prefix is exactly equal to the corresponding prefix of `N`;
- `e = 0`: the prefix is already smaller than the prefix of `N`.

Counting prefixes is not enough because the problem asks for their sum. Define
another DP:

```text
fp[c][e]
```

as the sum of the partial numeric values of all prefixes counted by
`dp[c][e]`. Every chosen bit is already multiplied by its value in the full
binary number.

Initially, the empty prefix is equal to the empty prefix of `N`, contains no
set bits, and has value zero:

```text
dp[0][1] = 1.
```

All other states are zero.

### Transition

Suppose the current bit of `N` is `v`, and choose `d` (`0` or `1`) for `x`.

If the current state is tight (`e = 1`) and `d > v`, this choice would make
`x > N`, so it is forbidden.

Otherwise, the next tightness is:

```text
ne = e
if e = 1 and d < v:
    ne = 0
```

The new popcount is:

```text
ni = c + d.
```

The number of prefixes is updated by:

```text
ndp[ni][ne] += dp[c][e].
```

Every old prefix keeps its previous partial value, so its existing sum is
copied:

```text
nfp[ni][ne] += fp[c][e].
```

If `d = 1`, the current bit contributes `2^(m-p-1)` to every one of the
`dp[c][e]` prefixes. Therefore add:

```text
2^(m-p-1) * dp[c][e].
```

The complete sum transition is:

```text
nfp[ni][ne] += fp[c][e]
if d = 1:
    nfp[ni][ne] += 2^(m-p-1) * dp[c][e].
```

All calculations are performed modulo `998244353`.

After finishing a bit, `dp` and `fp` are swapped with their next arrays, and
the old arrays are cleared. This rolling technique avoids storing a separate
DP layer for every bit.

### Final answer

After all bits are processed, every represented number is at most `N`.
Numbers with exactly `K` set bits may be either already smaller than `N` or
exactly equal to `N`, so the answer is:

```text
fp[K][0] + fp[K][1].
```

The constraints have `K >= 1`, so zero is never included in these states; all
counted numbers are positive.

### Correctness Proof

We prove by induction over the processed bit positions that:

- `dp[c][e]` equals the number of valid processed prefixes with popcount `c`
  and tightness `e`;
- `fp[c][e]` equals the sum of their partial numeric values.

The claim holds for the empty prefix because `dp[0][1] = 1`, its value is zero,
and every other state is empty.

Assume the claim holds before processing the current bit. For every state, the
algorithm tries both possible digits. It rejects exactly the choice that would
make a tight prefix exceed `N`, and computes the new tightness and popcount
directly from the chosen digit. Thus every valid next prefix is produced once
and assigned to the correct `ndp` state.

For the sums, every extension retains the old partial value. When the chosen
digit is one, the algorithm additionally adds the current bit value once for
each old prefix. Hence `nfp` is exactly the sum of the values represented by
the corresponding `ndp` state.

Therefore the invariant holds after every bit. At the end, the two tightness
states with popcount `K` contain exactly all positive integers `x <= N` with
`popcount(x) = K`, so their combined sum is the required answer.

### Complexity

There are at most `60` bits. For every bit, the algorithm processes at most
`K+1` popcount states, two tightness states, and two possible digits.

- Time: `O(K log N)` per test case
- Space: `O(K + log N)`
