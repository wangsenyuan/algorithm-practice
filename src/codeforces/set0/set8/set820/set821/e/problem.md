Okabe likes to take walks but knows that spies from the Organization could be anywhere; that's why he wants to know how
many different walks he can take in his city safely. Okabe's city can be represented as all points $(x, y)$ such that $x$
and $y$ are non-negative. Okabe starts at the origin (point $(0, 0)$), and needs to reach the point $(k, 0)$. If Okabe is
currently at the point $(x, y)$, in one step he can go to $(x + 1, y + 1)$, $(x + 1, y)$, or $(x + 1, y - 1)$.

Additionally, there are $n$ horizontal line segments, the $i$-th of which goes from $x = a_i$ to $x = b_i$ inclusive, and
is at $y = c_i$. It is guaranteed that $a_1 = 0$, $a_n \leq k \leq b_n$, and $a_i = b_{i-1}$ for $2 \leq i \leq n$. The
$i$-th line segment forces Okabe to walk with $y$-value in the range $0 \leq y \leq c_i$ when his $x$ value satisfies
$a_i \leq x \leq b_i$, or else he might be spied on. This also means he is required to be under two line segments when one
segment ends and another begins.

Okabe now wants to know how many walks there are from the origin to the point $(k, 0)$ satisfying these conditions, modulo
$10^9 + 7$.

## Input

The first line of input contains the integers $n$ and $k$ ($1 \leq n \leq 100$, $1 \leq k \leq 10^{18}$) — the number of
segments and the destination $x$ coordinate.

The next $n$ lines contain three space-separated integers $a_i$, $b_i$, and $c_i$ ($0 \leq a_i < b_i \leq 10^{18}$,
$0 \leq c_i \leq 15$) — the left and right ends of a segment, and its $y$ coordinate.

It is guaranteed that $a_1 = 0$, $a_n \leq k \leq b_n$, and $a_i = b_{i-1}$ for $2 \leq i \leq n$.

## Output

Print the number of walks satisfying the conditions, modulo $1000000007$ ($10^9 + 7$).

## Examples

### Example 1

**Input**
```
1 3
0 3 3
```

**Output**
```
4
```

### Example 2

**Input**
```
2 6
0 3 0
3 10 2
```

**Output**
```
4
```
