# E. Post Lamps

[Problem link](https://codeforces.com/problemset/problem/990/E)

**Contest:** [Codeforces Round #492 (Div. 1)](https://codeforces.com/contest/990)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Adilbek's house is on a street modeled as the OX axis. The street is dark, so he wants to install post
lamps to illuminate the segment `[0, n]`.

There are `n` positions where lamps may be placed, at integer coordinates `0, 1, ..., n-1`. Some
positions are blocked and cannot be used.

Lamps come in powers `1, 2, ..., k`. A lamp of power `l` placed at position `x` illuminates the
segment `[x, x + l]`. Adilbek must buy lamps of **exactly one** power. Each lamp of power `l` costs
`a_l`.

Find the minimum total cost to illuminate the entire segment `[0, n]`. Lamps may extend beyond `n`; only
full coverage of `[0, n]` is required.

## Input

The first line contains three integers `n`, `m`, and `k` (`1 <= k <= n <= 10^6`, `0 <= m <= n`).

The second line contains `m` integers `s_1, s_2, ..., s_m` (`0 <= s_1 < s_2 < ... < s_m < n`) — blocked
positions.

The third line contains `k` integers `a_1, a_2, ..., a_k` (`1 <= a_i <= 10^6`) — the cost of one lamp
for each power.

## Output

Print the minimum total cost using lamps of exactly one power. If covering `[0, n]` is impossible,
print `-1`.

## Example

### Input

```text
6 2 3
1 3
1 2 3
```

### Output

```text
6
```

### Input

```text
4 3 4
1 2 3
1 10 100 1000
```

### Output

```text
1000
```

### Input

```text
5 1 5
0
3 3 3 3 3
```

### Output

```text
-1
```

### Input

```text
7 4 3
2 4 5 6
3 14 15
```

### Output

```text
-1
```

## Solution

For one fixed lamp power `t`, the cost per lamp is fixed, so we only need to minimize the
number of lamps. A greedy choice works:

- assume the already covered prefix ends at position `i + t`, where the last lamp was placed
  at usable position `i`;
- if this already reaches `n`, we are done;
- otherwise, the next lamp must be placed at some unblocked position `j` with `i < j <= i + t`;
- choosing the largest such `j` is always optimal, because it covers at least as far as any
  smaller valid choice and never makes future choices worse.

The direct greedy simulation for every `t` would be too slow, so the implementation precomputes
two helper arrays:

- `before[x]`: the nearest unblocked position `<= x`;
- `next[x]`: the nearest blocked position `>= x`, or `n` if there is none.

For a fixed power `t`, start from position `0`. If position `0` is blocked, the answer is
impossible immediately, because every coverage must start at coordinate `0`.

At a current lamp position `i`, the farthest candidate for the next lamp is:

```text
j1 = before[i + t]
```

If `j1 == i`, then there is no unblocked position in `(i, i + t]`, so this power is impossible.

Otherwise, let:

```text
j2 = next[i]
```

If `j2 < j1`, then there is a blocked point before `j1`, but `j1` is still reachable and
unblocked, so the greedy move is simply `i = j1`.

If the next blocked point is not before `j1`, then the segment from `i` up to that blocked
point has no blocked positions. In that continuous free segment, the greedy positions are just:

```text
i, i + t, i + 2t, ...
```

So the implementation can count many lamps at once instead of moving one step at a time. It
computes the largest multiple step that still stays before the next blocked position, adds those
lamps to the count, and jumps directly there. If that multiple-step jump would stop before the
true greedy position `j1`, the code uses `j1` instead; this handles cases where `i + t` is blocked
but an earlier usable position is still reachable.

After counting the number of lamps needed for each power `t`, the total cost is:

```text
count(t) * a_t
```

The minimum valid total over all powers is the answer. If no power works, print `-1`.

## Correctness

For a fixed power `t`, consider any moment when the last chosen lamp is at `i` and the covered
prefix reaches `i + t`. Any valid next lamp must be placed at an unblocked position in
`(i, i + t]`; otherwise, a gap remains uncovered. Among all such positions, choosing the largest
one is never worse than choosing a smaller one, because it extends the covered prefix farthest and
does not invalidate any later position. Therefore the greedy rule gives the minimum number of lamps
for that power.

The arrays `before` and `next` do not change the greedy choices; they only find and skip over them
faster. `before[i+t]` gives exactly the farthest reachable unblocked next position, and the
multiple-step shortcut only compresses consecutive greedy moves inside a block-free interval.
Thus the counted number of lamps for each valid `t` is the greedy minimum for that power.

Since every solution must use exactly one power, checking all powers and taking the minimum
`count(t) * a_t` gives the global optimum.

## Complexity

Building `blocked`, `before`, and `next` takes `O(n)` time and `O(n)` memory.

For a fixed power `t`, each normal greedy move advances the current lamp position by about `t`,
and the shortcut counts a whole arithmetic progression inside a block-free interval in `O(1)`.
So the number of loop iterations for that power is `O(n / t)`. Summed over all powers:

```text
n/1 + n/2 + n/3 + ... + n/k = O(n log k)
```

Therefore the total time complexity is `O(n log k)`, and the memory complexity is `O(n)`.
