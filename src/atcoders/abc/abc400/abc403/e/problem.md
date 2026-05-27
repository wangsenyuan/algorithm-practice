# Problem E - Forbidden Prefix

https://atcoder.jp/contests/abc403/tasks/abc403_e

**Time Limit:** 2 sec / **Memory Limit:** 1024 MiB

**Score:** 500 points

## Problem Statement

There are two multisets of strings, `X` and `Y`, both initially empty.

You are given `Q` queries to process in order. In the `i`-th query, you receive
an integer `T_i` and a string `S_i`.

- If `T_i = 1`, insert `S_i` into `X`.
- If `T_i = 2`, insert `S_i` into `Y`.

After processing each query, print the number of strings in `Y` that have no
element of `X` as a prefix.

## Constraints

- `Q` is an integer between `1` and `2 * 10^5`, inclusive.
- `T_i in {1, 2}`.
- Each `S_i` is a string of length between `1` and `5 * 10^5`, inclusive,
  consisting of lowercase English letters.
- `sum |S_i| <= 5 * 10^5`.


## Solution

Maintain the answer directly:

```text
tot = number of strings in Y that currently have no prefix in X
```

There are two kinds of updates.

### Adding a String to `Y`

When a new string `s` is inserted into `Y`, it contributes to `tot` only if no
string already in `X` is a prefix of `s`.

So we keep a trie `tx` containing all strings inserted into `X`.

To check whether `s` is valid:

1. Walk along `s` in `tx`.
2. If the next trie edge does not exist, then no longer prefix from `X` can
   match, so `s` has no prefix in `X`.
3. If we ever reach a terminal node, then some string in `X` is a prefix of
   `s`, so `s` must not be counted.

If `s` is valid, increment `tot` and insert `s` into another trie `ty`.

The trie `ty` stores only the `Y` strings that are currently counted in `tot`.
For each trie node, `cnt[node]` is the number of counted `Y` strings passing
through that node.

### Adding a String to `X`

When a new string `s` is inserted into `X`, some previously counted strings in
`Y` may become invalid: exactly those strings that have `s` as a prefix.

Those strings form the subtree rooted at the node reached by `s` in `ty`.

So:

1. Walk along `s` in `ty`.
2. If the path does not exist, then no currently counted `Y` string starts with
   `s`, so `tot` does not change.
3. Otherwise, let `cur` be the node for `s`. Then `ty.cnt[cur]` is the number of
   counted `Y` strings with prefix `s`.
4. Subtract this number from `tot`.
5. Remove that whole subtree from `ty`, because those strings will never become
   valid again. The set `X` only grows, so once a `Y` string has a prefix in
   `X`, it can never return to the answer.

The code removes the subtree by deleting the edge from the parent of `cur`.
While walking to that parent, it subtracts the removed count from the ancestor
`cnt` values, so future prefix-subtree counts remain correct.

Finally, insert `s` into `tx` as a terminal string.

### Why This Is Efficient

Each query walks through its string in one or both tries. When a subtree is
removed from `ty`, it is simply detached; the code does not traverse every node
inside that subtree.

Because the total length of all query strings is bounded by `5 * 10^5`, the
total complexity is:

```text
O(sum |S_i|)
```

with the same order of memory usage for the trie nodes.
