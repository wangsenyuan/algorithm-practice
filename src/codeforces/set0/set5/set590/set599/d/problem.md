# Problem Statement

Spongebob is already tired trying to reason his weird actions and calculations, so he simply asked you to find all pairs of n and m, such that there are exactly x distinct squares in the table consisting of n rows and m columns. 

For example, in a 3 × 5 table there are:
- 15 squares with side one
- 8 squares with side two  
- 3 squares with side three

The total number of distinct squares in a 3 × 5 table is 15 + 8 + 3 = 26.

## Input

The first line of the input contains a single integer x (1 ≤ x ≤ 10^18) — the number of squares inside the tables Spongebob is interested in.

## Output

1. First print a single integer k — the number of tables with exactly x distinct squares inside.
2. Then print k pairs of integers describing the tables. Print the pairs in the order of increasing n, and in case of equality — in the order of increasing m.

## Examples

### Example 1
**Input:**
```
26
```

**Output:**
```
6
1 26
2 9
3 5
5 3
9 2
26 1
```

### Example 2
**Input:**
```
2
```

**Output:**
```
2
1 2
2 1
```

### Example 3
**Input:**
```
8
```

**Output:**
```
4
1 8
2 3
3 2
8 1
```


### ideas
1. 对于给定的(w <= h), 计算它中间的正方形的数量
2. 先考虑w, 1, 2, .... w 的边长的个数 = w, w - 1, w - 2, .... 1
3. 考虑h, 1, 2.... w 的边长的个数 = h, h - 1, h - 2, ... h - w + 1
4. 那么总的数量 = w * h + (w - 1) * (h - 1) + ... 1 * (h - w + 1) = x
5. 所以，对于给定的w，可以计算出对应的h
6. 