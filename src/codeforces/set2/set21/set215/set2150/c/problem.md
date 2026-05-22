# C. Limited Edition Shop

https://codeforces.com/problemset/problem/2150/C

**Time Limit:** 2 seconds

**Memory Limit:** 256 megabytes

## Problem

There is a shop with `n` objects numbered from `1` to `n`, and only one copy of
each object.

According to you, object `i` has value `v_i`, which may be negative.

Alice and Bob each have their own preference order:

- Alice: `a_1, a_2, ..., a_n`
- Bob: `b_1, b_2, ..., b_n`

Alice's favorite remaining object is the earliest unsold object in her order,
and Bob's favorite remaining object is the earliest unsold object in his order.

For `n` turns, either Alice or Bob goes to the shop and buys their own favorite
object that is still available. After all `n` turns, the shop is empty.

Among all possible sets of objects Alice could have bought, find the maximum
possible sum of their values.

## Input

Each test contains multiple test cases.

The first line contains `t` (`1 <= t <= 10^4`) -- the number of test cases.

For each test case:

- The first line contains `n` (`1 <= n <= 2 * 10^5`).
- The second line contains `n` integers `v_1, v_2, ..., v_n`
  (`-10^9 <= v_i <= 10^9`).
- The third line contains Alice's preference order
  `a_1, a_2, ..., a_n`.
- The fourth line contains Bob's preference order
  `b_1, b_2, ..., b_n`.

Both `a` and `b` are permutations of `1, 2, ..., n`.

It is guaranteed that the sum of `n` over all test cases does not exceed
`2 * 10^5`.

## Output

For each test case, output one integer: the maximum possible sum of values over
all sets of objects that Alice could have bought.

## Example

```text
Input
8
3
1 -1 1
3 1 2
2 3 1
3
-2 5 2
3 1 2
2 3 1
3
-1 -2 -3
3 1 2
2 3 1
3
1000000000 1000000000 1000000000
3 1 2
2 3 1
4
5 -15 10 -5
2 4 3 1
1 4 2 3
4
-5 -5 -5 100
2 3 1 4
4 1 2 3
4
-1 -100 5 10
1 2 3 4
2 3 4 1
12
-4 6 10 10 1 -8 6 2 -8 -4 0 -6
11 12 7 3 6 8 1 5 10 2 9 4
7 5 3 6 1 2 8 12 9 4 10 11

Output
2
5
0
3000000000
10
85
14
24
```

## Note

In the first test case, Alice can buy the sets:

```text
[]
[1]
[3]
[3, 1]
[3, 1, 2]
```

where the listed order is the purchase order.

For example, for Alice to buy `[1]`, Bob can first buy object `2`, then object
`3`, and then Alice buys object `1`.

The best set by value is `[3, 1]`, with total value `v_3 + v_1 = 2`.

## Solution

First, rename the objects by Alice's order. After this normalization, Alice's
preference order is:

```text
1, 2, ..., n
```

Let `pos[x]` be the position of object `x` in Bob's order.

Suppose Alice wants to buy exactly the subset `S`. A necessary condition is:

```text
for every x not in S and y in S:
    if x < y, then pos[x] < pos[y]
```

Why? If `x < y`, Alice prefers `x` before `y`. If also `pos[x] > pos[y]`, then
Bob prefers `y` before `x`. In that case:

- Alice cannot buy `y` before `x`, because `x` is still available and comes
  earlier in Alice's order.
- Bob cannot buy `x` before `y`, because `y` is still available and comes
  earlier in Bob's order.

So `y` can never belong to Alice while `x` belongs to Bob.

This condition is also sufficient. To realize such a subset `S`, process
Alice's chosen items in increasing order. Before Alice buys the next chosen
item, let Bob keep buying his favorite remaining items until that chosen item
becomes Alice's first remaining item. The condition guarantees Bob will not be
forced to take a future item from `S`.

Now the problem becomes: choose a maximum-weight subset satisfying the condition
above.

Define:

```text
dp[i][j]
```

as the maximum value using only the first `i` objects in Alice's normalized
order, where `j` is the maximum `pos[x]` among objects already assigned to Bob.

When considering object `i + 1`:

1. Alice buys it.

   This is allowed only if:

   ```text
   pos[i+1] > j
   ```

   because every object already assigned to Bob must have smaller Bob-position
   than this Alice object. The value increases by `v[i+1]`, and `j` does not
   change.

2. Bob buys it.

   This is always allowed. The value does not change, and:

   ```text
   j = max(j, pos[i+1])
   ```

The answer is the maximum value over all final states.

To make this efficient, store the DP values by `j` in a lazy segment tree.

- Alice-buy transition: for states `j < pos[i+1]`, add `v[i+1]`.
- Bob-buy transition: states with `j < pos[i+1]` may move into state
  `pos[i+1]`; this needs a range maximum query.
- States with `j >= pos[i+1]` stay at the same `j`.

All operations are range add, range max, and point/range update operations on
the segment tree.

The time complexity is `O(n log n)` per test case.
