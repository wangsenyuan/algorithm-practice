You are given two integer arrays `a` and `b`, both of length `n`.

You can choose any subset of indices and swap the elements at those positions (i.e. perform `swap(a_i, b_i)` for each `i` in the subset). A subset of indices is **good** if, after the swaps, both arrays are sorted in non-decreasing order.

Your task is to count the number of good subsets. Since the answer can be large, print it modulo `998244353`.

## Input

The first line contains a single integer `t` (`1 ≤ t ≤ 500`) — the number of test cases.

The first line of each test case contains a single integer `n` (`1 ≤ n ≤ 100`).

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 ≤ a_i ≤ 1000`).

The third line contains `n` integers `b_1, b_2, …, b_n` (`1 ≤ b_i ≤ 1000`).

**Additional constraint:** there is at least one good subset.

## Output

For each test case, print a single integer — the number of good subsets modulo `998244353`.

## Example

**Input**

```
3
3
2 1 4
1 3 2
1
4
4
5
2 3 3 4 4
1 1 3 5 6
```

**Output**

```
2
2
8
```

## Note

In the first example, there are `2` good subsets: `{1, 3}` and `{2}`.

In the second example, there are `2` good subsets: `{1}` and `{}`.

In the third example, there are `8` good subsets: `{1, 2, 3, 4, 5}`, `{1, 2, 3}`, `{1, 2, 4, 5}`, `{1, 2}`, `{3, 4, 5}`, `{3}`, `{4, 5}`, and `{}`.
