# Problem E: String Game

## Problem Description

Two people play the following string game. Initially the players have got some string s. The players move in turns, the player who cannot make a move loses.

Before the game began, the string is written on a piece of paper, one letter per cell.

**Example of the initial situation at s = "abacaba":**
```
a b a c a b a
1 2 3 4 5 6 7
```

## Game Rules

A player's move consists of the following sequence of actions:

1. **Choose a piece of paper**: The player chooses one of the available pieces of paper with some string written on it. Let's denote it as t. Note that initially, only one piece of paper is available.

2. **Choose a valid position**: The player chooses in the string t = t₁t₂...tₙ a character in position i (1 ≤ i ≤ n) such that for some positive integer l (0 < i - l; i + l ≤ n) the following equations hold:
   - tᵢ₋₁ = tᵢ₊₁
   - tᵢ₋₂ = tᵢ₊₂
   - ...
   - tᵢ₋ₗ = tᵢ₊ₗ
   
   In other words, position i must be the center of a palindrome of length at least 3.

3. **Cut the string**: The player cuts the cell with the chosen character. As a result of the operation, he gets three new pieces of paper:
   - First piece: string t₁t₂...tᵢ₋₁
   - Second piece: single character tᵢ
   - Third piece: string tᵢ₊₁tᵢ₊₂...tₙ

**Example of making action (i = 4) with string s = "abacaba":**
```
Before: a b a c a b a
        1 2 3 4 5 6 7
        (cut at position 4)

After:  [a b a] [c] [a b a]
        piece1  piece2 piece3
```

## Task

Your task is to determine the winner provided that both players play optimally well. If the first player wins, find the position of character that is optimal to cut in his first move. If there are multiple positions, print the minimal possible one.

## Input

The first line contains string s (1 ≤ |s| ≤ 5000). It is guaranteed that string s only contains lowercase English letters.

## Output

If the second player wins, print in the single line "Second" (without the quotes). Otherwise, print in the first line "First" (without the quotes), and in the second line print the minimal possible winning move — integer i (1 ≤ i ≤ |s|).

## Examples

### Example 1
**Input:**
```
abacaba
```

**Output:**
```
First
2
```

**Explanation:** In the first sample, the first player has multiple winning moves. The valid positions to cut are 2, 4, and 6 (all centers of palindromes). The minimum one is to cut the character in position 2.

### Example 2
**Input:**
```
abcde
```

**Output:**
```
Second
```

**Explanation:** In the second sample, the first player has no available moves because there are no palindromes of length 3 or more in the string "abcde".

## Notes

- A valid move requires cutting at the center of a palindrome of length at least 3
- Players play optimally (using game theory principles)
- If multiple winning moves exist, choose the one with the smallest position index


### ideas
1. dp[l][r] = x 
2. 如果 dp[0][n] = 0 => second, else first
3. first 的时候，找到dp[0][l] = dp[l+2][n]的第一个位置
4. dp[l][r] 如果它中间没有这样子的字符 = 0
5. else dp[l][r] = 如果存在某个位置i, 使的， dp[l][i-1] = dp[i+1][r], 那么 dp[l][r] = 1
6. 