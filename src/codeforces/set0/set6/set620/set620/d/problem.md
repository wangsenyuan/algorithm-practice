# Problem

Professor GukiZ has two integer arrays `a` and `b`. Let `sa` and `sb` be the sums
of `a` and `b`. He wants to minimize `v = |sa - sb|` by swapping elements
between the arrays.

In a single operation he may swap any element of `a` with any element of `b`.
After each swap the arrays are updated, and he will perform at most two swaps in
total. Your task is to find the minimal possible `v` and describe up to two
swaps that achieve it (in the order performed).

## Input

1. `n` — number of elements in `a` (`1 ≤ n ≤ 2000`)
2. `n` integers `ai` (`-10^9 ≤ ai ≤ 10^9`)
3. `m` — number of elements in `b` (`1 ≤ m ≤ 2000`)
4. `m` integers `bj` (`-10^9 ≤ bj ≤ 10^9`)

## Output

1. Minimal value `v = |sa - sb|` achievable with at most two swaps.
2. `k` — number of swaps used (`0 ≤ k ≤ 2`).
3. `k` lines: each contains indices `xp` and `yp` (`1 ≤ xp ≤ n`, `1 ≤ yp ≤ m`)
   describing the positions swapped during the `p`‑th operation.

If multiple optimal answers exist, print any of them.

## Examples

### Example 1
```
Input
5
5 4 3 2 1
4
1 1 1 1

Output
1
2
1 1
4 2
```

### Example 2
```
Input
5
1 2 3 4 5
1
15

Output
0
0
```

### Example 3
```
Input
5
1 2 3 4 5
4
1 2 3 4

Output
1
1
3 1
```


### ideas
1. 假设目标是使的 sum(b) >= sum(a)
2. 最多交换两个位置(i1, j1), (i2, j2)
3. s1 = sum(a) - a[i1] + b[j1] - a[i2] + b[j2]
4. s2 = sum(b) + a[i1] - b[j1] + a[i2] - b[j2]
5. s2 - s1 = sum(b) - sum(a) + 2 * a[i1] + 2 * a[i2] - 2 * b[j1] - 2 * b[j2]
6. 目标是使的 s2 - s1 最小 >= 0
7. 在选定i1, i2 的时候， 那么就是找b[j1] + b[j2]的某个值