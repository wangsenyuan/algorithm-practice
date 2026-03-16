### Problem

The Smart Beaver from ABBYY likes magic squares and wants to automate solving them.

A **magic square** is an `n × n` integer matrix such that:

- The sum of numbers in **each row** equals some number `s`.
- The sum of numbers in **each column** also equals `s`.
- The sum of numbers on the **main diagonal** equals `s`.
- The sum of numbers on the **secondary diagonal** equals `s`.

You are given a multiset of `n^2` integers `a_i`. Your task is to place these numbers into an `n × n` matrix so that they form a magic square, using each number exactly as many times as it appears in the input. It is guaranteed that at least one solution exists.

### Input

- First line: integer `n`.
- Second line: `n^2` integers `a_i` (`-10^8 ≤ a_i ≤ 10^8`), separated by spaces.

**Subtasks:**

- For 20 points: `1 ≤ n ≤ 3`
- For 50 points:
  - `1 ≤ n ≤ 4`
  - At most 9 distinct numbers among `a_i`
- For 100 points: `1 ≤ n ≤ 4`

### Output

- First line: the integer `s` — the common sum of each row, column, and both diagonals.
- Next `n` lines: `n` integers each, describing the magic square.
- If there are multiple valid magic squares, you may output **any** of them.

### Examples

**Input**
```text
3
1 2 3 4 5 6 7 8 9
```

**Output**
```text
15
2 7 6
9 5 1
4 3 8
```

**Input**
```text
3
1 0 -1 0 2 -1 -2 0 1
```

**Output**
```text
0
1 0 -1
-2 0 2
1 0 -1
```

**Input**
```text
2
5 5 5 5
```

**Output**
```text
10
5 5
5 5
```

### Short summary

- You must permute the given multiset of `n^2` integers into an `n × n` grid.
- All row sums, column sums, and both diagonal sums must be equal to the same value `s`.
- Since `n ≤ 4`, a carefully pruned backtracking or search over placements (respecting counts of each value) is feasible; any valid magic square is acceptable.


### ideas
1. s = sum(a) / n
2. backtrack