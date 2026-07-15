# A - Unusual-Constraint Knapsack

[Problem link](https://atcoder.jp/contests/arc223/tasks/arc223_a)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 400 points

## Problem Statement

There are `N` items `1, 2, ..., N`.
Item `i` has weight `w_i` and value `v_i`.
For `i = 2, 3, ..., N`, the following holds:

- `sum_{j=1}^{i-1} w_j < w_i`

Choose some (possibly zero) items such that the total weight is at most `W`.
Find the maximum possible total value of the chosen items.

Solve `T` test cases per input.

## Constraints

- `1 <= T <= 2 * 10^5`
- `1 <= N <= 60`
- `1 <= w_i <= W <= 10^18`
- `sum_{j=1}^{i-1} w_j < w_i` (`2 <= i <= N`)
- `1 <= v_i <= 10^16`
- The sum of `N` over all test cases is at most `2 * 10^5`
- All input values are integers

## Input

```
T
case_1
case_2
...
case_T
```

Each test case:

```
N W
w_1 v_1
w_2 v_2
...
w_N v_N
```

## Output

Output `T` lines. The `t`-th line should contain the answer for the `t`-th
test case.

## Samples

### Sample 1

Input:

```
1
3 10
1 2
3 1
9 3
```

Output:

```
5
```

Choosing items `1` and `3` gives total weight `10` and total value `5`.

### Sample 2

Input:

```
1
3 20
1 2
3 1
9 3
```

Output:

```
6
```

### Sample 3

Input:

```
1
4 15
1 1
3 2
5 3
10 4
```

Output:

```
7
```

## Solution

The unusual weight condition is:

```text
w[i] > w[0] + w[1] + ... + w[i-1]
```

This means that whenever item `i` fits into the remaining capacity, **all
earlier items together also fit into that same capacity**. This observation
turns the usual exponential knapsack recursion into a linear one.

### Prefix sums

The implementation builds two prefix arrays:

```text
pref1[i] = total weight of items 0 ... i-1, capped at W+1
pref2[i] = total value  of items 0 ... i-1
```

If `pref1[n] <= W`, every item fits. Since every value is positive, taking all
items is optimal, so the answer is immediately `pref2[n]`.

The cap at `W+1` is sufficient because `pref1` is used only to distinguish
whether the total weight is at most `W` or greater than `W`.

### Recurrence

Let:

```text
f(i, rest)
```

be the maximum value obtainable from items `0, 1, ..., i` with remaining
capacity `rest`.

If `i < 0`, there are no items, so the answer is `0`.

If `w[i] > rest`, item `i` cannot be selected:

```text
f(i, rest) = f(i-1, rest)
```

Otherwise, item `i` fits. There are two choices.

#### Do not take item `i`

Because

```text
sum(w[0 ... i-1]) < w[i] <= rest,
```

all earlier items fit together. Their values are positive, so the best choice
without item `i` is to take every earlier item:

```text
pref2[i]
```

There is no need to call `f(i-1, rest)` for this branch.

#### Take item `i`

The remaining capacity becomes `rest-w[i]`, and the best additional value is:

```text
v[i] + f(i-1, rest-w[i])
```

Therefore:

```text
f(i, rest) = max(
    pref2[i],
    v[i] + f(i-1, rest-w[i])
)
```

One subtle point is that comparing only `v[i]` with `pref2[i]` is not enough.
Even when `v[i] <= pref2[i]`, taking item `i` together with some earlier items
may be optimal. For example, if `(w, v)` is `[(1, 5), (3, 4)]` and `W = 4`,
taking both items gives value `9`.

### Why the recursion is linear

Ordinary knapsack recursion explores both “take” and “skip” recursively. Here,
the skip branch is known directly as `pref2[i]` whenever item `i` fits. Thus
each call makes at most one recursive call:

- skip recursively if item `i` does not fit;
- recurse only into the take branch if item `i` fits.

The index decreases by one on every call, so there are at most `N` calls and
no memoization is necessary.

### Correctness Proof

We prove that `f(i, rest)` returns the maximum value obtainable from items
`0 ... i` with capacity `rest`.

The claim is true for `i < 0`, where no item can be selected and the returned
value is `0`.

Assume the claim holds for `i-1`.

- If `w[i] > rest`, item `i` is impossible to select. The optimum is exactly
  the optimum among items `0 ... i-1`, which is `f(i-1, rest)` by the induction
  hypothesis.
- If `w[i] <= rest`, an optimal selection either excludes or includes item
  `i`. If it excludes item `i`, all earlier items fit together because their
  total weight is less than `w[i]`, and taking all of them gives `pref2[i]`.
  If it includes item `i`, the remaining problem has capacity `rest-w[i]`, so
  by the induction hypothesis its best value is
  `v[i] + f(i-1, rest-w[i])`. Taking the maximum covers both possibilities.

Thus the recurrence is correct. Calling `f(n-1, W)` therefore returns the
maximum obtainable value for the original problem.

### Complexity

Building the prefix arrays and following the single recursive path both take
linear time.

- Time: `O(N)` per test case
- Space: `O(N)` for the prefix arrays and recursion stack
