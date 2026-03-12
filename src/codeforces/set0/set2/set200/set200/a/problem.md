- **Given**: a theater with `n` rows and `m` seats, and `k` people, each with a favorite seat.
- **Process people in queue order**:
  - If their favorite seat is free, they take it.
  - Otherwise, among all free seats, they choose one with:
    1. minimal Manhattan distance \(|x_1 - x_2| + |y_1 - y_2|\),
    2. then minimal row index,
    3. then minimal seat index.
- **Output**: the final seat coordinates for each person, in the same order as in the queue.

## Input
The first input line contains three integers n, m, k (1 ≤ n, m ≤ 2000, 1 ≤ k ≤ min(n·m, 105) — the number of rows in the room, the number of seats in each row and the number of people in the line, correspondingly. Each of the next k lines contains two integers xi, yi (1 ≤ xi ≤ n, 1 ≤ yi ≤ m) — the coordinates of the seat each person has chosen. Numbers on the same line are separated by a space. The pairs of coordinates are located in the order, in which people stand in the line, starting from the head (the first person in the line who stands in front of the box office) to the tail (the last person in the line).

## Output
Print k lines, each containing a pair of integers. Print on the i-th line xi, yi — the coordinates of the seat, for which the person who stands i-th in the line will buy the ticket.

## Examples

### Example 1

Input:

```text
3 4 6
1 1
1 1
1 1
1 2
1 3
1 3
```

Output:

```text
1 1
1 2
2 1
1 3
1 4
2 3
```

### Example 2

Input:

```text
4 3 12
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
```

Output:

```text
2 2
1 2
2 1
2 3
3 2
1 1
1 3
3 1
3 3
4 2
4 1
4 3
```

---


## ideas
1. 对于每个(x, y), 需要知道离它最近的空的位置
2. 有4个方向的位置需要考虑，左上, 右上, 左下，右下
3. 左上的区域，应该找(x + y) 最大的空格位置（还没有被占的位置）
4. 因为 n * m 已经是1e5了，所以这里必须得是一个lgn的查找
5. 1, 4, 8, 12, 16, ...
6. 假设都堆在一起，那么最坏的检查 = sqrt(k) = sqrt(1e5) = 300?
7. 也就是300 * k ~ 3 * 1e8 (好像可以通过)

### editorial

In this problem were given the field, which size is n × m, and k queries. Each query is the cell of the field. You had to find closest free cell for given one (there was used manhattan metric). Than found cell was marked as used. The time complexity of the solution is supposed to be . Firstly, we show the main idea, than we show why solution has time complexity .

First of all, if n > m, than rotate matrix by 90 degrees (than we explain purpose of this transformation). We will have two matrices, size of each is n × m. In this matrices we will maintain for each cell the nearest free one to the left and to the right. Suppose we have query — cell (x, y). Iterate d — how much rows we will go above or below. Suppose we fix d, than look at the row x - d, for example (also we should do same things for row x + d). Lets find in this row nearest free cells to the left and to the right of (x - d, y) by using our matrices and try to update answer. When d will be greater than current found answer we will stop, because answer won't update in this case. After we find the answer, we have to change value in some cells of our matrices. It can be done by using structure DSU for each row.

Lets show why this solution has time complexity . Suppose all queries are equal, for example each query is (x, y). Than if field is big enough, all points will be placed in form of square, which is rotated by 45 degrees and have center in point (x, y). Also the side of the square will have length . Than diagonal will have length  too. It means that we see  rows in each query and we do O(1) operations in each row. If the square doesn't fit into the field, than field is too thin verticaly or horizontaly. So we will have figure looks like rectangle, and one side of this rectangle will be less than . In this case rotate field so, than least side of rectangle means rows. Than we will see not greater than  rows in each query and will perform not greater than O(1) operations in each row.

This problem is supposed to be the hardest problem of the contest.

