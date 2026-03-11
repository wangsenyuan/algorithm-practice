# Problem

**题意简述**：n 个蛋，每个蛋 A 要价 ai、G 要价 gi，且 ai + gi = 1000。把每个蛋恰好分给一个孩子（A 或 G），要求付给 A 的总价 Sa 与付给 G 的总价 Sg 满足 |Sa - Sg| ≤ 500。求一种分配方案（输出 n 个 'A' 或 'G'），不可能则输出 -1。



## Input

- The first line contains integer n (1 ≤ n ≤ 10^6) — the number of eggs.
- Next n lines contain two integers ai and gi each (0 ≤ ai, gi ≤ 1000; ai + gi = 1000): ai is the price said by A. for the i-th egg and gi is the price said by G. for the i-th egg.

## Output

- If it is impossible to assign the painting, print "-1" (without quotes).
- Otherwise print a string of n letters "G" and "A". The i-th letter is the child who gets the i-th egg ("A" = A., "G" = G.). Must hold: |Sa - Sg| ≤ 500.
- If there are several solutions, print any.

## Examples

### Example 1

**Input**

```text
2
1 999
999 1
```

**Output**

```text
AG
```

### Example 2

**Input**

```text
3
400 600
400 600
400 600
```

**Output**

```text
AGA
```

## ideas
1. 按照a[i]升序排列（也就是按g[i])降序排列, a[i] + g[i] = 1000
2. 当n = 1， 也可以（把它分配给价格小的那方）
3. 上面按照两头开始分配
4. 如果 a[i] <= 500, 那么就分配给A，否则分配给G
5. 100, 100, 100, 100, 100, .... 900
6. 