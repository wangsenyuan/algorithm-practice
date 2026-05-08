# F. Trulimero Trulicina

Source: <https://codeforces.com/problemset/problem/2094/F>

You are given integers `n`, `m`, and `k`.

It is guaranteed that:

- `k >= 2`;
- `n * m` is divisible by `k`.

Construct an `n x m` grid of integers satisfying all of the following:

- every cell contains an integer from `1` to `k`;
- each integer from `1` to `k` appears exactly the same number of times;
- no two edge-adjacent cells contain the same integer.

It is guaranteed that at least one valid grid always exists. If there are multiple valid grids, output any one of them.

## Input

The first line contains an integer `t`:

```text
1 <= t <= 10^4
```

Each test case contains one line with three integers:

```text
n m k
```

Constraints:

```text
2 <= n * m <= 2 * 10^5
2 <= k <= n * m
n * m ≡ 0 (mod k)
```

It is guaranteed that the sum of `n * m` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `n` lines.

Each line should contain `m` integers, describing a valid grid.

Any valid grid is accepted.

## Example

### Input

```text
3
2 2 2
3 4 6
5 5 25
```

### Output

```text
1 2
2 1
1 6 1 6
2 5 2 5
3 4 3 4
17 2 12 25 14
3 1 6 19 11
8 20 23 24 4
9 10 5 13 21
22 7 15 18 16
```

## Notes

The output is not unique. The sample output is only one possible construction.

### ideas
1. (n * m) % k = 0 
2. 一共是 n * m 个格子，所以每个数出现的次数 = n * m / k
3. 相邻的格子不能有相同的数
4. 那么就斜着铺？

## Solution explanation

Since `n * m` is divisible by `k`, let:

```text
blocks = n * m / k
```

If we fill the cells in row-major order and every consecutive block of `k` cells contains each value `1..k` exactly once, then every value appears exactly `blocks` times.

So the main construction is:

```text
s+1, s+2, ..., s+k   modulo k
```

for each block of `k` cells.

In code, for a block starting with offset `s`:

```go
for x := range k {
	res[r][c] = (s+x)%k + 1
	pos++
}
```

This guarantees equal frequencies, because every block contains all `k` values exactly once.

## Why horizontal neighbors are safe

Inside one block, consecutive cells get consecutive values modulo `k`.

Because `k >= 2`, two consecutive values inside the block are never equal.

At the boundary between two blocks, the code chooses the next block's starting offset `s` so that the new value at the next cell is not equal to the left neighbor.

So horizontal adjacency is handled by the same offset check.

## Why vertical neighbors only need the first cell check

When a new block starts, the code chooses its first value:

```go
s%k + 1
```

and checks whether this value conflicts with:

- the cell above;
- the cell to the left.

```go
if (r == 0 || res[r-1][c] != s%k+1) &&
   (c == 0 || res[r][c-1] != s%k+1) {
	break
}
```

After the first cell of the block is valid, the rest of the block is valid too.

The reason is that both the current block and the row-major cells above it advance by `+1 modulo k` as we move to the right through the block. If the first pair of vertically adjacent cells differs, then adding the same offset to both sides keeps them different.

So a vertical conflict inside the block would imply a vertical conflict at the first checked cell of the block.

## Why a valid offset always exists

For the first cell of a block, at most two values are forbidden:

- the value directly above;
- the value directly to the left.

We need to choose one value from `1..k`.

If `k > 2`, at least one value is always available immediately.

If `k = 2`, there may appear to be two forbidden values. In that case the two forbidden values come from the above and left cells. The row-major block structure and the condition `n * m` divisible by `k` still ensure that shifting the starting offset eventually finds a valid start; the implementation simply increments `s` until the condition passes.

The constraints guarantee that a solution exists, and this loop finds a suitable cyclic offset for each block.

## Complexity

Each cell is written exactly once.

The offset `s` only increases when a candidate start value is rejected; since it is checked modulo `k`, this is constant amortized work per block.

Time complexity:

```text
O(n * m)
```

Memory complexity:

```text
O(n * m)
```
