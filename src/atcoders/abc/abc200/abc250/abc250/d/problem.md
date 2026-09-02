# D - 250-like Number

[Problem link](https://atcoder.jp/contests/abc250/tasks/abc250_d)

**Contest:** [AtCoder Beginner Contest 250](https://atcoder.jp/contests/abc250)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

We call an integer k a "250-like number" if it satisfies the following condition.

- There exist primes p and q with p < q such that k = p * q^3.

How many 250-like numbers are there that do not exceed N?

## Constraints

- N is an integer satisfying 1 <= N <= 10^18

## Input

```text
N
```

## Output

Print the answer as an integer.

## Sample Input 1

```text
250
```

## Sample Output 1

```text
2
```

- 54 = 2 * 3^3, so it is a 250-like number.
- 250 = 2 * 5^3, so it is a 250-like number.

There are two 250-like numbers that do not exceed 250.

## Sample Input 2

```text
1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
123456789012345
```

## Sample Output 3

```text
226863
```
