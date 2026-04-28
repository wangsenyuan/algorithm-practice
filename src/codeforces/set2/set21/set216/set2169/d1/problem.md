# Problem

Polycarp has a sequence of all natural numbers from `1` to `10^12`. He modifies this sequence by performing the following action `x` times:

Simultaneously remove all numbers at positions `y`, `2 * y`, `3 * y`, ..., `m * y <= n`, where `n` is the length of the current sequence.

After that, Polycarp wants to find the `k`-th number in the remaining sequence, or determine that the length of the resulting sequence is less than `k`.

Help Polycarp solve this problem.

## Input

Each test contains multiple test cases.

The first line contains integer `t` (`1 <= t <= 10`) — the number of test cases.

Each test case contains one line with three integers `x`, `y`, and `k` (`1 <= x <= 10^5`, `1 <= y, k <= 10^12`).

## Output

For each test case, output a positive integer — the value at the `k`-th position in the resulting sequence, or `-1` if the length of the resulting sequence is less than `k`.

## Note

In the statement figure, numbers removed in the first operation are marked in one color and those removed in the second in another. For `x = 2`, `y = 3`, `k = 5`, the `k`-th remaining number is `10`.

## Example

**Input**

```text
6
2 3 5
2 5 1
20 2 1000000000000
175 10 28
100000 998244353 1999999999
1 1 1
```

**Output**

```text
10
1
-1
2339030304
2000199999
-1
```

### ideas

1. 考虑一次操作，删除第y个， 2*y个，3*y个。考虑一个数，在什么时候被删除掉？
2. 假设w第一次被删除掉了(w = y, 2y, 3y, ...)
3. 那么下一个被删除的数 v1 = w1 + 1, v2 = w2 + 2, v3 = w3 + 3, ...
4. 完全看不出来规律～～～～
5. 肯定是二分，检查某个数w，是否能在x次后被保留下来？不对，它保留下来了，它右边的数不一定保留下来
6. 考虑k处的数，一开始 a[k] = k 的
7. 一次操作后, a[k] = k + k / y (这个也不行，因为有些数字也会被删除掉)
8. 

## Extra Sample

For `x = 4`, `y = 5`, mark crossed-out numbers as `~~v~~`:

[0] `1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, ..., 60`

[1] `1, 2, 3, 4, ~~5~~, 6, 7, 8, 9, ~~10~~, 11, 12, 13, 14, ~~15~~, 16, 17, 18, 19, ~~20~~, ..., ~~60~~`

[2] `1, 2, 3, 4, ~~5~~, ~~6~~, 7, 8, 9, ~~10~~, 11, ~~12~~, 13, 14, ~~15~~, 16, 17, ~~18~~, 19, ~~20~~, ..., ~~56~~, ..., ~~60~~`

[3] `1, 2, 3, 4, ~~5~~, ~~6~~, ~~7~~, 8, 9, ~~10~~, 11, ~~12~~, 13, ~~14~~, ~~15~~, 16, 17, ~~18~~, 19, ~~20~~, ..., ~~53~~, ..., ~~60~~`

[4] `1, 2, 3, 4, ~~5~~, ~~6~~, ~~7~~, ~~8~~, 9, ~~10~~, 11, ~~12~~, 13, ~~14~~, ~~15~~, 16, ~~17~~, ~~18~~, 19, ~~20~~, ..., ~~57~~, ..., ~~60~~`

After `[4]`, the first 20 numbers in the remaining sequence are:

```text
1 2 3 4 9 11 13 16 19 21 23 26 28 32 33 34 39 41 42 44
```
