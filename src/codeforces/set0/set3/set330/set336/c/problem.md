# Problem Description

Vasily the bear has got a sequence of positive integers $a_1, a_2, ..., a_n$. Vasily the Bear wants to write out several numbers on a piece of paper so that the beauty of the numbers he wrote out was maximum.

## Beauty Definition

The beauty of the written out numbers $b_1, b_2, ..., b_k$ is such maximum non-negative integer $v$, that number $b_1$ and $b_2$ and ... and $b_k$ is divisible by number $2^v$ without a remainder. 

If such number $v$ doesn't exist (that is, for any non-negative integer $v$, number $b_1$ and $b_2$ and ... and $b_k$ is divisible by $2^v$ without a remainder), the beauty of the written out numbers equals $-1$.

## Task

Tell the bear which numbers he should write out so that the beauty of the written out numbers is maximum. If there are multiple ways to write out the numbers, you need to choose the one where the bear writes out as many numbers as possible.

**Note:** Expression $x$ and $y$ means applying the bitwise AND operation to numbers $x$ and $y$. In programming languages C++ and Java this operation is represented by "&", in Pascal — by "and".

## Input

- The first line contains integer $n$ ($1 \leq n \leq 10^5$)
- The second line contains $n$ space-separated integers $a_1, a_2, ..., a_n$ ($1 \leq a_1 < a_2 < ... < a_n \leq 10^9$)

## Output

- In the first line print a single integer $k$ ($k > 0$), showing how many numbers to write out
- In the second line print $k$ integers $b_1, b_2, ..., b_k$ — the numbers to write out

**Constraints:**
- You are allowed to print numbers $b_1, b_2, ..., b_k$ in any order, but all of them must be distinct
- If there are multiple ways to write out the numbers, choose the one with the maximum number of numbers to write out
- If there still are multiple ways, you are allowed to print any of them