# C. RemovevomeR

[Problem link](https://codeforces.com/problemset/problem/2241/C)

**Contest:** [Codeforces Round 1107 (Div. 3)](https://codeforces.com/contest/2241)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given a binary string `s` consisting only of the characters `0` and `1`.

In one operation, you can do the following:

- Choose a substring of `s` that is a palindrome of length at least `2`.
- Delete exactly one character from this chosen substring.

The remaining parts of the string are then concatenated to form the new string `s`.

Find the minimum possible length of the string `s` that can be achieved after applying this operation any number of times (possibly zero).

A string `a` is a substring of a string `b` if `a` can be obtained from `b` by the deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.

A string `a` of length `m` is said to be a palindrome if `a_i = a_{m+1-i}` for all `1 <= i <= m`.

## Input

The first line contains a single integer `t` (`1 <= t <= 100`) — the number of test cases.

The first line of each test case contains a single integer `n` (`1 <= n <= 100`) — the length of the binary string `s`.

The second line contains a binary string `s` of length `n`. It is guaranteed that each character of `s` is either `0` or `1`.

## Output

For each test case, print the minimum possible length of the string `s` that can be achieved after applying the operation any number of times.

## Example

### Input

```text
4
4
0000
3
110
6
110011
6
101100
```

### Output

```text
1
2
1
1
```

### Note

In the first test case, the initial string is `0000`. One sequence of operations is:

- Choose the palindromic substring `0000`. Delete one `0`. The string becomes `000`.
- Choose the palindromic substring `000`. Delete one `0`. The string becomes `00`.
- Choose the palindromic substring `00`. Delete one `0`. The string becomes `0`.

The string `0` contains no palindromic substrings of length at least `2`, so no further operations can be performed. The minimum possible length is `1`.

In the second test case, the initial string is `110`.

- Choose the palindromic substring `11`. Delete one `1`. The string becomes `10`.

The string `10` contains no palindromic substrings of length at least `2`, so no further operations can be performed. The minimum possible length is `2`.

## Solution

Any binary string of length at least `3` has a palindromic substring of
length at least `2`: a repeated pair `00`/`11`, or an alternating triple
`010`/`101`. Each operation deletes one character, so the string can
always be reduced to length `1` or `2`. Length `1` is terminal. Length
`2` is terminal only for `01` and `10`.

Those two strings arise precisely when `s` is two runs — a block of one
character followed by a block of the other. Inside a run you can shrink
`00`/`11`, but there is no wrapping palindrome `x…x` that can delete the
second run, so both characters survive. Any other string (one run, or
three or more) can be reduced to a single character.

The implementation walks the first run and, if a second run consumes the
rest of `s`, returns `2`; otherwise it returns `1`.

### Correctness sketch

Length `>= 3` always admits an operation, so the minimum is `1` or `2`.
Two runs compress to `01`/`10` and then stop. One run compresses to a
single character. With three or more runs a wrapping palindrome or an
extra run lets you delete until only one character remains.

### Complexity

One scan of `s`: `O(n)` time and `O(1)` extra memory. `n` is at most
`100`.
