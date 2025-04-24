# Problem F - Expected Value

You are given an integer sequence A = (A₁, A₂, ..., Aₙ) of length N.

## Operations

We will perform M operations on sequence A. For each operation i (from 1 to M):

1. Choose an integer p uniformly at random between Lᵢ and Rᵢ (inclusive)
2. Set Aₚ = Xᵢ

## Task

For the final sequence A after all M operations are complete:
- Calculate the expected value of Aᵢ for each i = 1, 2, ..., N
- All values should be taken modulo 998244353

### ideas
1. y1 = y0 + (x1 - y0) * n1 (n1 是 r - l + 1 的逆)
2. 