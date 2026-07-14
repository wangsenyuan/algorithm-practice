# G - Groups

[Problem link](https://atcoder.jp/contests/abc217/tasks/abc217_g)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 600 points

## Problem Statement

You are given positive integers `N` and `M`. For each `k = 1, ..., N`, solve
the following problem.

- Problem: Divide `N` people with ID numbers `1` through `N` into `k`
  non-empty groups. People whose ID numbers are equal modulo `M` cannot belong
  to the same group.
- How many such ways are there? Find the answer modulo `998244353`.

Two ways are considered different when there is a pair `(x, y)` such that
Person `x` and Person `y` belong to the same group in one way but not in the
other.

## Constraints

- `2 <= N <= 5000`
- `2 <= M <= N`
- All values in input are integers

## Input

```
N M
```

## Output

Print `N` lines. The `i`-th line should contain the answer for `k = i`.

## Samples

### Sample 1

Input:

```
4 2
```

Output:

```
0
2
4
1
```

### Sample 2

Input:

```
6 6
```

Output:

```
1
31
90
65
15
1
```

### Sample 3

Input:

```
20 5
```

Output:

```
0
0
0
331776
207028224
204931064
814022582
544352515
755619435
401403040
323173195
538468102
309259764
722947327
162115584
10228144
423360
10960
160
1
```


## Solution

Process the people in increasing order of ID and use dynamic programming.

Let `dp[j]` be the number of valid ways to divide the people processed so far
into exactly `j` non-empty, **unlabeled** groups.

Initially no person has been processed, so:

```text
dp[0] = 1
```

### How many existing groups can person `i` join?

Before processing person `i`, the earlier people with the same remainder
modulo `M` are:

```text
i-M, i-2M, i-3M, ...
```

Their number is:

```text
same = floor((i-1) / M)
```

In every valid partition, these `same` people must already be in different
groups. Therefore, if the current partition has `j` groups, exactly `same` of
them are forbidden for person `i`, while the other

```text
j - same
```

groups are available.

Notice that this is not `C(j, same)`: the groups occupied by the earlier
same-remainder people have already been determined by the current partition.
We only need to choose one available group for the new person.

### Transition

After inserting person `i`, a partition into `j` groups can be obtained in two
ways.

1. Put person `i` alone in a new group.

   The previous `i-1` people must form `j-1` groups, giving `dp[j-1]` ways.
   Because groups are unlabeled, there is only one way to add this singleton
   group.

2. Put person `i` into an existing group.

   The previous people already form `j` groups. Of those groups,
   `j - same` are available, giving `dp[j] * (j - same)` ways.

Thus the recurrence is:

```text
newDp[j] = dp[j-1] + dp[j] * (j - floor((i-1) / M))
```

All operations are taken modulo `998244353`.

The implementation stores `newDp` in the same array as `dp`. It iterates `j`
from large to small, so `dp[j-1]` still contains the value from before person
`i` was inserted. Updating in increasing order would reuse a value from the
current iteration and count invalid states.

The smallest possible number of groups grows as the largest remainder class
grows. The lower bound in the loop skips states that are already zero; at
`j = same`, both terms are zero, so that state also remains zero.

### Correctness Proof

We prove that after processing people `1, 2, ..., i`, `dp[j]` equals the number
of valid partitions of these people into exactly `j` non-empty groups.

The claim is true before processing anyone: the empty set has exactly one
partition into zero groups, represented by `dp[0] = 1`.

Assume the claim is true after processing the first `i-1` people. Consider any
valid partition of the first `i` people into `j` groups.

- If person `i` is alone, removing that singleton leaves a unique valid
  partition of the first `i-1` people into `j-1` groups. By the induction
  hypothesis, these partitions contribute exactly `dp[j-1]` possibilities.
- Otherwise, removing person `i` leaves a valid partition into `j` groups. The
  `floor((i-1)/M)` earlier people congruent to `i` modulo `M` occupy distinct
  groups, and person `i` may join any of the other
  `j - floor((i-1)/M)` groups. By the induction hypothesis, this contributes
  exactly `dp[j] * (j - floor((i-1)/M))` possibilities.

The two cases are disjoint and cover every valid partition. Therefore the
transition computes the correct number for every `j`. By induction, after all
`N` people are processed, `dp[k]` is the required answer for every
`k = 1, 2, ..., N`.

### Complexity

There are `N` iterations, each updating at most `N` states.

- Time: `O(N^2)`
- Space: `O(N)`
