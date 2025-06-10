# Zigzag Sequence and Function

## Problem Description

The court wizard Zigzag wants to become a famous mathematician. For that, he needs his own theorem, like the Cauchy theorem, or his sum, like the Minkowski sum. But most of all he wants to have his sequence, like the Fibonacci sequence, and his function, like the Euler's totient function.

### Zigzag Sequence Definition

The Zigag's sequence with the zigzag factor z is an infinite sequence Siz (i ≥ 1; z ≥ 2), that is determined as follows:

- Siz = 2, when i = 1
- Siz = (S(i-1)z + 1) mod z, when i > 1 and S(i-1)z < z - 1
- Siz = 0, when i > 1 and S(i-1)z = z - 1

Operation mod means taking the remainder from dividing number x by number y. For example, the beginning of sequence Si3 (zigzag factor 3) looks as follows: 1, 2, 3, 2, 1, 2, 3, 2, 1.

### Zigzag Function

Let's assume that we are given an array a, consisting of n integers. Let's define element number i (1 ≤ i ≤ n) of the array as ai. The Zigzag function is function Z(l, r, z), where l, r, z satisfy the inequalities 1 ≤ l ≤ r ≤ n, z ≥ 2.

### Operations

To become better acquainted with the Zigzag sequence and the Zigzag function, the wizard offers you to implement the following operations on the given array a:

1. **Assignment Operation**
   - Parameters: (p, v)
   - Description: Assigns value v to the p-th array element
   - Result: After the operation, the value of array element ap equals v

2. **Zigzag Operation**
   - Parameters: (l, r, z)
   - Description: Calculates the Zigzag function Z(l, r, z)

## Input Format

- First line: Integer n (1 ≤ n ≤ 10^5) — The number of elements in array a
- Second line: n space-separated integers a1, a2, ..., an (1 ≤ ai ≤ 10^9) — the elements of the array
- Third line: Integer m (1 ≤ m ≤ 10^5) — the number of operations
- Next m lines: Operation descriptions
  - If ti = 1 (assignment operation): Two space-separated integers pi, vi (1 ≤ pi ≤ n; 1 ≤ vi ≤ 10^9)
  - If ti = 2 (Zigzag operation): Three space-separated integers li, ri, zi (1 ≤ li ≤ ri ≤ n; 2 ≤ zi ≤ 6)

## Output Format

For each Zigzag operation, print the calculated value of the Zigzag function on a single line. Print the values in the order they are given in the input.

## Note

Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use cin, cout streams or the %I64d specifier.

## ideas
1. z比较小, <= 6. 所以可以直接使用6棵树
2. 问题出在给定z的情况下，如何维护数据，进而区间查询
3. 不考虑(z给定的情况下), let w = 2 * (z - 1)
   1. S[i] = 2  (i % w = 0)
   2. S[i] = i % w if i % w <= z 
   3. S[i] = 2 * z - i % w if i % w > z
4. 假设S[i] = i % w,            0,     1,    2, ... w - 1, 0,   1,   2 ... w - 1, 0
5. 令 R[i] = 2 * z - i % w      w + 2, w+1,  w, ... 3,     w+2, w+1, w, ....
6. tr[l...r] 表示对于区间l...r, i从l开始时， 这个区间的sum
7. 如何更新这个区间？ dp[l][0], dp[l][1], dp[l][2], 。。。 dp[l][i]表示对于l来说，它对应的位置为i时这个区间的sum
8. val[i][0] = val[2 * i + 1][0] + val[2 * i + 2][x] x和区间的长度有关系
9. 这个第二维，只需要处理 2 * (z - 1)个位置，所以不会超过10。 似乎是可以搞的 