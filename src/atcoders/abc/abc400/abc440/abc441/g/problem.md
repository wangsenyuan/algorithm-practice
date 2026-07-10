# G - Takoyaki and Flip

[Problem link](https://atcoder.jp/contests/abc441/tasks/abc441_g)

**Contest:** [AtCoder Beginner Contest 441](https://atcoder.jp/contests/abc441)

time limit: 2 sec

memory limit: 1024 MiB

score: 575 points

`N` plates are arranged in a straight line from left to right. Initially, all plates
are face-up and have no takoyaki on them.

Process `Q` queries of three types:

- Type 1: `1 L R X` — for each plate `i` in `[L, R]`, if it is face-up, place `X`
  takoyaki on it.
- Type 2: `2 L R` — for each plate `i` in `[L, R]`, if it has at least one takoyaki,
  eat all of them; then flip the plate.
- Type 3: `3 L R` — print the maximum number of takoyaki among plates `L..R`.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= Q <= 2 * 10^5`
- In all queries, `1 <= L <= R <= N`
- In type 1 queries, `1 <= X <= 10^9`
- There is at least one type 3 query
- All input values are integers

## Input

```text
N Q
query_1
...
query_Q
```

Type 1 queries are given as `1 L R X`. Type 2 and 3 queries are given as `t L R`.

## Output

Let `q` be the number of type 3 queries. Print `q` lines with the answers in order.

## Sample Input 1

```text
6 6
1 3 5 4
3 2 3
1 1 6 2
2 3 4
3 1 6
3 2 3
```

## Sample Output 1

```text
4
6
2
```

## Sample Input 2

```text
2 8
1 1 2 1000000000
1 1 2 1000000000
2 2 2
1 1 2 1000000000
1 1 2 1000000000
1 1 2 1000000000
3 2 2
3 1 2
```

## Sample Output 2

```text
0
5000000000
```

Note that face-down plates never receive takoyaki, and answers may exceed `2^32`.

## Sample Input 3

```text
24 30
1 11 24 4326
1 4 16 1149
1 14 20 2331
1 12 14 8930
1 22 23 6989
3 15 20
3 10 19
1 3 12 7988
1 18 23 8450
3 9 19
3 13 15
2 8 15
2 9 14
1 11 17 4062
1 6 15 1721
3 7 13
1 11 20 8541
1 8 10 3748
1 1 17 3252
2 9 23
2 1 23
3 2 22
1 5 23 7468
3 1 12
3 12 19
2 6 24
3 2 14
3 1 15
2 15 19
3 2 14
```

## Sample Output 3

```text
7806
16736
22393
16736
10858
0
7468
7468
0
0
0
```

## Solution

Use a lazy segment tree. For each segment, store:

- `mx`: the maximum number of takoyaki on any plate in the segment.
- `state`: the orientation state of the whole segment.

The code uses three states:

- `0`: every plate in the segment is face-up;
- `2`: every plate in the segment is face-down;
- `1`: the segment is mixed.

The merge rule is simple: if both children have the same `state`, the parent has that
state; otherwise the parent is mixed. The parent maximum is always the maximum of the two
children.

Two observations make the lazy operations possible:

1. A face-down plate always has `0` takoyaki. Type 1 never adds to face-down plates, and
   type 2 removes all takoyaki before flipping.
2. In a mixed segment, there is at least one face-up plate. Since face-down plates have
   value `0`, the segment maximum comes from face-up plates after any positive type-1 add.

For a type 1 query `1 L R X`, only face-up plates receive `X`.

- If a covered segment is entirely face-down (`state == 2`), nothing changes.
- Otherwise, the segment contains at least one face-up plate, and all candidate maximum
  values among face-up plates increase by `X`, so `mx += X`.

This is why `apply` adds `X` whenever `state != 2`.

For a type 2 query `2 L R`, every covered plate first loses all takoyaki, then flips.
Therefore, for a fully covered segment:

- `mx` becomes `0`;
- if the segment is uniform, its state flips between `0` and `2`;
- if the segment is mixed, it remains mixed.

The lazy tag stores:

- `add`: pending type-1 additions;
- `flipCnt`: pending type-2 operations.

When a flip is applied after previous additions, those additions no longer matter for the
covered segment because all takoyaki is eaten before flipping, so the flip tag replaces
older pending additions. When an addition is applied after pending flips, it can be merged
into the pending addition.

Type 3 is a standard range maximum query.

Each query touches `O(log N)` segment tree nodes, so the total time complexity is
`O(Q log N)`, and the memory complexity is `O(N)`.
