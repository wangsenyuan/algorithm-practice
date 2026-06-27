# D - Domino Covering XOR

[Problem link](https://atcoder.jp/contests/abc407/tasks/abc407_d)

**Contest:** [AtCoder Beginner Contest 407](https://atcoder.jp/contests/abc407)

time limit: 2 sec

memory limit: 1024 MiB

score: 425 points

There is a grid with `H` rows and `W` columns. Cell `(i, j)` (`1 <= i <= H`, `1 <= j <= W`) contains a
non-negative integer `A_{i,j}`.

Place zero or more dominoes on the grid. Each domino covers two adjacent cells:

- `(i, j)` and `(i, j+1)` for `1 <= i <= H`, `1 <= j < W`, or
- `(i, j)` and `(i+1, j)` for `1 <= i < H`, `1 <= j <= W`.

No cell may be covered by more than one domino.

For a domino placement, define its **score** as the bitwise XOR of all integers in cells **not**
covered by any domino.

Find the maximum possible score.

## Constraints

- `1 <= H`
- `1 <= W`
- `H * W <= 20`
- `0 <= A_{i,j} < 2^60` (`1 <= i <= H`, `1 <= j <= W`)
- All input values are integers

## Input

```text
H W
A_{1,1} A_{1,2} ... A_{1,W}
A_{2,1} A_{2,2} ... A_{2,W}
...
A_{H,1} A_{H,2} ... A_{H,W}
```

## Output

Print the answer.

## Sample Input 1

```text
3 4
1 2 3 8
4 0 7 10
5 2 4 2
```

## Sample Output 1

```text
15
```

## Sample Input 2

```text
1 11
1 2 4 8 16 32 64 128 256 512 1024
```

## Sample Output 2

```text
2047
```

You may place no dominoes at all.

## Sample Input 3

```text
4 5
74832 16944 58683 32965 97236
52995 43262 51959 40883 58715
13846 24919 65627 11492 63264
29966 98452 75577 40415 77202
```

## Sample Output 3

```text
131067
```

### ideas
