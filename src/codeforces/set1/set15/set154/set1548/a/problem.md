# A. Web of Lies

[Problem link](https://codeforces.com/problemset/problem/1548/A)

**Contest:** [Codeforces Round #736 (Div. 1)](https://codeforces.com/contest/1548)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

There are `n` nobles, numbered from `1` to `n`. Noble `i` has a power of `i`. There are also `m` "friendships". A friendship between nobles `a` and `b` is always mutual.

A noble is defined to be **vulnerable** if both of the following conditions are satisfied:

- the noble has at least one friend, and
- all of that noble's friends have a higher power.

You will have to process the following three types of queries.

1. Add a friendship between nobles `u` and `v`.
2. Remove a friendship between nobles `u` and `v`.
3. Calculate the answer to the following process.

The process: all vulnerable nobles are simultaneously killed, and all their friendships end. Then, it is possible that new nobles become vulnerable. The process repeats itself until no nobles are vulnerable. It can be proven that the process will end in finite time. After the process is complete, you need to calculate the number of remaining nobles.

Note that the results of the process are not carried over between queries, that is, every process starts with all nobles being alive!

## Input

The first line contains the integers `n` and `m` (`1 <= n <= 2 * 10^5`, `0 <= m <= 2 * 10^5`) — the number of nobles and number of original friendships respectively.

The next `m` lines each contain the integers `u` and `v` (`1 <= u, v <= n`, `u != v`), describing a friendship. No friendship is listed twice.

The next line contains the integer `q` (`1 <= q <= 2 * 10^5`) — the number of queries.

The next `q` lines contain the queries themselves, each query has one of the following three formats:

- `1 u v` (`1 <= u, v <= n`, `u != v`) — add a friendship between `u` and `v`. It is guaranteed that `u` and `v` are not friends at this moment.
- `2 u v` (`1 <= u, v <= n`, `u != v`) — remove a friendship between `u` and `v`. It is guaranteed that `u` and `v` are friends at this moment.
- `3` — print the answer to the process described in the statement.

## Output

For each type `3` query print one integer to a new line. It is guaranteed that there will be at least one type `3` query.

## Examples

### Input

```text
4 3
2 1
1 3
3 4
4
3
1 2 3
2 3 1
3
```

### Output

```text
2
1
```

### Input

```text
4 3
2 3
3 4
4 1
1
3
```

### Output

```text
1
```

### Note

In the first example, the first type `3` query kills noble `1` in round 1 (weaker than friends `2` and `3`), then noble `3` in round 2 (weaker than remaining friend `4`). Two nobles remain.

In the second type `3` query, the only surviving noble is `4`.

In the second example, two nobles die in the first round and one in the second, leaving a single survivor.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
