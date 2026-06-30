# A. Yaroslav and Sequence

[Problem link](https://codeforces.com/problemset/problem/301/A)

**Contest:** [Codeforces Round 182 (Div. 1)](https://codeforces.com/contest/301)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Yaroslav has an array consisting of `2 * n - 1` integers. In one operation, he can choose exactly `n` array
elements and multiply each of them by `-1`.

He may perform any number of such operations. Find the maximum possible sum of all array elements.

## Input

The first line contains an integer `n` (`2 <= n <= 100`).

The second line contains `2 * n - 1` integers — the array elements. Each element's absolute value does not
exceed `1000`.

## Output

Print one integer — the maximum sum Yaroslav can obtain.

## Example

### Input

```text
2
50 50 50
```

### Output

```text
150
```

### Input

```text
2
-1 -100 -1
```

### Output

```text
100
```

### Note

- In the first sample, no operation is needed; the sum is `150`.
- In the second sample, flip the sign of the first two elements to get sum `100`.

## Solution

Only the signs matter. Each operation flips exactly `n` elements, so applying operations is equivalent to choosing
which positions are flipped an odd number of times in the end.

If `n` is odd, flipping `n` elements changes an odd number of signs. Because the array length is `2 * n - 1`,
we can combine operations to make any single position change sign, and therefore any final sign pattern is
reachable. The maximum sum is simply:

```text
sum(abs(a[i]))
```

If `n` is even, every operation flips an even number of signs. Therefore the parity of the number of negative
values never changes. The best target is still to make every value positive, unless the initial number of
negative values is odd. In that case, one value must remain negative, so we lose the smallest possible amount by
leaving the element with minimum absolute value negative:

```text
sum(abs(a[i])) - 2 * min(abs(a[i]))
```

So the algorithm scans the array once, accumulating the absolute-value sum, counting negative values, and tracking
the minimum absolute value.

## Complexity

The solution runs in `O(n)` time and uses `O(1)` extra space.
