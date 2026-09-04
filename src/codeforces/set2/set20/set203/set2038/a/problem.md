# A. Bonus Project

[Problem link](https://codeforces.com/problemset/problem/2038/A)

**Contest:** [2024-2025 ICPC, NERC, Southern and Volga Russian Regional Contest (Unrated, Online Mirror, ICPC Rules, Preferably Teams)](https://codeforces.com/contest/2038)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

There is a team of `n` software engineers numbered from `1` to `n`. Their boss promises to give them a bonus if they complete an additional project. The project requires `k` units of work in total. The bonus promised to the `i`-th engineer is `a_i` burles. The boss doesn't assign specific tasks to engineers; it is expected that every engineer will voluntarily complete some integer amount of work units. The bonus will be paid to the entire team only if the project is completed; in other words, if the total amount of voluntary work units on the project is greater than or equal to `k`.

The amount of work that can be performed by each engineer is not limited. However, all engineers value their labour. The `i`-th engineer estimates one unit of their work as `b_i` burles. If the bonus is paid, the benefit `s_i` of the `i`-th engineer for completing `c` units of work is defined as `s_i = a_i - c · b_i`. If the bonus is not paid, the engineer will not volunteer to do any work.

Every engineer voices `c_i` in a way to maximize their own benefit `s_i`. If the expected benefit is going to be zero, an engineer will still agree to work to get the experience and to help their colleagues obtain the bonus. However, if the benefit is expected to be negative for some reason (an engineer needs to perform an excessive amount of work or the project is not going to be completed), that engineer will not work at all (completes zero amount of work units).

Print the amount of work completed by each engineer given that every engineer behaves optimally. The answer is unique.

## Input

The first line contains two integers `n` and `k` (`1 <= n <= 1000`; `1 <= k <= 10^6`) — the number of engineers and the number of work units the project requires.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 10^9`) — the bonus paid to each engineer if the project is completed.

The third line contains `n` integers `b_1, b_2, …, b_n` (`1 <= b_i <= 1000`) — the work unit cost for each engineer.

## Output

Print `n` integers `c_1, c_2, …, c_n` (`0 <= c_i <= k`) — the amount of work completed by each engineer.

## Examples

### Input

```text
3 6
4 7 6
1 2 3
```

### Output

```text
1 3 2
```

### Input

```text
3 12
4 7 6
1 2 3
```

### Output

```text
0 0 0
```

### Input

```text
3 11
6 7 8
1 2 3
```

### Output

```text
6 3 2
```

### Note

In the first example, engineers distributed the work across them and got the bonus, even though the benefit for the third engineer is zero.

In the second example, the bonus project requires too many work units to complete, so it's more beneficial for engineers not to work at all.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
