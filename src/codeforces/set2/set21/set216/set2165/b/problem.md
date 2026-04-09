# Problem

You are given a multiset $a$, which consists of $n$ integers $a_1, a_2, \ldots, a_n$. You would like to generate a new multiset $s$ through the following procedure:

1. Partition $a$ into any number of non-empty multisets $x_1, x_2, \ldots, x_k$, such that each element of $a$ belongs to exactly one of these multisets.
2. Initially, $s$ is empty. From each $x_i$, choose one of its modes[^mode] and insert it into $s$.

Please count the number of different multisets $s$ that can be generated through the procedure, modulo $998244353$.

Please note that the number of different multisets is counted, which means that the order of elements does not matter. However, the count of each element does matter, i.e. $\{1, 1, 2\}$, $\{1, 2\}$, $\{1, 1, 2, 2\}$ are all considered different.

[^mode]: The mode of a multiset is defined as the element which appears the most; if several elements are tied as the maximum, then all of them are considered modes.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 5000$). The description of the test cases follows.

The first line of each test case contains a single integer $n$ ($1 \le n \le 5000$) — the size of multiset $a$.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le n$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $5000$.

## Output

For each test case, print one line containing a single integer — the number of different multisets you can obtain, modulo $998244353$.

## Example

### Input

```
5
3
1 2 3
3
1 1 1
3
1 2 2
10
1 1 1 1 2 2 2 3 3 4
10
1 1 1 2 2 2 3 3 3 4
```

### Output

```
7
3
4
111
126
```

## Note

In the first test case, any non-empty subset of $\{1, 2, 3\}$ can be achieved, for a total of $7$ multisets.

In the third test case, we can generate $4$ different multisets:

- Partition the elements into multiset $\{1, 2, 2\}$, resulting in multiset $\{2\}$.
- Partition the elements into multisets $\{1, 2\}$, $\{2\}$, resulting in multiset $\{2, 2\}$.
- Partition the elements into multisets $\{1\}$, $\{2, 2\}$, resulting in multiset $\{1, 2\}$.
- Partition the elements into multisets $\{1\}$, $\{2\}$, $\{2\}$, resulting in multiset $\{1, 2, 2\}$.

It can be proven that no other multisets are possible.


### ideas
1. have no ideas at all!
2. Let `cnt[v]` be the frequency of value `v` in the original multiset, and let `mx = max(cnt[v])`.
3. Try to characterize a possible final multiset `s` only by its support `T`:
   the set of values that appear at least once in `s`.
4. Suppose the partition parts are `x_1, ..., x_k`. If the chosen mode in group `x_i`
   appears `r_i` times inside that group, then every value appears at most `r_i` times
   in that group.
5. Therefore for any fixed original value `v`, all `cnt[v]` copies of `v` must be placed
   across the groups, and group `i` can hold at most `r_i` copies of `v`. So:

   `cnt[v] <= r_1 + r_2 + ... + r_k`

6. Now ask: how large can `r_1 + ... + r_k` be?
   Every group chooses a mode from the support `T`. If a group chooses value `x`,
   its contribution to the sum is the number of copies of `x` used in that group.
   Summed over all groups choosing `x`, this cannot exceed `cnt[x]`.
   Hence:

   `r_1 + ... + r_k <= sum(cnt[x] for x in T)`

7. Combine the two facts for the most frequent value:

   `mx <= r_1 + ... + r_k <= sum(cnt[x] for x in T)`

   So a necessary condition is:

   `sum(cnt[x] for x in T) >= mx`

8. This condition is also sufficient. Think in a "capacity" way:
   if you choose any support `T` with total frequency at least `mx`, then you can split
   the copies of each `x in T` into positive parts; these parts become the mode-frequencies
   of groups labeled by `x`. The total capacity is enough to place every unchosen value.

9. Once `T` is fixed, how many different multisets `s` have exactly this support?
   For each `x in T`, its multiplicity in `s` can be any integer from `1` to `cnt[x]`.
   So the number of such `s` is:

   `prod(cnt[x] for x in T)`

10. That suggests a knapsack / DP over the distinct values:
    - skip value with frequency `c`
    - take it into the support, multiplying ways by `c` and increasing support-sum by `c`

11. Final counting formula:
    sum `prod(cnt[x])` over all support sets `T` such that

    `sum(cnt[x] for x in T) >= mx`
