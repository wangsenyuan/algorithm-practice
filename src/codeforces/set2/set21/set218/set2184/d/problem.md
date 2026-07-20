# D. Unfair Game

[Problem link](https://codeforces.com/problemset/problem/2184/D)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Bob has thought of a number from `1` to `n`, where `n = 2^d` for some
non-negative integer `d`. Initially, Alice knows whether the chosen number is
even or not.

Only Alice takes turns. In one move, she can either:

- subtract `1`, or
- if the current number is even, halve it.

After each move, Bob responds with either `-1` (the number became `0` and
Alice won), or the unique non-negative integer `x` such that the current
number `a` is divisible by `2^x` but not by `2^(x+1)`. In other words, Bob
reports the number of trailing zeroes in the binary representation of `a`.

The game lasts at most `k` moves. Count how many initial numbers `a` in
`[1, n]` are such that Alice, playing optimally, **cannot** win in at most
`k` moves.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n, k ≤ 10^9`
- `n = 2^d` for some non-negative integer `d`

## Input

```
t
n k
...
```

## Output

For each test case, print the number of integers from `1` to `n` for which
Alice cannot win in at most `k` moves.

## Samples

### Sample 1

Input:

```
1
4 1
```

Output:

```
3
```

`a = 2, 3, 4` need more than 1 move; `a = 1` wins in 1 move.

### Sample 2

Input:

```
1
4 2
```

Output:

```
2
```

### Sample 3

Input:

```
1
4 3
```

Output:

```
0
```

### Sample 4

Input:

```
1
4 4
```

Output:

```
0
```

### Sample 5

Input:

```
1
4 5
```

Output:

```
0
```

### Sample 6

Input:

```
1
16 5
```

Output:

```
4
```

### Sample 7

Input:

```
1
16 1
```

Output:

```
15
```

### ideas
1. given a, how alice can guess it?
2. if a is odd, then alice can only subtract 1 from it to get a - 1
3. if a is even, can half it.
4. then BOB response x, a % (1 << x), that x is then length of tailing zeros of a (after alice operation)
5. if x = 0, which means a is odd 
6. else x > 0, alice half it (x times)
7. 对于一个数w, 它的操作次数 = 它的长度(这个是half的数量) + 其中1的个数(这个是操作1的个数)
8. 那么当k > 60的时候 0, 肯定可以在60次内完成
9. 否则的话, 就是计算有多少个数的操作数 > k
10. 这个可以用digit dp?

## Solution

### How many moves does one number need?

For a positive number `a`, Alice's optimal action is forced by the lowest
binary bit:

- If `a` is odd, halving is forbidden, so she subtracts `1`. This changes one
  set bit into `0`.
- If `a` is even, halving removes one trailing zero, which is always better
  than subtracting `1` and making the number odd.

Bob's answer gives exactly the number of trailing zeroes, so Alice always
knows whether she can halve and can follow this strategy without knowing the
original number.

Suppose `a` has binary length `len(a)` and contains `popcount(a)` set bits.

- Exactly `popcount(a)` subtractions are needed: each subtraction removes one
  set bit.
- Exactly `len(a) - 1` halvings are needed: they remove every binary position
  below the leading bit.

Therefore,

```text
moves(a) = len(a) - 1 + popcount(a).
```

The problem is now purely combinatorial: count the numbers `a` in `[1,n]`
for which `moves(a) > k`.

### Counting by binary length

Let `m` be the binary length of `n`. First count all numbers whose binary
length is `i`, for `1 <= i < m`.

Every such number has a fixed leading `1` and `i-1` remaining positions. If
exactly `j` of those positions are `1`, then

```text
popcount(a) = j + 1
moves(a)    = (i - 1) + (j + 1) = i + j.
```

There are

```text
C(i-1, j)
```

numbers of this form. The code adds this binomial coefficient whenever
`i+j > k`.

The table `C` is Pascal's triangle:

```text
C(i,j) = C(i-1,j-1) + C(i-1,j).
```

### Numbers with the same binary length as `n`

The implementation also scans the bits of `n` from left to right in the
shape of a binary digit DP. At a `1` bit, choosing `0` there makes the number
strictly smaller than `n`; the remaining suffix can then be counted by how
many set bits it contains.

For this problem, however, `n` is guaranteed to be a power of two. Its binary
representation is `100...0`, so there is no later `1` bit at which such a
smaller `m`-bit number can branch. Consequently, after counting all shorter
lengths, the only remaining number is `n` itself.

For `n = 2^(m-1)`, we have `popcount(n) = 1`, hence

```text
moves(n) = m - 1 + 1 = m.
```

The final check adds `n` exactly when `m > k`.

The early return for `k > 2m` is safe because a number with at most `m`
binary digits can require fewer than `2m` moves.

## Correctness Proof

We prove that `solve(n,k)` returns the number of initial values Alice cannot
reduce to zero within `k` moves.

### Lemma 1

For every positive integer `a`, the minimum number of moves needed to reach
zero is `len(a) - 1 + popcount(a)`.

**Proof.** If the current number is odd, Alice must subtract `1`, removing its
lowest set bit. If it is even, halving removes one trailing binary position;
subtracting instead creates an odd number and cannot use fewer moves. During
the complete process, every set bit is removed by exactly one subtraction,
and every position below the leading bit is removed by exactly one halving.
Thus there are `popcount(a)` subtractions and `len(a)-1` halvings. ∎

### Lemma 2

For a fixed binary length `i < m` and a fixed `j`, the algorithm counts
exactly all `i`-bit numbers requiring `i+j` moves.

**Proof.** An `i`-bit number has one fixed leading set bit. Choosing `j` of
the remaining `i-1` positions produces exactly `C(i-1,j)` numbers, each with
`j+1` set bits. By Lemma 1, each requires
`(i-1)+(j+1)=i+j` moves. ∎

### Lemma 3

After the shorter binary lengths are counted, the algorithm correctly
handles every remaining valid number in `[1,n]`.

**Proof.** Because `n` is a power of two, it is the smallest number having
binary length `m`. Thus the only `m`-bit number not already handled is `n`
itself. It has one set bit and requires `m` moves by Lemma 1. The final check
counts it exactly when `m > k`. ∎

### Theorem

The algorithm returns the required answer.

**Proof.** By Lemma 2, the nested loops count exactly the shorter numbers
whose required move count exceeds `k`. By Lemma 3, the final check does the
same for the only remaining number, `n`. These groups are disjoint and cover
all integers in `[1,n]`. Therefore their total is exactly the number of games
Alice cannot win in at most `k` moves. ∎

## Complexity Analysis

Let `m = floor(log2(n)) + 1`. Under the constraints, `m <= 30`.

- Time: `O(m^2)` per test case.
- Space: `O(m^2)` for the binomial-coefficient table.
