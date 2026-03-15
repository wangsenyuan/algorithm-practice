### Problem

You are given two non-negative integers `x` and `y`.

Find two non-negative integers `p` and `q` such that:

- `p & q = 0`
- `|x - p| + |y - q|` is minimized

Here, `&` denotes the bitwise AND operation.

### Input

Each test contains multiple test cases.

- The first line contains the number of test cases `t` (`1 <= t <= 10^4`).
- Each test case contains two non-negative integers `x` and `y` (`0 <= x, y < 2^30`).

### Output

For each test case, output two non-negative integers `p` and `q`.

If there are multiple valid pairs, you may output any of them.

It can be proven that every valid solution satisfies:

- `max(p, q) < 2^31`

### Example

**Input**
```text
7
0 0
1 1
3 6
7 11
4 4
123 321
1073741823 1073741822
```

**Output**
```text
0 0
2 1
3 8
6 9
4 3
128 321
1073741824 1073741822
```

### Note

In the first test case, one valid pair is `p = 0` and `q = 0`, because:

- `0 & 0 = 0`
- `|x - p| + |y - q| = |0 - 0| + |0 - 0| = 0`

This is the minimum possible value.

In the third test case, one valid pair is `p = 3` and `q = 8`, because:

- `3 & 8 = 0`
- `|x - p| + |y - q| = |3 - 3| + |6 - 8| = 2`

This is optimal. Note that `(p, q) = (3, 4)` is also a valid answer.

### Key Insights

1. The condition `p & q = 0` means that for every bit position, we may choose only one of three patterns:
   - `(0, 0)`
   - `(1, 0)`
   - `(0, 1)`

2. The objective `|x - p| + |y - q|` is separable by bits only if we also know whether the already chosen higher bits make:
   - `p < x`, `p = x`, or `p > x`
   - `q < y`, `q = y`, or `q > y`

3. That suggests a digit-DP from the highest bit to the lowest bit.
   For each bit, try the three valid `(p_bit, q_bit)` choices and update the relation states for `p` vs `x` and `q` vs `y`.

4. Once a prefix is already smaller or larger, the contribution of each later bit to the absolute difference is fixed.
   This makes the DP transition straightforward.

5. The full state space is tiny:
   - 31 bit positions
   - 3 states for `p` relative to `x`
   - 3 states for `q` relative to `y`

   So the solution runs in constant time per test case.

6. After computing the minimum value, store parent transitions to reconstruct one optimal pair `(p, q)`.


