# Problem

**题意：** 给定非负整数 $a$, $b$。定义模方程 $a \bmod x = b$，其中 $x$ 为正整数变量，$a \bmod x$ 表示 $a$ 除以 $x$ 的余数。求该方程有多少个正整数解 $x$。

## Input

一行两个空格分隔的整数 $a$, $b$（$0 \le a, b \le 10^9$）。

## Output

若解有无穷多个，输出 `infinity`（不含引号）；否则输出解的个数。

## Examples

### Example 1

**Input:**

```
21 5
```

**Output:**

```
2
```

### Example 2

**Input:**

```
9435152 272
```

**Output:**

```
282
```

### Example 3

**Input:**

```
10 10
```

**Output:**

```
infinity
```

## Note

第一组样例中，$21 \bmod x = 5$ 的解为 $x = 8$ 和 $x = 16$。


## ideas
1. 如果 a = b => inf
2. a % x = b, 如果 a < b => 0
3. a > b, a % x = b => a = x * n + b => (a - b) % x = 0
4. 那么就是 a - b 的因数，且 > b 的部分