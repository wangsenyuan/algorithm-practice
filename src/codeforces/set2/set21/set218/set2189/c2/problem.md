# C2 - XOR-convenience (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2189/C2)

**Contest:** [Codeforces Round 1072 (Div. 1)](https://codeforces.com/contest/2189)

time limit: 2 seconds per test

memory limit: 256 MB

This is the hard version of XOR-convenience. Find a permutation `p` of length `n` such that for every
`i` with `1 <= i <= n - 1`, there exists `j` with `i <= j <= n` and

`p_i = p_j XOR i`

or report that no such permutation exists.

The easy version only requires the condition for `2 <= i <= n - 1`, and always has a solution.

## Constraints

- 1 <= t <= 10^4
- 3 <= n <= 2 * 10^5
- sum of `n` over all test cases <= 2 * 10^5

## Input

The first line contains `t` — the number of test cases.

Each test case contains one integer `n`.

## Output

For each test case, if a suitable permutation exists, print `n` integers `p_1, ..., p_n`. Otherwise
print `-1`.

If multiple answers exist, print any of them.

## Sample Input

```text
2
3
4
```

## Sample Output

```text
2 1 3
-1
```

## Note

For `n = 3`, permutation `[2, 1, 3]` works because `p_2 = 1` and `p_3 XOR 2 = 1`.

For `n = 4`, no suitable permutation exists.

## Solution

The condition can be rewritten as:

```text
p_j = p_i XOR i
```

for some `j >= i`. So for every position `i < n`, we only need to make sure that
the value `p_i XOR i` appears somewhere at or after position `i`.

### When there is no solution

If `n` is a power of two, there is no solution.

Let `n = 2^k`. For every `1 <= i < n`:

```text
n XOR i > n
```

because `n` has the `2^k` bit set and `i` uses only lower bits. Therefore, if
the value `n` were placed at any position `i < n`, then `p_i XOR i` would be
outside the permutation range `1..n`, which is impossible.

So `n` must be placed at the last position:

```text
p_n = n
```

Now check `i = n-1`. The only possible `j` values are `n-1` and `n`.

- `j = n-1` would require `p_{n-1} = p_{n-1} XOR (n-1)`, impossible because
  `n-1 != 0`.
- `j = n` would require `p_{n-1} = n XOR (n-1)`, but for `n = 2^k`,
  this equals `2n-1`, which is larger than `n`.

Thus powers of two are impossible.

### Main construction idea

For all other `n`, we construct a valid permutation directly.

The useful identity is:

```text
(i XOR 1) XOR i = 1
```

for any integer `i`.

So if we set:

```text
p_i = i XOR 1
```

then position `i` can use the later value `1` as its witness, because:

```text
p_i XOR i = 1
```

The construction keeps value `1` near the end, so many earlier positions can all
use the same witness.

### Odd `n`

For odd `n`, use:

```text
p_1 = n - 1
p_i = i XOR 1       for 2 <= i <= n-2
p_{n-1} = 1
p_n = n
```

This is a permutation:

- `p_1 = n-1`.
- The middle values `i XOR 1` for `2 <= i <= n-2` are exactly the numbers
  `2, 3, ..., n-2`, just swapped in pairs.
- The last two values are `1` and `n`.

Now verify the condition.

For `i = 1`:

```text
p_1 XOR 1 = (n-1) XOR 1 = n
```

because `n` is odd. Value `n` is at position `n`.

For every `2 <= i <= n-2`:

```text
p_i XOR i = (i XOR 1) XOR i = 1
```

and value `1` is at position `n-1`, which is after all such `i`.

For `i = n-1`:

```text
p_{n-1} XOR (n-1) = 1 XOR (n-1) = n
```

again because `n` is odd. Value `n` is at position `n`.

So the odd construction works.

### Even `n` that is not a power of two

For even non-power-of-two `n`, define:

```text
d = n XOR (n-1)
```

Here `d` is odd, `3 <= d <= n-1`, and the value `d-1` is even.

Use:

```text
p_1 = d XOR 1
p_i = i XOR 1       for 2 <= i <= n-2
p_d = n
p_{n-1} = 1
p_n = n - 2
```

The assignment `p_d = n` overwrites the middle value that would have been
`d XOR 1 = d-1`, and `p_1` supplies that missing value. The final value `n-2`
is also not produced by the middle range, so `p_n = n-2` supplies it. Therefore
all values `1..n` appear exactly once.

Now verify the condition.

For `i = 1`:

```text
p_1 XOR 1 = (d XOR 1) XOR 1 = d
```

Value `d` appears in the middle range at position `d-1`, so it is after position
`1`.

For `2 <= i <= n-2` and `i != d`:

```text
p_i XOR i = (i XOR 1) XOR i = 1
```

Value `1` is at position `n-1`, so it is after every such `i`.

For the special position `i = d`:

```text
p_d XOR d = n XOR d
```

Since `d = n XOR (n-1)`, this becomes:

```text
n XOR (n XOR (n-1)) = n-1
```

Value `n-1` appears at position `n-2`, and `d <= n-2`, so the witness is after
position `d`.

For `i = n-1`:

```text
p_{n-1} XOR (n-1) = 1 XOR (n-1) = n-2
```

Value `n-2` is at position `n`, so the witness is valid.

Thus the even non-power-of-two construction also works.

### Complexity

The permutation is built in one pass.

The time complexity is:

```text
O(n)
```

The memory complexity is:

```text
O(n)
```
