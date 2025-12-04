Polycarpus loves lucky numbers. Everybody knows that lucky numbers are positive integers, whose decimal representation (without leading zeroes) contain only the lucky digits x and y. For example, if x = 4, and y = 7, then numbers 47, 744, 4 are lucky.

## Undoubtedly Lucky Numbers

Polycarpus loves lucky numbers. A lucky number is a positive integer whose decimal representation (without leading zeroes) contains only two digits `x` and `y`. For example, if `x = 4` and `y = 7` then 47, 744 and 4 are lucky numbers.

We call a positive integer a undoubtedly lucky if there exist digits `x` and `y` (0 ≤ x, y ≤ 9) such that the decimal representation of the integer (without leading zeroes) contains only digits `x` and `y`. Given an integer `n`, count how many positive integers not exceeding `n` are undoubtedly lucky.

## Input

The first line contains a single integer `n` (1 ≤ n ≤ 10^9) — Polycarpus's number.

## Output

Print a single integer — the count of positive integers not exceeding `n` that are undoubtedly lucky.

Examples
### Example 1
Input
```
10
```
Output
```
10
```
### Example 2
Input
```
123
```
Output
```
113
```
Note
In the first test sample all numbers that do not exceed 10 are undoubtedly lucky.

In the second sample numbers 102, 103, 104, 105, 106, 107, 108, 109, 120, 123 are not undoubtedly lucky.

