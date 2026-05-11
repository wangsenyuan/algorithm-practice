# Codeforces 2218G - The 67th Iteration of "Counting is Fun"

https://codeforces.com/problemset/problem/2218/G

## Statement

You are given two integers `n` and `m`, and an array `b` of length `n`
(`0 <= bi < m`). It is guaranteed that every integer from `0` to `m - 1` appears
in `b` at least once.

A group of `n` people are standing in a line, in fixed positions numbered
`1, 2, ..., n`. They all want to sit down, but some are too socially awkward to
do so immediately.

For some arbitrary array `a` of length `n` (`0 <= ai < n`), each person `i` has
a social awkwardness level `ai`. Consider the seating process, which evolves in
discrete units of time `0, 1, 2, ...` as follows:

- If `ai = 0`, person `i` sits down at time `0`.
- Otherwise (`ai > 0`), person `i` sits down at the beginning of a time unit `t`
  if both conditions are satisfied:
  1. At least `ai` people have already sat down strictly before time `t`.
  2. At least one neighbor of person `i` has already sat down strictly before
     time `t`. The neighbors are `i - 1` and `i + 1`, if they exist.

At the beginning of every unit of time, each standing person checks these
conditions. All people for whom the conditions are satisfied sit down
simultaneously, and then the process advances to the next unit of time.

For each `i`, the value `bi` denotes the time unit at which person `i` sits
down.

Determine how many distinct arrays `a` are consistent with the given array `b`.
Since this number can be very large, output it modulo `676767677`.

The modulo `676767677` is prime.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

For each test case:

- The first line contains two integers `n` and `m`
  (`1 <= m <= n <= 2 * 10^5`) — the number of people and the total duration of
  the process.
- The second line contains `n` integers `b1, b2, ..., bn`
  (`0 <= bi < m`), where `bi` is the time unit at which person `i` sits down.

It is guaranteed that:

- the sum of `n` over all test cases does not exceed `2 * 10^5`;
- each provided array `b` contains every integer from `0` to `m - 1`.

## Output

For each test case, output a single integer: the number of distinct arrays `a`
that result in the given array `b`, modulo `676767677`.

## Sample

```text
7
4 3
0 1 2 0
8 4
0 1 2 3 1 2 0 1
9 5
1 0 1 3 4 3 2 1 0
15 14
3 0 1 2 3 4 5 6 7 8 9 10 11 12 13
5 5
4 3 0 1 2
5 2
0 1 1 1 0
3 2
0 1 1
```

```text
2
0
1920
138007136
8
0
0
```

## Note

In the first test case, the only two valid arrays are:

```text
[0, 1, 3, 0]
[0, 2, 3, 0]
```

At time `0`, people `1` and `4` sit down. Then person `2` sits down at time `1`.
If `a2` were `0`, person `2` would have sat down at time `0`; if `a2` were
greater than `2`, person `2` would not have been able to sit down at time `1`,
since only `2` people had sat down before time `1`.

It can be shown that the only possible value for `a3` is `3`.

In the second test case, there are no valid arrays because it is impossible for
person `5` to sit down before either of their neighbors.


### ideas
1. 如果不存在b[?] = 0, 那么答案是0（没有人在时刻0坐下，那么后续就不会发生）
2. 排序后的b，不能出现b[j] + 1 < b[j+1]的情况（假设在时刻3有一些人坐下了，但是在时刻4没有人坐下，到时刻5有人坐下， 不能出现这种情况）
3. 假设i, 在时刻b[i]坐下，那么b[i-1] <= b[i] or b[i+1] <= b[i]
4. a[i] > 在时刻b[i] - 2时坐下的人，a[i] <=在时刻b[i] - 1时坐下的人

## Solution

For each position `i`, we can determine how many values of `a[i]` are valid from
the desired sitting time `b[i]`. The implementation computes this count and
multiplies it into the answer immediately; it does not need to store separate
`lo`/`hi` arrays.

Let:

```text
cnt[t] = number of people with b[i] <= t
```

So `cnt[t]` is the number of people who have already sat down by the end of time
`t`.

At the beginning of time `d`, the people who have already sat down strictly
before that time are exactly those with:

```text
b[i] < d
```

Their count is:

```text
cnt[d - 1]
```

## People sitting at time 0

A person sits at time `0` if and only if:

```text
a[i] = 0
```

Therefore, for every position with `b[i] = 0`, the only valid value is:

```text
a[i] = 0
```

This contributes exactly one choice, so the implementation does not need to
multiply anything for time `0`.

## Necessary neighbor condition

For a person `i` with `b[i] = d > 0`, at least one neighbor must have already sat
down strictly before time `d`.

That means at least one existing neighbor must satisfy:

```text
b[neighbor] < d
```

If both sides either do not exist or have `b[neighbor] >= d`, then person `i`
could not sit at time `d`, so the whole array `b` is impossible.

This is the check in the code:

```text
(i == 0 || b[i-1] >= d) && (i == n-1 || b[i+1] >= d)
```

If it is true, return `0`.

## Upper bound for a[i]

At time `d`, the threshold condition must be true:

```text
a[i] <= number of people seated before time d
```

So:

```text
a[i] <= cnt[d - 1]
```

Thus the upper bound is implicit:

```text
hi = cnt[d - 1]
```

## Lower bound for a[i]

The person must not sit earlier than time `d`.

For `d = 1`, the only earlier time is `0`. Since `b[i] > 0`, we only need:

```text
a[i] != 0
```

So:

```text
lo = 1
```

For `d > 1`, consider time `d - 1`. The person must fail at least one of the two
conditions at the beginning of that time.

At the beginning of time `d - 1`, the count condition is:

```text
a[i] <= cnt[d - 2]
```

The neighbor condition is true if at least one neighbor has:

```text
b[neighbor] < d - 1
```

There are two cases.

If some neighbor already sat before time `d - 1`, then the neighbor condition
was already true at time `d - 1`. To prevent person `i` from sitting too early,
the count condition must be false:

```text
a[i] > cnt[d - 2]
```

So:

```text
lo = cnt[d - 2] + 1
```

Otherwise, no neighbor sat before time `d - 1`. Then person `i` already fails
the neighbor condition at time `d - 1`, so the count condition does not need to
fail. Since `d > 0`, we only need `a[i] > 0`:

```text
lo = 1
```

This matches the implementation:

```text
lo = 1
if neighbor with b < d - 1 exists:
    lo = cnt[d - 2] + 1
```

## Independence

Once `b` is fixed, both conditions for each person depend only on:

- `a[i]`;
- the global counts `cnt[t]`;
- the sitting times of the two neighbors in `b`.

They do not depend on the chosen values of `a[j]` for other positions.

Therefore every position contributes independently:

```text
number of choices for i = cnt[d - 1] - lo + 1
```

The current code multiplies this value directly into `ans` while processing the
positions in `open[d]`.

After the neighbor feasibility check, this count is the number of valid choices
for `a[i]` at this position. The statement guarantees every time value
`0..m-1` appears, so when the lower bound uses `cnt[d - 2] + 1`, there is at
least one person at time `d - 1`, making `cnt[d - 1]` large enough for the
interval endpoint arithmetic to stay valid.

## Complexity

Group positions by their sitting time:

```text
open[d] = all positions i with b[i] = d
```

Then process `d = 0..m-1`, maintaining the prefix counts `cnt[d]`.

Every position is processed once, so the total complexity is:

```text
O(n + m)
```

with `O(n + m)` memory.
