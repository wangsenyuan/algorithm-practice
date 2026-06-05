# C. Genetic engineering

[Problem link](https://codeforces.com/problemset/problem/86/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

Physicist Woll changed his subject of study to bioinformatics. We think of a DNA
sequence as an arbitrary string of uppercase letters `A`, `C`, `G`, and `T`.

Let `w` be a long DNA sequence and let `s_1, s_2, ..., s_m` be a collection of
short DNA sequences.

The collection **filters** `w` if every position of `w` is covered by at least one
occurrence of a sequence from the collection. Occurrences may intersect or even
cover each other.

Formally, for each position `i` in `w`, there must exist indices `l, r` such that:

- `1 <= l <= i <= r <= |w|`,
- the substring `w[l...r]` is equal to one of `s_1, s_2, ..., s_m`.

Find the number of different DNA sequences of length `n` filtered by the given
collection.

The answer can be very large, so output it modulo `1000000009`.

## Input

The first line contains two integers `n` and `m`
(`1 <= n <= 1000`, `1 <= m <= 10`) — the length of the string and the number of
sequences in the collection.

Each of the next `m` lines contains one sequence `s_i`. Each sequence is
non-empty, has length at most `10`, and consists only of uppercase letters
`A`, `C`, `G`, and `T`.

The collection may contain identical strings.

## Output

Print a single integer — the number of length-`n` strings filtered by the
collection, modulo `1000000009`.

## Examples

### Input

```text
2 1
A
```

### Output

```text
1
```

### Input

```text
6 2
CAT
TACT
```

### Output

```text
2
```

## Note

In the first sample, the string must be filtered by `A`, so the only possible
string is `AA`.

In the second sample, exactly two different strings satisfy the condition.

### ideas
1. 这里是要求任何一段l..r都能找到某个s[i], 能够覆盖w[l..r]

## Solution

Build the answer string from left to right. At every step, we need to know whether
the prefix built so far is fully covered by occurrences of the given patterns.

The direct condition is about every position, but while scanning left to right it
is enough to remember only the uncovered suffix length.

Define:

```text
gap = current_length - longest_fully_covered_prefix_length
```

So:

- `gap = 0` means the whole generated prefix is already covered.
- `gap = 1` means only the last character is not guaranteed covered.
- `gap = 2` means the last two characters are not guaranteed covered.

When appending one new character, this uncovered suffix first grows by one:

```go
ngap := gap + 1
```

If the new character completes a pattern ending at the current position, and that
pattern is long enough to cover the whole uncovered suffix, then the prefix becomes
fully covered again:

```go
if ac.best[to] >= ngap {
	ngap = 0
}
```

This is the main DP idea.

### Aho-Corasick State

To know which pattern lengths end at the current position, use an Aho-Corasick
automaton over the given DNA strings.

Each trie node represents a prefix of some pattern. While generating the answer
string, the current automaton node represents the longest suffix of the generated
string that is also a trie prefix.

The transition:

```go
to := ac.node[v][c]
```

means: if the current automaton state is `v` and we append character `c`, the next
automaton state is `to`.

The code stores:

```go
best[v]
```

as the longest pattern length that ends at the current position when the automaton
is at node `v`.

When a pattern is inserted, only its terminal node is marked:

```go
tr.best[cur] = max(tr.best[cur], len(s))
```

But a pattern can also end through a suffix. For example, if patterns include:

```text
CAT
AT
T
```

then when the current suffix is `CAT`, the patterns `CAT`, `AT`, and `T` all end
at the current position.

That is why failure links are needed.

### Failure Link

For a trie node `v`, `link[v]` points to the node representing the longest proper
suffix of `str(v)` that is also a trie prefix.

Example:

```text
str(v) = CAT
proper suffixes = AT, T, empty
```

If `AT` exists as a trie prefix, then:

```text
link[CAT] = AT
```

If `AT` also has suffix `T`, then:

```text
link[AT] = T
```

During `Build()`, the code propagates `best` through these links:

```go
tr.best[v] = max(tr.best[v], tr.best[tr.link[v]])
```

After this propagation, `best[v]` is not only the longest pattern ending exactly
at node `v`; it is the longest pattern that is a suffix of the current automaton
state. In other words, it is the longest pattern ending at the current generated
position.

The build also fills missing transitions:

```go
if u == 0 {
	tr.node[v][i] = tr.node[tr.link[v]][i]
}
```

So every `node[v][c]` transition is ready in `O(1)`.

### DP

The DP state is:

```text
dp[v][gap]
```

where:

- `v` is the current Aho-Corasick state,
- `gap` is the uncovered suffix length.

Initially:

```go
dp[0][0] = 1
```

The empty prefix is fully covered and starts at the root.

For each position and each next DNA character:

```go
to := ac.node[v][c]
ngap := gap + 1
if ac.best[to] >= ngap {
	ngap = 0
}
ndp[to][ngap] += dp[v][gap]
```

If `ngap` becomes larger than the maximum pattern length, this state is discarded,
because no future pattern ending later can cover a hole that starts too far back:
all pattern lengths are at most `maxLen`.

After generating exactly `n` characters, the whole string must be covered, so only
states with `gap = 0` are valid:

```go
sum(dp[v][0])
```

over all automaton states `v`.

### Complexity

There are at most `1 + 10 * 10 = 101` automaton nodes, because there are at most
`10` patterns of length at most `10`.

The DP has:

```text
n * states * maxLen * 4
```

transitions, where `maxLen <= 10`. This easily fits the limits.
