# F - Manhattan Cafe

[Problem link](https://atcoder.jp/contests/abc265/tasks/abc265_f)

**Contest:** [AtCoder Beginner Contest 265](https://atcoder.jp/contests/abc265)

time limit: 6 sec

memory limit: 1024 MiB

score: 500 points

For two points `x = (x_1, x_2, ..., x_N)` and `y = (y_1, y_2, ..., y_N)` in
`N`-dimensional space, their Manhattan distance is:

```text
d(x, y) = sum |x_i - y_i|
```

A point is a lattice point if every coordinate is an integer.

Given two lattice points `p = (p_1, p_2, ..., p_N)` and
`q = (q_1, q_2, ..., q_N)`, count the number of lattice points `r` such that:

```text
d(p, r) <= D
d(q, r) <= D
```

Print the count modulo `998244353`.

## Constraints

- `1 <= N <= 100`
- `0 <= D <= 1000`
- `-1000 <= p_i, q_i <= 1000`
- All input values are integers

## Input

```text
N D
p_1 p_2 ... p_N
q_1 q_2 ... q_N
```

## Output

Print the answer.

## Sample Input 1

```text
1 5
0
3
```

## Sample Output 1

```text
8
```

When `N = 1`, the valid points are `-2, -1, 0, 1, 2, 3, 4, 5`.

## Sample Input 2

```text
3 10
2 6 5
2 1 2
```

## Sample Output 2

```text
632
```

## Sample Input 3

```text
10 100
3 1 4 1 5 9 2 6 5 3
2 7 1 8 2 8 1 8 2 8
```

## Sample Output 3

```text
145428186
```

## Solution

### 1. DP state

Manhattan distance is the sum of the contributions from each coordinate, so we
process the coordinates from left to right.

After some coordinates have been processed, let:

```text
dp[i][j] = number of ways to choose those coordinates of r such that
           their total distance from p is exactly i, and
           their total distance from q is exactly j
```

Only distances at most `D` need to be stored. Initially no coordinate has been
chosen, so:

```text
dp[0][0] = 1
```

After all `N` coordinates have been processed, every state in the table is
valid, and the answer is:

```text
sum(dp[i][j]) for 0 <= i,j <= D
```

### 2. Choices for one coordinate

Consider one coordinate and write:

```text
a = p_t
b = q_t
s = |a-b|
```

We may assume `a <= b`; swapping them only exchanges the two distances. Let
`x = r_t`. There are three possible positions for `x`.

#### a. Between `a` and `b`

Write `x = a+k`, where `0 <= k <= s`. Its contribution to the two distances
is:

```text
(|x-a|, |x-b|) = (k, s-k)
```

Thus the possible pairs are:

```text
(0,s), (1,s-1), ..., (s,0)
```

The endpoints `x=a` and `x=b` are included here.

#### b. To the left of both points

Write `x = a-h`, where `h >= 1`. The contribution is:

```text
(h, s+h)
```

#### c. To the right of both points

Write `x = b+h`, where `h >= 1`. The contribution is:

```text
(s+h, h)
```

The two outside cases start at `h=1`, so the endpoints already counted by the
middle case are not counted again.

### 3. Direct transition and why it is too slow

Suppose the new accumulated distances are `(i,j)`.

For a middle choice with contribution `(k,s-k)`, the previous state is:

```text
(i-k, j-s+k), for 0 <= k <= s
```

For a left choice with contribution `(h,s+h)`, it is:

```text
(i-h, j-s-h), for h >= 1
```

For a right choice with contribution `(s+h,h)`, it is:

```text
(i-s-h, j-h), for h >= 1
```

Enumerating every `k` and `h` for every state would take `O(ND^3)`. The key
observation is that all three sums follow diagonals of the DP table, so diagonal
prefix sums make each transition `O(1)`.

### 4. Anti-diagonal prefix sums for the middle case

Define:

```text
dp2[i][j] = dp[i][j] + dp[i-1][j+1] + dp[i-2][j+2] + ...
```

It is computed by:

```go
dp2[i][j] = dp[i][j]
if i != 0 && j != D {
    dp2[i][j] += dp2[i-1][j+1]
}
```

For target `(i,j)`, the middle choices require:

```text
dp[i][j-s] + dp[i-1][j-s+1] + ... + dp[i-s][j]
```

This is a segment of one anti-diagonal. Start with the whole prefix:

```text
dp2[i][j-s]
```

and remove everything after the required `s+1` entries:

```text
dp2[i-s-1][j+1]
```

Therefore the transition is conceptually:

```text
nxt[i][j] += dp2[i][j-s]
nxt[i][j] -= dp2[i-s-1][j+1]
```

Any term whose index is outside `[0,D]` is omitted.

There is one extra boundary detail. If `j-s < 0`, the nominal first source
state `(i,j-s)` is invalid. We move along the same anti-diagonal until the
second index becomes zero. If `j-s = -z`, the first valid source is `(i-z,0)`.
This is what the code does with:

```go
si := i
sj := j - s
if sj < 0 {
    si += sj
    sj = 0
}
```

### 5. Main-diagonal prefix sums for the outside cases

Define:

```text
dp3[i][j] = dp[i][j] + dp[i-1][j-1] + dp[i-2][j-2] + ...
```

It is computed by:

```go
dp3[i][j] = dp[i][j]
if i != 0 && j != 0 {
    dp3[i][j] += dp3[i-1][j-1]
}
```

For a left-side choice, the contributions are `(h,s+h)` for every `h >= 1`.
Starting from all states represented by `dp3[i][j]`, their corresponding target
is `(i+1,j+s+1)`, so:

```go
nxt[i+1][j+s+1] += dp3[i][j]
```

Indeed, for a fixed target `(A,B)`, this adds the previous states:

```text
(A-1, B-s-1), (A-2, B-s-2), ...
```

For a right-side choice, the target is `(i+s+1,j+1)`, so:

```go
nxt[i+s+1][j+1] += dp3[i][j]
```

Both operations are additions: the left and right regions contain different,
valid choices of `r_t`.

### 6. Example

For the first sample, `p_1=0`, `q_1=3`, and `D=5`, so `s=3`.

The middle choices are:

| `r_1` | distance from `p` | distance from `q` |
|---:|---:|---:|
| 0 | 0 | 3 |
| 1 | 1 | 2 |
| 2 | 2 | 1 |
| 3 | 3 | 0 |

The valid choices to the left are `-1` and `-2`, producing `(1,4)` and
`(2,5)`. The valid choices to the right are `4` and `5`, producing `(4,1)` and
`(5,2)`. Hence there are `4+2+2=8` valid lattice points.

### 7. Correctness

We maintain the following invariant:

> After processing the first `t` coordinates, `dp[i][j]` is exactly the number
> of choices for those coordinates whose accumulated distances from `p` and
> `q` are `i` and `j`.

It is true initially because the empty choice has distances `(0,0)`. For the
next coordinate, every integer value lies in exactly one of the three regions:
between the two given coordinates, to their left, or to their right. The three
transitions add precisely the distance pair contributed by that value, and the
middle endpoints are not repeated in the outside regions. Thus every extended
partial point is counted exactly once, preserving the invariant.

After all coordinates are processed, the invariant says that `dp[i][j]` counts
the full lattice points with exact distances `(i,j)`. Summing the stored states
therefore counts exactly those with both distances at most `D`.

### 8. Complexity

For each coordinate, the algorithm builds and scans `O(D^2)` tables.

```text
Time:  O(ND^2)
Space: O(D^2)
```
