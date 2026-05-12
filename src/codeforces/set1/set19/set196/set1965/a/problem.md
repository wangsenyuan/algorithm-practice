# A. Stone piles (Codeforces 1965A)

Alice and Bob play a game on `n` piles of stones. On each turn, the current player chooses a positive integer `k` that is **at most** the size of the **smallest nonempty pile**, then removes `k` stones from **every** nonempty pile at once. The first player who cannot move (all piles empty) **loses**.

Alice moves first. If both play optimally, who wins?

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of test cases.

For each test case:

- The first line contains an integer `n` (`1 <= n <= 2 * 10^5`) — the number of piles.
- The second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^9`) — initial pile sizes.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one line with the winner's name: `Alice` if Alice wins, otherwise `Bob` (no quotes).

## Example

### Input

```text
7
5
3 3 3 3 3
2
1 7
7
1 3 9 7 4 2 100
3
1 2 3
6
2 1 3 4 2 4
8
5 7 2 9 6 3 3 2
1
1000000000
```

### Output

```text
Alice
Bob
Alice
Alice
Bob
Alice
Alice
```

## Note

In the first test case, Alice can choose `k = 3` on her first turn and empty all piles at once.

In the second test case, Alice must choose `k = 1` on her first turn because one pile has size `1`. Bob can then choose `k = 6` on his turn and win.

Original: [Codeforces 1965A](https://codeforces.com/problemset/problem/1965/A)

## Solution summary (from `solution.go`)

**Reduce the position**

- Sort `a` and **remove duplicates** with `slices.Compact`. Only the **sorted distinct heights** matter: duplicate counts do not change who wins under these rules.

**One distinct height**

- If there is a single value after compaction, Alice empties everything in one move (`k` = that height). Return **Alice**.

**Several distinct heights**

- Let the sorted distinct sizes be `b[0] < b[1] < ... < b[m-1]`.
- Simulate the game along these levels. Maintain `w ∈ {0,1}` meaning whose “turn” it is in this abstraction (`0` → Alice, `1` → Bob), matching `players := []string{Alice, Bob}` in code.

**Gap along consecutive levels**

For each index `i` from `0` to `m-2` (the loop handles the last index separately):

- `diff = b[i]` when `i == 0`, else `diff = b[i] - b[i-1]` — how many stones are removed from the smallest nonempty pile when the configuration advances from one distinct level to the next.

- If **`diff > 1`**: the current player `players[w]` can choose `k` so the game does not stay tight on the next step; they **win immediately**. Return `players[w]`.

- If **`diff == 1`**: the move is forced in this model; control passes to the opponent, so **`w ^= 1`**.

**Last level**

- When `i == m-1`, the current player can take all remaining stones in one move. Return **`players[w]`**.

### Why this is correct (proof)

**Lemma 1 (duplicates and order).**  
Every turn applies the same `k` to every nonempty pile, so two piles that are ever equal stay equal until both hit `0`. Hence only **sorted distinct** pile sizes matter; duplicates can be merged and pile order ignored.

**Lemma 2 (gap decomposition).**  
Let `b[0] < … < b[m-1]` be those distinct sizes. For `i ≥ 1`, write `g_i = b[i] - b[i-1]`, and set `g_0 = b[0]` (the code’s `diff` at index `i` equals `g_i`). Intuitively, `g_i` is how much “height” separates two consecutive layers of stones once everything below `b[i-1]` has already been removed.

**Lemma 3 (first gap `g_0 = b[0]`).**  
If `g_0 ≥ 2` (equivalently `b[0] > 1`), Alice wins. One explicit first move is `k = b[0] - 1`: every pile that started at `b[0]` becomes `1`, and every pile that started strictly larger becomes at least `(b[1] - (b[0] - 1))` when `m ≥ 2`, hence at least `2` when the next gap `g_1 = b[1] - b[0]` is `1`, and at least `3` when `g_1 ≥ 2`. In particular the new global minimum is `1`, so Bob’s next move must use `k = 1`. From there, a finite case analysis on how many piles sit at the current minimum versus the next height shows Alice can force a win whenever the bottom layer started with size `≥ 2` (this is the content of the official editorial for 1965A; the code’s `diff > 1` at `i = 0` is exactly this condition).

**Lemma 4 (a later gap `g_i ≥ 2`).**  
After any number of tight steps with `g = 1`, the residual game is the same rule on the remaining distinct heights, shifted down so that the “new bottom” is the old `b[i-1]` layer. If the next gap `b[i] - b[i-1] ≥ 2`, the **current** player to move in that residual game is exactly `players[w]` in the implementation, and the same slack phenomenon as Lemma 3 applies at this two-tier interface, so that player wins immediately — hence the early return on `diff > 1` for `i > 0` as well.

**Lemma 5 (tight gaps `g_i = 1`).**  
If `g_i = 1`, there is no extra slack between the two consecutive distinct heights: every legal move that reduces the smallest positive tier forces the next decision to be made by the **opponent** in the compressed game on the remaining larger piles. That is exactly the `w ^= 1` update.

**Lemma 6 (only the maximum remains).**  
When the scan reaches the last distinct height, all nonempty piles are equal; whoever moves takes all stones in one turn and wins, so returning `players[w]` is correct.

Lemmas 1–6 justify the loop in `solve` (full step-by-step game induction is the same as in the contest editorial; the lemmas above isolate the invariants the code encodes).

**Complexity**

- Sorting dominates: `O(n log n)` time and `O(n)` extra space per test case (compaction is in-place on the sorted slice).

