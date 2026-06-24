# E. Blackslex and Girls

[Problem link](https://codeforces.com/problemset/problem/2179/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

After failing to pick up a girl using De Bruijn sequence of fixed-length bitstrings, Blackslex
has turned his attention towards politics.

Due to his high charisma, he is now in charge of drawing borders for the `n` voting districts
of his country. In Blackslex's country, there are `x` voters for party A and `y` voters for
party B. Using his amazing drawing skills, he can allocate voters from any party into any
district of his choice.

His history with bitstrings has led him to wonder if he can allocate voters such that the
winner of each district follows a certain bitstring pattern. To avoid suspicion, he must also
allocate at least `p_i` voters into each district. Tell him if it is possible!

Formally, you are given a binary string `s` of length `n`, an array `p` of length `n`, and two
integers `x` and `y`.

You want to determine whether there exist two arrays of nonnegative integers `a` and `b` of
length `n` that satisfy the following conditions:

- `sum(a) = x` and `sum(b) = y`
- for every `1 <= i <= n`, `a_i + b_i >= p_i`
- for every `1 <= i <= n`:
  - if `s_i = 0` then `a_i > b_i`
  - if `s_i = 1` then `b_i > a_i`

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t`
(`1 <= t <= 10^4`).

The first line of each test case contains three integers `n`, `x`, and `y`
(`1 <= n <= 2 * 10^5`, `1 <= x, y <= 10^9`).

The second line contains a binary string `s` of length `n`.

The third line contains `n` integers `p_1, p_2, ..., p_n` (`1 <= p_i <= 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print (case-insensitive) `YES` if there exist arrays `a`, `b` satisfying
all conditions, or `NO` otherwise.

## Example

### Input

```text
6
3 5 5
010
2 4 3
2 4 3
00
3 4
2 4 2
00
3 3
3 4 2
011
1 1 2
4 2 2
1111
2 2 5 2
1 2 6
0
512
```

### Output

```text
YES
NO
YES
NO
NO
NO
```

## Note

In the first test case, one of the possible distributions of voters is: `a = [2, 0, 3]` and
`b = [0, 4, 1]`.

In the third test case, one of the possible distributions of voters is: `a = [2, 2]` and
`b = [1, 1]`.

For the other test cases, it can be shown that there are no distributions of voters that
satisfy the conditions.


### ideas
1. 如果s[i] = 0, a[i] > b[i], else a[i] < b[i], 且a[i] + b[i] >= p[i]
2. 假设最后的分配方案a[i], b[i], 
3. 假设有w个s[i] = 0, 那么有 n - w 个 s[i] = 1
4. 那么这w个 a[i] > b[i], 这些p[i], 最小的diff = 