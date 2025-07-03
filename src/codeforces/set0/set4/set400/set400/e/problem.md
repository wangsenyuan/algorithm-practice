# Inna and Binary Logic

## Problem Description

Inna is fed up with jokes about female logic. So she started using binary logic instead.

Inna has an array of `n` elements `a1[1], a1[2], ..., a1[n]`. Girl likes to train in her binary logic, so she does an exercise consisting of `n` stages:

- On the **first stage** Inna writes out all numbers from array `a1`
- On the **i-th** (i ≥ 2) stage girl writes all elements of array `ai`, which consists of `n - i + 1` integers
- The **k-th** integer of array `ai` is defined as follows: `ai[k] = ai-1[k] AND ai-1[k + 1]`

Here `AND` is bit-wise binary logical operation.

Dima decided to check Inna's skill. He asks Inna to change array, perform the exercise and say the sum of all elements she wrote out during the current exercise.

Help Inna to answer the questions!

## Input

The first line contains two integers `n` and `m` (1 ≤ n, m ≤ 10^5) — size of array `a1` and number of Dima's questions. 

Next line contains `n` integers `a1[1], a1[2], ..., a1[n]` (0 ≤ ai ≤ 10^5) — initial array elements.

Each of next `m` lines contains two integers — Dima's question description. Each question consists of two integers `pi, vi` (1 ≤ pi ≤ n; 0 ≤ vi ≤ 10^5). For this question Inna should make `a1[pi]` equals `vi`, and then perform the exercise. 

**Please, note that changes are saved from question to question.**

## Output

For each question print Inna's answer on a single line.

## ideas
1. a[1], a[2], a[3], a[4] => a[1] & a[2], a[2] & a[3], a[3] & a[4], 
2. 可以每一位单独处理，看它经过多少轮后，会变成0
3. a[i][d] 为1， 如果离他（同一位）最近的0的距离为w, 那么经过w次后，a[i][d]变成了0
4. 在这之前，它的贡献都是 1 << d （1 << d) * w
5. 修改一个数后，一个是它本身的贡献，会改变，还有一个是对它前面的数的影响
6. 比如原来停在它这里的，还有原来连接它这里的
7. 但是整体的更新应该可以在30次内完成
8. 假设有一段连续的1，长度为m， 那么总的贡献 = m + m - 1 + m - 2 + ... + 1
9. sum = (m + 1) * m / 2
10. got