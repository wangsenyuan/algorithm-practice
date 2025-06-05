# CodeForces 2114F - Problem Statement

## Problem Description

Given an integer `x` and an integer `k`. In one operation, you can perform one of two actions:

1. Choose an integer `1 ≤ a ≤ k` and assign `x = x · a`
2. Choose an integer `1 ≤ a ≤ k` and assign `x = x / a`, where the value of `x / a` must be an integer

Find the minimum number of operations required to make the number `x` equal to `y`, or determine that it is impossible.

## Input

The first line of the input contains one integer `t` (`1 ≤ t ≤ 10⁴`) — the number of test cases.

The only line of each test case contains three integers `x`, `y` and `k` (`1 ≤ x, y, k ≤ 10⁶`).

It is guaranteed that the sum of `x` and the sum of `y` across all test cases does not exceed `10⁸`.

## Output

For each test case, output `-1` if it is impossible to achieve `x = y` using the given operations, and the minimum number of required operations otherwise.

## Operations Summary

- **Multiplication**: `x = x · a` where `1 ≤ a ≤ k`
- **Division**: `x = x / a` where `1 ≤ a ≤ k` and `x / a` is an integer

## Constraints

- `1 ≤ t ≤ 10⁴`
- `1 ≤ x, y, k ≤ 10⁶`
- Sum of all `x` values ≤ `10⁸`
- Sum of all `y` values ≤ `10⁸`

## ideas
1. 乘法感觉没有啥限制，理论上可以一直乘（但是乘的太多，肯定有问题）
2. 如果对a的上限没有限制，那么可以先将x变成lcm(x, y), 然后将它变成y
3. 变成gcd(x, y), 然后再乘上来，可行吗？
4. 好像不一定，比如 6 / 3 * 2 是合理的
5. 6 / 2 * 3 也是合理的。
6. 上限是lcm(x, y)?
7. 如果 x能仅通过除法就能得到y，那么用任何a去乘x，都是不好的
8. 如果x是y的倍数，且这个倍数，可以用不超过k的序列，相乘组合出来，那么就应该直接这样操作
9. 否则x应该变大a。但是这个a，肯定不能在后面被除掉。 假设乘以2，得到x2, 那么x2 作为y的倍数，不能有2作为因子
10. x2 % (2 * y) != 0 才行;
11. 但是直接乘以2似乎也不对。那么只有当x没有2（或者2不够），但是y中有2的时候，才需要乘以2
12. 质数因子
13. 假设`x =` $p1 ^ {e1} * p2 ^ {e2}$, `y =` $p1 ^ {f1} * p2 ^ {f2}$
14. 如果e1 < f1， 那么必须有一个步骤要将e1 变成f1（乘法）
15. 如果e2 > f1， 那么必须有一个步骤要将e2 变成f2 (除法)
16. 如果不考虑最小操作数，那么就可以直接这样sum(abs(f1 - e1), abs(f2 - e2))
17. 最小操作数，就是要把一些乘法合并在一起,得到一个不超过k的数a
18. 这个是不是要dp才行？还是贪心？
19. 这时候用bfs，似乎是可以的。因为乘法后的结果，不会超过y（先尽量做除法）
20. 