# Problem C - Vasya and Beautiful Arrays

## Problem Statement

Vasya's got a birthday coming up and his mom decided to give him an array of positive integers `a` of length `n`.

Vasya thinks that an array's beauty is the greatest common divisor of all its elements. His mom, of course, wants to give him as beautiful an array as possible (with largest possible beauty). Unfortunately, the shop has only one array `a` left. On the plus side, the seller said that he could decrease some numbers in the array (no more than by `k` for each number).

The seller can obtain array `b` from array `a` if the following conditions hold:
- `bi > 0`
- `0 ≤ ai - bi ≤ k` for all `1 ≤ i ≤ n`

Help mom find the maximum possible beauty of the array she will give to Vasya (that seller can obtain).

## Input

The first line contains two integers `n` and `k` (`1 ≤ n ≤ 3·10⁵`; `1 ≤ k ≤ 10⁶`). The second line contains `n` integers `ai` (`1 ≤ ai ≤ 10⁶`) — array `a`.

## Output

In the single line print a single number — the maximum possible beauty of the resulting array.

## Examples

### Example 1
**Input:**
```
6 1
3 6 10 12 13 16
```

**Output:**
```
3
```

### Example 2
**Input:**
```
5 3
8 21 52 15 77
```

**Output:**
```
7
```

## Note

In the first sample we can obtain the array:
```
3 6 9 12 12 15
```

In the second sample we can obtain the next array:
```
7 21 49 14 77
```


### ideas
1. 假设最后的答案是c, (a[i] - b[i]) <= k, 且 b[i] % c == 0
2. let d[i] = a[i] - b[i] (d[i]是减少的数量) d[i] <= k
3. 那么 a[i] % c = d[i] % c
4. d[i] = 0, 1, 2, .... k
5. 假设到i为止，可以得到x, 那么就可以得到x的除数的部分, 比如 y * 2 = x ....
6. 所以，理论上x是在递减的，随着i的增加
7. 先按倒序排列
8. 似乎也不符合2分，比如例子2，是无法得到g = 6的情况的
9. 