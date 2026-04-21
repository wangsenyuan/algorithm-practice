# Problem

There is a new TV game on BerTV. In this game two players get a number `A` consisting of `2n` digits. Before each turn, players determine who will make the next move. Each player should make exactly `n` moves.

On a turn, the selected player takes the leftmost digit of `A` and appends it to their number (`S1` for Homer, `S2` for Marge). After that this leftmost digit is erased from `A`.

Initially the numbers of both players (`S1` and `S2`) are empty. Leading zeroes in numbers `A`, `S1`, and `S2` are allowed.

In the end of the game, the first player gets `S1` dollars, and the second gets `S2` dollars.

One day Homer and Marge came to play the game. They managed to know the number `A` beforehand. They want to find a sequence of moves such that both of them make exactly `n` moves and their total prize is maximized.

Help them.

## Input

The first line contains integer `n` (`1 ≤ n ≤ 18`).

The second line contains integer `A` consisting of exactly `2n` digits. This number may have leading zeroes.

## Output

Output a string of `2n` characters `'H'` and `'M'` — the sequence of moves of Homer and Marge that gives them the maximum possible total prize.

Each player must make exactly `n` moves.

If there are several optimal solutions, output any of them.

## Examples

### Example 1

**Input**

```text
2
1234
```

**Output**

```text
HHMM
```

### Example 2

**Input**

```text
2
9911
```

**Output**

```text
HMHM
```


### ideas
1. dp[i][j] 表示处理完前（i+j)个字符后的最大的结果
2. dp[i+1][j] 表示吧第i+j+1个字符分配给A，那么它的贡献 = s[i] * pow(10, (n - i)) + dp[i][j]

## key insights

1. Since digits are taken from left to right, deciding whether current digit goes to Homer or Marge is equivalent to assigning this digit to one of two length-`n` numbers.

2. If the original string length is `2n`, define:
   - `i`: processed prefix length,
   - `j`: how many digits have been assigned to Homer in that prefix.
   Then Marge has `i - j` digits in the prefix.

3. DP state in `solution.go`:
   - `dp[i][j]` = maximum value of `S1 + S2` after processing first `i` digits, with exactly `j` digits assigned to Homer.

4. Positional contribution is computed directly:
   - assigning current digit `w` to Homer contributes `w * 10^(n-1-j)`,
   - assigning it to Marge contributes `w * 10^(n-1-(i-1-j))`.
   This works because `j` and `i-1-j` are already-fixed counts before placing the current digit.

5. After filling DP, the code backtracks from `dp[2n][n]` to reconstruct one optimal `'H'/'M'` sequence.

6. Complexity:
   - Time: `O((2n) * n)`,
   - Space: `O((2n) * n)`,
   which is easily safe for `n <= 18`.