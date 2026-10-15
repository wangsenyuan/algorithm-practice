# P2448 无尽的生命

[Problem link](https://www.luogu.com.cn/problem/P2448)

time limit per test: 1.00s

memory limit per test: 125.00MB

The sequence starts as `S_i = i`: day `i` has ability `i`. Perform `k` swaps. The `i`-th swap exchanges the current abilities of day `x_i` and day `y_i`.

An abnormal pair is a pair of days `x < y` whose final abilities satisfy `S_x > S_y`. Count the abnormal pairs.

## Input

The first line contains an integer `k`, the number of swaps.

Each of the next `k` lines contains two integers `x_i` and `y_i`, the days whose abilities are swapped.

## Output

Print one integer, the number of abnormal pairs.

## Sample

### Input

```text
2
4 2
1 4
```

### Output

```text
4
```

The sequence changes as follows:

- `1, 2, 3, 4, 5, 6, …`
- `1, 4, 3, 2, 5, 6, …`
- `2, 4, 3, 1, 5, 6, …`

The abnormal pairs are `(1, 4)`, `(2, 3)`, `(2, 4)`, and `(3, 4)`.

## Constraints

- for 30% of the tests, `x_i, y_i <= 2 * 10^3`
- for 70% of the tests, `x_i, y_i <= 10^5`
- for 100% of the tests, `x_i, y_i <= 2^31 - 1` and `k <= 10^5`

## Solution

A day that never appears in a swap keeps ability `S_i = i`. Every swapped value is one of these touched days, so the final values on the touched days are a permutation of those days.

Sort the distinct touched days into `seen`. Slot `i` is day `seen[i]`, and `arr[i]` is the rank of the value currently there. `arr` starts as the identity. A swap of days `x` and `y` exchanges `arr` at the two fixed slots; the day-to-slot map does not move.

Two kinds of inversions remain.

- Both days are touched. Scan `arr` from left to right with a Fenwick tree of ranks. For the value now at a slot, add how many earlier touched values are strictly larger.
- One day is touched and the other is not. If rank `r` finishes in slot `i` and `r != i`, every untouched day strictly between `seen[r]` and `seen[i]` forms exactly one inversion with that value. There are `(seen[hi] - seen[lo] - 1) - (hi - lo - 1)` such days, where `lo = min(i, r)` and `hi = max(i, r)`. The subtracted term is the other touched days in that open interval; those pairs are already counted by the Fenwick tree.

A value that moved left is larger than each untouched day it passed, so it inverts with that day on its right. A value that moved right is smaller than each untouched day it passed, so that day inverts with it from the left.

The answer fits in `int64`. Sorting and the tree are `O(k log k)`.
