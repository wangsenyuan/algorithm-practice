# Problem

You are given a string `s` of length `n`, consisting of the first `k` lowercase English letters.

We define a `c`-repeat of some string `q` as a string consisting of `c` copies of `q`. For example, string `"acbacbacbacb"` is a `4`-repeat of string `"acb"`.

String `a` contains string `b` as a subsequence if `b` can be obtained from `a` by erasing some symbols.

Let `p` be a permutation of the first `k` lowercase English letters. Define `d(p)` as the smallest integer such that a `d(p)`-repeat of string `p` contains string `s` as a subsequence.

There are `m` operations of two types applied to `s` in order:

- `1 l r c`: replace all characters at positions from `l` to `r` by character `c`.
- `2 p`: for the given permutation `p` of the first `k` lowercase English letters, find `d(p)`.

Your task is to output answers for all operations of the second type.

## Input

The first line contains three integers `n`, `m`, and `k` (`1 <= n <= 200000`, `1 <= m <= 20000`, `1 <= k <= 10`) — the length of `s`, the number of operations, and the alphabet size.

The second line contains string `s`.

Each of the next `m` lines describes one operation:

- Type `1`: `1 l r c` (`1 <= l <= r <= n`, `c` is one of the first `k` lowercase English letters).
- Type `2`: `2 p`, where `p` is a permutation of the first `k` lowercase English letters.

## Output

For each query of the second type, output the value of function `d(p)`.

## Example

**Input**

```text
7 4 3
abacaba
1 3 5 b
2 abc
1 4 4 c
2 cba
```

**Output**

```text
6
5
```

## Note

After the first operation, string `s` becomes `abbbbba`.

In the second operation, the answer is `6`-repeat of `abc`:
`ABcaBcaBcaBcaBcAbc`.

After the third operation, string `s` becomes `abbcbba`.

In the fourth operation, the answer is `5`-repeat of `cba`:
`cbAcBacBaCBacBA`.

Uppercase letters denote occurrences of symbols from string `s`.

## Editorial

### 1. Reduce a query to adjacent pairs

For a fixed permutation `p`, imagine matching `s` inside:

```text
pppppp...
```

Let `pos[x]` be the index of character `x` inside `p`.

Suppose the previous matched character is `x = s[i-1]`, and the next character is `y = s[i]`.

- If `pos[x] < pos[y]`, then after matching `x`, `y` still appears later in the same copy of `p`.
- If `pos[x] >= pos[y]`, then `y` cannot be matched in the same copy anymore, so the matching must move to the next copy of `p`.

The first character always needs the first copy. After that, every adjacent pair that goes "backward" or stays at the same position starts one more copy.

Therefore:

```text
d(p) = 1 + count of adjacent pairs (x, y) in s where pos[x] >= pos[y]
```

Example after the first sample update:

```text
s = abbbbba
p = abc
pos[a]=0, pos[b]=1, pos[c]=2
```

Adjacent pairs:

```text
a->b: 0 < 1, same copy
b->b: 1 >= 1, new copy
b->b: 1 >= 1, new copy
b->b: 1 >= 1, new copy
b->b: 1 >= 1, new copy
b->a: 1 >= 0, new copy
```

There are `5` bad adjacent pairs, so `d(p) = 1 + 5 = 6`.

This means the real information needed for any query is only:

```text
cnt[i][j] = number of adjacent pairs i -> j in the current string
```

Then query `p` is:

```text
answer = 1 + sum cnt[i][j] over all i, j with pos[i] >= pos[j]
```

Since `k <= 10`, this summation is only `O(k^2)`.

### 2. Why direct simulation is too slow

The naive approach would be:

- For type `1`, update all positions in `[l, r]`.
- For type `2`, scan `s` from left to right while keeping a pointer in permutation `p`.

Both can take `O(n)`. With `n = 200000` and `m = 20000`, this is too slow.

We need to maintain the adjacent-pair counts under range assignment.

### 3. Aggregate for a substring

