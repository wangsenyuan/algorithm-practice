# Problem Description

The Cartesian coordinate system is set in the sky. There you can see $n$ stars, the $i$-th has coordinates $(x_i, y_i)$, a maximum brightness $c$, equal for all stars, and an initial brightness $s_i$ ($0 \leq s_i \leq c$).

Over time the stars twinkle. At moment $0$ the $i$-th star has brightness $s_i$. Let at moment $t$ some star has brightness $x$. Then at moment $(t + 1)$ this star will have brightness $x + 1$, if $x + 1 \leq c$, and $0$, otherwise.

You want to look at the sky $q$ times. In the $i$-th time you will look at the moment $t_i$ and you will see a rectangle with sides parallel to the coordinate axes, the lower left corner has coordinates $(x_{1i}, y_{1i})$ and the upper right — $(x_{2i}, y_{2i})$. For each view, you want to know the total brightness of the stars lying in the viewed rectangle.

A star lies in a rectangle if it lies on its border or lies strictly inside it.

## Input

The first line contains three integers $n$, $q$, $c$ ($1 \leq n, q \leq 10^5$, $1 \leq c \leq 10$) — the number of the stars, the number of the views and the maximum brightness of the stars.

The next $n$ lines contain the stars description. The $i$-th from these lines contains three integers $x_i$, $y_i$, $s_i$ ($1 \leq x_i, y_i \leq 100$, $0 \leq s_i \leq c \leq 10$) — the coordinates of $i$-th star and its initial brightness.

The next $q$ lines contain the views description. The $i$-th from these lines contains five integers $t_i$, $x_{1i}$, $y_{1i}$, $x_{2i}$, $y_{2i}$ ($0 \leq t_i \leq 10^9$, $1 \leq x_{1i} < x_{2i} \leq 100$, $1 \leq y_{1i} < y_{2i} \leq 100$) — the moment of the $i$-th view and the coordinates of the viewed rectangle.

## Output

For each view print the total brightness of the viewed stars.

## Examples

### Example 1

**Input:**

```text
2 3 3
1 1 1
3 2 0
2 1 1 2 2
0 2 1 4 5
5 1 1 5 5
```

**Output:**

```text
3
0
3
```

### Example 2

**Input:**

```text
3 4 5
1 1 2
2 3 0
3 3 1
0 1 1 100 100
1 2 2 4 4
2 2 1 4 7
1 50 50 51 51
```

**Output:**

```text
3
3
5
0
```

## Note

Let's consider the first example.

At the first view, you can see only the first star. At moment $2$ its brightness is $3$, so the answer is $3$.

At the second view, you can see only the second star. At moment $0$ its brightness is $0$, so the answer is $0$.

At the third view, you can see both stars. At moment $5$ brightness of the first is $2$, and brightness of the second is $1$, so the answer is $3$.
