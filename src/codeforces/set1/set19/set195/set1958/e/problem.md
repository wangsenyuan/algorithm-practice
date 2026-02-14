# Problem E

Suppose you have a permutation p of n integers — an array where each element is an integer from 1 to n, and every integer from 1 to n appears exactly once.

In one operation, you remove every element of this permutation which is less than at least one of its neighbors. For example, when you apply the operation to [3, 1, 2, 5, 4], you get [3, 5]. If you apply an operation again, you get [5].

It's easy to see that after applying a finite number of operations, you get an array consisting of a single integer n.

You are given two integers n and k. Find a permutation of n integers such that it becomes an array consisting of a single element n after exactly k operations (and not earlier).

## Input

The first line contains one integer t (1 ≤ t ≤ 2000) — the number of test cases.

Each test case consists of one line containing two integers n and k (2 ≤ n ≤ 100; 1 ≤ k ≤ n − 1).

## Output

For each test case, print the answer as follows:

- If a permutation of size n which becomes an array consisting of a single element n after exactly k operations does not exist, print −1;
- Otherwise, print n distinct integers from 1 to n — the requested permutation. If there are multiple such permutations, print any of them.

## Example

**Input:**
```
4
5 2
5 4
2 1
3 2
```

**Output:**
```
3 1 2 5 4
-1
1 2
2 1 3
```

## Solution

Let the current size of the array be m. When an operation is applied, at least ⌊m/2⌋ elements will be removed, because if an element is not removed, both of its neighbors will be removed.

So, for every value of k, we can calculate the minimum possible size of an array n (let's call it fₖ). For k = 1, fₖ = 2; for k > 1, this value is fₖ = 2·fₖ₋₁ − 1. We can actually derive an exact formula fₖ = 2ᵏ⁻¹ + 1, but we don't actually have to, since it can be calculated in O(n).

If n < fₖ, obviously, there is no solution. Let's show how to construct the answer for n = fₖ, and then let's deal with the case n > fₖ.

For k = 1, we can use an array [1, 2]. If we want to go from k to k+1, we can insert elements which will be removed on the first operation between the existing elements, while increasing the existing elements by the number of elements we added (so that they are actually greater). For example, [1, 2] → [2, 1, 3] → [4, 1, 3, 2, 5] (every element on even position is a newly inserted element). That's how we can resolve the case n = fₖ.

What if n > fₖ? Construct the answer for n = fₖ. Let's say we have m = n − fₖ "extra" elements. Let's insert them in such a way that they will all be deleted during the first operation, and will not affect the existing elements. For example, we can increase every existing element by m, and then add 1, 2, 3, …, m to the beginning of the array, so that they don't affect existing elements, and are all deleted during the first operation.

You have to take into account that fₖ might overflow. And you should first check that fₖ ≤ n, and only then construct the answer, since it takes O(fₖ + n).
