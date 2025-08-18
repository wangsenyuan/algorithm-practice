# Problem D: Rick and Morty's Acoustic Ridge

## Problem Description

Rick and Morty like to go to the ridge High Cry for crying loudly — there is an extraordinary echo. Recently they discovered an interesting acoustic characteristic of this ridge: if Rick and Morty begin crying simultaneously from different mountains, their cry would be heard between these mountains up to the height equal to the **bitwise OR** of mountains they've climbed and all the mountains between them.

## Bitwise OR Operation

**Bitwise OR** is a binary operation which is determined the following way:

Consider representation of numbers x and y in binary numeric system (possibly with leading zeroes):
- x = xₖ...x₁x₀
- y = yₖ...y₁y₀

Then z = x | y is defined as: z = zₖ...z₁z₀, where:
- zᵢ = 1, if xᵢ = 1 or yᵢ = 1
- zᵢ = 0, otherwise

In other words, a digit of bitwise OR of two numbers equals zero if and only if digits at corresponding positions in both numbers equal zero.

**Example:** Bitwise OR of numbers 10 = 1010₂ and 9 = 1001₂ equals 11 = 1011₂.

**Programming Note:** In C/C++/Java/Python this operation is defined as `|`, and in Pascal as `or`.

## Task

Help Rick and Morty calculate the number of ways they can select two mountains such that if they start crying from these mountains, their cry will be heard **above** these mountains and all mountains between them.

More formally, you should find the number of pairs (l, r) where 1 ≤ l < r ≤ n, such that the bitwise OR of heights of all mountains between l and r (inclusive) is **larger than** the height of any mountain in this interval.

## Input

- **First line:** Integer n (1 ≤ n ≤ 200,000) — the number of mountains in the ridge
- **Second line:** n integers aᵢ (0 ≤ aᵢ ≤ 10⁹) — the heights of mountains in order they are located in the ridge

## Output

Print a single integer — the number of ways to choose two different mountains.

## Examples

### Example 1
**Input:**
```
5
3 2 1 6 5
```

**Output:**
```
8
```

**Explanation:** All valid pairs are (numbering from one):
- (1, 4), (1, 5), (2, 3), (2, 4), (2, 5), (3, 4), (3, 5), (4, 5)

### Example 2
**Input:**
```
4
3 3 3 3
```

**Output:**
```
0
```

**Explanation:** There are no valid pairs because for any pair of mountains, the height of cry from them is 3, which equals the height of any mountain in the interval.


### ideas
1. a[l] | a[r] > max(a[l...r])
2. 对于给定的最大值x和a[r], 似乎是可以找到对应a[l]的个数的？
3. 