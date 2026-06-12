# B. Zoo

[Problem link](https://codeforces.com/problemset/problem/183/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

The Zoo is represented by an infinite two-dimensional grid. There are `n`
binoculars standing along the `OX` axis: for every `i` from `1` to `n` there is
exactly one binocular located at the point `(i, 0)`. There are also `m`
flamingos in the Zoo, each located at a point with strictly positive
coordinates; the flamingos do not move.

Each binocular can be independently rotated to face **any** straight ray. Once a
binocular is fixed to a direction, it sees every flamingo that lies on that ray.
For each binocular its angle is chosen so as to maximize the number of flamingos
it can see.

Compute the sum, over all `n` binoculars, of the maximum number of flamingos
each binocular can see.

## Input

The first line contains two integers `n` and `m`
(`1 <= n <= 10^6`, `1 <= m <= 250`) — the number of binoculars and the number of
flamingos.

Each of the next `m` lines contains two integers `x` and `y`
(`1 <= x, y <= 10^9`) — the coordinates of a flamingo. No two flamingos share the
same location.

## Output

Print a single integer — the sum over all binoculars of the maximum number of
flamingos that can be seen.

## Example

### Input

```text
5 5
2 1
4 1
3 2
4 3
4 4
```

### Output

```text
11
```

## Note

In the example the per-binocular maxima are:

- binocular `1`: flamingos `(2,1),(3,2),(4,3)` lie on the line `y = x - 1` → `3`
- binocular `2`: flamingos `(3,2),(4,4)` → `2`
- binocular `3`: no two flamingos are collinear with it → `1`
- binocular `4`: flamingos `(4,1),(4,3),(4,4)` lie on the vertical line `x = 4` → `3`
- binocular `5`: flamingos `(4,1),(3,2)` → `2`

The total is `3 + 2 + 1 + 3 + 2 = 11`.

## Solution summary

A binocular at `(b, 0)` sees a set of flamingos simultaneously iff they are all
collinear with `(b, 0)`. So for each binocular the answer is the largest group of
flamingos that lie on one line through it. Since every single flamingo is always
visible on its own, each binocular sees at least `1` (when `m >= 1`).

Key observation: a line can only matter for some binocular if it passes through
an **integer** point `(b, 0)` with `1 <= b <= n`. Two flamingos determine a line,
and there are only `O(m^2)` such pairs (`m <= 250`).

Algorithm (see [solution.go](solution.go)):

1. For every pair of flamingos `(i, j)` compute the `x`-intercept of the line
   through them (`play`):
   - same `y` (horizontal) → never meets `y = 0`, skip;
   - same `x` (vertical) → intercept is that `x`;
   - otherwise solve for `x` at `y = 0`; keep it only if it is an integer inside
     `[1, n]`.
2. Bucket each valid pair by `(intercept x, normalized direction)`. The direction
   is reduced by `gcd` and sign-normalized so collinear pairs map to the same key.
   `freq[x][dir]` then equals the number of pairs on that line, i.e. `C(k, 2)`
   where `k` is the number of flamingos on it.
3. For each binocular `x`, take the maximum bucket value `t = C(k, 2)` and recover
   `k` via a precomputed table `best[C(k,2)] = k`. Add `k` (or `1` if no pair maps
   to this binocular) to the answer.

Complexity: `O(m^2 * log + n)` time. Note the answer can be large, so a 64-bit
accumulator is appropriate.
