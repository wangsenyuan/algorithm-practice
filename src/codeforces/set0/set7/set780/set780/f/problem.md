# F — Axel and Marston

## Problem

A couple of friends, Axel and Marston, are travelling across Bitland.
There are `n` towns and `m` directed roads.
Each road has a type:

- `0`: pedestrian road
- `1`: bike road

There may be multiple roads between the same pair of towns, as long as their types differ.
Self-loops are also allowed.

The friends start in town `1`.
They must choose road types according to an infinite binary word defined as follows:

- start with `A0 = 0`
- for each `k >= 0`, define `Ak+1 = Ak + invert(Ak)`

where `invert` flips every bit `0 <-> 1`.

So the first few strings are:

- `A0 = 0`
- `A1 = 01`
- `A2 = 0110`
- `A3 = 01101001`
- ...

The friends traverse roads one by one, always following the next bit of this infinite word.
If at some step there is no outgoing road of the required type, the trip stops.

We need the maximum possible trip length.
If it can be made strictly larger than `10^18`, print `-1`.

## Input

- first line: `n m` (`1 <= n <= 500`, `0 <= m <= 2n^2`)
- next `m` lines: `v u t`
  - directed edge `v -> u`
  - type `t` is `0` or `1`

## Output

- the maximum possible length
- or `-1` if the length can be made strictly greater than `10^18`

## Examples

### Example 1

Input

```text
2 2
1 2 0
2 2 1
```

Output

```text
3
```

### Example 2

Input

```text
2 3
1 2 0
2 2 1
2 2 0
```

Output

```text
-1
```

## Detailed Explanation

The main difficulty is that the route type sequence is not arbitrary.
It is a very structured infinite word:

- `0`
- `01`
- `0110`
- `01101001`
- ...

This is the Thue-Morse sequence.

The solution works by precomputing which vertices can be connected by whole blocks of this sequence, then greedily extending the answer from large blocks to small blocks.

### 1. Two families of blocks

Define:

- `A_k` = the string after `k` inverse-append steps
- `B_k = invert(A_k)`

Then:

- `A_0 = 0`
- `B_0 = 1`

and the key recurrence is:

- `A_{k+1} = A_k B_k`
- `B_{k+1} = B_k A_k`

Each block has length:

- `|A_k| = |B_k| = 2^k`

This recurrence is the whole reason the DP works.

### 2. Reachability matrices

For every `k`, define two boolean matrices:

- `P_k[v][u] = 1` if there exists a path from `v` to `u` whose edge-type sequence is exactly `A_k`
- `Q_k[v][u] = 1` if there exists a path from `v` to `u` whose edge-type sequence is exactly `B_k`

Base case:

- `P_0` is just the adjacency matrix of type-`0` edges
- `Q_0` is just the adjacency matrix of type-`1` edges

Now use the block recurrence:

- to follow `A_{k+1}`, first follow `A_k`, then follow `B_k`
- to follow `B_{k+1}`, first follow `B_k`, then follow `A_k`

So:

- `P_{k+1} = P_k * Q_k`
- `Q_{k+1} = Q_k * P_k`

where `*` is boolean matrix multiplication:

- `(X * Y)[v][u] = 1` iff there exists `w` such that `X[v][w] = 1` and `Y[w][u] = 1`

This gives all reachability information for blocks of lengths `1, 2, 4, 8, ...`.

### 3. Why this is enough to build the answer

Suppose we already know:

- current traveled length is `L`
- current reachable set is `S`

meaning:

- starting from town `1`, after exactly `L` required steps, we can end in any town from `S`

Now ask whether we can extend the route by another block of length `2^k`.

Which block comes next?

That depends on the parity of the number of `1` bits in `L`.
For the Thue-Morse sequence, the suffix starting at position `L` begins with:

- `A_k` if `popcount(L)` is even
- `B_k` if `popcount(L)` is odd

So:

- if `popcount(L)` is even, we advance `S` through `P_k`
- if `popcount(L)` is odd, we advance `S` through `Q_k`

Let the new reachable set be `S'`.

If `S'` is non-empty, then we can safely increase:

- `L += 2^k`
- `S = S'`

Otherwise, this block cannot be appended.

### 4. Why greedy from large to small works

This is exactly the same idea as binary lifting or greedy binary construction of a maximum number.

We try `k` from large to small:

- first try to add `2^59`
- then `2^58`
- ...
- then `2^0`

Whenever a block is feasible, we take it.

Why is this correct?

Because after all `P_k` and `Q_k` are precomputed, each choice answers:

- “Can I extend the current prefix by exactly `2^k` more steps?”

So processing powers of two from large to small is just greedily constructing the largest possible length in binary.

At the end, `L` is the maximum achievable route length below `2^60`.

Since `2^60 > 10^18`, this is enough:

- if we can exceed `10^18`, print `-1`
- otherwise the exact answer is `L`

### 5. Why the previous solution failed on sample 4

The buggy approach tried to recursively compute the answer from a single current node.
That is not enough.

At a given prefix length, we may be able to reach many towns.
Some of them may fail on the next big block, while others may continue much farther.

So the state must be:

- a whole reachable set of towns

not just one town.

The fixed solution follows the editorial exactly:

- precompute block reachability matrices
- keep the whole reachable set
- greedily extend it

That is why sample 4 works now.

### 6. Bitset optimization

Naively, each boolean matrix multiplication is `O(n^3)`.
Doing that for about `60` levels would be too slow.

But `n <= 500`, so we can store each row as a bitset.

Then:

- a row union becomes fast bitwise OR
- matrix multiplication becomes “for every set bit in row `i`, OR the corresponding row from the other matrix”

This is fast enough in practice for `n = 500`.

In the implementation:

- each row is stored as `[]uint64`
- `P[k]` and `Q[k]` are bitset matrices
- `advance` takes a reachable-set bitset and applies one of these matrices

### 7. Complexity

Let `w = 64` be the machine word size.

Using bitsets, the complexity is roughly:

- `O(60 * n^3 / w)`

which is fine for `n <= 500`.

Memory usage is:

- `O(60 * n^2 / w)`

### 8. Summary

The full solution is:

1. Build `P_0` and `Q_0` from edges of type `0` and `1`
2. Precompute `P_k` and `Q_k` using
   - `P_{k+1} = P_k Q_k`
   - `Q_{k+1} = Q_k P_k`
3. Start with:
   - current length `L = 0`
   - reachable set `S = {1}`
4. For `k` from `59` down to `0`:
   - decide whether the next block is `A_k` or `B_k` from parity of `popcount(L)`
   - compute the next reachable set
   - if it is non-empty, accept this block
5. If `L > 10^18`, print `-1`, otherwise print `L`

The key insight is:

- do DP on whole blocks `A_k / B_k`, not on individual characters
- and keep the whole reachable set, not a single endpoint
