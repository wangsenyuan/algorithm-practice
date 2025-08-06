# Problem Description

## Story Context

Maxim loves to fill in a matrix in a special manner. Here is a pseudocode of filling in a matrix of size $(m + 1) \times (m + 1)$:

```pseudocode
for (j = 1; j <= m; j = j + 1) {
    a[1][j] = 0
}
a[1][m + 1] = 1

for (i = 2; i <= m + 1; i = i + 1) {
    for (j = 1; j <= m + 1; j = j + 1) {
        if (j equal 1) {
            a[i][j] = a[i - 1][j + 1]
        } else {
            if (j equal m + 1) {
                a[i][j] = a[i - 1][j - 1]
            } else {
                a[i][j] = a[i - 1][j - 1] xor a[i - 1][j + 1]
            }
        }
    }
}
```

## Problem Statement

Maxim asks you to count, how many numbers $m$ ($1 \leq m \leq n$) are there, such that the sum of values in the cells in the row number $m + 1$ of the resulting matrix equals $t$.

## Key Definitions

- **XOR Operation**: $(x \text{ xor } y)$ means applying the operation of bitwise excluding "OR" to numbers $x$ and $y$. The given operation exists in all modern programming languages. For example, in languages C++ and Java it is represented by character `^`, in Pascal — by `xor`.

## Input

A single line contains two integers $n$ and $t$ ($1 \leq n, t \leq 10^{12}$, $t \leq n + 1$).

**Note**: Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Output

In a single line print a single integer — the answer to the problem.

## Examples

### Example 1

**Input:**
```
1 1
```

**Output:**
```
1
```

### Example 2

**Input:**
```
3 2
```

**Output:**
```
1
```

### Example 3

**Input:**
```
3 3
```

**Output:**
```
0
```

### Example 4

**Input:**
```
1000000000000 1048576
```

**Output:**
```
118606527258
```


### ideas
1. 第一行，除了最后一列，其他都是0，最后一个是1
2. 后面所有的格子，都是它左上角、右上角的 xor 值
3. 左上角的元素全部是0，主轴线的元素，都是1

```
00000000001
00000000010
00000000101
00000001000
00000010100
00000100010
00001010101
00010100000
00100010000
01010101000
10000000100
```

4. 