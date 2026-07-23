# G - Flip Row or Col

[Problem link](https://atcoder.jp/contests/abc396/tasks/abc396_g)

**Contest:** [AtCoder Beginner Contest 396](https://atcoder.jp/contests/abc396)

time limit: 2 sec

memory limit: 1024 MiB

score: 600 points

## Problem

There is an `H × W` grid; each cell contains `0` or `1`. The cell in row `i`,
column `j` contains `A_{i,j}`.

You may perform the following operations any number of times in any order:

- Operation X: choose row `x` and flip every cell in that row
  (`0 ↔ 1`).
- Operation Y: choose column `y` and flip every cell in that column.

Find the minimum possible value of `sum of all A_{x,y}` after the operations.

## Constraints

- `1 <= H <= 2 * 10^5`
- `1 <= W <= 18`
- Each row is a length-`W` string of `0` and `1`

## Input

```text
H W
A_{1,1}A_{1,2}...A_{1,W}
...
A_{H,1}A_{H,2}...A_{H,W}
```

## Output

Print the answer.

## Samples

### Sample 1

```text
Input
3 3
100
010
110

Output
2
```

### Sample 2

```text
Input
3 4
1111
1111
1111

Output
0
```

### Sample 3

```text
Input
10 5
10000
00111
11000
01000
10110
01110
10101
00100
00100
10001

Output
13
```

## Solution Summary

Flipping the same row or column twice has no effect, and row and column flips
commute. Therefore, it is enough to choose:

- one `W`-bit mask describing which columns are flipped;
- independently for every row, whether that row is flipped.

Encode each row as a `W`-bit integer.

### Cost for a fixed column mask

Suppose the chosen column-flip mask is `x`, and a row has mask `a`. After the
column flips, the row is

```text
a XOR x.
```

Let

```text
d = popcount(a XOR x).
```

If we do not flip this row, it contributes `d` ones. If we flip the row, all
`W` bits are complemented, so it contributes `W-d` ones. The optimal
contribution of this row is therefore

```text
min(d, W-d).
```

Define the following arrays over all `W`-bit masks:

```text
freq[a] = number of rows equal to mask a
cost[t] = min(popcount(t), W-popcount(t)).
```

For a column mask `x`, the minimum total number of ones is

```text
answer[x] = sum(freq[a] * cost[a XOR x]) over all masks a.
```

This is the XOR convolution of `freq` and `cost`.

### Walsh-Hadamard transform

The Walsh-Hadamard transform converts XOR convolution into pointwise
multiplication. Transform both arrays, multiply corresponding entries, and
apply the inverse transform:

```text
F = FWHT(freq)
G = FWHT(cost)
H[i] = F[i] * G[i]
answer = inverseFWHT(H)
```

The resulting `answer[x]` is the optimum after choosing column mask `x` and
then choosing every row flip optimally. The final result is the minimum value
among all `answer[x]`.

The transform combines pairs using

```text
(p, q) -> (p+q, p-q).
```

Applying the same transform twice multiplies every entry by `2^W`, so the
inverse transform performs the same operations and then divides every entry
by `2^W`.

### Correctness Proof

We prove that the algorithm returns the minimum possible number of ones.

Fix any set of column flips and let its mask be `x`. For a row with mask `a`,
the column flips produce `a XOR x`. If this row contains `d` ones, leaving it
unchanged costs `d`, while flipping it costs `W-d`. Since row flips affect
different rows independently, choosing the smaller value for every row gives
the minimum possible cost for this fixed `x`.

Grouping equal row masks does not change this sum. Hence the minimum cost for
column mask `x` is exactly

```text
sum(freq[a] * cost[a XOR x]).
```

The XOR-convolution theorem guarantees that the forward transforms,
pointwise multiplication, and inverse transform compute this value for every
mask `x`. Therefore, the transformed array contains the optimal cost for each
possible choice of column flips.

Every sequence of operations is equivalent to some column mask followed by
independent row-flip choices. The algorithm considers every column mask and,
for each one, makes the optimal row choices. Taking the minimum over all masks
therefore gives the global minimum.

### Complexity

Parsing and encoding the rows takes `O(HW)` time. Each Walsh-Hadamard transform
takes `O(W * 2^W)` time.

- Time: `O(HW + W * 2^W)`.
- Space: `O(2^W)`.

## Solution Summary 2: Difference-Count DP

This is the method implemented by `solve`. It computes, for every possible
column-flip mask `X`, how many input rows are at each Hamming distance from
`X`.

As in the previous solution, encode every row as a `W`-bit mask `B`. If the
column-flip mask is `X`, then the number of ones in this row becomes

```text
c = popcount(B XOR X).
```

After optionally flipping the whole row, its optimal contribution is

```text
min(c, W-c).
```

Therefore, for every `X`, we need the number of rows whose Hamming distance
from `X` is exactly `c`, for each `c = 0..W`.

### DP definition

The implementation uses

```text
dp[c][X]
```

and processes mask bits from low to high.

After the first `i` bits have been processed, `dp[c][X]` is the number of
input row masks `B` such that:

- `B` and `X` are equal in every unprocessed bit `i..W-1`;
- they differ in exactly `c` of the processed bits `0..i-1`.

Equivalently, each original row `B` has generated every mask obtained by
toggling a subset of the first `i` bits, and `c` is the size of that subset.

Initially, no bit has been processed. A row can only generate its own mask
with zero differences, so

```text
dp[0][B] += 1.
```

### Transition

Suppose bit `i` is being processed. Every current state has two choices:

1. Keep bit `i` unchanged. The mask and difference count stay the same. This
   state is already present in `dp`, so no explicit update is necessary.
2. Toggle bit `i`. The new mask is `X XOR (1<<i)`, and the difference count
   increases by one:

```text
dp[c+1][X XOR (1<<i)] += dp[c][X].
```

The code iterates `c` downward:

```text
for c := i; c >= 0; c-- {
    for X := 0; X < 1<<W; X++ {
        dp[c+1][X^(1<<i)] += dp[c][X]
    }
}
```

Descending order is essential. The transition writes into layer `c+1`; if
the layers were processed upward, that newly written value could be used
again during the same bit, incorrectly toggling bit `i` more than once.

For example, start with one row `B = 101`. After processing the lowest two
bits, it contributes to:

```text
X = 101, distance 0
X = 100, distance 1
X = 111, distance 1
X = 110, distance 2
```

These are exactly the masks obtained by toggling any subset of those two
processed bits.

### Extracting the answer

After all `W` bits have been processed, there are no unprocessed bits left.
Thus

```text
dp[c][X]
```

is exactly the number of input rows `B` satisfying

```text
popcount(B XOR X) = c.
```

For a fixed column mask `X`, its minimum possible number of ones is therefore

```text
sum(dp[c][X] * min(c, W-c)) for c = 0..W.
```

Evaluate this expression for every `X` and take the minimum.

### Correctness Proof

We first prove the DP invariant by induction on the number of processed bits.

Before any bit is processed, each row `B` appears only in `dp[0][B]`. It
agrees with `B` in every bit and has zero differences, so the invariant holds.

Assume it holds after processing the first `i` bits. When processing bit `i`,
leaving the bit unchanged preserves both the generated mask and its difference
count. Toggling it changes exactly that bit, producing
`X XOR (1<<i)` and increasing the difference count by one. These two choices
generate every subset of the first `i+1` bits exactly once. Therefore, the
invariant also holds after processing bit `i`.

After all bits are processed, the invariant says that `dp[c][X]` counts
exactly the rows at Hamming distance `c` from `X`.

Now fix a column-flip mask `X`. A row at distance `c` contains `c` ones after
the column flips. Leaving it unchanged costs `c`; flipping the whole row costs
`W-c`. Hence its optimal contribution is `min(c,W-c)`. Multiplying this value
by `dp[c][X]` and summing over `c` gives the optimal result for this fixed
`X`.

The algorithm evaluates every possible column mask and takes the minimum, so
it returns the global optimum.

### Complexity

For bit `i`, the transition processes `i+1` distance layers and all `2^W`
masks. Therefore, the total transition work is

```text
2^W * (1 + 2 + ... + W) = O(W^2 * 2^W).
```

- Time: `O(HW + W^2 * 2^W)`.
- Space: `O(W * 2^W)`.
