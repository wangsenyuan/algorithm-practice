# Problem

Polycarpus has a family tree of `n` people numbered `1` to `n`. Each person has at most one direct parent (ancestor). Each person has a **name**; names need not be unique.

Definitions:

- Person `a` is a **1-ancestor** of person `b` if `a` is the direct parent of `b`.
- For `k > 1`, person `a` is a **k-ancestor** of person `b` if `b` has a parent and `a` is a **(k−1)-ancestor** of that parent.
- Person `a` is a **k-son** of person `b` if `b` is a **k-ancestor** of `a`.

The graph has no cycles: no one is their own ancestor.

Polycarpus wrote `m` pairs `(v_i, k_i)`. For each pair, find how many **distinct names** appear among all **k_i-sons** of person `v_i`.

## Input

- First line: integer `n` (`1 <= n <= 10^5`) — number of people.
- Next `n` lines: line `i` contains string `s_i` and integer `r_i` (`0 <= r_i <= n`):
  - `s_i` — name of person `i`
  - `r_i` — index of the direct parent of `i`, or `0` if person `i` has no parent
- Next line: integer `m` (`1 <= m <= 10^5`) — number of queries.
- Next `m` lines: two integers `v_i`, `k_i` (`1 <= v_i, k_i <= n`).

It is guaranteed that the family relations form a forest (no cycles). Each name is a non-empty string of at most `20` lowercase English letters.

## Output

Print `m` integers separated by whitespace — answers in the order of the queries.

## Examples

### Example 1

Input

```text
6
pasha 0
gerald 1
gerald 1
valera 2
igor 3
olesya 1
5
1 1
1 2
1 3
3 1
6 1
```

Output

```text
2
2
0
1
0
```

### Example 2

Input

```text
6
valera 0
valera 1
valera 1
gerald 0
valera 4
kolya 4
7
1 1
1 2
2 1
2 2
4 1
5 1
6 1
```

Output

```text
1
0
0
0
2
0
0
```
