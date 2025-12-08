# Problem Description

You are given an integer array 𝑎 of length 𝑛.

You can perform the following operation: choose a range [𝑙,𝑟] (1≤𝑙≤𝑟≤𝑛) and replace the value of elements 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟 with (𝑙+𝑟).

Your task is to calculate the maximum possible total array sum if you can perform the aforementioned operation at most once.

## Input

The first line contains a single integer 𝑡 (1≤𝑡≤10⁴) — the number of test cases.

The first line of each test case contains a single integer 𝑛 (1≤𝑛≤2⋅10⁵).

The second line contains 𝑛 integers 𝑎₁,𝑎₂,…,𝑎ₙ (0≤𝑎ᵢ≤2𝑛).

**Additional constraint on the input:** the sum of 𝑛 over all test cases doesn't exceed 2⋅10⁵.

## Output

For each test case, print a single integer — the maximum possible total array sum if you can perform the aforementioned operation at most once.

## Example

### Input
```
4
3
2 5 1
2
4 4
4
1 3 2 1
5
3 2 0 9 10
```

### Output
```
13
8
20
32
```

## Note

In the first example, you can perform the operation on the subarray [3,3], resulting in the array [2,5,6] and the sum 13.

In the second example, you don't need to perform any operation.

In the third example, you can perform the operation on the subarray [1,4], resulting in the array [5,5,5,5] and the sum 20.

In the fourth example, you can perform the operation on the subarray [2,3], resulting in the array [3,5,5,9,10] and the sum 32.


### ideas
1. (r - l + 1) * (r + l) - (sum[r] - sum(l-1))
2. (r + 1 - l) * (r + l) - (sum[r] - sum[l-1])
3. (r - l) *(r + l) + (r + l) - (sum[r] - sum[l-1])
4. r * r - l * l + r + l - sum[r] + sum[l-1]
5. (r * r + r - sum[r]) - (l * l - l - sum[l-1])