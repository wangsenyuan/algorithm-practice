# Problem

Little Gerald and his coach Mike play an interesting game.

At the start of the game there are:

- a pile of **n** candies
- a pile of **m** stones

Gerald and Mike move in turns, with **Mike moving first**.

On **each Mike move**:

- Mike looks at how many candies and stones Gerald has eaten so far.
- Suppose Gerald has eaten `a` candies and `b` stones.
- Mike gives Gerald `f(a, b)` prize points.

On **each Gerald move**:

- Gerald chooses to eat **one** item:
  - either one candy (if any candies remain), or
  - one stone (if any stones remain).

The game continues until Gerald has eaten **all but one candy and all but one stone**.  
When Mike sees that there is exactly one candy and one stone left, he awards points **one last time**, and then the game ends.

Constraints during the game:

- Gerald is **not allowed** to eat all candies.
- Gerald is **not allowed** to eat all stones.

Your task: tell Gerald how to play to **maximize** the total number of prize points he receives; i.e. you must find one optimal play sequence.

The scoring function is defined as:

\[
f(a, b) = (x_a + y_b) \bmod p
\]

using zero-based indices `a` and `b` into arrays `x` and `y`.

---

## Input

- The first line contains three integers `n`, `m`, `p`  
  (`1 ≤ n, m ≤ 20000`, `1 ≤ p ≤ 10^9`).

- The second line contains `n` integers: `x₀, x₁, …, xₙ₋₁`  
  (`0 ≤ xᵢ ≤ 20000`).

- The third line contains `m` integers: `y₀, y₁, …, yₘ₋₁`  
  (`0 ≤ yⱼ ≤ 20000`).

The value of `f(a, b)` is:

\[
f(a, b) = (x_a + y_b) \bmod p.
\]

---

## Output

Print:

1. On the first line: the **maximum total** prize points Gerald can earn.
2. On the second line: a string of length `n + m − 2` consisting only of characters `'C'` and `'S'`:
   - the `i`-th character is `'C'` if Gerald's `i`-th move is eating a **candy**
   - the `i`-th character is `'S'` if Gerald's `i`-th move is eating a **stone**.

This sequence must correspond to one optimal strategy.

---

## Examples

### Example 1

Input:

```text
2 2 10
0 0
0 1
```

Output:

```text
2
SC
```

### Example 2

Input:

```text
3 3 10
0 2 0
0 0 2
```

Output:

```text
10
CSSC
```

### Example 3

Input:

```text
3 3 2
0 1 1
1 1 0
```

Output:

```text
4
SCSC
```

---

## Note

In the first test:

- If Gerald first eats a **stone**, Mike gives him `f(0, 1) = (x₀ + y₁) mod p = 1` point.
- If Gerald first eats a **candy**, he gets `f(1, 0) = (x₁ + y₀) mod p = 0` points.

Before Gerald's first move, points = `f(0, 0) = (x₀ + y₀) mod p = 0`.  
After the optimal sequence (`S` then `C`), total points become `2`.  
So the maximum achievable is `2`, obtained by first eating a stone then a candy.

---

## Short summary

- **State:** after Gerald has eaten `a` candies and `b` stones, Mike awards `f(a, b) = (x_a + y_b) mod p`.
- **Moves:** Gerald makes `n + m − 2` moves total, each move is `'C'` or `'S'`, but he must leave exactly **one** candy and **one** stone uneaten.
- **Goal:** choose the order of `'C'` and `'S'` moves to **maximize** the sum of `f(a, b)` over all Mike moves (including the last one).

### ideas

1. 不考虑复杂性, `dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + f(x[i], y[j])`
1. `20000 * 20000 = 4 * 1e8`
1. 好像时间可以。但是空间不够。
1. 按照 `i + j` 的顺序处理？这样，最大值，可以省空间
1. 但是没法恢复出来。
1. 所以，还需要用 bitset? `0` 表示从左边来，`1` 表示从上部来？

---

## Key insights

1. Let `dp[i][j]` be the maximum total score after Gerald has eaten `i` candies and `j` stones.
1. Then the transition is:

```text
dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + f(i, j)
```

   because the last move into state `(i, j)` must be either:
   eat one candy from `(i-1, j)`, or eat one stone from `(i, j-1)`.

1. The answer is `dp[n-1][m-1]`, since the game ends when exactly one candy and one stone remain.
1. A full `n x m` DP table is too large when `n, m <= 20000`.
1. Observe that states with the same `i + j` lie on the same diagonal, and the transition into diagonal `w` only depends on diagonal `w - 1`.
1. So we can process diagonals one by one and keep only one DP array for the current frontier, which reduces score computation to `O(n)` memory.
1. The remaining problem is path reconstruction.
1. For each diagonal `w`, we store one bitset:

   - bit `i = 1` means state `(i, w-i)` came from `(i-1, w-i)` by taking a candy
   - bit `i = 0` means it came from `(i, w-i-1)` by taking a stone

1. In Go, this bitset is stored with `big.Int`, which is convenient and compact enough here.
1. After finishing the DP, start from `(n-1, m-1)` and walk backward through the diagonals:

   - if the bit for the current `i` is `1`, append `'C'` and decrease `i`
   - otherwise append `'S'`

1. Reverse the collected moves to obtain one optimal strategy.
1. This version is especially easy to read because the DP and the traceback both follow the same diagonal order.

### Complexity

- Time: `O(nm)`
- Memory:
- `O(n)` for the rolling DP over diagonals
- plus one `big.Int` bitset per diagonal for traceback

This is small enough for the memory limit, while still producing one optimal strategy.

## editorial

Essence of problem is that there is a board n × m in cell of wich placed numbers. And one must go from cell (0, 0) to cell (n - 1, m - 1), doing moves to one cell up and right (that is, increasing by 1 on of coordinates), maximizing sum o number on the cell in the path.
Gerald's problems is determine m + n - 1 cells of optimal path. Lets start with search one, middle cell. Lets . What cell of board can be k-th in path? It is cell, sum of coordinate of wich is equal to k, thus, it is diagonal of doard. And then we can calculate, what maximum sum Gerald can collect, came to every cell in the diagonal, by dynamic programming in lower triangle. And the same way, we can calculate, what maximum sum Gerald can collect, came to cell (n-1,m-1), starting from every cell in the diagonal. Sumed up this to values, we calculate, what maximum sum Gerald can collect, came to from cell (0, 0) to cell (n - 1, m - 1), travel throw every cell in the diagonal. And now we will find cell (x, y), in wich maximum is reached. It is k-th cell of the path. We used time O((n + m)2) and memory O(n + m).
Then we make recursive call on subproblems. In other word, will find optimal path from cell (0, 0) to cell (x, y) and from cell (x, y) to cell (n, m).
It is evident, this solution take memory O(n+m). Why it tkae time O((n+m)^2)?
Lets n + m is r.
Once we are find middle cell of path of length k. Twice we are find middle cell of path of length . Four times we are find middle cell of path of length . And so on. Therefore, time of program working will be . Thus, this solution take time O((n + m)2).
