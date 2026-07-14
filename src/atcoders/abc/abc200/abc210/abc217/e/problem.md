# E - Sorting Queries

[Problem link](https://atcoder.jp/contests/abc217/tasks/abc217_e)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 500 points

## Problem Statement

We have an empty sequence `A`. You will be given `Q` queries, which should be
processed in the order they are given. Each query is of one of the three kinds
below:

- `1 x` : Append `x` to the end of `A`.
- `2` : Print the element at the beginning of `A`. Then, delete that element.
  It is guaranteed that `A` will not be empty when this query is given.
- `3` : Sort `A` in ascending order.

## Constraints

- `1 <= Q <= 2 * 10^5`
- `0 <= x <= 10^9`
- `A` will not be empty when a query `2` is given.
- All values in input are integers.

## Input

```
Q
query_1
query_2
...
query_Q
```

Each query is in one of the three formats:

```
1 x
2
3
```

## Output

Print `q` lines, where `q` is the number of queries with type `2`.
The `j`-th line should contain the response for the `j`-th such query.

## Samples

### Sample 1

Input:

```
8
1 4
1 3
1 2
1 1
3
2
1 0
2
```

Output:

```
1
2
```

### Sample 2

Input:

```
9
1 5
1 5
1 3
2
3
2
1 6
3
2
```

Output:

```
5
3
5
```
