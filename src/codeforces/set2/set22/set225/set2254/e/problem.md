# E. Chronostasis

[Problem link](https://codeforces.com/problemset/problem/2254/E)

**Contest:** [Codeforces Round 1114 (Div. 3)](https://codeforces.com/contest/2254)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Yousef has a hidden array `a` of length `n` consisting entirely of strictly positive integers.

An operation was performed exactly once to create an array `b`:

- set `b_1 = a_1`;
- for every `i` from `2` to `n`, set `b_i = a_i - a_{i-1}`.

After this, the elements of `b` were completely shuffled.

You are given the shuffled array `b`. Reconstruct the lexicographically smallest original array `a`. If it is impossible for any arrangement of `b` to produce an array `a` of strictly positive integers, output `-1`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with an integer `n` (`1 <= n <= 2 * 10^5`) — the size of the array;
- one line with `n` integers `b_1, b_2, ..., b_n` (`-10^9 <= b_i <= 10^9`) — the shuffled array `b`.

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print `n` strictly positive integers `a_1, a_2, ..., a_n` (`a_i >= 1`) — the lexicographically smallest original array `a`. If it is impossible, print `-1`.

## Example

### Input

```text
8
1
5
4
-5 2 1 1
6
-3 4 2 -1 1 0
6
-2 -2 4 1 0 1
7
0 0 -2 3 0 -1 2
8
-1 -1 -1 -1 5 0 0 1
5
1000000000 500000000 750000000 100000000 900000000
10
1000000000 -1000000000 500000000 -500000000 1 1 -1 -1 2 -2
```

### Output

```text
5
-1
1 1 3 2 6 3
1 1 2 6 4 2
2 1 1 1 1 4 2
1 1 1 6 5 4 3 2
100000000 600000000 1350000000 2250000000 3250000000
-1
```

### Note

In the first test case, the only valid array is `a = [5]`.

In the second test case, no arrangement of `b` reconstructs an array of strictly positive integers, so the answer is `-1`.

In the third test case, one valid arrangement reconstructs `a = [1, 1, 3, 2, 6, 3]`. The differences `[1, 0, 2, -1, 4, -3]` are a permutation of the given `b`, and this `a` is lexicographically smallest among all valid reconstructions.

## Solution

Any arrangement of `b` as a sequence `d_1, d_2, ..., d_n` reconstructs

```text
a_1 = d_1
a_k = d_1 + d_2 + ... + d_k    (k >= 2)
```

so `a` is exactly the sequence of prefix sums of `d`. The array `a` is valid if and only if every prefix sum is at least `1`. Conversely, the differences of a candidate `a` recover `d` uniquely: `d_1 = a_1` and `d_k = a_k - a_{k-1}`. Therefore it is enough to permute `b` so that every prefix sum is `>= 1`, and among all such permutations to take the one whose prefix-sum sequence is lexicographically smallest.

### Existence

A valid permutation exists if and only if `sum(b) >= 1`.

Necessity is immediate: the last prefix sum is `a_n = sum(b)`.

Sufficiency is proved by the greedy below. Always append the smallest remaining value `x` such that the new prefix sum stays `>= 1`. The first step needs some `x >= 1`; if `sum(b) >= 1` then not every element can be `<= 0`, so such an `x` exists.

Suppose the greedy later fails after a valid prefix summing to `s >= 1`, with `r >= 1` values left. Failure means every remaining `x` satisfies `s + x < 1`. All values are integers, so `x <= -s`. The remaining sum is then at most `-s * r`. But the remaining sum is also `sum(b) - s >= 1 - s`, hence

```text
1 - s  <=  -s * r
1      <=  s (1 - r)
```

If `r = 1` the right-hand side is `0`. If `r >= 2` then `1 - r < 0`, so `s(1 - r) <= 1 - r <= -1`. In both cases `1 <=` a non-positive number, a contradiction.

Thus the greedy never gets stuck when `sum(b) >= 1`, and it returns `-1` precisely when no valid `a` exists.

### Lexicographically smallest `a`

It remains to show that this greedy produces the lexicographically smallest valid `a`.

Let `G` be the greedy prefix-sum sequence and let `L` be any valid sequence. The differences of a prefix of `a` are uniquely determined by that prefix, so if `G` and `L` agree on the first `i-1` entries then they have used the same multiset of differences and have the same remainder.

Proceed by induction. For the first position, every valid `a` must start with some element of `b` that is `>= 1`, so the smallest such element is the unique optimal `a_1`. The greedy picks it.

Assume `G` and `L` agree on `a_1, ..., a_{i-1}`, with common value `s`. The next difference of `L` is some remaining `y` with `y >= 1 - s`. The greedy picks the smallest remaining `x*` with the same lower bound, so `x* <= y` and the next greedy value `s + x*` is at most `L`'s next value. Because `L` is lexicographically smallest it cannot be strictly larger at this index, hence the two next values coincide and the induction continues.

Combined with the existence argument, each greedy choice is the unique lex-optimal next value and is always completable. Therefore `G = L`.

### Implementation

Sort `b` and store it in a range-max segment tree. At prefix sum `s` (with `s = 0` before the first pick), query the leftmost position whose value is at least `1 - s`; that is the smallest feasible remaining difference. Write it into `a`, delete it by setting the leaf to `-inf`, and set `s` to the new prefix sum. If a query finds nothing, output `-1`.

Each update and query is `O(log n)`, so the whole test case is `O(n log n)`. The sum of `n` over tests is `2 * 10^5`.