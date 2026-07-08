# D - 2-variable Function

[Problem link](https://atcoder.jp/contests/abc246/tasks/abc246_d)

**Contest:** [AtCoder Beginner Contest 246](https://atcoder.jp/contests/abc246)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

Given an integer `N`, find the smallest integer `X` that satisfies all of:

- `X >= N`
- There exist non-negative integers `(a, b)` such that `X = a^3 + a^2*b + a*b^2 + b^3`

## Constraints

- `N` is an integer
- `0 <= N <= 10^18`

## Input

```text
N
```

## Output

Print the answer as an integer.

## Sample Input 1

```text
9
```

## Sample Output 1

```text
15
```

For any integer `X` with `9 <= X <= 14`, no pair `(a, b)` satisfies the condition.
For `X = 15`, `(a, b) = (2, 1)` works.

## Sample Input 2

```text
0
```

## Sample Output 2

```text
0
```

`N` itself may already satisfy the condition.

## Sample Input 3

```text
999999999989449206
```

## Sample Output 3

```text
1000000000000000000
```

Input and output may not fit in a 32-bit integer type.

## Solution

Let

```text
f(a, b) = a^3 + a^2 b + a b^2 + b^3
        = (a + b)(a^2 + b^2)
```

Both `a` and `b` are non-negative, and `f(a, b)` is symmetric:

```text
f(a, b) = f(b, a)
```

So it is enough to enumerate the smaller side `a` and find the smallest `b >= a` such
that `f(a, b) >= N`. The important boundary is that `a` may be `0`; otherwise values such
as `1 = f(0, 1)`, `8 = f(0, 2)`, and `27 = f(0, 3)` are missed.

Because `N <= 10^18`, any needed value has `a, b <= 10^6`, since:

```text
f(10^6, 0) = 10^18
```

For each `a` from `0` to `10^6`:

1. Binary search the first `b` in `[0, 10^6]` with
   `(a + b) * (a * a + b * b) >= N`.
2. If `b < a`, stop. From this point on, the symmetric pair was already considered when
   the smaller value was enumerated.
3. Compute the exact value `f(a, b)` and minimize the answer.

The search predicate is monotonic in `b`, so binary search is valid.

The time complexity is `O(10^6 log 10^6)`, and the memory complexity is `O(1)`.
