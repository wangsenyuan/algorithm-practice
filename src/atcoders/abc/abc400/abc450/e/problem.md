# E - Fibonacci String (ABC450)

**Contest:** [ABC450](https://atcoder.jp/contests/abc450) — AtCoder Beginner Contest 450  
**Task:** [https://atcoder.jp/contests/abc450/tasks/abc450_e](https://atcoder.jp/contests/abc450/tasks/abc450_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

You are given strings `X` and `Y`. Define a sequence of strings
`S_1, S_2, ...` as follows.

- `S_1 = X`
- `S_2 = Y`
- For `i >= 3`, `S_i` is the concatenation of `S_{i-1}` and `S_{i-2}` in this
  order.

For each `i = 1, 2, ..., Q`, answer the following query.

You are given integers `L_i`, `R_i` and a character `C_i`. Find how many times
character `C_i` appears in the `L_i`-th through `R_i`-th characters of
`S_{10^18}`.

## Constraints

- `X` and `Y` are strings of lowercase English letters of length between `1` and
  `10^4`, inclusive.
- `1 <= Q <= 10^5`
- `1 <= L_i <= R_i <= 10^18`
- `C_i` is a lowercase English letter.
- All given numerical values are integers.

## Input

The input is given from Standard Input in the following format:

```text
X
Y
Q
L_1 R_1 C_1
L_2 R_2 C_2
...
L_Q R_Q C_Q
```

## Output

Output `Q` lines. The `i`-th line should contain how many times character `C_i`
appears in the `L_i`-th through `R_i`-th characters of `S_{10^18}`.

## Sample Input 1

```text
a
b
6
2 7 a
1 3 b
3 7 b
1 9 c
1 1000000000000000000 b
1000000000000000000 1000000000000000000 a
```

## Sample Output 1

```text
3
2
3
0
618033988749894848
1
```

`S_3`, `S_4`, and `S_5` are `ba`, `bab`, and `babba`, respectively.

`S_{10^18}` is `babbababbabba...`, and the second through seventh characters
contain three occurrences of `a`.


### ideas
1. 每个字符，也满足f[i] = f[i-1] + f[i-2] (fib 序列)
2. x, y, yx, yxy, yxyyx, yxyyxyxy, yxyyxyxyyxyyx
3. 可以看到奇数位置的x开始，偶数位置的y开始，所以，S(1e18)是以y开始的
4. 假设找到len(s[i]) >= R, 且是以y开始的序列i
5. f(i, r, c) 表示在s[i]中，找到在[1..r]中c的数量
6. 如果 r > len(s[i-1]), 那么 = c在s[i-1]中的数量 + f(i-2, r-len(s[i-1]), c)

## Solution

Use 0-based indices for the Fibonacci strings in the code:

```text
S[0] = X
S[1] = Y
S[i] = S[i-1] + S[i-2]
```

The query asks for character counts in a range of `S[10^18]`, but we never need
to build that huge string. For any character `c`, define:

```text
count(i, r, c) = number of c in the first r characters of S[i]
```

Then the answer to `[L, R]` is:

```text
count(j, R, c) - count(j, L-1, c)
```

where `j` is chosen large enough that the prefix of length `R` agrees with the
prefix of `S[10^18]`.

### Precomputation

For the base strings `X` and `Y`, build prefix character counts:

```go
pref0[pos][ch]
pref1[pos][ch]
```

These answer prefix queries inside `S[0]` and `S[1]` directly.

For larger Fibonacci strings, precompute:

```go
w[i]       = min(inf, len(S[i]))
freq[i][c] = min(inf, total count of c in S[i])
```

Both lengths and counts grow like Fibonacci numbers, so they are capped at
`inf = max(R) + 2`. Once values are above every query boundary, exact larger
values are unnecessary.

### Prefix Count Recursion

The key function is:

```go
play(i, r, c)
```

which returns `count(i, r, c)`.

For base strings:

```go
play(0, r, c) = pref0[r][c]
play(1, r, c) = pref1[r][c]
```

For `i >= 2`, remember the correct concatenation order:

```text
S[i] = S[i-1] + S[i-2]
```

So if the requested prefix lies fully inside the first part:

```go
if w[i-1] >= r {
	return play(i-1, r, c)
}
```

Otherwise take all occurrences from `S[i-1]` and recurse into the remaining
prefix of `S[i-2]`:

```go
return freq[i-1][c] + play(i-2, r-w[i-1], c)
```

This order matters. Since prefix queries depend on the left-to-right order,
using `S[i-2] + S[i-1]` would give correct total counts but wrong position-based
answers.

### Choosing `j`

The target is `S[10^18]`. With 0-based indexing, this corresponds to an odd index,
because:

```text
S[0] = X
S[1] = Y
```

The code finds the first precomputed length at least `R`:

```go
j := sort.Search(len(w), func(j int) bool {
	return w[j] >= R[i]
})
```

Then it adjusts to an odd index:

```go
if j&1 == 0 {
	j++
}
```

This works because once the Fibonacci strings are long enough, the prefix of a
larger string with the same parity is stable for the queried range.

In many cases this parity adjustment is redundant. If the search returns an even
`j > 0`, then after changing it to `j+1`, the recursion immediately sees
`w[j] >= R` and descends back into `S[j]`:

```go
if w[i-1] >= r {
	return play(i-1, r, c)
}
```

So the answer is the same. The adjustment is still useful for the base case: if
`R` is small enough that the first string `X` alone is long enough, the search may
return `j = 0`, but the huge target string starts like `Y`, not `X`. Changing to
the odd base side avoids using the wrong initial string.

Another way to think about this is simpler: instead of searching for the first
index whose length is at least `R`, we can start from a sufficiently large
precomputed index that has the same parity as `S[10^18]`. That is logically closer
to the statement, because we are directly asking for a prefix of the huge target
string. The recursive `play` function will then descend only as far as needed.

The implemented parity adjustment is a compact way to get the same effect while
keeping `j` small.

### Complexity

Only `O(log maxR)` Fibonacci levels are needed. Each query calls `play` twice,
and each call moves down by at least one level, so the time per query is
`O(log maxR)`.

The total complexity is:

```text
O((|X| + |Y|) * 26 + Q log maxR)
```

with small memory for the capped Fibonacci tables and base prefix counts.
