# Problem E: Sequence Merging

## Problem Description

You are given a sequence of positive integers a₁, a₂, ..., aₙ.

While possible, you perform the following operation:
- Find a pair of equal consecutive elements
- If there are more than one such pair, find the leftmost one (with the smallest indices of elements)
- If the two integers are equal to x, delete both and insert a single integer x + 1 in their place
- This way the number of elements in the sequence is decreased by 1 on each step

You stop performing the operation when there is no pair of equal consecutive elements.

## Example

If the initial sequence is `[5, 2, 1, 1, 2, 2]`, then:
- After the first operation: `[5, 2, 2, 2, 2]`
- After the second operation: `[5, 3, 2, 2]`
- After the third operation: `[5, 3, 3]`
- After the fourth operation: `[5, 4]`

After that, there are no equal consecutive elements left in the sequence, so you stop the process.

## Input

- The first line contains a single integer n (2 ≤ n ≤ 2·10⁵) — the number of elements in the sequence
- The second line contains the sequence of integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁹)

## Output

- In the first line print a single integer k — the number of elements in the sequence after you stop performing the operation
- In the second line print k integers — the sequence after you stop performing the operation

## Task

Determine the final sequence after you stop performing the operation.

