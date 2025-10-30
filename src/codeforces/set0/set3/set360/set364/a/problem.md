### Problem: Count Rectangles in Product Matrix

You are given a string of decimal digits `s`. Define a matrix `b` where `b[i][j] = s[i] * s[j]` (digit multiplication). Count the number of axis-aligned rectangles such that the sum of `b[i][j]` over all cells `(i, j)` inside the rectangle equals `a`.

### Rectangle Definition
A rectangle is a tuple `(x, y, z, t)` with `x ≤ y` and `z ≤ t`. Its elements are all cells `(i, j)` such that `x ≤ i ≤ y` and `z ≤ j ≤ t`.

### Input
- The first line contains an integer `a` (`0 ≤ a ≤ 10^9`).
- The second line contains a string `s` of decimal digits (`1 ≤ |s| ≤ 4000`).

### Output
Print a single integer — the number of rectangles whose cell-sum equals `a`.

### Notes
The original statement mentions C++ I/O specifiers; this is irrelevant outside C++.

### Examples

Input
```text
10
12345
```
Output
```text
6
```

Input
```text
16
439873893693495623498263984765
```
Output
```text
40
```


### ideas
1. 读起来像谜语一样
2. 对于第一个例子
```
   1 2 3 4 5
   2 4 6 8 10
   3 6 9 12 15
   4 8 12 16 20
   5 10 15 20 25
```
3. 2 + 2 + 2 = 6
4. si * sj <= 81 >= 0
5. 对于一个矩形(r1, c1, r2, c2)
6. = sum(r1...r2) * sum(c1 ... c2)
7. 貌似是成立的
8. 那么就好办了