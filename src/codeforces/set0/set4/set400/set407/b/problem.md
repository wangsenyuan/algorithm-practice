# Codeforces 407B - Long Path

## Problem Description

One day, little Vasya found himself in a maze consisting of `(n + 1)` rooms, numbered from `1` to `(n + 1)`. Initially, Vasya is at the first room and to get out of the maze, he needs to get to the `(n + 1)`-th one.

### Maze Structure

The maze is organized as follows. Each room of the maze has two one-way portals. Let's consider room number `i` (`1 ≤ i ≤ n`):
- Someone can use the **first portal** to move from it to room number `(i + 1)`
- Someone can use the **second portal** to move from it to room number `pi`, where `1 ≤ pi ≤ i`

### Vasya's Strategy

In order not to get lost, Vasya decided to act as follows:

1. Each time Vasya enters some room, he paints a cross on its ceiling
2. Initially, Vasya paints a cross at the ceiling of room 1
3. Let's assume that Vasya is in room `i` and has already painted a cross on its ceiling. Then:
   - If the ceiling now contains an **odd number** of crosses, Vasya uses the **second portal** (it leads to room `pi`)
   - Otherwise, Vasya uses the **first portal**

**Task:** Help Vasya determine the number of times he needs to use portals to get to room `(n + 1)` in the end.

## Input

- The first line contains integer `n` (`1 ≤ n ≤ 10³`) — the number of rooms
- The second line contains `n` integers `pi` (`1 ≤ pi ≤ i`). Each `pi` denotes the number of the room that someone can reach if he will use the second portal in the `i`-th room

## Output

Print a single number — the number of portal moves the boy needs to go out of the maze. As the number can be rather large, print it modulo `1000000007` (`10⁹ + 7`).


## ideas
1. 有点混乱
2. 理解错了，就是第一次进入一个房间， boy一定要移动到p[i],
3. 然后第二次进入这个房间时，他才能移动到i+1
4. 每次向后移动的时候，除了(p[i] = i)的情况，j = p[i], 位置j处已经有了偶数个cross
5. 那么可以当作是在空的走了一遍