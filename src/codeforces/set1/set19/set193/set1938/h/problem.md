# H. Pho Restaurant

[Problem link](https://codeforces.com/problemset/problem/1938/H)

**Contest:** [2024 ICPC Asia Pacific Championship — Online Mirror](https://codeforces.com/contest/1938)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

input: standard input

output: standard output

As you may know, pho is one of the most common dishes in Hanoi. You own a phở bò (beef pho)
restaurant in Vietnam with `n` tables, numbered `1` to `n`. ICPC contestants are in your restaurant;
each contestant is initially seated at one of the tables, and at least one contestant sits at each table.

Each contestant orders one of two kinds of pho:

- **phở tái** (rare beef pho)
- **phở chín** (well-done beef pho)

For each table, at least one of the following must hold after moving contestants:

- Every contestant at that table orders phở tái.
- Every contestant at that table orders phở chín.

You may move zero or more contestants to a different existing table (no new tables). There is no
limit on how many contestants can sit at the same table. After moving, each table is either empty or
all seated contestants want the same dish.

Find the **minimum number of contestants** you need to move.

The `i`-th input string describes table `i`: each character is `0` if that contestant wants phở tái,
or `1` if they want phở chín. The `j`-th character is the `j`-th contestant initially at that table
(from left to right in the string).

## Input

The first line contains one integer `n` (`2 <= n <= 100000`).

Each of the next `n` lines contains a binary string `S_i` (`1 <= |S_i| <= 200000`). The sum of
`|S_i|` over all `i` does not exceed `500000`.

## Output

Print one integer — the minimum number of contestants to move.

## Examples

### Input 1

```text
4
11101101
00
10001
10
```

### Output 1

```text
5
```

### Input 2

```text
2
101010
010101
```

### Output 2

```text
6
```

### Input 3

```text
5
0000
11
0
00000000
1
```

### Output 3

```text
0
```

## Note

In example 1, one optimal plan moves 5 contestants so that all remaining contestants at table 1 order
phở chín and every other non-empty table orders only phở tái. Fewer than 5 moves cannot work.

In example 3, every table is already homogeneous (all `0` or all `1`), so the answer is `0`.

## Thoughts Behind

A table that already contains only `0` or only `1` is already valid, so it should normally stay as it
is. The only difficult tables are mixed tables. For one mixed table, the final table can keep either:

- all `0` people, moving all `1` people away;
- all `1` people, moving all `0` people away.

So every mixed table is a binary choice. The first tempting idea is to make almost all mixed tables
choose the same final digit, with maybe one table choosing the other digit. That is not enough: if
several mixed tables have many `0`s and several others have many `1`s, the optimum may keep both
digits on multiple mixed tables.

Example:

```text
4
00001
00001
01111
01111
```

The best answer is `4`: keep `0` on the first two tables and keep `1` on the last two tables. A
solution that only allows one table to choose the minority side would overcount.

The remaining constraint is destination availability. If some `0` people are moved, there must be at
least one final `0` table to receive them. This destination may be:

- a mixed table whose final choice is `0`;
- an originally all-`0` table.

The same applies to moved `1` people. Therefore, after deciding final choices for all mixed tables, we
only need to know which digits appear among those final mixed tables. This is a 4-state mask:

- bit `0`: at least one mixed table ends as all `0`;
- bit `1`: at least one mixed table ends as all `1`.

For each mixed table with `cnt0` zeros and `cnt1` ones:

- choosing final `0` costs `cnt1` moves and sets bit `0`;
- choosing final `1` costs `cnt0` moves and sets bit `1`.

After processing all mixed tables, inspect every reachable mask. If all mixed tables end as `1`, then
the `0` people moved out of those tables need an originally all-`0` destination. If no such table
exists, we can convert the smallest originally all-`1` table into a `0` destination by moving all its
`1` people away; those moved `1` people can go to a mixed table that already ends as `1`.
Symmetrically, if all mixed tables end as `0`, but there is no originally all-`1` table, convert the
smallest originally all-`0` table.

If the mask contains both digits, both moved groups already have destinations and no extra cost is
needed. If there are no mixed tables, the initial arrangement is already valid.

### Correctness

In an optimal arrangement, every mixed original table can be assigned one final digit: keeping either
all original `0`s or all original `1`s is never worse than emptying the table. Choosing final `0`
forces all its original `1`s to move, and choosing final `1` forces all its original `0`s to move, so
the DP cost is exactly the number of moved people from mixed tables.

Originally homogeneous tables never need to change unless a moved group has no destination of its
digit. In that case changing any one opposite homogeneous table is sufficient, and the cheapest such
change is moving all people from the smallest opposite homogeneous table. The final mask check adds
exactly this necessary and sufficient extra cost.

Thus the DP enumerates all relevant final digit assignments for mixed tables, and the final adjustment
handles the only remaining destination constraint. Taking the minimum over all masks gives the minimum
number of moved contestants.

### Complexity

Let `L` be the total length of all strings. Counting table contents takes `O(L)`, and the DP has only
4 states per mixed table, so the total time is `O(L + n)`. The extra memory is `O(n)` for the counted
tables and `O(1)` for the DP states.
