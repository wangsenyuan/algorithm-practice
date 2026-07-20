# E - Takahashi is Slime 2

[Problem link](https://atcoder.jp/contests/abc384/tasks/abc384_e)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 450 points

## Problem Statement

There is a grid with `H` rows and `W` columns. Cell `(i, j)` is at the `i`-th
row from the top and `j`-th column from the left.

Initially, there is a slime with strength `S_{i,j}` in cell `(i, j)`, and
Takahashi is the slime in cell `(P, Q)`.

Find the maximum possible strength of Takahashi after performing the following
action any number of times (possibly zero):

- Among the slimes adjacent to him, choose one whose strength is **strictly
  less than** `1/X` times his strength and absorb it. The absorbed slime
  disappears, and Takahashi's strength increases by that slime's strength.

When a slime is absorbed, the gap is immediately filled by Takahashi, and
slimes that were adjacent to the disappeared one become newly adjacent to him.

## Constraints

- `1 <= H, W <= 500`
- `1 <= P <= H`
- `1 <= Q <= W`
- `1 <= X <= 10^9`
- `1 <= S_{i,j} <= 10^12`
- All input values are integers

## Input

```
H W X
P Q
S_1,1 ... S_1,W
...
S_H,1 ... S_H,W
```

## Output

Print the maximum possible strength of Takahashi after performing the action.

## Samples

### Sample 1

Input:

```
3 3 2
2 2
14 6 9
4 9 20
17 15 7
```

Output:

```
28
```

### Sample 2

Input:

```
3 4 1
1 1
5 10 1 1
10 1 1 1
1 1 1 1
```

Output:

```
5
```

### Sample 3

Input:

```
8 10 2
1 5
388 130 971 202 487 924 247 286 237 316
117 166 918 106 336 928 493 391 235 398
124 280 425 955 212 988 227 222 307 226
336 302 478 246 950 368 291 236 170 101
370 200 204 141 287 410 388 314 205 460
291 104 348 337 404 399 416 263 415 339
105 420 302 334 231 481 466 366 401 452
119 432 292 403 371 417 351 231 482 184
```

Output:

```
1343
```

## Solution

Takahashi occupies a connected region of cells that have already been
absorbed. Therefore, the only slimes that can be chosen next are the
unabsorbed cells adjacent to this region.

Maintain all such cells in a min-heap ordered by slime strength. At every
step, inspect the weakest reachable slime.

- If it is weak enough to absorb, add its strength to Takahashi's strength
  and insert its unvisited neighbors into the heap.
- If it cannot be absorbed, stop. Every other reachable slime is at least as
  strong, so none of them can be absorbed either.

The initial cell `(P, Q)` is also inserted into the heap. The code starts the
accumulated strength `sum` at zero, accepts this first cell without checking
the absorption condition, and then discovers its neighbors in the same way as
every other absorbed cell.

### Why choosing the weakest slime is optimal

Absorbing a slime only increases Takahashi's strength and expands, or leaves
unchanged, the set of reachable cells. It never makes a future absorption
impossible.

Suppose at some moment at least one reachable slime can be absorbed. The
weakest reachable slime can also be absorbed. Taking it immediately cannot
hurt, and its added strength may make more slimes absorbable. Thus, repeatedly
choosing the minimum-strength reachable slime obtains every absorption that is
possible.

Conversely, if the weakest reachable slime cannot be absorbed, every other
reachable slime is at least as strong and also fails the condition. There is no
valid next action, so the current strength is the maximum possible answer.

### Checking the strict inequality safely

For the current strength `sum` and a candidate of strength `v`, absorption is
allowed exactly when

```text
v < sum / X
```

or, without fractions,

```text
X * v < sum.
```

However, `X * v` can be as large as `10^21`, which does not fit in `int64`.
Since all values are positive integers, the same condition can be written as

```text
v <= (sum - 1) / X.
```

Therefore, after processing the initial cell, the algorithm stops when

```go
v > (sum - 1) / int64(x)
```

This preserves the required strict inequality without overflowing.

### Correctness Proof

We prove that the algorithm returns the maximum possible final strength.

At all times, the heap contains exactly the discovered, unabsorbed cells on
the boundary of Takahashi's absorbed region. This is initially true after the
starting cell is processed. Whenever a boundary cell is absorbed, all of its
previously undiscovered neighbors are inserted, so the invariant continues to
hold.

In each iteration, the algorithm removes the weakest boundary slime `v`.

- If `v <= (sum - 1) / X`, then `X * v < sum`, so absorbing `v` is a valid
  action. It increases `sum` and cannot remove any future option.
- Otherwise, `v` cannot be absorbed. Because `v` is the minimum element of the
  heap, every other boundary slime has strength at least `v` and cannot be
  absorbed either. Hence no valid action remains.

Thus every absorption performed by the algorithm is valid, and when it stops,
no continuation is possible. Since absorption only increases strength, the
returned value is the maximum achievable final strength.

### Complexity

Each of the `H * W` cells is marked and inserted into the heap at most once.
Every heap operation costs `O(log(H * W))`.

- Time: `O(H * W * log(H * W))`
- Space: `O(H * W)`
