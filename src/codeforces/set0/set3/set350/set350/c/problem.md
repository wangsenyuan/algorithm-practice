You've got a robot, its task is destroying bombs on a square plane. Specifically, the square plane contains **n** bombs, the *i*-th bomb is at point with coordinates \((x_i, y_i)\). We know that no two bombs are at the same point and that no bomb is at point with coordinates \((0, 0)\). Initially, the robot is at point with coordinates \((0, 0)\). Also, let's mark the robot's current position as \((x, y)\).

In order to destroy all the bombs, the robot can perform three types of operations:

- **Operation 1:**
  - Format: `1 k dir`
  - To perform the operation, the robot has to move in direction `dir` `k` (\(k \geq 1\)) times.
  - There are only 4 directions the robot can move in: `R`, `L`, `U`, `D`.
  - During one move, the robot can move from the current point to one of the following points:
    - `(x + 1, y)` (R)
    - `(x - 1, y)` (L)
    - `(x, y + 1)` (U)
    - `(x, y - 1)` (D)
  - It is forbidden to move from point `(x, y)` if at least one point on the path (besides the destination point) contains a bomb.

- **Operation 2:**
  - Format: `2`
  - To perform the operation, the robot has to pick a bomb at point `(x, y)` and put it in a special container.
  - The robot can carry the bomb from any point to any other point.
  - The operation cannot be performed if point `(x, y)` has no bomb.
  - It is forbidden to pick a bomb if the robot already has a bomb in its container.

- **Operation 3:**
  - Format: `3`
  - To perform the operation, the robot has to take a bomb out of the container and destroy it.
  - You are allowed to perform this operation only if the robot is at point `(0, 0)`.
  - It is forbidden to perform the operation if the container has no bomb.

---

**Help the robot and find the shortest possible sequence of operations he can perform to destroy all bombs on the coordinate plane.**

---

### Input

- The first line contains a single integer `n` (\(1 \leq n \leq 10^5\)) — the number of bombs on the coordinate plane.
- Next `n` lines contain two integers each. The *i*-th line contains numbers `(x_i, y_i)` (\(-10^9 \leq x_i, y_i \leq 10^9\)) — the coordinates of the *i*-th bomb.
- It is guaranteed that no two bombs are located at the same point and no bomb is at point `(0, 0)`.

### Output

- In a single line print a single integer `k` — the minimum number of operations needed to destroy all bombs.
- On the next lines print the descriptions of these `k` operations.
- If there are multiple sequences, you can print any of them.
- It is guaranteed that there is a solution where `k ≤ 10^6`.

### ideas
1. 每个robot都需要至少 2 * move + 1 + 1 次操作
2. 移动到位置(x, y)，然后pick，然后返回(0, 0)，然后销毁
3. 那么就必须保证路径上没有其他的bomb，所有要从近到远的处理