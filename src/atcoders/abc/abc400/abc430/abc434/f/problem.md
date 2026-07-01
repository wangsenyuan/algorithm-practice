# F - Concat (2nd)

[Problem link](https://atcoder.jp/contests/abc434/tasks/abc434_f)

**Contest:** [Sky Inc, Programming Contest 2025 (AtCoder Beginner Contest 434)](https://atcoder.jp/contests/abc434)

time limit: 2 sec

memory limit: 1024 MiB

score: 575 points

You are given `N` strings `S_1, S_2, ..., S_N` consisting of lowercase English letters.

For every permutation `P = (P_1, P_2, ..., P_N)` of `(1, 2, ..., N)`, form the string obtained by
concatenating `S_{P_1}, S_{P_2}, ..., S_{P_N}` in that order.

Let `A_1, A_2, ..., A_{N!}` be those `N!` strings sorted in lexicographical order. Output `A_2`.

You are given `T` test cases; solve each of them.

## Constraints

- `1 <= T <= 1.5 * 10^5`
- `2 <= N <= 3 * 10^5`
- `T`, `N` are integers
- `S_i` is a lowercase string with length between `1` and `10^6 - 1`, inclusive
- For a single input, the sum of `N` over all test cases does not exceed `3 * 10^5`
- For a single input, the sum of `|S_i|` over all test cases does not exceed `10^6`

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case is given in the following format:

```text
N
S_1
S_2
...
S_N
```

## Output

Print `T` lines. The `i`-th line should contain the answer for the `i`-th test case.

## Sample Input 1

```text
3
3
abc
ac
ahc
4
aaa
a
aaaa
a
15
ks
sy
k
ysk
yks
ky
ksy
sk
syk
s
kys
sky
ys
yk
y
```

## Sample Output 1

```text
abcahcac
aaaaaaaaa
kksksykykysskskyssyksyyksykyskyys
```

### Note

For the first test case, `S = ("abc", "ac", "ahc")`. The six concatenations in lex order are:

`abcacahc`, `abcahcac`, `acabcahc`, `acahcabc`, `ahcabcac`, `ahcacabc`

So `A_2 = abcahcac`.

## Solution

First sort the strings by the usual minimum-concatenation order:

```text
X < Y  iff  X + Y < Y + X
```

Let the sorted order be `S'`. Concatenating `S'` gives `A_1`.

There are three cases for `A_2`:

1. If `N = 2`, the only other permutation is the reversed order, so output `S'_2 + S'_1`.
2. If some adjacent pair satisfies `S'_i + S'_{i+1} = S'_{i+1} + S'_i`, swapping that pair produces the same
   concatenated string as `A_1`. Since equal strings are both present in the sorted list of all permutations,
   `A_2 = A_1`.
3. Otherwise, every adjacent pair is strictly ordered. Then the second string is the smaller of only these two
   candidates:

```text
S'_1 S'_2 ... S'_{N-2} S'_N S'_{N-1}
S'_1 S'_2 ... S'_{N-3} S'_{N-1} S'_{N-2} S'_N
```

So after sorting, the implementation checks for an equal adjacent pair, then compares the two candidate
concatenations from swapping the last pair or the pair before it.

To keep sorting fast, the comparator does not build `X + Y` and `Y + X` every time. For every string, precompute
its Z-array. When comparing a longer string with a shorter one, the Z-array lets us compare the shifted parts of
the longer string in `O(1)` after scanning only the shorter prefix. Thus each comparison costs
`O(min(|X|, |Y|))`.

### Z-array comparator

The comparator needs to decide whether:

```text
X + Y < Y + X
```

The implementation first ensures `len(X) >= len(Y)`. If not, it swaps the arguments and negates the answer:

```text
cmp(X, Y) = -cmp(Y, X)
```

Now write `|X| = a` and `|Y| = b`, where `a >= b`.

The two concatenations are:

```text
X + Y = X[0:b]     + X[b:a]       + Y
Y + X = Y          + X[0:a-b]     + X[a-b:a]
```

Compare them in three parts:

1. Compare `X[0:b]` with `Y` by scanning `b` characters. If they differ, the answer is known.
2. If they are equal, compare `X[b:a]` with `X[0:a-b]`. These are two substrings of the same string `X`.
   The Z-array of `X` gives the longest common prefix of `X[b:]` and `X` as `Z_X[b]`.
   - If `Z_X[b] < a-b`, the first differing character is:

```text
X[b + Z_X[b]]  vs  X[Z_X[b]]
```

   - Otherwise, the two middle substrings are equal.
3. If the middle part is also equal, compare the remaining `Y` with `X[a-b:a]` by scanning `b` characters.

Only the first and last parts scan at most `|Y|` characters. The middle part uses one Z-array lookup, so the whole
comparison costs `O(min(|X|, |Y|))`.

## Correctness

Sorting by `X + Y < Y + X` is the standard exchange argument for the lexicographically minimum concatenation, so
the sorted order gives `A_1`.

If two adjacent sorted strings commute, swapping them gives another permutation with exactly the same concatenated
string. Therefore the first two strings in the sorted list of all permutation results are both `A_1`, and the
answer is `A_1`.

Assume no adjacent pair commutes. Then every adjacent relation in the sorted order is strict:

```text
S'_i S'_{i+1} < S'_{i+1} S'_i
```

Consider any permutation different from `S'`. Its inversion number is the number of pairs that appear in the
opposite order from `S'`.

If the inversion number is at least `2`, it cannot be `A_2`. Choose one inverted adjacent pair inside that
permutation and swap it back toward the sorted order. Because the sorted adjacent relation is strict, this swap
strictly decreases the concatenated string. The original permutation still has at least one other smaller string
before it, so it cannot be the second string overall.

Therefore, when `A_2 != A_1`, only permutations with exactly one inversion can matter. A permutation with exactly
one inversion is just `S'` with one adjacent pair swapped.

Now compare adjacent swaps. Swapping pair `i` gives:

```text
S'_1 ... S'_{i-1} S'_{i+1} S'_i S'_{i+2} ... S'_N
```

If `i <= N-3`, this candidate differs from the following last-pair swap only at or before position `i`:

```text
S'_1 ... S'_{N-2} S'_N S'_{N-1}
```

At that first differing adjacent block, the earlier swap has `S'_{i+1} S'_i`, while the sorted-prefix candidate
still has `S'_i S'_{i+1}`. Since `S'_i S'_{i+1} < S'_{i+1} S'_i`, the earlier swap is strictly larger. Thus no
swap before `N-2` can be the answer.

The only remaining adjacent swaps are:

```text
swap(N-1, N):     S'_1 ... S'_{N-2} S'_N S'_{N-1}
swap(N-2, N-1):   S'_1 ... S'_{N-3} S'_{N-1} S'_{N-2} S'_N
```

Hence, after the equal-adjacent case is excluded, `A_2` is the smaller of exactly these two candidates.

The algorithm follows exactly these cases, so it returns `A_2`.

## Complexity

Let `L = sum |S_i|`. Z-array preprocessing costs `O(L)`. Sorting costs `O(L log N)` total with the optimized
comparator, and the final candidate construction costs `O(L)`. The memory usage is `O(L + N)`.
