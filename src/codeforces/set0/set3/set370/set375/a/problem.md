# Divisible by Seven

You have number a, whose decimal representation quite luckily contains digits 1, 6, 8, 9. Rearrange the digits in its decimal representation so that the resulting number will be divisible by 7.

Number a doesn't contain any leading zeroes and contains digits 1, 6, 8, 9 (it also can contain other digits). The resulting number also mustn't contain any leading zeroes.

## Input

The first line contains positive integer a in decimal notation. It is guaranteed that the representation of number a contains digits 1, 6, 8, 9. Number a doesn't contain any leading zeroes. The decimal representation contains at least 4 and at most 10^6 characters.

## Output

Print a number in decimal notation without leading zeroes — the result of the permutation.

If it is impossible to rearrange the digits of the number a in the required manner, print 0.

## Examples

### Example 1

**Input:**

```
1689
```

**Output:**

```
1869
```

### Example 2

**Input:**

```
18906
```

**Output:**

```
18690
```


### ideas
1. 包含1689，也可能包含其他字符; 没有前导0
2. 