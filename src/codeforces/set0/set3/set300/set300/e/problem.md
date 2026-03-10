# Problem

给定 $k$ 个正整数 $a_1, \dots, a_k$，要求找出最小的正整数 $n$，使得 `n! / (a1 * a2 * ... * ak)` 是整数（即 `a1 * a2 * ... * ak` 整除 `n!`）。输出这样的最小 $n$。

## Input

- The first line contains an integer $k$ ($1 \le k \le 10^6$).
- The second line contains $k$ integers $a_1, a_2, \dots, a_k$ ($1 \le a_i \le 10^7$).

## Output

Print the minimum positive integer `n` such that `n! / (a1 * a2 * ... * ak)` is an integer.

## Examples

### Example 1

**Input**

```text
2
1000 1000
```

**Output**

```text
2000
```

### Example 2

**Input**

```text
1
2
```

**Output**

```text
2
```


