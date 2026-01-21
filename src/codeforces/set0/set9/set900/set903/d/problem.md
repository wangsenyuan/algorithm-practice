Let's denote a function \(d(x, y)\).

You are given an array \(a\) consisting of \(n\) integers. You have to calculate the sum of \(d(a_i, a_j)\) over all pairs \((i, j)\) such that \(1 \le i \le j \le n\).

## Input

The first line contains one integer \(n\) \((1 \le n \le 200000)\) — the number of elements in \(a\).

The second line contains \(n\) integers \(a_1, a_2, \dots, a_n\) \((1 \le a_i \le 10^9)\) — elements of the array.

## Output

Print one integer — the sum of \(d(a_i, a_j)\) over all pairs \((i, j)\) such that \(1 \le i \le j \le n\).

## Examples

**Example 1**

Input:
```
5
1 2 3 1 3
```

Output:
```
4
```

**Example 2**

Input:
```
4
6 6 5 5
```

Output:
```
0
```

**Example 3**

Input:
```
4
6 6 4 4
```

Output:
```
-8
```

## Note

In the first example:

- \(d(a_1, a_2) = 0\);
- \(d(a_1, a_3) = 2\);
- \(d(a_1, a_4) = 0\);
- \(d(a_1, a_5) = 2\);
- \(d(a_2, a_3) = 0\);
- \(d(a_2, a_4) = 0\);
- \(d(a_2, a_5) = 0\);
- \(d(a_3, a_4) = -2\);
- \(d(a_3, a_5) = 0\);
- \(d(a_4, a_5) = 2\).