# Problem B: Arthur and Alexander

Arthur and Alexander are number busters. Today they've got a competition.

## Problem Description

Arthur took a group of four integers $a, b, w, x$ ($0 \leq b < w$, $0 < x < w$) and Alexander took integer $c$. Arthur and Alexander use distinct approaches to number bustings.

- **Alexander** is just a regular guy. Each second, he subtracts one from his number. In other words, he performs the assignment: $c = c - 1$.

- **Arthur** is a sophisticated guy. Each second Arthur performs a complex operation, described as follows:
  - If $b \geq x$: perform the assignment $b = b - x$
  - If $b < x$: then perform two consecutive assignments $a = a - 1$; $b = w - (x - b)$

You've got numbers $a, b, w, x, c$. Determine when Alexander gets ahead of Arthur if both guys start performing the operations at the same time. Assume that Alexander got ahead of Arthur if $c \leq a$.

## Input

The first line contains integers $a, b, w, x, c$ ($1 \leq a \leq 2 \cdot 10^9$, $1 \leq w \leq 1000$, $0 \leq b < w$, $0 < x < w$, $1 \leq c \leq 2 \cdot 10^9$).

## Output

Print a single integer — the minimum time in seconds Alexander needs to get ahead of Arthur. You can prove that the described situation always occurs within the problem's limits.

## Examples

### Example 1
**Input:**
```
4 2 3 1 6
```
**Output:**
```
2
```

### Example 2
**Input:**
```
4 2 3 1 7
```
**Output:**
```
4
```

### Example 3
**Input:**
```
1 2 3 2 6
```
**Output:**
```
13
```

### Example 4
**Input:**
```
1 1 2 1 1
```
**Output:**
```
0
```


### ideas
1. b < w, 且 b - x 直到 b < x => b % x
2. b = w - (x - b) 的时候，保证了b不会超过w
3. 所以，b肯定会进入一个循环
4. b < x
5. w - (x - b) < x => w + b < 2 * x 有没有可能？
6. 如果 b = 0, 那么 b = w - x < x 似乎是可能的
7. 