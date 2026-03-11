# Problem

**题意简述**：给定 n 个非负整数 a1, a2, ..., an（n ≤ 3，0 ≤ ai < 300）。两人轮流操作，BitLGM 先手。每次可选：(1) 选一个 ai，选 x 满足 1 ≤ x ≤ ai，令 ai = ai - x；(2) 选一个 x，令所有 ai = ai - x。无法操作者输。双方最优，问谁赢（输出 "BitLGM" 或 "BitAryo"）。

## Input

- The first line contains an integer n (1 ≤ n ≤ 3).
- The next line contains n integers a1, a2, ..., an (0 ≤ ai < 300).

## Output

Write the name of the winner (provided that both players play optimally well). Either "BitLGM" or "BitAryo" (without the quotes).

## Examples

### Example 1

**Input**

```text
2
1 1
```

**Output**

```text
BitLGM
```

### Example 2

**Input**

```text
2
1 2
```

**Output**

```text
BitAryo
```

### Example 3

**Input**

```text
3
1 2 1
```

**Output**

```text
BitLGM
```

## Editorial

- **n = 1**: If a1 = 0 then BitAryo wins, otherwise BitLGM wins.
- **n = 2**: Define `win[i][j]` = whether (i, j) is a winning position. Compute for all i, j using a loop (check all possible moves). O(n³).
- **n = 3**: Similar to NIM. (i, j, k) is winning iff (i xor j xor k) ≠ 0. O(1). (Don't forget parentheses in code.)
- **Alternative for n = 3**: DP with `lose[i][j]`, `lose2[i][j]`, `win[i][j][k]`. O(n³).
- **Total**: DP for n = 2, O(1) for n = 1 and n = 3 → O(n²).

### Proof: n = 3, why first wins when i xor j xor k ≠ 0

Let s = i xor j xor k (nim-sum).

**Lemma 1** (nim-sum = 0 → losing): From any position with s = 0, every move leads to s' ≠ 0.

- **Type 1** (decrease one pile): Change (i, j, k) to (i', j, k) with i' < i. New nim-sum = i' xor j xor k. Since i xor j xor k = 0, we have j xor k = i. So new sum = i' xor i ≠ 0 (because i' ≠ i).
- **Type 2** (decrease all by x): (i, j, k) → (i-x, j-x, k-x). One can verify that (i-x) xor (j-x) xor (k-x) ≠ 0 when i xor j xor k = 0 and x > 0 (the subtraction breaks the xor balance).

**Lemma 2** (nim-sum ≠ 0 → winning): From any position with s ≠ 0, there exists a move that reaches s' = 0.

- Let d be the highest bit where s has a 1. At least one pile has a 1 in bit d (otherwise s would be 0 there). WLOG assume pile i has a 1 in bit d.
- Set t = s xor i (= j xor k). We have t xor j xor k = 0, so moving pile i to t would yield nim-sum 0.
- **Claim**: t < i. Proof: In bit d, i has 1 and s has 1, so t = s xor i has 0. All bits above d: s is 0, so t = i there. Thus t is strictly smaller than i in bit d, hence t < i.
- **Type 1 move**: Decrease pile i by (i - t), leaving (t, j, k). New nim-sum = t xor j xor k = 0. ✓

**Conclusion**: Positions with s ≠ 0 are winning (can force opponent into s = 0); positions with s = 0 are losing. So first wins iff i xor j xor k ≠ 0.