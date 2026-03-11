# Problem

**题意简述**：给定两个 01 串 a、b。每次可任选 a 中相邻两个字符 x、y，将二者替换为 (x xor y, x or y) 或 (x or y, x xor y)。问能否通过若干次（可为 0 次）操作把 a 变成 b。

## Input

- The first line contains Bitlandish string a, the second line contains Bitlandish string b. The strings can have different lengths.
- It is guaranteed that the given strings only consist of characters "0" and "1". The strings are not empty, their length doesn't exceed 10^6.

## Output

Print "YES" if a can be transformed into b, otherwise print "NO". Please do not print the quotes.

## Examples

### Example 1

**Input**

```text
11
10
```

**Output**

```text
YES
```

### Example 2

**Input**

```text
1
01
```

**Output**

```text
NO
```

### Example 3

**Input**

```text
000
101
```

**Output**

```text
NO
```


### ideas
1. 11 => 10 or 01
2. (10, 01) => 11
3. 00 => 00
4. 如果有一个1, 那么可以把所有的变成1
5. b中，如果没有相邻的0，那么这样就ok了。但是如果存在相邻的0，就有点麻烦了
6. 10 =》 11 =》 01， 相当于移动1的位置
7. 而且，1的个数也可以增加（减少），移动到连续的地方，就可以变少