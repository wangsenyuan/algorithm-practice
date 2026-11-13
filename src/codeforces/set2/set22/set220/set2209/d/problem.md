# D. Ghostfires

Source: <https://codeforces.com/problemset/problem/2209/D>

Construct a longest string from at most the supplied red, green, and blue
ghostfires. Equal colors may not be adjacent or exactly three positions apart.

## Greedy construction

Keep the remaining counts of `R`, `G`, and `B`, then build the answer from
left to right. A color is unavailable when it equals the previous character or
the character three positions back. Among available colors, take the one with
the largest remaining count. On equal counts, prefer the character two
positions back.

The tie-break matters. From a suffix `ABA`, both `B` and the third color `C`
are legal. Choosing `B` on a tie preserves `ABAB`; choosing `C` can leave `B`
blocked by the distance-three restriction.

## Correctness

Let `T = r + g + b`, let `M` be the largest input count, and let
`O = T - M`. Since equal adjacent colors are forbidden, after the first use of
the majority color, every further use needs a non-majority separator. Hence no
valid answer can be longer than `min(T, 2*O + 1)`.

The last three characters are either `ABA` or `ABC`:

- After `ABA`, the legal choices are `B` and `C`; the greedy takes the larger
  remaining count, choosing `B` on a tie.
- After `ABC`, only `B` is legal, and the new suffix is `BCB`, again of the
  first form.

Thus the greedy always consumes the largest available supply and, when it
stops, either all ghostfires are used or only one color remains. In the latter
case it has used all `O` non-majority ghostfires and `O + 1` of the remaining
color, for length `2*O + 1`. It therefore reaches the upper bound and is
optimal.

## Complexity

The construction takes `O(r + g + b)` time and `O(r + g + b)` space for the
output.


## ideas
1. 考虑相同颜色的距离, 能否超过4? R????R, 那么中间使用2中颜色构造4个位置, BGBG好像是可以的
2. RBGBR 
