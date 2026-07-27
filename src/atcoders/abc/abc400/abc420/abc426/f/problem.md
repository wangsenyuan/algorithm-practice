# F - Clearance

[Problem link](https://atcoder.jp/contests/abc426/tasks/abc426_f)

**Contest:** [AtCoder Beginner Contest 426](https://atcoder.jp/contests/abc426)

time limit: 5 sec

memory limit: 1024 MiB

score: 525 points

## Problem

There are `N` products; product `i` has stock `A_i`.

Process `Q` orders in order. The `i`-th order is:

- Buy `k_i` units each of products `l_i, l_i+1, ..., r_i`. If a product has
  fewer than `k_i` units left, buy all remaining units. Report the total number
  of units bought in this order.

After each order, stocks are reduced before the next order.

## Constraints

- All input values are integers
- `1 <= N <= 3 * 10^5`
- `1 <= A_i <= 10^15`
- `1 <= Q <= 3 * 10^5`
- `1 <= l_i <= r_i <= N`
- `1 <= k_i <= 10^9`

## Input

```text
N
A_1 A_2 ... A_N
Q
l_1 r_1 k_1
...
l_Q r_Q k_Q
```

## Output

Print `Q` lines. The `i`-th line is the total units bought in the `i`-th order.

## Sample 1

```text
Input
6
2 6 4 5 7 5
5
1 6 1
3 5 4
4 4 1
2 5 1
1 6 100

Output
6
11
0
2
10
```

## Solution Summary

### Key observation

For one product with current stock `s`, an order asking for `k` units buys
`min(s, k)` units.

Instead of calculating this value separately for every product, first assume
that every product which has not yet been marked sold out can provide all `k`
units. If there are `c` such products in the query range, the preliminary
answer is:

```text
answer = c * k
```

Then subtract `k` from the stock of every product in the range.

- If `s >= k`, its new stock `s - k` is non-negative, so the preliminary
  contribution `k` was correct.
- If `s < k`, its new stock is the negative value `s - k`. Adding this value
  to the preliminary contribution corrects it:

```text
k + (s - k) = s
```

For example, if `s = 3` and `k = 5`, we first count `5`. The new stock is
`3 - 5 = -2`, so the corrected contribution is `5 + (-2) = 3`.

After applying this correction, the product is permanently marked sold out.
Therefore, each product needs this special handling at most once.

### Segment tree state

For every segment-tree node, maintain:

- `arr[i]`: the minimum remaining stock in the segment;
- `cnt[i]`: the number of products in the segment which have not yet been
  marked sold out;
- `lazy[i]`: a pending value to add to every stock in the segment.

A sold-out product is assigned the large value `inf` and its count becomes
zero. The large value prevents it from becoming the minimum again, while the
zero count prevents future orders from including it in their preliminary
answers.

The tree supports four operations:

1. `count(L, R)` returns the number of active products in the range.
2. `add(L, R, -k)` subtracts `k` from every stock in the range.
3. `firstNegative(L, R)` uses the stored minimums to locate a product whose
   stock became negative.
4. `remove(pos)` returns that negative stock, replaces it with `inf`, and
   changes its active count to zero.

### Processing an order

For an order `(L, R, k)`:

```text
answer = count(L, R) * k
add(L, R, -k)

while there is a negative stock in [L, R]:
    pos = firstNegative(L, R)
    answer += remove(pos)
```

Although the last line uses addition, `remove(pos)` returns a negative value.
Thus it subtracts exactly the amount by which the preliminary answer
overcounted that product.

If a product's stock becomes exactly zero, it is not removed immediately
because `firstNegative` searches only for negative values. This is harmless:
the next order covering that product tentatively counts `k`, changes its stock
from zero to `-k`, and then adds `-k` as the correction, so its actual
contribution is zero before it is marked sold out.

### Correctness

Consider any product in the current query range.

- If it was already marked sold out, its active count is zero, so it
  contributes nothing.
- If its stock is at least `k`, the preliminary answer counts `k`, which is
  exactly the amount bought, and its stock remains non-negative.
- If its stock is less than `k`, the preliminary answer counts `k`. Its new
  stock is `s - k < 0`, and adding this value changes its contribution to
  `k + (s - k) = s`, exactly all of its remaining stock. It is then marked sold
  out.

Therefore, every product contributes exactly `min(s, k)` to the answer, and
the segment tree represents the correct state for the next order.

### Complexity

The range count and range subtraction each take `O(log N)`.

Finding and removing one newly sold-out product takes `O(log N)`. Each product
is removed at most once over all queries, so all removal work costs
`O(N log N)` in total.

The total time complexity is:

```text
O((N + Q) log N)
```

The segment tree uses `O(N)` memory.

## Alternative Implementation: Fused Recursive Buy

The implementation above separates one order into several operations:

1. count active products;
2. subtract `k` from the range;
3. find every negative stock;
4. remove those products and correct the answer.

The current implementation uses the same central observation, but fuses all
four steps into one recursive `buy(L, R, v)` operation. This avoids explicitly
creating a preliminary answer and then correcting it.

### Invariant

Each segment-tree node still stores:

- `arr[i]`: the minimum stock among the products represented by the node;
- `cnt[i]`: the number of products which are not sold out;
- `lazy[i]`: the amount that still needs to be subtracted from every child.

When a product sells out, set its stock to `inf` and its count to zero.
Therefore, it no longer affects either the minimum stock or future answers.

All active products have strictly positive stock. A product whose stock is
exactly equal to the current order size is handled as sold out immediately.

### Recursive cases

Suppose the current node represents `[l, r]`.

#### Case 1: the segment is already sold out

If `cnt[i] == 0`, no product in this segment has stock, so its contribution is
zero:

```text
return 0
```

#### Case 2: every product can satisfy the order

If `[l, r]` is fully contained in `[L, R]` and:

```text
arr[i] > v
```

then even the product with minimum stock has more than `v` units. Consequently,
every active product in the segment can provide exactly `v` units.

Subtract `v` lazily from the whole segment and return:

```text
v * cnt[i]
```

No descent is necessary.

#### Case 3: a leaf cannot satisfy the full order

If the recursion reaches a leaf, Case 2 did not apply. Hence its remaining
stock is at most `v`.

The order buys all of that stock:

```text
answer = arr[i]
```

Then mark the product sold out:

```text
arr[i] = inf
cnt[i] = 0
```

Notice why Case 2 uses `arr[i] > v` rather than `arr[i] >= v`. If the stock is
exactly `v`, the recursion reaches the leaf, buys those `v` units, and
immediately removes the now-empty product.

#### Case 4: recurse into the children

Otherwise, push the lazy subtraction to the children and recurse only into
children intersecting `[L, R]`. Afterward, rebuild the current node:

```text
cnt[i] = cnt[left] + cnt[right]
arr[i] = min(arr[left], arr[right])
```

Thus the parent again stores the correct active count and minimum stock.

### Why this is equivalent

For each product with stock `s` in the query range:

- if `s > v`, the recursion eventually reaches a fully covered segment where
  the minimum is greater than `v`, and the product contributes `v`;
- if `s <= v`, the recursion reaches its leaf, the product contributes all
  `s` remaining units, and it is marked sold out;
- if it was already sold out, it contributes zero.

Therefore, every product contributes exactly:

```text
min(s, v)
```

which is precisely the required amount.

### Complexity

A query normally visits only the `O(log N)` nodes along its range boundaries
and fully covered nodes where the recursion stops.

Extra descent occurs only when a covered segment contains a product with stock
at most `v`. That descent eventually sells out at least one such product. Each
product can be sold out only once, so all such extra work over all queries is
`O(N log N)`.

The total time complexity remains:

```text
O((N + Q) log N)
```

The memory complexity is `O(N)`.
