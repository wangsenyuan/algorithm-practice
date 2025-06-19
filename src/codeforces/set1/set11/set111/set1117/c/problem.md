# Ship Navigation Problem

## Problem Description

You are a captain of a ship. Initially you are standing at point (x₁, y₁) (obviously, all positions in the sea can be described by cartesian plane) and you want to travel to a point (x₂, y₂).

You know the weather forecast — the string s of length n, consisting only of letters U, D, L and R. The letter corresponds to a direction of wind. Moreover, the forecast is periodic, e.g. the first day wind blows to the side s₁, the second day — s₂, the n-th day — sₙ and (n+1)-th day — s₁ again and so on.

## Ship Movement Rules

Ship coordinates change the following way:

- **U (Up)**: Ship moves from (x, y) to (x, y+1)
- **D (Down)**: Ship moves from (x, y) to (x, y-1)  
- **L (Left)**: Ship moves from (x, y) to (x-1, y)
- **R (Right)**: Ship moves from (x, y) to (x+1, y)

## Ship Control

The ship can also either go one of the four directions or stay in place each day. If it goes then it's exactly 1 unit of distance. Transpositions of the ship and the wind add up. If the ship stays in place, then only the direction of wind counts.

### Examples:
- If wind blows direction U and ship moves direction L: from point (x, y) → (x-1, y+1)
- If wind blows direction U and ship moves direction U: from point (x, y) → (x, y+2)

## Task

Determine the minimal number of days required for the ship to reach the point (x₂, y₂).

## Input

- **Line 1**: Two integers x₁, y₁ (0 ≤ x₁, y₁ ≤ 10⁹) — the initial coordinates of the ship
- **Line 2**: Two integers x₂, y₂ (0 ≤ x₂, y₂ ≤ 10⁹) — the coordinates of the destination point
- **Line 3**: Single integer n (1 ≤ n ≤ 10⁵) — the length of the string s
- **Line 4**: String s itself, consisting only of letters U, D, L and R

**Note**: It is guaranteed that the initial coordinates and destination point coordinates are different.

## Output

The only line should contain the minimal number of days required for the ship to reach the point (x₂, y₂).

If it's impossible then print "-1".

## ideas
1. 确实符合二分的条件。假设可以在x天到达，那么肯定可以在x+1天到达（因为可以在x+1天，选择和风相对的操作，从而抵消风的效果）
2. 假设要在x天到达，那么要怎么处理呢？肯定不能模拟这个过程吧
3. 当x很大的时候，比如 1e9， 那么久太慢了
4. 但是似乎是可以分开x方向和y水平方向来考虑吗？
5. 似乎是可以的
6. 只考虑x方向，那么在顺风的时候，就要往前划（似乎总是往前划，顺风的时候加速，逆风的时候，保留在原地）
7. 那么就是在dx中间，有多少个顺风的情况，和逆风的情况，还有无关的情况（上下）
8. 两个分开考虑似乎不大对
9. 考虑两个点的相对位置是左下/右上
10. 不考虑划动的情况，单独考虑风的作用，那么经过m秒后，船的位置是确定的
11. 假设现在的位置是x0, y0, 它要往目标划，dx + dy <= m