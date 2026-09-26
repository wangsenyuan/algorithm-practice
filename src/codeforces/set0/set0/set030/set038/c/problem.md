# C. Blinds

[Problem link](https://codeforces.com/problemset/problem/38/C)

**Tags:** brute force, \*1400

You have `n` wooden planks of lengths `a_1, ..., a_n`. You want to cut them into equal-height strips of some height `h` (`h >= l`) and hang those strips as blinds for a rectangular window. Each plank of length `a_i` yields `floor(a_i / h)` strips of height `h`. The window height is `h` and its width is the total number of strips; the area is therefore `h` times that total. Find the maximum possible window area, or `0` if no valid `h` exists.

## Constraints

- Single test (no leading `t`)
- First line: `n l` (`1 <= n, l <= 100`)
- Second line: `n` integers `a_i` (`1 <= a_i <= 100`)

## Samples

### Sample 1

Input:

```
4 2
1 2 3 4
```

Output:

```
8
```

### Sample 2

Input:

```
5 3
5 5 7 3 1
```

Output:

```
15
```

### Sample 3

Input:

```
2 3
1 2
```

Output:

```
0
```
