# Problem Description

Artsem has a friend Saunders from University of Chicago. Saunders presented him with the following problem.

Let [n] denote the set {1, ..., n}. We will also write f: [x] → [y] when a function f is defined in integer points 1, ..., x, and all its values are integers from 1 to y.

Now then, you are given a function f: [n] → [n]. Your task is to find a positive integer m, and two functions g: [n] → [m], h: [m] → [n], such that g(h(x)) = x for all x ∈ [m], and h(g(x)) = f(x) for all x ∈ [n], or determine that finding these is impossible.

## Input

The first line contains an integer n (1 ≤ n ≤ 10⁵).

The second line contains n space-separated integers — values f(1), ..., f(n) (1 ≤ f(i) ≤ n).

## Output

If there is no answer, print one integer -1.

Otherwise, on the first line print the number m (1 ≤ m ≤ 10⁶). On the second line print n numbers g(1), ..., g(n). On the third line print m numbers h(1), ..., h(m).

If there are several correct answers, you may output any of them. It is guaranteed that if a valid answer exists, then there is an answer satisfying the above restrictions.

## Examples

### Example 1

**Input:**

```text
3
1 2 3
```

**Output:**

```text
3
1 2 3
1 2 3
```

### Example 2

**Input:**

```text
3
2 2 2
```

**Output:**

```text
1
1 1 1
2
```

### Example 3

**Input:**

```text
2
2 1
```

**Output:**

```text
-1
```


### ideas
1. f[x] 是看作一条边, x -> f[x]的一条边
2. h(g(x)) = f(x)
3. g(h(x)) = x
4. 对于例子2， g(1) = 1, g(2) = 1, g(3) = 1
5. h(g(1...3)) = h(1) = f(1...3) = 2
6. m = 不同的f(?)的个数
7. 如果 f(a) = f(b), 那么可以让 g(a) = g(b) 具体的值先不用管
8. 然后 h(g(b)) = h(g(a)) = f(b), 那么就等于a就可以了
9. f(a)...f(b) 相等的集合中，取最小的a, 然后让g(a...b) = a
10. 让后 h(a) = a
11. 考虑例子3， 必须有 g(a) = a 