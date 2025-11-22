# Problem Description

Wherever the destination is, whoever we meet, let's render this song together.

On a Cartesian coordinate plane lies a rectangular stage of size w × h, represented by a rectangle with corners (0, 0), (w, 0), (w, h) and (0, h). It can be seen that no collisions will happen before one enters the stage.

On the sides of the stage stand n dancers. The i-th of them falls into one of the following groups:

- **Vertical**: stands at (xi, 0), moves in positive y direction (upwards);
- **Horizontal**: stands at (0, yi), moves in positive x direction (rightwards).

According to choreography, the i-th dancer should stand still for the first ti milliseconds, and then start moving in the specified direction at 1 unit per millisecond, until another border is reached. It is guaranteed that no two dancers have the same group, position and waiting time at the same time.

When two dancers collide (i.e. are on the same point at some time when both of them are moving), they immediately exchange their moving directions and go on.

Dancers stop when a border of the stage is reached. Find out every dancer's stopping position.

## Input

The first line of input contains three space-separated positive integers n, w and h (1 ≤ n ≤ 100 000, 2 ≤ w, h ≤ 100 000) — the number of dancers and the width and height of the stage, respectively.

The following n lines each describes a dancer: the i-th among them contains three space-separated integers gi, pi, and ti (1 ≤ gi ≤ 2, 1 ≤ pi ≤ 99 999, 0 ≤ ti ≤ 100 000), describing a dancer's group gi (gi = 1 — vertical, gi = 2 — horizontal), position, and waiting time. If gi = 1 then pi = xi; otherwise pi = yi. It's guaranteed that 1 ≤ xi ≤ w - 1 and 1 ≤ yi ≤ h - 1. It is guaranteed that no two dancers have the same group, position and waiting time at the same time.

## Output

Output n lines, the i-th of which contains two space-separated integers (xi, yi) — the stopping position of the i-th dancer in the input.

## Examples

### Example 1

**Input:**

```text
8 10 8
1 1 10
1 4 13
1 7 1
1 8 2
2 2 0
2 5 14
2 6 0
2 6 1
```

**Output:**

```text
4 8
10 5
8 8
10 6
10 2
1 8
7 8
10 6
```

### Example 2

**Input:**

```text
3 2 3
1 1 2
2 1 1
1 1 5
```

**Output:**

```text
1 3
2 1
1 3
```

## Note

The first example corresponds to the initial setup in the legend, and the tracks of dancers are marked with different colours in the following figure.

In the second example, no dancers collide.


### ideas
1. 考虑在y轴上的人， i第一次发生碰撞后，他会晚上运动；然后他有可能碰到第二个人(这个人是从左往右移动过来的，所以有可能是发生碰撞的某个在x轴上的人); 然后依次进行
2. 但是如果不考虑换人的问题，可以当作是某个x轴上的j,一直在往上运动
3. 单独考虑碰撞点和次数是不变的
4. 那么这样子就可以计算出发生碰撞的位置和次数 （这个次数大体应该是n+m）的
5. 从最右边的点开始