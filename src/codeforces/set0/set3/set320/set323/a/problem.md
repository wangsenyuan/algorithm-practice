# Problem

You are given a cube of size $k \times k \times k$, made of $k^3$ unit cubes. Two unit cubes are **neighbouring** if they share a face.

Paint each of the $k^3$ unit cubes either **black** or **white** so that:

- Each white cube has exactly **2** white neighbours.
- Each black cube has exactly **2** black neighbours.

If no solution exists, output `-1`. Otherwise, output the painting layer by layer (see Output format).

## Input

A single integer $k$ ($1 \le k \le 100$) — the side length of the cube.

## Output

If there is no solution, print a single line: `-1`.

Otherwise, print the painting **by layers**: for the first layer print a $k \times k$ matrix (each of $k$ lines has $k$ characters), then for the second layer, and so on until the $k$-th layer. Orientation of the cube in space does not matter.

- Use `w` for a white unit cube and `b` for a black one.
- You may print extra empty lines; they will be ignored.

## Examples

### Example 1

**Input:**

```
1
```

**Output:**

```
-1
```

### Example 2

**Input:**

```
2
```

**Output:**

```
bb
ww

bb
ww
```

(First two lines = layer 1, next two lines = layer 2.)


### ideas
1. 考虑k = 3, 貌似不行，因为中间的那个不好搞
```
bwb
bwb
bbb
```
2. 考虑k = 4 (貌似可以)
```
bbww
wwbb
bbww
wwbb
```
1. 考虑 k = 6
```
bbwwbb
wwbbww
bbwwbb
wwbbww
```
2. 
