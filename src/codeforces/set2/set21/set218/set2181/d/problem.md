# Problem D

The construction of the doorway for the Nonsense Engineering and Research Convention was delegated to one of the future attendees, who decided on a multi-layered sliding door design.

Each layer can be described as a horizontal interval, bounded by solid walls on the left and right, containing a number of sliding doors of fixed lengths. Within a layer, each door can move independently to the left or right, as long as it does not overlap other doors or the walls. All layers are parallel and stacked vertically.

After construction, the organizers noticed a problem: it is difficult to fully open the door, and since a large number of attendees are expected, they need to create the largest possible opening to allow everyone to pass through freely.

The size of the opening is defined as the total length of horizontal intervals such that, at every point of such an interval and in every layer, there is neither a door nor a wall. Your task is to determine the largest possible opening, given the doors' layout.

## Input

The first line contains an integer $n$ ($1 \leq n \leq 100000$) — the number of layers of the door.

Each of the next $n$ lines starts with three integers $k_i$, $x_{i,1}$, $x_{i,2}$ ($0 \leq k_i \leq 300000$; $0 \leq x_{i,1} < x_{i,2} \leq 10^9$) — the number of sliding doors on that layer and the $x$-coordinates $x_{i,1}$ and $x_{i,2}$ of the walls on that layer. There is a wall at $x_{i,1}$ and a wall at $x_{i,2}$; all positions with $x < x_{i,1}$ or $x > x_{i,2}$ are blocked by walls.

They are followed by $k_i$ integers $l_{i,1}, \ldots, l_{i,k_i}$ ($1 \leq l_{i,j}$; $\sum_{j=1}^{k_i} l_{i,j} \leq x_{i,2} - x_{i,1}$) — the lengths of the sliding doors on that layer given in order from the leftmost door to the rightmost.

It is guaranteed that $\sum_{i=1}^{n} k_i \leq 300000$.

## Output

Output a single integer — the size of the largest possible opening that can be achieved by moving the sliding doors on each layer.

## Examples

**Input:**

```text
2
2 2 11 3 2
3 4 12 1 1 2
```

**Output:**

```text
4
```

**Input:**

```text
2
2 0 7 2 4
1 4 9 4
```

**Output:**

```text
0
```

## Note

This illustration shows a solution for the first example. Walls are filled with black color, doors are filled with various shades of grey, the opening is white. When first doors on each layer are shifted to the left and the rest of the doors to the right, we get the largest opening of 4.


### ideas
1. 首先把所有的door都移动到右边，然后根据当前的位置（最右边的左墙）
2. 然后从队列中找出，它的左端点 + 长度，最小的那个
3. 然后更新，这段后面的所有位置的左端点（也可以不更新，好像可以一开始就放入）
4. 貌似可以搞