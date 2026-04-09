# Statement

You are given a grid `a` of **2** rows and **`n`** columns, where every cell has a value in **`1 .. 2n`**.

For integers `(l, r)` with `1 <= l <= r <= 2n`, define **`f(l, r)`** as a **binary** grid `b` of 2 rows and `n` columns such that:

- `b[i][j] = 1` if and only if `l <= a[i][j] <= r`,
- otherwise `b[i][j] = 0`.

Cell `(i, j)` is **`i`** rows from the top and **`j`** columns from the left.

Count the number of pairs **`(l, r)`** with `1 <= l <= r <= 2n` such that in **`f(l, r)`** there exists a **down-right path** of adjacent cells with value **1** from **`(1, 1)`** to **`(2, n)`**.

---

**Binary grid.** A grid is binary if and only if every cell is **0** or **1**.

**Down-right path.** A sequence of cells `(c1, c2, ...)` such that for each `k > 1`, cell `ck` shares either **its top side** or **its left side** with **some side of** `c(k-1)`. Equivalently, each step goes **one cell down** or **one cell right** (from `(1,1)` toward `(2,n)` in a 2×`n` grid).

---

## Input

Each test contains multiple test cases. The first line contains **`t`** (`1 <= t <= 10^4`).

For each test case:

- Line 1: integer **`n`** (`2 <= n <= 2 * 10^5`) — number of columns.
- Line 2: **`n`** integers `a[1][1], ..., a[1][n]` (`1 <= a[1][i] <= 2n`) — first row.
- Line 3: **`n`** integers `a[2][1], ..., a[2][n]` (`1 <= a[2][i] <= 2n`) — second row.

The sum of **`n`** over all test cases does not exceed **`2 * 10^5`**.

## Output

For each test case, print one integer: the number of pairs **`(l, r)`** with `1 <= l <= r <= 2n` such that **`f(l, r)`** admits a down-right path of **1**-cells from **`(1, 1)`** to **`(2, n)`**.

## Example

### Input

```text
5
2
1 3
3 1
3
1 2 3
3 2 1
4
1 5 5 5
5 3 1 2
4
8 8 8 8
8 8 8 8
6
6 6 5 7 9 12
1 4 2 8 5 6
```

### Output

```text
2
5
4
8
25
```

### ideas
1. 也就是说始终要存在一条边，从左上到右下的路径
2. 假设在第i列转向到下面一行，那么可以知道这个时候，需要cover的所有的数字的范围
3. 假设这个时候的最小值是a，最大值是b, 那么所有l <= a, b <= r的pair都可以满足条件
4. 那么这样子可以得到n对(a, b), 从而对它们进行排序；
5. 然后对这些pair进行处理，如果pair 1包含pair 2， 那么可以直接舍弃pair1, 
6. 因为满足1的条件，也满足了2的条件，反过来不成立
7. 这样子，就得到了一些相交但是不包含的区间; 处理逻辑是一样的