# Problem D

You are given an array a₁, a₂, …, aₙ consisting of n integers.

Your goal is to make it strictly increasing. To achieve that, you perform each of the following operations exactly once:

1. Remove any element;
2. Select any number of elements (possibly none or all n − 1) and add 1 to them.

Note that you are not allowed to rearrange the elements of the array.

For the resulting array a′, a′₁ < a′₂ < ⋯ < a′ₙ₋₁ should hold. Determine if it's possible to achieve that.

## Input

The first line contains a single integer t (1 ≤ t ≤ 10⁴) — the number of test cases.

The first line of each test case contains a single integer n (2 ≤ n ≤ 2·10⁵) — the number of elements of the array.

The second line contains n integers a₁, a₂, …, aₙ (1 ≤ aᵢ ≤ 10⁶).

The sum of n over all test cases doesn't exceed 2·10⁵.

## Output

For each test case, print YES if it's possible to remove one element and add 1 to some elements (possibly none or all), so that the array becomes strictly increasing. Otherwise, print NO.

## Example

**Input:**
```
8
4
4 4 1 5
5
4 4 1 5 5
2
10 5
3
1 2 3
3
2 1 1
4
1 1 1 1
4
1 3 1 2
5
1 1 3 3 1
```

**Output:**
```
YES
NO
YES
YES
YES
NO
YES
YES
```

## Note

In the first test case, you can remove the third element and add 1 to the second and the last element. a′ will become [4, 5, 6], which is strictly increasing.

In the second test case, there is no way to perform the operations so that the result is strictly increasing.

In the third test case, you can remove either of the elements.

In the fourth test case, you are already given a strictly increasing array, but you still have to remove an element. The result a′ can be [1, 3], for example.

## ideas
1. dp[i][0] 表示对第i个数，不+1情况下，后缀是否满足递增（没有删除的时候）
2. dp[i][1] 表示对第i个数 +1的情况下，后缀是否递增
3. dp[i][0] = if a[i] < a[i+1] && dp[i+1][0] or a[i] == a[i+1] && dp[i+1][1] else false
4. 那么第一个dp[i][0] = false的，就是必须删除的