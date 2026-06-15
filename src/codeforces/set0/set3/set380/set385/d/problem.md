# D. Bear and Floodlight

[Problem link](https://codeforces.com/problemset/problem/385/D)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

One day a bear lived on the Oxy axis. He was afraid of the dark, so he could not
move at night through plane points that are not lit.

The bear wanted a night walk from his house at point `(l, 0)` to his friend's
house at point `(r, 0)` along the segment of length `r - l`. To make this walk,
every point of the segment must be lit.

The Oxy axis contains `n` floodlights. Floodlight `i` is at point `(x_i, y_i)`
and can illuminate any angular sector of size `a_i` degrees with vertex at
`(x_i, y_i)`. The bear asked his friend to orient the floodlights so that during
the walk along the segment the bear can get as far away from his house as
possible. Help the bear find this distance.

Assume the plane has no obstacles and no light sources other than the
floodlights. The friend sets the floodlight directions once before the walk; they
do not move during the walk.

## Input

The first line contains three space-separated integers `n`, `l`, and `r`
(`1 <= n <= 20`, `-10^5 <= l <= r <= 10^5`).

Each of the next `n` lines contains three space-separated integers `x_i`, `y_i`,
and `a_i` (`-1000 <= x_i <= 1000`, `1 <= y_i <= 1000`, `1 <= a_i <= 90`) —
the floodlight descriptions.

Two floodlights may be at the same point.

## Output

Print one real number — the maximum possible distance from the bear's house that
he can reach during the walk. The answer is accepted if its absolute or relative
error does not exceed `10^-6`.

## Example

### Input

```text
2 3 5
3 1 45
5 1 45
1 0 1
1 1 30
1 0 1
1 1 45
1 0 2
0 2 90
```

### Output

```text
2.000000000
0.732050808
1.000000000
2.000000000
```

## Note

The statement figures illustrate possible floodlight orientations for the samples.


### ideas
1. dp[mask] 表示由mask表示的集合,最远(从l处出发)能到达的地方
2. 那么新的light i, 应该从这个远端开始,看能覆盖的位置