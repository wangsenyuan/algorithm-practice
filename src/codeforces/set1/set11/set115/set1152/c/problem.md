# Problem Statement

Neko loves divisors. During the latest number theory lesson, he got an interesting exercise from his math teacher.

Neko has two integers $a$ and $b$. His goal is to find a non-negative integer $k$ such that the least common multiple of $a + k$ and $b + k$ is the smallest possible. If there are multiple optimal integers $k$, he needs to choose the smallest one.

Given his mathematical talent, Neko had no trouble getting Wrong Answer on this problem. Can you help him solve it?

## Input

The only line contains two integers $a$ and $b$ $(1 \leq a, b \leq 10^9)$.

## Output

Print the smallest non-negative integer $k$ $(k \geq 0)$ such that the lowest common multiple of $a + k$ and $b + k$ is the smallest possible.

If there are many possible integers $k$ giving the same value of the least common multiple, print the smallest one.

## Examples

### Input 1

```
6 10
```

### Output 1

```
2
```

### Input 2

```
21 31
```

### Output 2

```
9
```

### Input 3

```
5 10
```

### Output 3

```
0
```

## Note

In the first test, one should choose $k = 2$, as the least common multiple of $6 + 2$ and $10 + 2$ is $24$, which is the smallest least common multiple possible.


### ideas
1. 这里要求的是 lcm(a + k, b + k) 最小
2. 不妨设 a <= b, 如果b % a == 0, 那么 k = 0
3. 