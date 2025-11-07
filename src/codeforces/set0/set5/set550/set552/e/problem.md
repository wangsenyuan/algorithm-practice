# Vanya and Brackets

Vanya is doing his maths homework. He has an expression of the form `x1 op1 x2 op2 ... opn-1 xn`, where `x1, x2, ..., xn` are digits from 1 to 9, and each operator `op` is either a plus `+` or multiplication `*`. Vanya needs to add exactly one pair of brackets to maximize the value of the resulting expression.

## Input

The first line contains expression `s` (`1 ≤ |s| ≤ 5001`, `|s|` is odd). Odd positions contain digits from 1 to 9, and even positions contain signs `+` or `*`.

The number of multiplication signs `*` doesn't exceed 15.

## Output

Print the maximum possible value of the expression.

## Examples

```
Input
3+5*7+8*4
Output
303
```

```
Input
2+3*5
Output
25
```

```
Input
3*4*5
Output
60
```

## Note

In the first test: `3 + 5 * (7 + 8) * 4 = 303`.

In the second test: `(2 + 3) * 5 = 25`.

In the third test: `(3 * 4) * 5 = 60` (many other placements are valid, for instance `(3) * 4 * 5 = 60`).


### ideas
1. 有意思
2. 如果括号加在为止i, j两边，那么需要计算出i前面的值, j前面的值，
3. 还有i,j中间的值， 然后再计算值
4. 好像是对的，两个下标，貌似也是 n * n
5. dp[i] 表示在i前面的结果（不考虑i后面的操作符）
6. 如果i前面是+， dp[i] = dp[i-1] + x
7. 如果i前面是*， dp[i] = prod[?...i] + dp[?-1]
8. *号不多，所以可以bruteforce，去找前面的+