# Problem G: Prime Painting

## Problem Description

In Berland, prime numbers are fashionable — the respectable citizens dwell only on the floors with numbers that are prime numbers. The numismatists value particularly high the coins with prime nominal values. All the prime days are announced holidays!

Yet even this is not enough to make the Berland people happy. On the main street of the capital stand n houses, numbered from 1 to n. The government decided to paint every house a color so that **the sum of the numbers of the houses painted every color is a prime number**.

However, it turned out that not all the citizens approve of this decision — many of them protest because they don't want many colored houses on the capital's main street. That's why it is decided to use the **minimal possible number of colors**. The houses don't have to be painted consecutively, but every one of n houses should be painted some color. The one-colored houses should not stand consecutively, any way of painting is acceptable.

There are no more than 5 hours left before the start of painting, help the government find the way when the sum of house numbers for every color is a prime number and the number of used colors is minimal.

## Problem Requirements

- Paint n houses (numbered 1 to n) with colors
- **Constraint**: The sum of house numbers for each color must be a prime number
- **Goal**: Use the minimal possible number of colors
- **Flexibility**: Houses don't need to be painted consecutively
- **Output**: Any valid solution is acceptable

## Input

The single input line contains an integer n (2 ≤ n ≤ 6000) — the number of houses on the main streets of the capital.

## Output

Print the sequence of n numbers, where the i-th number stands for the number of color for house number i. Number the colors consecutively starting from 1. Any painting order is allowed. If there are several solutions to that problem, print any of them. If there's no such way of painting print the single number -1.

## Examples

### Example 1
**Input:**
```
8
```

**Output:**
```
1 2 2 1 1 1 1 2
```

**Explanation:**
- Houses 1, 4, 5, 6, 7 are painted color 1
- Houses 2, 3, 8 are painted color 2
- Sum for color 1: 1 + 4 + 5 + 6 + 7 = 23 (prime)
- Sum for color 2: 2 + 3 + 8 = 13 (prime)
- Total colors used: 2

## Approach

This is a **partitioning problem** where we need to:
1. Partition the set {1, 2, ..., n} into subsets
2. Each subset's sum must be a prime number
3. Minimize the number of subsets

**Key insights:**
- We need to find prime numbers that can be expressed as sums of consecutive or non-consecutive integers from 1 to n
- The total sum 1 + 2 + ... + n = n(n+1)/2 must equal the sum of all prime numbers used
- We can use dynamic programming or greedy approaches to find valid partitions

## Notes

- The problem asks for any valid solution, not necessarily the lexicographically smallest
- If no valid partition exists, output -1
- Colors are numbered starting from 1
- Houses can be painted in any order (non-consecutive painting is allowed)

### ideas
1. 叉，又是个英语问题， 意思是同一种颜色的，房子编号的sum 要是一个质数
2. 比如 [1 2 2 1 1 1 1 2]
3. 对于颜色1， 1 + 4 + 5 + 6 + 7 = 23
4. 对于颜色2,  2 + 3 + 8 = 13