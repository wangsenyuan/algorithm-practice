# Problem D: Birthday Sweets

Little Petya has a birthday soon. Due to this wonderful event, Petya's friends decided to give him sweets. The total number of Petya's friends equals to n.

## Problem Definition

Let us remind you the definition of the greatest common divisor: GCD(a₁, ..., aₖ) = d, where d represents such a maximal positive number that each aᵢ (1 ≤ i ≤ k) is evenly divisible by d. At that, we assume that all aᵢ's are greater than zero.

Knowing that Petya is keen on programming, his friends have agreed beforehand that:
- The 1st friend gives a₁ sweets
- The 2nd friend gives a₂ sweets
- ...
- The nth friend gives aₙ sweets

## Constraints

At the same time, for any i and j (1 ≤ i, j ≤ n) they want:
1. GCD(aᵢ, aⱼ) ≠ 1
2. GCD(a₁, a₂, ..., aₙ) = 1
3. All aᵢ should be distinct

Help the friends to choose the suitable numbers a₁, ..., aₙ.

## Input

The first line contains an integer n (2 ≤ n ≤ 50).

## Output

If there is no answer, print "-1" without quotes. Otherwise print a set of n distinct positive numbers a₁, a₂, ..., aₙ.

**Output format:**
- Each line must contain one number
- Each number must consist of not more than 100 digits
- Numbers must not contain any leading zeros
- If there are several solutions, print any of them

## Required Conditions

Do not forget that all of the following conditions must be true:

1. For every i and j (1 ≤ i, j ≤ n): GCD(aᵢ, aⱼ) ≠ 1
2. GCD(a₁, a₂, ..., aₙ) = 1
3. For every i and j (1 ≤ i, j ≤ n, i ≠ j): aᵢ ≠ aⱼ

## Note

Please, do not use %lld specificator to read or write 64-bit integers in C++. It is preferred to use cout (also you may use %I64d).