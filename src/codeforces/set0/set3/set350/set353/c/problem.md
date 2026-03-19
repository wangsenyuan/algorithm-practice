# Problem

Valera has an array `a` consisting of `n` integers `a0, a1, ..., a(n-1)` and a function `f(x)` defined for integers `x` in the range `0 .. 2^n - 1`.

Let `bit(i)` be `1` if the binary representation of `x` has a `1` at position `i` (0-indexed from the least significant bit), otherwise `0`.

Then:

\[
f(x) = \sum_{i=0}^{n-1} a_i \cdot bit(i)
\]

For example, if `n = 4` and `x = 11` (`11 = 2^0 + 2^1 + 2^3`), then `f(x) = a0 + a1 + a3`.

Help Valera find the maximum value of `f(x)` among all `x` satisfying `0 ≤ x ≤ m`.

## Input

- First line: integer `n` (`1 ≤ n ≤ 10^5`) — the number of array elements.
- Second line: `n` space-separated integers `a0, a1, ..., a(n-1)` (`0 ≤ ai ≤ 10^4`) — the elements of array `a`.
- Third line: a binary string `s` of length `n` (characters `0` and `1`) — the binary representation of number `m` (most significant bit first).

## Output

Print a single integer — the maximum value of `f(x)` over all `x` such that `0 ≤ x ≤ m`.

## Examples

### Example 1

Input

```text
2
3 8
10
```

Output

```text
3
```

### Example 2

Input

```text
5
17 0 10 2 1
11010
```

Output

```text
27
```

## Note

In the first test case, `m = 2` (binary `10`), `f(0) = 0`, `f(1) = a0 = 3`.

In the second sample, `m = 2^4 + 2^3 + 2^1 = 26` (binary `11010`), the maximum value equals `f(5) = a0 + a2 = 17 + 10 = 27`.
