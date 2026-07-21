# A. Gellyfish and Flaming Peony

[Problem link](https://codeforces.com/problemset/problem/2115/A)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Gellyfish is given an array of `n` positive integers `a_1, a_2, ..., a_n`.

She needs to do the following two-step operation until all elements of `a` are equal:

1. Select two indexes `i`, `j` satisfying `1 <= i, j <= n` and `i != j`.
2. Replace `a_i` with `gcd(a_i, a_j)`.

Find the minimum number of operations to achieve her goal.

It can be proven that Gellyfish can always achieve her goal.

## Constraints

- `1 <= t <= 5000`
- `1 <= n <= 5000`
- `1 <= a_i <= 5000`
- The sum of `n` over all test cases does not exceed `5000`

## Input

The first line contains the number of test cases `t`.

For each test case:

- The first line contains a single integer `n` — the length of the array.
- The second line contains `n` integers `a_1, a_2, ..., a_n` — the elements of the array.

## Output

For each test case, output a single integer — the minimum number of operations.

## Example

```text
Input
3
3
12 20 30
6
1 9 1 9 8 1
3
6 14 15

Output
4
3
3
```

In the first test case, one minimal sequence is:

1. Choose `i = 3`, `j = 2`: `a` becomes `[12, 20, 10]`.
2. Choose `i = 1`, `j = 3`: `a` becomes `[2, 20, 10]`.
3. Choose `i = 2`, `j = 1`: `a` becomes `[2, 2, 10]`.
4. Choose `i = 3`, `j = 1`: `a` becomes `[2, 2, 2]`.


### ideas
1. 最后的数 x = gcd(a[1], ... a[n])
2. 先通过最快的次数得到x, 然后, 所有不是x的,通过一次就可以得到x
3. 但是下一步咋搞, 就不清楚了;
4. 所有的数都除以x, 那么就是要得到gcd = 1的情况
5. 
