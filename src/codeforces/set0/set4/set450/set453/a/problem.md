# Problem Description

Twilight Sparkle was playing Ludo with her friends Rainbow Dash, Apple Jack and Flutter Shy. But she kept losing. Having returned to the castle, Twilight Sparkle became interested in the dice that were used in the game.

The dice has m faces: the first face of the dice contains a dot, the second one contains two dots, and so on, the m-th face contains m dots. Twilight Sparkle is sure that when the dice is tossed, each face appears with probability 1/m. Also she knows that each toss is independent from others. Help her to calculate the expected maximum number of dots she could get after tossing the dice n times.

## Input

A single line contains two integers m and n (1 ≤ m, n ≤ 10^5).

## Output

Output a single real number corresponding to the expected maximum. The answer will be considered correct if its relative or absolute error doesn't exceed 10^(-4).

## Examples

### Example 1

**Input:**

```text
6 1
```

**Output:**

```text
3.500000000000
```

### Example 2

**Input:**

```text
6 3
```

**Output:**

```text
4.958333333333
```

### Example 3

**Input:**

```text
2 2
```

**Output:**

```text
1.750000000000
```

## Note

Consider the third test example. If you've made two tosses:

- You can get 1 in the first toss, and 2 in the second. Maximum equals to 2.
- You can get 1 in the first toss, and 1 in the second. Maximum equals to 1.
- You can get 2 in the first toss, and 1 in the second. Maximum equals to 2.
- You can get 2 in the first toss, and 2 in the second. Maximum equals to 2.

The probability of each outcome is 0.25, that is expectation equals to:
(1 × 0.25) + (2 × 0.25) + (2 × 0.25) + (2 × 0.25) = 0.25 + 0.5 + 0.5 + 0.5 = 1.75
