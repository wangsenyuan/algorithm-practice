# Problem C

Yet another round on DecoForces is coming! Grandpa Maks wanted to participate in it but someone has stolen his precious sofa! And how can one perform well with such a major loss?

Fortunately, the thief had left a note for Grandpa Maks. This note got Maks to the sofa storehouse. Still he had no idea which sofa belongs to him as they all looked the same!

The storehouse is represented as matrix $n \times m$. Every sofa takes two neighbouring by some side cells. No cell is covered by more than one sofa. There can be empty cells.

Sofa A is standing to the left of sofa B if there exist two such cells a and b that $x_a < x_b$, a is covered by A and b is covered by B. Sofa A is standing to the top of sofa B if there exist two such cells a and b that $y_a < y_b$, a is covered by A and b is covered by B. Right and bottom conditions are declared the same way.

Note that in all conditions $A \neq B$. Also some sofa A can be both to the top of another sofa B and to the bottom of it. The same is for left and right conditions.

The note also stated that there are $cnt_l$ sofas to the left of Grandpa Maks's sofa, $cnt_r$ — to the right, $cnt_t$ — to the top and $cnt_b$ — to the bottom.

Grandpa Maks asks you to help him to identify his sofa. It is guaranteed that there is no more than one sofa of given conditions.

Output the number of Grandpa Maks's sofa. If there is no such sofa that all the conditions are met for it then output -1.

## Input

The first line contains one integer number $d$ ($1 \leq d \leq 10^5$) — the number of sofas in the storehouse.

The second line contains two integer numbers $n, m$ ($1 \leq n, m \leq 10^5$) — the size of the storehouse.

Next $d$ lines contains four integer numbers $x_1, y_1, x_2, y_2$ ($1 \leq x_1, x_2 \leq n$, $1 \leq y_1, y_2 \leq m$) — coordinates of the $i$-th sofa. It is guaranteed that cells $(x_1, y_1)$ and $(x_2, y_2)$ have common side, $(x_1, y_1) \neq (x_2, y_2)$ and no cell is covered by more than one sofa.

The last line contains four integer numbers $cnt_l, cnt_r, cnt_t, cnt_b$ ($0 \leq cnt_l, cnt_r, cnt_t, cnt_b \leq d - 1$).

## Output

Print the number of the sofa for which all the conditions are met. Sofas are numbered 1 through $d$ as given in input. If there is no such sofa then print -1.

## Examples

### Example 1

**Input:**

```text
2
3 2
3 1 3 2
1 2 2 2
1 0 0 1
```

**Output:**

```text
1
```

### Example 2

**Input:**

```text
3
10 10
1 2 1 1
5 5 6 5
6 4 5 4
2 1 2 0
```

**Output:**

```text
2
```

### Example 3

**Input:**

```text
2
2 2
2 1 1 1
1 2 2 2
1 0 0 0
```

**Output:**

```text
-1
```

## Note

Let's consider the second example.

The first sofa has 0 to its left, 2 sofas to its right ($(1, 1)$ is to the left of both $(5, 5)$ and $(5, 4)$), 0 to its top and 2 to its bottom (both 2nd and 3rd sofas are below).

The second sofa has $cnt_l = 2$, $cnt_r = 1$, $cnt_t = 2$ and $cnt_b = 0$.

The third sofa has $cnt_l = 2$, $cnt_r = 1$, $cnt_t = 1$ and $cnt_b = 1$.

So the second one corresponds to the given conditions.

In the third example

The first sofa has $cnt_l = 1$, $cnt_r = 1$, $cnt_t = 0$ and $cnt_b = 1$.

The second sofa has $cnt_l = 1$, $cnt_r = 1$, $cnt_t = 1$ and $cnt_b = 0$.

And there is no sofa with the set $(1, 0, 0, 0)$ so the answer is -1.


### ideas
1. 在给定一个W时，计算有多少个在它的左边、上边、右边、下边
2. 左边是说A得x坐标，都必须在B的x坐标左边， max(xa) < min(xb)
3. 首先将sofa规整处理(x1 <= x2), (y1 <= y2)
4. 然后按照x1升序处理（从左到右）
5. 同时更新y方向的计数
6. 