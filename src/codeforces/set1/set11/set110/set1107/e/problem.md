# E. Vasya and Binary String

[Problem link](https://codeforces.com/problemset/problem/1107/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Vasya has a string `s` of length `n` consisting only of digits `0` and `1`.
Also he has an array `a` of length `n`.

Vasya performs the following operation until the string becomes empty:
choose some consecutive substring of equal characters, erase it from the
string and glue together the remaining parts (any of them can be empty).
For example, if he erases substring `111` from string `111110` he will get
the string `110`. Vasya gets `a_x` points for erasing a substring of length `x`.

Maximize the total points.

## Constraints

- `1 <= n <= 100`
- `s` consists only of `0` and `1`
- `1 <= a_i <= 10^9`

## Input

```
n
s
a1 a2 ... an
```

## Output

Print one integer — the maximum total points.

## Samples

### Sample 1

Input:

```
7
1101001
3 4 9 100 1 2 3
```

Output:

```
109
```

Note: one optimal sequence is
`1101001 → 111001 → 11101 → 1111 → ∅`.

### Sample 2

Input:

```
5
10101
3 10 15 15 15
```

Output:

```
23
```

Note: one optimal sequence is
`10101 → 1001 → 11 → ∅`.

## ideas
1. 怎么表示不同的状态?
2. dp[l][r] 表示l...r中间的被删除后, 能够得到的最优解?
3. 但是问题在于, 11100001110000111 那么可以删除中间的部分, 但是也可以保留中间的部分, 删除靠近两端中间的部分, 从而得到更优解
4. 把 s编码成{l1, l2, ... lk} 表示第一段相同的是l1, 第二种相同的是l2, ..., 相邻的都不相同
5. 对于一段内的, li, 要么把它全删除掉(通过使用最优的方案), 从而得到 [l0, ..l[i-2], l[i-1] + l[i+1], l[i+2], ...lk]
6. 存在某种保留部分状态的情况吗? (似乎是不存在的) 这种情况下, 要么就是先从其他状态出发处理
7. 感觉前进了一步~
8. dp[i][j][0/1] 表示到i为止, 前面的都被处理成了相同的字符(0或者1), 且长度为j时的最优解
9. 当前如果是和前面相同的字符, 那么可以增加j, dp[i+1][j+1][0/1]
10. 如果是和不一样的字符, 那么要么把前面的j个都用掉 => dp[i+1][1][1/0]
11. 或者把接下来相同的, 都用掉进入状态 dp[i+x][j][0/1] (消耗完下一段的, 保留当前的, 以获得远期收益)
12.

## Detailed solution

### 1. Why an ordinary interval DP is insufficient

A first idea might be:

```text
dp[l][r] = maximum score for completely deleting s[l:r]
```

However, this state loses information needed by the surrounding interval.
After deleting some characters in the middle, `s[l]` can become adjacent to
equal characters that originally lay outside or farther inside the interval.
The value of keeping `s[l]` depends on how many equal characters are already
connected to it.

For example, if several `1`s are separated by blocks of `0`s, deleting those
zero blocks can merge the `1`s into a larger group. Because `a[x]` is arbitrary
and need not be linear, the future value of that larger group cannot be derived
from `dp[l][r]` alone.

We therefore add the size of the already-connected group to the state.

### 2. DP state

Use a half-open interval and define:

```text
f(l, r, k) = maximum score obtainable from s[l:r], assuming that
             s[l] is already connected to k-1 equal characters
             immediately to its left
```

Including `s[l]` itself, the current leading group has size `k`.

The extra `k-1` characters are not physically inside `s[l:r]`; they represent
characters retained by an outer recursive state. They must eventually be
deleted together with `s[l]` and possibly more equal characters.

The answer is:

```text
f(0, n, 1)
```

Initially, `s[0]` has no previously attached characters, so the leading group
has size one.

### 3. Base cases

For an empty interval, there is nothing to delete:

```text
f(l, l, k) = 0
```

If the interval contains only `s[l]`, it must be deleted together with its
`k-1` attached characters. The total group size is `k`, so:

```text
f(l, l+1, k) = a[k]
```

The code stores `a` with zero-based indices, hence it uses `a[k-1]`.

### 4. Transition 1: delete the leading group now

One valid choice is to delete the group containing `s[l]` before it merges
with any later equal character. Its current size is `k`, giving `a[k]` points.
The remainder starts fresh at `s[l+1]` with a group size of one:

```text
f(l, r, k) = a[k] + f(l+1, r, 1)
```

In the code:

```go
res = a[prefix-1] + f(start+1, end, 1)
```

### 5. Transition 2: merge with a later equal character

Suppose `i` is a later position satisfying:

```text
l < i < r
s[i] = s[l]
```

If every character in `s[l+1:i]` is deleted first, then `s[l]` becomes
adjacent to `s[i]`. We deliberately keep the group containing `s[l]`, so its
`k` characters merge with `s[i]`. From position `i` onward, the leading group
therefore has size `k+1`.

The two parts contribute:

```text
f(l+1, i, 1)  // completely delete the characters between l and i
f(i, r, k+1)  // continue with s[i] attached to the retained leading group
```

Thus:

```text
f(l, r, k) = max(
    f(l, r, k),
    f(l+1, i, 1) + f(i, r, k+1)
)
```

The code tries every later `i` with `s[i] == s[l]`:

```go
for i := start + 1; i < end; i++ {
    if s[i] == s[start] {
        res = max(res,
            f(start+1, i, 1)+f(i, end, prefix+1))
    }
}
```

Notice that we do not score the leading group during this transition. It has
not been deleted yet; its score will be earned later by the recursive state
`f(i,r,k+1)`.

### 6. Why these transitions are complete

Consider an optimal strategy for `f(l,r,k)` and focus on the group containing
`s[l]`.

There are only two possibilities:

1. It is deleted without ever joining a later original character equal to
   `s[l]`. Deleting it immediately is at least as general: operations on the
   remaining characters do not depend on when this separate group is removed.
   This is Transition 1.
2. It eventually joins some later equal character. Let `i` be the first such
   original position that joins it. Before that can happen, every character in
   `s[l+1:i]` must be deleted. Their optimal contribution is
   `f(l+1,i,1)`. The retained group then gains `s[i]`, giving exactly the state
   `f(i,r,k+1)`. This is one of the choices in Transition 2.

Therefore every optimal strategy is covered by at least one transition. Every
transition is also achievable by the described deletion order, so the
recurrence returns exactly the optimum.

### 7. Memoization

The same state `(l,r,k)` can be reached through several different merge
choices. The three-dimensional table stores each computed result:

```text
dp[l][r][k] = f(l,r,k)
```

All rewards are positive, so zero safely means that a non-empty state has not
yet been computed. Empty intervals return immediately and do not need to be
stored.

### 8. Example of why keeping a character helps

For a pattern such as:

```text
1 ... 1 ... 1
```

we can delete the blocks between the `1`s first. The recursive states then
evolve conceptually as:

```text
f(first_1, r, 1)
    -> f(second_1, r, 2)
    -> f(third_1, r, 3)
```

Eventually the three connected `1`s can be removed for `a[3]` points. The
parameter `k` is precisely the information that an ordinary two-dimensional
interval DP would lose.

### 9. Complexity

There are `O(n^3)` possible states `(l,r,k)`. A state may scan `O(n)` later
positions looking for an equal character, so:

```text
Time:  O(n^4)
Space: O(n^3)
```

With `n <= 100`, this is within the intended bounds.

## About an `O(n^3)` solution

It is tempting to call the preceding algorithm `O(n^3)` because its memo table
has three indices. That is not its time complexity: every state also contains
the loop over `i`, so its worst-case running time is `O(n^4)`.

I could not derive a correct `O(n^3)` algorithm for the problem with arbitrary
values of `a`. The following last-operation formulation explains both where an
apparent cubic DP comes from and where the missing factor appears.

### 1. Describe the last deletion

Consider the last operation used to delete an interval `[l,r)`. The characters
removed by that operation:

- all have the same value, either `0` or `1`;
- form a subsequence of the original interval;
- become consecutive only after every gap between consecutive chosen
  characters has been completely deleted.

Suppose the retained subsequence uses positions

```text
p1 < p2 < ... < pk
```

Then the score has the form

```text
best(l, p1)
+ best(p1+1, p2)
+ ...
+ best(pk+1, r)
+ a[k]
```

where every `best(x,y)` is an already-solved interval that is deleted before
the final operation.

### 2. Auxiliary DP for one interval

Fix the final character `c` and scan from left to right. Define

```text
g[i][k] = maximum score after processing up to position i,
          while retaining k occurrences of c for the final deletion
```

At a position equal to `c`, we may retain it and move from `k` to `k+1`.
Alternatively, we may choose a whole following segment as a gap, add that
segment's already-known `best` value, and jump past it. After reaching the
right boundary, add `a[k]`.

There are `O(n^2)` states `(i,k)`. Trying the endpoint of the next deleted gap
takes another `O(n)`, so computing this auxiliary DP for one starting boundary
costs `O(n^3)`.

### 3. Why this is not an overall `O(n^3)` algorithm

The interval values used as gap scores are themselves subproblems. Repeating
the preceding `O(n^3)` computation independently for every interval would cost
`O(n^5)`.

Intervals with the same left boundary can share their scan states: while
extending the right boundary, the states needed for the shorter intervals have
already been computed. This removes one factor and gives `O(n^4)` overall, but
not `O(n^3)`.

The carry recurrence implemented in `solution.go` reaches the same bound in a
cleaner way:

```text
O(n^3) states * O(n) merge choices = O(n^4)
```

To reduce it to cubic time, the maximum over all matching merge positions
would have to be maintained in `O(1)` amortized time. For fixed `(l,r,k)`, that
maximum is

```text
max over i with s[i] = s[l] of
    f(l+1, i, 1) + f(i, r, k+1)
```

The reward array is arbitrary: it is not required to be monotone, convex, or
concave. Consequently, there is no justified monotonicity of the best `i`, nor
a Monge/convex property that would allow divide-and-conquer or Knuth-style
optimization. Merely removing the loop, or counting only the memo states,
would therefore give an incorrect complexity claim.

So the verified general solution described here is `O(n^4)`, not `O(n^3)`.
This also matches [Neal's alternative
editorial](https://codeforces.com/blog/entry/64833), which gives the same
three-dimensional carry DP in `O(n^4)` and explicitly asks whether an
`O(n^3)` solution exists.
