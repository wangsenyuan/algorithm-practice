# A. Adrenaline Rush

https://codeforces.com/problemset/problem/2052/A

**Time Limit:** 3 seconds

**Memory Limit:** 1024 megabytes

## Problem

There are `n` racing cars. Before the race starts, the cars stand in order:

```text
1, 2, ..., n
```

Car `1` is closest to the starting line, car `2` is second, and so on.

During the race, one car may overtake the car directly in front of it. This
swaps their positions. For any two cars `x` and `y`, car `x` may overtake car
`y` at most once, and car `y` may overtake car `x` at most once.

At the end of the race, the final order of cars is:

```text
c_1, c_2, ..., c_n
```

where `c_1` is the winner.

Given the final order, construct the longest possible sequence of overtake
events that could have happened.


## Solution

Think about adding cars one by one. Suppose we already constructed the longest
valid overtake sequence for cars `1..i-1`. Now add car `i`.

Initially, car `i` is behind every car `0..i-1` (using zero-based labels in the
code). To create as many overtakes as possible, first make car `i` overtake all
previous cars, from nearest to farthest:

```text
i overtakes current car at position i-1
i overtakes current car at position i-2
...
i overtakes current car at position 0
```

Now car `i` is at the front. This gives `i` new overtakes, one with every
previous car, and none of these ordered pairs has been used before.

But the final permutation may require some cars to finish before car `i`. Let
`pos1[x]` be the final position of car `x` in `c`. Scan the current prefix after
car `i` moved to the front. Whenever a car `x` should be before `i` in the final
order (`pos1[x] < pos1[i]`), let `x` overtake `i`.

This second overtake is also allowed: for this pair, the first event was
`i` overtaking `x`, and now the opposite direction `x` overtakes `i` happens
once.

The implementation does this iteratively:

```text
for i = 0..n-1:
    move car i to the front by overtaking all previous cars
    let every car that must finish before i overtake i back
```

The code maintains:

- `a`: current order of cars.
- `pos`: current position of each car.
- `pos1`: final position of each car in the target permutation.

Every printed operation swaps adjacent cars, so the sequence is valid. Each
ordered pair of cars is used at most once, because a pair can only be used when
the larger-numbered car is added: first larger over smaller, and possibly later
smaller over larger.

The construction is maximal because for each unordered pair of cars, there can
be at most two overtakes total, one in each direction. The algorithm uses both
directions exactly when needed to still end in the target final order, and uses
one direction otherwise.

The number of operations is `O(n^2)`, and constructing them directly is fine for
`n <= 1000`.