For any contiguous piece of the current string, store:

- `first`: its first character;
- `last`: its last character;
- `cnt[i][j]`: number of adjacent pairs `i -> j` fully inside that piece.

If two neighboring pieces `A` and `B` are joined, the new aggregate is:

```text
cnt = A.cnt + B.cnt
cnt[A.last][B.first] += 1
first = A.first
last = B.last
```

The extra pair is the single boundary pair between the last character of `A` and the first character of `B`.

For a uniform run `[l, r]` filled with character `c`, the aggregate is:

```text
first = c
last = c
cnt[c][c] = r - l
```

There are `r-l+1` characters in the run, so there are `r-l` adjacent pairs inside it.

### 4. Why not a dense segment tree

A segment tree can store this aggregate at every node and support range assignment with lazy propagation.

However, this implementation must fit Codeforces memory limits. A dense segment tree needs about `4*n` nodes. Each node would store a `10 x 10` matrix. In Go, using `int` for those counts makes this too large when `n = 200000`:

```text
4 * 200000 nodes * 100 ints per node
```

That is already hundreds of megabytes before pointers, tags, and runtime overhead.

### 5. Run-based treap

The submitted code instead stores the string as ordered runs of equal characters.

Each treap node represents one interval:

```text
[l, r] filled with character c
```

The treap is ordered by position. Its in-order traversal gives the current runs from left to right.

Each node stores the aggregate for its whole subtree:

- `first`, `last`;
- flattened `cnt[10*10]`, using `int32`;
- random priority `pri`;
- left and right children.

`int32` is enough because every count is at most `n-1 <= 199999`.

### 6. Recomputing one treap node

The subtree represented by a node is:

```text
left subtree + current run + right subtree
```

So `pull(node)` rebuilds the aggregate in that order.

When adding the next piece, if there was a previous piece, add the crossing pair:

```text
cnt[previous_last][next_first] += 1
```

Then add the next piece's internal matrix and update the current last character.

This is the same merge rule as the substring aggregate, applied to up to three pieces.

### 7. Split by position

`split(root, pos)` returns:

```text
left:  all runs strictly before pos
right: all runs starting at pos or after
```

If `pos` is outside the current root run, recurse into the left or right child.

If `pos` falls inside a run `[l, r]`, split that run into:

```text
[l, pos-1] and [pos, r]
```

Then merge the old children around those new run nodes. This preserves order and keeps every node as a uniform run.

The code may create zero-length boundary nodes if called exactly at a run boundary, but the earlier cases avoid that:

- `pos <= root.l` goes left without splitting the run;
- `pos > root.r` goes right without splitting the run;
- only `root.l < pos <= root.r` reaches the internal split.

### 8. Range assignment

For operation:

```text
1 l r c
```

Convert to zero-based positions and split twice:

```text
A, B = split(root, l)
B, C = split(B, r+1)
```

Now:

- `A` is before the update interval;
- `B` is the old content of `[l, r]`;
- `C` is after the update interval.

Discard `B`, create one new run `[l, r]` with character `c`, and merge:

```text
root = merge(merge(A, newRun), C)
```

All affected aggregates are rebuilt by `pull` while merging.

### 9. Query

For operation:

```text
2 p
```

Build `pos[]` for the permutation. The root already stores all adjacent-pair counts for the whole string, so compute:

```text
answer = 1
for i in alphabet:
    for j in alphabet:
        if pos[i] >= pos[j]:
            answer += root.cnt[i][j]
```

This exactly matches the formula from Section 1.

### 10. Complexity

Let `R` be the current number of treap nodes.

- Building from the initial string creates one node per maximal equal-character run.
- Type `1` does two splits and two merges, so it costs expected `O(k^2 log R)`.
- Type `2` costs `O(k^2)`.
- Memory is `O(k^2 * R)`.

The important practical difference from the dense segment tree is that `R` is based on runs created by updates, not `4*n` fixed tree nodes, and the matrix uses `int32` instead of Go `int`.
