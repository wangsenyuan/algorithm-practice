# D - Count Subgrid Sum = K (ABC461)

**Contest:** [ABC461](https://atcoder.jp/contests/abc461) — AtCoder Beginner Contest 461  
**Task:** [https://atcoder.jp/contests/abc461/tasks/abc461_d](https://atcoder.jp/contests/abc461/tasks/abc461_d)

**Time limit:** 4 sec / **Memory limit:** 1024 MiB  
**Score:** 425 points

## Problem Statement

There is an `H × W` grid; each cell contains `0` or `1`. Row `i` is given as string
`S_i` of length `W` (`S_i[j]` is the value in row `i`, column `j`, 1-indexed).

Count rectangular subgrids whose cell sum equals `K`. Equivalently, count integer
quadruples `(r_1, c_1, r_2, c_2)` such that:

- `1 <= r_1 <= r_2 <= H`
- `1 <= c_1 <= c_2 <= W`
- the sum of values in rows `r_1..r_2` and columns `c_1..c_2` equals `K`

## Constraints

- `1 <= H, W <= 500`
- `0 <= K <= H * W`
- each `S_i` is a length-`W` string of `0` and `1`

## Input

```text
H W K
S_1
S_2
⋮
S_H
```

## Output

Print the answer.

## Sample Input 1

```text
3 4 3
1001
1101
0110
```

## Sample Output 1

```text
8
```

Eight rectangles have sum `3` (listed in the official statement).

## Sample Input 2

```text
5 4 20
0101
1010
0101
1010
0101
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
15 20 17
10111101101100000100
01100000000010000011
01110010111000111000
11001100000111011000
10100001100011100010
01101000101010000101
10110001111110000100
10110011101100101101
01010001110011001001
01010110010001100110
01110100011110011110
01100000100111010010
11001101100111101100
10111100010101111011
00101101011100010000
```

## Sample Output 3

```text
448
```

## Solution

### Key idea

Fix the top and bottom rows `r_1..r_2`. For each column `c`, let `col[c]` be the
number of `1`s in rows `r_1..r_2` of that column. A rectangle with those row bounds
and columns `c_1..c_2` has sum `K` iff the subarray `col[c_1..c_2]` sums to `K`.

So for each row pair, the problem reduces to counting 1D subarrays with sum `K`.

### Row-pair enumeration + prefix sums

Enumerate all `r_1 <= r_2`, extending `col` as `r_2` increases. While scanning
columns left to right, maintain:

- `sum` — prefix sum of `col[0..c]`
- `todo` — all earlier prefix sums (`todo[0] = 0`, then append `sum` after each column)

A subarray ending at column `c` with sum `K` corresponds to a start index `l` with
`sum - todo[l] = K`, i.e. `todo[l] = sum - K`.

Because prefix sums are non-decreasing, use two pointers on `todo`:

- `pos1` — first index with `todo[pos1] >= sum - K`
- `pos2` — first index with `todo[pos2] > sum - K`

Then `pos2 - pos1` counts starts `l` with `todo[l] = sum - K`, i.e. valid subarrays
ending at `c`. Add this to the answer for each column.

### Rotation trick

If `H > W`, transpose the grid so the smaller side is the row count `n`. Row-pair
enumeration costs `O(n^2 · W)`; keeping `n = min(H, W)` minimizes that cost when the
grid is very rectangular.

### Complexity

- Time: `O(min(H, W)^2 · max(H, W))` — at most about `1.25 × 10^8` operations for
  `500 × 500`
- Space: `O(W)` for `col` and the prefix-sum list
