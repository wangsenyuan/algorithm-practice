# Codeforces 17C - Balance

https://codeforces.com/problemset/problem/17/C

## Statement

Nick likes strings very much. Once he wrote a random string of characters `a`,
`b`, and `c` on a piece of paper and began to perform the following operations:

- take two adjacent characters and replace the second character with the first
  one;
- take two adjacent characters and replace the first character with the second
  one.

For example, from `abc`, one operation can produce `bbc`, `abb`, or `acc`.

For a string, let `|a|`, `|b|`, and `|c|` be the frequencies of characters `a`,
`b`, and `c`. A string is balanced if every pair of these frequencies differs by
at most `1`.

Find the number of different balanced strings that can be obtained from the
given string after applying the operations zero or more times. The initial string
should also be counted if it is balanced.

Print the answer modulo `51123987`.

## Input

The first line contains an integer `n` (`1 <= n <= 150`) — the length of the
given string.

The second line contains the string `s`, consisting only of characters `a`, `b`,
and `c`.

## Output

Print one integer: the number of different balanced reachable strings modulo
`51123987`.

## Samples

```text
4
abca
```

```text
7
```

```text
4
abbc
```

```text
3
```

```text
2
ab
```

```text
1
```

## Note

In the first sample, it is possible to get 51 different strings, but only 7 of
them are balanced: `abca`, `bbca`, `bcca`, `bcaa`, `abcc`, `abbc`, `aabc`.

In the second sample, the balanced reachable strings are `abbc`, `aabc`,
`abcc`.

In the third sample, the only balanced reachable string is `ab`.

## editorial

Consider the input string A of length n. Let's perform some operations from the problem statement on this string; suppose we obtained some string B. Let compression of string X be a string X' obtained from X by replacing all consecutive equal letters with one such letter. For example, if S = "aabcccbbaa", then its compression S' = "abcba".
Now, consider compressions of strings A and B - A' and B'. It can be proven that if B can be obtained from some string A using some number of operations from the problem statement, then B' is a subsequence of A', and vice versa - if for any strings A and B of equal length B' is a subsequence of A', then string B can be obtained from string A using some number of operations.
Intuitively you can understand it in this manner: suppose we use some letter a from position i of A in order to put letter a at positions j..k of B (again, using problem statement operations). Then we can't use letters at positions 1..i - 1 of string A in order to influence positions k + 1..n of string B in any way; also, we can't use letters at positions i + 1..n of string A in order to influence positions 1..j - 1 of string B.

Now we have some basis for our solution, which will use dynamic programming. We'll form string B letter by letter, considering the fact that B' should still be a subsequence of A' (that is, we'll search for B' in A' while forming B). For this matter we'll keep some position i in string A' denoting where we stopped searching B' in A' at the moment, and three variables kA, kB, kC, denoting the frequences of characters a, b, c respectively in string B. Here are the transitions of our DP:
1) The next character of string B is a. Then we may go from the state (i, kA, kB, kC) to the state (nexti, 'a', kA + 1, kB, kC).
2) The next character of string B is b. Then we may go from the state (i, kA, kB, kC) to the state (nexti, 'b', kA, kB + 1, kC).
3) The next character of string B is c. Then we may go from the state (i, kA, kB, kC) to the state (nexti, 'c', kA, kB, kC + 1).
Where nexti, x is equal to minimal j such that j ≥ i and A'j = x (that is, the nearest occurrence of character x in A', starting from position i). Clearly, if in some case nexti, x is undefined, then the corresponding transition is impossible.
Having calculated f(i, kA, kB, kC), it's easy to find the answer:  for all triples kA, kB, kC for which the balance condition is fulfilled.

Such a solution exceeds time and memory limits. Note that if string B is going to be balanced, then kA, kB, kC ≤ (n + 2) / 3, so the number of states can be reduced up to 27 times. But it's also possible to decrease memory usage much storing matrix f by layers, where each layer is given by some value of kA (or kB or kC). It's possible since the value of kA is increased either by 0 or by 1 in each transition.

The overall complexity is O(N4) with a quite small constant.

## Implementation notes

The accepted code follows the editorial idea, but two details are important.

First compress the original string. Let the compressed string be `A'`, and let
`m = len(A')`. For every compressed position `i` and character `x`, precompute:

```text
next[i][x] = first position >= i in A' whose character is x
```

If there is no such position, the value is `m`.

While building the target string `B`, the DP keeps:

```text
i     = position in A' used by the last character of B'
last  = last character written to B
a, b  = number of a and b characters already written
```

The number of `c` characters does not need to be stored in the array. During the
layer for already-built length `l`, it is simply:

```text
c = l - a - b
```

The transition tries to append a new character `x`.

If `x == last`, the compressed target `B'` does not change: we are only
extending the current run of `B`. Therefore `i` stays unchanged.

If `x != last`, then a new run is created in `B'`. We must match that run after
the previous matched position in `A'`, so the new position is:

```text
j = next[i + 1][x]
```

If `j == m`, there is no valid occurrence, so this transition is impossible.

This `last` distinction is the part that is easiest to miss. Always jumping to
`next` for every appended character is wrong, because writing the same character
again should not consume another run from `A'`.

## Balance bound

Only balanced final strings matter. In a balanced string of length `n`, no
character appears more than:

```text
h = (n + 2) / 3
```

So any transition with `a > h`, `b > h`, or `c > h` can be skipped.

This bound is about the target string, not the original string. We must not
prune using the original frequencies, because the operations can change the
number of each character.

## Layered memory

A direct memoized DP over `(i, last, a, b, c)` is too large near `n = 150`.
The implementation uses layers by the current target length `l`.

At each layer it stores only:

```text
dp[a][b][i][last]
```

and derives `c = l - a - b`. After processing all states of length `l`, it
throws that layer away and keeps only the next layer.

This reduces memory from roughly:

```text
O(m * h^3 * 3)
```

to:

```text
O(m * h^2 * 3)
```

The time complexity remains `O(n * h^2 * m * 3 * 3)`, which is `O(n^4)` in the
worst case with a small constant.

Finally, after exactly `n` characters have been built, sum all states whose
counts `(a, b, c)` satisfy:

```text
max(a, b, c) - min(a, b, c) <= 1
```

All additions are taken modulo `51123987`.
