# B. Mushroom Scientists

[Problem link](https://codeforces.com/problemset/problem/185/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

As you very well know, the whole Universe traditionally uses a three-dimensional
Cartesian coordinate system. In this system each point corresponds to three real
coordinates `(x, y, z)`. In this coordinate system, the distance between the
center of the Universe and the point is calculated by the following formula:

```text
sqrt(x^2 + y^2 + z^2)
```

Mushroom scientists that work for the Great Mushroom King think that the
Universe is not exactly right, and the distance from the center of the Universe
to a point equals:

```text
x^a * y^b * z^c
```

To test the metric of mushroom scientists, the usual scientists offered them a
task: find such `x, y, z` (`0 <= x, y, z` and `x + y + z <= S`) that the
distance between the center of the Universe and the point `(x, y, z)` is as
large as possible in the metric of mushroom scientists. The mushroom scientists
are not good at maths, so they commissioned you to do the task.

Note that in this problem it is considered that `0^0 = 1`.

## Input

The first line contains a single integer `S` (`1 <= S <= 10^3`) — the maximum
sum of coordinates of the sought point.

The second line contains three space-separated integers `a`, `b`, `c`
(`0 <= a, b, c <= 10^3`) — the numbers that describe the metric of mushroom
scientists.

## Output

Print three real numbers — the coordinates of the point that reaches the
maximum value in the metric of mushroom scientists. If there are multiple
answers, print any of them that meets the limitations.

The natural logarithm of the distance from the center of the Universe to the
given point in the metric of mushroom scientists must not differ from the natural
logarithm of the maximum distance by more than `10^-6`. We think that
`ln(0) = -infinity`.

## Example

### Input

```text
3
1 1 1
3
2 0 0
```

### Output

```text
1.0 1.0 1.0
3.0 0.0 0.0
```


### ideas
1. 在满足 x, y, z, x + y + z <= sum 的情况下,
2. 找到 x^^a * y ^^ b * z ^^ c 最大的3个点
3. 迭代x, y? 快速的找到z? x和y确定的情况下,z不应该是越大越好吗?
4. 是不是问题出在实数上? 如果 x < 1, 除非a = 0, 否则x应该直接取0
5. 如果 a >= b, a >= c, 那么x越大越大.
6. 