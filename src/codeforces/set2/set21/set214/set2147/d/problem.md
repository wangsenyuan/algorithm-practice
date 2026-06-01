# D. Game on Array

[Problem link](https://codeforces.com/problemset/problem/2147/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

You are given an array `a` of `n` positive integers. Alice and Bob will play a game
with this array. They will take alternating turns, with Alice going first.

At each turn, the player must choose a value `x > 0` that appears in `a` at least
once. Then:

1. the player earns 1 point for each value `x` in the array,
2. each value `x` in the array is decreased by 1 and becomes `x - 1`.

Note that the player can choose `x` only if it is present in `a` at the moment,
so each valid move earns a positive amount of points. For example, if the array
is `[3, 8, 5, 8]` and Alice chooses `x = 8`, the array will become `[3, 7, 5, 7]`
and Alice will earn 2 points.

The game ends when no `x` can be chosen; that is, when all the elements in the
array are zero.

Given that both players want to maximize their points and play optimally, calculate
the amount of points that each player will end up with.

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 <= t <= 10^3`). The description of the test cases follows.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`) —
the size of the array.

The second line contains `n` integers `a1, a2, ..., an` (`1 <= ai <= 10^9`) — the
elements of the array.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output two integers, the amount of points Alice and Bob will
get if both play optimally.

## Example

### Input

```text
3
2 1 1
5
3 3 3 5 5
4
9 9 9 9
```

### Output

```text
3 1
10 9
20 16
```

## Note

In the first test case, Alice chooses `x = 1` with her first move. The array
becomes `[2, 0, 0]` and she earns 2 points. After that, Bob is forced to choose
`x = 2`. The array becomes `[1, 0, 0]` and Bob earns 1 point. Finally, Alice is
forced to choose `x = 1` and earn one more point. After that, the array is
`[0, 0, 0]`, so the game ends. In total, Alice finishes with 3 points and Bob
with 1 point.

In the third test case, each player will decrement all the elements and earn 4
points each move. Alice will earn `5 * 4 = 20` points while Bob will only earn
`4 * 4 = 16` points.


### ideas
1. 总的分数 = sum(arr) 
2. 所以只需要计算alice的得分就可以了
3. 如果所有的数字都一样，那么alice的得分 = (x + 1) / 2 * n
4. 对于alice来说，应该让bob没次得分越少越好，所以他的策略应该是从小开始，
5. 但是bob应该从大开始，这样子，可以产生更多相同的数。
6. 但是似乎总是有一个状态bob能够促使alice不得不配合自己？