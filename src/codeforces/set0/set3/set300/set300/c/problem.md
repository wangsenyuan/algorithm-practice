# Problem Description

Vitaly is a very weird man. He's got two favorite digits **a** and **b**. Vitaly calls a positive integer **good**, if the decimal representation of this integer only contains digits **a** and **b**. Vitaly calls a good number **excellent**, if the sum of its digits is a good number.

For example, let's say that Vitaly's favourite digits are 1 and 3, then:
- Number 12 isn't good (contains digit 2)
- Numbers 13 or 311 are good (only contain digits 1 and 3)
- Number 111 is excellent (sum of digits = 1+1+1 = 3, which is good)
- Number 11 isn't excellent (sum of digits = 1+1 = 2, which isn't good)

Now Vitaly is wondering, how many excellent numbers of length exactly **n** are there. As this number can be rather large, he asks you to count the remainder after dividing it by 1000000007 (10^9 + 7).

A number's length is the number of digits in its decimal representation without leading zeroes.

## Input

The first line contains three integers: **a**, **b**, **n** (1 ≤ a < b ≤ 9, 1 ≤ n ≤ 10^6).

## Output

Print a single integer — the answer to the problem modulo 1000000007 (10^9 + 7).

## Examples

### Example 1
**Input:**
```
1 3 3
```

**Output:**
```
1
```

### Example 2
**Input:**
```
2 3 10
```

**Output:**
```
165
```

### ideas
1. n (1e6), digit sum最多 = 9 * n = 1e7
2. 所以，需要计算出1e7中的good num
3. 然后dp[sum] = 表示组成digit sum = 是good的good的字符串的数量