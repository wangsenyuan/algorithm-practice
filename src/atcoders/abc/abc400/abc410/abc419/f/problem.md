# F - All Included (ABC419)

**Contest:** [ABC419](https://atcoder.jp/contests/abc419) — AtCoder Beginner Contest 419  
**Task:** [https://atcoder.jp/contests/abc419/tasks/abc419_f](https://atcoder.jp/contests/abc419/tasks/abc419_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 550 points

## Problem Statement

You are given `N` lowercase English strings `S_1, S_2, ..., S_N` and an integer
`L`.

Find the number, modulo `998244353`, of length-`L` lowercase English strings that
contain all of `S_1, S_2, ..., S_N` as substrings.

A **substring** of `S` is a string obtained by deleting zero or more characters from
the beginning and zero or more characters from the end of `S`. For example, `ab`,
`bc`, and `bcd` are substrings of `abcd`, while `ac`, `dc`, and `e` are not
substrings of `abcd`.

## Constraints

- `1 <= N <= 8`
- `1 <= L <= 100`
- `N` and `L` are integers
- Each `S_i` is a string of length between `1` and `10`, inclusive, consisting of
  lowercase English letters
- `S_i != S_j` for `i != j`


### ideas
1. 如果 s[i]被s[j]完全包含，那么可以把s[i]丢弃掉（因为结果中出现了s[j], 那么肯定也出现了s[i])
2. dp[i][mask][k]表示到i为止，已经完全包含了mask表示的集合，且最后一个是s[k]时的方案数
3. 现在加入字符s[w], 如果s[k]和s[w]有一段*公共*的部分，那么起点加入的长度，会不同
4. 但是有一个的问题，就是中间可以掺入一些随意的字符，会造成重复计算
5. dp[i][?] 这个状态，应该是和整个匹配有关系的
6. 用trie来定位状态？不对。
7. 比如对于abc, 再加入一个d，其实就切换到了cde, 如果是加入c, 就回到了c
8. 

## Solution

We need count all length-`L` strings that contain every pattern `S_i` as a
substring. Since `N <= 8`, it is natural to keep a bitmask of which patterns have
already appeared. The remaining problem is: after appending one more character,
how do we know which patterns newly become suffixes of the current string?

That is exactly what the Aho-Corasick automaton provides.

### Aho-Corasick Automaton

Build a trie containing all patterns. Each trie node represents a prefix of some
pattern, and during string construction the current node represents the longest
suffix of the built string that is also a trie prefix.

The important fields are:

```go
node[v][c]  // next automaton state after reading character c from state v
last[v]     // bitmask of patterns that end at this state
```

When inserting pattern `S_i`, the code walks down the trie and marks the terminal
node:

```go
tr.last[cur] |= 1 << i
```

But a pattern can also end through a failure link. For example, if the current
suffix is `abcd`, then pattern `bcd` or `cd` might also end at the same position.
So during `Build()`, the code propagates matched-pattern masks through failure
links:

```go
tr.last[v] |= tr.last[link[v]]
```

After this propagation, `last[v]` means: if the generated string currently ends at
automaton state `v`, then all patterns in `last[v]` have just appeared ending at
this position.

The other key part of `Build()` is filling missing transitions:

```go
if u == 0 {
	tr.node[v][i] = tr.node[link[v]][i]
}
```

This makes `node[v][c]` a complete automaton transition. The DP can append any
character in `O(1)` without manually following failure links.

### DP State

Let:

```text
dp[mask][v]
```

be the number of ways to build the current prefix such that:

- `mask` is the set of patterns already contained as substrings,
- `v` is the current Aho-Corasick state.

Initially:

```go
dp[0][0] = 1
```

The empty string has matched no patterns and starts at the trie root.

For each of the `L` positions, try every next lowercase letter:

```go
to := tr.node[v][i]
mask1 := mask | tr.last[to]
ndp[mask1][to] += dp[mask][v]
```

This transition has two jobs:

- `to` updates the suffix state after appending the new character.
- `tr.last[to]` tells which patterns newly appeared at this exact ending
  position.

So the DP never scans the built string and never checks all patterns manually.
The automaton has already compressed all suffix-matching information into the
state `v`.

### Answer

After exactly `L` characters have been generated, all patterns must have appeared,
so the required mask is:

```go
T - 1
```

where `T = 1 << N`. The final automaton state can be anything, so the answer is:

```go
sum(dp[T-1][v]) over all states v
```

### Complexity

The number of automaton states is at most the total length of all patterns plus
one, at most `8 * 10 + 1`.

For each of `L` positions, the DP iterates over:

- `2^N` masks,
- all automaton states,
- `26` next letters.

The total complexity is:

```text
O(L * 2^N * states * 26)
```

which fits easily under the constraints. The memory usage is
`O(2^N * states)`.
