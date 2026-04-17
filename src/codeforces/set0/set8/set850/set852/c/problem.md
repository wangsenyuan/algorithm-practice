# Problem

Bill is a famous mathematician in BubbleLand. Thanks to his discoveries, he earned enough money to build a beautiful house. Unfortunately, because he did not pay property tax on time, the court decided to punish him by taking part of his property.

Bill's property is modeled as a convex regular $2n$-gon:

$A_0, A_1, \ldots, A_{2n-1}, A_{2n}$, where $A_{2n} = A_0$,

and every side has length $1$.

The process is:

1. For each edge $A_kA_{k+1}$ ($k = 0 \ldots 2n-1$), split it into $n$ equal segments of length $\frac{1}{n}$, creating points $P_0, P_1, \ldots, P_{n-1}$.
2. On each edge $A_{2k}A_{2k+1}$ ($k = 0 \ldots n-1$), the court chooses one point $B_{2k} = P_i$ for some $i \in [0, n-1]$.
3. On each edge $A_{2k+1}A_{2k+2}$ ($k = 0 \ldots n-1$), Bill chooses one point $B_{2k+1} = P_i$ for some $i \in [0, n-1]$.

Bill keeps the area inside polygon:

$B_0, B_1, \ldots, B_{2n-1}$.

Bill already knows all court choices $B_{2k}$. Help him choose $B_{2k+1}$ to maximize the kept area.

## Input

The first line contains integer $n$ ($2 \le n \le 50000$), the half-number of polygon edges (total edges are $2n$).

The second line contains $n$ distinct integers $B_{2k}$ ($0 \le B_{2k} \le n-1$, $k = 0 \ldots n-1$), separated by spaces.
If $B_{2k} = i$, it means the court chose point $P_i$ on side $A_{2k}A_{2k+1}$.

## Output

Print $n$ distinct integers separated by spaces: $B_1, B_3, \ldots, B_{2n-1}$ — Bill's choices that maximize the area.

If multiple optimal solutions exist, print any of them.

## Example

### Input

```text
3
0 1 2
```

### Output

```text
0 2 1
```

## Note

For the sample, an optimal choice is:

- $B_1 = P_0$
- $B_3 = P_2$
- $B_5 = P_1$

## Solution Idea

The geometric statement looks intimidating, but after writing the area carefully it turns into a simple permutation optimization.

### 1. Parameterize each chosen point on an edge

On every side, the chosen point is one of:

- `P_0, P_1, ..., P_{n-1}`

So we can represent each choice by an integer from `0` to `n-1`.

The input gives:

- `b_i` for even edges

and we must choose:

- `x_i` for odd edges

where:

- `x_0, x_1, ..., x_{n-1}` must be a permutation of `0, 1, ..., n-1`

The permutation condition is exactly the statement that Bill must choose `n` distinct points.

### 2. Area as “full polygon minus corner triangles”

The polygon:

- `B_0, B_1, ..., B_{2n-1}`

lies inside the regular `2n`-gon.
So its area can be written as:

- area of the big polygon
- minus the sum of the small cut-off triangles near each vertex

The area of the big polygon is constant, so we only need to minimize the total removed area.

Each removed triangle depends only on:

- the chosen point on the left side of a vertex
- the chosen point on the right side of that vertex

Now let us compute the contribution of one variable `x_i` carefully.

Recall:

- on edge `A_{2i}A_{2i+1}`, the court chooses `B_{2i} = P_{b_i}`
- on edge `A_{2i+1}A_{2i+2}`, Bill chooses `B_{2i+1} = P_{x_i}`

Because `P_t` divides an edge into `n` equal parts:

- distance from the beginning of the edge to `P_t` is `t / n`
- distance from the end of the edge to `P_t` is `(n - t) / n`

#### Corner at `A_{2i+1}`

The two adjacent chosen points are:

- `B_{2i}` on `A_{2i}A_{2i+1}`
- `B_{2i+1}` on `A_{2i+1}A_{2i+2}`

So their distances from `A_{2i+1}` are:

- `dist(A_{2i+1}, B_{2i}) = (n - b_i) / n`
- `dist(A_{2i+1}, B_{2i+1}) = x_i / n`

Since all polygon angles are fixed, the triangle area is a common constant times the product of these two distances.
So the triangle at `A_{2i+1}` contributes proportionally to:

- `x_i * (n - b_i)`

#### Corner at `A_{2i+2}`

Now look at the next vertex `A_{2i+2}`.
Its adjacent chosen points are:

- `B_{2i+1}` on `A_{2i+1}A_{2i+2}`
- `B_{2i+2}` on `A_{2i+2}A_{2i+3}`

Their distances from `A_{2i+2}` are:

- `dist(A_{2i+2}, B_{2i+1}) = (n - x_i) / n`
- `dist(A_{2i+2}, B_{2i+2}) = b_{i+1} / n`

So this triangle contributes proportionally to:

- `(n - x_i) * b_{i+1}`

#### Total contribution involving `x_i`

The variable `x_i` appears only in these two neighboring corner triangles, so its whole contribution is:

- `x_i (n - b_i) + (n - x_i) b_{i+1}`

Expand:

- `= x_i n - x_i b_i + n b_{i+1} - x_i b_{i+1}`
- `= n b_{i+1} + x_i (n - b_i - b_{i+1})`

The term `n b_{i+1}` does not depend on Bill’s choices, so globally it is just part of a constant.

Therefore minimizing the total removed area is equivalent to minimizing:

- `sum x_i * (n - b_i - b_{i+1})`

which is the same as maximizing:

- `sum x_i * (b_i + b_{i+1})`

where indices are cyclic modulo `n`.

So after all geometry is eliminated, the whole problem becomes:

- choose a permutation `x_i` of `0..n-1`
- maximize a weighted sum

### 3. Reduced combinatorial problem

Define:

- `d_i = b_i + b_{i+1}`

Then we want to maximize:

- `sum x_i * d_i`

subject to:

- `x_i` is a permutation of `0, 1, ..., n-1`

This is now a standard rearrangement inequality situation.

### 4. Best permutation by sorting

If we have two positions `i` and `j` with:

- `d_i < d_j`

then to maximize the total sum, the smaller assigned value should go to `i` and the larger assigned value should go to `j`.

Why?

Suppose currently:

- `x_i > x_j`

Swap them.
The change in objective is:

- `x_j*d_i + x_i*d_j - (x_i*d_i + x_j*d_j)`
- `= (x_i - x_j) * (d_j - d_i)`

Since:

- `x_i - x_j > 0`
- `d_j - d_i > 0`

the change is positive.
So the swap improves the answer.

Therefore, in an optimal solution:

- if `d_i < d_j`, then `x_i < x_j`

That means:

- sort positions by `d_i`
- assign `0, 1, 2, ..., n-1` in that order

If some `d_i` are equal, any order among them is optimal.

### 5. Resulting algorithm

1. Compute:
   - `d_i = b_i + b_{i+1}` for all `i`
2. Sort positions `i` by `d_i`
3. Assign:
   - smallest number `0` to the smallest `d_i`
   - next number `1` to the next smallest `d_i`
   - ...
   - largest number `n-1` to the largest `d_i`

The resulting array is exactly Bill’s optimal choice.

### 6. Complexity

We only sort `n` numbers, so the complexity is:

- `O(n log n)`

with `O(n)` extra memory.

## Summary

The hard-looking geometry collapses to:

- maximize `sum x_i * (b_i + b_{i+1})`

over a permutation `x_i`.

Then the answer follows immediately from the rearrangement inequality:

- sort by `b_i + b_{i+1}`
- assign `0..n-1` in that order.
