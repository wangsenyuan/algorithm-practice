# B - Valid Arrays by K-Divisible Swaps

[Problem link](https://atcoder.jp/contests/arc223/tasks/arc223_b)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 500 points

## Problem Statement

You are given an integer sequence `A` of length `N` and a positive integer `K`.
You can perform the following operation on `A` any number of times:

- Choose an integer `i` with `1 <= i <= N-1` such that `A_i + A_{i+1}` is
  divisible by `K`, and swap `A_i` and `A_{i+1}`.

Find the total number, modulo `998244353`, of distinct sequences that `A` can
become after performing the operation zero or more times.

Solve `T` test cases per input.

## Constraints

- `1 <= T <= 10^5`
- `2 <= N <= 2 * 10^5`
- `1 <= K <= 10^9`
- `1 <= A_i <= 10^9`
- The sum of `N` over all test cases is at most `2 * 10^5`
- All input values are integers

## Input

```
T
case_1
case_2
...
case_T
```

Each test case:

```
N K
A_1 A_2 ... A_N
```

## Output

Output `T` lines. The `t`-th line should contain the answer for the `t`-th
test case.

## Samples

### Sample 1

Input:

```
1
4 3
1 2 4 7
```

Output:

```
4
```

The reachable sequences are `(1,2,4,7)`, `(1,4,2,7)`, `(1,4,7,2)`,
`(2,1,4,7)`.

### Sample 2

Input:

```
1
6 4
1 5 3 6 2 4
```

Output:

```
6
```

### Sample 3

Input:

```
1
3 2
2 2 3
```

Output:

```
1
```


## Solution

Whether two adjacent values can be swapped depends only on their remainders
modulo `K`.

Let the remainder of a value be `x`. It can swap only with a value whose
remainder is:

```text
y = K - x
```

where remainder `K` is represented by `0`. There are two different cases.

### Case 1: `x` and `K-x` are different

This is the normal case:

```text
x != 0 and 2*x != K
```

Starting at the current position, the code takes the maximal consecutive block
whose remainders are either `x` or `y = K-x`.

No element can cross either boundary of this block. A remainder outside
`{x, y}` is compatible with neither `x` nor `y`, so no matter how the block is
rearranged, the boundary pair can never be swapped. Therefore different
maximal blocks can be counted independently.

Inside one block:

- an `x`-element can swap with a `y`-element;
- two `x`-elements cannot swap;
- two `y`-elements cannot swap.

Consequently, the relative order of all `x`-elements is fixed, and the
relative order of all `y`-elements is also fixed. The only choice is how to
interleave these two ordered lists.

If the block contains `cnt[0]` elements of remainder `x` and `cnt[1]` elements
of remainder `y`, choose which `cnt[0]` of the total positions receive the
`x`-elements:

```text
C(cnt[0] + cnt[1], cnt[0])
```

Every such interleaving is reachable using adjacent swaps between the two
different remainder classes. Also, different interleavings have different
remainder sequences, so they represent different resulting arrays even when
equal values occur within one class.

For example, with `K = 3`, the values `1, 4, 2` have remainders `1, 1, 2`.
The values `1` and `4` must remain in that order, but `2` may be inserted in
any of the three positions:

```text
1 4 2
1 2 4
2 1 4
```

The order `4, 1, 2` is impossible because it reverses the two remainder-`1`
elements.

### Case 2: the remainder is its own complement

This happens when:

```text
x = 0
```

or, when `K` is even:

```text
x = K/2
```

Equivalently, the code checks `x == 0 || 2*x == K`.

Now any two elements with remainder `x` can be swapped, because their
remainder sum is divisible by `K`. Thus, inside a maximal consecutive block of
this remainder, arbitrary adjacent swaps are possible, so every permutation of
the block is reachable.

If the block length is `L`, there would be `L!` permutations if all values
were distinct. Equal values do not create distinct arrays, so for every value
appearing `freq[v]` times, divide by `freq[v]!`:

```text
L! / product(freq[v]!)
```

The implementation sorts the block, counts equal consecutive values, and
computes this multinomial coefficient as:

```text
F[L] * product(I[freq[v]])
```

where `F[t] = t!` and `I[t] = 1/t!` modulo `998244353`.

### Combining the blocks

The array is partitioned into maximal independent blocks of the two forms
above. Operations in one block never move an element into another block, while
every arrangement counted for a block is reachable inside it. Therefore the
choices for all blocks are independent, and the final answer is the product of
their contributions modulo `998244353`.

The helper `nCr` uses the precomputed factorial and inverse-factorial arrays to
evaluate each binomial coefficient in constant time.

### Correctness Proof

Consider the maximal block beginning at the current unprocessed position.

- If its first remainder `x` is different from `K-x`, exactly the remainders
  `x` and `K-x` can swap with each other. Elements of the same remainder can
  never cross, so their two internal orders are invariant. Conversely, any
  interleaving of the two orders can be produced by adjacent cross-remainder
  swaps. Hence this block contributes
  `C(cnt[0]+cnt[1], cnt[0])` distinct arrays.
- If `x` is `0` or `K/2`, every adjacent pair in the same-remainder block is
  swappable. Adjacent swaps generate every permutation, and removing duplicate
  permutations of equal values gives
  `L! / product(freq[v]!)` distinct arrays.

In either case, a remainder outside the block is incompatible with every
remainder inside it, so no operation can move an element across the block
boundary. Thus blocks are independent. Multiplying the correct contribution
of every block counts every reachable final array exactly once.

### Complexity

Factorials and inverse factorials are precomputed once in `O(MAXN)` time and
`O(MAXN)` space.

For one test case, complementary-remainder blocks are scanned linearly.
Self-complementary blocks are sorted; their total length is at most `N`.

- Time: `O(N log N)` per test case in the worst case
- Extra working space: `O(N)` in the worst case
