Dima loves representing an odd number as the sum of multiple primes, and Lisa loves it when there are at most three primes. Help them to represent the given number as the sum of at most three primes.

More formally, you are given an odd number $n$. Find a set of numbers $p_i$ $(1 \le i \le k)$ such that:

- $1 \le k \le 3$;
- each $p_i$ is a prime;
- the numbers $p_i$ do not necessarily have to be distinct.

It is guaranteed that at least one possible solution exists.

## Input

The single line contains an odd number $n$ $(3 \le n < 10^9)$.

## Output

In the first line print $k$ $(1 \le k \le 3)$, showing how many numbers are in the representation you found.

In the second line print numbers $p_i$ in any order. If there are multiple possible solutions, you can print any of them.

## Examples

### Input
```
27
```

### Output
```
3
5 11 11
```

## Note

A prime is an integer strictly larger than one that is divisible only by one and by itself.


### ideas
1. 如果n是一个质数 = n
2. 如果 n - 2 是一个质数 => n = 2 + pi
3. 如果 n - p1 肯定是一个偶数，
4. 