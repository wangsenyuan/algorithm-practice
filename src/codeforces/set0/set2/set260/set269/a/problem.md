# Problem

Emuskald is a well-known illusionist. One of his trademark tricks involves a set of magical boxes. The essence of the trick is packing boxes inside other boxes.

From the top view, each magical box looks like a square whose side length is `2^k` (`k` is an integer, `k >= 0`).

A magical box `v` can be put inside a magical box `u` if the side length of `v` is strictly smaller than the side length of `u`.

In particular, four boxes of side length `2^(k-1)` can be put into one box of side length `2^k`.

Emuskald is about to go on tour and needs to pack all his magical boxes into one larger magical box. Since boxes are expensive, help him find the smallest possible outer box.

## Input

- First line: integer `n` (`1 <= n <= 10^5`) — the number of different box sizes.
- Each of the next `n` lines contains two integers `k_i` and `a_i`:
  - `0 <= k_i <= 10^9`
  - `1 <= a_i <= 10^9`
  - meaning Emuskald has `a_i` boxes of side length `2^(k_i)`.

It is guaranteed that all `k_i` are distinct.

## Output

Print a single integer `p` such that the smallest magical box that can contain all boxes has side length `2^p`.

## Examples

### Example 1

Input

```text
2
0 3
1 5
```

Output

```text
3
```

### Example 2

Input

```text
1
0 4
```

Output

```text
1
```

### Example 3

Input

```text
2
1 10
2 2
```

Output

```text
3
```

## Note

In the first sample, if we have `3` boxes of side length `2` and `5` boxes of side length `1`, we can pack all of them into one box of side length `8` (`2^3`).

In the second sample, four boxes of side length `1` can be packed into one box of side length `2`.
