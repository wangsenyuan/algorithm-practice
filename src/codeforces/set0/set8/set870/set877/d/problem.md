# Problem D: Energy Drinks

## Problem Description

Olya loves energy drinks. She loves them so much that her room is full of empty cans from energy drinks.

Formally, her room can be represented as a field of n × m cells, each cell of which is empty or littered with cans.

Olya drank a lot of energy drink, so now she can run k meters per second. Each second she chooses one of the four directions (up, down, left or right) and runs from 1 to k meters in this direction. Of course, she can only run through empty cells.

Now Olya needs to get from cell (x₁, y₁) to cell (x₂, y₂). How many seconds will it take her if she moves optimally?

It's guaranteed that cells (x₁, y₁) and (x₂, y₂) are empty. These cells can coincide.

## Input

The first line contains three integers **n**, **m** and **k** (1 ≤ n, m, k ≤ 1000) — the sizes of the room and Olya's speed.

Then **n** lines follow containing **m** characters each, the i-th of them contains on j-th position:
- `#` if the cell (i, j) is littered with cans
- `.` otherwise

The last line contains four integers **x₁**, **y₁**, **x₂**, **y₂** (1 ≤ x₁, x₂ ≤ n, 1 ≤ y₁, y₂ ≤ m) — the coordinates of the first and the last cells.

## Output

Print a single integer — the minimum time it will take Olya to get from (x₁, y₁) to (x₂, y₂).

If it's impossible to get from (x₁, y₁) to (x₂, y₂), print **-1**.

## Examples

### Example 1
**Input:**
```
3 4 4
....
###.
....
1 1 3 1
```

**Output:**
```
3
```

**Explanation:** Olya should run 3 meters to the right in the first second, 2 meters down in the second second and 3 meters to the left in the third second.

### Example 2
**Input:**
```
3 4 1
....
###.
....
1 1 3 1
```

**Output:**
```
8
```

**Explanation:** Olya should run to the right for 3 seconds, then down for 2 seconds and then to the left for 3 seconds.

### Example 3
**Input:**
```
2 2 1
.#
#.
1 1 2 2
```

**Output:**
```
-1
```

**Explanation:** It's impossible to reach the destination due to blocked paths.

## Note

Olya does not recommend drinking energy drinks and generally believes that this is bad.


## ideas
1. 从一个位置，有4 * k个位置可以跳转
2. 如果直接跑，那么就是 n * n * k 的复杂性， 显然不大行
3. 用segment tree应该可以，range update + range query
4. 有没有更简单的方案？