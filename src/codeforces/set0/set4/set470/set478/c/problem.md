You have `r` red, `g` green and `b` blue balloons. To decorate a single table for the banquet you need exactly three balloons. Three balloons attached to some table shouldn't have the same color. What maximum number `t` of tables can be decorated if we know number of balloons of each color?

Your task is to write a program that for given values `r`, `g` and `b` will find the maximum number `t` of tables, that can be decorated in the required manner.

## Input

The single line contains three integers `r`, `g` and `b` (0 ≤ r, g, b ≤ 2·10⁹) — the number of red, green and blue balloons respectively. The numbers are separated by exactly one space.

## Output

Print a single integer `t` — the maximum number of tables that can be decorated in the required manner.

## Examples

### Input
```
5 4 3
```

### Output
```
4
```

### Input
```
1 1 1
```

### Output
```
1
```

### Input
```
2 3 3
```

### Output
```
2
```

## Note

In the first sample you can decorate the tables with the following balloon sets: "rgg", "gbb", "brr", "rrg", where "r", "g" and "b" represent the red, green and blue balls, respectively.



### ideas
1. 不能出现rrr， 其他的都ok
2. 如果 r = g = b => b
3. let r > g > b
4. 假设到x的时候，使的三个相等
5. (r - x) / 2 <= g - x + b - x
6. (r - x) <= 2 * (g + b - 2 * x)
7. r + 3 * x <= 2 * (g + b)
8. 3 * x <= 2 * (g + b) - r
9. （5， 4， 3）
10. 3 * x <= 2 * 7 - 5 = 9
11. => x = 3