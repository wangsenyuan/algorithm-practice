# Problem G

You are given a tree consisting of $n$ vertices. A number is written on each vertex; the number on vertex $i$ is equal to $a_i$.

Let's denote the function $g(x,y)$ as the greatest common divisor of the numbers written on the vertices belonging to the simple path from vertex $x$ to vertex $y$ (including these two vertices).

For every integer from $1$ to $2 \cdot 10^5$ you have to count the number of pairs $(x,y)$ $(1 \leq x \leq y \leq n)$ such that $g(x,y)$ is equal to this number.

## Input

The first line contains one integer $n$ — the number of vertices $(1 \leq n \leq 2 \cdot 10^5)$.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(1 \leq a_i \leq 2 \cdot 10^5)$ — the numbers written on vertices.

Then $n-1$ lines follow, each containing two integers $x$ and $y$ $(1 \leq x, y \leq n, x \neq y)$ denoting an edge connecting vertex $x$ with vertex $y$. It is guaranteed that these edges form a tree.

## Output

For every integer $i$ from $1$ to $2 \cdot 10^5$ do the following: if there is no pair $(x,y)$ such that $x \leq y$ and $g(x,y) = i$, don't output anything. Otherwise output two integers: $i$ and the number of aforementioned pairs. You have to consider the values of $i$ in ascending order.

See the examples for better understanding.

## Examples

### Example 1

**Input:**
```
3
1 2 3
1 2
2 3
```

**Output:**
```
1 4
2 1
3 1
```

### Example 2

**Input:**
```
6
1 2 4 8 16 32
1 6
6 3
3 4
4 2
6 5
```

**Output:**
```
1 6
2 5
4 6
8 1
16 2
32 1
```

### Example 3

**Input:**
```
4
9 16 144 6
1 3
2 3
4 3
```

**Output:**
```
1 1
2 1
3 1
6 2
9 2
16 2
144 1
```