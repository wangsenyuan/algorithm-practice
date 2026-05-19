# E. The Robotic Rush (Codeforces 2185E)

**Limits:** 3 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2185/E](https://codeforces.com/problemset/problem/2185/E)

---

There is an infinitely long number line.

On the line there are `n` **robots** and `m` **spikes**, each at a fixed integer position. Robot `i` is at `a_i`, spike `i` is at `b_i`. A robot **dies** if it ever occupies a position that contains a spike.

You are given `k` instructions. Each instruction tells **every** robot to move **one unit left** (`L`) or **one unit right** (`R`). Instructions are processed in order.

For each `i` from `1` to `k`, output how many robots are still **alive** after the first `i` instructions.

## Input

The first line contains an integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

For each test case:

- One line: `n`, `m`, `k` (`1 ≤ n, m, k ≤ 2·10^5`) — robots, spikes, instructions.
- One line: `n` distinct integers `a_1, …, a_n` (`0 ≤ a_i ≤ 10^9`) — robot positions.
- One line: `m` distinct integers `b_1, …, b_m` (`0 ≤ b_i ≤ 10^9`) — spike positions.
- One line: a string of length `k` over `{L, R}` — the instruction sequence.

### Prefix displacement bounds

Sort spike positions `b`.

Walk the instruction string once. Let `move` be the cumulative offset after each step (`L` → `-1`, `R` → `+1`). For each prefix length `i` (`1 … k`), store in `dp[i]`:

- `min_l` — minimum cumulative offset seen so far;
- `max_r` — maximum cumulative offset seen so far.

If a robot starts at `pos`, after the first `i` moves its position is always between `pos + dp[i].min_l` and `pos + dp[i].max_r` (at the step endpoints; the code uses these extremes to test spike contact).

### When does a robot die? (`check(pos, i)`)

Binary-search `spikes` to find the nearest spikes around `pos`:

- `l` — largest spike strictly to the left of `pos` (or `-∞` if none);
- `r` — smallest spike at or to the right of `pos` (or `+∞` if none).

With `dl = dp[i].min_l`, `dr = dp[i].max_r`, the robot **dies within the first `i` instructions** iff it reaches a dangerous extent:

```text
pos + dl <= l   OR   pos + dr >= r
```

(i.e. it moves to or past the neighboring spike on the left, or reaches the neighboring spike on the right.)

`check(pos, i) == true` means **dead by step `i`** (used that way in the code).

### Earliest death + difference array

- `diff[0] = n` (all alive before any instruction).
- For each robot at `pos`, if `check(pos, k)` (dies by the end), binary-search the **smallest** `j` with `check(pos, j)` and do `diff[j-1]--` (one fewer alive from step `j` onward).
- Prefix-sum `diff` to get alive counts after instructions `1 … k`.

### Complexity

- Sort spikes: `O(m log m)`.
- Build `dp`: `O(k)`.
- Per robot: `O(log m + log k)` for searches → `O(n log(n + m + k))` per test case.
- Total `O((n + m + k) log)` with sums of `n, m, k` bounded by `2·10^5`.

### Ideas (original notes)

1. Each robot can be handled alone: find the instruction step at which it dies.