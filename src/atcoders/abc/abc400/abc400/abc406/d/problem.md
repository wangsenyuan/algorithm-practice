# D - Garbage Removal

[Problem link](https://atcoder.jp/contests/abc406/tasks/abc406_d)

**Contest:** [Panasonic Programming Contest 2025 (AtCoder Beginner Contest 406)](https://atcoder.jp/contests/abc406)

time limit: 2.5 sec

memory limit: 1024 MiB

score: 400 points

There is an `H` by `W` grid. The `N` pieces of trash are at distinct cells `(X_i, Y_i)`.

Process `Q` queries in order. Each query is one of:

- Type 1: `1 x` — report how many trash pieces are in row `x`, then remove all trash in that row
- Type 2: `2 y` — report how many trash pieces are in column `y`, then remove all trash in that column

## Constraints

- 1 <= H, W, N <= 2 * 10^5
- 1 <= X_i <= H, 1 <= Y_i <= W
- (X_i, Y_i) != (X_j, Y_j) for i != j
- 1 <= Q <= 2 * 10^5
- For type 1 queries, 1 <= x <= H
- For type 2 queries, 1 <= y <= W
- All input values are integers

## Input

```text
H W N
X_1 Y_1
...
X_N Y_N
Q
query_1
...
query_Q
```

Each query is either `1 x` or `2 y`.

## Output

Print `Q` lines. The i-th line is the answer to the i-th query.

## Sample Input 1

```text
3 4 5
1 2
1 3
3 4
3 1
2 2
5
1 1
1 2
2 2
2 4
1 2
```

## Sample Output 1

```text
2
1
0
1
0
```

## Sample Input 2

```text
1 2 1
1 2
7
2 1
2 1
2 1
2 1
2 1
2 1
2 1
```

## Sample Output 2

```text
0
0
0
0
0
0
0
```

## Sample Input 3

```text
4 4 16
1 1
1 2
1 3
1 4
2 1
2 2
2 3
2 4
3 1
3 2
3 3
3 4
4 1
4 2
4 3
4 4
7
2 1
1 1
2 2
1 2
2 3
1 3
2 4
```

## Sample Output 3

```text
4
3
3
2
2
1
1
```

## Solution

Each piece of trash sits in exactly one row and one column. Instead of scanning the whole grid on
every query, keep two index sets:

- `rowTodo[r]`: trash indices still present in row `r`
- `colTodo[c]`: trash indices still present in column `c`

When trash `i` is placed at `(r, c)`, insert `i` into both sets.

### Processing queries

For a type-1 query on row `r`:

1. Answer with `len(rowTodo[r])`
2. For every trash index `i` still in that row, remove `i` from `colTodo[c]` where `(r, c)` is its cell
3. Clear `rowTodo[r]`

For a type-2 query on column `c`, do the symmetric operation using `colTodo[c]` and `rowTodo[r]`.

Example after the first query in sample 1:

- Row 1 contains trash at `(1, 2)` and `(1, 3)`, so the answer is `2`
- Those two indices are deleted from their column sets, and row 1 becomes empty
- Remaining trash stays indexed only in rows/columns where it has not been removed yet

### Correctness

A piece of trash is alive if and only if neither its row nor its column has been cleared yet. The
two sets always describe the same alive set, just grouped in two ways:

- `rowTodo[r]` lists alive trash in row `r`
- `colTodo[c]` lists alive trash in column `c`

When a row is cleared, every alive trash in that row is removed from both structures, so it can
never be counted again. The same holds for column clears. Therefore each query counts exactly the
alive trash in the requested row or column, then removes all of it.

### Complexity

Each trash index is inserted once and deleted at most once from each of the two sets.

Time: `O(N + Q + number of deletions) = O(N + Q)`  
Memory: `O(N + H + W)`
