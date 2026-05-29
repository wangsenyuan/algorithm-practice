# Problem E - Reflection on Grid

https://atcoder.jp/contests/abc431/tasks/abc431_e

**Time Limit:** 2 sec / **Memory Limit:** 1024 MiB

**Score:** 500 points

## Problem Statement

There is a grid with $H$ rows and $W$ columns. We will refer to the cell at the
$i$-th row from the top and $j$-th column from the left as cell $(i,j)$. Each cell
has at most one mirror placed on it.

Takahashi is standing on the left side of cell $(1,1)$, and Aoki is standing on
the right side of cell $(H,W)$. Takahashi has a flashlight and is shining light
from the left side of cell $(1,1)$ toward the right. Here, assume that the
flashlight's light does not diffuse and is a light ray that travels straight.

Takahashi's goal is to deliver the flashlight's light to Aoki by using the
mirrors in the grid.

There are three types of mirror placements. When light hits a mirror, the
direction of the light changes according to the mirror placement. For each
mirror placement, the exit direction for each entry direction is as shown in the
figures below.

- Type A (no mirror is placed)

![Type A](https://img.atcoder.jp/abc431/9a3821cb76d936b95b6979e084d56994.png)

- Type B (a mirror is placed on the diagonal connecting the upper-left and
  lower-right)

![Type B](https://img.atcoder.jp/abc431/283ea6eabada389e76e89518fcb9fb18.png)

- Type C (a mirror is placed on the diagonal connecting the upper-right and
  lower-left)

![Type C](https://img.atcoder.jp/abc431/9410773fd7615321f27f070d5f0b1844.png)

The mirror placement on the grid is represented by $H$ strings of length $W$:
$S_1,S_2,\ldots,S_H$. When the $j$-th character of $S_i$ is `A`, cell $(i,j)$ is
type A; when it is `B`, cell $(i,j)$ is type B; when it is `C`, cell $(i,j)$ is
type C.

Takahashi can perform the following operation any number of times to deliver the
light to Aoki:

- Choose one cell and change the mirror placement of that cell to a different
  type.

Find the minimum number of operations needed to deliver the light to Aoki.

You are given $T$ test cases; find the answer for each.

## Constraints

- $1 \leq T$
- $1 \leq H,W$
- $HW \leq 2 \times 10^5$
- $S_i$ is a string of length $W$ consisting of `A`, `B`, `C`.
- $T$, $H$, and $W$ are integers.
- The sum of $HW$ over all test cases is at most $2 \times 10^5$.

## Input

The input is given from Standard Input in the following format:

```
T
case_1
case_2
⋮
case_T
```

Each test case is given in the following format:

```
H W
S_1
S_2
⋮
S_H
```

## Output

Output $T$ lines. The $i$-th line should contain the answer for the $i$-th test
case.

## Sample Input 1

```
4
3 4
ABCB
CACC
BCBA
2 2
CB
AA
1 10
BCBCBCBCBC
10 10
CCAABACAAA
CCCBACACCA
BACAABCBBA
ACCCAACCCA
CCAAAACCBA
AACBBACCAA
BCCCACBBAB
CBBCAACCCC
CBBCCBCBCA
BBACABBACC
```

## Sample Output 1

```
0
2
10
5
```

For the first test case, the light can be delivered to Aoki without performing
any operations.

![Sample 1](https://img.atcoder.jp/abc431/f66a190627b09b9e6c4cbacb84b6d9bd.png)

For the second test case, by changing the mirror placement of cell $(1,1)$ to
type A and the mirror placement of cell $(2,2)$ to type B, the light can be
delivered to Aoki as shown in the figure below. It is impossible to deliver the
light to Aoki with one or fewer operations, so the answer is $2$.

![Sample 2](https://img.atcoder.jp/abc431/1aeac5f2fecf4319d369b8cc3067959a.png)
