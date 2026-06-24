# E. Friendly Gifts

[Problem link](https://codeforces.com/problemset/problem/2236/E)

**Contest:** [Codeforces Round 1103 (Div. 3)](https://codeforces.com/contest/2236)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Arseniy wants to give each of his two friends — Dabir and Egor — an array of numbers of the
**same length**.

An array `b` is called **good** if its elements can be rearranged into consecutive integers with
difference `1`, i.e. into something like `x, x+1, ..., x+k` for some integer `x`.

Arseniy already has an array `a` of length `n`. He will cut **two non-overlapping subsegments**
of equal length from `a` and give one subsegment to each friend. Both subsegments must be good
arrays.

Find the **maximum possible length** of these two arrays. If no such pair exists, the answer is
`0`.

Equivalently, a subsegment is good iff all its elements are distinct and
`max - min + 1` equals the segment length.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with an integer `n` (`1 <= n <= 5000`);
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `5000`.

## Output

For each test case, print one integer — the maximum possible length of the two chosen subsegments.

## Example

### Input

```text
5
1
1
2
1 2
2
1 4
4
2 1 4 3
5
1 2 3 4 5
```

### Output

```text
0
1
1
2
1
```

## Note

- In the first test case, it is impossible to choose two non-overlapping subsegments of equal
  positive length, so the answer is `0`.
- In the second test case, the maximum length is `1`; one choice is `[1]` and `[2]`.
- In the fourth test case, the maximum length is `2`; one choice is `[2, 1]` and `[4, 3]`.
- In the fifth test case, the maximum length is `1`; one choice is `[1]` and `[2]`.

## Solution

Find the maximum length `d` such that `a` contains **two disjoint good subsegments**, each of
length `d`.

### Short recap

- A subsegment is **good** iff its elements are distinct and form consecutive integers
  `w, w+1, …, w+d-1` for some minimum value `w`.
- Equivalently, with length `d = r - l + 1`, distinctness plus
  `max - min + 1 = d`, or the sum identity below.
- We need two **non-overlapping** subsegments of the **same** length `d`, both good.

### Key observations

1. **Sum test for consecutiveness.** If a subsegment has length `d`, minimum `w`, and distinct
   elements, it is good iff

   ```text
   sum = w·d + d·(d-1)/2
   ```

   because `w + (w+1) + … + (w+d-1) = w·d + (0+1+…+(d-1))`.

2. **Record minima by length.** For each starting index `i`, extend right while elements stay
   distinct. Whenever the sum identity holds, append the current minimum `w` to `todo[d]`, where
   `d` is the segment length. After the inner loop, clear the frequency marks for `[i, j)`.

3. **Pair good segments by value range.** Two good segments of length `d` with minima `v` and
   `v+d` use value sets `[v, v+d-1]` and `[v+d, v+2d-1]`, which are disjoint. If both minima
   appear among length-`d` good segments, one can choose two subsegments with those value ranges
   that do not overlap in the array (same idea as checking `good[l][l+d-1]` and
   `good[l+d][l+2d-1]` in the editorial).

4. **Search the maximum `d`.** For each `d` from `1` to `⌊n/2⌋`, sort `todo[d]` and check whether
   some `v` has a partner `v+d` via binary search. Keep the largest feasible `d`.

### Algorithm (matching `solve`)

**Step 1 — enumerate good segments.**

```text
todo[d] = list of minima w for good subsegments of length d
```

For each start `i`:

- Walk `j` right while `freq[a[j]] == 0`.
- Track `w = min` on the segment and `sum` of values.
- If `w·d + (d-1)·d/2 == sum` with `d = j-i+1`, append `w` to `todo[d]`.
- On duplicate, stop; reset `freq` for `[i, j)`.

**Step 2 — pair by consecutive value blocks.**

For `d = 1 … n/2`:

- Sort `todo[d]`.
- For each `v` in `todo[d]`, binary-search for `v+d`.
- If found, set `best = d` (later iterations overwrite with larger `d`).

**Step 3 — output `best`** (remains `0` if no pair exists).

### Walkthrough (sample 4)

`a = [2, 1, 4, 3]`, answer `2`.

- Length `2`: good segments include `[2,1]` (min `1`) and `[4,3]` (min `3`).  
  `todo[2]` contains `1` and `3`; `1 + 2 = 3` → pair found → `best = 2`.

### Complexity

- Enumerating segments: `O(n²)` time, `O(n)` extra space for `freq` and `todo`.
- Pairing: `O(n² log n)` from sorting each `todo[d]` (with `∑ d ≤ O(n²)` entries total).
- Fits `n ≤ 5000`, `∑ n ≤ 5000`.

### ideas

- Good segment = consecutive values; detect by distinctness + sum formula
  `w·d + d(d-1)/2`.
- Store minimum `w` per length in `todo[d]`; two blocks `[v, v+d-1]` and `[v+d, v+2d-1]`
  pair when both minima appear.
