# C - Brackets Stack Query

[Problem link](https://atcoder.jp/contests/abc428/tasks/abc428_c)

**Contest:** [AtCoder Beginner Contest 428](https://atcoder.jp/contests/abc428)

time limit: 3 sec

memory limit: 1024 MiB

score: 300 points

A string `T` is a **good bracket sequence** if it can be reduced to the empty string by repeatedly
removing a contiguous substring `()` from `T`.

For example, `()`, `(()())`, and the empty string are good; `)()(` and `)))` are not.

String `S` starts empty. Process `Q` queries in order. After each query, determine whether `S` is a
good bracket sequence.

Query types:

- `1 c` — append `c` to the end of `S`, where `c` is `(` or `)`
- `2` — remove the last character of `S` (guaranteed `S` is not empty)

## Constraints

- `1 <= Q <= 8 * 10^5`
- `c` in type-1 queries is `(` or `)`
- When a type-2 query is given, `S` is not empty
- `Q` is an integer

## Input

```text
Q
query_1
query_2
...
query_Q
```

Each query is one of:

```text
1 c
```

```text
2
```

## Output

Print `Q` lines. The `i`-th line should be `Yes` if `S` is a good bracket sequence immediately after
the `i`-th query, and `No` otherwise.

## Sample Input 1

```text
8
1 (
2
1 (
1 )
2
1 (
1 )
1 )
```

## Sample Output 1

```text
No
Yes
No
Yes
No
No
No
Yes
```

After each query, `S` becomes:

| Query | `S` | Good? |
|------:|-----|-------|
| 1 | `(` | No |
| 2 | `` | Yes |
| 3 | `(` | No |
| 4 | `()` | Yes |
| 5 | `(` | No |
| 6 | `((` | No |
| 7 | `(()` | No |
| 8 | `(())` | Yes |

### ideas
1. 
