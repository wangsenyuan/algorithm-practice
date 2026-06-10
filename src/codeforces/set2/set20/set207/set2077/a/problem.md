# A. Breach of Faith

[Problem link](https://codeforces.com/problemset/problem/2077/A)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

You and your team found a sequence `a_1, a_2, ..., a_{2n+1}` of positive integers
with these properties:

- `1 <= a_i <= 10^18` for all `1 <= i <= 2n + 1`
- all `2n + 1` values are pairwise distinct
- `a_1 = a_2 - a_3 + a_4 - a_5 + ... + a_{2n} - a_{2n+1}`

A collaborator deleted one element from `a` and shuffled the remaining `2n` values,
leaving you with sequence `b_1, b_2, ..., b_{2n}`. You forgot the original `a`.

Recover any valid sequence `a`. It is guaranteed that at least one exists.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`).

The second line contains `2n` distinct integers `b_1, b_2, ..., b_{2n}`
(`1 <= b_i <= 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `2n + 1` distinct integers — a valid sequence `a`
(`1 <= a_i <= 10^18`).

If multiple answers exist, output any of them. The sequence must satisfy the
conditions above, and `b` must be obtainable by deleting one element from `a` and
permuting the rest.

## Example

### Input

```text
4
2
1 8 4 6
3
9 7 2 1 8 4
4
86 33 14 77 21 6 3 2
1
6 2
```

### Output

```text
8 1 6 17 4
9 2 8 1 7 25 4
86 6 77 3 33 2 21 220 14
6 8 2
```

## Note

In the first test case, one valid sequence is `a = [8, 1, 6, 17, 4]`.

Check:

```text
a_1 = a_2 - a_3 + a_4 - a_5
8  = 1 - 6 + 17 - 4
```

If `17` is deleted, the remaining values `{8, 1, 6, 4}` can be rearranged into
`b = [1, 8, 4, 6]`.

In the last test case, `a = [6, 8, 2]` works because `6 = 8 - 2`, and deleting `8`
leaves `{6, 2}`.

Any valid output is accepted if it satisfies the constraints.


### ideas
1. a1 = a2 - a3 + a4 - a5 ... 
2. 把奇数移动到左边,满足
3. a1 + a3 + .. + a[2 * n + 1] = a2 + a4 + ... + a[2n]
4. a1 + a2 + ... + a[2 *n + 1] = sum
5. sum/2 = 两边的和
6. 假设缺少的是x, 且它在奇数位置上
7. s1 + x = sum 
8. 找到一个和s1同奇偶性的数, 就能满足sum 的条件
9. 所有的数不一样,为了让两边和一样,必须交替分配才行
10. 

## Solution Summary

The condition

```text
a_1 = a_2 - a_3 + a_4 - ... + a_{2n} - a_{2n+1}
```

is equivalent to saying that the sum of values placed at odd positions equals the
sum of values placed at even positions:

```text
a_1 + a_3 + ... + a_{2n+1} = a_2 + a_4 + ... + a_{2n}
```

So the task is to insert one new positive value `x` into the given `2n` values
and order all `2n + 1` values so the alternating sums are equal.

The solution sorts `b` and first tries the simplest construction: keep the sorted
array and append `x` at the end. If

```text
x = sum(values at 1-based even positions) - sum(values at 1-based odd positions)
```

is positive and not already present in `b`, then appending `x` produces a valid
answer.

If that candidate is invalid, the solution scans possible insertion positions
from right to left. For each position `i`, it computes how the alternating-sum
difference would change if a new value `x` were inserted before `b[i]`. Elements
to the right of `i` flip parity after insertion, so the code maintains their
signed contribution in `suf` and combines it with prefix parity sums. This gives
the only value of `x` that would make the final alternating sums equal. Once this
`x` is positive and distinct from all values in `b`, the solution inserts it at
that position and returns the result.

The map `nums` is used only to ensure the generated `x` is not one of the given
values, preserving the pairwise distinct requirement. The fallback inserts the
computed value at the front; the problem guarantees that a valid answer exists.

Sorting dominates the runtime, so each test case runs in `O(n log n)` time and
uses `O(n)` extra memory.
