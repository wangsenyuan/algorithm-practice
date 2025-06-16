# Problem C: Maximum Points

## Problem Description

You are playing a video game and you have just reached the bonus level, where the only possible goal is to score as many points as possible. Being a perfectionist, you've decided that you won't leave this level until you've gained the maximum possible number of points there.

### Game Mechanics

The bonus level consists of:
- n small platforms placed in a line, numbered from 1 to n from left to right
- (n - 1) bridges connecting adjacent platforms
- Each bridge has a known limit on how many times it can be crossed before collapsing

### Player Actions

1. Select a starting platform for the hero
2. Move the hero across platforms using undestroyed bridges
3. The level ends when the hero reaches a platform with no undestroyed bridges
4. Points are calculated as the number of transitions between platforms

**Important Note:** Once the hero starts moving in a direction across a bridge, they must continue in that direction until reaching a platform.

## Task

Find the maximum number of points possible to ensure no one can beat your record.

## Input Format

- First line: A single integer n (2 ≤ n ≤ 10^5) — number of platforms
- Second line: (n - 1) integers ai (1 ≤ ai ≤ 10^9, 1 ≤ i < n) — maximum number of times bridge i can be crossed

## Output Format

A single integer — the maximum possible points achievable on the bonus level.

## Technical Notes

- For C++ users: Use cin/cout or %I64d specifier for 64-bit integers
- Do not use %lld specifier

## ideas
1. dp[i][0] 表示从i开始，并且停止在i这里
2. dp[i][1] 表示从j(j >= i)开始，并且停止在i这里
3. dp[i][2] 表示从i开始，但是停止在(j >= i)处
4. dp[i][0] = a[i] / 2 * 2 + dp[i+1][0]
5. dp[i][1] = dp[i+1][1] + (a[i] or a[i] - 1), dp[i][0]
