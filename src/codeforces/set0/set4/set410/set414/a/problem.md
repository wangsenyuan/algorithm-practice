# Problem A - Mashmokh and Numbers

## Description

It's holiday. Mashmokh and his boss, Bimokh, are playing a game invented by Mashmokh.

In this game Mashmokh writes sequence of n distinct integers on the board. Then Bimokh makes several (possibly zero) moves. On the first move he removes the first and the second integer from the board, on the second move he removes the first and the second integer of the remaining sequence from the board, and so on. Bimokh stops when the board contains less than two numbers. When Bimokh removes numbers x and y from the board, he gets gcd(x, y) points. At the beginning of the game Bimokh has zero points.

Mashmokh wants to win in the game. For this reason he wants his boss to get exactly k points in total. But the guy doesn't know how to choose the initial sequence in the right way.

Please, help him. Find n distinct integers a₁, a₂, ..., aₙ such that his boss will score exactly k points. Also Mashmokh can't memorize too huge numbers. Therefore each of these integers must be at most 10⁹.

## Input

The first line of input contains two space-separated integers n, k (1 ≤ n ≤ 10⁵; 0 ≤ k ≤ 10⁸).

## Output

If such sequence doesn't exist output -1, otherwise output n distinct space-separated integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁹).

## Examples

### Example 1
**Input:**
```
5 2
```

**Output:**
```
1 2 3 4 5
```

### Example 2
**Input:**
```
5 3
```

**Output:**
```
2 4 3 7 1
```

### Example 3
**Input:**
```
7 2
```

**Output:**
```
-1
```

## Note

gcd(x, y) is greatest common divisor of x and y.


### ideas
1. 序列也可以自己选～
2. 那就好搞了
3. 如果 k < n / 2 => -1
4. 否则肯定有结果