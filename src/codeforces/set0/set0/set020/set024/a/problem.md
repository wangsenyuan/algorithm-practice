# A. Ring road

[Problem link](https://codeforces.com/problemset/problem/24/A)

**Contest:** [Codeforces Beta Round #24](https://codeforces.com/contest/24)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Nowadays one-way traffic is introduced all over the world in order to improve driving safety and
reduce traffic jams. The government of Berland decided to keep up with new trends. Formerly all `n`
cities of Berland were connected by `n` two-way roads in a ring, i.e. each city was connected
directly to exactly two other cities, and from each city it was possible to get to any other city.
The government of Berland introduced one-way traffic on all `n` roads, but it soon became clear that
it is impossible to get from some of the cities to some others. Now for each road it is known in
which direction the traffic is directed, and the cost of redirecting the traffic. What is the
smallest amount of money the government should spend on redirecting roads so that from every city
you can get to any other?

## Input

The first line contains integer `n` (`3 ≤ n ≤ 100`) — the number of cities (and roads) in Berland.
Next `n` lines contain the description of roads. Each road is described by three integers `ai`,
`bi`, `ci` (`1 ≤ ai, bi ≤ n`, `ai ≠ bi`, `1 ≤ ci ≤ 100`) — the road is directed from city `ai` to
city `bi`, and redirecting the traffic costs `ci`.

## Output

Output a single integer — the smallest amount of money the government should spend on redirecting
roads so that from every city you can get to any other.

## Examples

### Input

```text
3
1 3 1
1 2 1
3 2 1
```

### Output

```text
1
```

### Input

```text
3
1 3 1
1 2 5
3 2 1
```

### Output

```text
2
```

### Input

```text
6
1 5 4
5 3 8
2 4 15
1 6 16
2 3 23
4 6 42
```

### Output

```text
39
```

### Input

```text
4
1 2 9
2 3 8
3 4 7
4 1 5
```

### Output

```text
0
```
