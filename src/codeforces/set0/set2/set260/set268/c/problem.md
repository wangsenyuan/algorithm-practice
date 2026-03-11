# Problem

**题意简述**：在平面上考虑所有整点 `(x, y)`，满足 `0 ≤ x ≤ n`, `0 ≤ y ≤ m`, 且 `x + y > 0`。我们要选出一个最大的子集，使得任意两点间的欧氏距离都不是整数，称为“beautiful set”。输出这个子集的最大大小以及任意一种构造。


## Input

The single line contains two space-separated integers `n` and `m` (`1 ≤ n, m ≤ 100`).

## Output

In the first line print a single integer — the size `k` of the found beautiful set. In each of the next `k` lines print a pair of space-separated integers — the `x`- and `y`-coordinates, respectively, of a point from the set.

If there are several optimal solutions, you may print any of them.

## Examples

### Example 1

**Input**

```text
2 2
```

**Output**

```text
3
0 1
1 2
2 0
```

### Example 2

**Input**

```text
4 3
```

**Output**

```text
4
0 3
2 1
3 0
4 2
```


### ideas
1. 同一行，同一列，只能有一个点被选中
2. 应该能选中min(n, m) + 1 个点
3. 那么就是一个*皇后*放置问题？与其去检查*整数*距离，不如直接限定，不能在同一条斜线上？
4. 对于(n, m)比较小的情况，用brute force
5. 不行～
6. 