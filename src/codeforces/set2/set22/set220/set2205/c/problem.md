# Codeforces 2205C - Simons and Posting Blogs

https://codeforces.com/problemset/problem/2205/C

## Statement

There are `n` blogs. Blog `i` mentions `l_i` users in order:

```text
a_i = [a_{i,1}, a_{i,2}, ..., a_{i,l_i}]
```

You will post every blog exactly once, in any order. During the process, keep a
sequence `Q` of recently mentioned users.

When posting a blog, process its mentioned users from left to right:

- if the current user is already in `Q`, move that user to the front;
- otherwise, insert that user at the front.

After all blogs are posted, find the lexicographically smallest possible final
sequence `Q`.

## Input

Each test contains multiple test cases.

The first line contains `t` (`1 <= t <= 1000`).

For each test case:

- the first line contains `n` (`1 <= n <= 3000`);
- each of the next `n` lines starts with `l_i` (`1 <= l_i <= 3000`), followed by
  `l_i` integers `a_{i,1}, a_{i,2}, ..., a_{i,l_i}`
  (`1 <= a_{i,j} <= 10^6`).

It is guaranteed that the sum of `n` over all test cases is at most `3000`.
If `L = sum(l_i)` inside one test case, the sum of `L` over all test cases is
also at most `3000`.

## Output

For each test case, output the lexicographically smallest final `Q`. It must
contain each user that appears in at least one blog exactly once.

## Sample

```text
5
3
5 1 2 3 4 6
3 2 5 1
4 1 9 2 3
2
2 1 6
1 6
1
3 6 1 1
5
4 2 3 3 4
5 1 2 4 3 1
2 4 1
3 3 3 1
5 4 3 2 2 2
5
4 2 3 1 4
5 2 5 5 6 5
5 3 4 7 5 5
8 3 6 4 3 1 1 5 4
2 1 1
```

```text
1 5 2 3 9 6 4
6 1
1 6
1 3 2 4
1 4 3 2 5 6 7
```

## Solution

Think about one blog first. If we post it last, only the last occurrence of each
user inside this blog matters for the final order, because an earlier mention of
the same user will be moved again by the later mention.

So for every blog, reverse its list and remove duplicates while keeping the
first occurrence in this reversed order. For example:

```text
original: 6 1 1
reversed: 1 1 6
block:    1 6
```

After this conversion, each blog becomes a block of distinct users in exactly
the order it would contribute to the front of the final `Q` if no later blog
mentions those users again.

Now build the answer greedily.

At any step, already chosen users are fixed in the answer. For each remaining
blog block, remove users that are already fixed. Empty blocks can be ignored.
Among the non-empty blocks, choose the one that gives the smallest possible
prefix.

For two blocks `a` and `b`, the correct ordering test is:

```text
a should be before b iff a + b < b + a
```

This is the same comparison used in "smallest concatenation" problems. If `a+b`
is lexicographically smaller than `b+a`, then placing block `a` now gives a
better answer than placing block `b` now. The implementation sorts the current
blocks with this comparator and takes the first block.

After taking a block:

1. append all of its users to `res`;
2. mark them as fixed;
3. delete marked users from all remaining blocks;
4. repeat until no block remains.

This works because once a user is appended to the answer, every later blog that
mentions this user would move it again, so those later mentions must be ignored
when deciding the suffix after the fixed prefix.

The total number of mentioned users is at most `3000`, so repeatedly sorting and
filtering the blocks is fast enough. The comparator may scan two whole blocks,
so the implementation is comfortably within the constraints even with this
direct simulation.
