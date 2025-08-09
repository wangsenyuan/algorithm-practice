### Problem
You are given a permutation p1, p2, …, pn of length n.

You must build an array a1, a2, …, an in the following way: for each i (1≤i≤n), set either ai = pi or ai = 2n − pi. Find the minimum possible number of inversions in the array a.

### Definitions
- A permutation of length n is an array consisting of n distinct integers from 1 to n in arbitrary order. For example, [2,3,1,5,4] is a permutation, but [1,2,2] is not (2 appears twice), and [1,3,4] is not when n = 3.
- An inversion in the array a is a pair of indices (i, j) such that 1≤i<j≤n and ai > aj.

### Input
- The first line contains the number of test cases t (1≤t≤$10^3$).
- Each test case consists of:
  - A line with an integer n (2≤n≤$5\cdot 10^3$).
  - A line with n integers p1, p2, …, pn (1≤pi≤n), which form a permutation.
- The sum of n over all test cases does not exceed $5\cdot 10^3$.

### Output
For each test case, print a single integer — the minimum possible number of inversions in the array a.

### Example
Input
```text
5
2
2 1
3
2 1 3
4
4 3 2 1
5
2 3 1 5 4
6
2 3 4 1 5 6
```

Output
```text
0
1
0
2
2
```

### Explanation
- In the first test case, the only optimal array a is [2,3], with 0 inversions.
- In the second test case, one optimal array a is [2,5,3], with 1 inversion. Another possible optimal array a is [2,1,3].

## ideas
1. 感觉是两端，前面的不使用，后面的使用

