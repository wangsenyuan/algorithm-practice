# D. Lucky Segments (Codeforces 121D)

**Limits:** 4 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/121/D](https://codeforces.com/problemset/problem/121/D)

---

**Lucky numbers** are positive integers whose decimal representation contains only the digits **4** and **7**. For example, `47`, `744`, and `4` are lucky; `5`, `17`, and `467` are not.

Petya has `n` **closed segments** `[l_1; r_1], [l_2; r_2], …, [l_n; r_n]`.

In **one move**, he may pick any segment `i` and replace it with either:

- `[l_i + 1; r_i + 1]`, or  
- `[l_i - 1; r_i - 1]`

(i.e. shift that segment left or right by **1**).

A number `x` is called **full** if it lies in **every** segment: for all `i` (`1 ≤ i ≤ n`), `l_i ≤ x ≤ r_i`.

Petya makes **at most `k` moves**. After that, he counts how many **full lucky numbers** exist. Find the **maximum** possible count.

## Input

The first line contains two integers `n` and `k` (`1 ≤ n ≤ 10^5`, `1 ≤ k ≤ 10^18`) — the number of segments and the move budget.

Each of the next `n` lines contains two integers `l_i` and `r_i` (`1 ≤ l_i ≤ r_i ≤ 10^18`).

*(Statement note: in C++, avoid `%lld` for 64-bit I/O; `%I64d` is mentioned on the original page.)*

## Output

Print a single integer — the answer.

## Examples

### Example 1

**Input**

```text
4 7
1 4
6 9
4 7
3 5
```

**Output**

```text
1
```

### Example 2

**Input**

```text
2 7
40 45
47 74
```

**Output**

```text
2
```

## Note

In the first sample, shifting the second segment two units left gives `[4; 7]`, after which `4` is full.

In the second sample, shift the first segment two units right to `[42; 47]` and the second segment three units left to `[44; 71]`; then `44` and `47` are full.


## Solution

After all moves, every segment is still a segment with the same length as before.  If the final common intersection contains some interval `[L, R]`, then every original segment `[l_i, r_i]` must be shifted so that it covers `[L, R]`.

For one segment:

- if `l_i > L`, its left endpoint must move left to at most `L`, costing at least `l_i - L`;
- if `r_i < R`, its right endpoint must move right to at least `R`, costing at least `R - r_i`;
- if `l_i <= L` and `r_i >= R`, it already covers `[L, R]`, costing `0`.

So the minimum cost to make all segments cover `[L, R]` is:

```text
sum(max(0, l_i - L)) + sum(max(0, R - r_i))
```

This formula is exact. Moving a segment left helps only with the left boundary, and moving it right helps only with the right boundary. If a segment is too far right (`l_i > L`), moving it left by `l_i - L` is enough for the left boundary; because we only test intervals whose length is no larger than the minimum segment length, this shifted segment can also cover the right boundary. The symmetric argument applies when `r_i < R`.

Therefore an interval `[L, R]` is feasible iff:

1. `R - L + 1 <= minLen`, where `minLen = min(r_i - l_i + 1)`;
2. the cost above is at most `k`.

### Which Intervals Need Checking?

We only care about lucky numbers inside the final intersection. Suppose we want at least `w` full lucky numbers. Then it is enough to choose `w` consecutive lucky numbers:

```text
lucky[i], lucky[i+1], ..., lucky[i+w-1]
```

and ask whether all of them can be covered. The tightest interval that contains them is:

```text
L = lucky[i]
R = lucky[i+w-1]
```

If this tight interval is feasible, then those `w` lucky numbers are full. If it is not feasible, any larger interval containing the same lucky numbers is no easier to cover, because it only makes the boundary requirements stricter. So checking all consecutive lucky-number windows is sufficient.

The answer is monotonic: if `w` lucky numbers can be made full, then any smaller number can also be made full. Thus we binary search `w`.

### Lucky Number Range

The lucky numbers must be generated beyond the original maximum right endpoint.

Let:

```text
maxR = max(r_i)
```

No final full number can be greater than `maxR + k`, because even the segment with the largest original right endpoint would need more than `k` right moves to reach such a number. On the other hand, lucky numbers greater than `maxR` can be useful, because the whole intersection may be shifted right using the move budget.

So the correct generation limit is:

```text
maxR + k
```

not just `maxR`.

### Fast Cost Calculation

For a fixed left boundary `L`, the first cost part is:

```text
sum(max(0, l_i - L))
= sum(l_i - L for all l_i >= L)
= sum(l_i for all l_i >= L) - count(l_i >= L) * L
```

Sort segments by `l_i`. For every lucky number `L`, precompute:

```text
dp[L].sum = sum of l_i where l_i >= L
dp[L].cnt = count of such l_i
```

Then:

```text
leftCost = dp[L].sum - dp[L].cnt * L
```

Similarly, for a fixed right boundary `R`:

```text
sum(max(0, R - r_i))
= sum(R - r_i for all r_i <= R)
= count(r_i <= R) * R - sum(r_i for all r_i <= R)
```

Sort segments by `r_i`. For every lucky number `R`, precompute:

```text
fp[R].sum = sum of r_i where r_i <= R
fp[R].cnt = count of such r_i
```

Then:

```text
rightCost = fp[R].cnt * R - fp[R].sum
```

For a lucky window `[lucky[i], lucky[j]]`, the total cost is:

```text
leftCost + rightCost
```

If the interval length is at most `minLen` and this cost is at most `k`, then this window is feasible.

### Overflow Detail

Although `k`, `l_i`, and `r_i` fit in 64-bit integers, sums such as:

```text
sum(l_i)
sum(r_i)
```

can be as large as `10^5 * 10^18 = 10^23`, which does not fit in signed 64-bit integers. The implementation stores these prefix/suffix sums in a small unsigned 128-bit helper built from two `uint64` values. We only need addition, subtraction, multiplication of two ordinary `int` values, and comparison with `k`.

### Complexity

The number of lucky numbers up to about `2 * 10^18` is below `2^20`. Let this count be `m`.

```text
Generate lucky numbers: O(m)
Sort segments:          O(n log n)
Build dp/fp arrays:     O(n + m)
Binary search answer:   O(m log m)
Memory:                 O(m)
```
