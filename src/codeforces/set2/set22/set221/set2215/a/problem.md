# A. Interval Mod

[Problem link](https://codeforces.com/problemset/problem/2215/A)

**Contest:** [Codeforces Round 1092 (Unrated, Div. 1, Based on THUPC 2026 — Finals)](https://codeforces.com/contest/2215)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an array `a` consisting of `n` integers, as well as a parameter `k` and an integer set `M = {p, q}`.

You can perform the following operation on `a` an arbitrary number of times (possibly zero):

- First, choose an interval `[l, r]` (`1 ≤ l ≤ r ≤ n`) of length at least `k` (i.e., `r − l + 1 ≥ k`) and an integer `m ∈ M`;
- Then, set `a_i ← a_i mod m` for each `l ≤ i ≤ r`.

You have to find the minimum possible value of `∑_{i=1}^{n} a_i` after all operations.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

The first line of each test case contains four integers `n`, `k`, `p`, and `q` (`1 ≤ k ≤ n ≤ 10^5`, `1 ≤ p < q ≤ 10^9`) — the length of `a`, the parameter, and the elements of `M`.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 ≤ a_i ≤ 10^9`) — the elements of `a`.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, output a single integer — the minimum possible value of `∑_{i=1}^{n} a_i` after all operations.

## Example

### Input

```text
6
1 1 3 4
2026
3 2 10 20
31 41 59
4 3 3 4
1 2 3 4
6 4 9 20
18 27 180 9 45 99
7 4 3 5
6 7 14 12 100 78 4
9 4 244 353
9982 4435 3998 2443 5399 8244 3539 9824 4353
```

### Output

```text
1
11
3
0
4
569
```

### Note

In the second test case, a possible way to obtain `∑ a_i = 11` is to apply the following operation to `a`:

- Choose `[l, r] = [1, 3]` and `m = 10`, then `a` becomes `[1, 1, 9]`.

In the third test case, a possible way to obtain `∑ a_i = 3` is to apply the following operations to `a`:

- Choose `[l, r] = [1, 4]` and `m = 4`, then `a` becomes `[1, 2, 3, 0]`;
- Choose `[l, r] = [2, 4]` and `m = 3`, then `a` becomes `[1, 2, 0, 0]`.

## ideas

1. `a % m <= a` for `a >= 0`, so every op weakly decreases the sum; `p` is terminal because `a % p < p < q`.
2. Per index the only useful finals are `a % p` and `(a % q) % p`. The second can be better or worse; they coincide when `p | q`.
3. Always `p` the whole array at the end. Who gets `q` first is the only choice.
4. A `q`-run can be shorter than `k` if a `p`-block of length `k` is applied first and used as padding (sample 5).
5. Conversely, a first op paints `k` consecutive cells the same type, so some monochromatic run of length `k` is necessary.

## Summary

Each index ends as either `a_i % p` or `(a_i % q) % p`, according to whether its first successful mod is `p` or `q`. Later `p` never hurts, so it is enough to apply `p` on `[1, n]` last. The remaining question is which indices receive `q` while still fresh.

The first operation must cover `k` consecutive cells of one type. From that seed, a window of `k - 1` already-colored cells plus one neighbor can paint the next cell either type. Therefore any assignment is realizable as soon as it contains one monochromatic run of length `k`, and that condition is also necessary.

So force some window of length `k` to be all-`p` or all-`q`, and assign every other index independently to `min(a_i % p, (a_i % q) % p)`. Enumerate the window in `O(n)` with prefix sums: unconstrained prefix, uniform `min(p-sum, q-then-p-sum)` on the window, unconstrained suffix. Also take the two global assignments (all `p`, all `q` then `p`); both are dominated by some seed but are harmless.

If `p` divides `q`, the two residues are equal and the answer is just `∑ (a_i % p)`.

Complexity: `O(n)` per test case.
