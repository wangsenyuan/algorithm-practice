# Problem C

Ivan likes to learn different things about numbers, but he is especially interested in really big numbers. Ivan thinks that a positive integer number x is really big if the difference between x and the sum of its digits (in decimal representation) is not less than s. To prove that these numbers may have different special properties, he wants to know how rare (or not rare) they are — in fact, he needs to calculate the quantity of really big numbers that are not greater than n.

Ivan tried to do the calculations himself, but soon realized that it's too difficult for him. So he asked you to help him in calculations.

## Input

The first (and the only) line contains two integers n and s (1 ≤ n, s ≤ 10^18).

## Output

Print one integer — the quantity of really big numbers that are not greater than n.

## Examples

### Example 1

**Input:**

```text
12 1
```

**Output:**

```text
3
```

### Example 2

**Input:**

```text
25 20
```

**Output:**

```text
0
```

### Example 3

**Input:**

```text
10 9
```

**Output:**

```text
1
```

## Note

In the first example numbers 10, 11 and 12 are really big.

In the second example there are no really big numbers that are not greater than 25 (in fact, the first really big number is 30: 30 - 3 ≥ 20).

In the third example 10 is the only really big number (10 - 1 ≥ 9).


### ideas
1. 考虑一个数x， x - digit_sum(x) >= s
2. 如果x满足条件, x + 1 呢？
3. 如果x的个位数 < 9, 那么 digit_sum(x + 1) = digit_sum(x) + 1
4. 所以 (x + 1) - (digit_sum(x+1)) >= s 仍然成立
5. 如果x的个位数 = 9
6. digit_sum(x+1) <= digit_sum(x) - 9 + 1 所以， 仍然满足条件
7. 所以可以二分x
8. 