# F - Back and Forth Filling

[Problem link](https://atcoder.jp/contests/abc430/tasks/abc430_f)

**Contest:** [AtCoder Beginner Contest 430](https://atcoder.jp/contests/abc430)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

You are given an integer `N` and a string `S` of length `N - 1` consisting of `L` and `R`.

Consider writing integers into `N` cells arranged in a row so that the following conditions are satisfied:

- Every cell has one integer written on it.
- Every integer `1, 2, ..., N` appears in exactly one cell.
- When the `i`-th character of `S` is `L`, `i + 1` is written to the left of `i`.
- When the `i`-th character of `S` is `R`, `i + 1` is written to the right of `i`.

Let `C_i` be the number of integers that can be written in the `i`-th cell from the left. Find
`C_1, C_2, ..., C_N`.

You are given `T` test cases; find the answer for each of them.

## Constraints

- `1 <= T <= 20000`
- `2 <= N <= 3 * 10^5`
- `S` is a string of length `N - 1` consisting of `L` and `R`.
- For a single input, the sum of `N` does not exceed `3 * 10^5`.

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case is given in the following format:

```text
N
S
```

## Output

Print `T` lines. The `i`-th line should contain the answer for the `i`-th test case in the
following format:

```text
C_1 C_2 ... C_N
```

## Sample Input 1

```text
5
5
RLLR
3
RL
2
L
3
RR
20
RLLLLLLLLRLRRLLLRLR
```

## Sample Output 1

```text
2 4 3 4 2
2 2 1
1 1
1 1 1
5 9 13 14 15 17 18 19 19 20 20 19 19 18 17 16 14 12 9 5
```

### Note

- In the first test case, there are 11 valid fillings. For example, `(1, 4, 3, 2, 5)` is valid.
  Then `C = (2, 4, 3, 4, 2)`.
- In the second test case, the valid fillings are `(1, 3, 2)` and `(3, 1, 2)`, giving
  `C = (2, 2, 1)`.
- In the third test case, the only valid filling is `(2, 1)`, giving `C = (1, 1)`.
- In the fourth test case, the only valid filling is `(1, 2, 3)`, giving `C = (1, 1, 1)`.

## Solution

The key point is that `S[i]` describes the relative order of the **values** `i` and `i + 1`, not the
cell positions `i` and `i + 1`.

For example, if `S[4] = R`, it means value `5` must be somewhere to the right of value `4`:

```text
pos[5] > pos[4]
```

It does not mean value `5` must be written in the 5-th cell.

### Build a Directed Graph

Create one vertex for each value `1, 2, ..., N`.

For every adjacent value pair `(i, i + 1)`:

- if `S[i] = R`, then `i + 1` must be to the right of `i`, so add edge `i -> i + 1`;
- if `S[i] = L`, then `i + 1` must be to the left of `i`, so add edge `i + 1 -> i`.

An edge `u -> v` means:

```text
u must be placed before v
```

Since the original graph over values is just a path, this directed graph has no directed cycle.

### Possible Positions of One Value

For a value `x`, suppose:

- `before[x]` is the number of values that must be placed before `x`, including `x` itself;
- `after[x]` is the number of values that must be placed after `x`, including `x` itself.

Then `x` can be placed only after all `before[x] - 1` required predecessors, so the earliest
zero-based cell is:

```text
L = before[x] - 1
```

It also must leave enough room on the right for all `after[x] - 1` required successors, so the latest
zero-based cell is:

```text
R = N - after[x]
```

Therefore value `x` can appear in every cell in the interval:

```text
[before[x] - 1, N - after[x]]
```

The answer `C_i` asks how many values can be written in cell `i`, so each value contributes `+1` to
all cells in its interval. Use a difference array to add all intervals in `O(N)`.

### Computing `before`

Run topological DP on the directed graph.

For every source vertex, set:

```text
before[source] = 1
```

When processing an edge `u -> v`, add:

```text
before[v] += before[u]
```

After all incoming edges of `v` are processed, add `1` for `v` itself.

Because the underlying graph is a path, the predecessor parts flowing into the same vertex are
disjoint, so summing these counts gives exactly the number of forced predecessors plus itself.

The implementation uses indegrees and a queue for this topological traversal.

### Computing `after`

The number of required successors is symmetric.

Reverse every edge and run the same topological DP again. The result on the reversed graph is:

```text
after[x]
```

the number of values that must be placed after `x`, including `x`.

### Example

For the first sample:

```text
N = 5
S = RLLR
```

The constraints are:

```text
1 -> 2
3 -> 2
4 -> 3
4 -> 5
```

For value `4`, nothing must be before it, but values `3`, `2`, and `5` must be after it in the
partial order. Thus:

```text
before[4] = 1
after[4] = 4
```

So value `4` can be placed from:

```text
before[4] - 1 = 0
```

to:

```text
N - after[4] = 1
```

That is, value `4` can appear in cell `1` or cell `2`, which matches the listed valid fillings.

### Correctness

Every edge `u -> v` means `u` must appear to the left of `v`. Thus all ancestors of `x` in the
directed graph must be placed before `x`, and all descendants of `x` must be placed after `x`.
This proves that `x` cannot be placed outside `[before[x] - 1, N - after[x]]`.

Conversely, if a cell is inside this interval, there is enough room on the left for all required
predecessors of `x` and enough room on the right for all required successors of `x`. Since the
constraints form an oriented path, those predecessor and successor sets do not conflict, so such a
placement is possible in some valid filling.

Therefore each value contributes exactly to the cells in its interval, and the difference-array
accumulation gives every `C_i`.

### Complexity

Each test case builds `O(N)` edges and runs two topological traversals. The final difference-array
scan is also `O(N)`.

Total complexity is `O(N)` per test case, and memory usage is `O(N)`.
