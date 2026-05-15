# E. Adjusting Drones (Codeforces 2157E)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2157/E](https://codeforces.com/problemset/problem/2157/E)

## Problem Statement

You manage `n` drones. Drone `i` currently has energy level `a[i]`. You are also given an integer `k`, meaning that after the balancing process finishes, no energy level may appear more than `k` times.

The drones repeatedly perform one balancing operation while some energy level occurs strictly more than `k` times:

- Mark every drone `i` whose current energy value already appeared at some earlier position `j < i`.
- Increase the energy of every marked drone by `1`.
- Remove all marks.

The relative order of drones never changes. Only their energy values change.

Find the number of balancing operations performed before the process stops.

## Input

Each test contains multiple test cases.

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

For each test case:

- The first line contains two integers `n` and `k` (`1 <= k <= n <= 2 * 10^5`) — the number of drones and the maximum allowed frequency of any energy level.
- The second line contains `n` integers `a[1], a[2], ..., a[n]` (`1 <= a[i] <= 2n`) — the initial energy levels.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one integer: the number of balancing operations performed.

## Example

### Input

```text
5
6 3
1 1 1 1 1 1
5 1
1 3 2 1 4
6 2
1 1 1 2 3 3
4 1
8 8 8 8
2 2
1 2
```

### Output

```text
3
4
4
3
0
```

## Note

In the first test case, the energy levels evolve as:

```text
[1, 1, 1, 1, 1, 1]
[1, 2, 2, 2, 2, 2]
[1, 2, 3, 3, 3, 3]
[1, 2, 3, 4, 4, 4]
```

After `3` operations, each energy level appears at most `3` times, so the process stops.

In the second test case:

```text
[1, 3, 2, 1, 4]
[1, 3, 2, 2, 4]
[1, 3, 2, 3, 4]
[1, 3, 2, 4, 4]
[1, 3, 2, 4, 5]
```

The process stops after `4` operations.


## Solution Explanation

The operation only depends on the multiset of current energy values, not on the
actual identities of the drones.

For one value `x`, suppose there are `c` drones with energy `x` at some moment.
In the next operation, exactly one of them can stay at `x`, and the other
`c - 1` drones move to `x + 1`. Therefore a pile of equal values behaves like
this:

- one drone is left behind at the current value;
- all extra drones keep moving to the right;
- whenever they reach an empty value, one of them can stop there;
- the remaining extra drones continue moving.

So after at most `m` operations, the extra drones that started from value `x`
can only use values in the range `(x, x + m]`. They try the values from left to
right. If there is an empty value, one drone settles there. If all reachable
values are already occupied, all remaining extra drones are still together at
`x + m`.

This gives a check function:

> `play(m)`: after forcing `m` operations, is every final frequency at most
> `k`?

If yes, then the real process stops in at most `m` operations. If no, then more
than `m` operations are necessary.

### Why We Process Values From Large To Small

Consider an extra drone starting from value `x`. During the first step it can
only reach `x + 1`, during the second step it can only reach `x + 2`, and so on.

This means drones starting from larger values always have priority over drones
starting from smaller values for the same destination:

- a drone from `x` can reach `x + d` after `d` operations;
- a drone from `x - 1` reaches the same value `x + d` only after `d + 1`
  operations.

So when deciding which values are already occupied for a starting value `x`, all
values to the right of `x` that are caused by larger starting values are already
determined. Smaller starting values cannot arrive earlier and take those places.

That is why the implementation scans values in decreasing order.

### The Greedy Check

The code keeps an array `freq`, where `freq[v]` is the number of drones currently
assigned to value `v` in the simulated result after `m` operations.

During the descending scan, `zeros` stores empty values to the right of the
current value. Because the scan goes from large to small, `last(zeros)` is the
nearest empty value greater than the current `i`.

For each value `i`:

1. If `freq[i] == 0`, then `i` is empty, so it is added to `zeros`.
2. While `freq[i] > 1` and the nearest empty value `z` satisfies
   `z <= i + m`, move one extra drone from `i` to `z`.
3. If `freq[i]` is still greater than `1`, then the remaining extra drones
   cannot find an empty value within `m` steps, so all of them are placed at
   `i + m`.

This matches the real movement:

- an extra drone always stops at the first empty value it can reach;
- if no empty value appears within `m` steps, it is still moving after the
  `m`-th operation, so it stays at exactly `i + m`;
- one drone must remain at `i`, because the first occurrence of a value is never
  marked by that value itself.

After this simulation, if any `freq[v] > k`, then after `m` operations there is
still an overloaded energy level, so `m` is not enough. Otherwise `m` is enough.

### Why Binary Search Works

The predicate "`m` operations are enough" is monotonic.

If after `m` operations every frequency is at most `k`, then the real process
has already stopped by time `m`. Allowing more time cannot make the answer
larger, because the process would not continue after it stops.

Equivalently, while forcing extra operations, the maximum frequency never
increases:

- each old value contributes at most one drone that stays there;
- the other drones only move one step to the right;
- a new pile is formed from the one survivor at this value plus the moving
  extras from the previous value, so it cannot exceed the previous maximum pile
  size.

Therefore, once `play(m)` is true, `play(m + 1)`, `play(m + 2)`, and so on are
also true. The implementation binary-searches the smallest true `m`.

### Bounds

Initially `a[i] <= 2n`. The answer is at most `n`: from any starting value, among
the next `n` values there must be some free position for every extra drone,
because there are only `n` drones total. So the largest value we need to touch is
at most `3n`, and the implementation's `freq` array of size `4n + 1` is enough.

Each check scans `O(n)` possible values, and the binary search uses a fixed
number of bits. The total complexity is `O(n log n)` per test case in the usual
bound notation, with `O(n)` memory.
