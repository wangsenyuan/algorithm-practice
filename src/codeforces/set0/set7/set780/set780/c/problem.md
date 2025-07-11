# Problem C: Balloon Coloring

## Problem Description

Andryusha goes through a park each day. The squares and paths between them look boring to Andryusha, so he decided to decorate them.

The park consists of **n** squares connected with **(n - 1)** bidirectional paths in such a way that any square is reachable from any other using these paths. Andryusha decided to hang a colored balloon at each of the squares. The balloons' colors are described by positive integers, starting from 1.

In order to make the park varicolored, Andryusha wants to choose the colors in a special way. More precisely, he wants to use such colors that if **a**, **b** and **c** are distinct squares where:
- **a** and **b** have a direct path between them, and
- **b** and **c** have a direct path between them

Then the balloon colors on these three squares must be **distinct**.

Andryusha wants to use as few different colors as possible. Help him to choose the colors!

## Input

- The first line contains a single integer **n** (3 ≤ n ≤ 2·10⁵) — the number of squares in the park.
- Each of the next **(n - 1)** lines contains two integers **x** and **y** (1 ≤ x, y ≤ n) — the indices of two squares directly connected by a path.

**Note:** It is guaranteed that any square is reachable from any other using the paths.

## Output

- In the first line, print a single integer **k** — the minimum number of colors Andryusha has to use.
- In the second line, print **n** integers, where the **i-th** integer should be equal to the balloon color on the **i-th** square. Each of these numbers should be within the range from 1 to **k**.


## ideas
1. 假设u的子树都安排好了
2. 现在给u进行涂色（这个策略不行，因为子树之间也有限制）
3. 找到最大deg的节点，那么颜色不会少于deg + 1
4. 假设第一圈涂好颜色了，然后涂第二圈
5. 那么这一圈圈的处理，是可以的