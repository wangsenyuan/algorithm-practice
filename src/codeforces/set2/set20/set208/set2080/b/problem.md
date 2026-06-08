# B. Best Runner

[Problem link](https://codeforces.com/problemset/problem/2080/B)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

There are `n` running tracks in the stadium with lengths
`a_1, a_2, ..., a_n`. There are also `m` runners, and the `i`-th runner starts
at the beginning of track `b_i`.

All runners train for `T` seconds. The training of a runner proceeds as follows:

- Suppose the runner is currently at the beginning of track `i`.
- They run to the end of the current track in `a_i` seconds.
- After finishing that track, they can instantly:
  - return to the beginning of the current track,
  - move to the beginning of track `i - 1`, if `i > 1`,
  - move to the beginning of track `i + 1`, if `i < n`.
- They continue running from the chosen track.
- Once the training duration reaches `T` seconds, they finish training.

The best runner is a runner who runs the maximum number of full tracks during
the training time. There may be several best runners.

Determine how many full tracks the best runner can run.

## Input

The first line contains three integers `n`, `m`, and `T`
(`1 <= m <= n <= 300000`, `1 <= T <= 10^9`) — the number of tracks, the number
of runners, and the training duration.

The second line contains `n` integers `a_1, a_2, ..., a_n`
(`1 <= a_i <= 10^9`) — the lengths of the tracks.

The third line contains `m` integers `b_1, b_2, ..., b_m`
(`1 <= b_1 < b_2 < ... < b_m <= n`) — the track numbers from which the runners
start.

## Output

Output a single integer — the maximum number of full tracks that one of the
runners can run during the training time.

## Examples

### Input 1

```text
5 3 10
4 5 2 7 1
1 2 4
```

### Output 1

```text
4
```

### Input 2

```text
4 2 11
4 5 7 10
2 3
```

### Output 2

```text
2
```

## Note

In the first example, the runner starting on track `4` can run the most tracks:
they should run track `4`, then move to track `5` and run it `3` times.

In the second example, the runner starting on track `2` can run the second track
`2` times.

## Scoring

The tests consist of six groups. Points for each group are awarded only if all
tests of the group and all required groups are passed.

- Group 0: examples.
- Group 1: `n <= 1000`.
- Group 2: `a_i <= a_{i+1}` for all `1 <= i < n`.
- Group 3: `T <= 20`.
- Group 4: `a_i <= 20`.
- Group 5: `m = n`.
- Group 6: full constraints.


### ideas
1. 如果一个人已经在*最短的*trak（圈）上，那么他就在这个圈上完成比赛;
2. 且如果一个人选择往前(i-1)移动，那么他在到达*最优*的圈前，不会改变方向；
3. 假设对于i来说，j (j <= i)是他的*最优*圈
4. k * a[j] + sum[i] - sum[j] <= T (sum[i] - sum[j]是他到达j前花费的时间)
5. 总的圈数=k + i - j
6. k * a[j] - sum[j] <= T - sum[i]
7. 目标是使的k-j最大（不仅仅是k最大）
8. 我们固定j，看哪个人移动到这里能够跑最多。
9. 那么显然是最近的那个i（前后需要分开算)
10. 

## Solution summary

Think about the final behavior of one runner. If the runner decides that track
`j` is the best track to repeat, then before reaching `j` they should move
monotonically toward `j`; after reaching `j`, they should stay on `j` and keep
running it.

So we can enumerate the track `j` that will be repeated at the end.

Let `sum[i]` be the prefix sum of track lengths:

```text
sum[i] = a[0] + a[1] + ... + a[i-1]
```

### Runner from the left

Suppose the nearest starting runner on the left side of `j` starts at `b`.

To reach track `j`, this runner must first run tracks:

```text
b, b+1, ..., j-1
```

The time spent before reaching `j` is:

```text
sum[j] - sum[b]
```

The number of full tracks already completed is:

```text
j - b
```

After that, the runner stays on track `j`. If the remaining time is:

```text
T - (sum[j] - sum[b])
```

then the number of additional full runs on track `j` is:

```text
(T - (sum[j] - sum[b])) / a[j]
```

So the total number of full tracks is:

```text
j - b + (T - (sum[j] - sum[b])) / a[j]
```

This is exactly what the code computes as:

```go
k := calc(sum[j]-sum[b[lo]], j) + j - b[lo]
```

where `calc(t, j)` returns how many times track `j` can still be run after
spending `t` seconds to reach it.

### Runner from the right

If the nearest starting runner on the right side starts at `b`, then to reach
`j` they first run:

```text
b, b-1, ..., j+1
```

The time spent is:

```text
sum[b+1] - sum[j+1]
```

The number of full tracks already completed is:

```text
b - j
```

Then they repeat track `j`, giving:

```text
b - j + (T - (sum[b+1] - sum[j+1])) / a[j]
```

The code computes this with:

```go
k := calc(sum[b[lo+1]+1]-sum[j+1], j) + b[lo+1] - j
```

### Why only nearby starts are checked

For a fixed target track `j`, the solution checks:

- the closest starting position `<= j`,
- the closest starting position `> j`.

These two candidates are enough for this approach because any runner farther
away must spend more movement before reaching `j`. The code therefore scans
tracks from left to right and maintains `lo`, the index of the closest starting
runner not greater than the current `j`.

Then:

- if `b[lo] <= j`, check the left candidate;
- if `lo + 1 < m`, check the right candidate;
- if no left candidate exists yet, only the right candidate is checked.

### Complexity

The array `b` is already sorted by the statement. During the scan, pointer `lo`
only moves forward, so every track and every starting point is processed a
constant number of times.

Time complexity:

```text
O(n)
```

Space complexity:

```text
O(n)
```
