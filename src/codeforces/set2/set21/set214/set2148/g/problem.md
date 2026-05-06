Bessie has found an array `a` of length `n` on the floor. There appears to be a handwritten note lying next to the array, seemingly written by Farmer John. The note reads:

Help me, dear Bessie! Let `f(a)` denote the maximum integer `k` in the range `[1, n)` such that:

`gcd(a1, a2, ..., ak) > gcd(a1, a2, ..., a(k+1))`,

or `0` if no such `k` exists.

Bessie decides to help FJ. She defines `g(a)` to represent the maximum value of `f(a)` over all possible reorderings of `a`.

Bessie decides to not only find `g(a)`, but also the value of `g(p)` for all prefixes `p` of `a`.

Output `n` integers, the `i`-th of which is:

`g([a1, a2, ..., ai])`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`).

The following line contains `n` space-separated integers `a1, a2, ..., an` (`1 <= ai <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `n` integers on a new line: the `i`-th of which should be:

`g([a1, a2, ..., ai])`.

## Example

### Input

```text
3
8
2 4 3 6 5 7 8 6
6
6 6 6 6 6 6
9
8 4 2 6 3 9 5 7 8
```

### Output

```text
0 1 2 3 3 3 4 5
0 0 0 0 0 0
0 1 2 2 4 4 4 4 5
```


### ideas
1. g(a) = max(f(a.)), a. 是a的某个排列
2. 如果gcd(a) = x, 那么让a[i]/= x, 那么 gcd(a) = 1
3. 然后是否可以得到g(a) = n - 1 ?
4. 如果所有的a相同， 那么g(a) = 0, 否则g(a) = max(freq) ？
5. 不是，考虑 a = [4, 2, 3], 那么显然g(a) = 2
6. 但是如果 a = [4, 2, 3, 5], g(a) = 2 
7. 如果a中2的倍数，3的倍数，5的倍数，。。。 的最多个个数
8. 那么加入一个数的时候，只需要把它质因数分解即可

### Correct criterion

For a multiset `S`, consider choosing some integer divisor `d >= 2`.

If we put all numbers divisible by `d` first, then their prefix gcd is still divisible by `d`. If the next number is not divisible by `d`, the gcd strictly drops at exactly

`k = count of numbers divisible by d`.

So any `d` with `0 < cnt[d] < |S|` can achieve `f = cnt[d]`.

Conversely, suppose a permutation has a gcd drop at position `k`:

`gcd(first k numbers) > gcd(first k+1 numbers)`.

Let `D = gcd(first k numbers)`. Since the gcd drops after adding the next element, the next element is not divisible by `D`. All first `k` numbers are divisible by `D`, so at least `k` numbers are divisible by some divisor `D >= 2`, and not all numbers are divisible by `D`.

Therefore:

`g(S) = max cnt[d]` over all `d >= 2` such that `cnt[d] < |S|`.

This is the important correction: it is not enough to count only prime factors. We need all divisors, because if the whole prefix has gcd `6`, the useful drop may come from divisor `4`, `9`, `12`, etc. Example:

`S = [12, 6]`

Both numbers have prime factors `2` and `3`, so prime-frequency counting says no prime can create a drop. But `12` is divisible by `4` and `6` is not, so ordering `[12, 6]` gives a drop at `k = 1`.

### Implementation

Precompute all divisors `d >= 2` for every possible value.

While scanning the prefix, maintain:

- `freq[d]`: how many seen numbers are divisible by `d`;
- `cnt[c]`: how many divisors currently have frequency `c`;
- block sums over `cnt`, so we can find the largest non-empty frequency below the current prefix length quickly.

When a new number `x` arrives, for every divisor `d` of `x`:

1. remove the old `freq[d]` from `cnt`;
2. increment `freq[d]`;
3. insert the new `freq[d]` into `cnt`.

For prefix length `m = i + 1`, divisors with `freq[d] = m` divide every number in the prefix, so they cannot create a gcd drop. The answer is the maximum stored frequency strictly below `m`, i.e. the maximum in frequency range `[1, m-1]`.

In zero-based code, when processing index `i`, this query range is `[1, i]`.

The first fixed version used a segment tree for this range maximum query, but that is too slow here because every number may have many divisors and every divisor update paid an extra `log n`.

The optimized code uses square-root buckets:

- `cnt[c]` records whether any divisor has current frequency `c`;
- `sum[block]` records how many active frequencies exist inside that block;
- update is `O(1)`;
- query scans block summaries from high to low, then scans inside one non-empty block.

With block size about `sqrt(n)`, each query costs `O(sqrt(n))`, while divisor updates are constant-time. This is faster in practice than a recursive segment tree because the hot loop is the divisor update loop.

### Complexity

Let `tau(x)` be the number of divisors of `x`.

Each number updates all its divisors with `O(1)` work per divisor, then does one bucket query.

Total complexity:

`O(sum tau(a_i) + n sqrt n)`.
