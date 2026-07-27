# C. Minesweeper

[Problem link](https://codeforces.com/problemset/problem/2199/C)

**Contest:** [Codeforces contest 2199](https://codeforces.com/contest/2199)

## Problem

Construct a Minesweeper field with `2` rows and some number of columns. Each
cell is either empty (`.`) or a mine (`*`). Two cells are neighbors if they
share a side or a corner.

The field must satisfy:

- every empty cell neighbors at most one mine;
- the number of empty cells that neighbor at least one mine equals `k`;
- among all such fields, the number of columns `n` is minimized.

Construct any such minimum-width field, or report that it is impossible.

## Constraints

- `1 <= t <= 100`
- `1 <= k <= 100`

## Input

```text
t
k1
k2
...
kt
```

## Output

For each test case:

- if impossible, print `NO`;
- otherwise print:
  ```text
  YES
  n
  <row1 of length n>
  <row2 of length n>
  ```

## Sample 1

```text
Input
1

Output
YES
1
*
.
```

## Sample 2

```text
Input
4

Output
NO
```

## Sample 3

```text
Input
8

Output
YES
5
*....
...*.
```

## Sample 4

```text
Input
10

Output
YES
6
.*..*.
......
```

## Sample 5

```text
Input
9

Output
NO
```

## ideas
1. f(w) = permutations 的数量, permutation满足 
2. 如果 w[i] = 1, 那么存在(l, r)mex(p[l]... p[r]) = i
3. ..  w[i] = 0, ... 不存在
4. w[1] = 1, w[n] = 1
5. w[i] = 1, w[i+1] = 0
6. w[i] = 1, 说明存在一个区域包含[0....i-1], l1, r1 
7. w[i+1] = 0, 如果i现在就在区域0...i-1中, 那么mex(p[l1, r1]) = i + 1了 => i必须在l1, r1的外部, 
8. 假设r1 < r2 (p[r2] = i)
9. 在r1和r2中间, 必须至少存在i+1
10. 假设w[j] = 1 (j > i)
11. 那么i+1, i+2, ... j- 1, 这些必须在(r1...r2)之间吗?
12. 假设不是的, 比如我们把i+1移动到l1的前面, 那么马上就可以得出mex(l1...r2) = i+1 => w[i+1] = 1
13. 如果是i+2呢? =》w[i+2] = 1
14. 所以, i+1, i+2, j-1 必须在区间 (r1, r2)中间
15. 然后, 必须是i; 
16. [0....i-1, ???? i]
17. 但是, i也可以在l1的左端, 所以*2即可