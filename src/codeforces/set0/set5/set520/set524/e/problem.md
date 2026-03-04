# Problem

Polycarpus has a chessboard of size $n \times m$, where $k$ rooks are placed. Polycarpus hasn't yet invented the rules of the game he will play. However, he has already allocated $q$ rectangular areas of special strategic importance on the board — they must be protected well.

According to Polycarpus, a rectangular area of the board is **well protected** if all its vacant squares can be attacked by the rooks that stand on this area. The rooks on the rest of the board do not affect the area's defense. The position of the rooks is fixed and cannot be changed. We remind you that a rook attacks the squares located on the same vertical or horizontal line with it if there are no other pieces between the square and the rook.

Help Polycarpus determine whether all strategically important areas are protected.

## Input

The first line contains four integers $n$, $m$, $k$ and $q$ ($1 \le n, m \le 100\,000$, $1 \le k, q \le 200\,000$) — the sizes of the board, the number of rooks, and the number of strategically important areas.

We consider that the cells of the board are numbered by integers from $1$ to $n$ horizontally and from $1$ to $m$ vertically.

The next $k$ lines contain pairs of integers `x y`, describing the positions of the rooks ($1 \le x \le n$, $1 \le y \le m$). It is guaranteed that all the rooks are in distinct squares.

The next $q$ lines describe the strategically important areas as groups of four integers `x1 y1 x2 y2` ($1 \le x_1 \le x_2 \le n$, $1 \le y_1 \le y_2 \le m$). The corresponding rectangular area consists of cells $(x, y)$ for which $x_1 \le x \le x_2$, $y_1 \le y \le y_2$. Strategically important areas can intersect or coincide.

## Output

Print $q$ lines. For each strategically important area, print `YES` if it is well defended and `NO` otherwise.

## Examples

**Input**

```text
4 3 3 3
1 1
3 2
2 3
2 3 2 3
2 1 3 3
1 2 2 3
```

**Output**

```text
YES
YES
NO
```


### ideas
1. 给定一个区域内，看这个区域内，如果所有行都有rock，或者所有列都有rock，有没有第三种情况？
2. 没有。就这两种情况
3. 但是只能由这个区间内的rock。这个就麻烦了。
4. 感觉要用持久化树。考虑行维度的持久化，在tr[b].get(l, r) > tr[t-1].get(l, r)？
5. 不对。假设有一行没有变，但是它不是最小值，那么及时这个值返回true，也不对
6. get(l, r) 表示这个区间内每一列的设置的最后一个的最小值 get(l, r) >= b (这个是成立的)
7. 但是怎么处理呢？
8. 如果是行的过程，需要按照列的方向处理