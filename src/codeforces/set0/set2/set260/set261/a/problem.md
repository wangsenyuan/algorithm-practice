# A. Maxim and Discounts

[Problem link](https://codeforces.com/problemset/problem/261/A)

**Contest:** [Codeforces Round 160 (Div. 1)](https://codeforces.com/contest/261)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Maxim must buy `n` items. There are `m` discount types. Discount `i` requires a cart of exactly `q_i` items. With that cart he may also take up to two extra items for free, and each free item must cost no more than the cheapest item in the cart.

He may use the same discount any number of times, and he may buy items with no discount. There are enough carts. Find the minimum amount he must pay.

## Constraints

- `1 ≤ m ≤ 10^5`
- `1 ≤ q_i ≤ 10^5`
- `1 ≤ n ≤ 10^5`
- `1 ≤ a_i ≤ 10^4`

## Input

```text
m
q1 q2 ... qm
n
a1 a2 ... an
```

## Output

Print the minimum cost.

## Sample Input 1

```text
1
2
4
50 50 100 100
```

## Sample Output 1

```text
200
```

Put the two items that cost 100 in the cart and take the two items that cost 50 for free. He pays 200.

## Sample Input 2

```text
2
2 3
5
50 50 50 50 50
```

## Sample Output 2

```text
150
```

Buy 3 items and take the other 2 for free. He pays 150.

## Sample Input 3

```text
1
1
7
1 1 1 1 1 1 1
```

## Sample Output 3

```text
3
```

Each use of the discount pays for 1 item and takes at most 2 free items of equal price, so the seven items cost 3.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
