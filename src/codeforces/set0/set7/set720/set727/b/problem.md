# Problem B: Bill Calculator

## Problem Description

Vasily exited from a store and now he wants to recheck the total price of all purchases in his bill. The bill is a string in which the names of the purchases and their prices are printed in a row without any spaces.

The bill has the format: `"name1price1name2price2...namenpricen"`, where:
- `namei` (name of the i-th purchase) is a non-empty string of length not more than 10, consisting of lowercase English letters
- `pricei` (the price of the i-th purchase) is a non-empty string, consisting of digits and dots (decimal points)
- It is possible that purchases with equal names have different prices

## Price Format Rules

The price of each purchase is written in the following format:

1. **Integer prices**: If the price is an integer number of dollars, cents are not written
2. **Decimal prices**: After the number of dollars, a dot (decimal point) is written followed by cents in a two-digit format (if number of cents is between 1 and 9 inclusively, there is a leading zero)
3. **Thousands separator**: Every three digits (from less significant to the most) in dollars are separated by dot (decimal point)
4. **No leading zeros**: No extra leading zeroes are allowed
5. **Start and end**: The price always starts with a digit and ends with a digit

### Valid Price Examples:
- `"234"` - integer price
- `"1.544"` - decimal price
- `"149.431.10"` - price with thousands separator and cents
- `"0.99"` - decimal price less than 1 dollar
- `"123.05"` - decimal price with leading zero in cents

### Invalid Price Examples:
- `".333"` - starts with dot
- `"3.33.11"` - multiple dots in wrong positions
- `"12.00"` - trailing zeros in cents
- `".33"` - starts with dot
- `"0.1234"` - more than 2 decimal places
- `"1.2"` - single digit after decimal point

## Task

Write a program that will find the total price of all purchases in the given bill.

## Input

The only line of the input contains a non-empty string `s` with length not greater than 1000 — the content of the bill.

**Constraints:**
- It is guaranteed that the bill meets the format described above
- Each price in the bill is not less than one cent and not greater than $10^6$

## Output

Print the total price exactly in the same format as prices given in the input.

## Examples

### Example 1
**Input:**
```
chipsy48.32televizor12.390
```

**Output:**
```
12.438.32
```

**Explanation:**
- `chipsy` costs `48.32`
- `televizor` costs `12.390`
- Total: `48.32 + 12.390 = 60.71` → formatted as `12.438.32` (incorrect in example, should be `60.71`)

### Example 2
**Input:**
```
a1b2c3.38
```

**Output:**
```
6.38
```

**Explanation:**
- `a` costs `1`
- `b` costs `2`
- `c` costs `3.38`
- Total: `1 + 2 + 3.38 = 6.38`

### Example 3
**Input:**
```
aa0.01t0.03
```

**Output:**
```
0.04
```

**Explanation:**
- `aa` costs `0.01`
- `t` costs `0.03`
- Total: `0.01 + 0.03 = 0.04`

