# Problem Statement

A piece of paper contains an array of $n$ integers $a_1, a_2, \ldots, a_n$. Your task is to find a number that occurs the maximum number of times in this array.

However, before looking for such a number, you are allowed to perform **no more than $k$** of the following operations:
- Choose an arbitrary element from the array and add 1 to it.
- You may increase the same element of the array multiple times.

Your task is to find the maximum number of occurrences of some number in the array after performing no more than $k$ allowed operations. If there are several such numbers, your task is to find the minimum one.

---

## Input

The first line contains two integers $n$ and $k$ ($1 \leq n \leq 10^5$; $0 \leq k \leq 10^9$) — the number of elements in the array and the number of operations you are allowed to perform, correspondingly.

The second line contains a sequence of $n$ integers $a_1, a_2, \ldots, a_n$ ($|a_i| \leq 10^9$) — the initial array. The numbers in the lines are separated by single spaces.

---

## Output

In a single line, print two numbers:
- The maximum number of occurrences of some number in the array after at most $k$ allowed operations are performed.
- The minimum number that reaches the given maximum.

Separate the printed numbers by whitespaces.


## ideas

- 要在最多k次操作后，找到可以得到的最多出现的次数，以及这个数字（多个数字的时候，取最小的）
- 出现次数，符合二分的情况
- 就是如果能出现x次，那么必然可以出现x-1次