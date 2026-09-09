# G. Drowning

[Problem link](https://codeforces.com/problemset/problem/2227/G)

**Contest:** [Codeforces Round 1096 (Div. 3)](https://codeforces.com/contest/2227)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Yousef has an array `a` consisting of `n` positive integers.

He defines a reduction operation on any array `c` of length `|c| >= 3`:

- Choose an index `i` (`1 < i < |c|`) such that `c_{i-1} + c_{i+1} > c_i`.
- Replace the triplet `{c_{i-1}, c_i, c_{i+1}}` with a single integer `x = c_{i-1} - c_i + c_{i+1}`.

The new integer `x` occupies the position previously held by the triplet, and the length of the array decreases by `2`.

An array is considered good if it can be reduced to a single element by performing the operation above zero or more times. Note that an array of length `1` is always good.

Yousef wants you to count the number of pairs `(l, r)` (`1 <= l <= r <= n`) such that the subarray `a[l, r]` is good.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases. The description of each test case follows.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`) — the size of the array.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 10^9`) — the elements of the array.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output a single integer — the number of good subarrays.

## Example

### Input

```text
4
3
10 20 10
5
1 1 1 1 1
4
5 1 5 1
1
100
```

### Output

```text
3
9
5
1
```

### Note

In the first example, `a = [10, 20, 10]`. Subarrays `[10]`, `[20]`, and `[10]` are all good (`3` total). The subarray `[10, 20, 10]` is not good. To reduce it, we must pick `i = 2`. The condition `a_1 + a_3 > a_2` becomes `10 + 10 > 20`, which is `20 > 20` (false).

In the second example, `a = [1, 1, 1, 1, 1]`:

- All `5` subarrays of length `1` are good.
- All `4` subarrays of length `2` are not good.
- All `3` subarrays of length `3` (which are `[1, 1, 1]`) are good because `1 + 1 > 1`.
- All `2` subarrays of length `4` are not good.
- The subarray of length `5` is good: `[1, 1, 1, 1, 1]` with `i = 2` becomes `[1, 1, 1]`, then with `i = 2` becomes `[1]`.
- Total good subarrays `= 5 + 3 + 1 = 9`.

## ideas
1. c[i-1] + c[i+1] > c[i]
2. 然后替换为c[i-1] + c[i+1] - c[i]
3. 这里不变的东西, 好像是奇数位的sum - 偶数位的sum
4. [1, 2, 3] => 2
5. 如果奇数位的sum > 偶数位的sum, 那么一定可以吗?
6. 偶数长度的肯定不行(因为每次减少都是2)
7. 所以, 只有奇数长度的(有可能)变成1
8. 如果长度为11的能变成1, 那么长度为9的也可以吗? 不一定. 因为有可能在11的的地方有一个很大的数字
9. 也就是找到最长的, 奇数sum > 偶数sum的地方
10. 假设l...r (奇数长度满足这个条件)
11. a[l] + a[l+2] + .. + a[r] > a[l+1] + a[l+3] .. + a[r-1]
12. a[l] - a[l+1] + a[l+2] - a[l+3] ... + a[r-2] - a[r-1] + a[r] > 0
13. fix r 要计算有多少个这样的l
14.  let f[i] = f[i-2] + a[i] - a[i+1] (要区分奇偶性)
15.  s[l...r] = a[l] - a[l+1] + a[l+2] - a[l+3] ... + a[r-2] - a[r-1] + a[r]
16.   = f[r-2] - f[l-2] + a[r] > 0
17.   f[l-2] < f[r-2] + a[r] 满足这个的l的数量