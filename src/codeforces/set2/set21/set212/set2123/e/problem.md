# Problem E: MEX After Removing Elements

## Problem Description

Define the **MEX** (minimum excluded value) of an array to be the smallest nonnegative integer not present in the array. For example:

- `MEX([2,2,1]) = 0` because 0 is not in the array.
- `MEX([3,1,0,1]) = 2` because 0 and 1 are in the array but 2 is not.
- `MEX([0,3,1,2]) = 4` because 0, 1, 2, and 3 are in the array but 4 is not.

You are given an array $a$ of size $n$ of nonnegative integers.

For all $k$ ($0 \leq k \leq n$), count the number of possible values of `MEX(a)` after removing exactly $k$ values from $a$.

## Input

- The first line contains an integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases.
- The first line of each test case contains one integer $n$ ($1 \leq n \leq 2 \cdot 10^5$) — the size of the array $a$.
- The second line of each test case contains $n$ integers, $a_1, a_2, \ldots, a_n$ ($0 \leq a_i \leq n$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output a single line containing $n+1$ integers — the number of possible values of `MEX(a)` after removing exactly $k$ values, for $k = 0, 1, \ldots, n$.

## Example

### Input
```
5
5
1 0 0 1 2
6
3 2 0 4 5 1
6
1 2 0 1 3 2
4
0 3 4 1
5
0 0 0 0 0
```

### Output
```
1 2 4 3 2 1
1 6 5 4 3 2 1
1 3 5 4 3 2 1
1 3 3 2 1
1 1 1 1 1 1
```

## Note

In the first sample, consider $k = 1$. If you remove a 0, then you get the following array:

```
1 0 1 2
```

So we get `MEX(a) = 3`. Alternatively, if you remove the 2, then you get the following array:

```
1 0 0 1
```

So we get `MEX(a) = 2`. It can be shown that these are the only possible values of `MEX(a)` after removing exactly one value. So the output for $k = 1$ is 2.

## ideas
1. 考虑原来mex(a)的数，考虑x = 0,1... mex(a)
2. 如果删除k个数， 如果 cnt[x] > k, 肯定无法被删除，所以cnt[x]<= k要成立
3. 除此之外，suf(x) >= k(超过x，包括x的数的个数，要大于等于k)
4. 可以排序+heap