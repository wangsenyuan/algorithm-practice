# Problem

给定长度为 $n$ 的二进制数 $x$，定义排列 $P$：$P(i) = x \oplus i$，其中 $i \in [0, 2^n - 1]$。

求排列 $P$ 的**逆序对数量**，对 $10^9 + 7$ 取模。

其中 $\oplus$ 表示按位异或（XOR）。

## Input

一行，一个长度为 $n$ 的二进制串 $x$（$1 \le n \le 100$），可能有前导零。

## Output

逆序对数量对 $10^9 + 7$ 取模的结果。

## Examples

### Example 1

**Input:** `11`  
**Output:** `6`

### Example 2

**Input:** `01`  
**Output:** `2`

### Example 3

**Input:** `1`  
**Output:** `0`
