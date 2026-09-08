# C. Stamina and Tasks

[Problem link](https://codeforces.com/problemset/problem/2208/C)

**Contest:** [Codeforces Round 1086 (Div. 2)](https://codeforces.com/contest/2208)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

There are `n` tasks for you. Task `i` has an integer value of `c_i` and a difficulty of `p_i`. Also, you have an initial stamina of `1`, which is denoted as `S`. You should process the tasks from task `1` to task `n`. For each task, you have two choices:

- Give up the task. This way, nothing will happen.
- Complete the task. This way, you will gain `S * c_i` points. However, `S` will drop to `S * (1 - p_i / 100)` after the task is completed.

You need to maximize your points after you finish the process.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^3`). The description of the test cases follows.

The first line of each test case contains an integer `n` (`1 <= n <= 10^5`) denoting the number of tasks.

The following `n` lines contain two integers each, denoting `c_i` (`1 <= c_i <= 100`) and `p_i` (`0 <= p_i <= 100`).

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, output a single real number — the maximum possible points you can get. Your answer is considered correct if its absolute or relative error does not exceed `10^{-6}`.

Formally, let your answer be `a`, and the jury's answer be `b`. Your answer is accepted if and only if `|a - b| / max(1, |b|) <= 10^{-6}`.

## Example

### Input

```text
2
2
10 0
20 5
3
10 5
10 80
20 5
```

### Output

```text
30.0000000000
29.0000000000
```

### Note

In the first test case, it's optimal to complete task `1` and `2` in order, gaining points of `10 + 20 = 30`.

In the second test case, it's optimal to complete task `1`, give up task `2`, and complete task `3`. Before completing task `3`, your stamina has dropped to `1 - 5/100 = 0.95`. So your gain is `10 + 20 * 0.95 = 29` points in total.

## Solution

Stamina scales every later reward by the same factor, so the optimal
choices on a suffix do not depend on the stamina you arrive with. Let
`f` be the maximum score on the remaining suffix assuming `S = 1`. Scan
from the last task to the first. For task `i` the two options are:

- skip it and keep `f`;
- take it and score `c_i` now, then scale the old suffix by
  `(1 - p_i / 100)`.

So `f` becomes `max(f, c_i + f * (1 - p_i / 100))`. After the scan, `f`
is the answer for the whole array at stamina `1`.

### Correctness sketch

A suffix score at stamina `S` is exactly `S` times the same suffix at
stamina `1`, because every completed task multiplies the current stamina
and adds a linear reward. Therefore the skip/take comparison on a suffix
is independent of `S`, and the right-to-left recurrence enumerates both
choices against an already-optimal suffix.

### Complexity

One reverse pass: `O(n)` time and `O(1)` extra memory. The sum of `n`
over tests is at most `10^5`.
