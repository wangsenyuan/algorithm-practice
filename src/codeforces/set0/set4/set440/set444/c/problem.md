# C. DZY Loves Colors

[Problem link](https://codeforces.com/problemset/problem/444/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

DZY loves colors, and he enjoys painting.

On a colorful day, DZY gets a colorful ribbon, which consists of `n` units (they
are numbered from `1` to `n` from left to right). The color of the `i`-th unit
of the ribbon is `i` at first. It is colorful enough, but we still consider that
the colorfulness of each unit is `0` at first.

DZY loves painting, we know. He takes up a paintbrush with color `x` and uses it
to draw a line on the ribbon. In such a case some contiguous units are painted.
Imagine that the color of unit `i` currently is `y`. When it is painted by this
paintbrush, the color of the unit becomes `x`, and the colorfulness of the unit
increases by `|x - y|`.

DZY wants to perform `m` operations, each operation can be one of the following:

1. Paint all the units with numbers between `l` and `r` (both inclusive) with
   color `x`.
2. Ask the sum of colorfulness of the units between `l` and `r` (both inclusive).

Can you help DZY?

## Solution

Use a segment tree.

For every tree node representing interval `[l, r]`, maintain three values:

```go
mark[i]
add[i]
sum[i]
```

Their meaning:

- `mark[i] > 0`: every cell in this segment currently has exactly this color.
- `mark[i] == 0`: this segment is not uniform, so its children contain the
  detailed color information.
- `mark[i] == -1`: this segment's uniform color was cleared after being
  repainted; it should be treated as non-uniform for future repaint recursion.
- `add[i]`: colorfulness added uniformly to every cell in this segment by
  repainting this whole segment at once.
- `sum[i]`: total colorfulness contribution of this segment, including children
  and this node's own lazy `add`.

Initially cell `i` has color `i`, and colorfulness is `0`, so every leaf starts
as a uniform segment:

```go
mark[leaf] = i
```

### Repainting a Uniform Segment

Suppose a whole segment `[l, r]` is currently uniform with color `old`.

If we repaint the entire segment to color `x`, every cell in the segment gains:

```text
abs(old - x)
```

So:

```go
add[i] += abs(old - x)
sum[i] += abs(old-x) * (r-l+1)
mark[i] = x
```

This is the fast case.

### Repainting a Non-Uniform Segment

If the segment is not uniform, we cannot compute the repaint gain from one
color. We must recurse into its children.

Before recursing, if this node currently has a uniform mark, push that mark to
both children:

```go
if mark[i] > 0 {
    mark[left] = mark[i]
    mark[right] = mark[i]
    mark[i] = 0
}
```

After repainting the affected children, recompute:

```go
sum[i] = sum[left] + sum[right] + add[i] * (r-l+1)
```

The term `add[i] * (r-l+1)` is necessary because `add[i]` is a per-cell lazy
colorfulness contribution stored at this node.

### The `clear` Function

The function `clear(i, l, r, x)` means:

```text
make the whole segment [l, r] ready to be assigned color x,
while adding the required colorfulness increase first
```

If the segment is uniform, it applies the fast formula.

If the segment is not uniform, it clears both children recursively and rebuilds
`sum[i]`.

Finally it sets:

```go
mark[i] = -1
```

Then the caller sets:

```go
mark[i] = x
```

because after repainting the whole queried segment, it is uniform with color
`x`.

### Query

For a fully covered node, return:

```go
sum[i]
```

For a partially covered node, recurse into the children. But the queried cells
also inherit this node's lazy per-cell contribution `add[i]`.

If the current recursive query asks for subrange `[L, R]`, the contribution from
this node's lazy value is:

```go
add[i] * (R - L + 1)
```

This detail is easy to translate incorrectly. The C++ reference has:

```cpp
delta[k] * (r - l + 1)
```

where `l, r` are the current queried subrange. In Go, the matching expression
is:

```go
res := tr.add[i] * (R - L + 1)
```

It must be multiplication by the queried length. Adding `add[i]` only once, or
adding the length separately, gives wrong answers.

### Complexity

Every repaint either handles a uniform segment at once or descends to children.
The segment tree keeps enough uniform-color information to avoid visiting every
cell repeatedly.

Each operation runs in `O(log^2 N)` amortized in the intended analysis, and is
fast enough for `N, M <= 10^5`.
