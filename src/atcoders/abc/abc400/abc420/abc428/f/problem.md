# F - Pyramid Alignment

[Problem link](https://atcoder.jp/contests/abc428/tasks/abc428_f)

**Contest:** [AtCoder Beginner Contest 428](https://atcoder.jp/contests/abc428)

time limit: 2 sec

memory limit: 1024 MiB

score: 525 points

There are `N` intervals on a number line, numbered `1` to `N`.

Initially, interval `i` is `[0, W_i]` with `W_1 < W_2 < ... < W_N`.

Process `Q` queries in order. Each query is one of:

- **Type 1** (`1 v`): Let `l` be the current **left endpoint** of interval `v`. For every interval
  numbered `v` or less, translate it so its **left endpoint** becomes `l`.
- **Type 2** (`2 v`): Let `r` be the current **right endpoint** of interval `v`. For every interval
  numbered `v` or less, translate it so its **right endpoint** becomes `r`.
- **Type 3** (`3 x`): Output how many intervals currently contain coordinate `x + 1/2`.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= Q <= 2 * 10^5`
- `1 <= W_i <= 10^9`
- `W_1 < W_2 < ... < W_N`
- For type 1 and 2 queries, `1 <= v <= N`
- For type 3 queries, `0 <= x <= 10^9`
- At least one type 3 query is given
- All input values are integers

## Input

```text
N
W_1 ... W_N
Q
query_1
query_2
...
query_Q
```

Each query is one of:

```text
1 v
2 v
3 x
```

## Output

Let `q` be the number of type 3 queries. Print `q` lines; the `j`-th line is the answer to the
`j`-th type 3 query.

## Sample Input 1

```text
4
2 4 6 10
5
2 3
1 2
3 2
2 4
3 1
```

## Sample Output 1

```text
4
1
```

Initially: `[0,2], [0,4], [0,6], [0,10]`.

After query 1 (`2 3`): align right ends of intervals 1–3 to `6` → `[4,6], [2,6], [0,6], [0,10]`.

After query 2 (`1 2`): align left ends of intervals 1–2 to `2` → `[2,4], [2,6], [0,6], [0,10]`.

Query 3 (`3 2`): point `2.5` lies in all four intervals → output `4`.

After query 4 (`2 4`): align right ends of intervals 1–4 to `10` → `[8,10], [6,10], [4,10], [0,10]`.

Query 5 (`3 1`): point `1.5` lies only in interval 4 → output `1`.

## Solution

The intervals always keep the same lengths `W_i`. An operation only chooses an anchor endpoint and
moves a prefix of intervals so that each interval in that prefix is aligned to that anchor.

The useful invariant is the nesting order:

```text
L_1 >= L_2 >= ... >= L_N
R_1 <= R_2 <= ... <= R_N
```

Initially all left endpoints are `0`, and the right endpoints are increasing because `W_i` is
strictly increasing. A type `1 v` operation aligns intervals `1..v` to the left endpoint of interval
`v`, and a type `2 v` operation aligns them to the right endpoint of interval `v`; neither operation
breaks the nesting with larger intervals. Therefore, for a fixed query point `x + 1/2`, the intervals
containing it form a suffix of indices. We can binary search for the first containing interval.

### Which operation determines interval `i`?

For interval `i`, only operations with `v >= i` can affect it. Among those operations, the latest
relevant one determines whether interval `i` is placed by a left alignment or a right alignment.

Represent such an operation as a state:

```text
(v, dir, l, r)
```

where `v` is the operation index parameter, `dir` is `1` for left alignment or `2` for right
alignment, and `[l, r]` is the current position of interval `v` at the time of that operation.

If this state is the latest relevant state for interval `i`, then interval `i` is reconstructed as:

```text
dir = 1: [l, l + W_i]
dir = 2: [r - W_i, r]
```

If no operation with parameter at least `i` has happened, interval `i` is still `[0, W_i]`.

### Maintaining active states

When a new operation with parameter `v` arrives, operations with smaller parameter are no longer
needed as active outer frames:

```text
while stack top has parameter < v:
    pop it
```

The remaining stack top, if any, is the state that determines the current position of interval `v`.
So we first reconstruct interval `v` from that state, then store a new state for the current
operation. If the stack is empty, interval `v` is still `[0, W_v]`.

The stack is monotone by decreasing `v`, so each operation is pushed once and popped once.

To answer "latest relevant operation among parameters `>= i`", use a segment tree indexed by `v`.
At index `v`, store the id of the latest state whose operation parameter is `v`. A range maximum
query on `[i, N)` returns the latest state id affecting interval `i`, because state ids increase
with time.

### Querying a point

For type `3 x`, the actual point is `x + 1/2`. Since all endpoints are integers, an interval
`[l, r]` contains that point iff:

```text
l <= x and x < r
```

The right endpoint check is strict because `x = r` means the point is `r + 1/2`, already outside the
interval.

Define `check(i)` as "interval `i` contains `x + 1/2`":

1. Query the segment tree on `[i, N)` to find the latest state affecting interval `i`.
2. If no state exists, test the initial interval `[0, W_i]`.
3. Otherwise reconstruct interval `i` from that state and test `l <= x && x < r`.

By the nesting invariant, once some interval `i` contains the point, every larger interval also
contains it. Therefore `check(i)` is monotone, and binary search gives the first containing index
`p`. The answer is then:

```text
N - p
```

### Correctness

The nesting invariant ensures that the set of containing intervals is always a suffix. This makes
the binary search valid.

For any fixed interval `i`, operations with parameter `< i` never touch it. Among operations with
parameter `>= i`, later operations override earlier positions for interval `i`. The segment tree
range maximum query over `[i, N)` returns exactly this latest relevant operation state. Reconstructing
interval `i` from that state is correct because the operation aligns either the left endpoint or the
right endpoint while preserving length `W_i`.

Thus `check(i)` correctly decides whether interval `i` contains the query point, and binary search
counts exactly all containing intervals.

### Complexity

Each update operation performs one segment tree update and amortized `O(1)` stack work. Each type `3`
query performs a binary search over `N`, and each check uses one segment tree query, so it costs
`O(log^2 N)`. The total complexity is `O(Q log^2 N)`, with `O(Q + N)` memory.
