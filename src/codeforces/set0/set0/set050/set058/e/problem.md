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

### Key Insights

1. **Shortest-path on a digit DP graph.** Reverse all three strings so digits are processed LSB→MSB (matching how addition works). At each step, choose a digit for x and y; the z digit is forced by `dx + dy + carry`. Edges that don't match the next character of a/b/c cost 1 (an "extra" digit); matching edges cost 0. Dijkstra finds the minimum total extras, i.e. minimum total length.

2. **State design.** `(ia, ib, ic, carry, doneX, doneY, msdX, msdY)` where:
   - `ia/ib/ic` — how many characters of reversed a, b, c have been matched as subsequences of x, y, z.
   - `carry` — the addition carry (0 or 1).
   - `doneX/doneY` — whether x/y has finished emitting digits (allows different lengths).
   - `msdX/msdY` — the most-significant digit emitted so far for x/y, used to prevent leading zeros at termination.

3. **State space is small.** Since a, b, c ≤ 10^6, each has ≤ 7 digits. Total states ≤ 8³ × 2 × 4 × 11² ≈ 500K — easily fits in memory and runs fast with Dijkstra.

4. **Stopping an operand.** An operand x (or y) can stop only when all of a (or b) has been matched *and* its top digit is non-zero (no leading zero). Once stopped, it contributes 0 to every subsequent column. Both operands stopped + carry = 0 is the termination condition.

5. **Backtracking.** Store the parent state and chosen digits for each relaxation. Walk from the goal back to the start to collect x, y, z digit-by-digit (already in the right order since the last digit processed is the MSB).