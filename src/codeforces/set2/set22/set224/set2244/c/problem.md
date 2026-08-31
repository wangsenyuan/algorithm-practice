# C. Stepan and Permutation

[Problem link](https://codeforces.com/problemset/problem/2244/C)

**Contest:** [Codeforces Round 1109 (Div. 3)](https://codeforces.com/contest/2244)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Stepan found a permutation `p` of length `n`. Of course, he decided to sort it. To make the process more interesting, he chose two positive integers `x` and `y` (`x + y <= n`) and defined a rule for swapping elements.

In one move, Stepan can choose two indices `i` and `j` (`1 <= i, j <= n`) and swap the elements `p_i` and `p_j` if at least one of the following conditions holds:

- `|i - j| = x`
- `|i - j| = y`

Stepan wants to know whether it is possible to sort the permutation in ascending order using any number of such operations. Help him answer this question.

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of test cases.

The first line of each test case contains three integers `n`, `x`, and `y` (`1 <= x, y <= n <= 2 * 10^5`, `x + y <= n`) — the length of the array and the numbers chosen by Stepan.

The second line of each test case contains `n` integers `p_i` (`1 <= p_i <= n`) — the array `p`; it is guaranteed that `p` is a permutation.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output "YES" if it is possible to sort the permutation with the given `x` and `y`, and "NO" otherwise.

You may output each letter in any case (lowercase or uppercase). For example, the strings "yEs", "yes", "Yes", and "YES" will be accepted.

## Example

### Input

```text
4
5 2 3
5 4 3 2 1
6 2 4
2 1 4 3 6 5
4 2 2
1 2 3 4
5 2 3
1 2 3 5 4
```

### Output

```text
YES
NO
YES
YES
```

## Status

I/O and official samples are in place. `solve` is left as a TODO.
