### Summary

We have an `n × n` grid; initially, every cell is “evil”.  
Casting a **Purification** spell on cell `(r, c)` instantly purifies **all cells in row `r` and all cells in column `c`** (including `(r, c)` itself).  
Some cells are marked `'E'` (extra evil): you are **not allowed** to cast the spell on those cells, though they can still be purified via their row/column.

Goal:

- Purify **every** cell in the grid.
- Minimize the number of spell casts.
- If impossible, output `-1`.

Input:

- `n` (`1 ≤ n ≤ 100`)
- An `n × n` grid of characters:
  - `'.'` — a normal cell (you may cast on it).
  - `'E'` — an “extra evil” cell (you may **not** cast on it).

Output:

- If impossible: `-1`.
- Otherwise:
  - Let the minimum number of spells be `x`.
  - Print `x` lines, each with `row column`, giving cells where you should cast the spell.

Any valid plan with the minimum possible number of spells is accepted.***


### ideas
1. 如果存在一个cell，它同行全部是E，同列全部是是E， -1
2. 否则肯定存在一个答案，且答案 肯定 = n
3. 那些E是不能落子的地方，但是每行每列，都必须落子