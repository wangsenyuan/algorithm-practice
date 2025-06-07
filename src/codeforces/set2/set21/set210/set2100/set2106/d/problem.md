# Flower Collection Problem

## Problem Description

Flower boy has a garden of 𝑛 flowers that can be represented as an integer sequence 𝑎₁,𝑎₂,…,𝑎ₙ, where 𝑎ᵢ is the beauty of the 𝑖-th flower from the left.

Igor wants to collect exactly 𝑚 flowers. To do so, he will walk the garden from left to right and choose whether to collect the flower at his current position. The 𝑖-th flower among ones he collects must have a beauty of at least 𝑏ᵢ.

Igor noticed that it might be impossible to collect 𝑚 flowers that satisfy his beauty requirements, so before he starts collecting flowers, he can pick any integer 𝑘 and use his magic wand to grow a new flower with beauty 𝑘 and place it anywhere in the garden (between two flowers, before the first flower, or after the last flower). Because his magic abilities are limited, he may do this at most once.

## Task

Output the minimum integer 𝑘 Igor must pick when he performs the aforementioned operation to ensure that he can collect 𝑚 flowers. If he can collect 𝑚 flowers without using the operation, output 0. If it is impossible to collect 𝑚 flowers despite using the operation, output -1.

## Input

The first line of the input contains a single integer 𝑡 (1 ≤ 𝑡 ≤ 10⁴) — the number of test cases.

For each test case:
- The first line contains exactly two integers 𝑛 and 𝑚 (1 ≤ 𝑚 ≤ 𝑛 ≤ 2⋅10⁵) — the number of flowers in the garden and the number of flowers Igor wants to collect, respectively.
- The second line contains 𝑛 integers 𝑎₁,𝑎₂,...,𝑎ₙ (1 ≤ 𝑎ᵢ ≤ 10⁹) — where 𝑎ᵢ is the beauty of the 𝑖-th flower from the left in the garden.
- The third line contains 𝑚 integers 𝑏₁,𝑏₂,...,𝑏ₘ (1 ≤ 𝑏ᵢ ≤ 10⁹) — where 𝑏ᵢ is the minimum beauty the 𝑖-th flower must have that Igor will collect.

It is guaranteed that the sum of 𝑛 over all test cases does not exceed 2⋅10⁵.

## Output

For each test case, on a separate line, output the minimum integer 𝑘 Igor must pick when he performs the aforementioned operation to ensure that he can collect 𝑚 flowers. If he can collect 𝑚 flowers without using the operation, output 0. If it is impossible to collect 𝑚 flowers despite using the operation, output -1.

## ideas
1. 先考虑在不适用wand的情况下，要怎么处理
2. 假设目前已经收集了j个花朵，如果a[i] >= b[j], 那么就收集，否则就不收集（这个策略是否ok？）
3. 假设不收集能给出更有利的答案，没有收集a[i], 而是收集了 a[k] (i < k)
4. 那么也不能收集(i+1, i + 2, )直到k
5. 这时候，交换为收集i是更优的策略，因为i+1, i+2, ...中有符合后续要求的序列
6. 假设不放置无法找到答案(k = 0 不work)
7. 必须在某个位置放置一个花朵
8. k确实是符合2分的性质的（如果在某个位置放置k的可以，那么放置k+1的肯定也ok）
9. 但是问题是有n+1个位置。肯定不能每个位置都检查一遍
10. 那么这个k应该放在哪里呢？
11. 还真是要检查所有的位置
12. 假设在i，i+1中间放置了一个k
13. 那么在i的前面匹配掉的b的序列假设是b[j], 那么k必须消耗掉b[j+1] (k >= b[j+1]), 且后缀要能消耗掉b[j+2]开始的序列