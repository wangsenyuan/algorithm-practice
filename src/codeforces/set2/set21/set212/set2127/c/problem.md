# C. Trip Shopping

[Problem link](https://codeforces.com/problemset/problem/2127/C)

**Contest:** [Atto Round 1 (Codeforces Round 1041, Div. 1 + Div. 2)](https://codeforces.com/contest/2127)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Ali and Bahamin decided to spend their summer vacation on the beautiful southern coasts of Iran. They also agreed to do some shopping during the trip — but instead of setting a fixed budget, they decided to determine how much they would spend by playing a game.

The game is played on two arrays `a` and `b`, each containing `n` integers.

The game will last for `k` rounds. In one round:

- First, Ali selects two indices `i` and `j` (`1 <= i < j <= n`);
- Then, Bahamin rearranges the four integers `a_i`, `a_j`, `b_i`, and `b_j` arbitrarily. Note that Bahamin can swap numbers between two arrays. He can also keep the two arrays unchanged.

After all the `k` rounds, the value of the game is defined as `v = sum |a_i - b_i|` over `1 <= i <= n`. Ali and Bahamin will spend exactly `v` coins during their trip.

However, their goals are quite different:

- Ali wants to spend as little as possible, that is, to minimize `v`;
- Bahamin wants to spend as much as possible, that is, to maximize `v`.

You have to find the final amount of coins they will spend if both Ali and Bahamin play optimally.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `k` (`2 <= n <= 2 * 10^5`, `1 <= k <= n`) — the length of `a` and `b`, and the number of rounds;
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^9`) — the elements of `a`;
- one line with `n` integers `b_1, b_2, ..., b_n` (`1 <= b_i <= 10^9`) — the elements of `b`.

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one integer — the final amount of coins they will spend if both Ali and Bahamin play optimally.

## Example

### Input

```text
5
2 1
1 7
3 5
3 2
1 5 3
6 2 4
5 4
1 16 10 10 16
3 2 2 15 15
4 1
23 1 18 4
19 2 10 3
10 10
4 3 2 100 4 1 2 4 5 5
1 200 4 5 6 1 10 2 3 4
```

### Output

```text
8
9
30
16
312
```

### Note

In the first test case, Ali can only choose `(i, j) = (1, 2)`, and Bahamin can rearrange all four numbers. Thus, he can assign `a = [5, 1]` and `b = [3, 7]`. And the value of the game will be `v = |5 - 3| + |1 - 7| = 8`. It can be shown that this is the maximum possible value reachable for Bahamin — other arrangements like `a = [5, 7]`, `b = [1, 3]` are also possible, but they don't have larger values.

In the second test case, the best strategy for Bahamin is to keep the two arrays unchanged, regardless of what indices Ali selects. And the value of the game will be `v = |1 - 6| + |5 - 2| + |3 - 4| = 9`.

## Solution

Bahamin can always refuse to change a chosen pair, so Ali cannot decrease
the initial value `sum |a_i - b_i|`. Extra rounds after the first cannot
help Bahamin either: the four-number rearrange of a pair is idempotent, so
it is enough to consider a single round. `k` is unused.

Treat each index as a closed interval `[min(a_i, b_i), max(a_i, b_i)]`. The
initial value is the sum of those lengths. Sorting the four numbers of two
intervals yields two useful pictures:

- if the intervals overlap, every rearrange keeps the same total length;
- if they are disjoint, Bahamin can nest or interleave them and increase
  the total by twice the gap between them.

Ali therefore looks for a pair that Bahamin cannot improve. After sorting
intervals by left endpoint, any overlap appears between neighbors. If such
a neighbor pair exists, Ali always offers it and the answer is the initial
sum. Otherwise every pair is disjoint, Bahamin will take the smallest gap,
and the answer is the initial sum plus twice the minimum adjacent gap.

### Correctness sketch

The interval encoding is invariant under swapping `a_i` with `b_i`. Overlap
of any two intervals implies overlap of some consecutive pair in left-endpoint
order, so the linear scan detects Case 1 exactly. In the disjoint case the
closest pair is also adjacent after that sort, and the increase is
`2 · (left_{i} - right_{i-1})` as in the four-number calculation
`2a_j - 2b_i`. Ali cannot force a smaller increase, and Bahamin cannot get a
larger one than that minimum, so the value is optimal.

### Complexity

Sorting dominates: `O(n log n)` time and `O(n)` extra memory per test.
The sum of `n` over tests is at most `2 · 10^5`.


