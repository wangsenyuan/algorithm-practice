### Problem

Vasya wrote an arithmetic expression `a+b=c` in his notebook.  
The teacher checked it and found it **incorrect**.

Vasya now claims that he simply forgot to write down some digits in the numbers `a`, `b`, and `c`.  
You need to help him find numbers `x`, `y`, and `z` such that:

1. `x + y = z`
2. From the correct expression `x+y=z` you can erase **some digits** in `x`, `y`, and `z` (possibly zero digits) so that the remaining characters form exactly the given expression `a+b=c` as a subsequence.
3. The resulting correct expression `x+y=z` has the **minimal possible total length** (minimal number of characters).

### Input

- A single line: the expression `a+b=c`  
  - `1 ≤ a, b, c ≤ 10^6`  
  - `a`, `b`, and `c` have **no leading zeros**.

### Output

Print a **correct** expression `x+y=z`, where:

- `x`, `y`, `z` are non-negative integers without leading zeros.
- The original expression `a+b=c` appears in `x+y=z` as a **subsequence** (after erasing some digits).
- The printed expression has the **minimal possible number of characters**.
- If multiple minimal solutions exist, you may output **any** of them.

### Examples

**Input**
```text
2+4=5
```

**Output**
```text
21+4=25
```

**Input**
```text
1+1=3
```

**Output**
```text
1+31=32
```

**Input**
```text
1+1=2
```

**Output**
```text
1+1=2
```

### ideas
1. dp[i][j][k][c] 表示在处理为a[:i] + b[:j] = c[:k] 且a+b carry 了c时的最短长度
2. 最后的结果 = dp[n][m][l][0]