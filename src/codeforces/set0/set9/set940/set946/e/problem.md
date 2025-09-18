# Problem E: Beautiful Numbers

## Problem Description

Yes, that's another problem with definition of "beautiful" numbers.

Let's call a positive integer x **beautiful** if:
1. Its decimal representation (without leading zeroes) contains **even number of digits**
2. There exists a **permutation** of this representation which is **palindromic**

For example, 4242 is a beautiful number:
- It contains 4 digits (even)
- There exists a palindromic permutation: 2442

Given a positive integer s, find the **largest beautiful number** which is less than s.

## Definition of Beautiful Numbers

A number is beautiful if and only if:
- **Even digit count**: The number has an even number of digits
- **Palindromic permutation exists**: Some rearrangement of its digits forms a palindrome

**Examples:**
- `4242` → `2442` (palindrome) ✓ Beautiful
- `1234` → No palindromic permutation exists ✗ Not beautiful
- `1221` → Already a palindrome ✓ Beautiful
- `123` → Odd number of digits ✗ Not beautiful

## Input

The first line contains one integer t (1 ≤ t ≤ 10⁵) — the number of testcases you have to solve.

Then t lines follow, each representing one testcase and containing one string which is the decimal representation of number s. It is guaranteed that:
- This string has even length
- Contains no leading zeroes
- There exists at least one beautiful number less than s

The sum of lengths of s over all testcases doesn't exceed 2·10⁵.

## Output

For each testcase print one line containing the largest beautiful number which is less than s (it is guaranteed that the answer exists).

## Examples

### Example 1
**Input:**
```
4
89
88
1000
28923845
```

**Output:**
```
88
77
99
28923839
```

**Explanation:**
- For `89`: The largest beautiful number < 89 is `88` (can be permuted to `88` which is a palindrome)
- For `88`: The largest beautiful number < 88 is `77` (palindrome itself)
- For `1000`: The largest beautiful number < 1000 is `99` (palindrome itself)
- For `28923845`: The largest beautiful number < 28923845 is `28923839` (can be permuted to form a palindrome)

## Approach

This is a **constructive problem** where we need to:

1. **Understand the constraint**: A number is beautiful if it has even digits and can be permuted to form a palindrome
2. **Key insight**: For a number to be palindromic after permutation, each digit (except possibly one) must appear an even number of times
3. **Strategy**: Find the largest number less than s that satisfies the beautiful number criteria

**Algorithm:**
- Start from s-1 and work downwards
- For each candidate number, check if it has even number of digits
- Check if it can be permuted to form a palindrome (each digit appears even times, except at most one)
- Return the first valid number found

## Notes

- The problem guarantees that at least one beautiful number exists less than s
- We need to find the **largest** such number
- A palindrome can have at most one digit with odd frequency (the middle character)
- For even-length palindromes, all digits must appear even number of times

## ideas
1. let n = len(s)
2. 如果 n > 1, 那么如果无法得到一个和s等长的结果, 那么 9..9 长度为n-1就是结果
3. 然后找长度为n,且 < s, 且能组成回文的字符串
4. dp[mask][0/1]表示到目前为止,是否小于(0)等于(1)s, 且由mask组成的字符串（只考虑奇数次的字符）