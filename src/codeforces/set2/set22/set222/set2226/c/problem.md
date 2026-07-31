# C. Mental Monumental (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2226/C)

**Contest:** [Codeforces contest 2226](https://codeforces.com/contest/2226)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

This is the easy version of this problem. The difference between the versions is
that in this version, you are required to find only the value of `f(a)`.

For any array `[c_1, c_2, ..., c_m]`, define `f(c)` as the maximum possible
`mex(c)` that can be achieved by performing the following operation **exactly
once**:

- Choose an integer array `[b_1, b_2, ..., b_m]` such that `b_i >= 1` for all
  `1 <= i <= m`;
- Set `c_i := c_i mod b_i` for every `1 <= i <= m`.

You are given an array `a` consisting of `n` non-negative integers. Determine
the value of `f(a)`.

Here `mex(c)` is the smallest non-negative integer that does not appear in `c`.

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 <= t <= 10^4`).

Each test case:

- the first line contains a single integer `n` (`1 <= n <= 2 * 10^5`);
- the second line contains `n` integers `a_1, ..., a_n` (`0 <= a_i <= 10^6`).

It is guaranteed that the sum of `n` over all test cases does not exceed
`2 * 10^5`, and the sum of `max(a_1, ..., a_n)` over all test cases does not
exceed `10^6`.

## Output

For each test case, output a single integer — the value of `f(a)`.

## Sample Input 1

```text
4
4
0 1 2 3
2
6 7
6
8 1 7 6 4 3
9
9 8 2 4 4 3 5 3 4
```

## Sample Output 1

```text
4
2
5
6
```

## Note

- First test case: choosing `b = [1, 2, 3, 4]` leaves `a` unchanged and
  `mex(a) = 4`.
- Second test case: choosing `b = [3, 3]` makes `a = [0, 1]`, so `mex(a) = 2`.


## ideas
1. 假设需要mex = w, 那么需要找到a[?] = 0, 1, 2, .. w - 1
2. 假设某个数v < w 不存在, 那么a[i] % b[i] = v
3.  a[i] = v + b[i] * k
4.  比如2不存在, 6 % 4 = 2, 5 % 3 = 2,
5.  那么应该选择哪个呢? 这个时候, 应该应该选择最小的那个数, 5?
6.  好像不对, 因为5后面用得到
7.  从v可以得到余数0, 1, 2, ... (v / 2 - 1)
8.  好像确实应该使用更小的那个数, 但是应该是多余的那个数

## Solution

### When can `x` become `r`?

There are two possibilities:

- If `x = r`, choose any `b > x`, so `x mod b = x = r`.
- If `x > r`, choose `b = x-r`. This is valid exactly when
  `b > r`, or equivalently `x > 2r`.

Conversely, if `x mod b = r` and `x != r`, then

```text
x = r + k * b,  k >= 1,  b > r,
```

so necessarily `x > 2r`. Therefore:

```text
x can produce r  <=>  x = r or x > 2r.
```

### Checking whether MEX `w` is achievable

To obtain MEX `w`, the final array must contain every value
`0, 1, ..., w-1`. Any unused element can be changed to `0` with `b = 1`,
so it does not introduce `w`.

First sort the array. For every value `v < w`, keep one occurrence of `v`
to represent itself. Put all duplicate occurrences and all values `v >= w`
into a sorted list `todo`.

Now process every missing value `r` from `0` to `w-1`:

1. Remove every `x <= 2r` from the front of `todo`.
   Such an `x` cannot produce `r`. It also cannot produce any later missing
   value, because the condition only becomes stricter as `r` increases.
2. If `todo` is empty, MEX `w` is impossible.
3. Otherwise, consume its smallest element. It satisfies `x > 2r`, so it can
   produce `r`.

Keeping an existing `v` as `v` is safe. If some construction instead used
`v` for a smaller residue and used a larger `x` to produce `v`, the two
assignments can be swapped: keep `v` unchanged and use `x` for the smaller
residue.

Consuming the smallest eligible spare element is also optimal. A larger
element can do everything that the smaller element can do, and may still be
needed for a later residue whose threshold `2r` is higher.

### Binary search

If MEX `w` is achievable, then every positive `k < w` is also achievable:
keep the elements assigned to `0, 1, ..., k-1` and change every other element
to `0`. The resulting array contains the required prefix but not `k`.
Thus the predicate used by the search is monotonic.

Binary-search the first `w` for which the check fails. The answer is `w-1`.
Since an array of length `n` cannot have MEX greater than `n`, it is sufficient
to search through `0..n`.

### Complexity

Sorting costs `O(n log n)`. One feasibility check is `O(n)`, and binary search
performs `O(log n)` checks, so the total complexity is:

```text
Time:  O(n log n)
Space: O(n)
```
