# Problem

You are given an array `a` of `n` integers.

In one operation, you may remove any chosen set of elements, but the chosen set must satisfy one of the following:

- all chosen elements are equal; or
- all chosen elements are pairwise distinct.

(Choosing exactly one element is always valid.)

For each `x` from `0` to `n - 1`, find the minimum number of operations needed to make the array size exactly `x`.

## Input

The first line contains integer `t` (`1 <= t <= 10^4`) — the number of test cases.

For each test case:

- The first line contains integer `n` (`1 <= n <= 3 * 10^5`) — the array size.
- The second line contains `n` integers `a1, a2, ..., an` (`1 <= ai <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `3 * 10^5`.

## Output

For each test case, print `n` integers `c0, c1, ..., c_{n-1}`, where `c_i` is the minimum number of operations required to reduce the array size to exactly `i`.

## Example

**Input**

```text
5
11
5 5 5 5 2 2 2 8 6 1 7
6
3 3 3 3 3 3
5
2 1 3 5 4
8
1 1 1 2 3 4 5 6
1
1
```

**Output**

```text
3 3 2 2 2 1 1 1 1 1 1
1 1 1 1 1 1
1 1 1 1 1
2 2 1 1 1 1 1 1
1
```

## Note

In the first sample, one possible way to realize the answers:

- `c0 = 3`: remove `a8, a9, a10, a11`, then `a1, a2, a3, a4`, then `a5, a6, a7`.
- `c1 = 3`: remove `a8, a9, a10, a11`, then `a1, a2, a3, a4`, then `a5, a6`.
- `c2 = 2`: remove `a7, a8, a9, a10, a11`, then `a1, a2, a3, a4`.
- `c3 = 2`: remove `a7, a8, a9, a10, a11`, then `a1, a2, a3`.
- `c4 = 2`: remove `a7, a8, a9, a10, a11`, then `a1, a2`.
- `c5 = 1`: remove `a1, a7, a8, a9, a10, a11`.
- `c6 = 1`: remove `a7, a8, a9, a10, a11`.
- `c7 = 1`: remove `a1, a2, a3, a4`.
- `c8 = 1`: remove `a1, a2, a3`.
- `c9 = 1`: remove `a1, a2`.
- `c10 = 1`: remove `a7`.

## Key insights (based on `solution.go`)

Let the array have value frequencies `f1, f2, ..., fm` (sum of all frequencies is `n`). The operation type does not care about positions—only about how many equal elements and how many distinct values you choose.

### 1. Model the operation via “remove distinct” `k` times

Consider using exactly `k` operations of the form “remove a set of pairwise distinct elements”.

In one such operation, you can remove **at most one** element from any value `v`. Therefore, across `k` distinct-removal operations:

- a value with frequency `f <= k` can be fully removed (you can remove one from it in each of the `k` operations);
- a value with frequency `f > k` will still have `f - k` elements remaining **if you never delete this value afterward**.

So for a fixed `k`, each value with `f > k` contributes a “leftover payload” `f - k` that we might keep or delete later.

### 2. Delete any whole value via one “remove equal” operation

Now consider operations of the form “remove a set of equal elements”.

If a value currently has `t > 0` remaining elements, you can remove **all** remaining `t` elements in **one** “equal” operation by picking exactly those remaining elements (they are all equal).

Therefore, for each value with `f > k`, we can decide:

- keep it: it contributes `f - k` to the final remaining size;
- delete it: it costs `1` extra operation, and contributes `0` remaining elements.

### 3. Reduce it to choosing which `f > k` values to keep

Fix `k`. Let `cntGreater = # { values with f > k }`.

If we keep some subset of these values, then:

- remaining size is `s = sum (f - k)` over kept values;
- number of “equal” deletions is `cntGreater - keptCount`;
- total operations is `d = k + (cntGreater - keptCount)`.

In the implementation, values with frequencies are sorted into an array `c`. For each `k`, it finds the first index `j` such that `c[j] > k`; all indices `>= j` are exactly the “values with f > k” (including zeros for unused values).

Then it iterates `i` over possible cut points and uses the convention:

- keep frequencies `c[j], c[j+1], ..., c[i-1]`;
- delete frequencies `c[i], ..., c[n-1]` (each deleted value costs one “equal” operation).

Under this convention:

- `s` (in code) is the cumulative sum `sum_{t=j..i-1} (c[t] - k)` (leftover elements);
- `d` (in code) is `k + (n - i)` which equals `k + (#deleted values among those > k)`.

Finally, the code updates:

- `res[s] = min(res[s], d)`,

meaning: the minimum number of operations to leave exactly `s` elements.

### 4. Why there is a final prefix-min step

The implementation finishes with:

- `res[x] = min(res[x], res[x-1])` for `x = 1..n-1`.

This enforces the natural monotonicity: if you can achieve a smaller remaining size with `d` operations, then allowing a larger remaining size cannot require more operations, so `res[x]` should be non-increasing as `x` grows.

## How the code maps to these ideas

The Kotlin/Go logic in `src/codeforces/set2/set21/set214/set2141/f/solution.go` does exactly the following:

1. Build the frequency list `c` of length `n` where `c[v-1]` is the count of value `v` in the array, then sort `c`.
2. Initialize `res[x] = n` for all `x` (worst-case upper bound).
3. For each `k` from `0` to `n`:
   - find `j` where `c[j] > k` (all indices before `j` have `c[idx] <= k`);
   - maintain `s` as cumulative leftover `sum(c[t] - k)` while moving a cut `i` from `j` upward;
   - compute `d = k + n - i`;
   - update `res[s] = min(res[s], d)` when `s < n`.
4. Apply the monotone improvement `res[x] = min(res[x], res[x-1])`.

## Editorial (dual formulation)

Rephrase the problem: for each number of operations, compute the **maximum** number of elements that can be removed. From that, the original answer follows: if you can remove `m` elements in `k` operations, you can also remove any smaller number of elements in no more than `k` operations.

In this version, each operation either removes one occurrence of **every** remaining distinct value, or removes **all** occurrences of the current maximum-frequency value.

Build `C_1, C_2, ..., C_n` where `C_i` is the count of value `i` in `a`. Only the multiset of counts matters, so sort `C`.

Each operation on `C` does one of:

- subtract `1` from every positive entry (floor at `0`); or
- set the current maximum entry to `0`.

Order does not matter: if you use `x` operations of the first type and `y` of the second, then `y` maxima become `0` and the rest decrease by `x`, regardless of interleaving.

Iterate over candidate pairs `(x, y)`. Naively `O(n^2)` pairs, but only `O(n)` matter: after removing `y` maxima, there is no reason to take `x` larger than the next maximum; across all `y`, the relevant `x` values number at most `(sum of C_i) + n ≤ 2n`.

To evaluate `(x, y)` on sorted `C`: a two-pointer finds the first index where `C_i > x` — the prefix before that becomes `0`, suffix entries become `C_i - x`. Prefix sums help sum the `y` largest values among the modified array.
