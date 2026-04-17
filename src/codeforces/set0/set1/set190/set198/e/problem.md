# Problem

One day Qwerty the Ranger witnessed two transport ships collide. As a result, all cargo was scattered in space. Qwerty wants to collect as many lost items as possible and sell them later.

Among the cargo are gravitational grippers. A gripper can be installed on a spaceship and can pull items in space to the ship's cargo hold.

There are `n` lost grippers. For the `i`-th gripper:

- position: `(x_i, y_i)`
- mass: `m_i`
- power: `p_i`
- radius: `r_i`

A gripper can pull any item if:

- item mass `<=` gripper power, and
- distance to the item `<=` gripper radius.

Qwerty's ship is at point `(x, y)` and initially has one old gripper with power `p` and radius `r`.

Qwerty may:

- collect grippers that can be pulled by the currently active gripper,
- install any collected gripper as the new active gripper at any time,
- reuse any available gripper any number of times.

At any moment only one gripper is active. The ship and all items are stationary, except that a pulled item moves into the cargo hold.

Find the maximum number of lost grippers Qwerty can collect.

## Input

The first line contains five integers:

`x y p r n`

Constraints:

- `-10^9 <= x, y <= 10^9`
- `1 <= p, r <= 10^9`
- `1 <= n <= 250000`

Each of the next `n` lines contains:

`x_i y_i m_i p_i r_i`

Constraints:

- `-10^9 <= x_i, y_i <= 10^9`
- `1 <= m_i, p_i, r_i <= 10^9`

It is guaranteed that:

- all grippers are at distinct points,
- no gripper is at Qwerty's ship position.

## Output

Print one integer — the maximum number of grippers Qwerty can pull to his ship.

Do **not** count the initial old gripper.

## Example

### Input

```text
0 0 5 10 5
5 4 7 11 5
-7 1 4 7 8
0 2 13 5 6
2 -3 9 3 4
13 5 1 9 9
```

### Output

```text
3
```

## Note

In the sample:

1. collect the second gripper,
2. use it to collect the first gripper,
3. use the first gripper to collect the fourth gripper.

The third gripper is too heavy, and the fifth gripper is too far.

## Solution Summary (matching `solution.go`)

The implementation treats each lost gripper as reachable if there exists an already available active gripper `(power, radius)` that satisfies:

- `m_i <= power`
- `dist((x,y), (x_i,y_i)) <= radius`

where distance is Euclidean distance from the ship position (implemented with squared distance to avoid floating point).

### Key idea

Use a BFS-like process over available grippers:

1. Start queue with the initial gripper `(p, r)`.
2. Pop one capability pair `(power, radius)`.
3. Fetch all *currently uncollected* grippers that satisfy mass and distance constraints.
4. Each fetched gripper increases answer by `1` and contributes a new capability `(p_i, r_i)` into the queue.

So the core problem is efficient retrieval/deletion under two filters:

- `mass <= power`
- `distanceSquared <= radius^2`

### Data structure used

1. Compute `d_i = (x_i-x)^2 + (y_i-y)^2` for every gripper.
2. Coordinate-compress all distinct `d_i`.
3. Build a segment tree over compressed distance index.
4. At each leaf (one fixed distance), store grippers sorted by mass.
5. Each segment tree node maintains `minMass` of its interval.

This allows pruning:

- if node's `minMass > power`, no gripper in this node is usable now;
- if node interval is beyond allowed distance index, skip it.

When a gripper is collected, it is removed from the leaf list, and `minMass` is updated upward.  
Therefore each gripper is removed once, and never processed again.

### Complexity

Let `n` be number of lost grippers.

- sorting/compression/build: `O(n log n)`
- each gripper is inserted once and removed once
- segment-tree traversals are logarithmic with pruning

Overall complexity is `O(n log n)` (amortized), with `O(n)` memory for stored grippers plus segment-tree overhead.