# E - A += v

[Problem link](https://atcoder.jp/contests/abc449/tasks/abc449_e)

## Problem Summary

Given integers `N`, `M`, and an array `A` of length `N` whose values are in
`[1, M]`, repeat the following operation a very large number of times:

- Choose the value `v` in `[1, M]` with the minimum current occurrence count in
  `A`. If multiple values have that minimum count, choose the smallest value.
- Append `v` to the end of `A`.

For each query `X_i`, output the value at position `X_i` after all operations.

## Constraints

- `1 <= N, M <= 5 * 10^5`
- `1 <= A_i <= M`
- `1 <= Q <= 2 * 10^5`
- `1 <= X_i <= 10^18`

## Sample 1

```text
3 3
1 1 2
8
1
2
3
4
5
6
7
8
```

```text
1
1
2
3
2
3
1
2
```

## Sample 2

```text
7 30
20 26 3 14 4 4 9
10
31
9
21
23
97
99
30
79
57
3
```

```text
30
2
18
21
7
9
29
19
27
3
```

## Solution Summary

Let `cnt[v]` be the initial occurrence count of value `v`.

After a value `v` has been appended `t` times, the next time it can be chosen
corresponds to the pair:

```text
(cnt[v] + t, v)
```

The operation always chooses the smallest pair lexicographically: first by
current count, then by value. Therefore, the infinite appended suffix is exactly
all pairs `(cnt[v] + t, v)` for every `v` and `t >= 0`, sorted
lexicographically.

For a query position inside the original prefix, the answer is just `A[X]`.
Otherwise let `k = X - N`, the 1-based position in the appended suffix.

Define:

```text
F(L) = number of appended pairs with count < L
     = sum over v of max(0, L - cnt[v])
```

Binary search the smallest `T` such that `F(T) >= k`. Then the answer lies at
count level `T - 1`. Its rank among values active at that level is:

```text
k - F(T - 1)
```

The active values are exactly those with `cnt[v] <= T - 1`, and they appear in
increasing value order. After computing all query levels, process them offline in
increasing level order. Add values to a Fenwick tree when their initial count
becomes active, then use order statistics on the Fenwick tree to find the
required rank.

Sorting the counts lets us compute `F(L)` in `O(log M)` using prefix sums.

## Complexity

For each query, binary search costs `O(log X * log M)`. The offline Fenwick phase
costs `O((M + Q) log M)`. Memory usage is `O(M + Q)`.


## ideas
1. 肯定不能按照题目的方式,不断添加v上去
2. 应该计算在k-1时,最少的数时多少, 那么第k个数就是它
3. 会不会在某个点开始,就是从1,2...m一次添加的?
4. 有这样的一个位置,假设最多的数是x, 且freq(x) = w, 那么经过 w * m - len(a) 次后, 所有的数都是w
5. 在这之后,都是 1,2, 3... 
6. 那么在这之前的变化, 按照freq 升序排列, 当freq相同时,按照数字升序
7. 

## Explanation

### Current Solution: Level Counts Until the Maximum

Let `freq[v]` be the initial count of value `v`. Sort all values by
`(freq[v], v)`. Suppose the maximum initial count is `w`.

Before every value reaches count `w`, the process only raises lower-frequency
values. The total number of appended values needed to make all `M` values have
count `w` is:

```text
w * M - N
```

After that point, all values always have equal counts before each full round, so
the appended sequence becomes periodic:

```text
1, 2, 3, ..., M, 1, 2, 3, ..., M, ...
```

So each query after the original prefix is split into two cases:

- If `k > w * M - N`, answer directly from the final cycle.
- Otherwise, answer from the leveling phase before all counts reach `w`.

For the leveling phase, consider the sorted array `arr`. Before `arr[i]` becomes
active, the earlier `i` values must all be raised to `arr[i].first`. The number
of append operations needed for that is:

```text
i * arr[i].first - pref[i]
```

where `pref[i]` is the sum of the first `i` frequencies in sorted order.

Binary search the first `i` where this value is at least `k`. Then the answer is
inside the previous active group `arr[0..i-1]`. Inside such a group, all active
values have equal current frequency, so they are chosen by increasing value.
The segment tree stores the active values by original label and supports finding
the `k`-th smallest active value.

This is the approach used by `solution.go`.

### Alternative Solution: Query Count Levels Directly

Another way to view the same process is to say that every appended occurrence of
value `v` corresponds to a pair:

```text
(freq[v] + t, v)
```

where `t >= 0` is how many times `v` has already been appended. The whole
appended suffix is all such pairs sorted lexicographically.

For a count level `L`, the number of appended pairs with count less than `L` is:

```text
F(L) = sum(max(0, L - freq[v]))
```

Using sorted frequencies and prefix sums, `F(L)` can be computed in `O(log M)`.
For a suffix query `k`, binary search the smallest level `T` with
`F(T) >= k`. The answer is at level `T - 1`, with rank:

```text
k - F(T - 1)
```

among values satisfying `freq[v] <= T - 1`, ordered by value. Process all
queries offline by increasing level, add active values into a Fenwick tree, and
use order statistics to find the required rank.

This alternative does not special-case the final `1..M` cycle. It treats the
cycle as just more count levels.

### Comparison

Both solutions rely on the same invariant: at any moment, the next appended
value is the smallest label among those with minimum current frequency.

The current solution separates the process into two phases:

1. Raise every value to the initial maximum count `w`.
2. Use the simple periodic cycle after all counts are equal.

The alternative solution keeps one uniform model for all future positions by
binary searching the count level that contains each query.

The current solution has a simpler tail case and avoids binary searching up to
large query values. The alternative solution is more uniform and directly
matches the infinite sorted-pair interpretation. Both use an order-statistics
data structure over active labels and run in roughly `O((M + Q) log M)` after
sorting.
