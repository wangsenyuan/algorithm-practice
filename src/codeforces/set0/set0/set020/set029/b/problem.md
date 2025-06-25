# Traffic Light Problem

## Problem Description

A car moves from point A to point B at speed $v$ meters per second. The action takes place on the X-axis. At the distance $d$ meters from A there are traffic lights. Starting from time 0, for the first $g$ seconds the green light is on, then for the following $r$ seconds the red light is on, then again the green light is on for the $g$ seconds, and so on.

The car can be instantly accelerated from 0 to $v$ and vice versa, can instantly slow down from the $v$ to 0. Consider that it passes the traffic lights at the green light instantly. If the car approaches the traffic lights at the moment when the red light has just turned on, it doesn't have time to pass it. But if it approaches the traffic lights at the moment when the green light has just turned on, it can move. The car leaves point A at the time 0.

**What is the minimum time for the car to get from point A to point B without breaking the traffic rules?**

## Input

The first line contains integers $l$, $d$, $v$, $g$, $r$ ($1 \leq l, d, v, g, r \leq 1000$, $d < l$) — the distance between A and B (in meters), the distance from A to the traffic lights, car's speed, the duration of green light and the duration of red light.

## Output

Output a single number — the minimum time that the car needs to get from point A to point B. Your output must have relative or absolute error less than $10^{-6}$.


### ideas
1. 如果 d % (g + r) < g => l / d
2. else d % (g + r) >= g 也就是说，车辆在到达d的时候需要等待
3. d / v + (g + r) - d % (g + r), 现在是绿灯 + (l - d) / v
4. 要转换成t
5. 