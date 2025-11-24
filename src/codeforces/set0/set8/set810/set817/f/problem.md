# Problem F

You are given a set of integer numbers, initially it is empty. You should perform n queries.

There are three different types of queries:

1. `1 l r` — Add all missing numbers from the interval [l, r]
2. `2 l r` — Remove all present numbers from the interval [l, r]
3. `3 l r` — Invert the interval [l, r] — add all missing and remove all present numbers from the interval [l, r]

After each query you should output MEX of the set — the smallest positive (MEX ≥ 1) integer number which is not presented in the set.

## Input

The first line contains one integer number n (1 ≤ n ≤ 10^5).

Next n lines contain three integer numbers t, l, r (1 ≤ t ≤ 3, 1 ≤ l ≤ r ≤ 10^18) — type of the query, left and right bounds.

## Output

Print MEX of the set after each query.

## Examples

### Example 1

**Input:**

```text
3
1 3 4
3 1 6
2 1 3
```

**Output:**

```text
1
3
1
```

### Example 2

**Input:**

```text
4
1 1 3
3 5 6
2 4 4
3 1 6
```

**Output:**

```text
4
4
4
1
```

## Note

Here are contents of the set after each query in the first example:

- {3, 4} — the interval [3, 4] is added
- {1, 2, 5, 6} — numbers {3, 4} from the interval [1, 6] got deleted and all the others are added
- {5, 6} — numbers {1, 2} got deleted



### ideas
1. segment tree支持三种操作，区间内全部设置为0，或者全部设置为1，或者revert
2. 然后找到从左边开始为0的第一个位置