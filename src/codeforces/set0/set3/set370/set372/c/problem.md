# Fireworks Festival

## Problem Description

A festival will be held in a town's main street. There are $n$ sections in the main street. The sections are numbered 1 through $n$ from left to right. The distance between each adjacent sections is 1.

In the festival $m$ fireworks will be launched. The $i$-th ($1 \leq i \leq m$) launching is on time $t_i$ at section $a_i$. If you are at section $x$ ($1 \leq x \leq n$) at the time of $i$-th launching, you'll gain happiness value $b_i - |a_i - x|$ (note that the happiness value might be a negative value).

You can move up to $d$ length units in a unit time interval, but it's prohibited to go out of the main street. Also you can be in an arbitrary section at initial time moment (time equals to 1), and want to maximize the sum of happiness that can be gained from watching fireworks. Find the maximum total happiness.

**Note:** Two or more fireworks can be launched at the same time.

## Input

The first line contains three integers $n$, $m$, $d$ ($1 \leq n \leq 150000$; $1 \leq m \leq 300$; $1 \leq d \leq n$).

Each of the next $m$ lines contains integers $a_i$, $b_i$, $t_i$ ($1 \leq a_i \leq n$; $1 \leq b_i \leq 10^9$; $1 \leq t_i \leq 10^9$). The $i$-th line contains description of the $i$-th launching.

It is guaranteed that the condition $t_i \leq t_{i+1}$ ($1 \leq i < m$) will be satisfied.

## Output

Print a single integer — the maximum sum of happiness that you can gain from watching all the fireworks.

## Notes

Please, do not write the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## ideas
1. 假设在时刻t时，You处在位置x处, 此时正好在发射烟火i，且 a[i] > x
2. 那么收益 = b[i] - a[i] + x
3. 如果 a[i] < x, 收益 = b[i] + a[i] - x
4. 所以反过来，在x处的收益  = 它右边 b[i] - a[i] + x, 它左边 b[i] + a[i] - x
5. 还需要计算上次烟火的收益，假设上次如果在位置y, 如果 abs(y - x) <= d * (t[i] - t[j])
6. 那么y的收益，就可以增加到x中（且这个和本次烟火时没有关系的）
7. 涉及到区间查询和单点更新
8. 按照烟火发射时间进行升序处理（有相同时刻的烟火，这个还需要特别考虑）
9. 