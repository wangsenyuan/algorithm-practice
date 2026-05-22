# D. Division Versus Addition

https://codeforces.com/problemset/problem/2152/D

**Time Limit:** 2 seconds

**Memory Limit:** 1024 megabytes

## Problem

For an array `b = [b_1, b_2, ..., b_m]` with every `b_i >= 2`, consider a
two-player game played by Poby and Rekkles.

- Poby moves first.
- On Poby's turn, he chooses an element `x >= 2` and replaces it with
  `floor(x / 2)`.
- On Rekkles's turn, he chooses an element `x >= 2` and replaces it with
  `x + 1`.
- The game ends once all elements are equal to `1`.

The score is the number of moves made by Poby. Poby wants to minimize the
score, while Rekkles wants to maximize it.

The value of an array `b` is the final score when both players play optimally.

You are given an array `a` of length `n`, where every `a_i >= 2`. Answer `q`
independent queries. For each query `[l, r]`, find the value of the subarray:

```text
[a_l, a_{l+1}, ..., a_r]
```

## Solution

For one number `x`, the important question is who touches it first.

Define `play(x)` as the number of Poby moves needed if the current value is `x`
and it is Poby's turn to operate on this number. The local code simulates the
forced sequence:

```text
Poby:    x -> floor(x / 2)
Rekkles: x -> x + 1      if the value is still greater than 1
Poby:    ...
```

until the number becomes `1`.

So for an original value `a_i`:

- If Poby is the first player to touch this element, its cost is `play(a_i)`.
- If Rekkles is the first player to touch this element, he first changes it to
  `a_i + 1`, so its cost is `play(a_i + 1)`.

The difference:

```text
play(a_i + 1) - play(a_i)
```

is always either `0` or `1`. Rekkles only gains anything by being first on
elements where this difference is `1`.

For a query interval, let:

```text
A = sum play(a_i)
B = sum play(a_i + 1)
diff = B - A
```

Here `diff` is the number of elements where Rekkles moving first would add one
extra Poby move.

The game alternates turns and Poby moves first. On those `diff` valuable
elements, Poby can choose to touch about half of them first and prevent the
extra cost. More precisely, Poby prevents:

```text
ceil(diff / 2)
```

of the possible extra moves. Starting from the pessimistic total `B`, the value
of the interval is therefore:

```text
B - ceil(diff / 2)
```

The code computes this as:

```text
res = B
res -= (diff + 1) / 2
```

To answer range queries quickly, build prefix sums of both costs:

```text
dp[i]  = sum play(a_1 ... a_i)
dp2[i] = sum play(a_1 + 1 ... a_i + 1)
```

Then each query obtains `A` and `B` by subtraction in `O(1)`.

The preprocessing is `O(n log max(a_i))`, because `play(x)` repeatedly halves
the number. Each query is answered in `O(1)`.
