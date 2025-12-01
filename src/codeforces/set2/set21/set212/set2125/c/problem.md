A prime number is a positive integer that has exactly two divisors: 1 and itself. The first several prime numbers are 2, 3, 5, 7, 11, ….

Prime factorization of a positive integer is representing it as a product of prime numbers. For example:

- the prime factorization of 111 is 3⋅37;
- the prime factorization of 43 is 43;
- the prime factorization of 12 is 2⋅2⋅3.

For every positive integer, its prime factorization is unique (if you don't consider the order of primes in the product).

We call a positive integer good if all primes in its factorization consist of at least two digits. For example:

- 343 = 7⋅7⋅7 is not good;
- 111 = 3⋅37 is not good;
- 1111 = 11⋅101 is good;
- 43 = 43 is good.

You have to calculate the number of good integers from 𝑙 to 𝑟 (endpoints included).

## Input

The first line contains one integer 𝑡 (1 ≤ 𝑡 ≤ 10³) — the number of test cases.

Each test case consists of one line containing two integers 𝑙 and 𝑟 (2 ≤ 𝑙 ≤ 𝑟 ≤ 10¹⁸).

## Output

For each test case, print one integer — the number of good integers from 𝑙 to 𝑟.

## Example

**Input:**
```
4
2 100
2 1000
13 37
2 1000000000000000000
```

**Output:**
```
21
227
7
228571428571428570
```
