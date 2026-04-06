# E. Easy Chess

Elma is learning chess figures. She learned that a **rook** can move either horizontally or vertically. To enhance her understanding of rook movement, Elma's grandmother gave Elma an \(8 \times 8\) chessboard and asked her to find a way to move the rook from **a1** to **h8** making exactly \(n\) moves, so that all visited cells are different.

A **visited cell** is the initial cell **a1** and each cell on which the rook lands after a move.

## Input

The input contains a single integer \(n\) (\(2 \le n \le 63\)) — the desired number of moves.

## Output

Output a space-separated list of \(n + 1\) visited cells in the order they are visited by the rook. All cells must be different. The list should start with **a1** and end with **h8**. A solution always exists.

## key insights

1. The board size is fixed: only `8 x 8 = 64` cells.
   - So we do not need a complicated formula.
   - It is enough to build one universal route and then extract the required length from it.

2. Build a Hamiltonian path from `a1` to `h8`.
   - Visit the first 6 columns in a vertical snake:
     - `a1 -> a2 -> ... -> a8`
     - `b8 -> b7 -> ... -> b1`
     - and so on
   - Then finish the last two columns `g` and `h` with a fixed 16-cell path from `g1` to `h8`.
   - This gives one path visiting all 64 cells exactly once.

3. Now we only need exactly `n` moves, i.e. exactly `n + 1` visited cells.
   - Instead of using the whole Hamiltonian path, choose a subsequence of it.
   - Keep the first cell `a1` and the last cell `h8`.
   - For any two consecutive kept cells, they must lie on the same row or the same column, so the rook can jump directly between them in one move.

4. Since the master path has only 64 cells, a tiny DP is enough.
   - `reach[i][len]` = whether we can end at the `i`-th cell of the master path using exactly `len` kept cells.
   - Transition from `j` to `i` if:
     - `j < i`
     - the two cells are in the same row or column
     - `reach[j][len-1]` is true

5. Store predecessors and reconstruct one valid subsequence of length `n + 1`.
   - The output sequence itself is the visited cells.
   - All cells are distinct because the master path uses each board cell once.

6. Why this always works.
   - For this specific 64-cell master path, every required length from `2` to `63` is reachable by such a subsequence.
   - Therefore a valid answer always exists, exactly as the statement guarantees.
