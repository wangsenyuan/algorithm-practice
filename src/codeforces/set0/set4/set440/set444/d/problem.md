# D. DZY Loves Strings

[Problem link](https://codeforces.com/problemset/problem/444/D)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

DZY loves strings, and he enjoys collecting them.

In China, many people like to use strings containing their names' initials, for
example: `xyz`, `jcvb`, `dzy`, `dyh`.

Once DZY found a lucky string `s`. A lot of pairs of good friends came to DZY
when they heard about the news. The first member of the `i`-th pair has name
`a_i`, the second one has name `b_i`. Each pair wondered if there is a substring
of the lucky string containing both of their names. If so, they want to find the
one with minimum length, which can give them good luck and make their friendship
last forever.

Please help DZY for each pair find the minimum length of the substring of `s`
that contains both `a_i` and `b_i`, or point out that such substring does not
exist.

A substring of `s` is a string `s_l s_{l+1} ... s_r` for some integers `l, r`
(`1 <= l <= r <= |s|`). The length of such substring is `r-l+1`.

A string `p` contains another string `q` if there is a substring of `p` equal to
`q`.

## Solution

All query strings have length at most `4`. Therefore, the total number of
different strings that can matter is small enough to enumerate from `s`:

```text
all substrings of s with length 1, 2, 3, or 4
```

For each such string `x`, store the sorted list:

```text
pos[x] = all starting positions where x appears in s
```

If either query string does not appear in `s`, the answer is immediately `-1`.

### Length of a Substring Covering Two Occurrences

Suppose string `a` starts at position `i`, and string `b` starts at position
`j`.

The shortest substring covering these two particular occurrences starts at:

```text
min(i, j)
```

and ends at:

```text
max(i + len(a), j + len(b)) - 1
```

So its length is:

```go
max(i+len(a), j+len(b)) - min(i, j)
```

This is implemented as `unionLen`.

The whole query asks for the minimum of this value over all occurrence pairs of
`a` and `b`.

### Light Query

If both strings appear only a moderate number of times, answer the query
directly from the two sorted occurrence lists.

Let:

```text
u = pos[a]
v = pos[b]
```

Iterate over the shorter list. For each occurrence `i` in `u`, the best
occurrence of `b` is one of the closest starts around `i` in `v`:

```text
first v[j] >= i
last  v[j] <  i
```

These are found by binary search:

```go
j := sort.SearchInts(v, i)
```

Then check `v[j]` and `v[j-1]` if they exist.

Why are only these two candidates enough?

For fixed `i`, moving `j` farther to the right only increases the right end of
the covering substring. Moving `j` farther to the left only decreases the left
end. Therefore, the best partner is the nearest occurrence on either side.

This gives:

```text
O(min(|pos[a]|, |pos[b]|) * log N)
```

per light query.

### Why This Can Still TLE

The bad case is a very frequent short string, for example:

```text
s = "aaaaaaaaaa..."
queries = ("a", "aaaa"), ("a", "aaaa"), ...
```

The occurrence list of `"a"` has size `O(N)`. If we scan it for every query, the
total complexity becomes `O(NQ)`, which is too slow.

So we split strings into:

- **light** strings: occurrence count at most `threshold`;
- **heavy** strings: occurrence count greater than `threshold`.

The code uses:

```go
const threshold = 320
```

There can be only `O(N / threshold)` heavy strings, because every heavy string
has more than `threshold` occurrences, and the total number of stored
occurrences over all length `1..4` substrings is `O(4N)`.

### Heavy Precomputation

For every heavy string `h` that appears in at least one query, precompute:

```text
heavyAns[h][x] = answer for pair (h, x)
```

for every string `x` that appears in `s`.

To do this efficiently, build two arrays over positions of `s`:

```text
prev[i] = nearest occurrence start of h at or before i
next[i] = nearest occurrence start of h at or after i
```

They are computed by two linear scans.

Then for every string `x`, iterate over all its occurrences `p`.

For this occurrence of `x`, the best occurrence of heavy string `h` is again one
of the nearest starts around `p`:

```text
prev[p]
next[p]
```

So update:

```go
ans[x] = min(ans[x], unionLen(prev[p], h, p, x))
ans[x] = min(ans[x], unionLen(next[p], h, p, x))
```

Doing this for all `x` costs `O(total occurrences) = O(N)` for one heavy string,
because every occurrence of every short substring is visited once.

### Answering Queries

For a query `(a, b)`:

1. If either string does not occur, answer `-1`.
2. If `a` is heavy, answer `heavyAns[a][b]`.
3. Else if `b` is heavy, answer `heavyAns[b][a]`.
4. Otherwise, use the light-query binary-search method.

This avoids revisiting a huge occurrence list for many queries.

### Complexity

Let `T` be the threshold.

Light queries cost:

```text
O(T log N)
```

each, because the shorter occurrence list has size at most `T`.

The number of heavy strings is:

```text
O(N / T)
```

Each heavy string costs `O(N)` to precompute, so heavy preprocessing costs:

```text
O(N^2 / T)
```

With `T` around `sqrt(N)`, this is easily fast enough for:

```text
N <= 50000, Q <= 100000
```

The implementation uses `T = 320`.
