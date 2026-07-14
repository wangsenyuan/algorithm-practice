# F - Make Pair

[Problem link](https://atcoder.jp/contests/abc217/tasks/abc217_f)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 500 points

## Problem Statement

There are `2N` students standing in a row, numbered `1, 2, ..., 2N` from left
to right. For all pairs of two students, they are on good or bad terms.
Specifically, for each `1 <= i <= M`, Student `A_i` and Student `B_i` are on
good terms; for the remaining pairs of two students, they are on bad terms.

The teacher is going to do the following operation `N` times to form `N` pairs
of two students.

- Choose two adjacent students who are on good terms, pair them, and remove
  them from the row.
- If the removed students were not at an end of the row, close up the gap so
  that the two students who were to the left and right of the removed students
  are now adjacent.

Find the number, modulo `998244353`, of possible ways to do the operation `N`
times. Two ways are considered different when there exists some `i` such that
the pair of students chosen in the `i`-th operation differs.

## Constraints

- `1 <= N <= 200`
- `0 <= M <= N(2N-1)`
- `1 <= A_i < B_i <= 2N`
- All pairs `(A_i, B_i)` are distinct
- All values in input are integers

## Input

```
N M
A_1 B_1
A_2 B_2
...
A_M B_M
```

## Output

Print the number of possible ways to complete the procedure, modulo
`998244353`.

## Samples

### Sample 1

Input:

```
2 3
1 2
1 4
2 3
```

Output:

```
1
```

### Sample 2

Input:

```
2 2
1 2
3 4
```

Output:

```
2
```

### Sample 3

Input:

```
2 2
1 3
2 4
```

Output:

```
0
```


## ideas
1. 如果 (ai, bi) 要pair在一起, 那么他们中间的都必须被pair掉
2. dp[start, end] 表示把start...end中间的人都安排好的方案数
3. dp[start, end] = dp[start + 1, end - 1] 如果 start 和 end 匹配
4.   + dp[start + 1, i - 1] * dp[i+1, end] 如果start 和i可以匹配
5. 但是这个不不能处理sample 2 (按照这个方式, f(sample2) = 1)
6. 还需要 * (i+1, end) 中的集合数 + 1 (因为 start, i) 可以插入到任何一个地方

## Solution

### 1. The matching must have an interval structure

Suppose Student `start` is eventually paired with Student `i`, where
`start < i`.

Before these two students can be removed, every student originally between
them must already have been removed. Those middle students therefore have to
form pairs entirely inside the interval:

```text
[start+1, i-1]
```

No student inside this interval can be paired with a student to the right of
`i`: such a pair would cross `(start,i)`, and its two endpoints could never
become adjacent while `start` and `i` were both still present.

Consequently, after choosing the partner of the leftmost student, the problem
splits into two independent even-length intervals:

```text
inside: [start+1, i-1]
right:  [i+1, end]
```

This is the reason interval DP works.

Because the inside interval must contain an even number of students,
`i-start-1` must be even. Equivalently, `i-start` is odd. The code enumerates
exactly those positions:

```go
for i := start + 1; i <= end; i += 2
```

### 2. DP state

Define, using inclusive endpoints:

```text
f(start,end) = number of valid operation sequences that remove every
               student in positions start through end
```

The positions refer to the original row. Every state reached by the recurrence
has an even number of students.

For an empty interval:

```text
f(start,end) = 1, when start > end
```

There is exactly one way to remove nothing. This value is also the
multiplicative identity when one side of a split is empty.

The answer is:

```text
f(0, 2N-1)
```

### 3. Choose the partner of `start`

Try every `i` such that:

```text
start < i <= end
i-start is odd
Student start and Student i are on good terms
```

The boolean table `can` stores the last condition. Since `start < i` in the
recurrence, each friendship is normalized and stored as:

```go
can[min(u,v)][max(u,v)] = true
```

Once `(start,i)` is chosen, the inside interval can be removed in:

```text
x = f(start+1, i-1)
```

ways. After all inside students disappear, `start` and `i` become adjacent,
so removing them is forced as the last operation belonging to this left
part. Thus `x` already determines the complete relative order of all
operations involving `[start,i]`; no extra factor is needed for the pair
`(start,i)` itself.

The right interval can independently be removed in:

```text
y = f(i+1, end)
```

ways.

### 4. Why multiplication by a binomial coefficient is necessary

The left and right subproblems are independent, so their operations may be
interleaved.

Let:

```text
cnt1 = (i-start+1)/2  // pairs removed from [start,i]
cnt2 = (end-i)/2      // pairs removed from [i+1,end]
```

After choosing one valid internal order for each side, the relative order of
the `cnt1` left operations must be preserved, and so must the relative order
of the `cnt2` right operations. We only need to choose which `cnt1` positions
among all `cnt1+cnt2` operation slots are occupied by left operations:

```text
C(cnt1+cnt2, cnt1)
```

Therefore, the contribution of pairing `start` with `i` is:

```text
f(start+1, i-1)
* f(i+1, end)
* C(cnt1+cnt2, cnt1)
```

For Sample 2, the pairs are `(1,2)` and `(3,4)`. Either pair may be removed
first, so the answer is:

```text
C(2,1) = 2
```

This is precisely the factor missing from a recurrence that only multiplies
the two interval answers.

### 5. Full recurrence

Combining all possible partners of `start` gives:

```text
f(start,end) = sum over valid i:
    f(start+1, i-1)
    * f(i+1, end)
    * C((end-start+1)/2, (i-start+1)/2)
```

The code writes the two pair counts separately as `cnt1` and `cnt2`, but:

```text
cnt1 + cnt2 = (end-start+1)/2
```

All additions and multiplications are performed modulo `998244353`.

The binomial coefficients are precomputed with Pascal's identity:

```text
C[n][0] = C[n][n] = 1
C[n][k] = C[n-1][k-1] + C[n-1][k]
```

### 6. Correctness proof

We prove by induction on the even length of `[start,end]` that
`f(start,end)` equals the number of valid operation sequences removing that
interval.

#### Base case

If `start > end`, the interval is empty. The algorithm returns `1`, which is
the unique empty operation sequence.

#### Inductive step

Consider a non-empty even-length interval and any valid complete operation
sequence for it. Student `start` is paired with one unique student `i`.
They must be on good terms. All students between them must disappear before
that pair is removed, so the students in `[start+1,i-1]` form a complete
independent subproblem. The remaining students `[i+1,end]` form the other
complete independent subproblem. Both intervals have smaller even lengths.

By the induction hypothesis, the recurrence counts their internal operation
orders exactly with `f(start+1,i-1)` and `f(i+1,end)`. Once the two orders are
fixed, choosing the left operation slots among all operation slots uniquely
determines their interleaving, giving the binomial factor. Hence the recurrence
counts the considered sequence exactly once.

Conversely, take any term generated by the recurrence. Its inside sequence
removes every student between `start` and `i`; then the left sequence can
remove the now-adjacent good pair `(start,i)`. Its right sequence is valid by
the induction hypothesis. Any order-preserving interleaving of these two
sequences remains valid because operations in the disjoint intervals do not
change the relative order within the other interval. Thus every sequence
counted by the recurrence is achievable.

Every valid sequence chooses exactly one partner `i` for `start`, so different
terms cannot count the same sequence. Therefore the recurrence is correct.

### 7. Memoization and complexity

There are `O(N^2)` intervals among the `2N` students. Each interval tries
`O(N)` possible partners for its left endpoint, so the interval DP costs
`O(N^3)` time. Pascal's table takes `O(N^2)` time, which does not change the
overall bound.

The `can`, `dp`, and binomial-coefficient tables all use `O(N^2)` space.

```text
Time:  O(N^3)
Space: O(N^2)
```
