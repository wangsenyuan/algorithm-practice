# C - ~

[Problem link](https://atcoder.jp/contests/abc406/tasks/abc406_c)

**Contest:** [Panasonic Programming Contest 2025 (AtCoder Beginner Contest 406)](https://atcoder.jp/contests/abc406)

time limit: 2 sec

memory limit: 1024 MiB

score: 350 points

For an integer sequence A = (A_1, A_2, ..., A_{|A|}), we say A is **tilde-shaped** if all of the
following hold:

- |A| >= 4
- A_1 < A_2
- There is exactly one index i with 2 <= i < |A| such that A_{i-1} < A_i > A_{i+1}
- There is exactly one index i with 2 <= i < |A| such that A_{i-1} > A_i < A_{i+1}

You are given a permutation P = (P_1, ..., P_N) of (1, 2, ..., N). Count how many contiguous
subarrays of P are tilde-shaped.

## Constraints

- 4 <= N <= 3 * 10^5
- P is a permutation of (1, 2, ..., N)
- All input values are integers

## Input

```text
N
P_1 P_2 ... P_N
```

## Output

Print the number of tilde-shaped contiguous subarrays.

## Sample Input 1

```text
6
1 3 6 4 2 5
```

## Sample Output 1

```text
2
```

The tilde-shaped subarrays are (3, 6, 4, 2, 5) and (1, 3, 6, 4, 2, 5).

## Sample Input 2

```text
6
1 2 3 4 5 6
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
12
11 3 8 9 5 2 10 4 1 6 12 7
```

## Sample Output 3

```text
4
```


## Solution

A tilde-shaped subarray has a fixed shape:

```text
increase ... peak decrease ... valley increase ...
```

So once the peak index `i` is fixed, the subarray is determined by:

- a left boundary inside the strictly increasing run ending at `i`
- a right boundary inside the strictly increasing run starting after the valley

### Precomputation

Scan left to right and compute `dp[i]`: the length of the longest strictly increasing contiguous
segment ending at `i`. Reset the counter when `P[i-1] > P[i]`.

Scan right to left and compute:

- `r`: the length of the strictly decreasing segment that starts at `i+1`
- `fp[i]`: the length of the strictly increasing segment that starts at `i`

Reset `r` when `P[i] < P[i+1]`, and reset the increasing counter when `P[i] > P[i+1]`.

### Counting

Only positions with `P[i] > P[i+1]` can be the peak of a tilde subarray.

For such an index `i`, let `j = i + 1 + r`. Then:

- `dp[i] - 1` is the number of valid left boundaries
- `fp[j]` is the number of valid right boundaries

Add `(dp[i] - 1) * fp[j]` to the answer for every peak candidate `i` with `j < N`.

Example for `P = (1, 3, 6, 4, 2, 5)`:

- Peak at `6` (index 2): `dp[2] = 3`, so 2 left starts `(1, ...)` and `(3, ...)`
- Decrease `4, 2`, then increase `5`, so `j = 5` and `fp[5] = 1`
- Contribution: `2 * 1 = 2`

### Correctness

Every tilde subarray has exactly one peak and one valley, so its internal shape is uniquely an
increase, then a decrease, then an increase. The peak must satisfy `P[i-1] < P[i] > P[i+1]`, and
the first condition `A_1 < A_2` means the left boundary must lie in the increasing run ending at
the peak, but not at the peak itself. That gives `dp[i] - 1` choices.

After the peak, the subarray must contain the full decreasing run up to the valley and then continue
within the increasing run after the valley. The variable `r` measures the forced decreasing length,
so `j = i + 1 + r` is exactly the first index of the final increasing part. Extending the subarray
right within that run gives `fp[j]` choices.

Different pairs `(left, right)` produce different subarrays, and every tilde subarray is counted
once at its peak. Therefore the sum equals the required count.

### Complexity

Both scans are linear, and each index is processed once.

Time: `O(N)`  
Memory: `O(N)`