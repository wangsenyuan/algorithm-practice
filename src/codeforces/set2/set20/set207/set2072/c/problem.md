# Problem

Akito still has nowhere to live, so he got a job at a bank as a key creator for storages.

In this world, the key for storage code `(n, x)` is an array `a` of length `n` such that:

`a1 | a2 | ... | an = x`,

where `|` is bitwise OR.

Among all such arrays, `MEX({a1, a2, ..., an})` must be maximized.

For each given `n` and `x`, construct any valid key array.

`MEX(S)` is the minimum non-negative integer `z` such that `z` is not in `S`, and all integers `0 <= y < z` are in `S`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains one line with two integers `n` and `x` (`1 <= n <= 2 * 10^5`, `0 <= x < 2^30`) — the length of the array and the required bitwise OR.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `n` integers `ai` (`0 <= ai < 2^30`) — elements of a key array satisfying all conditions.

If there are multiple valid arrays, output any.

## Example

**Input**

```text
9
1 69
7 7
5 7
7 3
8 7
3 52
9 11
6 15
2 3
```

**Output**

```text
69
6 0 3 4 1 2 5
4 1 3 0 2
0 1 2 3 2 1 0
7 0 6 1 5 2 4 3
0 52 0
0 1 8 3 0 9 11 2 10
4 0 3 8 1 2
0 3
```
