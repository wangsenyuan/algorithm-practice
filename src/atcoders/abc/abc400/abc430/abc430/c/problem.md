# C - Truck Driver

[Problem link](https://atcoder.jp/contests/abc430/tasks/abc430_c)

**Contest:** [AtCoder Beginner Contest 430](https://atcoder.jp/contests/abc430)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

> In AtCoder Country, there is a rule that "a truck driver must take a break of at least B minutes
> when driving for A minutes or more."

You are given a string `S` of length `N` consisting of `a` and `b`, and positive integers `A` and
`B`. Find the number of integer pairs `(l, r)` that satisfy all of the following conditions:

- `1 <= l <= r <= N`
- The number of `a` in the substring from the `l`-th character through the `r`-th character of `S`
  is greater than or equal to `A`.
- The number of `b` in the substring from the `l`-th character through the `r`-th character of `S`
  is less than `B`.

## Constraints

- `1 <= N <= 3 * 10^5`
- `1 <= A, B <= N`
- `S` is a string of length `N` consisting of `a` and `b`.
- All input numbers are integers.

## Input

```text
N A B
S
```

## Output

Print the answer.

## Sample Input 1

```text
11 4 2
abbaaabaaba
```

## Sample Output 1

```text
3
```

The pairs `(l, r)` that satisfy the conditions are `(4, 8)`, `(4, 9)`, and `(5, 9)`.

## Sample Input 2

```text
13 1 2
bbbbbbbbbbbbb
```

## Sample Output 2

```text
0
```

There are no pairs `(l, r)` that satisfy the conditions.

## Solution

Use prefix counts and count valid left endpoints for every right endpoint.

Let:

- `pa[i]` be the number of `a` in the prefix `S[0:i]`;
- `pb[i]` be the number of `b` in the prefix `S[0:i]`.

For a substring `[l, r]` in zero-based half-open form `[l, r+1)`, the conditions are:

```text
pa[r+1] - pa[l] >= A
pb[r+1] - pb[l] < B
```

The first condition is equivalent to:

```text
pa[l] <= pa[r+1] - A
```

For a fixed right endpoint, all valid starts for the `a` condition form a prefix of possible starts.
The implementation maintains those starts with a moving pointer `l`.

### Maintaining starts with enough `a`

During the scan, `s1` is the current number of `a` in `S[0:r+1]`.

The pointer `l` is advanced while removing `S[l]` still leaves at least `A` letters `a` in the
suffix ending at the current `r`. Every such `l + 1` becomes a start position that satisfies the
`a` condition for this and all later right endpoints.

For every active start position `p`, store its prefix `b` count in a Fenwick tree. In the code this
value is:

```text
p - number_of_a_before_p
```

because among the first `p` characters, everything that is not `a` is `b`.

### Filtering by the number of `b`

For the current right endpoint, the prefix `b` count is:

```text
pb[r+1] = (r + 1) - s1
```

The substring must have fewer than `B` letters `b`:

```text
pb[r+1] - pb[p] <= B - 1
```

So the active start must satisfy:

```text
pb[p] >= pb[r+1] - (B - 1)
```

The code calls this lower bound `w`:

```text
w = (r + 1 - s1) - (B - 1)
```

The Fenwick tree stores counts by `pb[p]`, so:

- first count all active starts with `get(n)`;
- if `w >= 0`, subtract starts with `pb[p] <= w - 1`;
- if `w < 0`, subtract nothing, because every prefix `b` count is already at least `0`.

This boundary case is important. When `w < 0`, all currently active starts satisfy the `b`
condition. For example, with `S = "aa"` and `A = 1, B = 2`, every substring containing an `a` is
valid, and the answer is `3`.

### Correctness

For each right endpoint `r`, the moving pointer inserts exactly the start positions whose substring
ending at `r` has at least `A` letters `a`. Once a start position becomes valid for the `a`
condition, it remains valid for all later right endpoints, so inserting it once is sufficient.

Among those active starts, the Fenwick tree counts them by their prefix `b` count. The derived
inequality `pb[p] >= pb[r+1] - (B - 1)` is exactly the condition that the substring has fewer than
`B` letters `b`. Therefore the Fenwick query adds exactly the starts that satisfy both conditions
for this fixed `r`.

Summing this count over all right endpoints counts every valid pair `(l, r)` once, at its own right
endpoint, and counts no invalid pair.

### Complexity

Each start position is inserted into the Fenwick tree at most once. Each right endpoint performs
`O(log N)` Fenwick work.

Total complexity is `O(N log N)`, and memory usage is `O(N)`.
