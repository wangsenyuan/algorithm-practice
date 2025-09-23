# Problem D: Linear Function Composition

## Problem Statement

Consider a linear function $f(x) = Ax + B$. Let's define:
- $g^{(0)}(x) = x$
- $g^{(n)}(x) = f(g^{(n-1)}(x))$ for $n > 0$

For the given integer values $A$, $B$, $n$ and $x$, find the value of $g^{(n)}(x) \bmod (10^9 + 7)$.

## Input

The only line contains four integers $A$, $B$, $n$ and $x$ ($1 \leq A, B, x \leq 10^9$, $1 \leq n \leq 10^{18}$) — the parameters from the problem statement.

**Note:** The given value $n$ can be too large, so you should use 64-bit integer type to store it. In C++ you can use the `long long` integer type and in Java you can use `long` integer type.

## Output

Print the only integer $s$ — the value $g^{(n)}(x) \bmod (10^9 + 7)$.

## Examples

### Example 1
**Input:**
```
3 4 1 1
```

**Output:**
```
7
```

### Example 2
**Input:**
```
3 4 2 1
```

**Output:**
```
25
```

### Example 3
**Input:**
```
3 4 3 1
```

**Output:**
```
79
```


### ideas
1. 考虑一般的形式, $g_0(x)$ = x
2. $g_1(x)$ = $f(g_0(x))$ = $f(x)$ = $A*x + B$
3. $g_2(x)$ = $f(g_1)x))$ = $f(A*x + B)$ = $A * (A * x + B) + B = A * A * x + A * B + B$ = $A^2*x + A*B + B$
4. $g_3(x)$ = $A^3 * x + A^2*B + A * B + B$
5. $g_i(x)$ = $A^i * x + A^{i-1}*B + .... + A*B + B$
6. 后半部分是个等比数列 = $B * (1 + A + A^2 + ... + A^i)$
  
###
我来为您解释等比数列的求和公式。

等比数列是指每一项与前一项的比值都相等的数列。设首项为 $a_1$，公比为 $r$，则等比数列的通项公式为：
$$a_n = a_1 \cdot r^{n-1}$$

等比数列的求和公式有两种情况：

## 1. 当 $r \neq 1$ 时：
$$S_n = a_1 \cdot \frac{1-r^n}{1-r} = a_1 \cdot \frac{r^n-1}{r-1}$$

## 2. 当 $r = 1$ 时：
$$S_n = n \cdot a_1$$

## 无穷等比数列求和（当 $|r| < 1$ 时）：
$$S_{\infty} = \frac{a_1}{1-r}$$

## 推导过程：

设等比数列的前 $n$ 项和为 $S_n$：
$$S_n = a_1 + a_1r + a_1r^2 + \cdots + a_1r^{n-1}$$

两边同时乘以公比 $r$：
$$rS_n = a_1r + a_1r^2 + a_1r^3 + \cdots + a_1r^n$$

两式相减：
$$S_n - rS_n = a_1 - a_1r^n$$
$$S_n(1-r) = a_1(1-r^n)$$

当 $r \neq 1$ 时：
$$S_n = a_1 \cdot \frac{1-r^n}{1-r}$$

## 示例：
- 数列：2, 4, 8, 16, 32（首项 $a_1=2$，公比 $r=2$）
- 前5项和：$S_5 = 2 \cdot \frac{1-2^5}{1-2} = 2 \cdot \frac{1-32}{-1} = 2 \cdot 31 = 62$

这个公式在数学、物理、工程等领域都有广泛应用，特别是在计算复利、人口增长模型等方面。