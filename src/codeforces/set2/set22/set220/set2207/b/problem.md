# Codeforces 2207B - One Night At Freddy's

https://codeforces.com/problemset/problem/2207/B

## Statement

Let `n`, `m`, and `l` be positive integers.

You have made the unfortunate decision to work as a night guard at Freddy
Fazbear's Pizzeria, where there are `m` animatronics numbered `1..m` waiting to
jumpscare you.

The night consists of `l` seconds, numbered `1..l`. The `j`-th animatronic has a
danger level `dj`, and initially:

```text
d1 = d2 = ... = dm = 0
```

Every second, the danger level of exactly one animatronic increases by `1`.
Throughout the night, you can observe all current values of `dj`.

For example, if `m = 2`, one possible list of danger levels after `5` seconds is:

```text
[d1, d2] = [2, 3]
```

You are not defenseless. At each of the `n` fixed times `ai`
(`1 <= a1 < a2 < ... < an <= l`), you can shine your flashlight on exactly one
animatronic `ji` of your choice. This happens immediately after the `ai`-th
second and sets its danger level back to zero:

```text
d[ji] := 0
```

The choice is made independently for each flashlight use.

Continuing the example above, if `a1 = 5` and you choose to flash the second
animatronic at that time, then the danger levels after `5` seconds become:

```text
[d1, d2] = [2, 0]
```

The overall danger is the maximum danger across all animatronics:

```text
max(dj) for 1 <= j <= m
```

You lose if the overall danger at the end of the night, after `l` seconds, is
greater than `x`.

Find the minimum value of `x` such that, regardless of the animatronics'
actions, you can guarantee that the final overall danger is at most `x`.

## Input

Each test contains multiple test cases.

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

For each test case:

- The first line contains three integers `n`, `m`, and `l`
  (`1 <= n, m, l <= 2 * 10^5`, `n <= l`, `1 <= m * l <= 2 * 10^5`) — the number
  of flashlight actions, the number of animatronics, and the length of the
  night.
- The second line contains `n` integers `a1, a2, ..., an`
  (`1 <= a1 < a2 < ... < an <= l`) — the times at which you can shine the
  flashlight.

It is guaranteed that the sum of `m * l` over all test cases does not exceed
`2 * 10^5`.

## Output

For each test case, output a single integer: the minimum `x` such that you can
guarantee the final danger level is at most `x`.

## Sample

```text
7
1 2 10
10
5 1 32
1 4 9 16 25
2 3 40
13 37
2 2 7
6 7
8 5 60
3 17 20 28 36 44 45 50
6 7 1987
6 7 66 77 666 777
1 1 1
1
```

```text
5
7
19
1
19
1477
0
```

## Note

In the first test case, there are `2` animatronics and the night length is `10`.
You can flash after the `10`-th second. For `x = 5`, after `10` seconds, one
animatronic must have at least `5` danger and the other has at most `5` danger.
Flash the one with higher danger, and the final danger is at most `5`. It can be
shown that `5` is minimum.

In the second test case, there is only one animatronic and the night length is
`32`. The animatronic's danger increases by `1` every second. After the `25`-th
second, reset it to `0`; then `7` more seconds pass before the night ends, so the
final danger is always `7`.

In the third test case, it can be proven that the minimum possible value of `x`
is `19`.

### ideas
1. 假设a[m]的时候，将某个怪物照射了，那么在接下来的l - a[m]+1的时间内，就没有机会，那么就让它生长是最优的
2. x >= l - a[m] + 1
3. 然后考虑a[m-1], 如果它是最后一次（虽然不是），那么x >= l - a[m-1] + 1, 
4. 但是因为还有一次照射，最好的方案是各张一半，x/2, x - x/2， 然后多的那个被照射，剩下的（a[m] - a[m-1] + 1）/2 + l - a[m] + 1
5. 考虑m-2. 一样的逻辑？
6. 怪兽的数量没有考虑？

## Solution

At any flashlight time, the best choice is to reset an animatronic with maximum
current danger. Resetting a smaller danger while a larger one remains can only
make the final maximum worse.

Now suppose there are `k` flashlight uses remaining. Only the largest `k + 1`
danger levels matter:

- At most `k` animatronics can still be reset before the end.
- So even if we reset the largest `k` among the dangerous ones, the final answer
  is still controlled by the next one.
- Any animatronic outside the largest `k + 1` cannot become the final maximum
  unless it first enters this tracked set.

Therefore, while there are `k` flashes remaining, we maintain only:

```text
min(m, k + 1)
```

danger levels, sorted in descending order.

## Adversary move

Each second, the animatronics choose one danger level to increase by `1`.

To maximize the eventual final maximum among the tracked values, the adversary
should increase the smallest tracked value. This is the usual balancing
strategy: increasing a larger value can be erased by a future flashlight, while
raising the smallest tracked value increases the lower part of the top
`k + 1` values.

So every second:

```text
increase the last element of the descending sorted tracked list
sort descending again
```

The constraints allow this direct simulation, because:

```text
sum(m * l) <= 2 * 10^5
```

## Flashlight move

If a flashlight is available at the current second:

1. Remove the largest tracked danger level. This represents resetting the most
   dangerous animatronic to `0`.
2. Decrease the number of remaining flashes.
3. If the tracked list should still contain more values, append a `0`. This
   represents another currently harmless animatronic becoming relevant.
4. Sort descending again.

The target tracked size after the flash is:

```text
min(m, remaining_flashes + 1)
```

At the end of all `l` seconds, the answer is the largest value in the tracked
list.

## Complexity

The tracked list has length at most `m`. Each second sorts this list, so the
implementation is easily fast enough under the given constraint:

```text
O(l * m log m)
```

over each test case, with:

```text
sum(m * l) <= 2 * 10^5
```
