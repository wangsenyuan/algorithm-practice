# Problem Description

Alexander is learning how to convert numbers from the decimal system to any other. However, he doesn't know English letters, so he writes any number only as a decimal number. This means that instead of the letter A, he will write the number 10.

For example, by converting the number 475 from decimal to hexadecimal system, he gets 11311 (475 = 1·16² + 13·16¹ + 11·16⁰).

Alexander lived calmly until he tried to convert the number back to the decimal number system.

Alexander remembers that he worked with little numbers, so he asks to find the minimum decimal number such that by converting it to the system with base n, he will get the number k.

## Input

- The first line contains the integer **n** (2 ≤ n ≤ 10⁹)
- The second line contains the integer **k** (0 ≤ k < 10⁶⁰), it is guaranteed that the number k contains no more than 60 symbols
- All digits in the second line are strictly less than n
- Alexander guarantees that the answer exists and does not exceed 10¹⁸
- The number k doesn't contain leading zeros

## Output

Print the number **x** (0 ≤ x ≤ 10¹⁸) — the answer to the problem.