# D. Christmas Tree Decoration

[Problem link](https://codeforces.com/problemset/problem/2182/D)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

A group of `n` people decided to decorate the Christmas tree. They have `(n+1)`
boxes of decorations, numbered from `0` to `n`. Initially, the `i`-th box contains
`a_i` decorations.

A permutation `p` of size `n` is **fair** if it is possible to hang all decorations
using this process:

- person `p_1` takes a decoration from box `0` or box `p_1`, and hangs it;
- person `p_2` takes a decoration from box `0` or box `p_2`, and hangs it;
- and so on;
- after person `p_n`, person `p_1` follows, repeating until all decorations are
  hung.

During this process, there must never be a step where person `i` needs a decoration
but both box `0` and box `i` are empty. If people can choose which box to take from
at each step so this never happens, the permutation is fair.

Count the number of fair permutations modulo `998244353`.

## Input

The first line contains an integer `t` (`1 <= t <= 5000`) — the number of test cases.

Each test case:

- One line: `n` (`1 <= n <= 50`).
- One line: `n+1` integers `a_0, a_1, …, a_n` (`0 <= a_i <= 10^6`).

## Output

For each test case, print one integer — the number of fair permutations modulo
`998244353`.

## Example

### Input

```text
4
3
1 2 1 0
3
1 0 2 0
1
2 5
4
6 1 4 2 1
```

### Output

```text
2
0
1
12
```

## Note

In the first example, the fair permutations are `[1, 2, 3]` and `[1, 3, 2]`.

For permutation `[1, 3, 2]`:

- person `p_1 = 1` hangs a decoration from box `1`;
- person `p_2 = 3` hangs a decoration from box `0`;
- person `p_3 = 2` hangs a decoration from box `2`;
- person `p_1 = 1` hangs a decoration from box `1`.

If person `p_1 = 1` takes from box `0` on the first step, person `p_2 = 3` cannot
perform the next step (boxes `0` and `3` are both empty). Since this can be avoided,
the permutation is fair.

## ideas
1. 好奇怪的题目. 有多少种排列,可以得到一个fair的分配.
2. 排序a[1:n]
3. 这个完全没想法, 怎么才1600~
4. 考虑n = 2, [1, 2, 0]  (这个答案 = 1, [2, 1])
5. n = 2, [1, 2, 1] 答案是2, [2, 1]和[1, 2]都可以