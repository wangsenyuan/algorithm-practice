# D. XOR, Expression and Two Binary Numbers

[Problem link](https://codeforces.com/problemset/problem/2234/D)

**Contest:** [Codeforces Round 1102 (Div. 2)](https://codeforces.com/contest/2234)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an integer `k`. There is a sequence of `n`-bit binary numbers
`a_1, a_2, ..., a_{2^k + 1}`. The numbers `a_1` and `a_{2^k + 1}` are given to you, and the
others are unknown. Then the unknown numbers are filled in as follows over `k` steps:

- Suppose that before the `i`-th step, the numbers with indices `p_1 < p_2 < ... < p_m` have
  already been filled. (Before the first step, these are the numbers with indices `1` and
  `2^k + 1`.)
- Then for each `j` from `1` to `m - 1`, the assignment
  `a_{(p_j + p_{j + 1}) / 2} := a_{p_j} XOR a_{p_{j + 1}}` is performed.
- These assignments happen simultaneously, and after that all these numbers also become filled.

It can be shown that this process always fills all numbers completely.

This is what the process looks like for `k = 2` and `n = 3`, where initially
`a_1 = 010` and `a_5 = 110`:

- Before the first step, the numbers with indices `1, 5` are filled. Therefore, this operation
  will perform `a_3 := a_1 XOR a_5 = 010 XOR 110 = 100`.
- Before the second step, the numbers with indices `1, 3, 5` are filled. Therefore, this operation
  will perform `a_2 := a_1 XOR a_3 = 110` and `a_4 := a_3 XOR a_5 = 010`.

You need to compute the following expression:

```text
x_1 * y_1 + x_2 * y_2 + ... + x_{2^k + 1} * y_{2^k + 1}
```

where `x_i` is the number of set bits in the `i`-th number, and `y_i` is the number of zero bits in
the `i`-th number.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t`
(`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `k` (`1 <= n <= 10^5`,
`1 <= k <= 30`) — the length of the binary numbers and the number determining the length of the
sequence.

The second line of each test case contains a binary string `s` of length `n` — the value of `a_1`.

The third line of each test case contains a binary string `z` of length `n` — the value of
`a_{2^k + 1}`.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, output one integer — the value of the expression from the statement.

## Example

### Input

```text
4
3 2
010
110
1 1
0
0
2 2
01
00
7 30
1010111
0011010
```

### Output

```text
10
0
3
12169074016
```

### Note

- In the first test case, the process was described in the statement. The resulting sequence of
  binary numbers is `[010, 110, 100, 010, 110]`. Then the expression equals
  `1 * 2 + 2 * 1 + 1 * 2 + 1 * 2 + 2 * 1 = 10`.
- In the second test case, at the first step we have `a_2 = a_1 XOR a_3 = 0 XOR 0 = 0`. Therefore,
  the resulting sequence is `[0, 0, 0]`, and the expression is zero.

## Solution

For one generated binary number, if it has `x` ones and `y` zeroes, then `x * y` counts unordered
pairs of bit positions whose bits are different. So instead of trying to build every whole
`n`-bit number, count the contribution of every pair of columns.

Take two bit positions `p` and `q`. For every generated number `a_i`, this pair contributes `1`
exactly when:

```text
a_i[p] XOR a_i[q] = 1
```

The midpoint construction is also based on XOR, so this value evolves independently as a one-bit
version of the same process. Its two endpoints are:

```text
left  = s[p] XOR s[q]
right = z[p] XOR z[q]
```

Therefore, for all column pairs with the same `(left, right)`, the total contribution is the same:
the number of ones in the generated one-bit sequence.

### Counting column pairs

Each original column has only four possible types, based on its endpoint bits:

```text
type 0: (s = 0, z = 0)
type 1: (s = 0, z = 1)
type 2: (s = 1, z = 0)
type 3: (s = 1, z = 1)
```

Let `cnt[type]` be the number of columns of this type. For every pair of types `(i, j)`, the number
of unordered column pairs is:

- `cnt[i] * cnt[j]`, if `i != j`;
- `cnt[i] * (cnt[i] - 1) / 2`, if `i == j`.

For this pair of types, the one-bit endpoints are:

```text
left  = s_i XOR s_j
right = z_i XOR z_j
```

The answer adds:

```text
number_of_pairs * score[left][right]
```

where `score[a][b]` is the number of ones produced by the one-bit process after `k` steps, starting
from endpoints `a` and `b`.

### Computing `score`

Let `dp[a][b]` be the number of ones in the full one-bit sequence generated from endpoints `a` and
`b` after the current number of steps.

For zero steps, the sequence is just `[a, b]`, so:

```text
dp[a][b] = a + b
```

When one more step is added, the first midpoint is:

```text
x = a XOR b
```

The final sequence is the left generated sequence from `(a, x)` plus the right generated sequence
from `(x, b)`. The midpoint `x` is counted in both halves, so subtract it once:

```text
next[a][b] = dp[a][x] + dp[x][b] - x
```

Repeating this transition `k` times gives all four `score` values.

### Correctness

For any fixed generated number, `x_i * y_i` counts exactly the unordered bit-position pairs whose
two bits are different in that number. Summing this over all generated numbers is therefore the same
as summing, for every unordered column pair, how many generated numbers make those two columns
different.

For a fixed column pair `(p, q)`, the value `a_i[p] XOR a_i[q]` follows the same recursive midpoint
rule as the original numbers: the midpoint value is the XOR of the two endpoint values. Thus its
whole behavior depends only on the two endpoint differences `s[p] XOR s[q]` and `z[p] XOR z[q]`.
The `score` recurrence counts exactly the number of ones in that one-bit generated sequence.

The algorithm groups all column pairs by the four possible column endpoint types, counts how many
pairs belong to each group, and multiplies by the matching `score`. This covers every unordered
pair of bit positions exactly once, so the produced sum equals the required expression.

### Complexity

Counting column types takes `O(n)`. Computing `score` takes `O(k)` over only four states. The final
pair-type loop has constant size.

Total complexity per test case is `O(n + k)`, and the memory usage is `O(1)`.
