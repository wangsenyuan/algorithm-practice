# E. Mental Monumental (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2226/E)

**Contest:** [Codeforces Round 1095 (Div. 2)](https://codeforces.com/contest/2226)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

This is the hard version of this problem. In this version, you are required to
find the values of `f(·)` for each prefix of `a`.

For any array `[c_1, c_2, ..., c_m]`, define `f(c)` as the maximum possible
`mex(c)` that can be achieved by performing the following operation **exactly
once**:

- Choose an integer array `[b_1, b_2, ..., b_m]` such that `b_i >= 1` for all
  `1 <= i <= m`;
- Set `c_i := c_i mod b_i` for every `1 <= i <= m`.

You are given an array `a` consisting of `n` non-negative integers. For each
prefix `a^{(i)} = [a_1, a_2, ..., a_i]`, determine the value of `f(a^{(i)})`.

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

For each test case, output `n` integers — the `i`-th integer is
`f([a_1, ..., a_i])`.

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
9 9 8 2 4 4 3 5 3
```

## Sample Output 1

```text
1 2 3 4
1 2
1 2 3 4 5 5
1 2 3 4 5 5 5 6 6
```

## Note

First test case:

- `[0]` with `b = [1]` gives `mex = 1`
- `[0, 1]` with `b = [1, 2]` gives `mex = 2`
- `[0, 1, 2]` with `b = [1, 2, 3]` gives `mex = 3`
- `[0, 1, 2, 3]` with `b = [1, 2, 3, 4]` gives `mex = 4`

Second test case:

- `[6]` with `b = [6]` gives `mex = 1`
- `[6, 7]` with `b = [3, 3]` gives `mex = 2`


## ideas
1. 在C的基础上, 这时候再binary search肯定不行了
2. 假设f(i-1) = w 计算好了.
3. f(i), 这个时候, 如果添加进去了一个w, 那么f(i) = w+1了
4. 但是它还可以增加吗?
5. f(i-1) = w, 表示没有v > 2 * w (或者 v = w)
6. 这时候,加入一个v, 能够得到w的时候, 似乎不会触发连锁反应?
7. 不对, 当存在w+1的时候, 还是会连锁反应的. w+1无法构造出w, 但是遇到w后, 就会继续增长
8. 但是, 问题是w+1可能会在某个时刻被用掉了(为了得到w)

## Solution

### Why maintaining the assignments directly is difficult

As in the easy version, a value `x` can produce:

```text
x itself, or any r satisfying x > 2r.
```

The answer is non-decreasing as the prefix grows, so it is natural to keep a
current answer `mex` and try to increase it after every insertion.

However, one insertion can require many assignments to be rearranged. For
example, before inserting `6`, the value `8` might be used to produce `0`.
After `6` arrives, it is better to use `6` for `0` and release `8` for a larger
missing value. Explicitly maintaining this chain of replacements can be too
slow.

Instead of maintaining one particular matching, maintain only whether a
matching exists.

### Fixed values, offerers, and receivers

When trying to cover all values through `mex`:

- A value `x <= mex` is **fixed** if one occurrence of `x` is kept as `x`.
- Every other array element is an **offerer**.
- Every required value in `0..mex` that is not fixed is a **receiver**.

An offerer `x` can serve a receiver `r` exactly when:

```text
x >= 2r + 1.
```

For a receiver threshold `q`, define:

```text
balance[q]
    = number of offerers x with x >= 2q + 1
    - number of receivers r with r >= q.
```

Consider the receivers from largest to smallest. Every suffix whose smallest
receiver is `q` needs at least as many eligible offerers as receivers.
Therefore, a matching exists exactly when every relevant `balance[q]` is
non-negative.

If `q` is fixed, it is not a receiver and does not start a relevant suffix.
The implementation sets its segment-tree position to `inf`, excluding it from
minimum queries.

Thus feasibility is checked with:

```text
min(balance[0..mex]) >= 0.
```

### How an inserted value changes the balances

Suppose the new value is `x`.

#### Case 1: `x` is an extra offerer

This happens when:

- `x` has already appeared, so this occurrence is a duplicate; or
- `x > mex`, so it is not currently required as a fixed value.

It can serve every receiver

```text
q <= floor((x-1)/2),
```

so add `1` to:

```text
[0, floor((x-1)/2)].
```

#### Case 2: the first `x <= mex` appears

Before the insertion, `x` was a receiver. Now it can be fixed:

- remove receiver `x`, which adds `1` to `balance[0..x]`;
- set position `x` to `inf`, because `x` is no longer a receiver.

### Increasing the answer

Initially, `mex = 0`, and receiver `0` is still missing, so:

```text
balance[0] = -1.
```

After processing an insertion, while

```text
min(balance[0..mex]) >= 0,
```

all values `0..mex` can be produced, so the answer increases by one.

After incrementing `mex`, the new value `mex` becomes a receiver. Adding this
receiver subtracts `1` from:

```text
[0, mex].
```

If `mex` has already appeared, fix one occurrence immediately:

1. Set position `mex` to `inf`.
2. Remove that occurrence from the offerers, subtracting `1` from
   `[0, floor((mex-1)/2)]`.
3. Remove receiver `mex`, adding `1` to `[0, mex]`.

Then check the minimum again. This loop naturally handles a single insertion
causing the answer to increase several times.

### Segment tree

The segment tree stores:

- the minimum `balance` in each interval;
- a lazy range-add tag.

It supports:

- adding a value to an interval;
- setting one position to `inf`;
- querying the minimum on `[0, mex]`.

Every range operation costs `O(log n)`. The answer increases at most `n` times
over the entire test case, so all `while`-loop iterations together are `O(n)`.

```text
Time:  O(n log n)
Space: O(n + max(a))
```
