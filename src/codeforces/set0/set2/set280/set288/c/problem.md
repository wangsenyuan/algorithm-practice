# Solution

To maximize the target expression, we want a permutation that minimizes lost bits.

A bit is considered **lost** for pair `(i, p[i])` if that bit is `1` in both numbers, so it becomes `0` in `i XOR p[i]`.

The key fact: for every `n`, we can build a permutation where no bits are lost in the processed block.

## Construction idea

Process while `n > 0`:

1. Let `b` be the highest set-bit position in `n` (0-indexed).
2. Let `m` be the smallest integer in `[0, n]` whose `b`-th bit is also `1`.
3. Pair numbers symmetrically inside segment `[m, n]`:
   - `p[m] = n`
   - `p[m+1] = n-1`
   - ...
   - `p[n] = m`

Equivalently:

- `p[m] = m - 1`
- `p[m-1] = m`
- `p[m+1] = m - 2`
- `p[m-2] = m + 1`
- and so on (same mirror pairing pattern in that block).

After finishing this segment, continue with the remaining prefix:

`n = m - (n - m + 1) - 1`

and repeat.

This iterative block pairing gives the required optimal permutation.