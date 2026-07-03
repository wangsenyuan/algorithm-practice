# B. Restore Cube

[Problem link](https://codeforces.com/problemset/problem/464/B)

**Contest:** [Codeforces Round #268 (Div. 2)](https://codeforces.com/contest/464)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Peter wrote down the eight vertices of a cube in 3D space, one vertex per line. Each line contains three
integers — the coordinates of that vertex. The cube has non-zero side length; its vertices lie at integer
points (the cube need not be axis-aligned).

Nick then independently permuted the three numbers inside each line (any number of swaps per line).

Given the scrambled eight lines, restore valid cube vertex coordinates, or report that it is impossible.

## Input

Eight lines, each with three space-separated integers (`|value| <= 10^6`).

## Output

If restoration is possible, print `YES`, then eight lines of restored coordinates. Line `i` in the output
must be a permutation of line `i` in the input. The eight points must be vertices of a cube with
non-zero side length. Print any valid restoration.

Otherwise print `NO`.

## Example

### Input

```text
0 0 0
0 0 1
0 0 1
0 0 1
1 1 1
1 1 1
1 1 1
1 1 1
```

### Output

```text
YES
0 0 0
0 0 1
0 1 0
1 0 0
0 1 1
1 0 1
1 1 0
1 1 1
```

### Input

```text
0 0 0
0 0 0
0 0 0
0 0 0
1 1 1
1 1 1
1 1 1
1 1 1
```

### Output

```text
NO
```

## Solution

Each input line keeps its own three numbers, but those three numbers may be permuted. Since a line has
only three coordinates, it has at most `3! = 6` different permutations. Therefore we can try all
choices by DFS:

1. Process the eight lines from top to bottom.
2. Sort the current line first, then enumerate its unique permutations with `nextPermutation`.
3. Append the chosen permutation as the restored point for this line.
4. Recurse to the next line.

If two restored lines become the same point, they cannot be eight distinct cube vertices, so the code
skips that choice immediately.

After choosing one permutation for all eight lines, the helper `check` verifies whether the eight
points form a cube. Pick the first point as one cube vertex and compute the squared distances from it
to the other seven points.

For any cube vertex:

1. There are exactly `3` adjacent vertices at distance `s^2`.
2. There are exactly `3` face-diagonal vertices at distance `2s^2`.
3. There is exactly `1` opposite vertex at distance `3s^2`.

Here `s` is the non-zero side length of the cube. So the smallest positive squared distance from the
first point must appear `3` times, twice that distance must appear `3` times, and three times that
distance must appear `1` time. Any other distance pattern is impossible.

Once the DFS finds a permutation assignment that passes `check`, it prints `YES` and those eight
restored points. If every assignment fails, it prints `NO`.

### Correctness

Every valid output must choose, for each input line, one permutation of exactly the three numbers from
that line. The DFS enumerates all unique permutations for each row and combines them across all eight
rows, so it eventually considers every possible valid restoration.

If the DFS accepts a restoration, `check` has found the cube distance pattern from one vertex: three
edges, three face diagonals, and one body diagonal with squared lengths in ratio `1 : 2 : 3`. This is
exactly the distance structure of a cube from any vertex, so the accepted points form a cube.

If there exists a valid cube restoration, the DFS will consider that same row-by-row permutation
assignment. For that assignment, the eight points are cube vertices, so their distances from the first
point have the required `3, 3, 1` counts. Thus `check` returns true, and the algorithm finds a valid
answer.

Therefore the algorithm prints `YES` exactly when the scrambled rows can be restored into cube
vertices.

### Complexity

Each row has at most `6` unique permutations, so the DFS tries at most `6^8` assignments. Checking one
assignment uses only the eight points, so it is `O(1)`.

The total time complexity is `O(6^8)`, and the extra memory usage is `O(8)`.
