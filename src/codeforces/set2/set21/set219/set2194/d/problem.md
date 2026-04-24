# Problem

Given a table of size `n × m`, where each cell contains either `0` or `1`, divide it into two parts by a cut that goes from the top-left corner to the bottom-right corner.

The cut can move only **right** or **down**.

Let:

- `a` = number of ones in one part,
- `b` = number of ones in the other part.

Your goal is to maximize `a * b`.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^4`) — number of test cases.

Each test case:

- First line: two integers `n` and `m` (`1 ≤ n, m ≤ 3 * 10^5`, `2 ≤ n * m ≤ 3 * 10^5`) — table dimensions.
- Next `n` lines: each contains `m` integers `ai,j` (`0 ≤ ai,j ≤ 1`) — table values.

It is guaranteed that the sum of `n * m` over all test cases does not exceed `3 * 10^5`.

## Output

For each test case:

1. Print one integer — the maximum product value.
2. Print one string describing the cut path:
   - consisting of exactly `n - 1` characters `'D'` and `m - 1` characters `'R'`,
   - where `'D'` means move cut down, and `'R'` means move cut right.

If there are multiple optimal cuts, print any.

## Example

**Input**

```text
3
5 5
1 0 1 1 0
0 1 0 1 1
1 0 1 0 0
0 1 0 1 0
0 0 0 0 1
5 4
0 0 1 0
0 1 1 1
1 0 0 1
0 1 0 1
0 0 1 0
3 2
1 0
0 1
1 1
```

**Output**

```text
30
RDRDRDRDDR
20
DRRDRDDDR
4
DRDRD
```

## Note

The statement figures show valid optimal cuts for the first and second test cases.


### ideas
1. 这个咋这么难的～