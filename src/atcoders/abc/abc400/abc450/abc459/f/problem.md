# F - -1, +1

**Contest:** [ABC459](https://atcoder.jp/contests/abc459) — AtCoder Beginner Contest 459  
**Task:** [https://atcoder.jp/contests/abc459/tasks/abc459_f](https://atcoder.jp/contests/abc459/tasks/abc459_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

You are given a sequence of non-negative integers
`A = (A_1, A_2, ..., A_N)` of length `N`.

You can perform the following operation on `A` zero or more times:

- Choose an integer `i` with `1 <= i <= N - 1`, decrease `A_i` by `1`, and
  increase `A_{i+1}` by `1`.

Find the minimum number of operations required to make `A` strictly increasing.

It can be proved that the answer is less than `2^63`.

You are given `T` test cases; solve each.

## Constraints

- `1 <= T <= 3 * 10^5`
- `1 <= N <= 2 * 10^5`
- `0 <= A_i <= 10^9`
- The sum of `N` across all test cases is at most `6 * 10^5`.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
T
case_1
case_2
...
case_T
```

The `i`-th test case is given in the following format:

```text
N
A_1 A_2 ... A_N
```

## Output

Output the answers for the test cases in order, separated by newlines.

## Sample Input 1

```text
4
3
0 1 0
4
4 6 3 5
7
1 2 3 4 5 6 7
10
11 9 1 3 17 19 10 19 17 3
```

## Sample Output 1

```text
3
5
0
78
```

## Sample Explanation

For the first test case, the array can be made strictly increasing in three
operations:

1. Choose `i = 1`. `A` becomes `(-1, 2, 0)`.
2. Choose `i = 2`. `A` becomes `(-1, 1, 1)`.
3. Choose `i = 2`. `A` becomes `(-1, 0, 2)`.

It is impossible to make `A` strictly increasing in fewer than three operations.

### ideas
1. 贪心？ dp[i] 等于将后缀变成递增序列的操作数
2. 但是问题在于，如果a[i]不一样，操作数也不一样
3. sum不变
4. 始终是将前面的数,移动到后面的数,操作数 = 移动的距离
5. 假设 a[i]是a[0]后面的最小的数, 那么必须将a[0] - (a[i] - 1)个数,移动到i去
6. 这样子a[0]就变成了最小的数.
7. 似乎是说得通的,如果a[0]变成了最小的数,那么就就需要再减少它了

## Solution explanation

One operation moves `1` unit from position `i` to position `i+1`. So mass can
only move to the right.

Let the final array be `B`. Because operations only move values to the right:

- total sum is preserved;
- for every prefix, the final prefix sum cannot exceed the original prefix sum:

```text
B_1 + ... + B_i <= A_1 + ... + A_i
```

The cost of transforming `A` into `B` is:

```text
sum over i: i * (B_i - A_i)
```

using 0-based indexing. If `B_i - A_i` is positive, that many units arrived at
position `i`; if it is negative, that many units left position `i`. This weighted
sum equals the total distance moved to the right.

### Remove strictness

We need:

```text
B_0 < B_1 < ... < B_{N-1}
```

Define:

```text
X_i = A_i - i
Y_i = B_i - i
```

Then `B_i < B_{i+1}` is equivalent to:

```text
Y_i <= Y_{i+1}
```

So the target becomes: change `X` into a nondecreasing integer array `Y`, while
preserving total sum and respecting the prefix constraints.

The code first performs this normalization:

```go
for i := range n {
    a[i] -= i
}
```

After this point, `a` means the normalized array `X`.

### Block merging

The optimal nondecreasing integer array can be built with a pool-adjacent
violators style process.

Maintain blocks. Each block stores:

```text
(length, sum)
```

When a new value `v` is appended, it starts as a block:

```text
(1, v)
```

If the previous block can be placed before the current block while keeping the
expanded values nondecreasing, we keep them separate. Otherwise, the two blocks
must be merged.

For a block with length `l` and sum `s`, the lexicographically smallest
nondecreasing integer sequence with that sum is:

```text
floor(s / l), floor((s + 1) / l), ..., floor((s + l - 1) / l)
```

So:

- the minimum value in the block is `floor(s / l)`;
- the maximum value in the block is `ceil(s / l)`.

Two adjacent blocks are valid if:

```text
ceil(left.sum / left.length) <= floor(right.sum / right.length)
```

That is the merge condition in the code:

```go
if divCeil(tail.second, tail.first) <= divFloor(now.second, now.first) {
    break
}
```

If this condition is false, the two blocks overlap in value order, so they must
be merged:

```go
now.first += tail.first
now.second += tail.second
```

This repeats until the block sequence is valid again.

### Why floor and ceil are explicit

The normalized values can be negative. Go integer division truncates toward
zero, but the algorithm needs mathematical floor and ceiling.

For example, `-1 / 2` in Go is `0`, but mathematically:

```text
floor(-1 / 2) = -1
```

So the implementation uses:

```go
divFloor()
divCeil()
```

This matters for cases such as `A = [0, 1, 1]`, where the normalized values are
`[0, 0, -1]`.

### Expanding blocks

After all blocks are fixed, each block `(l, s)` is expanded as:

```go
for i := range l {
    b = append(b, divFloor(s+i, l))
}
```

This produces the nondecreasing normalized target array `Y`.

Finally, compute the operation count:

```go
res += i * (b[i] - a[i])
```

Here `a` is the original normalized array `X`, and `b` is the optimized
normalized array `Y`. The subtraction by index cancels out, so this weighted
sum is the same as the movement cost in the original array.

### Complexity

Each element creates one block and each block is merged at most once. Therefore,
the block process is linear.

Time complexity per test case:

```text
O(N)
```

Space complexity:

```text
O(N)
```
