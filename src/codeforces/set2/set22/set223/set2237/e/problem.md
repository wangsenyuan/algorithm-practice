# E. Permutation Commutation

[Problem link](https://codeforces.com/problemset/problem/2237/E)

**Contest:** [Order Capital Round 2 (Codeforces Round 1104, Div. 1 + Div. 2)](https://codeforces.com/contest/2237)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Quack the Duck has a permutation `a` of length `n` and an incomplete sequence
`b_1, b_2, ..., b_n`.

Each element of `b` is either `-1` or an integer from `1` to `n`. Each integer from
`1` to `n` appears at most once in `b`.

Quack hopes to complete `b` into a permutation that commutes with `a`. In other words,
after replacing every `-1` in `b`, the equality `a_{b_i} = b_{a_i}` should hold for every
`1 <= i <= n`.

Ja the Ghost wants to help Quack. Among all possible ways to complete `b`, he wants to
find the lexicographically smallest one.

Determine whether such a completion exists. If it exists, output the lexicographically
smallest valid permutation `b`. Otherwise, report that it is impossible.

A permutation of length `n` is an array consisting of `n` distinct integers from `1` to
`n` in arbitrary order.

An array `p` is lexicographically smaller than an array `q` of the same size if and only
if `p != q`, and in the first position where `p` and `q` differ, the array `p` has a
smaller element than the corresponding element in `q`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases
`t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`) — the
length of the permutation.

The second line of each test case contains `n` integers `a_1, a_2, ..., a_n`
(`1 <= a_i <= n`) — the permutation `a`.

The third line of each test case contains `n` integers `b_1, b_2, ..., b_n`
(`b_i = -1` or `1 <= b_i <= n`) — the incomplete sequence `b`.

It is guaranteed that each integer from `1` to `n` appears at most once in `b`.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print `"YES"` if the answer exists, and `"NO"` otherwise.

You can output the answer in any case (upper or lower).

If the answer exists, in the next line output `n` integers `p_1, p_2, ..., p_n` — the
lexicographically smallest valid sequence after replacing every `-1` in `b`.

The sequence `p` must be a permutation, that is, each integer from `1` to `n` must appear
exactly once in `p`. Also, it must satisfy `a_{p_i} = p_{a_i}` for every `1 <= i <= n`.

## Example

### Input

```text
12
3
2 3 1
-1 -1 -1
4
2 1 4 3
-1 -1 4 -1
4
2 1 4 3
3 1 -1 -1
4
2 1 4 3
1 -1 -1 2
5
2 3 1 5 4
2 -1 -1 -1 -1
5
2 3 1 5 4
4 -1 -1 -1 -1
6
2 3 1 5 6 4
4 -1 -1 -1 -1 -1
6
2 1 4 3 6 5
-1 3 -1 -1 -1 -1
6
3 5 6 2 1 4
-1 -1 -1 3 6 -1
7
2 3 1 5 4 6 7
-1 -1 -1 -1 -1 7 -1
8
2 3 4 1 6 7 8 5
5 7 -1 -1 -1 -1 -1 -1
8
2 3 4 1 6 7 8 5
5 -1 -1 -1 -1 -1 -1 -1
```

### Output

```text
YES
1 2 3
YES
1 2 4 3
NO
NO
YES
2 3 1 4 5
NO
YES
4 5 6 1 2 3
YES
4 3 1 2 5 6
NO
YES
1 2 3 4 5 7 6
NO
YES
5 6 7 8 1 2 3 4
```

### Note

- In the first test case, `b = [1, 2, 3]` commutes with any permutation `a`. Since all
  elements of `b` are unknown, this is also the lexicographically smallest possible valid
  permutation.
- In the second test case, `a = [2, 1, 4, 3]` and `b_3 = 4`. Since `a_3 = 4`, the
  condition for `i = 3` gives `a_{b_3} = b_{a_3}`, so `a_4 = b_4`, hence `b_4 = 3`. The
  remaining values are `1` and `2`, and the lexicographically smallest valid choice is
  `b_1 = 1`, `b_2 = 2`. Thus the answer is `[1, 2, 4, 3]`.
- In the third test case, `a = [2, 1, 4, 3]`, `b_1 = 3`, and `b_2 = 1`. For `i = 1`, the
  condition requires `a_{b_1} = b_{a_1}`. However, `a_{b_1} = a_3 = 4`, while
  `b_{a_1} = b_2 = 1`. Since `4 != 1`, no valid completion exists.

## Summary

The key observation is that `a` splits the positions into directed cycles. If
`b` commutes with `a`, then applying `b` to one element of an `a`-cycle forces
the whole cycle:

```text
b[a[i]] = a[b[i]]
```

So if `b[i]` is known, then `b[a[i]]` must be `a[b[i]]`, then
`b[a[a[i]]]` must be `a[a[b[i]]]`, and so on. The first pass follows each
cycle that already has a known value and propagates all forced values. If a
position is already assigned to a different value, the completion is impossible.
After this propagation, the code also counts images and rejects any duplicate,
because the final `b` must be a permutation.

Now every still-unassigned domain cycle has no fixed value at all. Such a cycle
of length `k` can only be mapped to an unused target cycle of `a` with the same
length `k`; otherwise the commutation equation cannot stay consistent after
walking around the cycle.

For the lexicographically smallest answer, process cycles by the first position
where they are discovered. For a free domain cycle, that first position is the
smallest index in the cycle. For a free target cycle, the discovered
representative is the smallest unused value in that cycle. Pairing same-length
domain cycles and target cycles in this order makes the earliest still-free
position receive the smallest possible value. Inside a paired cycle, map the
representative `x` to representative `y`, then continue with
`b[a[x]] = a[y]`, `b[a[a[x]]] = a[a[y]]`, and so on.

The implementation does this in four steps:

1. Convert to zero-based indexing, keeping `-1` as unknown.
2. Propagate every known assignment around its `a`-cycle and reject conflicts.
3. Count used images and reject duplicates.
4. Group unassigned domain cycles and unused target cycles by length, then pair
   equal-length cycles in discovery order and fill them.

The work is linear in the permutation size. Each element is visited a constant
number of times, so the complexity is `O(n)` time and `O(n)` memory per test
case.
