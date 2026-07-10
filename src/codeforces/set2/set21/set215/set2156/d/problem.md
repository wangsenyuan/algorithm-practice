# D. Find the Last Number

[Problem link](https://codeforces.com/problemset/problem/2156/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

This is an **interactive** problem.

There is a hidden permutation `p` of length `n`. You may ask at most `2n` queries of
the form:

- Choose `i` and `x` with `1 <= i <= n - 1` and `1 <= x <= 10^9`.
- The grader returns `0` if `p_i & x == 0`, and `1` otherwise.

You **cannot** query the last element `p_n` (because `i <= n - 1`).

Find the value of `p_n` using at most `2n` queries.

## Constraints

- `1 <= t <= 10^3`
- `2 <= n <= 2 * 10^4`
- Sum of `n` over all test cases does not exceed `2 * 10^4`

## Interaction

First read `t`. For each test case:

1. Read `n`.
2. Make queries by printing `? i x` and flushing. Then read one integer `0` or `1`.
3. Print the answer as `! x` and flush, then proceed to the next test case.

## Example

```text
Input (interleaved)
2
2
0
1
3
1
0
1

Output (interleaved)
? 1 1
? 1 2
! 1
? 1 3
? 1 2
? 2 3
! 2
```

### Note

- Test 1: hidden permutation `[2, 1]`, so `p_2 = 1`.
- Test 2: hidden permutation `[1, 3, 2]`, so `p_3 = 2`.

Empty lines in the sample are for readability only; they do not appear in real
interaction.

## Solution

Although the original problem is interactive, the implemented `solve` receives an `ask`
function so tests can simulate the judge.

The hidden array is a permutation of `1..n`, and we can query every position except the
last one. So the task is equivalent to finding the one value from `1..n` that is missing
among `p_1, ..., p_{n-1}`.

Find this missing value bit by bit.

Maintain:

- `res`: the low bits of `p_n` that have already been determined.
- `arr`: positions among `1..n-1` whose values match those low bits.

Initially `res = 0`, and every known position is still possible.

For bit `d`, query each position in `arr` with:

```text
x = 2^d
```

This splits `arr` into:

- `arr0`: known values with bit `d = 0`;
- `arr1`: known values with bit `d = 1`.

Now compare the observed size of `arr0` with how many numbers in `1..n` should have the
same low bits as `res` and bit `d = 0`.

Let:

```text
mod = 2^(d+1)
```

Numbers whose lowest `d+1` bits are equal to `res` are exactly:

```text
res, res + mod, res + 2*mod, ...
```

inside `1..n`, except that when `res = 0`, the first positive such number is `mod`, not
`0`.

This is what `count(n, mod, res)` computes:

```go
func count(n int, mod int, rem int) int {
    first := rem
    if first == 0 {
        first = mod
    }
    if first > n {
        return 0
    }
    return (n-first)/mod + 1
}
```

So `cnt0 = count(n, mod, res)` is the expected number of values in the full set `1..n`
whose current low bits would be `res` with bit `d = 0`.

Because exactly one value is missing:

- if `len(arr0) < cnt0`, the missing value is in the zero-bit group, so keep `arr0`;
- otherwise the missing value has bit `d = 1`, so set that bit in `res` and keep `arr1`.

After processing all bits up to `bits.Len(n)`, `res` is the missing value, which is
`p_n`.

The number of queries is bounded by:

```text
(n-1) + about half + about quarter + ... < 2n
```

because after each bit, only the matching group is kept. The time complexity is `O(n)`,
and the memory complexity is `O(n)`.
