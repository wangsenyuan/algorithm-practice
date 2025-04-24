# Grid Transformation Problem

## Problem Description

You are given two grids, A and B, each with H rows and W columns.

### Initial Setup

For each pair of integers (i,j) satisfying 1≤i≤H and 1≤j≤W:
- Let (i,j) denote the cell in the i-th row and j-th column
- In grid A, cell (i,j) contains the integer A[i,j]
- In grid B, cell (i,j) contains the integer B[i,j]

### Operations Allowed

You can repeat the following operations any number of times (including zero):

1. Choose an integer i where 1≤i≤H−1 and swap the i-th and (i+1)-th rows in grid A
2. Choose an integer i where 1≤i≤W−1 and swap the i-th and (i+1)-th columns in grid A

### Goal

Determine if it's possible to make grid A identical to grid B using these operations. If possible, find the minimum number of operations required.

Note: Grid A is considered identical to grid B if and only if for all pairs (i,j) where 1≤i≤H and 1≤j≤W, the integer in cell (i,j) of grid A equals the integer in cell (i,j) of grid B.

### Constraints

- All input values are integers
- 2 ≤ H,W ≤ 5
- 1 ≤ A[i,j], B[i,j] ≤ 10^9

### ideas
1. brute force and check？
2. 5 * 4 * 3 * 2 * 1 = 120
3. 120 * 120