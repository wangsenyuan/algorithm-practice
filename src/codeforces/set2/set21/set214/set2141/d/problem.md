# Problem

You are given an integer array `a` with `n` elements. You need to make all elements equal.

You may perform the following operation at most `k` times:

- choose any index `i` (`1 <= i <= n`) and increase `a[i]` by `1`.

If the chosen element `a[i]` is strictly greater than the current minimum element of the array before the increment, you earn one coin.

Find the maximum number of coins you can earn while making all array elements equal.

## Input

The first line contains integer `t` (`1 <= t <= 1000`) — the number of test cases.

For each test case:

- First line: two integers `n` and `k` (`2 <= n <= 3 * 10^5`, `1 <= k <= 10^12`) — array length and maximum operations.
- Second line: `n` integers `a1, a2, ..., an` (`1 <= ai <= 10^9`) — the array.

It is guaranteed that the sum of `n` over all test cases does not exceed `3 * 10^5`.

## Output

For each test case:

- print `-1` if it is impossible to make all elements equal within at most `k` operations;
- otherwise, print the maximum possible number of coins.

## Example

**Input**

```text
4
3 16
1 10 2
4 20
6 2 4 9
5 9
7 7 7 7 7
2 1000000000000
1000000000 1000000000
```

**Output**

```text
-1
11
0
499999999999
```

## Note

In the first test case, at least `17` operations are needed to make all elements equal (for example, to `[10, 10, 10]`), so the answer is `-1`.

In the second test case, one optimal sequence is:

- increase `a3 = 4` to `10`,
- increase `a1 = 6` to `10`,
- increase `a4 = 9` to `10`,
- increase `a2 = 2` to `10`.

This uses `19` operations and earns `(10 - 4) + (10 - 6) + (10 - 9) = 11` coins, because increments on the current minimum element (`a2` here) do not give coins.

In the third test case, you can keep the array unchanged or make all elements `8`; both give `0` coins.


### ideas
1. 这个题目咋像谜语一样呢？
2. 假设目前的最小值是w
3. 操作次数有限制，最多k次
4. 而且最后需要大家一样，那么就有 n * avg <= sum + k
5. 如果avg < max(a) => -1, avg >= max(a)
6.  