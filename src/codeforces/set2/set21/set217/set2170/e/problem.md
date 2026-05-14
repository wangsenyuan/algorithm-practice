# E. Binary Strings and Blocks (Codeforces 2170E)

**Limits:** 2 seconds per test, 512 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2170/E](https://codeforces.com/problemset/problem/2170/E)

## Problem Statement

Define a **block** in a binary string as a continuous substring of characters of the same type that cannot be extended either to the left or to the right.

For example, in the string `110001111`, there are three blocks:

- `11`, from the `1`-st character to the `2`-nd character;
- `000`, from the `3`-rd character to the `5`-th character;
- `1111`, from the `6`-th character to the `9`-th character.

The substring from the `7`-th character to the `9`-th character, `111`, is not a block because it can be extended to the left.

The substring from the `1`-st character to the `5`-th character, `11000`, is not a block because it contains characters of different types.

A string is called **beautiful** if we can remove exactly one block from it so that the resulting string consists of an odd number of blocks.

For example:

- `110001111` is beautiful because we can remove the block from the `3`-rd to the `5`-th character, resulting in `111111`, which consists of one block;
- `1010` is beautiful because we can remove the block from the `1`-st to the `1`-st character, resulting in `010`, which consists of three blocks;
- `0000` is not beautiful because the only way to remove a block from it results in an empty string, which consists of `0` blocks.

You are given an integer `n` and `m` constraints. The `i`-th constraint is described by a pair of integers `l_i, r_i`.

For a string `s`, denote the substring from character `l` to character `r` inclusive as:

```text
s[l:r] = s_l s_{l+1} ... s_r
```

Count the number of binary strings `s` of length `n` such that for every `i` from `1` to `m`, the string `s[l_i:r_i]` is beautiful.

Since the answer may be huge, output it modulo `998244353`.

## Solution

First characterize when a binary string is beautiful.

Suppose a string has `k` blocks.

- If `k = 1`, removing the only block leaves the empty string, which has `0` blocks, so the string is not beautiful.
- If `k >= 2` and `k` is even, remove one end block.  The remaining string has `k - 1` blocks, which is odd.
- If `k >= 3` and `k` is odd, remove any middle block.  Its two neighboring blocks have the same character, so they merge, and the remaining string has `k - 2` blocks, which is odd.

Therefore:

```text
a binary string is beautiful  <=>  it has at least 2 blocks
                                <=> it contains both 0 and 1
                                <=> it has at least one adjacent change
```

So every constraint `[l, r]` simply says:

```text
there exists some j with l < j <= r and s[j-1] != s[j]
```

In other words, every constrained interval must contain at least one boundary where adjacent characters differ.

### Convert Constraints

For every right endpoint `r`, keep only the largest left endpoint among constraints ending at `r`:

```text
L[r] = max l over constraints [l, r]
```

Then take prefix maximums:

```text
L[r] = max(L[1], L[2], ..., L[r])
```

After this, when we build a prefix ending at position `r`, all constraints with right endpoint at most `r` are satisfied if the latest adjacent change is after `L[r]`.

Why?  A constraint `[l, r]` requires a change in:

```text
(l, r]
```

If the latest change position is greater than `l`, then this interval contains a change.  For all constraints already seen, it is enough to compare against the maximum such `l`.

The implementation uses `0`-based endpoints, so the code stores:

```go
L[r] = max(L[r], l)
```

and later uses `L[i-1]` while processing prefix length `i`.

### DP Meaning

Think of a string as being split into runs.  Let the last run start after the latest adjacent change.

Define:

```text
dp[i] = number of valid binary strings of length i
        whose last run starts at position i
```

Equivalently, when extending from some previous prefix, we choose where the last adjacent change before the end occurs.  The actual final bit is determined by alternating from the previous run, while the first run has two choices (`0` or `1`).

The code stores only the current `dp` value and a prefix sum array:

```text
fp[i] = dp[0] + dp[1] + ... + dp[i-1]
```

The initialization:

```go
dp = 2
fp[1] = 2
```

represents the two one-character strings: `0` and `1`.

### Transition

When processing prefix length `i`, the last change must be after `L[i-1]`.

So the previous split point can be any valid position in:

```text
L[i-1] + 1, L[i-1] + 2, ..., i
```

Using prefix sums, the number of choices is:

```text
dp = fp[i] - fp[L[i-1] + 1]
```

This is exactly the implementation:

```go
dp = sub(fp[i], fp[L[i-1]+1])
fp[i+1] = add(fp[i], dp)
```

All operations are modulo `998244353`.

### Why This Counts All Valid Strings

For any valid prefix, only the latest adjacent change matters for future constraints, because every future condition asks whether there is at least one change after some left endpoint.

If the latest change is too early, then some interval ending at the current position would be constant, so the string is invalid.

If the latest change is after the required boundary, all constraints ending here or earlier are satisfied.  Future characters only need to know this latest change position, and the prefix-sum DP sums over all possible valid previous positions.

Thus the recurrence counts exactly the binary strings satisfying all constraints.

### Complexity

For each test case:

```text
Build L: O(n + m)
DP:      O(n)
Memory:  O(n)
```

The total `n` and `m` over all test cases are bounded by `3 * 10^5`, so this fits easily.
