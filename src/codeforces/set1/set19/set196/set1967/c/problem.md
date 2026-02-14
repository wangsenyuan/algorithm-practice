# Problem C

Let lowbit(x) denote the value of the lowest binary bit of x, e.g. lowbit(12) = 4, lowbit(8) = 8.

For an array a of length n, if an array s of length n satisfies

sₖ = (Σᵢ₌ₖ₋ₗₒᵥᵦᵢₜ₍ₖ₎₊₁ᵏ aᵢ) mod 998244353

for all k, then s is called the **Fenwick Tree** of a. Let's denote it as s = f(a).

For a positive integer k and an array a, fᵏ(a) is defined as follows:

- fᵏ(a) = f(a) if k = 1
- fᵏ(a) = f(fᵏ⁻¹(a)) otherwise

You are given an array b of length n and a positive integer k. Find an array a that satisfies 0 ≤ aᵢ < 998244353 and fᵏ(a) = b. It can be proved that an answer always exists. If there are multiple possible answers, you may print any of them.

## Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10⁴). The description of the test cases follows.

The first line of each test case contains two positive integers n (1 ≤ n ≤ 2·10⁵) and k (1 ≤ k ≤ 10⁹), representing the length of the array and the number of times the function f is performed.

The second line of each test case contains an array b₁, b₂, …, bₙ (0 ≤ bᵢ < 998244353).

It is guaranteed that the sum of n over all test cases does not exceed 2·10⁵.

## Output

For each test case, print a single line, containing a valid array a of length n.

## Example

**Input:**
```
2
8 1
1 2 1 4 1 2 1 8
6 2
1 4 3 17 5 16
```

**Output:**
```
1 1 1 1 1 1 1 1
1 2 3 4 5 6
```

## Note

In the first test case, it can be seen that f¹([1,1,1,1,1,1,1,1]) = [1,2,1,4,1,2,1,8].

In the second test case, it can be seen that f²([1,2,3,4,5,6]) = f¹([1,3,3,10,5,11]) = [1,4,3,17,5,16].


### ideas
1. 感觉很复杂～
2. 