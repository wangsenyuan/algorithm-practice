# P3188 [HNOI2007] 梦幻岛宝珠

[Problem link](https://www.luogu.com.cn/problem/P3188)

time limit per test: 1.00s

memory limit per test: 125.00MB

There are `n` gems. Each gem has a weight and a value. Choose a subset whose total weight is at most `W` and whose total value is as large as possible.

## Input

The input contains several test cases. Each case is:

- one line with two positive integers `n` and `W`, the number of gems and the weight limit
- `n` lines, each with two integers `weight_i` and `value_i`

After the last case there is a line `-1 -1`. That line is a terminator, not a case, and produces no output.

## Output

For each case, print one integer: the maximum total value.

## Sample

### Input

```text
4 10
8 9
5 8
4 6
2 5
4 13
8 9
5 8
4 6
2 5
16 75594681
393216 5533
2 77
32768 467
29360128 407840
112 68
24576 372
768 60
33554432 466099
16384 318
33554432 466090
2048 111
24576 350
9216 216
12582912 174768
16384 295
1024 76
-1 -1
```

### Output

```text
14
19
1050650
```

## Constraints

- at most `20` test cases
- `1 <= n <= 100`
- `1 <= W <= 2^30`
- `1 <= weight_i <= 2^30`, `0 <= value_i <= 2^30`
- each `weight_i` can be written as `a * 2^b` with `1 <= a <= 10` and `0 <= b <= 30`
- the answer does not exceed `2^30`

## Status

I/O and the official sample cases are in place. `solve` is left as a TODO.
