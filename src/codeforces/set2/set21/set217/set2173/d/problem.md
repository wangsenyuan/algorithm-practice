# Problem

After classes at Ohashi High School, Ryuuji gives Taiga a positive integer `n` and a challenge.

They will play for exactly `k` moves. In one move, Taiga chooses a non-negative integer `l` and updates:

`n <- n + 2^l`.

The score of a move is the number of binary carries that occur when adding `2^l` to the current `n` in base 2.

The total score is the sum of move scores over all `k` moves.

Taiga wants to maximize the total score. Find the maximum possible total score.

## Input

Each test contains multiple test cases.

The first line contains integer `t` (`1 <= t <= 1000`) — the number of test cases.

Each test case contains one line with two integers `n` and `k` (`1 <= n < 2^30`, `0 <= k <= 10^9`) — the initial value and the number of moves.

## Output

For each test case, output one integer — the maximum total score Taiga can achieve.

## Example

**Input**

```text
6
7 1
13 2
42 2
1048576 100
23 2
371 1
```

**Output**

```text
3
4
3
100
5
3
```

## Note

In the first test case, `(n, k) = (7, 1)` and `7 = 111_2`.
Adding `2^0` gives `111 + 1 = 1000`, which produces carries at bits `0`, `1`, and `2`.
So the total score is `3`.

In the second test case, `(n, k) = (13, 2)` and `13 = 1101_2`.
First add `2^0`: `1101 + 0001 = 1110` (one carry).
Then add `2^1`: `1110 + 0010 = 10000` (carries through bits `1`, `2`, `3`).
Total score is `1 + 3 = 4`.

In the third test case, `(n, k) = (42, 2)` and `42 = 101010_2`.
First add `2^1`: `101010 + 000010 = 101100` (one carry).
Then add `2^2`: `101100 + 000100 = 110000` (carries at bits `2` and `3`).
Total score is `1 + 2 = 3`.

In the fifth test case, `(n, k) = (23, 2)` and `23 = 10111_2`.
First add `2^0`: `10111 + 00001 = 11000` (carries at bits `0`, `1`, `2`).
Then add `2^3`: `11000 + 01000 = 100000` (carries at bits `3`, `4`).
Total score is `3 + 2 = 5`.


### ideas
1. 考虑n是1<<m, 那么此时l = m, 进位一次，且从次之后，是不是一直都只能进位一次呢？
2. 假设 n = 4, 还能操作2次, 那么按照上面的操作，得分 = 2
3. 如果 n = 4, + 2, (得分0)， + 2 (得分2)，分数还是2
4. 考虑 n = 0001101100
5. 可以操作2次, 那么先填中间的0, 编程 00011111000 然后，再最后1的位置添加 => 5
6. 如果先填最后一位, 得到 0001110000 (得分 2)， 然后，再处理一次, 0010.. => 2 + 3 = 5
7. 如果k大于中间的0的个数 => k - cnt0 + （第一个1和最后一个1之间的距离）？
8. 如果 k < ... 那么需要特殊处理？
9. 如果是这样，是不是太简单了？