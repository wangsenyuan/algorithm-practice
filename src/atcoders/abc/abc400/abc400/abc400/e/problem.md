# E - Ringo's Favorite Numbers 3 (ABC400)

**Contest:** [ABC400](https://atcoder.jp/contests/abc400) — AtCoder Beginner Contest 400  
**Task:** [https://atcoder.jp/contests/abc400/tasks/abc400_e](https://atcoder.jp/contests/abc400/tasks/abc400_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 425 points

## Problem Statement

A positive integer N is a **400 number** if and only if it satisfies both of the following two conditions:

- N has exactly 2 distinct prime factors.
- For each prime factor p of N, p divides N an even number of times. More formally, the maximum non-negative integer k such that p^k divides N is even.

For example, 400 is a 400 number because its prime factors are 2 and 5 (exactly two distinct primes), 2 divides 400 four times, and 5 divides it twice — all exponents are even.

Process Q queries. Each query gives you an integer A, so find the largest 400 number not exceeding A. Under the constraints of this problem, a 400 number not exceeding A always exists.

## Constraints

- 1 ≤ Q ≤ 2 × 10^5
- For each query, 36 ≤ A ≤ 10^12
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
Q
query_1
query_2
⋮
query_Q
```

Here, query_i is the i-th query, given in the following format:

```text
A
```

## Output

Print Q lines. The i-th line should contain the answer to the i-th query.

## Sample Input 1

```text
5
404
36
60
1000000000000
123456789
```

## Sample Output 1

```text
400
36
36
1000000000000
123454321
```

There are exactly 2 prime factors of 400: 2 and 5. Also, 2 divides 400 four times and 5 divides it twice, so 400 is a 400 number. None of 401, 402, 403, and 404 is a 400 number, so the answer is 400.

### ideas
